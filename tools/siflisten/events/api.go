package events

import "context"

type Client interface {
	/*GetCursor the position of named cursor from cursors table */
	GetCursor(ctx context.Context, name string) (int64, error)

	SetCursor(ctx context.Context, name string, position int64) error

	/*CreateEvent creates event in events table. */
	CreateEvent(ctx context.Context, ev *Event) error

	/*GetEvent gets one event after the last processed cursor position. */
	GetEvent(ctx context.Context, cursorPosition int64) (*Event, error)

	/*GetEvents gets all events after the last processed cursor position. */
	GetEvents(ctx context.Context, cursorPosition int64) ([]*Event, error)

	/*ConsumeForever consumes events continuously,
	updating cursor after successfully running consumer. */
	ConsumeForever(ctx context.Context, cursorName string, f Consumer)
}

type Consumer func(context.Context, *Event) error

type Event struct {
	ID         int64  `sql:"id"`
	EventType  string `sql:"type"`
	Height     int32  `sql:"height"`
	Attributes []Attribute
	Metadata   []byte `sql:"metadata"`
}

type Attribute struct {
	Key   string
	Value string
}