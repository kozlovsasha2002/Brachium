package entity

type Task struct {
	TaskId                 int64  `json:"taskId"`
	Description            string `json:"description"`
	MaximumExecutionValue  int32  `json:"maximumExecutionValue"`
	ActualExecutionValue   int32  `json:"actualExecutionValue"`
	PercentageOfCompletion int8   `json:"percentageOfCompletion"`
}
