package constants

type TaskState string

const (
	StatePending    TaskState = "pending"
	StateInProgress TaskState = "in_progress"
	StateDone       TaskState = "done"
	StateArchived   TaskState = "archived"
)
