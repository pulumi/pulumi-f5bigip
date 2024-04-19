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

// `ltm.Snat` Manages a SNAT configuration
//
// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource.For example `/Common/test-snat`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := ltm.NewSnat(ctx, "test-snat", &ltm.SnatArgs{
//				Name:        pulumi.String("/Common/test-snat"),
//				Translation: pulumi.String("/Common/136.1.1.2"),
//				Sourceport:  pulumi.String("preserve"),
//				Origins: ltm.SnatOriginArray{
//					&ltm.SnatOriginArgs{
//						Name: pulumi.String("0.0.0.0/0"),
//					},
//				},
//				Vlans: pulumi.StringArray{
//					pulumi.String("/Common/internal"),
//				},
//				Vlansdisabled: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type Snat struct {
	pulumi.CustomResourceState

	// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringOutput `pulumi:"autolasthop"`
	// Fullpath
	FullPath pulumi.StringPtrOutput `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringOutput `pulumi:"mirror"`
	// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
	Origins SnatOriginArrayOutput `pulumi:"origins"`
	// Partition or path to which the SNAT belongs
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
	Snatpool pulumi.StringPtrOutput `pulumi:"snatpool"`
	// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
	Sourceport pulumi.StringPtrOutput `pulumi:"sourceport"`
	// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
	Translation pulumi.StringPtrOutput `pulumi:"translation"`
	// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
	Vlans pulumi.StringArrayOutput `pulumi:"vlans"`
	// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
	Vlansdisabled pulumi.BoolPtrOutput `pulumi:"vlansdisabled"`
}

// NewSnat registers a new resource with the given unique name, arguments, and options.
func NewSnat(ctx *pulumi.Context,
	name string, args *SnatArgs, opts ...pulumi.ResourceOption) (*Snat, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Origins == nil {
		return nil, errors.New("invalid value for required argument 'Origins'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snat
	err := ctx.RegisterResource("f5bigip:ltm/snat:Snat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnat gets an existing Snat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatState, opts ...pulumi.ResourceOption) (*Snat, error) {
	var resource Snat
	err := ctx.ReadResource("f5bigip:ltm/snat:Snat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snat resources.
type snatState struct {
	// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop *string `pulumi:"autolasthop"`
	// Fullpath
	FullPath *string `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror *string `pulumi:"mirror"`
	// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
	Name *string `pulumi:"name"`
	// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
	Origins []SnatOrigin `pulumi:"origins"`
	// Partition or path to which the SNAT belongs
	Partition *string `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
	Sourceport *string `pulumi:"sourceport"`
	// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
	Translation *string `pulumi:"translation"`
	// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
	Vlans []string `pulumi:"vlans"`
	// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
	Vlansdisabled *bool `pulumi:"vlansdisabled"`
}

type SnatState struct {
	// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringPtrInput
	// Fullpath
	FullPath pulumi.StringPtrInput
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringPtrInput
	// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
	Name pulumi.StringPtrInput
	// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
	Origins SnatOriginArrayInput
	// Partition or path to which the SNAT belongs
	Partition pulumi.StringPtrInput
	// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
	Snatpool pulumi.StringPtrInput
	// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
	Sourceport pulumi.StringPtrInput
	// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
	Translation pulumi.StringPtrInput
	// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
	Vlans pulumi.StringArrayInput
	// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
	Vlansdisabled pulumi.BoolPtrInput
}

func (SnatState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatState)(nil)).Elem()
}

type snatArgs struct {
	// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop *string `pulumi:"autolasthop"`
	// Fullpath
	FullPath *string `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror *string `pulumi:"mirror"`
	// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
	Name string `pulumi:"name"`
	// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
	Origins []SnatOrigin `pulumi:"origins"`
	// Partition or path to which the SNAT belongs
	Partition *string `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
	Sourceport *string `pulumi:"sourceport"`
	// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
	Translation *string `pulumi:"translation"`
	// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
	Vlans []string `pulumi:"vlans"`
	// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
	Vlansdisabled *bool `pulumi:"vlansdisabled"`
}

// The set of arguments for constructing a Snat resource.
type SnatArgs struct {
	// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringPtrInput
	// Fullpath
	FullPath pulumi.StringPtrInput
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringPtrInput
	// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
	Name pulumi.StringInput
	// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
	Origins SnatOriginArrayInput
	// Partition or path to which the SNAT belongs
	Partition pulumi.StringPtrInput
	// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
	Snatpool pulumi.StringPtrInput
	// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
	Sourceport pulumi.StringPtrInput
	// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
	Translation pulumi.StringPtrInput
	// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
	Vlans pulumi.StringArrayInput
	// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
	Vlansdisabled pulumi.BoolPtrInput
}

func (SnatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatArgs)(nil)).Elem()
}

type SnatInput interface {
	pulumi.Input

	ToSnatOutput() SnatOutput
	ToSnatOutputWithContext(ctx context.Context) SnatOutput
}

func (*Snat) ElementType() reflect.Type {
	return reflect.TypeOf((**Snat)(nil)).Elem()
}

func (i *Snat) ToSnatOutput() SnatOutput {
	return i.ToSnatOutputWithContext(context.Background())
}

func (i *Snat) ToSnatOutputWithContext(ctx context.Context) SnatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatOutput)
}

// SnatArrayInput is an input type that accepts SnatArray and SnatArrayOutput values.
// You can construct a concrete instance of `SnatArrayInput` via:
//
//	SnatArray{ SnatArgs{...} }
type SnatArrayInput interface {
	pulumi.Input

	ToSnatArrayOutput() SnatArrayOutput
	ToSnatArrayOutputWithContext(context.Context) SnatArrayOutput
}

type SnatArray []SnatInput

func (SnatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snat)(nil)).Elem()
}

func (i SnatArray) ToSnatArrayOutput() SnatArrayOutput {
	return i.ToSnatArrayOutputWithContext(context.Background())
}

func (i SnatArray) ToSnatArrayOutputWithContext(ctx context.Context) SnatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatArrayOutput)
}

// SnatMapInput is an input type that accepts SnatMap and SnatMapOutput values.
// You can construct a concrete instance of `SnatMapInput` via:
//
//	SnatMap{ "key": SnatArgs{...} }
type SnatMapInput interface {
	pulumi.Input

	ToSnatMapOutput() SnatMapOutput
	ToSnatMapOutputWithContext(context.Context) SnatMapOutput
}

type SnatMap map[string]SnatInput

func (SnatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snat)(nil)).Elem()
}

func (i SnatMap) ToSnatMapOutput() SnatMapOutput {
	return i.ToSnatMapOutputWithContext(context.Background())
}

func (i SnatMap) ToSnatMapOutputWithContext(ctx context.Context) SnatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatMapOutput)
}

type SnatOutput struct{ *pulumi.OutputState }

func (SnatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snat)(nil)).Elem()
}

func (o SnatOutput) ToSnatOutput() SnatOutput {
	return o
}

func (o SnatOutput) ToSnatOutputWithContext(ctx context.Context) SnatOutput {
	return o
}

// Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
func (o SnatOutput) Autolasthop() pulumi.StringOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringOutput { return v.Autolasthop }).(pulumi.StringOutput)
}

// Fullpath
func (o SnatOutput) FullPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringPtrOutput { return v.FullPath }).(pulumi.StringPtrOutput)
}

// Enables or disables mirroring of SNAT connections.
func (o SnatOutput) Mirror() pulumi.StringOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringOutput { return v.Mirror }).(pulumi.StringOutput)
}

// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
func (o SnatOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
func (o SnatOutput) Origins() SnatOriginArrayOutput {
	return o.ApplyT(func(v *Snat) SnatOriginArrayOutput { return v.Origins }).(SnatOriginArrayOutput)
}

// Partition or path to which the SNAT belongs
func (o SnatOutput) Partition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringPtrOutput { return v.Partition }).(pulumi.StringPtrOutput)
}

// Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
func (o SnatOutput) Snatpool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringPtrOutput { return v.Snatpool }).(pulumi.StringPtrOutput)
}

// Specifies how the SNAT object handles the client's source port. The default is `preserve`.
func (o SnatOutput) Sourceport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringPtrOutput { return v.Sourceport }).(pulumi.StringPtrOutput)
}

// Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
func (o SnatOutput) Translation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringPtrOutput { return v.Translation }).(pulumi.StringPtrOutput)
}

// Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
func (o SnatOutput) Vlans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snat) pulumi.StringArrayOutput { return v.Vlans }).(pulumi.StringArrayOutput)
}

// Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
func (o SnatOutput) Vlansdisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snat) pulumi.BoolPtrOutput { return v.Vlansdisabled }).(pulumi.BoolPtrOutput)
}

type SnatArrayOutput struct{ *pulumi.OutputState }

func (SnatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snat)(nil)).Elem()
}

func (o SnatArrayOutput) ToSnatArrayOutput() SnatArrayOutput {
	return o
}

func (o SnatArrayOutput) ToSnatArrayOutputWithContext(ctx context.Context) SnatArrayOutput {
	return o
}

func (o SnatArrayOutput) Index(i pulumi.IntInput) SnatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snat {
		return vs[0].([]*Snat)[vs[1].(int)]
	}).(SnatOutput)
}

type SnatMapOutput struct{ *pulumi.OutputState }

func (SnatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snat)(nil)).Elem()
}

func (o SnatMapOutput) ToSnatMapOutput() SnatMapOutput {
	return o
}

func (o SnatMapOutput) ToSnatMapOutputWithContext(ctx context.Context) SnatMapOutput {
	return o
}

func (o SnatMapOutput) MapIndex(k pulumi.StringInput) SnatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snat {
		return vs[0].(map[string]*Snat)[vs[1].(string)]
	}).(SnatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnatInput)(nil)).Elem(), &Snat{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatArrayInput)(nil)).Elem(), SnatArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatMapInput)(nil)).Elem(), SnatMap{})
	pulumi.RegisterOutputType(SnatOutput{})
	pulumi.RegisterOutputType(SnatArrayOutput{})
	pulumi.RegisterOutputType(SnatMapOutput{})
}
