package scanner

type Result struct {
	Name    string
	Status  string
	Details string
}

type Scanner interface {
	Scan(host string) []Result
}
