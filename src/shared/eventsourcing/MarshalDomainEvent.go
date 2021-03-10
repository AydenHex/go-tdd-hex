package eventsourcing

type MarshalDomainEvent func(event DomainEvent) ([]byte, error)
