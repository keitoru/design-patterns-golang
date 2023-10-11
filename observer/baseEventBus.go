package observer

import "sync"

type BaseEventBus struct {
	mux       sync.Mutex
	observers map[string]map[Observer]struct{}
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()

	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}

	b.observers[topic][o] = struct{}{}
}

func (b *BaseEventBus) UnSubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()

	delete(b.observers[topic], o)
}
