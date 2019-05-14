# Changelog

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
