package watcher

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var executionTime = prometheus.NewHistogram(prometheus.HistogramOpts{Name: "task_execution_time_histogram_ms"})

func init() {
	prometheus.MustRegister(executionTime)
}

type Timer struct {
	baseWatcher

	startTime *time.Time
}

var _ Watcher = Timer{}

func NewTimer() Timer {
	return Timer{}
}

func (t Timer) OnTaskStart(ctx context.Context, meta MetaInfo) {
	now := time.Now()
	t.startTime = &now
}

func (t Timer) OnTaskFinish(ctx context.Context, meta MetaInfo, res Result) {
	if t.startTime == nil {
		log.Errorf("start time not found")
		return
	}

	duration := time.Since(*t.startTime).Seconds()
	executionTime.Observe(duration)
}
