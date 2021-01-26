# Changelog

## Unreleased

### Added 

* Update signer operation builder
* Update asset operation builder

### Added

* Remove asset operation builder
* Redemption reviewable request

### Changed

* `signcontrol.IsSigned` now check if v2 auth headers presented.

## 3.10.0

### Added

* `manage_poll_op` operation with `ClosePoll` arm

## 3.9.0

### Added

* signcontrol.Error implements BadRequest() bool
* doorman.Error implements NotAllowed() bool

## 3.8.0

### Added
* Ability to set custom signer constraints for request
* Cancel change role request operation
* Initiate KYC Recovery builder

### Changed

* Atomic swap bid and ask operations

## 3.6.0

### Added
* InitiateKYCRecoveryOp
* CreateKYCRecoveryRequestOp
* KYCRecoveryRequest

## 3.5.0

### Added
* Switched to xdr 3.3.0 (new error codes, new ledger version)

## 3.4.0

### Added

* Remove Account Specific Rule builder

## 3.3.4

### Added

* Manage Account Specific Rule Op

## 3.3.3

###Added

* Cancel poll, update poll end time XDRs

## 3.3.2

### Fixed

* Support for prefixed endpoint signature

## 3.3.1

### Fixed

* doorman now properly handles case where master is not account signer

## 3.3.0

### Added

* Doorman with SignChecker (interface with GetSkipCheck method for getting passAllChecks value)

## 3.2.0

* Voting XDRs

## 3.1.0

### Added

* Valid until logic for request signatures

### Deprecated

* `signcontrol.ErrNotAllowed`

### Fixed

* `Date` header value format to conform with RFC
* `doorman` and `signcontrol` errors are now easily distinguishable

## 3.0.2-x.2

### Added

* KeyValue op builders
* LicenseOp, StampOp builders
* New result code for manage signer

### Removed

* Doorman: signer extensions via names

## 3.0.2-x.1

### Changed

Bumped XDR to `3.0.1-x.0`

## 3.0.2-x.0

### Changed

* Squashed ledger version;

### Removed

* Payment and Direct Debit Operations
