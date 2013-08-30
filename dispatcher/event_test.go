package dispatcher

import (
	"testing"
)

func TestStoped(t *testing.T) {
	e := NewEvent()

	//a new event must not be stoped
	if e.Stopped() {
		t.Error()
	}
}

func TestStop(t *testing.T) {
	e := NewEvent()
	e.Stop()
	if !e.Stopped() {
		t.Error()
	}
}

func TestSetDispatcher(t *testing.T) {
	d := NewDispatcher()
	e := NewEvent()
	e.SetDispatcher(d)

	if d != e.Dispatcher() {
		t.Error()
	}
}

func TestGetDispatcher(t *testing.T) {
	e := NewEvent()
	if nil != e.Dispatcher() {
		t.Error()
	}
}

func TestGetName(t *testing.T) {
	e := NewEvent()
	if "" != e.Name() {
		t.Error()
	}
}

func TestSetName(t *testing.T) {
	e := NewEvent()
	e.SetName("Foo")
	if "Foo" != e.Name() {
		t.Error()
	}
}

func TestSubject(t *testing.T) {
	e := NewEvent()
	s := new(interface{})
	e.SetSubject(s)
	if e.Subject() != s {
		t.Error()
	}

}
