// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.PoolAttachment` Manages nodes membership in pools
//
// ## Example Usage
//
// There are two ways to use ltmPoolAttachment resource, where we can take node reference from ltmNode or we can specify node directly with ip:port/fqdn:port which will also create node and atach to pool.
//
// Pool attachment with node directly taking ip:port/fqdn:port
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
// 			Send: pulumi.String("GET /some/path\n"),
// 			Timeout:  pulumi.Int(999),
// 			Interval: pulumi.Int(998),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := ltm.NewPool(ctx, "pool", &ltm.PoolArgs{
// 			Name:              pulumi.String("/Common/terraform-pool"),
// 			LoadBalancingMode: pulumi.String("round-robin"),
// 			Monitors: pulumi.StringArray{
// 				monitor.Name,
// 			},
// 			AllowSnat: pulumi.String("yes"),
// 			AllowNat:  pulumi.String("yes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewPoolAttachment(ctx, "attachNode", &ltm.PoolAttachmentArgs{
// 			Pool: pool.Name,
// 			Node: pulumi.String("1.1.1.1:80"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Pool attachment with node reference from ltmNode
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		monitor, err := ltm.NewMonitor(ctx, "monitor", &ltm.MonitorArgs{
// 			Name:   pulumi.String("/Common/terraform_monitor"),
// 			Parent: pulumi.String("/Common/http"),
// 			Send: pulumi.String("GET /some/path\n"),
// 			Timeout:  pulumi.Int(999),
// 			Interval: pulumi.Int(998),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := ltm.NewPool(ctx, "pool", &ltm.PoolArgs{
// 			Name:              pulumi.String("/Common/terraform-pool"),
// 			LoadBalancingMode: pulumi.String("round-robin"),
// 			Monitors: pulumi.StringArray{
// 				monitor.Name,
// 			},
// 			AllowSnat: pulumi.String("yes"),
// 			AllowNat:  pulumi.String("yes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		node, err := ltm.NewNode(ctx, "node", &ltm.NodeArgs{
// 			Name:    pulumi.String("/Common/terraform_node"),
// 			Address: pulumi.String("192.168.30.2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewPoolAttachment(ctx, "attachNode", &ltm.PoolAttachmentArgs{
// 			Pool: pool.Name,
// 			Node: node.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", name, ":80"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PoolAttachment struct {
	pulumi.CustomResourceState

	// Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
	ConnectionLimit pulumi.IntOutput `pulumi:"connectionLimit"`
	// Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
	ConnectionRateLimit pulumi.StringOutput `pulumi:"connectionRateLimit"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntOutput `pulumi:"dynamicRatio"`
	// Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
	FqdnAutopopulate pulumi.StringPtrOutput `pulumi:"fqdnAutopopulate"`
	// Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
	Node pulumi.StringOutput `pulumi:"node"`
	// Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
	Pool pulumi.StringOutput `pulumi:"pool"`
	// Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
	PriorityGroup pulumi.IntOutput `pulumi:"priorityGroup"`
	// "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
	Ratio pulumi.IntOutput `pulumi:"ratio"`
}

// NewPoolAttachment registers a new resource with the given unique name, arguments, and options.
func NewPoolAttachment(ctx *pulumi.Context,
	name string, args *PoolAttachmentArgs, opts ...pulumi.ResourceOption) (*PoolAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Node == nil {
		return nil, errors.New("invalid value for required argument 'Node'")
	}
	if args.Pool == nil {
		return nil, errors.New("invalid value for required argument 'Pool'")
	}
	var resource PoolAttachment
	err := ctx.RegisterResource("f5bigip:ltm/poolAttachment:PoolAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPoolAttachment gets an existing PoolAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPoolAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolAttachmentState, opts ...pulumi.ResourceOption) (*PoolAttachment, error) {
	var resource PoolAttachment
	err := ctx.ReadResource("f5bigip:ltm/poolAttachment:PoolAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PoolAttachment resources.
type poolAttachmentState struct {
	// Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
	ConnectionRateLimit *string `pulumi:"connectionRateLimit"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio *int `pulumi:"dynamicRatio"`
	// Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
	FqdnAutopopulate *string `pulumi:"fqdnAutopopulate"`
	// Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
	Node *string `pulumi:"node"`
	// Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
	Pool *string `pulumi:"pool"`
	// Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
	PriorityGroup *int `pulumi:"priorityGroup"`
	// "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
	Ratio *int `pulumi:"ratio"`
}

type PoolAttachmentState struct {
	// Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
	ConnectionLimit pulumi.IntPtrInput
	// Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
	ConnectionRateLimit pulumi.StringPtrInput
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntPtrInput
	// Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
	FqdnAutopopulate pulumi.StringPtrInput
	// Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
	Node pulumi.StringPtrInput
	// Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
	Pool pulumi.StringPtrInput
	// Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
	PriorityGroup pulumi.IntPtrInput
	// "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
	Ratio pulumi.IntPtrInput
}

func (PoolAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolAttachmentState)(nil)).Elem()
}

