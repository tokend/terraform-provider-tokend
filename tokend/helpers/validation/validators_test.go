package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSource(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		source := ""
		_, errors := ValidateSource(source, "")
		assert.NotEqual(t, 0, len(errors))
	})

	t.Run("invalid", func(t *testing.T) {
		source := "SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4"
		_, errors := ValidateSource(source, "")
		assert.NotEqual(t, 0, len(errors))
	})

	t.Run("valid", func(t *testing.T) {
		source := "GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB"
		_, errors := ValidateSource(source, "")
		assert.Equal(t, 0, len(errors))
	})
}

func TestValidateSigner(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		signer := ""
		_, errors := ValidateSigner(signer, "")
		assert.NotEqual(t, 0, len(errors))
	})

	t.Run("invalid", func(t *testing.T) {
		signer := "GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB"
		_, errors := ValidateSigner(signer, "")
		assert.NotEqual(t, 0, len(errors))
	})

	t.Run("valid", func(t *testing.T) {
		signer := "SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4"
		_, errors := ValidateSigner(signer, "")
		assert.Equal(t, 0, len(errors))
	})
}

func TestNonEmptyString(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		str := ""
		_, errors := NonEmptyString(str, "")
		assert.Equal(t, 1, len(errors))
	})

	t.Run("valid", func(t *testing.T) {
		str := "str"
		_, errors := NonEmptyString(str, "")
		assert.Equal(t, 0, len(errors))
	})
}
