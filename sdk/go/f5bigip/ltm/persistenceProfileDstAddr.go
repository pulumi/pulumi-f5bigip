// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures a cookie persistence profile
//
// ## Example
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
//			_, err := ltm.NewPersistenceProfileDstAddr(ctx, "dstaddr", &ltm.PersistenceProfileDstAddrArgs{
//				DefaultsFrom:        pulumi.String("/Common/dest_addr"),
//				HashAlgorithm:       pulumi.String("carp"),
//				Mask:                pulumi.String("255.255.255.255"),
//				MatchAcrossPools:    pulumi.String("enabled"),
//				MatchAcrossServices: pulumi.String("enabled"),
//				MatchAcrossVirtuals: pulumi.String("enabled"),
//				Mirror:              pulumi.String("enabled"),
//				Name:                pulumi.String("/Common/terraform_ppdstaddr"),
//				OverrideConnLimit:   pulumi.String("enabled"),
//				Timeout:             pulumi.Int(3600),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Reference
//
// `name` - (Required) Name of the virtual address
//
// `defaultsFrom` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.
//
// `matchAcrossPools` (Optional) (enabled or disabled) match across pools with given persistence record
//
// `matchAcrossServices` (Optional) (enabled or disabled) match across services with given persistence record
//
// `matchAcrossVirtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
//
// `mirror` (Optional) (enabled or disabled) mirror persistence record
//
// `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
//
// `overrideConnLimit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
//
// ## Importing
//
// An dest-addr persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
// An example is below:
// ```sh
// $ terraform import bigip_ltm_persistence_profile_dstaddr.dstaddr "/Common/terraform_ppdstaddr"
// ```
type PersistenceProfileDstAddr struct {
	pulumi.CustomResourceState

	AppService pulumi.StringOutput `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringOutput `pulumi:"hashAlgorithm"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringOutput `pulumi:"mask"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringOutput `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringOutput `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringOutput `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror pulumi.StringOutput `pulumi:"mirror"`
	// Name of the persistence profile
	Name pulumi.StringOutput `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringOutput `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout pulumi.IntOutput `pulumi:"timeout"`
}

// NewPersistenceProfileDstAddr registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileDstAddr(ctx *pulumi.Context,
	name string, args *PersistenceProfileDstAddrArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileDstAddr, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultsFrom == nil {
		return nil, errors.New("invalid value for required argument 'DefaultsFrom'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PersistenceProfileDstAddr
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileDstAddr gets an existing PersistenceProfileDstAddr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileDstAddr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileDstAddrState, opts ...pulumi.ResourceOption) (*PersistenceProfileDstAddr, error) {
	var resource PersistenceProfileDstAddr
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileDstAddr resources.
type persistenceProfileDstAddrState struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask *string `pulumi:"mask"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals *string `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror *string `pulumi:"mirror"`
	// Name of the persistence profile
	Name *string `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit *string `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout *int `pulumi:"timeout"`
}

type PersistenceProfileDstAddrState struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringPtrInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringPtrInput
	// To enable _ disable
	Mirror pulumi.StringPtrInput
	// Name of the persistence profile
	Name pulumi.StringPtrInput
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringPtrInput
	// Timeout for persistence of the session
	Timeout pulumi.IntPtrInput
}

func (PersistenceProfileDstAddrState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileDstAddrState)(nil)).Elem()
}

type persistenceProfileDstAddrArgs struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask *string `pulumi:"mask"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals *string `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror *string `pulumi:"mirror"`
	// Name of the persistence profile
	Name string `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit *string `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a PersistenceProfileDstAddr resource.
type PersistenceProfileDstAddrArgs struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringPtrInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringPtrInput
	// To enable _ disable
	Mirror pulumi.StringPtrInput
	// Name of the persistence profile
	Name pulumi.StringInput
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringPtrInput
	// Timeout for persistence of the session
	Timeout pulumi.IntPtrInput
}

