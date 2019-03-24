// +build !wasm

package vugu

// NOTE: Dummy implementation allows code that does JS things to compile in non-wasm,
// even though that code will never get called - needed so we can do server side rendering of components.

var DOMEventStub = &DOMEvent{} // used as a value to mean "replace this with the actual event that came in"

type DOMEvent struct {
}

func (e *DOMEvent) RequestRender() {
}

func (e *DOMEvent) PreventDefault() {

}

func (e *DOMEvent) EventEnv() EventEnv {
	return nil // do we need to do anything with this? in static html rendering the events are never called... right?
}
