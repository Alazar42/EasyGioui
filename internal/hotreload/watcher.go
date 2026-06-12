package hotreload

import (
	"context"
	"time"
)

type Watcher struct {
	Interval time.Duration
}

func New() *Watcher {
	return &Watcher{Interval: 500 * time.Millisecond}
}

func (w *Watcher) Run(ctx context.Context, onTick func() error) error {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := onTick(); err != nil {
				return err
			}
		}
	}
}
