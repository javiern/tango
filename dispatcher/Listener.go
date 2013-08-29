package dispatcher

type Listener interface {
	Listen(Event) Event
}
