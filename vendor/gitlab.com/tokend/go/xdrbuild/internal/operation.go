package internal

import "gitlab.com/tokend/go/xdr"

type Operation interface {
	XDR() (*xdr.Operation, error)
}
