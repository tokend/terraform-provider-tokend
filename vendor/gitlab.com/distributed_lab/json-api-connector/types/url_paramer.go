package types

type QueryParamer interface {
	// Encode must return filters, page parameters and includes from struct passed
	// as an interface implementation without leading and trailing slashes
	//
	// Look into url.Values type in purposes of correct filters, page params and includes encoding
	Encode() string
}

type PathParamer interface {
	// Path must return string assembled from parameters passed
	// as an interface implementation without leading and trailing slashes
	//
	// E.g. implementation for
	// type AssetPairPath struct {
	// 	   ID string
	// }
	// (endpoint /asset_pairs/{id})
	// must return string with only with value of the field ID
	//
	// Implementation for
	// type AccountSignersPathParams struct {
	//     AccountID string
	// }
	// (endpoint /accounts/{id}/signers)
	// must return string like "${AccountID}/signers"
	Path() string
}
