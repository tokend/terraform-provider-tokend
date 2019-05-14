package responses

import "gitlab.com/tokend/go/resources"

type AccountSigners struct {
	Signers []resources.Signer `json:"signers"`
}
