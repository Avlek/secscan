package scanner

type Statuses string

const (
	StatusOk       Statuses = "ok"
	StatusWarning  Statuses = "warning"
	StatusCritical Statuses = "critical"
)

type Result struct {
	Name    string
	Status  Statuses
	Details string
}

type Scanner interface {
	Scan(host string) []Result
}
