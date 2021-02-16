// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.Node` Manages a node configuration
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
// 		_, err := ltm.NewNode(ctx, "node", &ltm.NodeArgs{
// 			Address:         pulumi.String("192.168.30.1"),
// 			ConnectionLimit: pulumi.Int(0),
// 			Description:     pulumi.String("Test-Node"),
// 			DynamicRatio:    pulumi.Int(1),
// 			Fqdn: &ltm.NodeFqdnArgs{
// 				AddressFamily: pulumi.String("ipv4"),
// 				Interval:      pulumi.String("3000"),
// 			},
// 			Monitor:   pulumi.String("/Common/icmp"),
// 			Name:      pulumi.String("/Common/terraform_node1"),
// 			RateLimit: pulumi.String("disabled"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Node struct {
	pulumi.CustomResourceState

	// IP or hostname of the node
	Address pulumi.StringOutput `pulumi:"address"`
	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit pulumi.IntOutput `pulumi:"connectionLimit"`
	// User-defined description give ltm_node
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntOutput  `pulumi:"dynamicRatio"`
	Fqdn         NodeFqdnPtrOutput `pulumi:"fqdn"`
	// specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor pulumi.StringPtrOutput `pulumi:"monitor"`
	// Name of the node
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit pulumi.StringOutput `pulumi:"rateLimit"`
	// Sets the ratio number for the node.
	Ratio pulumi.IntOutput `pulumi:"ratio"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	State pulumi.StringPtrOutput `pulumi:"state"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Node
	err := ctx.RegisterResource("f5bigip:ltm/node:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("f5bigip:ltm/node:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
	// IP or hostname of the node
	Address *string `pulumi:"address"`
	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// User-defined description give ltm_node
	Description *string `pulumi:"description"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio *int      `pulumi:"dynamicRatio"`
	Fqdn         *NodeFqdn `pulumi:"fqdn"`
	// specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor *string `pulumi:"monitor"`
	// Name of the node
	Name *string `pulumi:"name"`
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit *string `pulumi:"rateLimit"`
	// Sets the ratio number for the node.
	Ratio *int `pulumi:"ratio"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	State *string `pulumi:"state"`
}

