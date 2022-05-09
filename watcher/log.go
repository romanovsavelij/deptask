package watcher

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	baseWatcher
}

var _ Watcher = Logger{}

func NewLogger() Logger {
	return Logger{}
}

func (l Logger) OnTaskFinish(ctx context.Context, meta MetaInfo, res Result) {
	name, ok := meta[MetaName]
	if !ok {
		log.Errorf("failed to find name property for a task")
		return
	}

	if res.Err != nil {
		log.Errorf("failed to execute task %v: %v", name, res.Err)
		return
	}
}
