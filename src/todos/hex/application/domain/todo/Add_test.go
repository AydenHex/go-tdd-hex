package todo_test

import (
	"testing"

	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo"
	"github.com/aydenhex/go-tdd-hex/service/todos/hex/application/domain/todo/value"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegister(t *testing.T) {
	Convey("Prepare test artifact", t, func() {
		var err error

		id := value.GenerateTodoID()

		task, err := value.BuildTask("Cooking")
		So(err, ShouldBeNil)

		add := domain.BuildAddTodo(
			id,
			task,
			false,
		)

		Convey("\nSCENARIO: Add a Todo", func() {
			Convey("When AddTodo", func() {
				added := todo.Add(add)

				Convey("Then TodoAdded", func() {
					So(added.TodoID().Equals(add.TodoID()), ShouldBeTrue)
					So(added.Task().Equals(add.Task()), ShouldEqual)
					So(added.IsDone() == add.IsDone(), ShouldBeTrue)
					So(added.IsFailureEvent(), ShouldBeFalse)
					So(added.FailureReason(), ShouldBeNil)
					So(added.Meta().StreamVersion(), ShouldEqual, uint(1))
				})
			})
		})
	})
}
