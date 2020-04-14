// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.IRule` Creates iRule on BIG-IP F5 device
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
	if args == nil || args.Irule == nil {
		return nil, errors.New("missing required argument 'Irule'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &IRuleArgs{}
	}
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
