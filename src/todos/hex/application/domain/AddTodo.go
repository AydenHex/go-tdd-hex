package domain

import "github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"

type AddTodo struct {
	todoID value.TodoID
	task   value.Task
	isDone bool
}

func BuildAddTodo(
	todoID value.TodoID,
	task value.Task,
	isDone bool) AddTodo {

	add := AddTodo{
		todoID: todoID,
		task:   task,
		isDone: isDone,
	}

	return add
}

func (command AddTodo) TodoID() value.TodoID {
	return command.todoID
}

func (command AddTodo) Task() value.Task {
	return command.task
}

func (command AddTodo) IsDone() bool {
	return command.isDone
}
