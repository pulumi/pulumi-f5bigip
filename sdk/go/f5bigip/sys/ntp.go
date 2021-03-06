// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `sys.Ntp` provides details about a specific bigip
//
// This resource is helpful when configuring NTP server on the BIG-IP.
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/sys"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sys.NewNtp(ctx, "ntp1", &sys.NtpArgs{
// 			Description: pulumi.String("/Common/NTP1"),
// 			Servers: pulumi.StringArray{
// 				pulumi.String("time.facebook.com"),
// 			},
// 			Timezone: pulumi.String("America/Los_Angeles"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Ntp struct {
	pulumi.CustomResourceState

	// Name of the ntp Servers
	Description pulumi.StringOutput `pulumi:"description"`
	// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
	Servers pulumi.StringArrayOutput `pulumi:"servers"`
	// Specifies the time zone that you want to use for the system time.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewNtp registers a new resource with the given unique name, arguments, and options.
func NewNtp(ctx *pulumi.Context,
	name string, args *NtpArgs, opts ...pulumi.ResourceOption) (*Ntp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	var resource Ntp
	err := ctx.RegisterResource("f5bigip:sys/ntp:Ntp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNtp gets an existing Ntp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NtpState, opts ...pulumi.ResourceOption) (*Ntp, error) {
	var resource Ntp
	err := ctx.ReadResource("f5bigip:sys/ntp:Ntp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ntp resources.
type ntpState struct {
	// Name of the ntp Servers
	Description *string `pulumi:"description"`
	// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
	Servers []string `pulumi:"servers"`
	// Specifies the time zone that you want to use for the system time.
	Timezone *string `pulumi:"timezone"`
}

type NtpState struct {
	// Name of the ntp Servers
	Description pulumi.StringPtrInput
	// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
	Servers pulumi.StringArrayInput
	// Specifies the time zone that you want to use for the system time.
	Timezone pulumi.StringPtrInput
}

func (NtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*ntpState)(nil)).Elem()
}

type ntpArgs struct {
	// Name of the ntp Servers
	Description string `pulumi:"description"`
	// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
	Servers []string `pulumi:"servers"`
	// Specifies the time zone that you want to use for the system time.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a Ntp resource.
type NtpArgs struct {
	// Name of the ntp Servers
	Description pulumi.StringInput
	// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
	Servers pulumi.StringArrayInput
	// Specifies the time zone that you want to use for the system time.
	Timezone pulumi.StringPtrInput
}

func (NtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ntpArgs)(nil)).Elem()
}

type NtpInput interface {
	pulumi.Input

	ToNtpOutput() NtpOutput
	ToNtpOutputWithContext(ctx context.Context) NtpOutput
}

func (*Ntp) ElementType() reflect.Type {
	return reflect.TypeOf((*Ntp)(nil))
}

func (i *Ntp) ToNtpOutput() NtpOutput {
	return i.ToNtpOutputWithContext(context.Background())
}

func (i *Ntp) ToNtpOutputWithContext(ctx context.Context) NtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpOutput)
}

func (i *Ntp) ToNtpPtrOutput() NtpPtrOutput {
	return i.ToNtpPtrOutputWithContext(context.Background())
}

func (i *Ntp) ToNtpPtrOutputWithContext(ctx context.Context) NtpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpPtrOutput)
}

type NtpPtrInput interface {
	pulumi.Input

	ToNtpPtrOutput() NtpPtrOutput
	ToNtpPtrOutputWithContext(ctx context.Context) NtpPtrOutput
}

type ntpPtrType NtpArgs

func (*ntpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ntp)(nil))
}

func (i *ntpPtrType) ToNtpPtrOutput() NtpPtrOutput {
	return i.ToNtpPtrOutputWithContext(context.Background())
}

func (i *ntpPtrType) ToNtpPtrOutputWithContext(ctx context.Context) NtpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpPtrOutput)
}

