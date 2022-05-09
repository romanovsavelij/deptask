package task

import (
	"Pipelining/watcher"
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WaitGroup(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var counter int32 = 0
		ctx := context.Background()

		task1 := NewTask(func(ctx context.Context) watcher.Result {
			atomic.AddInt32(&counter, 1)

			return watcher.NewResult(nil, nil)
		})

		task2 := NewTask(func(ctx context.Context) watcher.Result {
			atomic.AddInt32(&counter, 1)

			return watcher.NewResult(nil, nil)
		})

		task3 := NewTask(func(ctx context.Context) watcher.Result {
			atomic.AddInt32(&counter, 1)

			return watcher.NewResult(nil, nil)
		})

		taskGroup := WaitGroup(task1, task2, task3)

		res := taskGroup.Wait(ctx)
		require.NoError(t, res.Err)
		assert.Equal(t, nil, res.Value)
		assert.Equal(t, int32(3), counter)
	})

	t.Run("context cancel", func(t *testing.T) {
		// TODO
	})
}

func Test_ErrGroup(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// TODO
	})

	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		someErr := errors.New("some error")

		task1 := NewTask(func(ctx context.Context) watcher.Result {
			return watcher.NewResult(nil, someErr)
		})

		task2 := NewTask(func(ctx context.Context) watcher.Result {
			return watcher.NewResult(nil, nil)
		})

		task := ErrGroup(task1, task2)

		res := task.Wait(ctx)
		require.Empty(t, res.Value)
		require.Error(t, res.Err)
		assert.Equal(t, someErr, res.Err)
	})
}
