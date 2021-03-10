package application

import (
	"github.com/aydenhex/go-tdd-hex/service/shared"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
	"github.com/cockroachdb/errors"
)

const maxTodoCommandHandlerRetries = uint8(10)

type TodoCommandHandler struct {
	retrieveTodoEventStream ForRetrievingTodoEventStreams
	startTodoEventStream    ForStartingTodoEventStreams
	appendToTodoEventStream ForAppendingToTodoEventStreams
}

func NewTodoCommandHandler(
	retrieveTodoEventStream ForRetrievingTodoEventStreams,
	startTodoEventStream ForStartingTodoEventStreams,
	appendToTodoEventStream ForAppendingToTodoEventStreams,
) *TodoCommandHandler {

	return &TodoCommandHandler{
		retrieveTodoEventStream: retrieveTodoEventStream,
		startTodoEventStream:    startTodoEventStream,
		appendToTodoEventStream: appendToTodoEventStream,
	}
}

func (h *TodoCommandHandler) AddTodo(task string) (value.TodoID, error) {
	var err error
	var command domain.AddTodo
	wrapWithMsg := "todoCommandHandler.AddTodo"

	taskValue, err := value.BuildTask(task)
	if err != nil {
		return value.TodoID{}, errors.Wrap(err, wrapWithMsg)
	}

	command = domain.BuildAddTodo(
		value.GenerateTodoID(),
		taskValue,
		false,
	)

	doAdd := func() error {
		todoAdded := todo.Add(command)

		if err = h.startTodoEventStream(todoAdded); err != nil {
			return err
		}

		return nil
	}

	if err = shared.RetryOnConcurrencyConflict(doAdd, maxTodoCommandHandlerRetries); err != nil {
		return value.TodoID{}, errors.Wrap(err, wrapWithMsg)
	}

	return command.TodoID(), nil
}
