package resources

import "gitlab.com/tokend/regources"

// TODO Comment
// TODO Consider moving the type into listener package
type TransactionEvent struct {
	Transaction *regources.Transaction
	Meta        regources.PageMeta
}
