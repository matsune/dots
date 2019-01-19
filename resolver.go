package dots

type resolver interface {
	ReadTargets() ([]target, error)
	do(target) error
}
