package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain"
)

var validate *validator.Validate

func init() {
	panic("TODO(reto): import")
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult struct {
	Valid  bool             `json:"valid"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func ValidateTask(title, description string, status domain.TaskStatus) ValidationResult {
	panic("TODO(reto): ValidateTask")
}

func isValidStatus(status domain.TaskStatus) bool {
	panic("TODO(reto): isValidStatus")
}

func ValidateTaskInput(title, description string, status domain.TaskStatus) error {
	panic("TODO(reto): ValidateTaskInput")
}