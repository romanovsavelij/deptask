package task

import (
	"Pipelining/watcher"
	"context"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Task(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var counter int32 = 0
		ctx := context.Background()

		task1 := NewTask(func(ctx context.Context) watcher.Result {
			atomic.AddInt32(&counter, 1)

			return watcher.NewResult(nil, nil)
		})

		res := task1.Wait(ctx)
		require.NoError(t, res.Err)
		assert.Equal(t, nil, res.Value)
		assert.Equal(t, int32(1), counter)
	})

	t.Run("default task", func(t *testing.T) {
		taskName := "task with name"
		task := NewDefaultTask(func(ctx context.Context) watcher.Result {
			return watcher.NewResult(nil, nil)
		}, taskName)

		assert.Equal(t, 2, len(task.watchers))
		assert.Equal(t, 1, len(task.meta))
		assert.Equal(t, taskName, task.meta[watcher.MetaName])
	})
}
