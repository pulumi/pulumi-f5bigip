// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mypool, err := ltm.NewPool(ctx, "mypool", &ltm.PoolArgs{
// 			Name:              pulumi.String("/Common/test-pool"),
// 			AllowNat:          pulumi.String("yes"),
// 			AllowSnat:         pulumi.String("yes"),
// 			LoadBalancingMode: pulumi.String("round-robin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewPolicy(ctx, "test_policy", &ltm.PolicyArgs{
// 			Name:     pulumi.String("/Common/test-policy"),
// 			Strategy: pulumi.String("first-match"),
// 			Requires: pulumi.StringArray{
// 				pulumi.String("http"),
// 			},
// 			Controls: pulumi.StringArray{
// 				pulumi.String("forwarding"),
// 			},
// 			Rules: ltm.PolicyRuleArray{
// 				&ltm.PolicyRuleArgs{
// 					Name: pulumi.String("rule6"),
// 					Actions: ltm.PolicyRuleActionArray{
// 						&ltm.PolicyRuleActionArgs{
// 							Forward: pulumi.Bool(true),
// 							Pool:    mypool.Name,
// 						},
// 					},
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			mypool,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Specifies the controls
	Controls pulumi.StringArrayOutput `pulumi:"controls"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringOutput `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy pulumi.StringPtrOutput `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires pulumi.StringArrayOutput `pulumi:"requires"`
	// Rules can be applied using the policy
	Rules PolicyRuleArrayOutput `pulumi:"rules"`
	// Specifies the match strategy
	Strategy pulumi.StringPtrOutput `pulumi:"strategy"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Policy
	err := ctx.RegisterResource("f5bigip:ltm/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("f5bigip:ltm/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Specifies the controls
	Controls []string `pulumi:"controls"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name *string `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy *string `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires []string `pulumi:"requires"`
	// Rules can be applied using the policy
	Rules []PolicyRule `pulumi:"rules"`
	// Specifies the match strategy
	Strategy *string `pulumi:"strategy"`
}

type PolicyState struct {
	// Specifies the controls
	Controls pulumi.StringArrayInput
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringPtrInput
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy pulumi.StringPtrInput
	// Specifies the protocol
	Requires pulumi.StringArrayInput
	// Rules can be applied using the policy
	Rules PolicyRuleArrayInput
	// Specifies the match strategy
	Strategy pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Specifies the controls
	Controls []string `pulumi:"controls"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name string `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy *string `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires []string `pulumi:"requires"`
	// Rules can be applied using the policy
	Rules []PolicyRule `pulumi:"rules"`
	// Specifies the match strategy
	Strategy *string `pulumi:"strategy"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies the controls
	Controls pulumi.StringArrayInput
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringInput
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy pulumi.StringPtrInput
	// Specifies the protocol
	Requires pulumi.StringArrayInput
	// Rules can be applied using the policy
	Rules PolicyRuleArrayInput
	// Specifies the match strategy
	Strategy pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i *Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//          PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Policy)(nil))
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//          PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Policy)(nil))
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyT(func(v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}

type PolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Policy {
		return vs[0].([]Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Policy {
		return vs[0].(map[string]Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
