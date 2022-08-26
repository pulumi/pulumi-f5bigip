// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `IpsecPolicy` Manage IPSec policies on a BIG-IP
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-policy)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := f5bigip.NewIpsecPolicy(ctx, "test-policy", &f5bigip.IpsecPolicyArgs{
//				AuthAlgorithm:       pulumi.String("sha1"),
//				Description:         pulumi.String("created by terraform provider"),
//				EncryptAlgorithm:    pulumi.String("3des"),
//				Ipcomp:              pulumi.String("deflate"),
//				Lifetime:            pulumi.Int(3),
//				Mode:                pulumi.String("tunnel"),
//				Name:                pulumi.String("/Common/test-policy"),
//				Protocol:            pulumi.String("esp"),
//				TunnelLocalAddress:  pulumi.String("192.168.1.1"),
//				TunnelRemoteAddress: pulumi.String("10.10.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpsecPolicy struct {
	pulumi.CustomResourceState

	// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
	// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
	AuthAlgorithm pulumi.StringOutput `pulumi:"authAlgorithm"`
	// Description of the IPSec policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
	// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
	EncryptAlgorithm pulumi.StringOutput `pulumi:"encryptAlgorithm"`
	// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
	Ipcomp pulumi.StringOutput `pulumi:"ipcomp"`
	// Specifies the length of time before the IKE security association expires, in kilobytes.
	KbLifetime pulumi.IntOutput `pulumi:"kbLifetime"`
	// Specifies the length of time before the IKE security association expires, in minutes.
	Lifetime pulumi.IntOutput `pulumi:"lifetime"`
	// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
	// modp4096, modp6144, modp8192`
	PerfectForwardSecrecy pulumi.StringOutput `pulumi:"perfectForwardSecrecy"`
	// Specifies the IPsec protocol. Valid choices are: `ah, esp`
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelLocalAddress pulumi.StringOutput `pulumi:"tunnelLocalAddress"`
	// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelRemoteAddress pulumi.StringOutput `pulumi:"tunnelRemoteAddress"`
}

// NewIpsecPolicy registers a new resource with the given unique name, arguments, and options.
func NewIpsecPolicy(ctx *pulumi.Context,
	name string, args *IpsecPolicyArgs, opts ...pulumi.ResourceOption) (*IpsecPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource IpsecPolicy
	err := ctx.RegisterResource("f5bigip:index/ipsecPolicy:IpsecPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsecPolicy gets an existing IpsecPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsecPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsecPolicyState, opts ...pulumi.ResourceOption) (*IpsecPolicy, error) {
	var resource IpsecPolicy
	err := ctx.ReadResource("f5bigip:index/ipsecPolicy:IpsecPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsecPolicy resources.
type ipsecPolicyState struct {
	// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
	// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// Description of the IPSec policy.
	Description *string `pulumi:"description"`
	// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
	// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
	EncryptAlgorithm *string `pulumi:"encryptAlgorithm"`
	// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
	Ipcomp *string `pulumi:"ipcomp"`
	// Specifies the length of time before the IKE security association expires, in kilobytes.
	KbLifetime *int `pulumi:"kbLifetime"`
	// Specifies the length of time before the IKE security association expires, in minutes.
	Lifetime *int `pulumi:"lifetime"`
	// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
	Mode *string `pulumi:"mode"`
	// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
	Name *string `pulumi:"name"`
	// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
	// modp4096, modp6144, modp8192`
	PerfectForwardSecrecy *string `pulumi:"perfectForwardSecrecy"`
	// Specifies the IPsec protocol. Valid choices are: `ah, esp`
	Protocol *string `pulumi:"protocol"`
	// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelLocalAddress *string `pulumi:"tunnelLocalAddress"`
	// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelRemoteAddress *string `pulumi:"tunnelRemoteAddress"`
}

type IpsecPolicyState struct {
	// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
	// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
	AuthAlgorithm pulumi.StringPtrInput
	// Description of the IPSec policy.
	Description pulumi.StringPtrInput
	// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
	// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
	EncryptAlgorithm pulumi.StringPtrInput
	// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
	Ipcomp pulumi.StringPtrInput
	// Specifies the length of time before the IKE security association expires, in kilobytes.
	KbLifetime pulumi.IntPtrInput
	// Specifies the length of time before the IKE security association expires, in minutes.
	Lifetime pulumi.IntPtrInput
	// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
	Mode pulumi.StringPtrInput
	// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
	Name pulumi.StringPtrInput
	// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
	// modp4096, modp6144, modp8192`
	PerfectForwardSecrecy pulumi.StringPtrInput
	// Specifies the IPsec protocol. Valid choices are: `ah, esp`
	Protocol pulumi.StringPtrInput
	// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelLocalAddress pulumi.StringPtrInput
	// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelRemoteAddress pulumi.StringPtrInput
}

func (IpsecPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecPolicyState)(nil)).Elem()
}

type ipsecPolicyArgs struct {
	// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
	// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// Description of the IPSec policy.
	Description *string `pulumi:"description"`
	// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
	// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
	EncryptAlgorithm *string `pulumi:"encryptAlgorithm"`
	// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
	Ipcomp *string `pulumi:"ipcomp"`
	// Specifies the length of time before the IKE security association expires, in kilobytes.
	KbLifetime *int `pulumi:"kbLifetime"`
	// Specifies the length of time before the IKE security association expires, in minutes.
	Lifetime *int `pulumi:"lifetime"`
	// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
	Mode *string `pulumi:"mode"`
	// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
	Name string `pulumi:"name"`
	// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
	// modp4096, modp6144, modp8192`
	PerfectForwardSecrecy *string `pulumi:"perfectForwardSecrecy"`
	// Specifies the IPsec protocol. Valid choices are: `ah, esp`
	Protocol *string `pulumi:"protocol"`
	// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelLocalAddress *string `pulumi:"tunnelLocalAddress"`
	// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelRemoteAddress *string `pulumi:"tunnelRemoteAddress"`
}

// The set of arguments for constructing a IpsecPolicy resource.
type IpsecPolicyArgs struct {
	// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
	// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
	AuthAlgorithm pulumi.StringPtrInput
	// Description of the IPSec policy.
	Description pulumi.StringPtrInput
	// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
	// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
	EncryptAlgorithm pulumi.StringPtrInput
	// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
	Ipcomp pulumi.StringPtrInput
	// Specifies the length of time before the IKE security association expires, in kilobytes.
	KbLifetime pulumi.IntPtrInput
	// Specifies the length of time before the IKE security association expires, in minutes.
	Lifetime pulumi.IntPtrInput
	// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
	Mode pulumi.StringPtrInput
	// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
	Name pulumi.StringInput
	// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
	// modp4096, modp6144, modp8192`
	PerfectForwardSecrecy pulumi.StringPtrInput
	// Specifies the IPsec protocol. Valid choices are: `ah, esp`
	Protocol pulumi.StringPtrInput
	// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelLocalAddress pulumi.StringPtrInput
	// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
	TunnelRemoteAddress pulumi.StringPtrInput
}

func (IpsecPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsecPolicyArgs)(nil)).Elem()
}

