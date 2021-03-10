package application

import (
	"github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
)

type ForRetrievingTodoEventStreams func(id value.TodoID) (eventsourcing.EventStream, error)
