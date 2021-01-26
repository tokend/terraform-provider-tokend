package data

import "gitlab.com/tokend/go/resources"

//go:generate mockery -case underscore -name AccountQ -testonly  -outpkg=doorman -output=../

// AccountQ data access interface for getting account details
type AccountQ interface {
	// Signers get account signers, nil slice is expected if account is not found
	Signers(address string) ([]resources.Signer, error)
}
