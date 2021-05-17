// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `sys.SnmpTraps` provides details bout how to enable snmpTraps resource on BIG-IP
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/sys"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sys.NewSnmpTraps(ctx, "snmpTraps", &sys.SnmpTrapsArgs{
// 			Community:   pulumi.String("f5community"),
// 			Description: pulumi.String("Setup snmp traps"),
// 			Host:        pulumi.String("195.10.10.1"),
// 			Name:        pulumi.String("snmptraps"),
// 			Port:        pulumi.Int(111),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SnmpTraps struct {
	pulumi.CustomResourceState

	// Encrypted password
	AuthPasswordencrypted pulumi.StringPtrOutput `pulumi:"authPasswordencrypted"`
	// Specifies the protocol used to authenticate the user.
	AuthProtocol pulumi.StringPtrOutput `pulumi:"authProtocol"`
	// Specifies the community string used for this trap.
	Community pulumi.StringPtrOutput `pulumi:"community"`
	// The port that the trap will be sent to.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the authoritative security engine for SNMPv3.
	EngineId pulumi.StringPtrOutput `pulumi:"engineId"`
	// The host the trap will be sent to.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// Name of the snmp trap.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// User defined description.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
	PrivacyPassword pulumi.StringPtrOutput `pulumi:"privacyPassword"`
	// Specifies the encrypted password used to encrypt traffic.
	PrivacyPasswordEncrypted pulumi.StringPtrOutput `pulumi:"privacyPasswordEncrypted"`
	// Specifies the protocol used to encrypt traffic.
	PrivacyProtocol pulumi.StringPtrOutput `pulumi:"privacyProtocol"`
	// Specifies whether or not traffic is encrypted and whether or not authentication is required.
	SecurityLevel pulumi.StringOutput `pulumi:"securityLevel"`
	// Security name used in conjunction with SNMPv3.
	SecurityName pulumi.StringPtrOutput `pulumi:"securityName"`
	// SNMP version used for sending the trap.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSnmpTraps registers a new resource with the given unique name, arguments, and options.
func NewSnmpTraps(ctx *pulumi.Context,
	name string, args *SnmpTrapsArgs, opts ...pulumi.ResourceOption) (*SnmpTraps, error) {
	if args == nil {
		args = &SnmpTrapsArgs{}
	}

	var resource SnmpTraps
	err := ctx.RegisterResource("f5bigip:sys/snmpTraps:SnmpTraps", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnmpTraps gets an existing SnmpTraps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnmpTraps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnmpTrapsState, opts ...pulumi.ResourceOption) (*SnmpTraps, error) {
	var resource SnmpTraps
	err := ctx.ReadResource("f5bigip:sys/snmpTraps:SnmpTraps", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnmpTraps resources.
type snmpTrapsState struct {
	// Encrypted password
	AuthPasswordencrypted *string `pulumi:"authPasswordencrypted"`
	// Specifies the protocol used to authenticate the user.
	AuthProtocol *string `pulumi:"authProtocol"`
	// Specifies the community string used for this trap.
	Community *string `pulumi:"community"`
	// The port that the trap will be sent to.
	Description *string `pulumi:"description"`
	// Specifies the authoritative security engine for SNMPv3.
	EngineId *string `pulumi:"engineId"`
	// The host the trap will be sent to.
	Host *string `pulumi:"host"`
	// Name of the snmp trap.
	Name *string `pulumi:"name"`
	// User defined description.
	Port *int `pulumi:"port"`
	// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
	PrivacyPassword *string `pulumi:"privacyPassword"`
	// Specifies the encrypted password used to encrypt traffic.
	PrivacyPasswordEncrypted *string `pulumi:"privacyPasswordEncrypted"`
	// Specifies the protocol used to encrypt traffic.
	PrivacyProtocol *string `pulumi:"privacyProtocol"`
	// Specifies whether or not traffic is encrypted and whether or not authentication is required.
	SecurityLevel *string `pulumi:"securityLevel"`
	// Security name used in conjunction with SNMPv3.
	SecurityName *string `pulumi:"securityName"`
	// SNMP version used for sending the trap.
	Version *string `pulumi:"version"`
}

type SnmpTrapsState struct {
	// Encrypted password
	AuthPasswordencrypted pulumi.StringPtrInput
	// Specifies the protocol used to authenticate the user.
	AuthProtocol pulumi.StringPtrInput
	// Specifies the community string used for this trap.
	Community pulumi.StringPtrInput
	// The port that the trap will be sent to.
	Description pulumi.StringPtrInput
	// Specifies the authoritative security engine for SNMPv3.
	EngineId pulumi.StringPtrInput
	// The host the trap will be sent to.
	Host pulumi.StringPtrInput
	// Name of the snmp trap.
	Name pulumi.StringPtrInput
	// User defined description.
	Port pulumi.IntPtrInput
	// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
	PrivacyPassword pulumi.StringPtrInput
	// Specifies the encrypted password used to encrypt traffic.
	PrivacyPasswordEncrypted pulumi.StringPtrInput
	// Specifies the protocol used to encrypt traffic.
	PrivacyProtocol pulumi.StringPtrInput
	// Specifies whether or not traffic is encrypted and whether or not authentication is required.
	SecurityLevel pulumi.StringPtrInput
	// Security name used in conjunction with SNMPv3.
	SecurityName pulumi.StringPtrInput
	// SNMP version used for sending the trap.
	Version pulumi.StringPtrInput
}

func (SnmpTrapsState) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpTrapsState)(nil)).Elem()
}

