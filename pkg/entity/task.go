package entity

type task struct {
	taskId                 int64
	description            string
	maximumExecutionValue  int32
	actualExecutionValue   int32
	percentageOfCompletion int8
}

func New() *task {

	return nil
}
