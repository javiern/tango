package ev

type ListenersMap map[string][]Listener

type Dispatcher interface {
	Dispatch(name string) *Event
	DispatchEvent(name string, ev *Event) *Event
	AddListener(name string, listener Listener) Listener
	AddListenerOnce(name string, listener Listener) Listener
	AddFunc(name string, listener func(ev *Event)) Listener
	AddFuncOnce(name string, listener func(ev *Event)) Listener
	On(name string, listener func(ev *Event)) Listener
	Once(name string, listener func(ev *Event)) Listener
	RemoveListener(name string, listener Listener)
	RemoveListeners(name string)
	ClearListeners()
	Listeners(name string) []Listener
	AllListeners() (ListenersMap, ListenersMap)
	HasListeners(name string) bool
}

type EventDispatcher struct {
	listeners     ListenersMap
	listenersOnce ListenersMap
}

func NewDispatcher() *EventDispatcher {
	d := new(EventDispatcher)
	d.listeners = make(ListenersMap)
	d.listenersOnce = make(ListenersMap)
	return d
}

func (d *EventDispatcher) Dispatch(name string) *Event {
	ev := NewEvent(name, d)
	return d.DispatchEvent(name, ev)
}

func (d *EventDispatcher) DispatchEvent(name string, ev *Event) *Event {
	ev.SetDispatcher(d)
	ev.SetName(name)

	//reviso si esta registrado el evento
	if d.HasListeners(name) {
		d.doDispatch(d.Listeners(name), ev)
	}

	delete(d.listenersOnce, name)

	return ev
}

func (d *EventDispatcher) doDispatch(listeners []Listener, ev *Event) {
	for _, l := range listeners {
		l.Listen(ev)
	}
}

func (d *EventDispatcher) AddListener(name string, listener Listener) Listener {
	switch t := listener.(type) {
	default:
		d.listeners[name] = append(d.listeners[name], t)
		return t
	case ListenerFunc:
		d.listeners[name] = append(d.listeners[name], &t)
		return &t
	}
}

func (d *EventDispatcher) AddFunc(name string, listener func(ev *Event)) Listener {
	return d.AddListener(name, ListenerFunc(listener))
}

func (d *EventDispatcher) On(name string, listener func(ev *Event)) Listener {
	return d.AddListener(name, ListenerFunc(listener))
}

func (d *EventDispatcher) AddListenerOnce(name string, listener Listener) Listener {
	switch t := listener.(type) {
	default:
		d.listenersOnce[name] = append(d.listenersOnce[name], t)
		return t
	case ListenerFunc:
		d.listenersOnce[name] = append(d.listenersOnce[name], &t)
		return &t
	}
}

func (d *EventDispatcher) AddFuncOnce(name string, listener func(ev *Event)) Listener {
	return d.AddListenerOnce(name, ListenerFunc(listener))
}

func (d *EventDispatcher) Once(name string, listener func(ev *Event)) Listener {
	return d.AddListenerOnce(name, ListenerFunc(listener))
}

func (d *EventDispatcher) RemoveListener(name string, listener Listener) {
	for k, l := range d.listeners[name] {
		if l == listener {
			d.listeners[name] = append(d.listeners[name][:k], d.listeners[name][k+1:]...)
		}
	}
	for k1, l1 := range d.listenersOnce[name] {
		if l1 == listener {
			d.listenersOnce[name] = append(d.listenersOnce[name][:k1], d.listenersOnce[name][k1+1:]...)
		}
	}

	//if the array is empty (len(s) == 0), remove the key from the map.
	if len(d.listeners[name]) == 0 {
		delete(d.listeners, name)
	}

	if len(d.listenersOnce[name]) == 0 {
		delete(d.listenersOnce, name)
	}

}

func (d *EventDispatcher) RemoveListeners(name string) {
	delete(d.listeners, name)
	delete(d.listenersOnce, name)
}

func (d *EventDispatcher) ClearListeners() {
	d.listeners = make(ListenersMap)
	d.listenersOnce = make(ListenersMap)
}

func (d *EventDispatcher) Listeners(name string) []Listener {
	return append(d.listeners[name], d.listenersOnce[name]...)
}

func (d *EventDispatcher) AllListeners() (ListenersMap, ListenersMap) {
	return d.listeners, d.listenersOnce
}

func (d *EventDispatcher) HasListeners(name string) bool {
	_, ok := d.listeners[name]
	_, okOnce := d.listenersOnce[name]
	return ok || okOnce
}
