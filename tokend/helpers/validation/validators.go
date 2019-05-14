package validation

import (
	"fmt"

	"gitlab.com/tokend/go/strkey"
)

func NonEmptyString(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}

	if v == "" {
		return nil, []error{fmt.Errorf("%q must not be empty", k)}
	}

	return nil, nil
}

func ValidateSource(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}
	_, err := strkey.Decode(strkey.VersionByteAccountID, v)
	if err != nil {
		return nil, []error{fmt.Errorf("%q illegal base32 data at input byte 48", k)}
	}
	return nil, nil
}

func ValidateSigner(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}
	_, err := strkey.Decode(strkey.VersionByteSeed, v)
	if err != nil {
		return nil, []error{fmt.Errorf("%q illegal base32 data at input byte 48", k)}
	}
	return nil, nil
}
