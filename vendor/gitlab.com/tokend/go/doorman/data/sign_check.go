package data

import "context"

// SkipChecker check if it's needed to disable constraints validation completely
type SkipChecker interface {
	GetSkipCheck(context.Context) (bool, error)
}

type Checker struct {
	passAllChecks bool
}

func NewChecker(passAllCheck bool) SkipChecker {
	return &Checker{
		passAllChecks: passAllCheck,
	}
}

func (c *Checker) GetSkipCheck(ctx context.Context) (bool, error) {
	return c.passAllChecks, nil
}
