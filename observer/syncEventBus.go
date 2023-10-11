package observer

import (
	"context"
	"fmt"
)

type SyncEventBus struct {
	BaseEventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
}

// 通知订阅了该topic的观察者
func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.Lock()
	subscribers := s.observers[e.Topic]
	s.mux.Unlock()

	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}

	s.HandleErr(ctx, errs)

}

// 处理 publish 失败的 observer
func (s *SyncEventBus) HandleErr(ctx context.Context, errs map[Observer]error) {
	for o, err := range errs {
		fmt.Printf("observer: %v, err: %v\n", o, err)
	}
}
