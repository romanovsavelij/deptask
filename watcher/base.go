package watcher

import (
	"context"
)

type baseWatcher struct{}

var _ Watcher = baseWatcher{}

func (b baseWatcher) OnTaskStart(ctx context.Context, meta MetaInfo)              {}
func (b baseWatcher) OnTaskFinish(ctx context.Context, meta MetaInfo, res Result) {}