type IpsecPolicyInput interface {
	pulumi.Input

	ToIpsecPolicyOutput() IpsecPolicyOutput
	ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput
}

func (*IpsecPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecPolicy)(nil)).Elem()
}

func (i *IpsecPolicy) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return i.ToIpsecPolicyOutputWithContext(context.Background())
}

func (i *IpsecPolicy) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyOutput)
}

// IpsecPolicyArrayInput is an input type that accepts IpsecPolicyArray and IpsecPolicyArrayOutput values.
// You can construct a concrete instance of `IpsecPolicyArrayInput` via:
//
//	IpsecPolicyArray{ IpsecPolicyArgs{...} }
type IpsecPolicyArrayInput interface {
	pulumi.Input

	ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput
	ToIpsecPolicyArrayOutputWithContext(context.Context) IpsecPolicyArrayOutput
}

type IpsecPolicyArray []IpsecPolicyInput

func (IpsecPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return i.ToIpsecPolicyArrayOutputWithContext(context.Background())
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyArrayOutput)
}

// IpsecPolicyMapInput is an input type that accepts IpsecPolicyMap and IpsecPolicyMapOutput values.
// You can construct a concrete instance of `IpsecPolicyMapInput` via:
//
//	IpsecPolicyMap{ "key": IpsecPolicyArgs{...} }
type IpsecPolicyMapInput interface {
	pulumi.Input

	ToIpsecPolicyMapOutput() IpsecPolicyMapOutput
	ToIpsecPolicyMapOutputWithContext(context.Context) IpsecPolicyMapOutput
}

