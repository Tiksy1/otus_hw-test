package storage

import "context"

type Event struct {
	ID          string
	Title       string
	StartDate   uint64
	EndDate     uint64
	Description string
	OwnerID     string
	RemindIn    int64
}

type EventDataStore interface {
	NewEvent(ctx context.Context, e Event) error
	UpdateEvent(ctx context.Context, e Event) error
	RemoveEvent(ctx context.Context, id string) error
	EventList(ctx context.Context, from uint64, to uint64) ([]Event, error)
}
