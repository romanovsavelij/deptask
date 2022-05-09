package task

import (
	"context"

	"Pipelining/watcher"
)

// Waitable wait for task to finish executing.
type Waitable interface {
	Wait(ctx context.Context) watcher.Result
}

type RunFunc func(ctx context.Context) watcher.Result

type Task struct {
	f        RunFunc
	watchers watcher.Watchers
	meta     watcher.MetaInfo

	Children []*Task
}

var _ Waitable = Task{}

func NewTask(f RunFunc) Task {
	return Task{
		f:    f,
		meta: make(watcher.MetaInfo),
	}
}

func NewDefaultTask(f RunFunc, name string) Task {
	return NewTask(f).
		WithName(name).
		WithWatcher(watcher.NewLogger()).
		WithWatcher(watcher.NewTimer())
}

func (t Task) Wait(ctx context.Context) watcher.Result {
	t.watchers.OnTaskStart(ctx, t.meta)
	res := t.f(ctx)
	t.watchers.OnTaskFinish(ctx, t.meta, res)

	return res
}