type IpsecPolicyMap map[string]IpsecPolicyInput

func (IpsecPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyMap) ToIpsecPolicyMapOutput() IpsecPolicyMapOutput {
	return i.ToIpsecPolicyMapOutputWithContext(context.Background())
}

func (i IpsecPolicyMap) ToIpsecPolicyMapOutputWithContext(ctx context.Context) IpsecPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyMapOutput)
}

type IpsecPolicyOutput struct{ *pulumi.OutputState }

func (IpsecPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return o
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return o
}

// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
func (o IpsecPolicyOutput) AuthAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.AuthAlgorithm }).(pulumi.StringOutput)
}

// Description of the IPSec policy.
func (o IpsecPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
func (o IpsecPolicyOutput) EncryptAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.EncryptAlgorithm }).(pulumi.StringOutput)
}

// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
func (o IpsecPolicyOutput) Ipcomp() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.Ipcomp }).(pulumi.StringOutput)
}

// Specifies the length of time before the IKE security association expires, in kilobytes.
func (o IpsecPolicyOutput) KbLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.IntOutput { return v.KbLifetime }).(pulumi.IntOutput)
}

// Specifies the length of time before the IKE security association expires, in minutes.
func (o IpsecPolicyOutput) Lifetime() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.IntOutput { return v.Lifetime }).(pulumi.IntOutput)
}

// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
func (o IpsecPolicyOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
func (o IpsecPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
// modp4096, modp6144, modp8192`
func (o IpsecPolicyOutput) PerfectForwardSecrecy() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.PerfectForwardSecrecy }).(pulumi.StringOutput)
}

// Specifies the IPsec protocol. Valid choices are: `ah, esp`
func (o IpsecPolicyOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
func (o IpsecPolicyOutput) TunnelLocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.TunnelLocalAddress }).(pulumi.StringOutput)
}

// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
func (o IpsecPolicyOutput) TunnelRemoteAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsecPolicy) pulumi.StringOutput { return v.TunnelRemoteAddress }).(pulumi.StringOutput)
}

type IpsecPolicyArrayOutput struct{ *pulumi.OutputState }

func (IpsecPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) Index(i pulumi.IntInput) IpsecPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsecPolicy {
		return vs[0].([]*IpsecPolicy)[vs[1].(int)]
	}).(IpsecPolicyOutput)
}

type IpsecPolicyMapOutput struct{ *pulumi.OutputState }

func (IpsecPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyMapOutput) ToIpsecPolicyMapOutput() IpsecPolicyMapOutput {
	return o
}

func (o IpsecPolicyMapOutput) ToIpsecPolicyMapOutputWithContext(ctx context.Context) IpsecPolicyMapOutput {
	return o
}

func (o IpsecPolicyMapOutput) MapIndex(k pulumi.StringInput) IpsecPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsecPolicy {
		return vs[0].(map[string]*IpsecPolicy)[vs[1].(string)]
	}).(IpsecPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecPolicyInput)(nil)).Elem(), &IpsecPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecPolicyArrayInput)(nil)).Elem(), IpsecPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsecPolicyMapInput)(nil)).Elem(), IpsecPolicyMap{})
	pulumi.RegisterOutputType(IpsecPolicyOutput{})
	pulumi.RegisterOutputType(IpsecPolicyArrayOutput{})
	pulumi.RegisterOutputType(IpsecPolicyMapOutput{})
}
