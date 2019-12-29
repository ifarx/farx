package event

import (
	"src/instruct"
	"time"
)

//参考HTML标准

//暂时未想到是否需要事件冒泡

type EventType string

type EventError error

type Event interface {
	Type() EventType
	Target() EventTarget
	Prevent() EventError
	Timestamp() time.Time
}

type EventListener interface {
	HandleEvent(event Event) EventError
}

type EventTarget interface {
	DispatchEvent(event Event) EventError
	AddEventListener(typ EventType, listener EventListener) EventError
	RemoveEventListener(typ EventType, listener EventListener) EventError
}

type FuncEventListener struct {
	_func func(event Event) EventError
}

func (f *FuncEventListener) HandleEvent(event Event) EventError {
	return f._func(event)
}

func NewFuncEventListener(f func(event Event) EventError) EventListener {
	return &FuncEventListener{_func: f}
}

type EventListeners struct {
	_els []EventListener
}

func (ls *EventListeners) AddEventListener(el EventListener) bool {
	size := len(ls._els)
	for index := 0; index < size; index++ {
		if el == ls._els[index] {
			return true
		}
	}
	ls._els = append(ls._els, el)
	return true
}

func (ls *EventListeners) RemoveEventListener(el EventListener) bool {
	var index int = 0
	size := len(ls._els)
	if size <= 0 {
		return false
	}
	if el == ls._els[0] {
		ls._els = ls._els[:0]
		return true
	}
	if el == ls._els[size-1] {
		ls._els = ls._els[:size-1]
		return true
	}
	for ; index < size; index++ {
		if el == ls._els[index] {
			ls._els = append(ls._els[:index], ls._els[index+1:]...)
			return true
		}
	}
	return false
}

func (ls *EventListeners) ForeachEvent(event Event) EventError {
	size := len(ls._els)
	var err EventError = nil
	for index := 0; index < size; index++ {
		err = ls._els[index].HandleEvent(event)
		if err != nil {
			return err
		}
	}
	return err
}

func NewEventListeners() *EventListeners {
	return &EventListeners{_els: make([]EventListener, 0, 10)}
}

/**
 */
type EventTargetAgent struct {
	_lamp map[EventType]*EventListeners
}

func (elt *EventTargetAgent) DispatchEvent(event Event) EventError {
	els := elt._lamp[event.Type()]
	if els == nil {
		return nil
	}
	return els.ForeachEvent(event)
}

func (elt *EventTargetAgent) AddEventListener(typ EventType, listener EventListener) EventError {
	els := elt._lamp[typ]
	if els == nil {
		els = NewEventListeners()
		elt._lamp[typ] = els
	}
	add := els.AddEventListener(listener)
	if !add {
		return instruct.NewErrcode(-99, "add event listener not ok")
	}
	return nil
}

func (elt *EventTargetAgent) RemoveEventListener(typ EventType, listener EventListener) EventError {
	els := elt._lamp[typ]
	if els != nil {
		remove := els.RemoveEventListener(listener)
		if !remove {
			return instruct.NewErrcode(-100, "remove event listener not ok")
		}
	}
	return nil
}

func NewEventTargetAgent() *EventTargetAgent {
	return &EventTargetAgent{_lamp: make(map[EventType]*EventListeners)}
}
