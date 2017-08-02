package event

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// Key .
// type Key struct {
// 	Type string
// 	// If Name == "":
// 	//     during listener register it means registered at type (subscribe all event with the same type).
// 	//     during event pubish it means that only listener who registered at type can received it.
// 	Name string
// }

// Key .
// Now the structure of `Key` changed to `[2]string{Type, Name}` .
type Key [2]string

func (k *Key) String() string {
	return fmt.Sprintf(`{"type": "%s", "name": "%s"}`, k[0], k[1])
}

// Event .
type Event struct {
	Key
	Data interface{}
}

func (e *Event) String() string {
	bytes, err := json.Marshal(e.Data)
	if err != nil {
		return fmt.Sprintf(`{"key": %s, "data": %v}`, e.Key.String(), e.Data)
	}
	return fmt.Sprintf(`{"key": %s, "data": %s}`, e.Key.String(), string(bytes))
}

// Send to the given pool.
func (e *Event) Send(p *Pool) { p.Publish(e) }

// Handler .
type Handler func(Event) error

// Listener .
type Listener struct {
	ID string
	Key
	Handler
}

// Bind on the given pool.
func (l *Listener) Bind(p *Pool) { p.Register(l) }

// Unbind from the given pool.
func (l *Listener) Unbind(p *Pool) { p.Remove(l) }

// Pool .
type Pool struct {
	// {"type": {"id": handler}}
	types map[string]map[string]Handler
	// {"type": {"name": {"id": handler}}}
	names map[string]map[string]map[string]Handler

	// ErrHandler would be called when a listener's handler executed and return an error.
	// The param `e` is the event, `id` is the `ID` field of the failed listener.
	ErrHandler func(e Event, id string, err error)
	// NilHandler would be called if the key has not found when `emit` or `on` executed.
	// You can implement it and return a bool value to decide if the key can be permit.
	NilHandler func(Key) bool

	wg sync.WaitGroup
}

// Wait for all handlers finished.
func (p *Pool) Wait() { p.wg.Wait() }

// NewPool create an event pool, you should manage your event keys manual.
// Don't create event keys repeatedly and infinitely, witch may increase the pressure of GC.
// In deployment environment, you should override the `NilHandler` check the legally keys.
// Also, I provide `NewRestrictPool` to make it more convenient to manual keys.
func NewPool() *Pool {
	// TODO
	return &Pool{
		types: map[string]map[string]Handler{},
		names: map[string]map[string]map[string]Handler{},
		ErrHandler: func(e Event, id string, err error) {
			log.Printf(`handler execute failed, event: %s, listener id: %s .\n`, e, id)
			log.Println(err)
		},
		NilHandler: func(Key) bool { return true },
		wg:         sync.WaitGroup{},
	}
}

// NewRestrictPool create an event pool with init keys, and would reject any new keys insert.
// Although looks a bit complexity, this function is more recommend to use.
// You should best to maintain your keys to emulates. for example:
/*
var eks = []event.Key{
	event.Key{"SYS", "start"},
	event.Key{"SYS", "stop"},
	...
}
type EKeyIndex int
const (
	EKeyStart EKeyIndex = iota
	EKeyStop
	...
)
*/
func NewRestrictPool(keys ...Key) *Pool {
	p := NewPool()
	for _, key := range keys {
		p.find(key)
	}
	p.NilHandler = func(k Key) bool { return false }
	return p
}

func (p *Pool) find(k Key) map[string]Handler {
	var handlers map[string]Handler

	if k[1] == "" {
		var ok bool
		handlers, ok = p.types[k[0]]
		if !ok {
			if !p.NilHandler(k) {
				return nil
			}
			handlers = map[string]Handler{}
			p.types[k[0]] = handlers
		}
	} else {
		handlersMap, ok := p.names[k[0]]
		if !ok {
			if !p.NilHandler(k) {
				return nil
			}
			handlersMap = map[string]map[string]Handler{}
			p.names[k[0]] = handlersMap
		}
		handlers, ok = handlersMap[k[1]]
		if !ok {
			if !p.NilHandler(k) {
				return nil
			}
			handlers = map[string]Handler{}
			handlersMap[k[1]] = handlers
		}
	}

	return handlers
}

// Publish an event.
func (p *Pool) Publish(e *Event) {
	// If name == "", publish an event on type at the same time.
	if e.Key[1] != "" {
		p.Publish(&Event{Key{e.Key[0], ""}, e.Data})
	}

	handlers := p.find(e.Key)
	if handlers == nil {
		return
	}

	p.wg.Add(len(handlers))

	for id, handler := range handlers {
		go func(id string, handler Handler) {
			defer p.wg.Done()
			err := handler(*e)
			if err != nil {
				p.ErrHandler(*e, id, err)
			}
		}(id, handler)
	}

	p.wg.Wait()
}

// Emit : send some data with the given key.
func (p *Pool) Emit(k Key, d interface{}) *Event {
	e := &Event{k, d}
	p.Publish(e)
	return e
}

// Register a listener.
func (p *Pool) Register(l *Listener) {
	handlers := p.find(l.Key)
	if handlers == nil {
		return
	}
	// While `ID` is "", it will auto generate an id for listener.
	// However, notice that's an inner implement and may be changed.
	// You should not use this function to register a listener without certain unique id.
	// If you don't want to manage `ID` manual, you can use `On` function.
	if l.ID == "" {
		i := len(handlers)
		for l.ID == "" {
			id := fmt.Sprintf("listener%02d", i)
			i++
			if _, ok := (handlers)[id]; ok {
				continue
			}
			l.ID = id
		}
	}
	handlers[l.ID] = l.Handler
}

// On : bind a handler on the given key.
func (p *Pool) On(k Key, h Handler) *Listener {
	l := &Listener{"", k, h}
	p.Register(l)
	return l
}

// Remove a listener.
func (p *Pool) Remove(l *Listener) {
	p.Off(l.Key, l.ID)
}

// Off : unbind a handler from the given key.
func (p *Pool) Off(k Key, id string) *Listener {
	handlers := p.find(k)
	if handlers == nil {
		return nil
	}
	l := &Listener{id, k, handlers[id]}
	delete(handlers, id)
	return l
}
