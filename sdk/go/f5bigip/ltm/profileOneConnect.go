// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.ProfileOneConnect` Configures a custom profileOneconnect for use by health checks.
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_profile_oneconnect.html.markdown.
type ProfileOneConnect struct {
	pulumi.CustomResourceState

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
	IdleTimeoutOverride pulumi.StringPtrOutput `pulumi:"idleTimeoutOverride"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
	MaxAge pulumi.IntPtrOutput `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
	MaxReuse pulumi.IntPtrOutput `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
	MaxSize pulumi.IntPtrOutput `pulumi:"maxSize"`
	// Name of the profile_oneconnect
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specify if you want to share the pool, default value is "disabled"
	SharePools pulumi.StringPtrOutput `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringPtrOutput `pulumi:"sourceMask"`
}

// NewProfileOneConnect registers a new resource with the given unique name, arguments, and options.
func NewProfileOneConnect(ctx *pulumi.Context,
	name string, args *ProfileOneConnectArgs, opts ...pulumi.ResourceOption) (*ProfileOneConnect, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileOneConnectArgs{}
	}
	var resource ProfileOneConnect
	err := ctx.RegisterResource("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileOneConnect gets an existing ProfileOneConnect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileOneConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileOneConnectState, opts ...pulumi.ResourceOption) (*ProfileOneConnect, error) {
	var resource ProfileOneConnect
	err := ctx.ReadResource("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileOneConnect resources.
type profileOneConnectState struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
	IdleTimeoutOverride *string `pulumi:"idleTimeoutOverride"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
	MaxAge *int `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
	MaxReuse *int `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
	MaxSize *int `pulumi:"maxSize"`
	// Name of the profile_oneconnect
	Name *string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specify if you want to share the pool, default value is "disabled"
	SharePools *string `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask *string `pulumi:"sourceMask"`
}

type ProfileOneConnectState struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
	IdleTimeoutOverride pulumi.StringPtrInput
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
	MaxAge pulumi.IntPtrInput
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
	MaxReuse pulumi.IntPtrInput
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
	MaxSize pulumi.IntPtrInput
	// Name of the profile_oneconnect
	Name pulumi.StringPtrInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specify if you want to share the pool, default value is "disabled"
	SharePools pulumi.StringPtrInput
	// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringPtrInput
}

func (ProfileOneConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileOneConnectState)(nil)).Elem()
}

type profileOneConnectArgs struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
	IdleTimeoutOverride *string `pulumi:"idleTimeoutOverride"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
	MaxAge *int `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
	MaxReuse *int `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
	MaxSize *int `pulumi:"maxSize"`
	// Name of the profile_oneconnect
	Name string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specify if you want to share the pool, default value is "disabled"
	SharePools *string `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask *string `pulumi:"sourceMask"`
}

// The set of arguments for constructing a ProfileOneConnect resource.
type ProfileOneConnectArgs struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
	IdleTimeoutOverride pulumi.StringPtrInput
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
	MaxAge pulumi.IntPtrInput
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
	MaxReuse pulumi.IntPtrInput
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
	MaxSize pulumi.IntPtrInput
	// Name of the profile_oneconnect
	Name pulumi.StringInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specify if you want to share the pool, default value is "disabled"
	SharePools pulumi.StringPtrInput
	// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringPtrInput
}

func (ProfileOneConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileOneConnectArgs)(nil)).Elem()
}

