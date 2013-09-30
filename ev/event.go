package ev

type Event struct {
	dispatcher Dispatcher
	name       string
	subject    interface{}
}

func NewEvent(name string, d Dispatcher) *Event {
	ev := new(Event)
	ev.name = name
	ev.dispatcher = d
	return ev
}

func (ev *Event) SetDispatcher(d Dispatcher) {
	ev.dispatcher = d
}

func (ev *Event) Dispatcher() Dispatcher {
	return ev.dispatcher
}

func (ev *Event) SetName(s string) {
	ev.name = s
}

func (ev *Event) Name() string {
	return ev.name
}

func (ev *Event) SetSubject(sub interface{}) {
	ev.subject = sub
}

func (ev *Event) Subject() interface{} {
	return ev.subject
}
