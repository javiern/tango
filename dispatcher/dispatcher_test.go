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
