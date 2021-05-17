// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FastApplication` This resource will create and manage FAST applications on BIG-IP from provided JSON declaration.
type FastApplication struct {
	pulumi.CustomResourceState

	// A FAST application name.
	Application pulumi.StringOutput `pulumi:"application"`
	// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
	FastJson pulumi.StringOutput `pulumi:"fastJson"`
	// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// A FAST tenant name on which you want to manage application.
	Tenant pulumi.StringOutput `pulumi:"tenant"`
}

// NewFastApplication registers a new resource with the given unique name, arguments, and options.
func NewFastApplication(ctx *pulumi.Context,
	name string, args *FastApplicationArgs, opts ...pulumi.ResourceOption) (*FastApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FastJson == nil {
		return nil, errors.New("invalid value for required argument 'FastJson'")
	}
	var resource FastApplication
	err := ctx.RegisterResource("f5bigip:index/fastApplication:FastApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastApplication gets an existing FastApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastApplicationState, opts ...pulumi.ResourceOption) (*FastApplication, error) {
	var resource FastApplication
	err := ctx.ReadResource("f5bigip:index/fastApplication:FastApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FastApplication resources.
type fastApplicationState struct {
	// A FAST application name.
	Application *string `pulumi:"application"`
	// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
	FastJson *string `pulumi:"fastJson"`
	// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
	Template *string `pulumi:"template"`
	// A FAST tenant name on which you want to manage application.
	Tenant *string `pulumi:"tenant"`
}

type FastApplicationState struct {
	// A FAST application name.
	Application pulumi.StringPtrInput
	// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
	FastJson pulumi.StringPtrInput
	// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
	Template pulumi.StringPtrInput
	// A FAST tenant name on which you want to manage application.
	Tenant pulumi.StringPtrInput
}

func (FastApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastApplicationState)(nil)).Elem()
}

type fastApplicationArgs struct {
	// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
	FastJson string `pulumi:"fastJson"`
	// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a FastApplication resource.
type FastApplicationArgs struct {
	// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
	FastJson pulumi.StringInput
	// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
	Template pulumi.StringPtrInput
}

func (FastApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastApplicationArgs)(nil)).Elem()
}

type FastApplicationInput interface {
	pulumi.Input

	ToFastApplicationOutput() FastApplicationOutput
	ToFastApplicationOutputWithContext(ctx context.Context) FastApplicationOutput
}

func (*FastApplication) ElementType() reflect.Type {
	return reflect.TypeOf((*FastApplication)(nil))
}

func (i *FastApplication) ToFastApplicationOutput() FastApplicationOutput {
	return i.ToFastApplicationOutputWithContext(context.Background())
}

func (i *FastApplication) ToFastApplicationOutputWithContext(ctx context.Context) FastApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationOutput)
}

func (i *FastApplication) ToFastApplicationPtrOutput() FastApplicationPtrOutput {
	return i.ToFastApplicationPtrOutputWithContext(context.Background())
}

func (i *FastApplication) ToFastApplicationPtrOutputWithContext(ctx context.Context) FastApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationPtrOutput)
}

type FastApplicationPtrInput interface {
	pulumi.Input

	ToFastApplicationPtrOutput() FastApplicationPtrOutput
	ToFastApplicationPtrOutputWithContext(ctx context.Context) FastApplicationPtrOutput
}

type fastApplicationPtrType FastApplicationArgs

func (*fastApplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FastApplication)(nil))
}

func (i *fastApplicationPtrType) ToFastApplicationPtrOutput() FastApplicationPtrOutput {
	return i.ToFastApplicationPtrOutputWithContext(context.Background())
}

func (i *fastApplicationPtrType) ToFastApplicationPtrOutputWithContext(ctx context.Context) FastApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationPtrOutput)
}

// FastApplicationArrayInput is an input type that accepts FastApplicationArray and FastApplicationArrayOutput values.
// You can construct a concrete instance of `FastApplicationArrayInput` via:
//
//          FastApplicationArray{ FastApplicationArgs{...} }
type FastApplicationArrayInput interface {
	pulumi.Input

	ToFastApplicationArrayOutput() FastApplicationArrayOutput
	ToFastApplicationArrayOutputWithContext(context.Context) FastApplicationArrayOutput
}

