package task

import (
	"Pipelining/watcher"
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// WaitGroup creates new a task that waits for all the tasks to finish.
// error is always nil.
func WaitGroup(tasks ...Task) Task {
	return NewTask(func(ctx context.Context) watcher.Result {
		wg := sync.WaitGroup{}
		wg.Add(len(tasks))

		for _, task := range tasks {
			go func() {
				defer wg.Done()

				task.Wait(ctx)
			}()
		}
		wg.Wait()

		return watcher.NewResult(nil, nil)
	})
}

// ErrGroup waits for all the tasks to finish or until the first error.
func ErrGroup(tasks ...Task) Task {
	return NewTask(func(ctx context.Context) watcher.Result {
		errGroup, ctx := errgroup.WithContext(ctx)
		for _, task := range tasks {
			task := task

			errGroup.Go(func() error {
				// TODO: handle result?
				res := task.Wait(ctx)

				return res.Err
			})
		}
		err := errGroup.Wait()

		return watcher.NewResult(nil, err)
	})
}
