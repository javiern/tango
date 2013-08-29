package dispatcher

type Subscriber interface {
	Events() map[string]Event
}
