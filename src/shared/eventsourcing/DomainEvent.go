package eventsourcing

type DomainEvent interface {
	Meta() EventMeta
	IsFailureEvent() bool
	FailureReason() error
}
