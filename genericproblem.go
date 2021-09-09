package problemdetails

type GenericProblem struct {
	Type     string `json:"type"`               // url to the problem type
	Title    string `json:"title"`              // problem title
	Status   int    `json:"status,omitempty"`   // http status code
	Detail   string `json:"detail,omitempty"`   // a human-readable explaination
	Instance string `json:"instance,omitempty"` // A URI reference that identifies the specific ocurrence. Does not need to derefrence to something
}

func (problem GenericProblem) WithType(t string) GenericProblem {
	problem.Type = t
	return problem
}

func (problem *GenericProblem) WithTitle(title string) *GenericProblem {
	problem.Title = title
	return problem
}

func (problem *GenericProblem) WithStatus(status int) *GenericProblem {
	problem.Status = status
	return problem
}

func (problem *GenericProblem) WithDetail(detail string) *GenericProblem {
	problem.Detail = detail
	return problem
}

func (problem *GenericProblem) WithInstance(instance string) *GenericProblem {
	problem.Instance = instance
	return problem
}
