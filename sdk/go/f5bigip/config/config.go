// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/pulumi/pulumi/sdk/go/pulumi/config"
)

// Domain name/IP of the BigIP
func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:address")
}

// Login reference for token authentication (see BIG-IP REST docs for details)
func GetLoginRef(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:loginRef")
}

// The user's password
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:password")
}

// Enable to use an external authentication source (LDAP, TACACS, etc)
func GetTokenAuth(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "f5bigip:tokenAuth")
}

// Username with API access to the BigIP
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "f5bigip:username")
}
