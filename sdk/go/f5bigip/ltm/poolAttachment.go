// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PoolAttachment struct {
	pulumi.CustomResourceState

	// Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
	Node pulumi.StringOutput `pulumi:"node"`
	// Name of the pool in /Partition/Name format
	Pool pulumi.StringOutput `pulumi:"pool"`
}

// NewPoolAttachment registers a new resource with the given unique name, arguments, and options.
func NewPoolAttachment(ctx *pulumi.Context,
	name string, args *PoolAttachmentArgs, opts ...pulumi.ResourceOption) (*PoolAttachment, error) {
	if args == nil || args.Node == nil {
		return nil, errors.New("missing required argument 'Node'")
	}
	if args == nil || args.Pool == nil {
		return nil, errors.New("missing required argument 'Pool'")
	}
	if args == nil {
		args = &PoolAttachmentArgs{}
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
	// Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
	Node *string `pulumi:"node"`
	// Name of the pool in /Partition/Name format
	Pool *string `pulumi:"pool"`
}

type PoolAttachmentState struct {
	// Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
	Node pulumi.StringPtrInput
	// Name of the pool in /Partition/Name format
	Pool pulumi.StringPtrInput
}

func (PoolAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolAttachmentState)(nil)).Elem()
}

type poolAttachmentArgs struct {
	// Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
	Node string `pulumi:"node"`
	// Name of the pool in /Partition/Name format
	Pool string `pulumi:"pool"`
}

// The set of arguments for constructing a PoolAttachment resource.
type PoolAttachmentArgs struct {
	// Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
	Node pulumi.StringInput
	// Name of the pool in /Partition/Name format
	Pool pulumi.StringInput
}

func (PoolAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolAttachmentArgs)(nil)).Elem()
}
