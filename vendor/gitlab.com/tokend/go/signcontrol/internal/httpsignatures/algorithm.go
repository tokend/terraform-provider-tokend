package httpsignatures

import (
	"errors"
	"sync"
)

var (
	mutex              = sync.Mutex{}
	algorithms         = map[string]Algorithm{}
	ErrNoSuchAlgorithm = errors.New("no such algorithm")
	ErrUnexpectedKey   = errors.New("unexpected key")
)

func RegisterAlgorithm(algo Algorithm) {
	mutex.Lock()
	defer mutex.Unlock()
	algorithms[algo.Name()] = algo
}

func GetAlgorithm(name string) (Algorithm, error) {
	mutex.Lock()
	defer mutex.Unlock()
	algorithm, ok := algorithms[name]
	if !ok {
		return nil, ErrNoSuchAlgorithm
	}
	return algorithm, nil
}

type Algorithm interface {
	Name() string
	Sign(key interface{}, digest []byte) ([]byte, error)
	Verify(signature Signature, digest []byte) bool
}
