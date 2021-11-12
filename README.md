[![Actions Status](https://github.com/pulumi/pulumi-f5bigip/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-f5bigip/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Ff5bigip.svg)](https://www.npmjs.com/package/@pulumi/f5bigip)
[![Python version](https://badge.fury.io/py/pulumi-f5bigip.svg)](https://pypi.org/project/pulumi-f5bigip)
[![NuGet version](https://badge.fury.io/nu/pulumi.f5bigip.svg)](https://badge.fury.io/nu/pulumi.f5bigip)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-f5bigip/sdk/v3/go)](https://pkg.go.dev/github.com/pulumi/pulumi-f5bigip/sdk/v3/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-f5bigip/blob/master/LICENSE)

# F5 BigIP Provider

This provider allows management of F5 BigIP resources using Pulumi. This provider uses the iControlREST API to
perform management tasks, so it will need to be installed and enabled on your F5 device before proceeding.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/f5bigip

or `yarn`:

    $ yarn add @pulumi/f5bigip

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_f5bigip

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-f5bigip/sdk/v3
    
### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.F5bigip
    
## Configuration

The following configuration points are available:

- `f5bigip:address` - Domain name/IP of the BigIP. May be set via the `BIGIP_HOST` environment variable.
- `f5bigip:port` - Management Port to connect to BigIP.
- `f5bigip:username` - Username with API access to the BigIP. May be set via the `BIGIP_USER` environment variable.
- `f5bigip:password` - Password for API access to the BigIP. May be set via the `BIGIP_PASSWORD` environment variable.
- `f5bigip:tokenAuth` - Enable to use an external authentication source (LDAP, TACACS, etc). May be set via the `BIGIP_TOKEN_AUTH` environment variable.
- `f5bigip:tokenAuth` - Enable to use an external authentication source (LDAP, TACACS, etc). May be set via the `BIGIP_TOKEN_AUTH` environment variable.
- `f5bigip:loginRef` - Login reference for token authentication (see BIG-IP REST docs for details) May be set via the `BIGIP_LOGIN_REF` environment variable.

## Reference

For further information, please visit [the F5bigip provider docs](https://www.pulumi.com/docs/intro/cloud-providers/f5bigip) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/f5bigip).
