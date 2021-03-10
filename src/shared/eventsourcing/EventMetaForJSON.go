package eventsourcing

type EventMetaForJSON struct {
	EventName   string `json:"eventName"`
	OccuredAt   string `json:"occuredAt"`
	MessageID   string `json:"messageID"`
	CausationID string `json:"causationID"`
}
