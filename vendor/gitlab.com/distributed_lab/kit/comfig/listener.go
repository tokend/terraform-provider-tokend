package comfig

import (
	"net"
	"sync"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type Listenerer interface {
	Listener() net.Listener
}

func NewListenerer(getter kv.Getter) Listenerer {
	return &listener{
		getter: getter,
	}
}

type listener struct {
	getter kv.Getter
	once   sync.Once
	value  net.Listener
	err    error
}

func (l *listener) Listener() net.Listener {
	l.once.Do(func() {
		var config struct {
			Addr string `fig:"addr,required"`
		}
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(l.getter, "listener")).
			Please()
		if err != nil {
			l.err = errors.Wrap(err, "failed to figure out listener")
			return
		}
		l.value, l.err = net.Listen("tcp", config.Addr)
	})
	if l.err != nil {
		panic(l.err)
	}
	return l.value
}
