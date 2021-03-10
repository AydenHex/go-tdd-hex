package serialization

import "github.com/aydenhex/go-tdd-hex/src/shared/eventsourcing"

type TodoAddedForJSON struct {
	TodoID string                         `json:"todoID"`
	Task   string                         `json:"task"`
	IsDone bool                           `json:"isDone"`
	Meta   eventsourcing.EventMetaForJSON `json:"meta"`
}
