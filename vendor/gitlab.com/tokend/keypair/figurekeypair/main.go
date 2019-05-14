package figurekeypair

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/tokend/keypair"
)

var Hooks = figure.Hooks{
	"keypair.Address": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			kp, err := keypair.ParseAddress(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse kp")
			}
			return reflect.ValueOf(kp), nil
		case nil:
			return reflect.ValueOf(nil), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
	"keypair.Full": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			kp, err := keypair.ParseSeed(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse kp")
			}
			kpFull, ok := kp.(keypair.Full)
			if !ok {
				return reflect.Value{}, errors.Wrap(err,
					"failed to cast kp to keypair.Full; string must be a Seed")
			}
			return reflect.ValueOf(kpFull), nil
		case nil:
			return reflect.ValueOf(nil), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
