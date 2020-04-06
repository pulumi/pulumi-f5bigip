// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.Pool` Manages a pool configuration.
//
// Resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_pool.html.markdown.
type Pool struct {
	pulumi.CustomResourceState

	// Allow NAT
	AllowNat pulumi.StringOutput `pulumi:"allowNat"`
	// Allow SNAT
	AllowSnat pulumi.StringOutput `pulumi:"allowSnat"`
	// Userdefined value to describe the pool
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Possible values: round-robin, ...
	LoadBalancingMode pulumi.StringOutput `pulumi:"loadBalancingMode"`
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayOutput `pulumi:"monitors"`
	// Name of the pool
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of times the system tries to select a new pool member after a failure.
	ReselectTries pulumi.IntOutput `pulumi:"reselectTries"`
	// Possible values: none, reset, reselect, drop
	ServiceDownAction pulumi.StringOutput `pulumi:"serviceDownAction"`
	// Slow ramp time for pool members
	SlowRampTime pulumi.IntOutput `pulumi:"slowRampTime"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PoolArgs{}
	}
	var resource Pool
	err := ctx.RegisterResource("f5bigip:ltm/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("f5bigip:ltm/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// Allow NAT
	AllowNat *string `pulumi:"allowNat"`
	// Allow SNAT
	AllowSnat *string `pulumi:"allowSnat"`
	// Userdefined value to describe the pool
	Description *string `pulumi:"description"`
	// Possible values: round-robin, ...
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// List of monitor names to associate with the pool
	Monitors []string `pulumi:"monitors"`
	// Name of the pool
	Name *string `pulumi:"name"`
	// Number of times the system tries to select a new pool member after a failure.
	ReselectTries *int `pulumi:"reselectTries"`
	// Possible values: none, reset, reselect, drop
	ServiceDownAction *string `pulumi:"serviceDownAction"`
	// Slow ramp time for pool members
	SlowRampTime *int `pulumi:"slowRampTime"`
}

type PoolState struct {
	// Allow NAT
	AllowNat pulumi.StringPtrInput
	// Allow SNAT
	AllowSnat pulumi.StringPtrInput
	// Userdefined value to describe the pool
	Description pulumi.StringPtrInput
	// Possible values: round-robin, ...
	LoadBalancingMode pulumi.StringPtrInput
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayInput
	// Name of the pool
	Name pulumi.StringPtrInput
	// Number of times the system tries to select a new pool member after a failure.
	ReselectTries pulumi.IntPtrInput
	// Possible values: none, reset, reselect, drop
	ServiceDownAction pulumi.StringPtrInput
	// Slow ramp time for pool members
	SlowRampTime pulumi.IntPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Allow NAT
	AllowNat *string `pulumi:"allowNat"`
	// Allow SNAT
	AllowSnat *string `pulumi:"allowSnat"`
	// Userdefined value to describe the pool
	Description *string `pulumi:"description"`
	// Possible values: round-robin, ...
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// List of monitor names to associate with the pool
	Monitors []string `pulumi:"monitors"`
	// Name of the pool
	Name string `pulumi:"name"`
	// Number of times the system tries to select a new pool member after a failure.
	ReselectTries *int `pulumi:"reselectTries"`
	// Possible values: none, reset, reselect, drop
	ServiceDownAction *string `pulumi:"serviceDownAction"`
	// Slow ramp time for pool members
	SlowRampTime *int `pulumi:"slowRampTime"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Allow NAT
	AllowNat pulumi.StringPtrInput
	// Allow SNAT
	AllowSnat pulumi.StringPtrInput
	// Userdefined value to describe the pool
	Description pulumi.StringPtrInput
	// Possible values: round-robin, ...
	LoadBalancingMode pulumi.StringPtrInput
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayInput
	// Name of the pool
	Name pulumi.StringInput
	// Number of times the system tries to select a new pool member after a failure.
	ReselectTries pulumi.IntPtrInput
	// Possible values: none, reset, reselect, drop
	ServiceDownAction pulumi.StringPtrInput
	// Slow ramp time for pool members
	SlowRampTime pulumi.IntPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}
