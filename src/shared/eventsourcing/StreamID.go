package eventsourcing

type StreamID struct {
	value string
}

func NewStreamID(from string) StreamID {
	if from == "" {
		panic("newStreamID: emptyInputGiven")
	}

	return StreamID{value: from}
}

func (streamID StreamID) String() string {
	return streamID.value
}
