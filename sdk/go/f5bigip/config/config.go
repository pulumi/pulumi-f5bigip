// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Domain name/IP of the BigIP
func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:address")
}

// Login reference for token authentication (see BIG-IP REST docs for details)
func GetLoginRef(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:loginRef")
}

// The user's password. Leave empty if using token_value
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:password")
}

// Management Port to connect to Bigip
func GetPort(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:port")
}

// If this flag set to true,sending telemetry data to TEEM will be disabled
func GetTeemDisable(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "f5bigip:teemDisable")
}

// Enable to use an external authentication source (LDAP, TACACS, etc)
func GetTokenAuth(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "f5bigip:tokenAuth")
}

// A token generated outside the provider, in place of password
func GetTokenValue(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:tokenValue")
}

// Valid Trusted Certificate path
func GetTrustedCertPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:trustedCertPath")
}

// Username with API access to the BigIP
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:username")
}

// If set to true, Disables TLS certificate check on BIG-IP. Default : True
func GetValidateCertsDisable(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "f5bigip:validateCertsDisable")
}
