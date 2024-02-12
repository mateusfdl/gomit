package gomit

import (
	"log"
	"sync"
)

type HandlerCallback[T any] func(data ...any) error

type EventMap struct {
	lock            sync.RWMutex
	listeners       map[string][]HandlerCallback[any]
	ActiveListeners int
}

var listenersMap *EventMap

func AddListener(name string, h HandlerCallback[any]) {
        // We lock since we are modifying the map
	listenersMap.lock.Lock()
	defer listenersMap.lock.Unlock()

	if ls := listenersMap.listeners; ls == nil {
		listenersMap.listeners = make(map[string][]HandlerCallback[any])
		listenersMap.listeners[name] = make([]HandlerCallback[any], 0)
	}

	if ok := listenersMap.listeners[name]; ok == nil {
		listenersMap.listeners[name] = make([]HandlerCallback[any], 1)
		listenersMap.ActiveListeners++
	}

	listenersMap.listeners[name] = append(listenersMap.listeners[name], h)
}

func Listeners() map[string][]HandlerCallback[any] {
	return listenersMap.listeners
}

func Emit(en string, data ...any) {
	handlers := listenersMap.listeners[en]

	for _, v := range handlers {
		go func(f HandlerCallback[any]) {
			err := f()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}(v)
	}
}

func init() {
	listenersMap = &EventMap{
		ActiveListeners: 0,
	}
}
