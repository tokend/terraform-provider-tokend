package horizon

import (
	goresources "gitlab.com/tokend/go/resources"
	"gitlab.com/tokend/horizon-connector/internal/listener"
	"gitlab.com/tokend/horizon-connector/internal/operation"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"gitlab.com/tokend/regources"
	generated "gitlab.com/tokend/regources/generated"
)

// don't blame me, just make sure all exported types are really exported

type TransactionEvent = resources.TransactionEvent
type TXPacket = listener.TXPacket

// DEPRECATED: use regources directly
type Request = regources.ReviewableRequest

type WithdrawRequest = regources.WithdrawalRequest

// DEPRECATED: use regources directly
//type KYCRequest = regources.UpdateKYCRequest

type ReviewableRequestEvent = listener.ReviewableRequestEvent
type WithdrawalRequestStreamingOpts = listener.WithdrawalRequestStreamingOpts
type IssuanceRequestStreamingOpts = listener.IssuanceRequestStreamingOpts

// DEPRECATED: use regources directly
type Info = regources.Info

type Signer = goresources.Signer
type Sale = resources.Sale
type CoreSale = resources.CoreSale
type SaleDetails = resources.SaleDetails
type User = resources.User
type UserAttributes = resources.UserAttributes
type Balance = resources.Balance
type ChoppedBalance = resources.ChoppedBalance
type CheckSaleState = operations.CheckSaleState
type CheckSaleStateResponse = listener.CheckSaleStateResponse
type Blob = resources.Blob
type Document = resources.Document
type Reference = resources.Reference
type Account = resources.Account
type Wallet = resources.Wallet
type CreateKYCRequestOp = operations.CreateKYCRequest
type CreateKYCRequestOpResponse = listener.CreateKYCRequestOpResponse
type PaymentV2Op = operations.PaymentV2
type PaymentV2OpResponse = listener.PaymentOpV2Response
type ReviewRequestOp = operations.ReviewRequest
type ReviewRequestOpResponse = listener.ReviewRequestOpResponse
type Poll = generated.Poll
type Included = generated.Included
type PollResponse = generated.PollResponse

// DEPRECATED: use regources directly
//type RequestKYCDetails = regources.UpdateKYCRequest

type ReviewableRequestType = operation.ReviewableRequestType

// DEPRECATED: use regources directly
type KYCData = regources.KYCData

type KeyValue = resources.KeyValue

// DEPRECATED: use regources directly
type Asset = regources.Asset

// DEPRECATED: use regources directly
type Amount = regources.Amount
