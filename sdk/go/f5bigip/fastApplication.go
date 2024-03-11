// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FastApplication` This resource will create and manage FAST applications on BIG-IP from provided JSON declaration.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := f5bigip.NewFastApplication(ctx, "foo-app", &f5bigip.FastApplicationArgs{
//				FastJson: readFileOrPanic("new_fast_app.json"),
//				Template: pulumi.String("examples/simple_http"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type FastApplication struct {
	pulumi.CustomResourceState

	// A FAST application name.
	//
	// * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	//
	// * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
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
	//
	// * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
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
	return reflect.TypeOf((**FastApplication)(nil)).Elem()
}

func (i *FastApplication) ToFastApplicationOutput() FastApplicationOutput {
	return i.ToFastApplicationOutputWithContext(context.Background())
}

func (i *FastApplication) ToFastApplicationOutputWithContext(ctx context.Context) FastApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationOutput)
}

// FastApplicationArrayInput is an input type that accepts FastApplicationArray and FastApplicationArrayOutput values.
// You can construct a concrete instance of `FastApplicationArrayInput` via:
//
//	FastApplicationArray{ FastApplicationArgs{...} }
type FastApplicationArrayInput interface {
	pulumi.Input

	ToFastApplicationArrayOutput() FastApplicationArrayOutput
	ToFastApplicationArrayOutputWithContext(context.Context) FastApplicationArrayOutput
}

type FastApplicationArray []FastApplicationInput

func (FastApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastApplication)(nil)).Elem()
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
//	FastApplicationMap{ "key": FastApplicationArgs{...} }
type FastApplicationMapInput interface {
	pulumi.Input

	ToFastApplicationMapOutput() FastApplicationMapOutput
	ToFastApplicationMapOutputWithContext(context.Context) FastApplicationMapOutput
}

type FastApplicationMap map[string]FastApplicationInput

func (FastApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastApplication)(nil)).Elem()
}

func (i FastApplicationMap) ToFastApplicationMapOutput() FastApplicationMapOutput {
	return i.ToFastApplicationMapOutputWithContext(context.Background())
}

func (i FastApplicationMap) ToFastApplicationMapOutputWithContext(ctx context.Context) FastApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastApplicationMapOutput)
}

type FastApplicationOutput struct{ *pulumi.OutputState }

func (FastApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastApplication)(nil)).Elem()
}

func (o FastApplicationOutput) ToFastApplicationOutput() FastApplicationOutput {
	return o
}

func (o FastApplicationOutput) ToFastApplicationOutputWithContext(ctx context.Context) FastApplicationOutput {
	return o
}

// A FAST application name.
//
// * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
func (o FastApplicationOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *FastApplication) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
func (o FastApplicationOutput) FastJson() pulumi.StringOutput {
	return o.ApplyT(func(v *FastApplication) pulumi.StringOutput { return v.FastJson }).(pulumi.StringOutput)
}

// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
func (o FastApplicationOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastApplication) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// A FAST tenant name on which you want to manage application.
func (o FastApplicationOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *FastApplication) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

type FastApplicationArrayOutput struct{ *pulumi.OutputState }

func (FastApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastApplication)(nil)).Elem()
}

func (o FastApplicationArrayOutput) ToFastApplicationArrayOutput() FastApplicationArrayOutput {
	return o
}

func (o FastApplicationArrayOutput) ToFastApplicationArrayOutputWithContext(ctx context.Context) FastApplicationArrayOutput {
	return o
}

func (o FastApplicationArrayOutput) Index(i pulumi.IntInput) FastApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FastApplication {
		return vs[0].([]*FastApplication)[vs[1].(int)]
	}).(FastApplicationOutput)
}

type FastApplicationMapOutput struct{ *pulumi.OutputState }

func (FastApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastApplication)(nil)).Elem()
}

func (o FastApplicationMapOutput) ToFastApplicationMapOutput() FastApplicationMapOutput {
	return o
}

func (o FastApplicationMapOutput) ToFastApplicationMapOutputWithContext(ctx context.Context) FastApplicationMapOutput {
	return o
}

func (o FastApplicationMapOutput) MapIndex(k pulumi.StringInput) FastApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FastApplication {
		return vs[0].(map[string]*FastApplication)[vs[1].(string)]
	}).(FastApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastApplicationInput)(nil)).Elem(), &FastApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastApplicationArrayInput)(nil)).Elem(), FastApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastApplicationMapInput)(nil)).Elem(), FastApplicationMap{})
	pulumi.RegisterOutputType(FastApplicationOutput{})
	pulumi.RegisterOutputType(FastApplicationArrayOutput{})
	pulumi.RegisterOutputType(FastApplicationMapOutput{})
}