func (PersistenceProfileDstAddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileDstAddrArgs)(nil)).Elem()
}

type PersistenceProfileDstAddrInput interface {
	pulumi.Input

	ToPersistenceProfileDstAddrOutput() PersistenceProfileDstAddrOutput
	ToPersistenceProfileDstAddrOutputWithContext(ctx context.Context) PersistenceProfileDstAddrOutput
}

func (*PersistenceProfileDstAddr) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileDstAddr)(nil)).Elem()
}

func (i *PersistenceProfileDstAddr) ToPersistenceProfileDstAddrOutput() PersistenceProfileDstAddrOutput {
	return i.ToPersistenceProfileDstAddrOutputWithContext(context.Background())
}

func (i *PersistenceProfileDstAddr) ToPersistenceProfileDstAddrOutputWithContext(ctx context.Context) PersistenceProfileDstAddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileDstAddrOutput)
}

// PersistenceProfileDstAddrArrayInput is an input type that accepts PersistenceProfileDstAddrArray and PersistenceProfileDstAddrArrayOutput values.
// You can construct a concrete instance of `PersistenceProfileDstAddrArrayInput` via:
//
//	PersistenceProfileDstAddrArray{ PersistenceProfileDstAddrArgs{...} }
type PersistenceProfileDstAddrArrayInput interface {
	pulumi.Input

	ToPersistenceProfileDstAddrArrayOutput() PersistenceProfileDstAddrArrayOutput
	ToPersistenceProfileDstAddrArrayOutputWithContext(context.Context) PersistenceProfileDstAddrArrayOutput
}

type PersistenceProfileDstAddrArray []PersistenceProfileDstAddrInput

func (PersistenceProfileDstAddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileDstAddr)(nil)).Elem()
}

func (i PersistenceProfileDstAddrArray) ToPersistenceProfileDstAddrArrayOutput() PersistenceProfileDstAddrArrayOutput {
	return i.ToPersistenceProfileDstAddrArrayOutputWithContext(context.Background())
}

func (i PersistenceProfileDstAddrArray) ToPersistenceProfileDstAddrArrayOutputWithContext(ctx context.Context) PersistenceProfileDstAddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileDstAddrArrayOutput)
}

// PersistenceProfileDstAddrMapInput is an input type that accepts PersistenceProfileDstAddrMap and PersistenceProfileDstAddrMapOutput values.
// You can construct a concrete instance of `PersistenceProfileDstAddrMapInput` via:
//
//	PersistenceProfileDstAddrMap{ "key": PersistenceProfileDstAddrArgs{...} }
type PersistenceProfileDstAddrMapInput interface {
	pulumi.Input

	ToPersistenceProfileDstAddrMapOutput() PersistenceProfileDstAddrMapOutput
	ToPersistenceProfileDstAddrMapOutputWithContext(context.Context) PersistenceProfileDstAddrMapOutput
}

type PersistenceProfileDstAddrMap map[string]PersistenceProfileDstAddrInput

func (PersistenceProfileDstAddrMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileDstAddr)(nil)).Elem()
}

func (i PersistenceProfileDstAddrMap) ToPersistenceProfileDstAddrMapOutput() PersistenceProfileDstAddrMapOutput {
	return i.ToPersistenceProfileDstAddrMapOutputWithContext(context.Background())
}

func (i PersistenceProfileDstAddrMap) ToPersistenceProfileDstAddrMapOutputWithContext(ctx context.Context) PersistenceProfileDstAddrMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileDstAddrMapOutput)
}

type PersistenceProfileDstAddrOutput struct{ *pulumi.OutputState }

func (PersistenceProfileDstAddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileDstAddr)(nil)).Elem()
}

func (o PersistenceProfileDstAddrOutput) ToPersistenceProfileDstAddrOutput() PersistenceProfileDstAddrOutput {
	return o
}