type snmpTrapsArgs struct {
	// Encrypted password
	AuthPasswordencrypted *string `pulumi:"authPasswordencrypted"`
	// Specifies the protocol used to authenticate the user.
	AuthProtocol *string `pulumi:"authProtocol"`
	// Specifies the community string used for this trap.
	Community *string `pulumi:"community"`
	// The port that the trap will be sent to.
	Description *string `pulumi:"description"`
	// Specifies the authoritative security engine for SNMPv3.
	EngineId *string `pulumi:"engineId"`
	// The host the trap will be sent to.
	Host *string `pulumi:"host"`
	// Name of the snmp trap.
	Name *string `pulumi:"name"`
	// User defined description.
	Port *int `pulumi:"port"`
	// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
	PrivacyPassword *string `pulumi:"privacyPassword"`
	// Specifies the encrypted password used to encrypt traffic.
	PrivacyPasswordEncrypted *string `pulumi:"privacyPasswordEncrypted"`
	// Specifies the protocol used to encrypt traffic.
	PrivacyProtocol *string `pulumi:"privacyProtocol"`
	// Specifies whether or not traffic is encrypted and whether or not authentication is required.
	SecurityLevel *string `pulumi:"securityLevel"`
	// Security name used in conjunction with SNMPv3.
	SecurityName *string `pulumi:"securityName"`
	// SNMP version used for sending the trap.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a SnmpTraps resource.
type SnmpTrapsArgs struct {
	// Encrypted password
	AuthPasswordencrypted pulumi.StringPtrInput
	// Specifies the protocol used to authenticate the user.
	AuthProtocol pulumi.StringPtrInput
	// Specifies the community string used for this trap.
	Community pulumi.StringPtrInput
	// The port that the trap will be sent to.
	Description pulumi.StringPtrInput
	// Specifies the authoritative security engine for SNMPv3.
	EngineId pulumi.StringPtrInput
	// The host the trap will be sent to.
	Host pulumi.StringPtrInput
	// Name of the snmp trap.
	Name pulumi.StringPtrInput
	// User defined description.
	Port pulumi.IntPtrInput
	// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
	PrivacyPassword pulumi.StringPtrInput
	// Specifies the encrypted password used to encrypt traffic.
	PrivacyPasswordEncrypted pulumi.StringPtrInput
	// Specifies the protocol used to encrypt traffic.
	PrivacyProtocol pulumi.StringPtrInput
	// Specifies whether or not traffic is encrypted and whether or not authentication is required.
	SecurityLevel pulumi.StringPtrInput
	// Security name used in conjunction with SNMPv3.
	SecurityName pulumi.StringPtrInput
	// SNMP version used for sending the trap.
	Version pulumi.StringPtrInput
}

func (SnmpTrapsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpTrapsArgs)(nil)).Elem()
}

type SnmpTrapsInput interface {
	pulumi.Input

	ToSnmpTrapsOutput() SnmpTrapsOutput
	ToSnmpTrapsOutputWithContext(ctx context.Context) SnmpTrapsOutput
}

func (*SnmpTraps) ElementType() reflect.Type {
	return reflect.TypeOf((*SnmpTraps)(nil))
}

func (i *SnmpTraps) ToSnmpTrapsOutput() SnmpTrapsOutput {
	return i.ToSnmpTrapsOutputWithContext(context.Background())
}

func (i *SnmpTraps) ToSnmpTrapsOutputWithContext(ctx context.Context) SnmpTrapsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpTrapsOutput)
}

func (i *SnmpTraps) ToSnmpTrapsPtrOutput() SnmpTrapsPtrOutput {
	return i.ToSnmpTrapsPtrOutputWithContext(context.Background())
}

func (i *SnmpTraps) ToSnmpTrapsPtrOutputWithContext(ctx context.Context) SnmpTrapsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpTrapsPtrOutput)
}

type SnmpTrapsPtrInput interface {
	pulumi.Input

	ToSnmpTrapsPtrOutput() SnmpTrapsPtrOutput
	ToSnmpTrapsPtrOutputWithContext(ctx context.Context) SnmpTrapsPtrOutput
}

type snmpTrapsPtrType SnmpTrapsArgs

func (*snmpTrapsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnmpTraps)(nil))
}

func (i *snmpTrapsPtrType) ToSnmpTrapsPtrOutput() SnmpTrapsPtrOutput {
	return i.ToSnmpTrapsPtrOutputWithContext(context.Background())
}

func (i *snmpTrapsPtrType) ToSnmpTrapsPtrOutputWithContext(ctx context.Context) SnmpTrapsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpTrapsPtrOutput)
}

