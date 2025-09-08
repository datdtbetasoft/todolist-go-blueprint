package request

import "time"

type RegisterRequest struct {
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Birthday time.Time `json:"birthday,omitempty" time_format:"2006-01-02"`
	Password string    `json:"password" validate:"required"`
	Provider string    `json:"provider,omitempty"`
}
