package internal

import "io"

// Client exists for testing purpose only
//go:generate mockery -case underscore -name Client
type Client interface {
	Get(string) ([]byte, error)
	PostJSON(endpoint string, req interface{}) (statusCode int, response []byte, err error)
	Put(string, io.Reader) ([]byte, error)
	Delete(string) ([]byte, error)
}
