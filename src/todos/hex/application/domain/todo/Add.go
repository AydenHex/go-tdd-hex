package todo

import "github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"

func Add(with domain.AddTodo) domain.TodoAdded {
	event := domain.BuildTodoAdded(
		with.TodoID(),
		with.Task(),
		with.IsDone(),
		1,
	)

	return event
}
