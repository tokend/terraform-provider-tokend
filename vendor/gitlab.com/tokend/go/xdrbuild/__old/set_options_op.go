package __old

import (
	"reflect"

	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

var (
	ErrEmptyOp = errors.New("at least one value was expected")
)

const maxSignerWeight = 255

type Signer struct {
	PublicKey  string
	Weight     uint32
	SignerType uint32
	Identity   uint32
	Name       string
}

func (s Signer) Validate() error {
	errs := Errors{
		"public_key": Validate(s.PublicKey, Required),
	}

	if s.Weight != 0 {
		errs = Errors{
			"signer_type": Validate(s.SignerType, Required),
			"identity":    Validate(s.Identity, Required),
			"name":        Validate(s.Name, Required),
			"weight":      Validate(int64(s.Weight), Required, Max(maxSignerWeight)),
		}
	}

	return errs.Filter()
}

type SetOptions struct {
	Signer        *Signer
	MasterWeight  *uint32
	LowThreshold  *uint32
	MedThreshold  *uint32
	HighThreshold *uint32
}

func (op SetOptions) Validate() error {
	errs := Errors{}

	if op.Signer != nil {
		errs["/signer"] = op.Signer.Validate()
	}

	if op.MasterWeight != nil {
		errs["/master_weight"] = Validate(int64(*op.MasterWeight), Max(maxSignerWeight))
	}

	if op.LowThreshold != nil {
		errs["/low_threshold"] = Validate(int64(*op.LowThreshold), Max(maxSignerWeight))
	}

	if op.MedThreshold != nil {
		errs["/med_threshold"] = Validate(int64(*op.MedThreshold), Max(maxSignerWeight))
	}

	if op.HighThreshold != nil {
		errs["/high_threshold"] = Validate(int64(*op.HighThreshold), Max(maxSignerWeight))
	}

	if reflect.DeepEqual(op, SetOptions{}) {
		errs["/"] = ErrEmptyOp
	}

	return errs.Filter()
}

func (op SetOptions) XDR() (*xdr.Operation, error) {
	xdrOp := xdr.SetOptionsOp{
		MasterWeight:  (*xdr.Uint32)(op.MasterWeight),
		LowThreshold:  (*xdr.Uint32)(op.LowThreshold),
		MedThreshold:  (*xdr.Uint32)(op.MedThreshold),
		HighThreshold: (*xdr.Uint32)(op.HighThreshold),
	}

	if op.Signer != nil {
		var signerPubKey xdr.AccountId
		if err := signerPubKey.SetAddress(op.Signer.PublicKey); err != nil {
			return nil, errors.Wrap(err, "failed to set signer public key")
		}

		xdrOp.Signer = &xdr.Signer{
			PubKey:     signerPubKey,
			Weight:     xdr.Uint32(op.Signer.Weight),
			SignerType: xdr.Uint32(op.Signer.SignerType),
			Identity:   xdr.Uint32(op.Signer.Identity),
			Name:       xdr.String256(op.Signer.Name),
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:         xdr.OperationTypeSetOptions,
			SetOptionsOp: &xdrOp,
		},
	}, nil
}

//DeleteSigner create new SetOptions without signer weight
//by default if signer weight is zero, it mean delete signer
func DeleteSigner(publicKey string) *SetOptions {
	return &SetOptions{
		Signer: &Signer{
			PublicKey: publicKey,
			Weight:    0,
		},
	}
}

func AddSigner(publicKey, name string, weight, signerType, identity uint32) *SetOptions {
	return &SetOptions{
		Signer: &Signer{
			PublicKey:  publicKey,
			Weight:     weight,
			SignerType: signerType,
			Identity:   identity,
			Name:       name,
		},
	}
}

func SetThresholds(masterWeight, lowThreshold, medThreshold, highThreshold uint32) *SetOptions {
	return &SetOptions{
		MasterWeight:  &masterWeight,
		LowThreshold:  &lowThreshold,
		MedThreshold:  &medThreshold,
		HighThreshold: &highThreshold,
	}
}
