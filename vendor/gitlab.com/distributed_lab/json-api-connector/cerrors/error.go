package cerrors

func (e Error) Status() int {
	return int(e.status)
}

func (e Error) Body() []byte {
	return e.body
}

func (e Error) Path() string {
	return string(e.path)
}
