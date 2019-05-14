package comfig

import (
	"sync"
)

// Once like `sync.Once` but also stores return and panic values
type Once struct {
	m     sync.Mutex
	err   interface{}
	value interface{}
}

func (o *Once) Do(f func() interface{}) interface{} {
	o.m.Lock()
	defer o.m.Unlock()
	if o.err != nil {
		panic(o.err)
	}
	if o.value != nil {
		return o.value
	}
	defer func() {
		if rvr := recover(); rvr != nil {
			o.err = rvr
			panic(rvr)
		}
	}()
	o.value = f()
	return o.value
}
