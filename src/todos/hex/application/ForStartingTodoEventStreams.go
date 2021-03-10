package application

import "github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"

type ForStartingTodoEventStreams func(todoRegistered domain.TodoAdded) error
