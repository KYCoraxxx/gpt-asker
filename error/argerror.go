package error

type ArgError struct {
	msg string
}

func NewArgError(msg string) ArgError {
	err := ArgError{msg: msg}
	return err
}

func (e ArgError) Error() string {
	panic("Fatal Error: Argument resolving failed. " + e.msg)
}