type poolAttachmentArgs struct {
	// Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
	ConnectionRateLimit *string `pulumi:"connectionRateLimit"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio *int `pulumi:"dynamicRatio"`
	// Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
	FqdnAutopopulate *string `pulumi:"fqdnAutopopulate"`
	// Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
	Node string `pulumi:"node"`
	// Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
	Pool string `pulumi:"pool"`
	// Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
	PriorityGroup *int `pulumi:"priorityGroup"`
	// "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
	Ratio *int `pulumi:"ratio"`
}

// The set of arguments for constructing a PoolAttachment resource.
type PoolAttachmentArgs struct {
	// Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
	ConnectionLimit pulumi.IntPtrInput
	// Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
	ConnectionRateLimit pulumi.StringPtrInput
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntPtrInput
	// Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
	FqdnAutopopulate pulumi.StringPtrInput
	// Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
	Node pulumi.StringInput
	// Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
	Pool pulumi.StringInput
	// Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
	PriorityGroup pulumi.IntPtrInput
	// "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
	Ratio pulumi.IntPtrInput
}

func (PoolAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolAttachmentArgs)(nil)).Elem()
}

type PoolAttachmentInput interface {
	pulumi.Input

	ToPoolAttachmentOutput() PoolAttachmentOutput
	ToPoolAttachmentOutputWithContext(ctx context.Context) PoolAttachmentOutput
}

func (*PoolAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolAttachment)(nil)).Elem()
}

func (i *PoolAttachment) ToPoolAttachmentOutput() PoolAttachmentOutput {
	return i.ToPoolAttachmentOutputWithContext(context.Background())
}

func (i *PoolAttachment) ToPoolAttachmentOutputWithContext(ctx context.Context) PoolAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolAttachmentOutput)
}

// PoolAttachmentArrayInput is an input type that accepts PoolAttachmentArray and PoolAttachmentArrayOutput values.
// You can construct a concrete instance of `PoolAttachmentArrayInput` via:
//
//          PoolAttachmentArray{ PoolAttachmentArgs{...} }
type PoolAttachmentArrayInput interface {
	pulumi.Input

	ToPoolAttachmentArrayOutput() PoolAttachmentArrayOutput
	ToPoolAttachmentArrayOutputWithContext(context.Context) PoolAttachmentArrayOutput
}

type PoolAttachmentArray []PoolAttachmentInput

func (PoolAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PoolAttachment)(nil)).Elem()
}

func (i PoolAttachmentArray) ToPoolAttachmentArrayOutput() PoolAttachmentArrayOutput {
	return i.ToPoolAttachmentArrayOutputWithContext(context.Background())
}

func (i PoolAttachmentArray) ToPoolAttachmentArrayOutputWithContext(ctx context.Context) PoolAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolAttachmentArrayOutput)
}

// PoolAttachmentMapInput is an input type that accepts PoolAttachmentMap and PoolAttachmentMapOutput values.
// You can construct a concrete instance of `PoolAttachmentMapInput` via:
//
//          PoolAttachmentMap{ "key": PoolAttachmentArgs{...} }
type PoolAttachmentMapInput interface {
	pulumi.Input

	ToPoolAttachmentMapOutput() PoolAttachmentMapOutput
	ToPoolAttachmentMapOutputWithContext(context.Context) PoolAttachmentMapOutput
}

type PoolAttachmentMap map[string]PoolAttachmentInput

func (PoolAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PoolAttachment)(nil)).Elem()
}

func (i PoolAttachmentMap) ToPoolAttachmentMapOutput() PoolAttachmentMapOutput {
	return i.ToPoolAttachmentMapOutputWithContext(context.Background())
}

func (i PoolAttachmentMap) ToPoolAttachmentMapOutputWithContext(ctx context.Context) PoolAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolAttachmentMapOutput)
}

type PoolAttachmentOutput struct{ *pulumi.OutputState }

func (PoolAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolAttachment)(nil)).Elem()
}

func (o PoolAttachmentOutput) ToPoolAttachmentOutput() PoolAttachmentOutput {
	return o
}

func (o PoolAttachmentOutput) ToPoolAttachmentOutputWithContext(ctx context.Context) PoolAttachmentOutput {
	return o
}

type PoolAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PoolAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PoolAttachment)(nil)).Elem()
}

func (o PoolAttachmentArrayOutput) ToPoolAttachmentArrayOutput() PoolAttachmentArrayOutput {
	return o
}

func (o PoolAttachmentArrayOutput) ToPoolAttachmentArrayOutputWithContext(ctx context.Context) PoolAttachmentArrayOutput {
	return o
}

func (o PoolAttachmentArrayOutput) Index(i pulumi.IntInput) PoolAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PoolAttachment {
		return vs[0].([]*PoolAttachment)[vs[1].(int)]
	}).(PoolAttachmentOutput)
}

type PoolAttachmentMapOutput struct{ *pulumi.OutputState }

func (PoolAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PoolAttachment)(nil)).Elem()
}

func (o PoolAttachmentMapOutput) ToPoolAttachmentMapOutput() PoolAttachmentMapOutput {
	return o
}

func (o PoolAttachmentMapOutput) ToPoolAttachmentMapOutputWithContext(ctx context.Context) PoolAttachmentMapOutput {
	return o
}

func (o PoolAttachmentMapOutput) MapIndex(k pulumi.StringInput) PoolAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PoolAttachment {
		return vs[0].(map[string]*PoolAttachment)[vs[1].(string)]
	}).(PoolAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PoolAttachmentInput)(nil)).Elem(), &PoolAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolAttachmentArrayInput)(nil)).Elem(), PoolAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolAttachmentMapInput)(nil)).Elem(), PoolAttachmentMap{})
	pulumi.RegisterOutputType(PoolAttachmentOutput{})
	pulumi.RegisterOutputType(PoolAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PoolAttachmentMapOutput{})
}