// NtpArrayInput is an input type that accepts NtpArray and NtpArrayOutput values.
// You can construct a concrete instance of `NtpArrayInput` via:
//
//          NtpArray{ NtpArgs{...} }
type NtpArrayInput interface {
	pulumi.Input

	ToNtpArrayOutput() NtpArrayOutput
	ToNtpArrayOutputWithContext(context.Context) NtpArrayOutput
}

type NtpArray []NtpInput

func (NtpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Ntp)(nil))
}

func (i NtpArray) ToNtpArrayOutput() NtpArrayOutput {
	return i.ToNtpArrayOutputWithContext(context.Background())
}

func (i NtpArray) ToNtpArrayOutputWithContext(ctx context.Context) NtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpArrayOutput)
}

// NtpMapInput is an input type that accepts NtpMap and NtpMapOutput values.
// You can construct a concrete instance of `NtpMapInput` via:
//
//          NtpMap{ "key": NtpArgs{...} }
type NtpMapInput interface {
	pulumi.Input

	ToNtpMapOutput() NtpMapOutput
	ToNtpMapOutputWithContext(context.Context) NtpMapOutput
}

type NtpMap map[string]NtpInput

func (NtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Ntp)(nil))
}

func (i NtpMap) ToNtpMapOutput() NtpMapOutput {
	return i.ToNtpMapOutputWithContext(context.Background())
}

func (i NtpMap) ToNtpMapOutputWithContext(ctx context.Context) NtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpMapOutput)
}

type NtpOutput struct {
	*pulumi.OutputState
}

func (NtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ntp)(nil))
}

func (o NtpOutput) ToNtpOutput() NtpOutput {
	return o
}

func (o NtpOutput) ToNtpOutputWithContext(ctx context.Context) NtpOutput {
	return o
}

func (o NtpOutput) ToNtpPtrOutput() NtpPtrOutput {
	return o.ToNtpPtrOutputWithContext(context.Background())
}

func (o NtpOutput) ToNtpPtrOutputWithContext(ctx context.Context) NtpPtrOutput {
	return o.ApplyT(func(v Ntp) *Ntp {
		return &v
	}).(NtpPtrOutput)
}

type NtpPtrOutput struct {
	*pulumi.OutputState
}

func (NtpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ntp)(nil))
}

func (o NtpPtrOutput) ToNtpPtrOutput() NtpPtrOutput {
	return o
}

func (o NtpPtrOutput) ToNtpPtrOutputWithContext(ctx context.Context) NtpPtrOutput {
	return o
}

type NtpArrayOutput struct{ *pulumi.OutputState }

func (NtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ntp)(nil))
}

func (o NtpArrayOutput) ToNtpArrayOutput() NtpArrayOutput {
	return o
}

func (o NtpArrayOutput) ToNtpArrayOutputWithContext(ctx context.Context) NtpArrayOutput {
	return o
}

func (o NtpArrayOutput) Index(i pulumi.IntInput) NtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ntp {
		return vs[0].([]Ntp)[vs[1].(int)]
	}).(NtpOutput)
}

type NtpMapOutput struct{ *pulumi.OutputState }

func (NtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ntp)(nil))
}

func (o NtpMapOutput) ToNtpMapOutput() NtpMapOutput {
	return o
}

func (o NtpMapOutput) ToNtpMapOutputWithContext(ctx context.Context) NtpMapOutput {
	return o
}

func (o NtpMapOutput) MapIndex(k pulumi.StringInput) NtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ntp {
		return vs[0].(map[string]Ntp)[vs[1].(string)]
	}).(NtpOutput)
}

func init() {
	pulumi.RegisterOutputType(NtpOutput{})
	pulumi.RegisterOutputType(NtpPtrOutput{})
	pulumi.RegisterOutputType(NtpArrayOutput{})
	pulumi.RegisterOutputType(NtpMapOutput{})
}
