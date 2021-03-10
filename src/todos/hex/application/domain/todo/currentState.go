package todo

import (
	"github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
)

type currentState struct {
	id                   value.TodoID
	task                 value.Task
	isDone               bool
	isDeleted            bool
	currentStreamVersion uint
}

func buildCurrentStateFrom(eventStream eventsourcing.EventStream) currentState {
	todo := currentState{}

	for _, event := range eventStream {
		switch actualEvent := event.(type) {
		case domain.TodoAdded:
			todo.id = actualEvent.TodoID()
			todo.task = actualEvent.Task()
		}

		todo.currentStreamVersion = event.Meta().StreamVersion()
	}
	return todo
}
