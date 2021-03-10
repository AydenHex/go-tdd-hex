package application

import (
	"github.com/aydenhex/go-tdd-hex/service/shared"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
	"github.com/cockroachdb/errors"
)

type TodoQueryHandler struct {
	retrieveTodoEventStream ForRetrievingTodoEventStreams
}

func NewTodoQueryHandler(retrieveTodoEventStream ForRetrievingTodoEventStreams) *TodoQueryHandler {
	return &TodoQueryHandler{
		retrieveTodoEventStream: retrieveTodoEventStream,
	}
}

func (h *TodoQueryHandler) TodoViewByID(todoID string) (todo.View, error) {
	var err error
	var todoIDValue value.TodoID
	wrapWithMsg := "todoQueryHandler.TodoViewByID"

	if todoIDValue, err = value.BuildTodoID(todoID); err != nil {
		return todo.View{}, errors.Wrap(err, wrapWithMsg)
	}

	eventStream, err := h.retrieveTodoEventStream(todoIDValue)
	if err != nil {
		return todo.View{}, errors.Wrap(err, wrapWithMsg)
	}

	todoView := todo.BuildViewFrom(eventStream)

	if todoView.IsDeleted {
		err := errors.New("todo not found")

		return todo.View{}, shared.MarkAndWrapError(err, shared.ErrNotFound, wrapWithMsg)
	}

	return todoView, nil
}
