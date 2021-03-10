package eventsourcing

import (
	"reflect"
	"strings"
	"time"
)

const (
	metaTimestampFormat = time.RFC3339
)

type EventMeta struct {
	eventName     string
	occuredAt     string
	streamVersion uint
}

func BuildEventMeta(
	event DomainEvent,
	streamVersion uint,
) EventMeta {

	eventType := reflect.TypeOf(event).String()
	eventTypePart := strings.Split(eventType, ".")
	eventName := eventTypePart[len(eventTypePart)-1]

	meta := EventMeta{
		eventName:     eventName,
		occuredAt:     time.Now().Format(metaTimestampFormat),
		streamVersion: streamVersion,
	}

	return meta
}

func RebuildEventMeta(
	eventName string,
	occuredAt string,
	streamVersion uint,
) EventMeta {
	return EventMeta{
		eventName:     eventName,
		occuredAt:     occuredAt,
		streamVersion: streamVersion,
	}
}

func (eventMeta EventMeta) EventName() string {
	return eventMeta.eventName
}

func (eventMeta EventMeta) OccuredAt() string {
	return eventMeta.occuredAt
}

func (eventMeta EventMeta) StreamVersion() uint {
	return eventMeta.streamVersion
}
