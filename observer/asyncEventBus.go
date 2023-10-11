package observer

import (
	"context"
	"fmt"
)

type observerWithErr struct {
	o   Observer
	err error
}

type AsyncEventBus struct {
	BaseEventBus
	errC chan *observerWithErr
	ctx  context.Context
	stop context.CancelFunc
}

func NewAsyncEventBus() *AsyncEventBus {
	a := AsyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}

	a.ctx, a.stop = context.WithCancel(context.Background())

	// 处理处理错误的异步守护协程
	go a.handleErr()

	return &a
}

func (a *AsyncEventBus) Stop() {
	a.stop()
}

func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
	a.mux.Lock()
	subscribers := a.observers[e.Topic]
	a.mux.Unlock()

	for s := range subscribers {
		s := s
		go func() {
			if err := s.OnChange(ctx, e); err != nil {
				select {
				case <-a.ctx.Done():
				case a.errC <- &observerWithErr{
					o:   s,
					err: err,
				}:
				}
			}
		}()
	}
}

func (a *AsyncEventBus) handleErr() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case resp := <-a.errC:
			// 处理 publish 失败的 observer
			fmt.Printf("observer: %v, err: %v", resp.o, resp.err)
		}
	}
}
