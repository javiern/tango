package dispatcher

import (
	"testing"
)

type test struct {
	s string
}

func (t *test) Listen(e Event) Event {
	return e
}

func imAListener(e Event) Event {
	return e
}

func TestInitialState(t *testing.T) {
	d := NewDispatcher()
	if len(d.listeners) != 0 {
		t.Error()
	}
}

func TestAddListener(t *testing.T) {
	d := NewDispatcher()
	//	e := NewEvent()
	l1 := new(test)
	//the normal addListener
	d.AddListener("struct.method", l1)
	d.AddListener("function", ListenerFunc(imAListener))
	d.AddListener("closure", ListenerFunc(func(e Event) Event {
		return e
	}))
	if len(d.listeners) != 3 {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}
}

func TestOn(t *testing.T) {
	d := NewDispatcher()
	//	e := NewEvent()
	l1 := new(test)
	//the normal addListener
	d.On("struct.method", l1)
	d.On("function", ListenerFunc(imAListener))
	d.On("closure", ListenerFunc(func(e Event) Event {
		return e
	}))
	if len(d.listeners) != 3 {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}
}

func TestRmListener(t *testing.T) {
	d := NewDispatcher()
	//	e := NewEvent()
	l1 := new(test)
	//the normal addListener
	d.AddListener("foo", l1)
	l2 := d.AddListener("foo", ListenerFunc(imAListener))
	l3 := d.AddListener("foo", ListenerFunc(func(e Event) Event {
		return e
	}))
	if len(d.Listeners("foo")) != 3 {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}

	d.RemoveListener("foo", l1)
	if len(d.Listeners("foo")) != 2 {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}

	d.RemoveListener("foo", l2)
	if len(d.Listeners("foo")) != 1 {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}
	expected := []Listener{l3}

	if !arrayEqual(expected, d.Listeners("foo")) {
		t.Log(len(d.listeners))
		t.Error(d.listeners)
	}

}

func arrayEqual(a []Listener, b []Listener) bool {
	if len(a) == len(b) {
		for i := range a {
			return b[i] == a[i]
		}
	}
	return false
}
