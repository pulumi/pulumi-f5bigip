// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `IpsecProfile` Manage IPSec Profiles on a BIG-IP
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := f5bigip.NewIpsecProfile(ctx, "azurevWANProfile", &f5bigip.IpsecProfileArgs{
// 			Description:     pulumi.String("mytestipsecprofile"),
// 			Name:            pulumi.String("/Common/Mytestipsecprofile"),
// 			TrafficSelector: pulumi.String("test-trafficselector"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IpsecProfile struct {
	pulumi.CustomResourceState

	// Specifies descriptive text that identifies the IPsec interface tunnel profile.
	Description pulumi.StringOutput `pulumi:"description"`
	// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
	ParentProfile pulumi.StringPtrOutput `pulumi:"parentProfile"`
	// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
	TrafficSelector pulumi.StringOutput `pulumi:"trafficSelector"`
}

// NewIpsecProfile registers a new resource with the given unique name, arguments, and options.
func NewIpsecProfile(ctx *pulumi.Context,
	name string, args *IpsecProfileArgs, opts ...pulumi.ResourceOption) (*IpsecProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource IpsecProfile
	err := ctx.RegisterResource("f5bigip:index/ipsecProfile:IpsecProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsecProfile gets an existing IpsecProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsecProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsecProfileState, opts ...pulumi.ResourceOption) (*IpsecProfile, error) {
	var resource IpsecProfile
	err := ctx.ReadResource("f5bigip:index/ipsecProfile:IpsecProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsecProfile resources.
type ipsecProfileState struct {
	// Specifies descriptive text that identifies the IPsec interface tunnel profile.
	Description *string `pulumi:"description"`
	// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
	Name *string `pulumi:"name"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
	ParentProfile *string `pulumi:"parentProfile"`
	// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
	TrafficSelector *string `pulumi:"trafficSelector"`
}

type IpsecProfileState struct {
	// Specifies descriptive text that identifies the IPsec interface tunnel profile.
	Description pulumi.StringPtrInput
	// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
	Name pulumi.StringPtrInput
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
	ParentProfile pulumi.StringPtrInput
	// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
	TrafficSelector pulumi.StringPtrInput
}

func (IpsecProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecProfileState)(nil)).Elem()
}

type ipsecProfileArgs struct {
	// Specifies descriptive text that identifies the IPsec interface tunnel profile.
	Description *string `pulumi:"description"`
	// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
	Name string `pulumi:"name"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
	ParentProfile *string `pulumi:"parentProfile"`
	// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
	TrafficSelector *string `pulumi:"trafficSelector"`
}

// The set of arguments for constructing a IpsecProfile resource.
type IpsecProfileArgs struct {
	// Specifies descriptive text that identifies the IPsec interface tunnel profile.
	Description pulumi.StringPtrInput
	// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
	Name pulumi.StringInput
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
	ParentProfile pulumi.StringPtrInput
	// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
	TrafficSelector pulumi.StringPtrInput
}

func (IpsecProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecProfileArgs)(nil)).Elem()
}

type IpsecProfileInput interface {
	pulumi.Input

	ToIpsecProfileOutput() IpsecProfileOutput
	ToIpsecProfileOutputWithContext(ctx context.Context) IpsecProfileOutput
}

func (*IpsecProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecProfile)(nil)).Elem()
}

func (i *IpsecProfile) ToIpsecProfileOutput() IpsecProfileOutput {
	return i.ToIpsecProfileOutputWithContext(context.Background())
}

func (i *IpsecProfile) ToIpsecProfileOutputWithContext(ctx context.Context) IpsecProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecProfileOutput)
}

// IpsecProfileArrayInput is an input type that accepts IpsecProfileArray and IpsecProfileArrayOutput values.
// You can construct a concrete instance of `IpsecProfileArrayInput` via:
//
//          IpsecProfileArray{ IpsecProfileArgs{...} }
type IpsecProfileArrayInput interface {
	pulumi.Input

	ToIpsecProfileArrayOutput() IpsecProfileArrayOutput
	ToIpsecProfileArrayOutputWithContext(context.Context) IpsecProfileArrayOutput
}

type IpsecProfileArray []IpsecProfileInput

func (IpsecProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsecProfile)(nil)).Elem()
}

func (i IpsecProfileArray) ToIpsecProfileArrayOutput() IpsecProfileArrayOutput {
	return i.ToIpsecProfileArrayOutputWithContext(context.Background())
}

func (i IpsecProfileArray) ToIpsecProfileArrayOutputWithContext(ctx context.Context) IpsecProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecProfileArrayOutput)
}

// IpsecProfileMapInput is an input type that accepts IpsecProfileMap and IpsecProfileMapOutput values.
// You can construct a concrete instance of `IpsecProfileMapInput` via:
//
//          IpsecProfileMap{ "key": IpsecProfileArgs{...} }
type IpsecProfileMapInput interface {
	pulumi.Input

	ToIpsecProfileMapOutput() IpsecProfileMapOutput
	ToIpsecProfileMapOutputWithContext(context.Context) IpsecProfileMapOutput
}

type IpsecProfileMap map[string]IpsecProfileInput

func (IpsecProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsecProfile)(nil)).Elem()
}

func (i IpsecProfileMap) ToIpsecProfileMapOutput() IpsecProfileMapOutput {
	return i.ToIpsecProfileMapOutputWithContext(context.Background())
}

func (i IpsecProfileMap) ToIpsecProfileMapOutputWithContext(ctx context.Context) IpsecProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecProfileMapOutput)
}

type IpsecProfileOutput struct{ *pulumi.OutputState }

func (IpsecProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecProfile)(nil)).Elem()
}

func (o IpsecProfileOutput) ToIpsecProfileOutput() IpsecProfileOutput {
	return o
}

func (o IpsecProfileOutput) ToIpsecProfileOutputWithContext(ctx context.Context) IpsecProfileOutput {
	return o
}

type IpsecProfileArrayOutput struct{ *pulumi.OutputState }

func (IpsecProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsecProfile)(nil)).Elem()
}

func (o IpsecProfileArrayOutput) ToIpsecProfileArrayOutput() IpsecProfileArrayOutput {
	return o
}

func (o IpsecProfileArrayOutput) ToIpsecProfileArrayOutputWithContext(ctx context.Context) IpsecProfileArrayOutput {
	return o
}

func (o IpsecProfileArrayOutput) Index(i pulumi.IntInput) IpsecProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsecProfile {
		return vs[0].([]*IpsecProfile)[vs[1].(int)]
	}).(IpsecProfileOutput)
}

type IpsecProfileMapOutput struct{ *pulumi.OutputState }

func (IpsecProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsecProfile)(nil)).Elem()
}

func (o IpsecProfileMapOutput) ToIpsecProfileMapOutput() IpsecProfileMapOutput {
	return o
}

func (o IpsecProfileMapOutput) ToIpsecProfileMapOutputWithContext(ctx context.Context) IpsecProfileMapOutput {
	return o
}

func (o IpsecProfileMapOutput) MapIndex(k pulumi.StringInput) IpsecProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsecProfile {
		return vs[0].(map[string]*IpsecProfile)[vs[1].(string)]
	}).(IpsecProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecProfileInput)(nil)).Elem(), &IpsecProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecProfileArrayInput)(nil)).Elem(), IpsecProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecProfileMapInput)(nil)).Elem(), IpsecProfileMap{})
	pulumi.RegisterOutputType(IpsecProfileOutput{})
	pulumi.RegisterOutputType(IpsecProfileArrayOutput{})
	pulumi.RegisterOutputType(IpsecProfileMapOutput{})
}
