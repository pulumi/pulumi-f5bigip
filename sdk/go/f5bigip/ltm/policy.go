// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.Policy` Configures Virtual Server
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewPolicy(ctx, "test_policy", &ltm.PolicyArgs{
// 			Name:     pulumi.String("test-policy"),
// 			Strategy: pulumi.String("/Common/first-match"),
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
// 							Pool:    pulumi.Any(bigip_ltm_pool.Pool.Name),
// 						},
// 					},
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			bigip_ltm_pool.Mypool,
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
	// Name of the Policy
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
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PolicyArgs{}
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
	// Name of the Policy
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
	// Name of the Policy
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
	// Name of the Policy
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
	// Name of the Policy
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
