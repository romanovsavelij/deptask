package watcher

import (
	"context"
)

// Watcher handles task events.
type Watcher interface {
	OnTaskStart(ctx context.Context, meta MetaInfo)
	OnTaskFinish(ctx context.Context, meta MetaInfo, res Result)
}

type MetaInfo map[any]any

type Watchers []Watcher

var _ Watcher = Watchers{}

func (w Watchers) OnTaskStart(ctx context.Context, meta MetaInfo) {
	for _, watcher := range w {
		watcher.OnTaskStart(ctx, meta)
	}
}

func (w Watchers) OnTaskFinish(ctx context.Context, meta MetaInfo, res Result) {
	for _, watcher := range w {
		watcher.OnTaskFinish(ctx, meta, res)
	}
}
