package value

import (
	"github.com/aydenhex/go-tdd-hex/service/shared"
	"github.com/cockroachdb/errors"
)

type Task struct {
	value string
}

func BuildTask(input string) (Task, error) {
	if input == "" {
		err := errors.New("empty input for Task")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, "BuildTask")

		return Task{}, err
	}

	task := Task{value: input}

	return task, nil
}

func RebuildTask(input string) Task {
	return Task{value: input}
}

func (task Task) String() string {
	return task.value
}

func (task Task) Equals(other Task) bool {
	return task.value == other.value
}
