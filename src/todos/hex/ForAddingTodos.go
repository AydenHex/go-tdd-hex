package hex

import "github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"

type ForAddingTodos func(task string) (value.TodoID, error)
