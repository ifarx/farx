package event

import (
	"fmt"
	"src/event"
	"testing"
)

func TestEvent(t *testing.T) {
	//el := &EEventTarget{}
	//el.AddEventListener()
	el := event.NewFuncEventListener(func(e event.Event) event.EventError { return nil })
	el1 := event.NewFuncEventListener(func(e event.Event) event.EventError { return nil })
	el2 := event.NewFuncEventListener(func(e event.Event) event.EventError { return nil })
	el3 := event.NewFuncEventListener(func(e event.Event) event.EventError { return nil })
	els := &event.EventListeners{}
	els.AddEventListener(el)
	els.AddEventListener(el1)
	els.AddEventListener(el2)
	els.AddEventListener(el3)
	els.RemoveEventListener(el3)
	fmt.Println("fasdas")
}

func TestEvent1(t *testing.T) {
	//et := event.NewSimpleEventTarget()
	//el := event.NewFuncEventListener(func(e event.Event) event.EventError {
	//	return nil
	//})
	//et.AddEventListener("t", el)
	//et.DispatchEvent(nil)
}