type FastApplicationArray []FastApplicationInput

func (FastApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FastApplication)(nil))
}

func (i FastApplicationArray) ToFastApplicationArrayOutput() FastApplicationArrayOutput {
	return i.ToFastApplicationArrayOutputWithContext(context.Background())
}

func (i FastApplicationArray) ToFastApplicationArrayOutputWithContext(ctx context.Context) FastApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationArrayOutput)
}

// FastApplicationMapInput is an input type that accepts FastApplicationMap and FastApplicationMapOutput values.
// You can construct a concrete instance of `FastApplicationMapInput` via:
//
//          FastApplicationMap{ "key": FastApplicationArgs{...} }
type FastApplicationMapInput interface {
	pulumi.Input

	ToFastApplicationMapOutput() FastApplicationMapOutput
	ToFastApplicationMapOutputWithContext(context.Context) FastApplicationMapOutput
}

type FastApplicationMap map[string]FastApplicationInput

func (FastApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FastApplication)(nil))
}

func (i FastApplicationMap) ToFastApplicationMapOutput() FastApplicationMapOutput {
	return i.ToFastApplicationMapOutputWithContext(context.Background())
}

func (i FastApplicationMap) ToFastApplicationMapOutputWithContext(ctx context.Context) FastApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationMapOutput)
}

type FastApplicationOutput struct {
	*pulumi.OutputState
}

func (FastApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FastApplication)(nil))
}

func (o FastApplicationOutput) ToFastApplicationOutput() FastApplicationOutput {
	return o
}

func (o FastApplicationOutput) ToFastApplicationOutputWithContext(ctx context.Context) FastApplicationOutput {
	return o
}

func (o FastApplicationOutput) ToFastApplicationPtrOutput() FastApplicationPtrOutput {
	return o.ToFastApplicationPtrOutputWithContext(context.Background())
}

func (o FastApplicationOutput) ToFastApplicationPtrOutputWithContext(ctx context.Context) FastApplicationPtrOutput {
	return o.ApplyT(func(v FastApplication) *FastApplication {
		return &v
	}).(FastApplicationPtrOutput)
}

type FastApplicationPtrOutput struct {
	*pulumi.OutputState
}

func (FastApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastApplication)(nil))
}

func (o FastApplicationPtrOutput) ToFastApplicationPtrOutput() FastApplicationPtrOutput {
	return o
}

func (o FastApplicationPtrOutput) ToFastApplicationPtrOutputWithContext(ctx context.Context) FastApplicationPtrOutput {
	return o
}

type FastApplicationArrayOutput struct{ *pulumi.OutputState }

func (FastApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FastApplication)(nil))
}

func (o FastApplicationArrayOutput) ToFastApplicationArrayOutput() FastApplicationArrayOutput {
	return o
}

func (o FastApplicationArrayOutput) ToFastApplicationArrayOutputWithContext(ctx context.Context) FastApplicationArrayOutput {
	return o
}

func (o FastApplicationArrayOutput) Index(i pulumi.IntInput) FastApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FastApplication {
		return vs[0].([]FastApplication)[vs[1].(int)]
	}).(FastApplicationOutput)
}

type FastApplicationMapOutput struct{ *pulumi.OutputState }

func (FastApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FastApplication)(nil))
}

func (o FastApplicationMapOutput) ToFastApplicationMapOutput() FastApplicationMapOutput {
	return o
}

func (o FastApplicationMapOutput) ToFastApplicationMapOutputWithContext(ctx context.Context) FastApplicationMapOutput {
	return o
}

func (o FastApplicationMapOutput) MapIndex(k pulumi.StringInput) FastApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FastApplication {
		return vs[0].(map[string]FastApplication)[vs[1].(string)]
	}).(FastApplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(FastApplicationOutput{})
	pulumi.RegisterOutputType(FastApplicationPtrOutput{})
	pulumi.RegisterOutputType(FastApplicationArrayOutput{})
	pulumi.RegisterOutputType(FastApplicationMapOutput{})
}
