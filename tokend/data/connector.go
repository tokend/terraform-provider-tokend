package data

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/connector/tx"
)

//go:generate mockery -case underscore -name Connector
type Connector interface {
	KeyValues() keyvalues.KeyValues
	Submitter() tx.HorizonSubmitter
}
