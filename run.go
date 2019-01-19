package dots

const (
	exitOK = iota
	exitError
)

func Run(c cmd) int {
	return exitOK
}
