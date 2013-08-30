package dispatcher

type Listener interface {
	Listen(Event) Event
}

type ListenerFunc func(Event) Event

func (l ListenerFunc) Listen(e Event) Event {
	return l(e)
}
