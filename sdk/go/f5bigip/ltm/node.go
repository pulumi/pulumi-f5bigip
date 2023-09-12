// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `ltm.Node` Manages a node configuration
//
// For resources should be named with their `full path`.The full path is the combination of the `partition + name` of the resource( example: `/Common/my-node` ) or `partition + Direcroty + name` of the resource ( example: `/Common/test/my-node` ).
// When including directory in `full path` we have to make sure it is created in the given partition before using it.
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
//			_, err := ltm.NewNode(ctx, "node", &ltm.NodeArgs{
//				Address:         pulumi.String("192.168.30.1"),
//				ConnectionLimit: pulumi.Int(0),
//				Description:     pulumi.String("Test-Node"),
//				DynamicRatio:    pulumi.Int(1),
//				Fqdn: &ltm.NodeFqdnArgs{
//					AddressFamily: pulumi.String("ipv4"),
//					Interval:      pulumi.String("3000"),
//				},
//				Monitor:   pulumi.String("/Common/icmp"),
//				Name:      pulumi.String("/Common/terraform_node1"),
//				RateLimit: pulumi.String("disabled"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
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
	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session pulumi.StringOutput `pulumi:"session"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	//
	// > *NOTE* Below attributes needs to be configured under fqdn option.
	State pulumi.StringOutput `pulumi:"state"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session *string `pulumi:"session"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	//
	// > *NOTE* Below attributes needs to be configured under fqdn option.
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
	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session pulumi.StringPtrInput
	// Default is "user-up" you can set to "user-down" if you want to disable
	//
	// > *NOTE* Below attributes needs to be configured under fqdn option.
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
	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session *string `pulumi:"session"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	//
	// > *NOTE* Below attributes needs to be configured under fqdn option.
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
	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session pulumi.StringPtrInput
	// Default is "user-up" you can set to "user-down" if you want to disable
	//
	// > *NOTE* Below attributes needs to be configured under fqdn option.
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
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

func (i *Node) ToOutput(ctx context.Context) pulumix.Output[*Node] {
	return pulumix.Output[*Node]{
		OutputState: i.ToNodeOutputWithContext(ctx).OutputState,
	}
}

// NodeArrayInput is an input type that accepts NodeArray and NodeArrayOutput values.
// You can construct a concrete instance of `NodeArrayInput` via:
//
//	NodeArray{ NodeArgs{...} }
type NodeArrayInput interface {
	pulumi.Input

	ToNodeArrayOutput() NodeArrayOutput
	ToNodeArrayOutputWithContext(context.Context) NodeArrayOutput
}

type NodeArray []NodeInput

func (NodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Node)(nil)).Elem()
}

func (i NodeArray) ToNodeArrayOutput() NodeArrayOutput {
	return i.ToNodeArrayOutputWithContext(context.Background())
}

func (i NodeArray) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeArrayOutput)
}

func (i NodeArray) ToOutput(ctx context.Context) pulumix.Output[[]*Node] {
	return pulumix.Output[[]*Node]{
		OutputState: i.ToNodeArrayOutputWithContext(ctx).OutputState,
	}
}

// NodeMapInput is an input type that accepts NodeMap and NodeMapOutput values.
// You can construct a concrete instance of `NodeMapInput` via:
//
//	NodeMap{ "key": NodeArgs{...} }
type NodeMapInput interface {
	pulumi.Input

	ToNodeMapOutput() NodeMapOutput
	ToNodeMapOutputWithContext(context.Context) NodeMapOutput
}

type NodeMap map[string]NodeInput

func (NodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Node)(nil)).Elem()
}

func (i NodeMap) ToNodeMapOutput() NodeMapOutput {
	return i.ToNodeMapOutputWithContext(context.Background())
}

func (i NodeMap) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeMapOutput)
}

func (i NodeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Node] {
	return pulumix.Output[map[string]*Node]{
		OutputState: i.ToNodeMapOutputWithContext(ctx).OutputState,
	}
}

type NodeOutput struct{ *pulumi.OutputState }

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func (o NodeOutput) ToOutput(ctx context.Context) pulumix.Output[*Node] {
	return pulumix.Output[*Node]{
		OutputState: o.OutputState,
	}
}

// IP or hostname of the node
func (o NodeOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Specifies the maximum number of connections allowed for the node or node address.
func (o NodeOutput) ConnectionLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Node) pulumi.IntOutput { return v.ConnectionLimit }).(pulumi.IntOutput)
}

// User-defined description give ltm_node
func (o NodeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the fixed ratio value used for a node during ratio load balancing.
func (o NodeOutput) DynamicRatio() pulumi.IntOutput {
	return o.ApplyT(func(v *Node) pulumi.IntOutput { return v.DynamicRatio }).(pulumi.IntOutput)
}

func (o NodeOutput) Fqdn() NodeFqdnPtrOutput {
	return o.ApplyT(func(v *Node) NodeFqdnPtrOutput { return v.Fqdn }).(NodeFqdnPtrOutput)
}

// specifies the name of the monitor or monitor rule that you want to associate with the node.
func (o NodeOutput) Monitor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.StringPtrOutput { return v.Monitor }).(pulumi.StringPtrOutput)
}

// Name of the node
func (o NodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
func (o NodeOutput) RateLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.RateLimit }).(pulumi.StringOutput)
}

// Sets the ratio number for the node.
func (o NodeOutput) Ratio() pulumi.IntOutput {
	return o.ApplyT(func(v *Node) pulumi.IntOutput { return v.Ratio }).(pulumi.IntOutput)
}

// Enables or disables the node for new sessions. The default value is user-enabled.
func (o NodeOutput) Session() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.Session }).(pulumi.StringOutput)
}

// Default is "user-up" you can set to "user-down" if you want to disable
//
// > *NOTE* Below attributes needs to be configured under fqdn option.
func (o NodeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type NodeArrayOutput struct{ *pulumi.OutputState }

func (NodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Node)(nil)).Elem()
}

func (o NodeArrayOutput) ToNodeArrayOutput() NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Node] {
	return pulumix.Output[[]*Node]{
		OutputState: o.OutputState,
	}
}

func (o NodeArrayOutput) Index(i pulumi.IntInput) NodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Node {
		return vs[0].([]*Node)[vs[1].(int)]
	}).(NodeOutput)
}

type NodeMapOutput struct{ *pulumi.OutputState }

func (NodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Node)(nil)).Elem()
}

func (o NodeMapOutput) ToNodeMapOutput() NodeMapOutput {
	return o
}

func (o NodeMapOutput) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return o
}

func (o NodeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Node] {
	return pulumix.Output[map[string]*Node]{
		OutputState: o.OutputState,
	}
}

func (o NodeMapOutput) MapIndex(k pulumi.StringInput) NodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Node {
		return vs[0].(map[string]*Node)[vs[1].(string)]
	}).(NodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInput)(nil)).Elem(), &Node{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeArrayInput)(nil)).Elem(), NodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeMapInput)(nil)).Elem(), NodeMap{})
	pulumi.RegisterOutputType(NodeOutput{})
	pulumi.RegisterOutputType(NodeArrayOutput{})
	pulumi.RegisterOutputType(NodeMapOutput{})
}
