package dromaius

var (
	errCommandNotAllowed     = errors.New("command not allowed")
	errCommandNotImplemented = errors.New("command not implemented")
	errHostNotRecognized     = errors.New("incorrect hostname supplied")
)
