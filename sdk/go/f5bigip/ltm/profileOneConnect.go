// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.ProfileOneConnect` Configures a custom profileOneconnect for use by health checks.
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ltm.NewProfileOneConnect(ctx, "test-oneconnect", &ltm.ProfileOneConnectArgs{
//				Name: pulumi.String("/Common/test-oneconnect"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// BIG-IP LTM oneconnect profiles can be imported using the `name` , e.g.
//
// ```sh
//
//	$ pulumi import f5bigip:ltm/profileOneConnect:ProfileOneConnect test-oneconnect /Common/test-oneconnect
//
// ```
type ProfileOneConnect struct {
	pulumi.CustomResourceState

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
	IdleTimeoutOverride pulumi.StringOutput `pulumi:"idleTimeoutOverride"`
	// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
	LimitType pulumi.StringOutput `pulumi:"limitType"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
	MaxAge pulumi.IntOutput `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
	MaxReuse pulumi.IntOutput `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringOutput `pulumi:"partition"`
	// Specify if you want to share the pool, default value is `disabled`.
	SharePools pulumi.StringOutput `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringOutput `pulumi:"sourceMask"`
}

// NewProfileOneConnect registers a new resource with the given unique name, arguments, and options.
func NewProfileOneConnect(ctx *pulumi.Context,
	name string, args *ProfileOneConnectArgs, opts ...pulumi.ResourceOption) (*ProfileOneConnect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
	IdleTimeoutOverride *string `pulumi:"idleTimeoutOverride"`
	// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
	LimitType *string `pulumi:"limitType"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
	MaxAge *int `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
	MaxReuse *int `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
	MaxSize *int `pulumi:"maxSize"`
	// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
	Name *string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specify if you want to share the pool, default value is `disabled`.
	SharePools *string `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask *string `pulumi:"sourceMask"`
}

type ProfileOneConnectState struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
	IdleTimeoutOverride pulumi.StringPtrInput
	// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
	LimitType pulumi.StringPtrInput
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
	MaxAge pulumi.IntPtrInput
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
	MaxReuse pulumi.IntPtrInput
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
	MaxSize pulumi.IntPtrInput
	// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
	Name pulumi.StringPtrInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specify if you want to share the pool, default value is `disabled`.
	SharePools pulumi.StringPtrInput
	// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringPtrInput
}

func (ProfileOneConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileOneConnectState)(nil)).Elem()
}

type profileOneConnectArgs struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
	IdleTimeoutOverride *string `pulumi:"idleTimeoutOverride"`
	// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
	LimitType *string `pulumi:"limitType"`
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
	MaxAge *int `pulumi:"maxAge"`
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
	MaxReuse *int `pulumi:"maxReuse"`
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
	MaxSize *int `pulumi:"maxSize"`
	// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
	Name string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specify if you want to share the pool, default value is `disabled`.
	SharePools *string `pulumi:"sharePools"`
	// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask *string `pulumi:"sourceMask"`
}

// The set of arguments for constructing a ProfileOneConnect resource.
type ProfileOneConnectArgs struct {
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
	IdleTimeoutOverride pulumi.StringPtrInput
	// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
	LimitType pulumi.StringPtrInput
	// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
	MaxAge pulumi.IntPtrInput
	// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
	MaxReuse pulumi.IntPtrInput
	// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
	MaxSize pulumi.IntPtrInput
	// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
	Name pulumi.StringInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specify if you want to share the pool, default value is `disabled`.
	SharePools pulumi.StringPtrInput
	// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
	SourceMask pulumi.StringPtrInput
}

func (ProfileOneConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileOneConnectArgs)(nil)).Elem()
}

type ProfileOneConnectInput interface {
	pulumi.Input

	ToProfileOneConnectOutput() ProfileOneConnectOutput
	ToProfileOneConnectOutputWithContext(ctx context.Context) ProfileOneConnectOutput
}

func (*ProfileOneConnect) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileOneConnect)(nil)).Elem()
}

func (i *ProfileOneConnect) ToProfileOneConnectOutput() ProfileOneConnectOutput {
	return i.ToProfileOneConnectOutputWithContext(context.Background())
}

func (i *ProfileOneConnect) ToProfileOneConnectOutputWithContext(ctx context.Context) ProfileOneConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOneConnectOutput)
}

// ProfileOneConnectArrayInput is an input type that accepts ProfileOneConnectArray and ProfileOneConnectArrayOutput values.
// You can construct a concrete instance of `ProfileOneConnectArrayInput` via:
//
//	ProfileOneConnectArray{ ProfileOneConnectArgs{...} }
type ProfileOneConnectArrayInput interface {
	pulumi.Input

	ToProfileOneConnectArrayOutput() ProfileOneConnectArrayOutput
	ToProfileOneConnectArrayOutputWithContext(context.Context) ProfileOneConnectArrayOutput
}

type ProfileOneConnectArray []ProfileOneConnectInput

func (ProfileOneConnectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileOneConnect)(nil)).Elem()
}

