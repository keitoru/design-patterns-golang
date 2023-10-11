package observer

import "context"

// 事件
type Event struct {
	Topic string
	Val   interface{}
}

// 观察者
type Observer interface {
	OnChange(ctx context.Context, e *Event) error
}

// 事件总线
type EventBus interface {
	Subscribe(topic string, o Observer)
	UnSubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}
