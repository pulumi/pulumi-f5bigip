// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `net.Vlan` Manages a vlan configuration
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/net"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := net.NewVlan(ctx, "vlan1", &net.VlanArgs{
// 			Interfaces: net.VlanInterfaceArray{
// 				&net.VlanInterfaceArgs{
// 					Tagged:   pulumi.Bool(false),
// 					Vlanport: pulumi.String("1.2"),
// 				},
// 			},
// 			Name: pulumi.String("/Common/Internal"),
// 			Tag:  pulumi.Int(101),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Vlan struct {
	pulumi.CustomResourceState

	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces VlanInterfaceArrayOutput `pulumi:"interfaces"`
	// Name of the vlan
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag pulumi.IntPtrOutput `pulumi:"tag"`
}

// NewVlan registers a new resource with the given unique name, arguments, and options.
func NewVlan(ctx *pulumi.Context,
	name string, args *VlanArgs, opts ...pulumi.ResourceOption) (*Vlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Vlan
	err := ctx.RegisterResource("f5bigip:net/vlan:Vlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlan gets an existing Vlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanState, opts ...pulumi.ResourceOption) (*Vlan, error) {
	var resource Vlan
	err := ctx.ReadResource("f5bigip:net/vlan:Vlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vlan resources.
type vlanState struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces []VlanInterface `pulumi:"interfaces"`
	// Name of the vlan
	Name *string `pulumi:"name"`
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag *int `pulumi:"tag"`
}

type VlanState struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces VlanInterfaceArrayInput
	// Name of the vlan
	Name pulumi.StringPtrInput
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag pulumi.IntPtrInput
}

func (VlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanState)(nil)).Elem()
}

type vlanArgs struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces []VlanInterface `pulumi:"interfaces"`
	// Name of the vlan
	Name string `pulumi:"name"`
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag *int `pulumi:"tag"`
}

// The set of arguments for constructing a Vlan resource.
type VlanArgs struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces VlanInterfaceArrayInput
	// Name of the vlan
	Name pulumi.StringInput
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag pulumi.IntPtrInput
}

func (VlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanArgs)(nil)).Elem()
}

type VlanInput interface {
	pulumi.Input

	ToVlanOutput() VlanOutput
	ToVlanOutputWithContext(ctx context.Context) VlanOutput
}

func (Vlan) ElementType() reflect.Type {
	return reflect.TypeOf((*Vlan)(nil)).Elem()
}

func (i Vlan) ToVlanOutput() VlanOutput {
	return i.ToVlanOutputWithContext(context.Background())
}

func (i Vlan) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanOutput)
}

type VlanOutput struct {
	*pulumi.OutputState
}

func (VlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VlanOutput)(nil)).Elem()
}

func (o VlanOutput) ToVlanOutput() VlanOutput {
	return o
}

func (o VlanOutput) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VlanOutput{})
}
