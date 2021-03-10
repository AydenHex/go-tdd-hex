package application

import (
	"github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
)

type ForAppendingToTodoEventStreams func(recordedEvents eventsourcing.RecordedEvents, id value.TodoID) error
