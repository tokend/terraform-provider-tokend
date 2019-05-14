package internal

type Validatable interface {
	Validate() error
}
