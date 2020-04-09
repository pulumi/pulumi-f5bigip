// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ssl

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ssl.Key` This resource will import SSL certificate key on BIG-IP LTM.
// Certificate key can be imported from certificate key files on the local disk, in PEM format
//
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ssl_key.html.markdown.
type Key struct {
	pulumi.CustomResourceState

	// Content of SSL certificate key present on local Disk
	Content pulumi.StringOutput `pulumi:"content"`
	// Name of SSL Certificate key with .key extension
	Name pulumi.StringOutput `pulumi:"name"`
	// Partition on to SSL Certificate key to be imported
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &KeyArgs{}
	}
	var resource Key
	err := ctx.RegisterResource("f5bigip:ssl/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("f5bigip:ssl/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// Content of SSL certificate key present on local Disk
	Content *string `pulumi:"content"`
	// Name of SSL Certificate key with .key extension
	Name *string `pulumi:"name"`
	// Partition on to SSL Certificate key to be imported
	Partition *string `pulumi:"partition"`
}

type KeyState struct {
	// Content of SSL certificate key present on local Disk
	Content pulumi.StringPtrInput
	// Name of SSL Certificate key with .key extension
	Name pulumi.StringPtrInput
	// Partition on to SSL Certificate key to be imported
	Partition pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Content of SSL certificate key present on local Disk
	Content string `pulumi:"content"`
	// Name of SSL Certificate key with .key extension
	Name string `pulumi:"name"`
	// Partition on to SSL Certificate key to be imported
	Partition *string `pulumi:"partition"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Content of SSL certificate key present on local Disk
	Content pulumi.StringInput
	// Name of SSL Certificate key with .key extension
	Name pulumi.StringInput
	// Partition on to SSL Certificate key to be imported
	Partition pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}
