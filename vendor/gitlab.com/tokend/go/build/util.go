package build

import (
	"errors"

	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdr"
)

func setAccountId(addressOrSeed string, aid *xdr.AccountId) error {
	kp, err := keypair.Parse(addressOrSeed)
	if err != nil {
		return err
	}

	if aid == nil {
		return errors.New("aid is nil in setAccountId")
	}

	return aid.SetAddress(kp.Address())
}
