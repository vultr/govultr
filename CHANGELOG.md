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