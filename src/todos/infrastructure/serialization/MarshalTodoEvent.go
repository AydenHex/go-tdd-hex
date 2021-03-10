package serialization

import (
	"github.com/aydenhex/go-tdd-hex/src/shared"
	"github.com/aydenhex/go-tdd-hex/src/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/src/todos/hex/application/domain"
	"github.com/cockroachdb/errors"
	jsoniter "github.com/json-iterator/go"
)

func MarshalTodoEvent(event eventsourcing.DomainEvent) ([]byte, error) {
	var err error
	var json []byte

	switch actualEvent := event.(type) {
	case domain.TodoAdded:
		json = marshalTodoAdded(actualEvent)
	default:
		err = errors.Wrapf(errors.New("event is unknown"), "marshalTodoEvent [%s] failed", event.Meta().EventName())
		return nil, errors.Mark(err, shared.ErrMarshalingFailed)
	}

	return json, nil
}

func marshalTodoAdded(event domain.TodoAdded) []byte {

	data := TodoAddedForJSON{
		TodoID: event.TodoID(),
		Task:   event.Task(),
		IsDone: event.IsDone(),
		Meta:   marshalEventMeta(event),
	}

	json, _ := jsoniter.ConfigFatest.Marshal(data)
	return json
}
