package problemdetails

type Problem interface {
	WithType(t string) *Problem
	WithTitle(title string) *Problem
	WithStatus(status string) *Problem
	WithDetail(detail string) *Problem
	WithInstance(detail string) *Problem
}