type NodeState struct {
	// IP or hostname of the node
	Address pulumi.StringPtrInput
	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit pulumi.IntPtrInput
	// User-defined description give ltm_node
	Description pulumi.StringPtrInput
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntPtrInput
	Fqdn         NodeFqdnPtrInput
	// specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor pulumi.StringPtrInput
	// Name of the node
	Name pulumi.StringPtrInput
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit pulumi.StringPtrInput
	// Sets the ratio number for the node.
	Ratio pulumi.IntPtrInput
	// Default is "user-up" you can set to "user-down" if you want to disable
	State pulumi.StringPtrInput
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	// IP or hostname of the node
	Address string `pulumi:"address"`
	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// User-defined description give ltm_node
	Description *string `pulumi:"description"`
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio *int      `pulumi:"dynamicRatio"`
	Fqdn         *NodeFqdn `pulumi:"fqdn"`
	// specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor *string `pulumi:"monitor"`
	// Name of the node
	Name string `pulumi:"name"`
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit *string `pulumi:"rateLimit"`
	// Sets the ratio number for the node.
	Ratio *int `pulumi:"ratio"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	// IP or hostname of the node
	Address pulumi.StringInput
	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit pulumi.IntPtrInput
	// User-defined description give ltm_node
	Description pulumi.StringPtrInput
	// Specifies the fixed ratio value used for a node during ratio load balancing.
	DynamicRatio pulumi.IntPtrInput
	Fqdn         NodeFqdnPtrInput
	// specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor pulumi.StringPtrInput
	// Name of the node
	Name pulumi.StringInput
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit pulumi.StringPtrInput
	// Sets the ratio number for the node.
	Ratio pulumi.IntPtrInput
	// Default is "user-up" you can set to "user-down" if you want to disable
	State pulumi.StringPtrInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}

type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(ctx context.Context) NodeOutput
}

func (*Node) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

func (i *Node) ToNodePtrOutput() NodePtrOutput {
	return i.ToNodePtrOutputWithContext(context.Background())
}

func (i *Node) ToNodePtrOutputWithContext(ctx context.Context) NodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePtrOutput)
}

type NodePtrInput interface {
	pulumi.Input

	ToNodePtrOutput() NodePtrOutput
	ToNodePtrOutputWithContext(ctx context.Context) NodePtrOutput
}

type nodePtrType NodeArgs

func (*nodePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil))
}

func (i *nodePtrType) ToNodePtrOutput() NodePtrOutput {
	return i.ToNodePtrOutputWithContext(context.Background())
}

func (i *nodePtrType) ToNodePtrOutputWithContext(ctx context.Context) NodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePtrOutput)
}

// NodeArrayInput is an input type that accepts NodeArray and NodeArrayOutput values.
// You can construct a concrete instance of `NodeArrayInput` via:
//
//          NodeArray{ NodeArgs{...} }
type NodeArrayInput interface {
	pulumi.Input

	ToNodeArrayOutput() NodeArrayOutput
	ToNodeArrayOutputWithContext(context.Context) NodeArrayOutput
}

type NodeArray []NodeInput

func (NodeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Node)(nil))
}

func (i NodeArray) ToNodeArrayOutput() NodeArrayOutput {
	return i.ToNodeArrayOutputWithContext(context.Background())
}

func (i NodeArray) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeArrayOutput)
}

// NodeMapInput is an input type that accepts NodeMap and NodeMapOutput values.
// You can construct a concrete instance of `NodeMapInput` via:
//
//          NodeMap{ "key": NodeArgs{...} }
type NodeMapInput interface {
	pulumi.Input

	ToNodeMapOutput() NodeMapOutput
	ToNodeMapOutputWithContext(context.Context) NodeMapOutput
}

type NodeMap map[string]NodeInput

func (NodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Node)(nil))
}

func (i NodeMap) ToNodeMapOutput() NodeMapOutput {
	return i.ToNodeMapOutputWithContext(context.Background())
}

func (i NodeMap) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeMapOutput)
}

type NodeOutput struct {
	*pulumi.OutputState
}

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func (o NodeOutput) ToNodePtrOutput() NodePtrOutput {
	return o.ToNodePtrOutputWithContext(context.Background())
}

func (o NodeOutput) ToNodePtrOutputWithContext(ctx context.Context) NodePtrOutput {
	return o.ApplyT(func(v Node) *Node {
		return &v
	}).(NodePtrOutput)
}

type NodePtrOutput struct {
	*pulumi.OutputState
}

func (NodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil))
}

func (o NodePtrOutput) ToNodePtrOutput() NodePtrOutput {
	return o
}

func (o NodePtrOutput) ToNodePtrOutputWithContext(ctx context.Context) NodePtrOutput {
	return o
}

type NodeArrayOutput struct{ *pulumi.OutputState }

func (NodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Node)(nil))
}

func (o NodeArrayOutput) ToNodeArrayOutput() NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) Index(i pulumi.IntInput) NodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Node {
		return vs[0].([]Node)[vs[1].(int)]
	}).(NodeOutput)
}

type NodeMapOutput struct{ *pulumi.OutputState }

func (NodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Node)(nil))
}

func (o NodeMapOutput) ToNodeMapOutput() NodeMapOutput {
	return o
}

func (o NodeMapOutput) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return o
}

func (o NodeMapOutput) MapIndex(k pulumi.StringInput) NodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Node {
		return vs[0].(map[string]Node)[vs[1].(string)]
	}).(NodeOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeOutput{})
	pulumi.RegisterOutputType(NodePtrOutput{})
	pulumi.RegisterOutputType(NodeArrayOutput{})
	pulumi.RegisterOutputType(NodeMapOutput{})
}
