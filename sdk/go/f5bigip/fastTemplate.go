// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FastTemplate` This resource will import and create FAST template sets on BIG-IP LTM.
// Template set can be imported from zip archive files on the local disk.
type FastTemplate struct {
	pulumi.CustomResourceState

	// MD5 hash of the zip archive file containing FAST template
	Md5Hash pulumi.StringOutput `pulumi:"md5Hash"`
	// Name of the FAST template set to be created on to BIGIP
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Path to the zip archive file containing FAST template set on Local Disk
	Source pulumi.StringOutput `pulumi:"source"`
}

// NewFastTemplate registers a new resource with the given unique name, arguments, and options.
func NewFastTemplate(ctx *pulumi.Context,
	name string, args *FastTemplateArgs, opts ...pulumi.ResourceOption) (*FastTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Md5Hash == nil {
		return nil, errors.New("invalid value for required argument 'Md5Hash'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource FastTemplate
	err := ctx.RegisterResource("f5bigip:index/fastTemplate:FastTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastTemplate gets an existing FastTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastTemplateState, opts ...pulumi.ResourceOption) (*FastTemplate, error) {
	var resource FastTemplate
	err := ctx.ReadResource("f5bigip:index/fastTemplate:FastTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FastTemplate resources.
type fastTemplateState struct {
	// MD5 hash of the zip archive file containing FAST template
	Md5Hash *string `pulumi:"md5Hash"`
	// Name of the FAST template set to be created on to BIGIP
	Name *string `pulumi:"name"`
	// Path to the zip archive file containing FAST template set on Local Disk
	Source *string `pulumi:"source"`
}

type FastTemplateState struct {
	// MD5 hash of the zip archive file containing FAST template
	Md5Hash pulumi.StringPtrInput
	// Name of the FAST template set to be created on to BIGIP
	Name pulumi.StringPtrInput
	// Path to the zip archive file containing FAST template set on Local Disk
	Source pulumi.StringPtrInput
}

func (FastTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastTemplateState)(nil)).Elem()
}

type fastTemplateArgs struct {
	// MD5 hash of the zip archive file containing FAST template
	Md5Hash string `pulumi:"md5Hash"`
	// Name of the FAST template set to be created on to BIGIP
	Name *string `pulumi:"name"`
	// Path to the zip archive file containing FAST template set on Local Disk
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a FastTemplate resource.
type FastTemplateArgs struct {
	// MD5 hash of the zip archive file containing FAST template
	Md5Hash pulumi.StringInput
	// Name of the FAST template set to be created on to BIGIP
	Name pulumi.StringPtrInput
	// Path to the zip archive file containing FAST template set on Local Disk
	Source pulumi.StringInput
}

func (FastTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastTemplateArgs)(nil)).Elem()
}

type FastTemplateInput interface {
	pulumi.Input

	ToFastTemplateOutput() FastTemplateOutput
	ToFastTemplateOutputWithContext(ctx context.Context) FastTemplateOutput
}

func (*FastTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**FastTemplate)(nil)).Elem()
}

func (i *FastTemplate) ToFastTemplateOutput() FastTemplateOutput {
	return i.ToFastTemplateOutputWithContext(context.Background())
}

func (i *FastTemplate) ToFastTemplateOutputWithContext(ctx context.Context) FastTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastTemplateOutput)
}

// FastTemplateArrayInput is an input type that accepts FastTemplateArray and FastTemplateArrayOutput values.
// You can construct a concrete instance of `FastTemplateArrayInput` via:
//
//	FastTemplateArray{ FastTemplateArgs{...} }
type FastTemplateArrayInput interface {
	pulumi.Input

	ToFastTemplateArrayOutput() FastTemplateArrayOutput
	ToFastTemplateArrayOutputWithContext(context.Context) FastTemplateArrayOutput
}

type FastTemplateArray []FastTemplateInput

func (FastTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastTemplate)(nil)).Elem()
}

func (i FastTemplateArray) ToFastTemplateArrayOutput() FastTemplateArrayOutput {
	return i.ToFastTemplateArrayOutputWithContext(context.Background())
}

func (i FastTemplateArray) ToFastTemplateArrayOutputWithContext(ctx context.Context) FastTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastTemplateArrayOutput)
}

// FastTemplateMapInput is an input type that accepts FastTemplateMap and FastTemplateMapOutput values.
// You can construct a concrete instance of `FastTemplateMapInput` via:
//
//	FastTemplateMap{ "key": FastTemplateArgs{...} }
type FastTemplateMapInput interface {
	pulumi.Input

	ToFastTemplateMapOutput() FastTemplateMapOutput
	ToFastTemplateMapOutputWithContext(context.Context) FastTemplateMapOutput
}

type FastTemplateMap map[string]FastTemplateInput

func (FastTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastTemplate)(nil)).Elem()
}

func (i FastTemplateMap) ToFastTemplateMapOutput() FastTemplateMapOutput {
	return i.ToFastTemplateMapOutputWithContext(context.Background())
}

func (i FastTemplateMap) ToFastTemplateMapOutputWithContext(ctx context.Context) FastTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastTemplateMapOutput)
}

type FastTemplateOutput struct{ *pulumi.OutputState }

func (FastTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastTemplate)(nil)).Elem()
}

func (o FastTemplateOutput) ToFastTemplateOutput() FastTemplateOutput {
	return o
}

func (o FastTemplateOutput) ToFastTemplateOutputWithContext(ctx context.Context) FastTemplateOutput {
	return o
}

// MD5 hash of the zip archive file containing FAST template
func (o FastTemplateOutput) Md5Hash() pulumi.StringOutput {
	return o.ApplyT(func(v *FastTemplate) pulumi.StringOutput { return v.Md5Hash }).(pulumi.StringOutput)
}

// Name of the FAST template set to be created on to BIGIP
func (o FastTemplateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastTemplate) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Path to the zip archive file containing FAST template set on Local Disk
func (o FastTemplateOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FastTemplate) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

type FastTemplateArrayOutput struct{ *pulumi.OutputState }

func (FastTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastTemplate)(nil)).Elem()
}

func (o FastTemplateArrayOutput) ToFastTemplateArrayOutput() FastTemplateArrayOutput {
	return o
}

func (o FastTemplateArrayOutput) ToFastTemplateArrayOutputWithContext(ctx context.Context) FastTemplateArrayOutput {
	return o
}

func (o FastTemplateArrayOutput) Index(i pulumi.IntInput) FastTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FastTemplate {
		return vs[0].([]*FastTemplate)[vs[1].(int)]
	}).(FastTemplateOutput)
}

type FastTemplateMapOutput struct{ *pulumi.OutputState }

func (FastTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastTemplate)(nil)).Elem()
}

func (o FastTemplateMapOutput) ToFastTemplateMapOutput() FastTemplateMapOutput {
	return o
}

func (o FastTemplateMapOutput) ToFastTemplateMapOutputWithContext(ctx context.Context) FastTemplateMapOutput {
	return o
}

func (o FastTemplateMapOutput) MapIndex(k pulumi.StringInput) FastTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FastTemplate {
		return vs[0].(map[string]*FastTemplate)[vs[1].(string)]
	}).(FastTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastTemplateInput)(nil)).Elem(), &FastTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastTemplateArrayInput)(nil)).Elem(), FastTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastTemplateMapInput)(nil)).Elem(), FastTemplateMap{})
	pulumi.RegisterOutputType(FastTemplateOutput{})
	pulumi.RegisterOutputType(FastTemplateArrayOutput{})
	pulumi.RegisterOutputType(FastTemplateMapOutput{})
}
