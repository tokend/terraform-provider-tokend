package horizon

type Error interface {
	error
	Status() int
	Body() []byte
	Path() string
}
