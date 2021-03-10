package value

import (
	"github.com/aydenhex/go-tdd-hex/service/shared"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type TodoID struct {
	value string
}

func GenerateTodoID() TodoID {
	return TodoID{value: uuid.New().String()}
}

func BuildTodoID(value string) (TodoID, error) {
	if value == "" {
		err := errors.New("empty input for TodoID")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, "BuildTodoID")

		return TodoID{}, err
	}

	id := TodoID{value: value}

	return id, nil
}

func RebuildTodoID(value string) TodoID {
	return TodoID{value: value}
}

func (id TodoID) String() string {
	return id.value
}

func (id TodoID) Equals(other TodoID) bool {
	return id.String() == other.String()
}
