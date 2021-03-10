package postgres

import (
	"database/sql"
	"math"
	"strings"

	"github.com/aydenhex/go-tdd-hex/service/shared"
	"github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
	"github.com/cockroachdb/errors"
	"github.com/lib/pq"
)

const streamPrefix = "todo"

type TodoEventStore struct {
	db                   *sql.DB
	eventStoreTableName  string
	marshalDomainEvent   eventsourcing.MarshalDomainEvent
	unmarshalDomainEvent eventsourcing.UnmarshalDomainEvent
}

func NewTodoEventStore(
	db *sql.DB,
	eventStoreTableName string,
	marshalDomainEvent eventsourcing.MarshalDomainEvent,
	unmarshalDomainEvent eventsourcing.UnmarshalDomainEvent,
) *TodoEventStore {

	return &TodoEventStore{
		db:                   db,
		eventStoreTableName:  eventStoreTableName,
		marshalDomainEvent:   marshalDomainEvent,
		unmarshalDomainEvent: unmarshalDomainEvent,
	}
}

func (s *TodoEventStore) RetrieveEventStream(id value.TodoID) (eventsourcing.EventStream, error) {
	wrapWithMsg := "todoEventStore.RetrieveEventStream"

	eventStream, err := s.loadEventStream(s.streamID(id), 0, math.MaxUint32)
	if err != nil {
		return nil, errors.Wrap(err, wrapWithMsg)
	}

	if len(eventStream) == 0 {
		err := errors.New("todo not found")
		return nil, shared.MarkAndWrapError(err, shared.ErrNotFound, wrapWithMsg)
	}

	return eventStream, nil
}

func (s *TodoEventStore) StartEventStream(todoAdded domain.TodoAdded) error {
	var err error
	wrapWithMsg := "todoEventStore.StartEventStream"

	tx, err := s.db.Begin()
	if err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	if err = s.appendEventsToStream(tx, s.streamID(todoAdded.TodoID()), todoAdded); err != nil {
		_ = tx.Rollback()

		if errors.Is(err, shared.ErrConcurrencyConflict) {
			return shared.MarkAndWrapError(errors.New("found duplicated todo"), shared.ErrDuplicate, wrapWithMsg)
		}

		return errors.Wrap(err, wrapWithMsg)
	}

	if err = tx.Commit(); err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	return nil

}

func (s *TodoEventStore) AppendToEventStream(recordedEvents eventsourcing.RecordedEvents, id value.TodoID) error {
	var err error
	wrapWithMsg := "todoEventStore.AppendToEventStream"

	tx, err := s.db.Begin()
	if err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	if err = s.appendEventsToStream(tx, s.streamID(id), recordedEvents...); err != nil {
		_ = tx.Rollback()

		return errors.Wrap(err, wrapWithMsg)
	}

	if err = tx.Commit(); err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	return nil
}

func (s *TodoEventStore) PurgeEventStream(id value.TodoID) error {
	var err error
	wrapWithMsg := "todoEventStore.PurgeEventStream"

	tx, err := s.db.Begin()
	if err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	if err = tx.Commit(); err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	if err := s.purgeEventStream(s.streamID(id)); err != nil {
		return errors.Wrap(err, wrapWithMsg)
	}

	return nil
}

func (s *TodoEventStore) streamID(id value.TodoID) eventsourcing.StreamID {
	return eventsourcing.NewStreamID(streamPrefix + "-" + id.String())
}

func (s *TodoEventStore) loadEventStream(
	streamID eventsourcing.StreamID,
	fromVersion uint,
	maxEvents uint,
) (eventsourcing.EventStream, error) {

	var err error
	wrapWithMsg := "loadEventStream"

	queryTemplate := `SELECT event_name, payload, stream_version from %name%
											WHERE stream_id = $1 AND stream_version >= $2
											ORDER BY stream_version ASC
											LIMIT $3`

	query := strings.Replace(queryTemplate, "%name%", s.eventStoreTableName, 1)

	eventRows, err := s.db.Query(query, streamID.String(), fromVersion, maxEvents)
	if err != nil {
		return nil, shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
	}

	var eventStream eventsourcing.EventStream
	var eventName string
	var payload string
	var streamVersion uint
	var domainEvent eventsourcing.DomainEvent

	for eventRows.Next() {
		if eventRows.Err() != nil {
			return nil, shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
		}

		if err = eventRows.Scan(&eventName, &payload, &streamVersion); err != nil {
			return nil, shared.MarkAndWrapError(err, shared.ErrTechnical, wrapWithMsg)
		}

		if domainEvent, err = s.unmarshalDomainEvent(eventName, []byte(payload), streamVersion); err != nil {
			return nil, shared.MarkAndWrapError(err, shared.ErrUnmarshalingFailed, wrapWithMsg)
		}

		eventStream = append(eventStream, domainEvent)
	}

	return eventStream, nil
}

func (s *TodoEventStore) appendEventsToStream(
	tx *sql.Tx,
	streamID eventsourcing.StreamID,
	events ...eventsourcing.DomainEvent,
) error {
	var err error
	wrapWithMsg := "appendEventsToStream"

	queryTemplate := `INSERT INTO %name% (stream_id, stream_version, event_name, occured_at, payload)
											VALUES ($1, $2, $3, $4, $5)`
	query := strings.Replace(queryTemplate, "%name%", s.eventStoreTableName, 1)

	for _, event := range events {
		var eventJson []byte

		eventJson, err = s.marshalDomainEvent(event)
		if err != nil {
			return shared.MarkAndWrapError(err, shared.ErrMarshalingFailed, wrapWithMsg)
		}

		_, err := tx.Exec(
			query,
			streamID.String(),
			event.Meta().StreamVersion(),
			event.Meta().OccuredAt(),
			eventJson,
		)

		if err != nil {
			return errors.Wrap(s.mapEventStorePostgresErrors(err), wrapWithMsg)
		}
	}

	return nil
}

func (s *TodoEventStore) purgeEventStream(streamID eventsourcing.StreamID) error {
	queryTemplate := `DELETE FROM %name% WHERE stream_id = $1`
	query := strings.Replace(queryTemplate, "%name%", s.eventStoreTableName, 1)

	if _, err := s.db.Exec(query, streamID.String()); err != nil {
		return shared.MarkAndWrapError(err, shared.ErrTechnical, "purgeEventStream")
	}

	return nil
}

func (s *TodoEventStore) mapEventStorePostgresErrors(err error) error {
	switch actualErr := err.(type) {
	case *pq.Error:
		switch actualErr.Code {
		case "23505":
			return errors.Mark(err, shared.ErrConcurrencyConflict)
		}
	}

	return errors.Mark(err, shared.ErrTechnical)
}
