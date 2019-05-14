package operation

type ReviewableRequestType string

const (
	// WithdrawalsReviewableRequestType stands for both Withdraw and TwoStepWithdraw ReviewableRequests
	WithdrawalsReviewableRequestType ReviewableRequestType = "withdrawals"
	KYCReviewableRequestType         ReviewableRequestType = "update_kyc"

	// For RequestTypes AssetCreate and AssetUpdate
	AssetReviewableRequestType ReviewableRequestType = "assets"
	// For PreIssuanceCreate RequestType
	PreIssuanceReviewableRequestType ReviewableRequestType = "preissuances"
	// For IssuanceCreate RequestType
	IssuanceReviewableRequestType ReviewableRequestType = "issuances"
	// For Sale RequestType
	SaleReviewableRequestType ReviewableRequestType = "sales"
	// For LimitsUpdate RequestType
	LimitsUpdateReviewableRequestType ReviewableRequestType = "limits_updates"
)
