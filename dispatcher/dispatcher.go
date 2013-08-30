package dispatcher

type Dispatcher interface {
	Dispatch(name string, e Event) Event
	AddListener(name string, listener Listener) Listener
	On(name string, listener Listener) Listener
	//Once(name string, listener Listener) Listener
	RemoveListener(name string, listener Listener)
	Listeners(name string) []Listener
	HasListeners(name string) bool
}

type dispatcher struct {
	listeners map[string][]Listener
}

func NewDispatcher() *dispatcher {
	d := new(dispatcher)
	d.listeners = make(map[string][]Listener)
	return d
}

func (d *dispatcher) Dispatch(name string, e Event) Event {
	e.SetName(name)
	e.SetDispatcher(d)
	if d.HasListeners(name) {
		listeners := d.Listeners(name)
		for _, l := range listeners {
			e = l.Listen(e)
		}
	}
	return e
}

func (d *dispatcher) AddListener(name string, listener Listener) Listener {
	e := NewEvent()
	e.SetSubject(listener)
	d.listeners[name] = append(d.listeners[name], listener)
	return listener
}
func (d *dispatcher) On(name string, listener Listener) Listener {
	return d.AddListener(name, listener)
}

//func (d *dispatcher) Once(name string, listener Listener) Listener {
//@TODO find a way to remove the listener after dispatched
//	return d.AddListener(name, listener)
//}
func (d *dispatcher) RemoveListener(name string, listener Listener) {
	//@TODO this way can get really slow in case of much events gets registered
	//and deleted, i need to find a better way to do this
	//one option, is to turn the []Listener into map[Listener]bool, but
	//it will give me problems returning an iterable of Listeners
	for i, v := range d.Listeners(name) {
		if v == listener {
			copy(d.listeners[name][i:], d.listeners[name][i+1:])
			d.listeners[name][len(d.listeners[name])-1] = nil // or the zero value of T
			d.listeners[name] = d.listeners[name][:len(d.listeners[name])-1]
		}
	}
}
func (d *dispatcher) Listeners(name string) []Listener {
	return d.listeners[name]
}
func (d *dispatcher) HasListeners(name string) bool {
	if len(d.listeners) > 0 {
		return true
	} else {
		return false
	}
}
