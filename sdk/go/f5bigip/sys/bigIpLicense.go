// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BigIpLicense struct {
	pulumi.CustomResourceState

	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringOutput `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringOutput `pulumi:"registrationKey"`
}

// NewBigIpLicense registers a new resource with the given unique name, arguments, and options.
func NewBigIpLicense(ctx *pulumi.Context,
	name string, args *BigIpLicenseArgs, opts ...pulumi.ResourceOption) (*BigIpLicense, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Command == nil {
		return nil, errors.New("invalid value for required argument 'Command'")
	}
	if args.RegistrationKey == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BigIpLicense
	err := ctx.RegisterResource("f5bigip:sys/bigIpLicense:BigIpLicense", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigIpLicense gets an existing BigIpLicense resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigIpLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigIpLicenseState, opts ...pulumi.ResourceOption) (*BigIpLicense, error) {
	var resource BigIpLicense
	err := ctx.ReadResource("f5bigip:sys/bigIpLicense:BigIpLicense", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigIpLicense resources.
type bigIpLicenseState struct {
	// Tmsh command to execute tmsh commands like install
	Command *string `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey *string `pulumi:"registrationKey"`
}

type BigIpLicenseState struct {
	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringPtrInput
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringPtrInput
}

func (BigIpLicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIpLicenseState)(nil)).Elem()
}

