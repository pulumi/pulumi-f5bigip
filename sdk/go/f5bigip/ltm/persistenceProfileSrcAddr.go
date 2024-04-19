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

// Configures a source address persistence profile
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
//			_, err := ltm.NewPersistenceProfileSrcAddr(ctx, "srcaddr", &ltm.PersistenceProfileSrcAddrArgs{
//				DefaultsFrom:        pulumi.String("/Common/source_addr"),
//				HashAlgorithm:       pulumi.String("carp"),
//				MapProxies:          pulumi.String("enabled"),
//				Mask:                pulumi.String("255.255.255.255"),
//				MatchAcrossPools:    pulumi.String("enabled"),
//				MatchAcrossServices: pulumi.String("enabled"),
//				MatchAcrossVirtuals: pulumi.String("enabled"),
//				Mirror:              pulumi.String("enabled"),
//				Name:                pulumi.String("/Common/terraform_srcaddr"),
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
// `defaultsFrom` - (Required) Parent cookie persistence profile
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
// `hashAlgorithm` (Optional) Specify the hash algorithm
//
// `mask` (Optional) Identify a range of source IP addresses to manage together as a single source address affinity persistent connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
//
// `mapProxies` (Optional) (enabled or disabled) Directs all to the same single pool member
//
// ## Importing
//
// An source-addr persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
// An example is below:
// ```sh
// $ terraform import bigip_ltm_persistence_profile_srcaddr.srcaddr "/Common/terraform_srcaddr"
// ```
type PersistenceProfileSrcAddr struct {
	pulumi.CustomResourceState

	AppService pulumi.StringOutput `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringOutput `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringOutput `pulumi:"mapProxies"`
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

// NewPersistenceProfileSrcAddr registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileSrcAddr(ctx *pulumi.Context,
	name string, args *PersistenceProfileSrcAddrArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileSrcAddr, error) {
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
	var resource PersistenceProfileSrcAddr
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileSrcAddr gets an existing PersistenceProfileSrcAddr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileSrcAddr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileSrcAddrState, opts ...pulumi.ResourceOption) (*PersistenceProfileSrcAddr, error) {
	var resource PersistenceProfileSrcAddr
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileSrcAddr resources.
type persistenceProfileSrcAddrState struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies *string `pulumi:"mapProxies"`
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

type PersistenceProfileSrcAddrState struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringPtrInput
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

func (PersistenceProfileSrcAddrState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSrcAddrState)(nil)).Elem()
}

type persistenceProfileSrcAddrArgs struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies *string `pulumi:"mapProxies"`
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

// The set of arguments for constructing a PersistenceProfileSrcAddr resource.
type PersistenceProfileSrcAddrArgs struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringPtrInput
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

func (PersistenceProfileSrcAddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSrcAddrArgs)(nil)).Elem()
}

type PersistenceProfileSrcAddrInput interface {
	pulumi.Input

	ToPersistenceProfileSrcAddrOutput() PersistenceProfileSrcAddrOutput
	ToPersistenceProfileSrcAddrOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrOutput
}

func (*PersistenceProfileSrcAddr) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileSrcAddr)(nil)).Elem()
}

func (i *PersistenceProfileSrcAddr) ToPersistenceProfileSrcAddrOutput() PersistenceProfileSrcAddrOutput {
	return i.ToPersistenceProfileSrcAddrOutputWithContext(context.Background())
}

func (i *PersistenceProfileSrcAddr) ToPersistenceProfileSrcAddrOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSrcAddrOutput)
}

// PersistenceProfileSrcAddrArrayInput is an input type that accepts PersistenceProfileSrcAddrArray and PersistenceProfileSrcAddrArrayOutput values.
// You can construct a concrete instance of `PersistenceProfileSrcAddrArrayInput` via:
//
//	PersistenceProfileSrcAddrArray{ PersistenceProfileSrcAddrArgs{...} }
type PersistenceProfileSrcAddrArrayInput interface {
	pulumi.Input

	ToPersistenceProfileSrcAddrArrayOutput() PersistenceProfileSrcAddrArrayOutput
	ToPersistenceProfileSrcAddrArrayOutputWithContext(context.Context) PersistenceProfileSrcAddrArrayOutput
}

type PersistenceProfileSrcAddrArray []PersistenceProfileSrcAddrInput

func (PersistenceProfileSrcAddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileSrcAddr)(nil)).Elem()
}

func (i PersistenceProfileSrcAddrArray) ToPersistenceProfileSrcAddrArrayOutput() PersistenceProfileSrcAddrArrayOutput {
	return i.ToPersistenceProfileSrcAddrArrayOutputWithContext(context.Background())
}

func (i PersistenceProfileSrcAddrArray) ToPersistenceProfileSrcAddrArrayOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSrcAddrArrayOutput)
}