// SnmpTrapsArrayInput is an input type that accepts SnmpTrapsArray and SnmpTrapsArrayOutput values.
// You can construct a concrete instance of `SnmpTrapsArrayInput` via:
//
//          SnmpTrapsArray{ SnmpTrapsArgs{...} }
type SnmpTrapsArrayInput interface {
	pulumi.Input

	ToSnmpTrapsArrayOutput() SnmpTrapsArrayOutput
	ToSnmpTrapsArrayOutputWithContext(context.Context) SnmpTrapsArrayOutput
}

type SnmpTrapsArray []SnmpTrapsInput

func (SnmpTrapsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SnmpTraps)(nil))
}

func (i SnmpTrapsArray) ToSnmpTrapsArrayOutput() SnmpTrapsArrayOutput {
	return i.ToSnmpTrapsArrayOutputWithContext(context.Background())
}

func (i SnmpTrapsArray) ToSnmpTrapsArrayOutputWithContext(ctx context.Context) SnmpTrapsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpTrapsArrayOutput)
}

// SnmpTrapsMapInput is an input type that accepts SnmpTrapsMap and SnmpTrapsMapOutput values.
// You can construct a concrete instance of `SnmpTrapsMapInput` via:
//
//          SnmpTrapsMap{ "key": SnmpTrapsArgs{...} }
type SnmpTrapsMapInput interface {
	pulumi.Input

	ToSnmpTrapsMapOutput() SnmpTrapsMapOutput
	ToSnmpTrapsMapOutputWithContext(context.Context) SnmpTrapsMapOutput
}

type SnmpTrapsMap map[string]SnmpTrapsInput

func (SnmpTrapsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SnmpTraps)(nil))
}

func (i SnmpTrapsMap) ToSnmpTrapsMapOutput() SnmpTrapsMapOutput {
	return i.ToSnmpTrapsMapOutputWithContext(context.Background())
}

func (i SnmpTrapsMap) ToSnmpTrapsMapOutputWithContext(ctx context.Context) SnmpTrapsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpTrapsMapOutput)
}

type SnmpTrapsOutput struct {
	*pulumi.OutputState
}

func (SnmpTrapsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnmpTraps)(nil))
}

func (o SnmpTrapsOutput) ToSnmpTrapsOutput() SnmpTrapsOutput {
	return o
}

func (o SnmpTrapsOutput) ToSnmpTrapsOutputWithContext(ctx context.Context) SnmpTrapsOutput {
	return o
}

func (o SnmpTrapsOutput) ToSnmpTrapsPtrOutput() SnmpTrapsPtrOutput {
	return o.ToSnmpTrapsPtrOutputWithContext(context.Background())
}

func (o SnmpTrapsOutput) ToSnmpTrapsPtrOutputWithContext(ctx context.Context) SnmpTrapsPtrOutput {
	return o.ApplyT(func(v SnmpTraps) *SnmpTraps {
		return &v
	}).(SnmpTrapsPtrOutput)
}

type SnmpTrapsPtrOutput struct {
	*pulumi.OutputState
}

func (SnmpTrapsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnmpTraps)(nil))
}

func (o SnmpTrapsPtrOutput) ToSnmpTrapsPtrOutput() SnmpTrapsPtrOutput {
	return o
}

func (o SnmpTrapsPtrOutput) ToSnmpTrapsPtrOutputWithContext(ctx context.Context) SnmpTrapsPtrOutput {
	return o
}

type SnmpTrapsArrayOutput struct{ *pulumi.OutputState }

func (SnmpTrapsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SnmpTraps)(nil))
}

func (o SnmpTrapsArrayOutput) ToSnmpTrapsArrayOutput() SnmpTrapsArrayOutput {
	return o
}

func (o SnmpTrapsArrayOutput) ToSnmpTrapsArrayOutputWithContext(ctx context.Context) SnmpTrapsArrayOutput {
	return o
}

func (o SnmpTrapsArrayOutput) Index(i pulumi.IntInput) SnmpTrapsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SnmpTraps {
		return vs[0].([]SnmpTraps)[vs[1].(int)]
	}).(SnmpTrapsOutput)
}

type SnmpTrapsMapOutput struct{ *pulumi.OutputState }

func (SnmpTrapsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SnmpTraps)(nil))
}

func (o SnmpTrapsMapOutput) ToSnmpTrapsMapOutput() SnmpTrapsMapOutput {
	return o
}

func (o SnmpTrapsMapOutput) ToSnmpTrapsMapOutputWithContext(ctx context.Context) SnmpTrapsMapOutput {
	return o
}

func (o SnmpTrapsMapOutput) MapIndex(k pulumi.StringInput) SnmpTrapsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SnmpTraps {
		return vs[0].(map[string]SnmpTraps)[vs[1].(string)]
	}).(SnmpTrapsOutput)
}

func init() {
	pulumi.RegisterOutputType(SnmpTrapsOutput{})
	pulumi.RegisterOutputType(SnmpTrapsPtrOutput{})
	pulumi.RegisterOutputType(SnmpTrapsArrayOutput{})
	pulumi.RegisterOutputType(SnmpTrapsMapOutput{})
}
