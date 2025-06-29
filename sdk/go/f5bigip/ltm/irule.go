// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.IRule` Creates iRule on BIG-IP F5 device
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "myirule.tcl",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Loading from a file is the preferred method
//			_, err = ltm.NewIRule(ctx, "rule", &ltm.IRuleArgs{
//				Name:  pulumi.String("/Common/terraform_irule"),
//				Irule: pulumi.String(invokeFile.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ltm.NewIRule(ctx, "rule2", &ltm.IRuleArgs{
//				Name:  pulumi.String("/Common/terraform_irule2"),
//				Irule: pulumi.String("when CLIENT_ACCEPTED {\n     log local0. \"test\"\n   }\n"),
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
// ##myirule.tcl
type IRule struct {
	pulumi.CustomResourceState

	// Body of the iRule
	Irule pulumi.StringOutput `pulumi:"irule"`
	// Name of the iRule
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIRule registers a new resource with the given unique name, arguments, and options.
func NewIRule(ctx *pulumi.Context,
	name string, args *IRuleArgs, opts ...pulumi.ResourceOption) (*IRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Irule == nil {
		return nil, errors.New("invalid value for required argument 'Irule'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IRule
	err := ctx.RegisterResource("f5bigip:ltm/iRule:IRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIRule gets an existing IRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IRuleState, opts ...pulumi.ResourceOption) (*IRule, error) {
	var resource IRule
	err := ctx.ReadResource("f5bigip:ltm/iRule:IRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IRule resources.
type iruleState struct {
	// Body of the iRule
	Irule *string `pulumi:"irule"`
	// Name of the iRule
	Name *string `pulumi:"name"`
}

type IRuleState struct {
	// Body of the iRule
	Irule pulumi.StringPtrInput
	// Name of the iRule
	Name pulumi.StringPtrInput
}

func (IRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iruleState)(nil)).Elem()
}

type iruleArgs struct {
	// Body of the iRule
	Irule string `pulumi:"irule"`
	// Name of the iRule
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a IRule resource.
type IRuleArgs struct {
	// Body of the iRule
	Irule pulumi.StringInput
	// Name of the iRule
	Name pulumi.StringInput
}

func (IRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iruleArgs)(nil)).Elem()
}

type IRuleInput interface {
	pulumi.Input

	ToIRuleOutput() IRuleOutput
	ToIRuleOutputWithContext(ctx context.Context) IRuleOutput
}

func (*IRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IRule)(nil)).Elem()
}

func (i *IRule) ToIRuleOutput() IRuleOutput {
	return i.ToIRuleOutputWithContext(context.Background())
}

func (i *IRule) ToIRuleOutputWithContext(ctx context.Context) IRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IRuleOutput)
}

// IRuleArrayInput is an input type that accepts IRuleArray and IRuleArrayOutput values.
// You can construct a concrete instance of `IRuleArrayInput` via:
//
//	IRuleArray{ IRuleArgs{...} }
type IRuleArrayInput interface {
	pulumi.Input

	ToIRuleArrayOutput() IRuleArrayOutput
	ToIRuleArrayOutputWithContext(context.Context) IRuleArrayOutput
}

type IRuleArray []IRuleInput

func (IRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IRule)(nil)).Elem()
}

func (i IRuleArray) ToIRuleArrayOutput() IRuleArrayOutput {
	return i.ToIRuleArrayOutputWithContext(context.Background())
}

func (i IRuleArray) ToIRuleArrayOutputWithContext(ctx context.Context) IRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IRuleArrayOutput)
}

// IRuleMapInput is an input type that accepts IRuleMap and IRuleMapOutput values.
// You can construct a concrete instance of `IRuleMapInput` via:
//
//	IRuleMap{ "key": IRuleArgs{...} }
type IRuleMapInput interface {
	pulumi.Input

	ToIRuleMapOutput() IRuleMapOutput
	ToIRuleMapOutputWithContext(context.Context) IRuleMapOutput
}

type IRuleMap map[string]IRuleInput

func (IRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IRule)(nil)).Elem()
}

func (i IRuleMap) ToIRuleMapOutput() IRuleMapOutput {
	return i.ToIRuleMapOutputWithContext(context.Background())
}

func (i IRuleMap) ToIRuleMapOutputWithContext(ctx context.Context) IRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IRuleMapOutput)
}

type IRuleOutput struct{ *pulumi.OutputState }

func (IRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IRule)(nil)).Elem()
}

func (o IRuleOutput) ToIRuleOutput() IRuleOutput {
	return o
}

func (o IRuleOutput) ToIRuleOutputWithContext(ctx context.Context) IRuleOutput {
	return o
}

// Body of the iRule
func (o IRuleOutput) Irule() pulumi.StringOutput {
	return o.ApplyT(func(v *IRule) pulumi.StringOutput { return v.Irule }).(pulumi.StringOutput)
}

// Name of the iRule
func (o IRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type IRuleArrayOutput struct{ *pulumi.OutputState }

func (IRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IRule)(nil)).Elem()
}

func (o IRuleArrayOutput) ToIRuleArrayOutput() IRuleArrayOutput {
	return o
}

func (o IRuleArrayOutput) ToIRuleArrayOutputWithContext(ctx context.Context) IRuleArrayOutput {
	return o
}

func (o IRuleArrayOutput) Index(i pulumi.IntInput) IRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IRule {
		return vs[0].([]*IRule)[vs[1].(int)]
	}).(IRuleOutput)
}

type IRuleMapOutput struct{ *pulumi.OutputState }

func (IRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IRule)(nil)).Elem()
}

func (o IRuleMapOutput) ToIRuleMapOutput() IRuleMapOutput {
	return o
}

func (o IRuleMapOutput) ToIRuleMapOutputWithContext(ctx context.Context) IRuleMapOutput {
	return o
}

func (o IRuleMapOutput) MapIndex(k pulumi.StringInput) IRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IRule {
		return vs[0].(map[string]*IRule)[vs[1].(string)]
	}).(IRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IRuleInput)(nil)).Elem(), &IRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*IRuleArrayInput)(nil)).Elem(), IRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IRuleMapInput)(nil)).Elem(), IRuleMap{})
	pulumi.RegisterOutputType(IRuleOutput{})
	pulumi.RegisterOutputType(IRuleArrayOutput{})
	pulumi.RegisterOutputType(IRuleMapOutput{})
}
