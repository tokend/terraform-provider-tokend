# Changelog

## 4.7.0

### Added 

* `KYCRecoveryStatus`
* `AccountAttributes`

### Fixed

* Atomic swap ask and bid typo issue

## 4.6.0

### Added

* `AccountKyc` resource to handle kyc data for account
* Operation to remove asset pair
* `Atomic swap ask` resource

## 4.5.0

### Added

* `trailing_digits_count` to asset creation request attributes
* `Asset` relationship to create withdraw request relationships

## 4.4.1

### Fixed

* CreateKycRecoveryRequestOpAttributes

## 4.4.0

### Added

* `ManageAccountSpecificRuleOp` resource to handle `/v3/history`

## 4.3.4

### Added

* `AccessDefinitionType` attribute to sale

### Changed

* Switched to go 3.5.0

## 4.3.3

### Added

* `BaseHardCap` to `SaleAttributes` resource
* `SoftCap` and `HardCap` `CreateSaleRequestAttributes`

## 4.3.2

### Added

* `Asset` relationship to `ConvertedBalancesCollection`

## 4.3.1

### Added
* `Match` resource

## 4.3.0

### Added
* `Statistics` resource
* `LimitsWithStats` resource

## Removed
* `Asset` relation from `ConvertedBalanceStateRelationships`

## 4.2.3
### Added
* `ConvertedBalancesCollection` and `ConvertedBalanceState` resources

## 4.2.2
### Fixed
* Updated tokend/go to 3.3.4
## 4.2.1
### Added
* Manage Account Specific Rules for Sale
## 4.2.0

### Added 

* `Owner` relation to `Balance` (ID of the owner's account)
* `Horizon-State` resource for `/v3` endpoint (info regarding horizon and it's dependencies state)
* `SubmitTransactionBody` resource for `POST /v3/transactions` endpoint (now possible to specify if we should wait for tx to be ingested)
* `Cancelled` poll state
* `UpdatePollEndTimeOp` type
* `PollId` moved to `ManagePollOp`

## 4.1.0

### Added

* generated `OrderBook`, `OrderBookRelationships`, `OrderBookResponse`, `OrderBooksResponse` types to represent order book

## 4.0.0

### Added

* generated all existing models from openapi spec

## 3.1.0

### Added
* `CreatePollRequest` type to represent reviewable request for poll creation
* `Poll`, `Vote` types
* `PollResponse`, `PollsResponse`,`VoteResponse`, `VotesResponse`
* `PollParticipation` type to represent voting result


## 3.0.3

### Added

* `TransactionsResponse`, `Transaction`,  types to represent the get transactions response
* `TransactionResponseMeta`, `TransactionAttrs`, `TransactionRelations` helping structures
* `LedgerEntryChange`, `LedgerEntryChangeAttrs` types to represent the ledger entry change resource
* `ledger-entry-changes` and `transactions` resource types

## 3.0.2

## 3.0.2-x.0

### Added

* `BalancesResponse` type to represent the response on `/balances`
* `PublicKeyEntryResponse` type to represent the response on `/public_key_entries`
* `PublicKeyEntry`, `PublicKeyEntryRelationships` types

## 3.0.1-x.1

### Added

* Asset type to asset, manage asset, create asset request resources
* Relations `Limits` and `ExternalSystemIDs` to Account
* Extended Fee type to it as flag

### Fixed

* Use xdr.AccountRuleAction and xdr.SignerRuleAction
* (internal) proper types for reviewable requests amounts
* (internal) removed KYCData from ChangeRoleRequest

### Changed

* Rename is_forbid to forbids
* Rename details to creator_details in ops:
    create_aml_alert_request
    create_change_role_request
    create_issuance_request
    create_manage_limits_request
    create_pressuance_request
    create_sale_request
    create_withdraw_request
    manage_asset

## 3.0.1-x.0

### Changed

* Bumped tokend/go to `3.0.2-x.0` (switched to XDR `> 3.0.1`)

## 3.0.0-x.1

### Changed

* `PreIssuanceRequestAttrs.Amount` now has `Amount` type
* `WithdrawalRequestAttrs.Fee` is now type `Fee` instead of `FeeStr`

### Removed

* `FeeStr` type
