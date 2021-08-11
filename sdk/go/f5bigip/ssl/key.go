// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ssl.Key` This resource will import SSL certificate key on BIG-IP LTM.
// Certificate key can be imported from certificate key files on the local disk, in PEM format
type Key struct {
	pulumi.CustomResourceState

	// Content of SSL certificate key present on local Disk
	Content pulumi.StringOutput `pulumi:"content"`
	// Full Path Name of ssl key
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// Name of the SSL Certificate key to be Imported on to BIGIP
	Name pulumi.StringOutput `pulumi:"name"`
	// Partition of ssl certificate key
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
	// Full Path Name of ssl key
	FullPath *string `pulumi:"fullPath"`
	// Name of the SSL Certificate key to be Imported on to BIGIP
	Name *string `pulumi:"name"`
	// Partition of ssl certificate key
	Partition *string `pulumi:"partition"`
}

type KeyState struct {
	// Content of SSL certificate key present on local Disk
	Content pulumi.StringPtrInput
	// Full Path Name of ssl key
	FullPath pulumi.StringPtrInput
	// Name of the SSL Certificate key to be Imported on to BIGIP
	Name pulumi.StringPtrInput
	// Partition of ssl certificate key
	Partition pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Content of SSL certificate key present on local Disk
	Content string `pulumi:"content"`
	// Full Path Name of ssl key
	FullPath *string `pulumi:"fullPath"`
	// Name of the SSL Certificate key to be Imported on to BIGIP
	Name string `pulumi:"name"`
	// Partition of ssl certificate key
	Partition *string `pulumi:"partition"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Content of SSL certificate key present on local Disk
	Content pulumi.StringInput
	// Full Path Name of ssl key
	FullPath pulumi.StringPtrInput
	// Name of the SSL Certificate key to be Imported on to BIGIP
	Name pulumi.StringInput
	// Partition of ssl certificate key
	Partition pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

func (i *Key) ToKeyPtrOutput() KeyPtrOutput {
	return i.ToKeyPtrOutputWithContext(context.Background())
}

func (i *Key) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPtrOutput)
}

type KeyPtrInput interface {
	pulumi.Input

	ToKeyPtrOutput() KeyPtrOutput
	ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput
}

type keyPtrType KeyArgs

func (*keyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil))
}

func (i *keyPtrType) ToKeyPtrOutput() KeyPtrOutput {
	return i.ToKeyPtrOutputWithContext(context.Background())
}

func (i *keyPtrType) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPtrOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//          KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Key)(nil))
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//          KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Key)(nil))
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct {
	*pulumi.OutputState
}

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) ToKeyPtrOutput() KeyPtrOutput {
	return o.ToKeyPtrOutputWithContext(context.Background())
}

func (o KeyOutput) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return o.ApplyT(func(v Key) *Key {
		return &v
	}).(KeyPtrOutput)
}

type KeyPtrOutput struct {
	*pulumi.OutputState
}

func (KeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil))
}

func (o KeyPtrOutput) ToKeyPtrOutput() KeyPtrOutput {
	return o
}

func (o KeyPtrOutput) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return o
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Key)(nil))
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Key {
		return vs[0].([]Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Key)(nil))
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Key {
		return vs[0].(map[string]Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyPtrOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
