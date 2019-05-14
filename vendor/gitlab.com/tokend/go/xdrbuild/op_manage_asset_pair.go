package xdrbuild

import (
	"gitlab.com/tokend/go/xdr"
)

type CreateAssetPair struct {
	Base                    string
	Quote                   string
	PhysicalPrice           int64
	PhysicalPriceCorrection int64
	MaxPriceStep            int64
	Policies                int32
}

type UpdateAssetPairPrice struct {
	Base          string
	Quote         string
	PhysicalPrice int64
}

type UpdateAssetPairPolicies struct {
	Base                    string
	Quote                   string
	PhysicalPriceCorrection int64
	MaxPriceStep            int64
	Policies                int32
}

func (ap CreateAssetPair) XDR() (*xdr.Operation, error) {
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAssetPair,
			ManageAssetPairOp: &xdr.ManageAssetPairOp{
				Action:                  xdr.ManageAssetPairActionCreate,
				Base:                    xdr.AssetCode(ap.Base),
				Quote:                   xdr.AssetCode(ap.Quote),
				PhysicalPrice:           xdr.Int64(ap.PhysicalPrice),
				Policies:                xdr.Int32(ap.Policies),
				PhysicalPriceCorrection: xdr.Int64(ap.PhysicalPriceCorrection),
				MaxPriceStep:            xdr.Int64(ap.MaxPriceStep),
			},
		},
	}
	return op, nil
}

func (ap UpdateAssetPairPrice) XDR() (*xdr.Operation, error) {
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAssetPair,
			ManageAssetPairOp: &xdr.ManageAssetPairOp{
				Action:        xdr.ManageAssetPairActionUpdatePrice,
				Base:          xdr.AssetCode(ap.Base),
				Quote:         xdr.AssetCode(ap.Quote),
				PhysicalPrice: xdr.Int64(ap.PhysicalPrice),
			},
		},
	}
	return op, nil
}

func (ap UpdateAssetPairPolicies) XDR() (*xdr.Operation, error) {
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAssetPair,
			ManageAssetPairOp: &xdr.ManageAssetPairOp{
				Action:                  xdr.ManageAssetPairActionUpdatePolicies,
				Base:                    xdr.AssetCode(ap.Base),
				Quote:                   xdr.AssetCode(ap.Quote),
				Policies:                xdr.Int32(ap.Policies),
				PhysicalPriceCorrection: xdr.Int64(ap.PhysicalPriceCorrection),
				MaxPriceStep:            xdr.Int64(ap.MaxPriceStep),
			},
		},
	}
	return op, nil
}