func (o PersistenceProfileDstAddrOutput) ToPersistenceProfileDstAddrOutputWithContext(ctx context.Context) PersistenceProfileDstAddrOutput {
	return o
}

func (o PersistenceProfileDstAddrOutput) AppService() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.AppService }).(pulumi.StringOutput)
}

// Inherit defaults from parent profile
func (o PersistenceProfileDstAddrOutput) DefaultsFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.DefaultsFrom }).(pulumi.StringOutput)
}

// Specify the hash algorithm
func (o PersistenceProfileDstAddrOutput) HashAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.HashAlgorithm }).(pulumi.StringOutput)
}

// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
func (o PersistenceProfileDstAddrOutput) Mask() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.Mask }).(pulumi.StringOutput)
}

// To enable _ disable match across pools with given persistence record
func (o PersistenceProfileDstAddrOutput) MatchAcrossPools() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.MatchAcrossPools }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileDstAddrOutput) MatchAcrossServices() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.MatchAcrossServices }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileDstAddrOutput) MatchAcrossVirtuals() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.MatchAcrossVirtuals }).(pulumi.StringOutput)
}

// To enable _ disable
func (o PersistenceProfileDstAddrOutput) Mirror() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.Mirror }).(pulumi.StringOutput)
}

// Name of the persistence profile
func (o PersistenceProfileDstAddrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (o PersistenceProfileDstAddrOutput) OverrideConnLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.StringOutput { return v.OverrideConnLimit }).(pulumi.StringOutput)
}

// Timeout for persistence of the session
func (o PersistenceProfileDstAddrOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *PersistenceProfileDstAddr) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

type PersistenceProfileDstAddrArrayOutput struct{ *pulumi.OutputState }

func (PersistenceProfileDstAddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileDstAddr)(nil)).Elem()
}

func (o PersistenceProfileDstAddrArrayOutput) ToPersistenceProfileDstAddrArrayOutput() PersistenceProfileDstAddrArrayOutput {
	return o
}

func (o PersistenceProfileDstAddrArrayOutput) ToPersistenceProfileDstAddrArrayOutputWithContext(ctx context.Context) PersistenceProfileDstAddrArrayOutput {
	return o
}

func (o PersistenceProfileDstAddrArrayOutput) Index(i pulumi.IntInput) PersistenceProfileDstAddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistenceProfileDstAddr {
		return vs[0].([]*PersistenceProfileDstAddr)[vs[1].(int)]
	}).(PersistenceProfileDstAddrOutput)
}

type PersistenceProfileDstAddrMapOutput struct{ *pulumi.OutputState }

func (PersistenceProfileDstAddrMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileDstAddr)(nil)).Elem()
}

func (o PersistenceProfileDstAddrMapOutput) ToPersistenceProfileDstAddrMapOutput() PersistenceProfileDstAddrMapOutput {
	return o
}

func (o PersistenceProfileDstAddrMapOutput) ToPersistenceProfileDstAddrMapOutputWithContext(ctx context.Context) PersistenceProfileDstAddrMapOutput {
	return o
}

func (o PersistenceProfileDstAddrMapOutput) MapIndex(k pulumi.StringInput) PersistenceProfileDstAddrOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistenceProfileDstAddr {
		return vs[0].(map[string]*PersistenceProfileDstAddr)[vs[1].(string)]
	}).(PersistenceProfileDstAddrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileDstAddrInput)(nil)).Elem(), &PersistenceProfileDstAddr{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileDstAddrArrayInput)(nil)).Elem(), PersistenceProfileDstAddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileDstAddrMapInput)(nil)).Elem(), PersistenceProfileDstAddrMap{})
	pulumi.RegisterOutputType(PersistenceProfileDstAddrOutput{})
	pulumi.RegisterOutputType(PersistenceProfileDstAddrArrayOutput{})
	pulumi.RegisterOutputType(PersistenceProfileDstAddrMapOutput{})
}
