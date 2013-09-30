package ev

import (
	"testing"
)

type TestListener struct{}

func (tl *TestListener) Listen(ev *Event) {}

func setUp() (d Dispatcher) {
	d = NewDispatcher()
	l := new(TestListener)
	d.AddListener("test.l.add", l)
	d.AddListenerOnce("test.l.once", l)
	d.On("test.on", func(ev *Event) {})
	d.Once("test.once", func(ev *Event) {})
	d.AddFunc("test.func", func(ev *Event) {})
	d.AddFuncOnce("test.func.once", func(ev *Event) {})

	return
}

func checkListeners(d Dispatcher, t *testing.T, has bool) {
	if d.HasListeners("test.l.add") == has {
		t.Error("test.l.add")
	}

	if d.HasListeners("test.l.once") == has {
		t.Error("test.l.once")
	}

	if d.HasListeners("test.on") == has {
		t.Error("test.on")
	}

	if d.HasListeners("test.once") == has {
		t.Error("test.once")
	}

	if d.HasListeners("test.func") == has {
		t.Error("test.func")
	}

	if d.HasListeners("test.func.once") == has {
		t.Error("test.func.once")
	}
}

func TestAddListener(t *testing.T) {
	d := setUp()                //agrega listeners
	checkListeners(d, t, false) //si no tiene da error
}

func TestRemoveAllListener(t *testing.T) {
	d := setUp()               //agrega listeners
	d.ClearListeners()         // borra todos los listeners
	checkListeners(d, t, true) // si tiene da error
}

func TestRemoveAListener(t *testing.T) {
	d := setUp()
	l := new(TestListener)

	if d.HasListeners("Listener") {
		t.Error("Initial state failed, it must not have a listener here")
	}

	d.AddListener("Listener", l)

	if !d.HasListeners("Listener") {
		t.Error("it must have a listener")
	}

	d.RemoveListener("Listener", l)

	if d.HasListeners("Listener") {
		t.Error("Remove listener failed", d.Listeners("Listener"), d.HasListeners("Listener"), d)
	}
}

func TestRemoveAllListenersFromEvent(t *testing.T) {
	d := setUp()
	l := new(TestListener)
	d.AddListener("test.l.add", l)
	d.RemoveListeners("test.l.add")
	if d.HasListeners("test.l.add") {
		t.Error("Remove listener failed", d.Listeners("test.l.add"), d.HasListeners("test.l.add"), d)
	}
}

func TestRemoveClosureFromEvent(t *testing.T) {
	d := setUp()

	//for removal of a listener, i need a reference to ir
	//this is my closure definition, i save it in a variable
	l := ListenerFunc(func(ev *Event) {})
	//l := func(ev *Event) {}

	//check not having listeners
	if d.HasListeners("closure") {
		t.Error("Initial state failed, it must not have a listener here")
	}

	//i add the closure via the add listener
	//this will give me back a reference for the added event, it can be used later
	//to remove it
	//it works with on, and once as well
	ref := d.AddListener("closure", l)

	// i check if i have the listener registered
	if !d.HasListeners("closure") {
		t.Error("It must have a listener at this poit")
	}

	//remove the listener
	d.RemoveListener("closure", ref)

	if d.HasListeners("closure") {
		t.Error("Remove listener failed", d.Listeners("closure"), d.HasListeners("closure"), d)
	}
}
