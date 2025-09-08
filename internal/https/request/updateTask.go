package request

type UpdateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	State       string `json:"state" validate:"required"`
	Completed   bool   `json:"completed"`
	StartDate   string `json:"start_date" validate:"required"` // yyyy-MM-dd
	DueDate     string `json:"due_date" validate:"required"`   // yyyy-MM-dd
}
