package scanner

type Statuses string

const (
	StatusOk       Statuses = "ok"
	StatusWarning  Statuses = "warning"
	StatusCritical Statuses = "critical"
)

var statusPriority = map[Statuses]int{
	StatusCritical: 0,
	StatusWarning:  1,
	StatusOk:       2,
}

type Result struct {
	Name    string
	Status  Statuses
	Details string
}

type Scanner interface {
	Scan(host string) []Result
}