type bigIpLicenseArgs struct {
	// Tmsh command to execute tmsh commands like install
	Command string `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey string `pulumi:"registrationKey"`
}

// The set of arguments for constructing a BigIpLicense resource.
type BigIpLicenseArgs struct {
	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringInput
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringInput
}

func (BigIpLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIpLicenseArgs)(nil)).Elem()
}

type BigIpLicenseInput interface {
	pulumi.Input

	ToBigIpLicenseOutput() BigIpLicenseOutput
	ToBigIpLicenseOutputWithContext(ctx context.Context) BigIpLicenseOutput
}

func (*BigIpLicense) ElementType() reflect.Type {
	return reflect.TypeOf((**BigIpLicense)(nil)).Elem()
}

func (i *BigIpLicense) ToBigIpLicenseOutput() BigIpLicenseOutput {
	return i.ToBigIpLicenseOutputWithContext(context.Background())
}

func (i *BigIpLicense) ToBigIpLicenseOutputWithContext(ctx context.Context) BigIpLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigIpLicenseOutput)
}

func (i *BigIpLicense) ToOutput(ctx context.Context) pulumix.Output[*BigIpLicense] {
	return pulumix.Output[*BigIpLicense]{
		OutputState: i.ToBigIpLicenseOutputWithContext(ctx).OutputState,
	}
}

// BigIpLicenseArrayInput is an input type that accepts BigIpLicenseArray and BigIpLicenseArrayOutput values.
// You can construct a concrete instance of `BigIpLicenseArrayInput` via:
//
//	BigIpLicenseArray{ BigIpLicenseArgs{...} }
type BigIpLicenseArrayInput interface {
	pulumi.Input

	ToBigIpLicenseArrayOutput() BigIpLicenseArrayOutput
	ToBigIpLicenseArrayOutputWithContext(context.Context) BigIpLicenseArrayOutput
}

type BigIpLicenseArray []BigIpLicenseInput

func (BigIpLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BigIpLicense)(nil)).Elem()
}

func (i BigIpLicenseArray) ToBigIpLicenseArrayOutput() BigIpLicenseArrayOutput {
	return i.ToBigIpLicenseArrayOutputWithContext(context.Background())
}

func (i BigIpLicenseArray) ToBigIpLicenseArrayOutputWithContext(ctx context.Context) BigIpLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigIpLicenseArrayOutput)
}

func (i BigIpLicenseArray) ToOutput(ctx context.Context) pulumix.Output[[]*BigIpLicense] {
	return pulumix.Output[[]*BigIpLicense]{
		OutputState: i.ToBigIpLicenseArrayOutputWithContext(ctx).OutputState,
	}
}

// BigIpLicenseMapInput is an input type that accepts BigIpLicenseMap and BigIpLicenseMapOutput values.
// You can construct a concrete instance of `BigIpLicenseMapInput` via:
//
//	BigIpLicenseMap{ "key": BigIpLicenseArgs{...} }
type BigIpLicenseMapInput interface {
	pulumi.Input

	ToBigIpLicenseMapOutput() BigIpLicenseMapOutput
	ToBigIpLicenseMapOutputWithContext(context.Context) BigIpLicenseMapOutput
}

type BigIpLicenseMap map[string]BigIpLicenseInput

func (BigIpLicenseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BigIpLicense)(nil)).Elem()
}

func (i BigIpLicenseMap) ToBigIpLicenseMapOutput() BigIpLicenseMapOutput {
	return i.ToBigIpLicenseMapOutputWithContext(context.Background())
}

func (i BigIpLicenseMap) ToBigIpLicenseMapOutputWithContext(ctx context.Context) BigIpLicenseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigIpLicenseMapOutput)
}

func (i BigIpLicenseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BigIpLicense] {
	return pulumix.Output[map[string]*BigIpLicense]{
		OutputState: i.ToBigIpLicenseMapOutputWithContext(ctx).OutputState,
	}
}

type BigIpLicenseOutput struct{ *pulumi.OutputState }

func (BigIpLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigIpLicense)(nil)).Elem()
}

func (o BigIpLicenseOutput) ToBigIpLicenseOutput() BigIpLicenseOutput {
	return o
}

func (o BigIpLicenseOutput) ToBigIpLicenseOutputWithContext(ctx context.Context) BigIpLicenseOutput {
	return o
}

func (o BigIpLicenseOutput) ToOutput(ctx context.Context) pulumix.Output[*BigIpLicense] {
	return pulumix.Output[*BigIpLicense]{
		OutputState: o.OutputState,
	}
}

// Tmsh command to execute tmsh commands like install
func (o BigIpLicenseOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIpLicense) pulumi.StringOutput { return v.Command }).(pulumi.StringOutput)
}

// A unique Key F5 provides for Licensing BIG-IP
func (o BigIpLicenseOutput) RegistrationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIpLicense) pulumi.StringOutput { return v.RegistrationKey }).(pulumi.StringOutput)
}

type BigIpLicenseArrayOutput struct{ *pulumi.OutputState }

func (BigIpLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BigIpLicense)(nil)).Elem()
}

func (o BigIpLicenseArrayOutput) ToBigIpLicenseArrayOutput() BigIpLicenseArrayOutput {
	return o
}

func (o BigIpLicenseArrayOutput) ToBigIpLicenseArrayOutputWithContext(ctx context.Context) BigIpLicenseArrayOutput {
	return o
}

func (o BigIpLicenseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BigIpLicense] {
	return pulumix.Output[[]*BigIpLicense]{
		OutputState: o.OutputState,
	}
}

func (o BigIpLicenseArrayOutput) Index(i pulumi.IntInput) BigIpLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BigIpLicense {
		return vs[0].([]*BigIpLicense)[vs[1].(int)]
	}).(BigIpLicenseOutput)
}

type BigIpLicenseMapOutput struct{ *pulumi.OutputState }

func (BigIpLicenseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BigIpLicense)(nil)).Elem()
}

func (o BigIpLicenseMapOutput) ToBigIpLicenseMapOutput() BigIpLicenseMapOutput {
	return o
}

func (o BigIpLicenseMapOutput) ToBigIpLicenseMapOutputWithContext(ctx context.Context) BigIpLicenseMapOutput {
	return o
}

func (o BigIpLicenseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BigIpLicense] {
	return pulumix.Output[map[string]*BigIpLicense]{
		OutputState: o.OutputState,
	}
}

func (o BigIpLicenseMapOutput) MapIndex(k pulumi.StringInput) BigIpLicenseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BigIpLicense {
		return vs[0].(map[string]*BigIpLicense)[vs[1].(string)]
	}).(BigIpLicenseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigIpLicenseInput)(nil)).Elem(), &BigIpLicense{})
	pulumi.RegisterInputType(reflect.TypeOf((*BigIpLicenseArrayInput)(nil)).Elem(), BigIpLicenseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BigIpLicenseMapInput)(nil)).Elem(), BigIpLicenseMap{})
	pulumi.RegisterOutputType(BigIpLicenseOutput{})
	pulumi.RegisterOutputType(BigIpLicenseArrayOutput{})
	pulumi.RegisterOutputType(BigIpLicenseMapOutput{})
}
