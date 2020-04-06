CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to support Go modules

---

## 1.5.0 (2020-03-23)
* Upgrade to v1.1.2 of the BigIP Terraform Provider
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.4.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.3.0 (2020-01-02)
* Namespace names in .NET SDK are adjusted to PascalCase
([#38](https://github.com/pulumi/pulumi-f5bigip/pull/38)).
* Upgrade to v1.1.1 of the BigIP Terraform Provider
* Upgrade to pulumi-terraform-bridge v1.5.2

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform@dotnet

## 1.1.0 (2019-11-13)
* Remove README.rst from the Python package and replace it with README.md

## 1.0.0 (2019-11-13)
* Add support for DotNet SDK Generation

## 0.2.11 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.2.10 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.2.9 (2019-08-20)
* Upgrade to v0.12.4 of the BigIP Terraform Provider
* Depend on latest pulumi package

## 0.2.8 (2019-08-09)
* Upgrade to pulumi-terraform@9db2fc93cd

## 0.2.7 (2019-08-08)
* Update to pulumi-terraform@013b95b1c8

## 0.2.6 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.2.5 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.2.4 (2019-06-21)
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.2.3 (2019-06-06)
* Add TypeScript type guards for each resource class ([7ace3e9b5f](https://github.com/pulumi/pulumi-terraform/commit/7ace3e9b5f2dcd4692b029ba4b80360582d7949a))
* Update to v0.12.3 of the Big IP Terraform provider

## 0.2.2 (2019-05-31)
* Update to v0.12.2 of the Big IP Terraform Provider

## 0.2.1 (2019-04-22)
* Update to the latest master branch of the Terraform F5 Big IP provider

## 0.2.0 (2019-03-06)
* Update to the latest version of the `pulumi` SDK

## 0.1.2 (2019-02-20)
* Add support for the `deleteBeforeReplace` resource option and improved delete-before-replace behaviour introduced in Pulumi v0.16.14

## 0.1.1 (2019-02-01)
* Update to v0.12.0 of the Big IP Terraform Provider
* Add documentation and examples

## 0.1.0 (2019-02-01)
* Initial release of F5 Big IP provider
