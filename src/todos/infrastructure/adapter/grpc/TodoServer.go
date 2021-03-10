package todogrpc

import (
	"context"

	"github.com/aydenhex/go-tdd-hex/service/todos/hex"
)

type todoServer struct {
	add hex.ForAddingTodos
}

func NewTodoServer(
	add hex.ForAddingTodos,
) *todoServer {
	server := &todoServer{
		add: add,
	}

	return server
}

func (server *todoServer) Add(
	_ context.Context,
	req *AddRequest,
) (*AddResponse, error) {

	todoId, err := server.add(req.task)
	if err != nil {
		return nil, MapToGRPCErrors(err)
	}
	

}
