package domain

import (
	"github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
)

type TodoAdded struct {
	todoID value.TodoID
	task   value.Task
	isDone bool
	meta   eventsourcing.EventMeta
}

func BuildTodoAdded(
	todoID value.TodoID,
	task value.Task,
	isDone bool,
	streamVersion uint,
) TodoAdded {

	event := TodoAdded{
		todoID: todoID,
		task:   task,
		isDone: isDone,
	}

	event.meta = eventsourcing.BuildEventMeta(event, streamVersion)

	return event
}

func RebuildTodoAdded(
	todoID value.TodoID,
	task value.Task,
	isDone bool,
	meta eventsourcing.EventMeta,
) TodoAdded {

	event := TodoAdded{
		todoID: todoID,
		task:   task,
		isDone: isDone,
		meta:   meta,
	}

	return event
}

func (event TodoAdded) TodoID() value.TodoID {
	return event.todoID
}

func (event TodoAdded) Task() value.Task {
	return event.task
}

func (event TodoAdded) IsDone() bool {
	return event.isDone
}

func (event TodoAdded) Meta() eventsourcing.EventMeta {
	return event.meta
}

func (event TodoAdded) IsFailureEvent() bool {
	return false
}

func (event TodoAdded) FailureReason() error {
	return nil
}
