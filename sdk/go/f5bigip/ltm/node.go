// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.Node` Manages a node configuration
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_node.html.markdown.
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
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
	// 'disabled'.
	RateLimit pulumi.StringOutput `pulumi:"rateLimit"`
	// Sets the ratio number for the node.
	Ratio pulumi.IntOutput `pulumi:"ratio"`
	// Default is "user-up" you can set to "user-down" if you want to disable
	State pulumi.StringPtrOutput `pulumi:"state"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil || args.Address == nil {
		return nil, errors.New("missing required argument 'Address'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &NodeArgs{}
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
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
	// 'disabled'.
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
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
	// 'disabled'.
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
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
	// 'disabled'.
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
	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
	// 'disabled'.
	RateLimit pulumi.StringPtrInput
	// Sets the ratio number for the node.
	Ratio pulumi.IntPtrInput
	// Default is "user-up" you can set to "user-down" if you want to disable
	State pulumi.StringPtrInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}
