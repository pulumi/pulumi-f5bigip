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
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil {
		args = &NtpArgs{}
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

func (Ntp) ElementType() reflect.Type {
	return reflect.TypeOf((*Ntp)(nil)).Elem()
}

func (i Ntp) ToNtpOutput() NtpOutput {
	return i.ToNtpOutputWithContext(context.Background())
}

func (i Ntp) ToNtpOutputWithContext(ctx context.Context) NtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NtpOutput)
}

type NtpOutput struct {
	*pulumi.OutputState
}

func (NtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NtpOutput)(nil)).Elem()
}

func (o NtpOutput) ToNtpOutput() NtpOutput {
	return o
}

func (o NtpOutput) ToNtpOutputWithContext(ctx context.Context) NtpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NtpOutput{})
}