// PersistenceProfileSrcAddrMapInput is an input type that accepts PersistenceProfileSrcAddrMap and PersistenceProfileSrcAddrMapOutput values.
// You can construct a concrete instance of `PersistenceProfileSrcAddrMapInput` via:
//
//	PersistenceProfileSrcAddrMap{ "key": PersistenceProfileSrcAddrArgs{...} }
type PersistenceProfileSrcAddrMapInput interface {
	pulumi.Input

	ToPersistenceProfileSrcAddrMapOutput() PersistenceProfileSrcAddrMapOutput
	ToPersistenceProfileSrcAddrMapOutputWithContext(context.Context) PersistenceProfileSrcAddrMapOutput
}

type PersistenceProfileSrcAddrMap map[string]PersistenceProfileSrcAddrInput

func (PersistenceProfileSrcAddrMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileSrcAddr)(nil)).Elem()
}

func (i PersistenceProfileSrcAddrMap) ToPersistenceProfileSrcAddrMapOutput() PersistenceProfileSrcAddrMapOutput {
	return i.ToPersistenceProfileSrcAddrMapOutputWithContext(context.Background())
}

func (i PersistenceProfileSrcAddrMap) ToPersistenceProfileSrcAddrMapOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSrcAddrMapOutput)
}

type PersistenceProfileSrcAddrOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSrcAddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileSrcAddr)(nil)).Elem()
}

func (o PersistenceProfileSrcAddrOutput) ToPersistenceProfileSrcAddrOutput() PersistenceProfileSrcAddrOutput {
	return o
}

func (o PersistenceProfileSrcAddrOutput) ToPersistenceProfileSrcAddrOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrOutput {
	return o
}

func (o PersistenceProfileSrcAddrOutput) AppService() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.AppService }).(pulumi.StringOutput)
}

// Inherit defaults from parent profile
func (o PersistenceProfileSrcAddrOutput) DefaultsFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.DefaultsFrom }).(pulumi.StringOutput)
}

// Specify the hash algorithm
func (o PersistenceProfileSrcAddrOutput) HashAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.HashAlgorithm }).(pulumi.StringOutput)
}

// To enable _ disable directs all to the same single pool member
func (o PersistenceProfileSrcAddrOutput) MapProxies() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.MapProxies }).(pulumi.StringOutput)
}

// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
func (o PersistenceProfileSrcAddrOutput) Mask() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.Mask }).(pulumi.StringOutput)
}

// To enable _ disable match across pools with given persistence record
func (o PersistenceProfileSrcAddrOutput) MatchAcrossPools() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.MatchAcrossPools }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileSrcAddrOutput) MatchAcrossServices() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.MatchAcrossServices }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileSrcAddrOutput) MatchAcrossVirtuals() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.MatchAcrossVirtuals }).(pulumi.StringOutput)
}

// To enable _ disable
func (o PersistenceProfileSrcAddrOutput) Mirror() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.Mirror }).(pulumi.StringOutput)
}

// Name of the persistence profile
func (o PersistenceProfileSrcAddrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (o PersistenceProfileSrcAddrOutput) OverrideConnLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.StringOutput { return v.OverrideConnLimit }).(pulumi.StringOutput)
}

// Timeout for persistence of the session
func (o PersistenceProfileSrcAddrOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *PersistenceProfileSrcAddr) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

type PersistenceProfileSrcAddrArrayOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSrcAddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileSrcAddr)(nil)).Elem()
}

func (o PersistenceProfileSrcAddrArrayOutput) ToPersistenceProfileSrcAddrArrayOutput() PersistenceProfileSrcAddrArrayOutput {
	return o
}

func (o PersistenceProfileSrcAddrArrayOutput) ToPersistenceProfileSrcAddrArrayOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrArrayOutput {
	return o
}

func (o PersistenceProfileSrcAddrArrayOutput) Index(i pulumi.IntInput) PersistenceProfileSrcAddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistenceProfileSrcAddr {
		return vs[0].([]*PersistenceProfileSrcAddr)[vs[1].(int)]
	}).(PersistenceProfileSrcAddrOutput)
}

type PersistenceProfileSrcAddrMapOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSrcAddrMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileSrcAddr)(nil)).Elem()
}

func (o PersistenceProfileSrcAddrMapOutput) ToPersistenceProfileSrcAddrMapOutput() PersistenceProfileSrcAddrMapOutput {
	return o
}

func (o PersistenceProfileSrcAddrMapOutput) ToPersistenceProfileSrcAddrMapOutputWithContext(ctx context.Context) PersistenceProfileSrcAddrMapOutput {
	return o
}

func (o PersistenceProfileSrcAddrMapOutput) MapIndex(k pulumi.StringInput) PersistenceProfileSrcAddrOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistenceProfileSrcAddr {
		return vs[0].(map[string]*PersistenceProfileSrcAddr)[vs[1].(string)]
	}).(PersistenceProfileSrcAddrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSrcAddrInput)(nil)).Elem(), &PersistenceProfileSrcAddr{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSrcAddrArrayInput)(nil)).Elem(), PersistenceProfileSrcAddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSrcAddrMapInput)(nil)).Elem(), PersistenceProfileSrcAddrMap{})
	pulumi.RegisterOutputType(PersistenceProfileSrcAddrOutput{})
	pulumi.RegisterOutputType(PersistenceProfileSrcAddrArrayOutput{})
	pulumi.RegisterOutputType(PersistenceProfileSrcAddrMapOutput{})
}
