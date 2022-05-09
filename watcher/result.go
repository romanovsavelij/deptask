package watcher

type Result struct {
	Value any
	Err   error
}

func NewResult(value any, err error) Result {
	return Result{
		Value: value,
		Err:   err,
	}
}
