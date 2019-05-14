package data

import "github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"

//go:generate mockery -case underscore -name Connector
type Connector interface {
	KeyValues() keyvalues.KeyValues
}
