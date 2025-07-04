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

// `ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
//
// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-policy`.
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
//			mypool, err := ltm.NewPool(ctx, "mypool", &ltm.PoolArgs{
//				Name:              pulumi.String("/Common/test-pool"),
//				AllowNat:          pulumi.String("yes"),
//				AllowSnat:         pulumi.String("yes"),
//				LoadBalancingMode: pulumi.String("round-robin"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ltm.NewPolicy(ctx, "test-policy", &ltm.PolicyArgs{
//				Name:     pulumi.String("/Common/test-policy"),
//				Strategy: pulumi.String("first-match"),
//				Requires: pulumi.StringArray{
//					pulumi.String("http"),
//				},
//				Controls: pulumi.StringArray{
//					pulumi.String("forwarding"),
//				},
//				Rules: ltm.PolicyRuleArray{
//					&ltm.PolicyRuleArgs{
//						Name: pulumi.String("rule6"),
//						Actions: ltm.PolicyRuleActionArray{
//							&ltm.PolicyRuleActionArgs{
//								Forward:    pulumi.Bool(true),
//								Connection: pulumi.Bool(false),
//								Pool:       mypool.Name,
//							},
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				mypool,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Importing
//
// An existing policy can be imported into this resource by supplying policy Name in `full path` as `id`.
// An example is below:
// ```sh
// $ terraform import bigip_ltm_policy.policy-import-test /Common/policy2
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Specifies the controls
	Controls pulumi.StringArrayOutput `pulumi:"controls"`
	// Specifies descriptive text that identifies the ltm policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringOutput `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
	//
	// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
	PublishedCopy pulumi.StringPtrOutput `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires pulumi.StringArrayOutput `pulumi:"requires"`
	// List of Rules can be applied using the policy. Each rule is block type with following arguments.
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Specifies descriptive text that identifies the ltm policy.
	Description *string `pulumi:"description"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name *string `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
	//
	// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
	PublishedCopy *string `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires []string `pulumi:"requires"`
	// List of Rules can be applied using the policy. Each rule is block type with following arguments.
	Rules []PolicyRule `pulumi:"rules"`
	// Specifies the match strategy
	Strategy *string `pulumi:"strategy"`
}

type PolicyState struct {
	// Specifies the controls
	Controls pulumi.StringArrayInput
	// Specifies descriptive text that identifies the ltm policy.
	Description pulumi.StringPtrInput
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringPtrInput
	// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
	//
	// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
	PublishedCopy pulumi.StringPtrInput
	// Specifies the protocol
	Requires pulumi.StringArrayInput
	// List of Rules can be applied using the policy. Each rule is block type with following arguments.
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
	// Specifies descriptive text that identifies the ltm policy.
	Description *string `pulumi:"description"`
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name string `pulumi:"name"`
	// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
	//
	// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
	PublishedCopy *string `pulumi:"publishedCopy"`
	// Specifies the protocol
	Requires []string `pulumi:"requires"`
	// List of Rules can be applied using the policy. Each rule is block type with following arguments.
	Rules []PolicyRule `pulumi:"rules"`
	// Specifies the match strategy
	Strategy *string `pulumi:"strategy"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies the controls
	Controls pulumi.StringArrayInput
	// Specifies descriptive text that identifies the ltm policy.
	Description pulumi.StringPtrInput
	// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
	Name pulumi.StringInput
	// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
	//
	// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
	PublishedCopy pulumi.StringPtrInput
	// Specifies the protocol
	Requires pulumi.StringArrayInput
	// List of Rules can be applied using the policy. Each rule is block type with following arguments.
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
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
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
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Specifies the controls
func (o PolicyOutput) Controls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Controls }).(pulumi.StringArrayOutput)
}

// Specifies descriptive text that identifies the ltm policy.
func (o PolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
//
// Deprecated: This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
func (o PolicyOutput) PublishedCopy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.PublishedCopy }).(pulumi.StringPtrOutput)
}

// Specifies the protocol
func (o PolicyOutput) Requires() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Requires }).(pulumi.StringArrayOutput)
}

// List of Rules can be applied using the policy. Each rule is block type with following arguments.
func (o PolicyOutput) Rules() PolicyRuleArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyRuleArrayOutput { return v.Rules }).(PolicyRuleArrayOutput)
}

// Specifies the match strategy
func (o PolicyOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Strategy }).(pulumi.StringPtrOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