func (i ProfileOneConnectArray) ToProfileOneConnectArrayOutput() ProfileOneConnectArrayOutput {
	return i.ToProfileOneConnectArrayOutputWithContext(context.Background())
}

func (i ProfileOneConnectArray) ToProfileOneConnectArrayOutputWithContext(ctx context.Context) ProfileOneConnectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOneConnectArrayOutput)
}

// ProfileOneConnectMapInput is an input type that accepts ProfileOneConnectMap and ProfileOneConnectMapOutput values.
// You can construct a concrete instance of `ProfileOneConnectMapInput` via:
//
//	ProfileOneConnectMap{ "key": ProfileOneConnectArgs{...} }
type ProfileOneConnectMapInput interface {
	pulumi.Input

	ToProfileOneConnectMapOutput() ProfileOneConnectMapOutput
	ToProfileOneConnectMapOutputWithContext(context.Context) ProfileOneConnectMapOutput
}

type ProfileOneConnectMap map[string]ProfileOneConnectInput

func (ProfileOneConnectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileOneConnect)(nil)).Elem()
}

func (i ProfileOneConnectMap) ToProfileOneConnectMapOutput() ProfileOneConnectMapOutput {
	return i.ToProfileOneConnectMapOutputWithContext(context.Background())
}

func (i ProfileOneConnectMap) ToProfileOneConnectMapOutputWithContext(ctx context.Context) ProfileOneConnectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOneConnectMapOutput)
}

type ProfileOneConnectOutput struct{ *pulumi.OutputState }

func (ProfileOneConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileOneConnect)(nil)).Elem()
}

func (o ProfileOneConnectOutput) ToProfileOneConnectOutput() ProfileOneConnectOutput {
	return o
}

func (o ProfileOneConnectOutput) ToProfileOneConnectOutputWithContext(ctx context.Context) ProfileOneConnectOutput {
	return o
}

// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
func (o ProfileOneConnectOutput) DefaultsFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.DefaultsFrom }).(pulumi.StringOutput)
}

// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
func (o ProfileOneConnectOutput) IdleTimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.IdleTimeoutOverride }).(pulumi.StringOutput)
}

// Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
func (o ProfileOneConnectOutput) LimitType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.LimitType }).(pulumi.StringOutput)
}

// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
func (o ProfileOneConnectOutput) MaxAge() pulumi.IntOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.IntOutput { return v.MaxAge }).(pulumi.IntOutput)
}

// Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
func (o ProfileOneConnectOutput) MaxReuse() pulumi.IntOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.IntOutput { return v.MaxReuse }).(pulumi.IntOutput)
}

// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
func (o ProfileOneConnectOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
func (o ProfileOneConnectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Displays the administrative partition within which this profile resides
func (o ProfileOneConnectOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.Partition }).(pulumi.StringOutput)
}

// Specify if you want to share the pool, default value is `disabled`.
func (o ProfileOneConnectOutput) SharePools() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.SharePools }).(pulumi.StringOutput)
}

// Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
func (o ProfileOneConnectOutput) SourceMask() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileOneConnect) pulumi.StringOutput { return v.SourceMask }).(pulumi.StringOutput)
}

type ProfileOneConnectArrayOutput struct{ *pulumi.OutputState }

func (ProfileOneConnectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileOneConnect)(nil)).Elem()
}

func (o ProfileOneConnectArrayOutput) ToProfileOneConnectArrayOutput() ProfileOneConnectArrayOutput {
	return o
}

func (o ProfileOneConnectArrayOutput) ToProfileOneConnectArrayOutputWithContext(ctx context.Context) ProfileOneConnectArrayOutput {
	return o
}

func (o ProfileOneConnectArrayOutput) Index(i pulumi.IntInput) ProfileOneConnectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProfileOneConnect {
		return vs[0].([]*ProfileOneConnect)[vs[1].(int)]
	}).(ProfileOneConnectOutput)
}

type ProfileOneConnectMapOutput struct{ *pulumi.OutputState }

func (ProfileOneConnectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileOneConnect)(nil)).Elem()
}

func (o ProfileOneConnectMapOutput) ToProfileOneConnectMapOutput() ProfileOneConnectMapOutput {
	return o
}

func (o ProfileOneConnectMapOutput) ToProfileOneConnectMapOutputWithContext(ctx context.Context) ProfileOneConnectMapOutput {
	return o
}

func (o ProfileOneConnectMapOutput) MapIndex(k pulumi.StringInput) ProfileOneConnectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProfileOneConnect {
		return vs[0].(map[string]*ProfileOneConnect)[vs[1].(string)]
	}).(ProfileOneConnectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileOneConnectInput)(nil)).Elem(), &ProfileOneConnect{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileOneConnectArrayInput)(nil)).Elem(), ProfileOneConnectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileOneConnectMapInput)(nil)).Elem(), ProfileOneConnectMap{})
	pulumi.RegisterOutputType(ProfileOneConnectOutput{})
	pulumi.RegisterOutputType(ProfileOneConnectArrayOutput{})
	pulumi.RegisterOutputType(ProfileOneConnectMapOutput{})
}
