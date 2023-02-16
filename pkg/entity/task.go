package entity

type Task struct {
	TaskId                 int    `json:"taskId"`
	Description            string `json:"description"`
	MaximumExecutionValue  int    `json:"maximumExecutionValue"`
	ActualExecutionValue   int    `json:"actualExecutionValue"`
	PercentageOfCompletion int8   `json:"percentageOfCompletion"`
	CheckListId            int    `json:"checkListId"`
}
