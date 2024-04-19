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

// Configures an SSL persistence profile
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
//			_, err := ltm.NewPersistenceProfileSsl(ctx, "ppssl", &ltm.PersistenceProfileSslArgs{
//				DefaultsFrom:        pulumi.String("/Common/ssl"),
//				MatchAcrossPools:    pulumi.String("enabled"),
//				MatchAcrossServices: pulumi.String("enabled"),
//				MatchAcrossVirtuals: pulumi.String("enabled"),
//				Mirror:              pulumi.String("enabled"),
//				Name:                pulumi.String("/Common/terraform_ssl"),
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
// ## Importing
//
// An ssl persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
// An example is below:
// ```sh
// $ terraform import bigip_ltm_persistence_profile_ssl.ppssl "/Common/terraform_ssl"
// ```
type PersistenceProfileSsl struct {
	pulumi.CustomResourceState

	AppService pulumi.StringPtrOutput `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
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
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewPersistenceProfileSsl registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileSsl(ctx *pulumi.Context,
	name string, args *PersistenceProfileSslArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileSsl, error) {
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
	var resource PersistenceProfileSsl
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileSsl gets an existing PersistenceProfileSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileSslState, opts ...pulumi.ResourceOption) (*PersistenceProfileSsl, error) {
	var resource PersistenceProfileSsl
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileSsl resources.
type persistenceProfileSslState struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
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

type PersistenceProfileSslState struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
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

func (PersistenceProfileSslState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSslState)(nil)).Elem()
}

type persistenceProfileSslArgs struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
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

// The set of arguments for constructing a PersistenceProfileSsl resource.
type PersistenceProfileSslArgs struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
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

func (PersistenceProfileSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSslArgs)(nil)).Elem()
}

type PersistenceProfileSslInput interface {
	pulumi.Input

	ToPersistenceProfileSslOutput() PersistenceProfileSslOutput
	ToPersistenceProfileSslOutputWithContext(ctx context.Context) PersistenceProfileSslOutput
}

func (*PersistenceProfileSsl) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileSsl)(nil)).Elem()
}

func (i *PersistenceProfileSsl) ToPersistenceProfileSslOutput() PersistenceProfileSslOutput {
	return i.ToPersistenceProfileSslOutputWithContext(context.Background())
}

func (i *PersistenceProfileSsl) ToPersistenceProfileSslOutputWithContext(ctx context.Context) PersistenceProfileSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSslOutput)
}

// PersistenceProfileSslArrayInput is an input type that accepts PersistenceProfileSslArray and PersistenceProfileSslArrayOutput values.
// You can construct a concrete instance of `PersistenceProfileSslArrayInput` via:
//
//	PersistenceProfileSslArray{ PersistenceProfileSslArgs{...} }
type PersistenceProfileSslArrayInput interface {
	pulumi.Input

	ToPersistenceProfileSslArrayOutput() PersistenceProfileSslArrayOutput
	ToPersistenceProfileSslArrayOutputWithContext(context.Context) PersistenceProfileSslArrayOutput
}

type PersistenceProfileSslArray []PersistenceProfileSslInput

func (PersistenceProfileSslArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileSsl)(nil)).Elem()
}

func (i PersistenceProfileSslArray) ToPersistenceProfileSslArrayOutput() PersistenceProfileSslArrayOutput {
	return i.ToPersistenceProfileSslArrayOutputWithContext(context.Background())
}

func (i PersistenceProfileSslArray) ToPersistenceProfileSslArrayOutputWithContext(ctx context.Context) PersistenceProfileSslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSslArrayOutput)
}

// PersistenceProfileSslMapInput is an input type that accepts PersistenceProfileSslMap and PersistenceProfileSslMapOutput values.
// You can construct a concrete instance of `PersistenceProfileSslMapInput` via:
//
//	PersistenceProfileSslMap{ "key": PersistenceProfileSslArgs{...} }
type PersistenceProfileSslMapInput interface {
	pulumi.Input

	ToPersistenceProfileSslMapOutput() PersistenceProfileSslMapOutput
	ToPersistenceProfileSslMapOutputWithContext(context.Context) PersistenceProfileSslMapOutput
}

type PersistenceProfileSslMap map[string]PersistenceProfileSslInput

func (PersistenceProfileSslMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileSsl)(nil)).Elem()
}

func (i PersistenceProfileSslMap) ToPersistenceProfileSslMapOutput() PersistenceProfileSslMapOutput {
	return i.ToPersistenceProfileSslMapOutputWithContext(context.Background())
}

func (i PersistenceProfileSslMap) ToPersistenceProfileSslMapOutputWithContext(ctx context.Context) PersistenceProfileSslMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileSslMapOutput)
}

type PersistenceProfileSslOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileSsl)(nil)).Elem()
}

func (o PersistenceProfileSslOutput) ToPersistenceProfileSslOutput() PersistenceProfileSslOutput {
	return o
}

func (o PersistenceProfileSslOutput) ToPersistenceProfileSslOutputWithContext(ctx context.Context) PersistenceProfileSslOutput {
	return o
}

func (o PersistenceProfileSslOutput) AppService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringPtrOutput { return v.AppService }).(pulumi.StringPtrOutput)
}

// Inherit defaults from parent profile
func (o PersistenceProfileSslOutput) DefaultsFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.DefaultsFrom }).(pulumi.StringOutput)
}

// To enable _ disable match across pools with given persistence record
func (o PersistenceProfileSslOutput) MatchAcrossPools() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.MatchAcrossPools }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileSslOutput) MatchAcrossServices() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.MatchAcrossServices }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileSslOutput) MatchAcrossVirtuals() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.MatchAcrossVirtuals }).(pulumi.StringOutput)
}

// To enable _ disable
func (o PersistenceProfileSslOutput) Mirror() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.Mirror }).(pulumi.StringOutput)
}

// Name of the persistence profile
func (o PersistenceProfileSslOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (o PersistenceProfileSslOutput) OverrideConnLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.StringOutput { return v.OverrideConnLimit }).(pulumi.StringOutput)
}

// Timeout for persistence of the session
func (o PersistenceProfileSslOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistenceProfileSsl) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type PersistenceProfileSslArrayOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileSsl)(nil)).Elem()
}

func (o PersistenceProfileSslArrayOutput) ToPersistenceProfileSslArrayOutput() PersistenceProfileSslArrayOutput {
	return o
}

func (o PersistenceProfileSslArrayOutput) ToPersistenceProfileSslArrayOutputWithContext(ctx context.Context) PersistenceProfileSslArrayOutput {
	return o
}

func (o PersistenceProfileSslArrayOutput) Index(i pulumi.IntInput) PersistenceProfileSslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistenceProfileSsl {
		return vs[0].([]*PersistenceProfileSsl)[vs[1].(int)]
	}).(PersistenceProfileSslOutput)
}

type PersistenceProfileSslMapOutput struct{ *pulumi.OutputState }

func (PersistenceProfileSslMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileSsl)(nil)).Elem()
}

func (o PersistenceProfileSslMapOutput) ToPersistenceProfileSslMapOutput() PersistenceProfileSslMapOutput {
	return o
}

func (o PersistenceProfileSslMapOutput) ToPersistenceProfileSslMapOutputWithContext(ctx context.Context) PersistenceProfileSslMapOutput {
	return o
}

func (o PersistenceProfileSslMapOutput) MapIndex(k pulumi.StringInput) PersistenceProfileSslOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistenceProfileSsl {
		return vs[0].(map[string]*PersistenceProfileSsl)[vs[1].(string)]
	}).(PersistenceProfileSslOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSslInput)(nil)).Elem(), &PersistenceProfileSsl{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSslArrayInput)(nil)).Elem(), PersistenceProfileSslArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileSslMapInput)(nil)).Elem(), PersistenceProfileSslMap{})
	pulumi.RegisterOutputType(PersistenceProfileSslOutput{})
	pulumi.RegisterOutputType(PersistenceProfileSslArrayOutput{})
	pulumi.RegisterOutputType(PersistenceProfileSslMapOutput{})
}
