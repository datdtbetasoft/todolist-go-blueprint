package request

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty" validate:"omitempty,oneof=pending in_progress done archived"`
	Completed   bool   `json:"completed,omitempty"`
	StartDate   string `json:"start_date,omitempty" time_format:"2006-01-02"`
	DueDate     string `json:"due_date" validate:"required" time_format:"2006-01-02"`
}
