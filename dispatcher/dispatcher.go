package dispatcher

type Listener interface {
	Listen(Event) Event
}

type Dispatcher interface {
	Dispatch(name string, e Event) Event
	On(name string, e Event) Event
	Once(name string, e Event) Event
	AddListener(name string, listener callable, priority int)
	RemoveListener(name string, listener callable)
	Listeners(name string) map[int][]callable
	HasListeners(name string) bool
}

type dispatcher struct{}

func (d *dispatcher) Dispatch(name string, e Event) Event                      {}
func (d *dispatcher) AddListener(name string, listener callable, priority int) {}
func (d *dispatcher) RemoveListener(name string, l callable)                   {}
func (d *dispatcher) Listeners(name string) map[int][]callable                 {}
func (d *dispatcher) HasListeners(name string) bool                            {}
func (d *dispatcher) doDispatch(listeners map[int][]callable, e Event)         {}
