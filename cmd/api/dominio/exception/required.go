package exception

type Required interface {
	Error() string
	Required() bool
}

type DataRequired struct {
	ErrMessage string
}

func (required DataRequired) Error() string {
	return required.ErrMessage
}

func (required DataRequired) Required() bool {
	return true
}
