CHANGELOG
=========

## Notice (2022-01-06)

*As of this notice, using CHANGELOG.md is DEPRECATED. We will be using [GitHub Releases](https://github.com/pulumi/pulumi-f5bigip/releases) for this repository*

## HEAD (Unreleased)
_(none)_

---

## 3.6.1 (2021-12-30)
* Upgrade terraform-provider-bigip to v1.12.1

## 3.6.0 (2021-11-15)
* Upgrade to terraform-bridge 3.11.0
* Upgrade to pulumi 3.17.0

## 3.5.0 (2021-10-29)
* Upgrade to v1.12.0 of the BigIP Terraform Provider

## 3.4.1 (2021-09-27)
* Upgrade to v1.11.1 of the BigIP Terraform Provider

## 3.4.0 (2021-08-12)
* Upgrade to v1.11.0 of the BigIP Terraform Provider

## 3.3.0 (2021-07-12)
* Upgrade to v1.10.0 of the BigIP Terraform Provider
  **Please Note:** This includes a breaking change to remove `defaultsFrom` from `f5bigip.ltm.Monitor`.

## 3.2.0 (2021-05-27)
* Upgrade to v3.2.1 of the pulumi-terraform-bridge

## 3.1.0 (2021-05-21)
* Upgrade to v1.9.0 of the BigIP Terraform Provider

## 3.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 2.11.0 (2021-04-12)
* Upgrade to v1.8.0 of the BigIP Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.23.0

## 2.10.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 2.10.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 2.9.0 (2021-02-19)
* Upgrade to v1.7.0 of the BigIP Terraform Provider

## 2.8.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider

## 2.8.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.1

## 2.7.1 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.7.0 (2021-01-08)
* Upgrade to v1.6.0 of the BigIP Terraform Provider

## 2.6.0 (2020-12-08)
* Upgrade to v1.5.0 of the BigIP Terraform Provider

## 2.5.2 (2020-11-23)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.5.1 (2020-11-06)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.5.0 (2020-10-26)
* Upgrade to v1.4.0 of the BigIP Terraform Provider
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.4.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.3.0 (2020-07-24)
* Upgrade to v1.3.0 of the BigIP Terraform Provider

## 2.2.0 (2020-07-07)
* Upgrade to v1.2.1 of the BipIP Terraform Provider

## 2.1.3 (2020-06-11)
* Switch to GitHub actions for build

## 2.1.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.1.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.6.0 (2020-04-14)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to support Go modules

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
