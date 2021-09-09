package problemdetails

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string
	Error string
}

type ValidationProblem struct {
	Type     string `json:"type"`               // url to the problem type
	Title    string `json:"title"`              // problem title
	Status   int    `json:"status,omitempty"`   // http status code
	Detail   string `json:"detail,omitempty"`   // a human-readable explaination
	Instance string `json:"instance,omitempty"` // A URI reference that identifies the specific ocurrence. Does not need to derefrence to something

	// Fields added by VQ
	ValidationErrors []ValidationError `json:"validationerrors,omitempty"` // a list of the validation errors that occurred
}

func (problem ValidationProblem) WithType(t string) ValidationProblem {
	problem.Type = t
	return problem
}

func (problem *ValidationProblem) WithTitle(title string) *ValidationProblem {
	problem.Title = title
	return problem
}

func (problem *ValidationProblem) WithStatus(status int) *ValidationProblem {
	problem.Status = status
	return problem
}

func (problem *ValidationProblem) WithDetail(detail string) *ValidationProblem {
	problem.Detail = detail
	return problem
}

func (problem *ValidationProblem) WithInstance(instance string) *ValidationProblem {
	problem.Instance = instance
	return problem
}

func (problem *ValidationProblem) ProblemType() string {
	return problem.Type
}

func (problem *ValidationProblem) ProblemTitle() string {
	return problem.Title
}

func (problem *ValidationProblem) ProblemStatus() int {
	return problem.Status
}

func (problem *ValidationProblem) ProblemDetail() string {
	return problem.Detail
}

func (problem *ValidationProblem) ProblemInstance() string {
	return problem.Instance
}

func FromValidationErrors(validationErrors validator.ValidationErrors) ValidationProblem {
	var errs []ValidationError

	for _, e := range validationErrors {
		errs = append(errs, ValidationError{
			Field: e.Field(),
			Error: e.Error(),
		})
	}

	return ValidationProblem{
		// TODO: Add status code handling here later on, for now about:blank is fine by the RFC
		Type:             "about:blank",
		Status:           400,
		Detail:           "Invalid request",
		ValidationErrors: errs,
	}
}
