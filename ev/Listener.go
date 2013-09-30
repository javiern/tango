package ev

type Listener interface {
	Listen(ev *Event)
}

type ListenerFunc func(ev *Event)

// ServeHTTP calls f(w, r).
func (f ListenerFunc) Listen(ev *Event) {
	f(ev)
}
