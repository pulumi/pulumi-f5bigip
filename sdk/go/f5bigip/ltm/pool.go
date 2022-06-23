// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.Pool` Manages F5 BIG-IP LTM pools via iControl REST API.
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
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
// 		monitor, err := ltm.NewMonitor(ctx, "monitor", &ltm.MonitorArgs{
// 			Name:   pulumi.String("/Common/terraform_monitor"),
// 			Parent: pulumi.String("/Common/http"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewPool(ctx, "pool", &ltm.PoolArgs{
// 			Name:                 pulumi.String("/Common/Axiom_Environment_APP1_Pool"),
// 			LoadBalancingMode:    pulumi.String("round-robin"),
// 			MinimumActiveMembers: pulumi.Int(1),
// 			Monitors: pulumi.StringArray{
// 				monitor.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Pool struct {
	pulumi.CustomResourceState

	// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
	AllowNat pulumi.StringOutput `pulumi:"allowNat"`
	// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
	AllowSnat pulumi.StringOutput `pulumi:"allowSnat"`
	// Specifies descriptive text that identifies the pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the load balancing method. The default is Round Robin.
	LoadBalancingMode pulumi.StringOutput `pulumi:"loadBalancingMode"`
	// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
	MinimumActiveMembers pulumi.IntOutput `pulumi:"minimumActiveMembers"`
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayOutput `pulumi:"monitors"`
	// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of times the system tries to contact a new pool member after a passive failure.
	ReselectTries pulumi.IntOutput `pulumi:"reselectTries"`
	// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
	ServiceDownAction pulumi.StringOutput `pulumi:"serviceDownAction"`
	// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
	SlowRampTime pulumi.IntOutput `pulumi:"slowRampTime"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
	// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
	AllowNat *string `pulumi:"allowNat"`
	// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
	AllowSnat *string `pulumi:"allowSnat"`
	// Specifies descriptive text that identifies the pool.
	Description *string `pulumi:"description"`
	// Specifies the load balancing method. The default is Round Robin.
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
	MinimumActiveMembers *int `pulumi:"minimumActiveMembers"`
	// List of monitor names to associate with the pool
	Monitors []string `pulumi:"monitors"`
	// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
	Name *string `pulumi:"name"`
	// Specifies the number of times the system tries to contact a new pool member after a passive failure.
	ReselectTries *int `pulumi:"reselectTries"`
	// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
	ServiceDownAction *string `pulumi:"serviceDownAction"`
	// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
	SlowRampTime *int `pulumi:"slowRampTime"`
}

type PoolState struct {
	// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
	AllowNat pulumi.StringPtrInput
	// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
	AllowSnat pulumi.StringPtrInput
	// Specifies descriptive text that identifies the pool.
	Description pulumi.StringPtrInput
	// Specifies the load balancing method. The default is Round Robin.
	LoadBalancingMode pulumi.StringPtrInput
	// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
	MinimumActiveMembers pulumi.IntPtrInput
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayInput
	// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
	Name pulumi.StringPtrInput
	// Specifies the number of times the system tries to contact a new pool member after a passive failure.
	ReselectTries pulumi.IntPtrInput
	// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
	ServiceDownAction pulumi.StringPtrInput
	// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
	SlowRampTime pulumi.IntPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
	AllowNat *string `pulumi:"allowNat"`
	// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
	AllowSnat *string `pulumi:"allowSnat"`
	// Specifies descriptive text that identifies the pool.
	Description *string `pulumi:"description"`
	// Specifies the load balancing method. The default is Round Robin.
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
	MinimumActiveMembers *int `pulumi:"minimumActiveMembers"`
	// List of monitor names to associate with the pool
	Monitors []string `pulumi:"monitors"`
	// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
	Name string `pulumi:"name"`
	// Specifies the number of times the system tries to contact a new pool member after a passive failure.
	ReselectTries *int `pulumi:"reselectTries"`
	// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
	ServiceDownAction *string `pulumi:"serviceDownAction"`
	// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
	SlowRampTime *int `pulumi:"slowRampTime"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
	AllowNat pulumi.StringPtrInput
	// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
	AllowSnat pulumi.StringPtrInput
	// Specifies descriptive text that identifies the pool.
	Description pulumi.StringPtrInput
	// Specifies the load balancing method. The default is Round Robin.
	LoadBalancingMode pulumi.StringPtrInput
	// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
	MinimumActiveMembers pulumi.IntPtrInput
	// List of monitor names to associate with the pool
	Monitors pulumi.StringArrayInput
	// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
	Name pulumi.StringInput
	// Specifies the number of times the system tries to contact a new pool member after a passive failure.
	ReselectTries pulumi.IntPtrInput
	// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
	ServiceDownAction pulumi.StringPtrInput
	// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
	SlowRampTime pulumi.IntPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//          PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//          PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
func (o PoolOutput) AllowNat() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.AllowNat }).(pulumi.StringOutput)
}

// Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
func (o PoolOutput) AllowSnat() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.AllowSnat }).(pulumi.StringOutput)
}

// Specifies descriptive text that identifies the pool.
func (o PoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the load balancing method. The default is Round Robin.
func (o PoolOutput) LoadBalancingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LoadBalancingMode }).(pulumi.StringOutput)
}

// Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
func (o PoolOutput) MinimumActiveMembers() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.MinimumActiveMembers }).(pulumi.IntOutput)
}

// List of monitor names to associate with the pool
func (o PoolOutput) Monitors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringArrayOutput { return v.Monitors }).(pulumi.StringArrayOutput)
}

// Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of times the system tries to contact a new pool member after a passive failure.
func (o PoolOutput) ReselectTries() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.ReselectTries }).(pulumi.IntOutput)
}

// Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
func (o PoolOutput) ServiceDownAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ServiceDownAction }).(pulumi.StringOutput)
}

// Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
func (o PoolOutput) SlowRampTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.SlowRampTime }).(pulumi.IntOutput)
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].([]*Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].(map[string]*Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PoolInput)(nil)).Elem(), &Pool{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolArrayInput)(nil)).Elem(), PoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolMapInput)(nil)).Elem(), PoolMap{})
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
