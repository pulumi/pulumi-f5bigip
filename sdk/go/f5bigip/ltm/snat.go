// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.Snat` Manages a snat configuration
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewSnat(ctx, "test_snat", &ltm.SnatArgs{
// 			Autolasthop: pulumi.String("default"),
// 			FullPath:    pulumi.String("/Common/test-snat"),
// 			Mirror:      pulumi.String("disabled"),
// 			Name:        pulumi.String("TEST_SNAT_NAME"),
// 			Origins: ltm.SnatOriginArray{
// 				&ltm.SnatOriginArgs{
// 					Name: pulumi.String("2.2.2.2"),
// 				},
// 				&ltm.SnatOriginArgs{
// 					Name: pulumi.String("3.3.3.3"),
// 				},
// 			},
// 			Partition:     pulumi.String("Common"),
// 			Translation:   pulumi.String("/Common/136.1.1.1"),
// 			Vlansdisabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Snat struct {
	pulumi.CustomResourceState

	// -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringPtrOutput `pulumi:"autolasthop"`
	// Fullpath
	FullPath pulumi.StringPtrOutput `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringPtrOutput `pulumi:"mirror"`
	// Name of the snat
	Name pulumi.StringOutput `pulumi:"name"`
	// IP or hostname of the snat
	Origins SnatOriginArrayOutput `pulumi:"origins"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
	Snatpool pulumi.StringPtrOutput `pulumi:"snatpool"`
	// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
	Sourceport pulumi.StringPtrOutput `pulumi:"sourceport"`
	// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
	Translation pulumi.StringPtrOutput `pulumi:"translation"`
	// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
	Vlans pulumi.StringArrayOutput `pulumi:"vlans"`
	// Disables the SNAT on all VLANs.
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
	// -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop *string `pulumi:"autolasthop"`
	// Fullpath
	FullPath *string `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror *string `pulumi:"mirror"`
	// Name of the snat
	Name *string `pulumi:"name"`
	// IP or hostname of the snat
	Origins []SnatOrigin `pulumi:"origins"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
	Sourceport *string `pulumi:"sourceport"`
	// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
	Translation *string `pulumi:"translation"`
	// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
	Vlans []string `pulumi:"vlans"`
	// Disables the SNAT on all VLANs.
	Vlansdisabled *bool `pulumi:"vlansdisabled"`
}

type SnatState struct {
	// -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringPtrInput
	// Fullpath
	FullPath pulumi.StringPtrInput
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringPtrInput
	// Name of the snat
	Name pulumi.StringPtrInput
	// IP or hostname of the snat
	Origins SnatOriginArrayInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
	Snatpool pulumi.StringPtrInput
	// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
	Sourceport pulumi.StringPtrInput
	// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
	Translation pulumi.StringPtrInput
	// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
	Vlans pulumi.StringArrayInput
	// Disables the SNAT on all VLANs.
	Vlansdisabled pulumi.BoolPtrInput
}

func (SnatState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatState)(nil)).Elem()
}

type snatArgs struct {
	// -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop *string `pulumi:"autolasthop"`
	// Fullpath
	FullPath *string `pulumi:"fullPath"`
	// Enables or disables mirroring of SNAT connections.
	Mirror *string `pulumi:"mirror"`
	// Name of the snat
	Name string `pulumi:"name"`
	// IP or hostname of the snat
	Origins []SnatOrigin `pulumi:"origins"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
	// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
	Sourceport *string `pulumi:"sourceport"`
	// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
	Translation *string `pulumi:"translation"`
	// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
	Vlans []string `pulumi:"vlans"`
	// Disables the SNAT on all VLANs.
	Vlansdisabled *bool `pulumi:"vlansdisabled"`
}

// The set of arguments for constructing a Snat resource.
type SnatArgs struct {
	// -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
	Autolasthop pulumi.StringPtrInput
	// Fullpath
	FullPath pulumi.StringPtrInput
	// Enables or disables mirroring of SNAT connections.
	Mirror pulumi.StringPtrInput
	// Name of the snat
	Name pulumi.StringInput
	// IP or hostname of the snat
	Origins SnatOriginArrayInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
	// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
	Snatpool pulumi.StringPtrInput
	// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
	Sourceport pulumi.StringPtrInput
	// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
	Translation pulumi.StringPtrInput
	// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
	Vlans pulumi.StringArrayInput
	// Disables the SNAT on all VLANs.
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

func (Snat) ElementType() reflect.Type {
	return reflect.TypeOf((*Snat)(nil)).Elem()
}

func (i Snat) ToSnatOutput() SnatOutput {
	return i.ToSnatOutputWithContext(context.Background())
}

func (i Snat) ToSnatOutputWithContext(ctx context.Context) SnatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatOutput)
}

type SnatOutput struct {
	*pulumi.OutputState
}

func (SnatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnatOutput)(nil)).Elem()
}

func (o SnatOutput) ToSnatOutput() SnatOutput {
	return o
}

func (o SnatOutput) ToSnatOutputWithContext(ctx context.Context) SnatOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnatOutput{})
}
