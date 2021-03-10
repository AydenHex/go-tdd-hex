package todo

import "github.com/aydenhex/go-tdd-hex/service/shared/eventsourcing"

type View struct {
	ID        string
	Task      string
	IsDone    bool
	IsDeleted bool
	Version   uint
}

func BuildViewFrom(eventStream eventsourcing.EventStream) View {
	todo := buildCurrentStateFrom(eventStream)

	todoView := View{
		ID:        todo.id.String(),
		Task:      todo.task.String(),
		IsDone:    todo.isDone,
		IsDeleted: todo.isDeleted,
		Version:   todo.currentStreamVersion,
	}

	return todoView
}
