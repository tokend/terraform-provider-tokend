package doorman

import (
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/doorman/data"
	"gitlab.com/tokend/go/resources"
)

// Doorman interface only purpose is to simplify tests
type Doorman interface {
	Check(*http.Request, ...SignerConstraint) error
	AccountSigners(string) ([]resources.Signer, error)
	DefaultSignerOfConstraints() []SignerOfExt
}

type doorman struct {
	// SkipChecker check if it's needed to disable constraints validation completely, any request will succeed
	SkipChecker data.SkipChecker
	// AccountQ used to get account details during constraint checks
	AccountQ data.AccountQ
	// signerOfExts are used to specify list of constraints for SignerOf
	signerOfExts []SignerOfExt
}

func (d *doorman) AccountSigners(address string) ([]resources.Signer, error) {
	return d.AccountQ.Signers(address)
}

// Check ensures request passes at least one constraint
// return non-doorman error if checker failed to get skip_check value
func (d *doorman) Check(r *http.Request, constraints ...SignerConstraint) error {
	passAllChecks, err := d.SkipChecker.GetSkipCheck(r.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get skip_check value")
	}

	if passAllChecks {
		return nil
	}

	for _, constraint := range constraints {
		switch err := constraint(r, d); err {
		case nil:
			// request passed constraint check
			return nil
		case ErrNotAllowed:
			// check failed, let's try next one
			continue
		default:
			// probably runtime issue
			return err
		}
	}

	// request failed all checks
	return ErrNotAllowed
}

func (d *doorman) DefaultSignerOfConstraints() []SignerOfExt {
	return d.signerOfExts
}

type RoleConstraint struct {
	RoleID uint64
}

func (c *RoleConstraint) Check(signer resources.Signer) bool {
	return c.RoleID == signer.Role
}

type RestrictedRoleConstraint struct {
	RoleID uint64
}

func (c *RestrictedRoleConstraint) Check(signer resources.Signer) bool {
	return c.RoleID != signer.Role
}
