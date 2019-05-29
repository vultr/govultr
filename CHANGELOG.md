# Change Log

## v0.1.0 (2019-05-10)
### Features
* Initial release
* Supports all available API endpoints that Vultr has to offer


## [v0.1.1](https://github.com/vultr/govultr/compare/v0.1.0..v0.1.1) (2019-05-20)
### Features
* add `SnapshotID` to ServerOptions as an option during server creation
* bumped default RateLimit from `.2` to `.6` seconds
### Breaking Changes
* Iso
  * Renamed all instances of `Iso` to `ISO`.  
* BlockStorage
  * Renamed `Cost` to `CostPerMonth`
  * Renamed `Size` to `SizeGB` 
* BareMetal & Server 
  * Change `SSHKeyID` to `SSHKeyIDs` which are now `[]string` instead of `string`
  * Renamed `OS` to `Os`

## [v0.1.2](https://github.com/vultr/govultr/compare/v0.1.1..v0.1.2) (2019-05-29)
### Fixes
* Fixed Server Option `NotifyActivate` bug that ignored a `false` value
* Fixed Bare Metal Server Option `UserData` to be based64encoded 
### Breaking Changes
* Renamed all methods named `GetList` to `List`
* Renamed all methods named `Destroy` to `Delete`
* Server Service
    * Renamed `GetListByLabel` to `ListByLabel`
    * `GetListByMainIP` to `ListByMainIP`
    * `GetListByTag` to `ListByTag`
* Bare Metal Server Service
    * Renamed `GetListByLabel` to `ListByLabel`
    * `GetListByMainIP` to `ListByMainIP`
    * `GetListByTag` to `ListByTag`