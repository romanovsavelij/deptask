package task

import "Pipelining/platform/watcher"

func (t Task) WithName(name string) Task {
	t.meta[watcher.MetaName] = name

	return t
}

func (t Task) WithWatcher(w watcher.Watcher) Task {
	t.watchers = append(t.watchers, w)

	return t
}
