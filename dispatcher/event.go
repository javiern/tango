package dispatcher

type Event interface {
	Stopped() bool
	Stop()
	Dispatcher() Dispatcher
	SetDispatcher(d Dispatcher)
	Name() string
	SetName(s string)
	Subject() interface{}
	SetSubject(s interface{})
}

type event struct {
	name       string
	dispatcher Dispatcher
	stopped    bool
	subject    interface{}
}

func NewEvent() *event {
	e := new(event)
	return e
}

func (e *event) Stopped() bool {
	return e.stopped
}

func (e *event) Stop() {
	e.stopped = true
}

func (e *event) Dispatcher() Dispatcher {
	return e.dispatcher
}

func (e *event) SetDispatcher(d Dispatcher) {
	e.dispatcher = d
}

func (e *event) Name() string {
	return e.name
}

func (e *event) String() string {
	return e.Name()
}

func (e *event) SetName(s string) {
	e.name = s
}

func (e *event) Subject() interface{} {
	return e.subject
}

func (e *event) SetSubject(s interface{}) {
	e.subject = s
}
