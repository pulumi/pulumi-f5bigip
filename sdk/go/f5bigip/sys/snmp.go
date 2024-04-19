// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `sys.Snmp` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/sys"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sys.NewSnmp(ctx, "snmp", &sys.SnmpArgs{
//				SysContact:  pulumi.String(" NetOPsAdmin s.shitole@f5.com"),
//				SysLocation: pulumi.String("SeattleHQ"),
//				Allowedaddresses: pulumi.StringArray{
//					pulumi.String("202.10.10.2"),
//				},
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
type Snmp struct {
	pulumi.CustomResourceState

	// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
	Allowedaddresses pulumi.StringArrayOutput `pulumi:"allowedaddresses"`
	// Specifies the contact information for the system administrator.
	SysContact pulumi.StringPtrOutput `pulumi:"sysContact"`
	// Describes the system's physical location.
	SysLocation pulumi.StringPtrOutput `pulumi:"sysLocation"`
}

// NewSnmp registers a new resource with the given unique name, arguments, and options.
func NewSnmp(ctx *pulumi.Context,
	name string, args *SnmpArgs, opts ...pulumi.ResourceOption) (*Snmp, error) {
	if args == nil {
		args = &SnmpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snmp
	err := ctx.RegisterResource("f5bigip:sys/snmp:Snmp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnmp gets an existing Snmp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnmp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnmpState, opts ...pulumi.ResourceOption) (*Snmp, error) {
	var resource Snmp
	err := ctx.ReadResource("f5bigip:sys/snmp:Snmp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snmp resources.
type snmpState struct {
	// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
	Allowedaddresses []string `pulumi:"allowedaddresses"`
	// Specifies the contact information for the system administrator.
	SysContact *string `pulumi:"sysContact"`
	// Describes the system's physical location.
	SysLocation *string `pulumi:"sysLocation"`
}

type SnmpState struct {
	// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
	Allowedaddresses pulumi.StringArrayInput
	// Specifies the contact information for the system administrator.
	SysContact pulumi.StringPtrInput
	// Describes the system's physical location.
	SysLocation pulumi.StringPtrInput
}

func (SnmpState) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpState)(nil)).Elem()
}

type snmpArgs struct {
	// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
	Allowedaddresses []string `pulumi:"allowedaddresses"`
	// Specifies the contact information for the system administrator.
	SysContact *string `pulumi:"sysContact"`
	// Describes the system's physical location.
	SysLocation *string `pulumi:"sysLocation"`
}

// The set of arguments for constructing a Snmp resource.
type SnmpArgs struct {
	// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
	Allowedaddresses pulumi.StringArrayInput
	// Specifies the contact information for the system administrator.
	SysContact pulumi.StringPtrInput
	// Describes the system's physical location.
	SysLocation pulumi.StringPtrInput
}

func (SnmpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snmpArgs)(nil)).Elem()
}

type SnmpInput interface {
	pulumi.Input

	ToSnmpOutput() SnmpOutput
	ToSnmpOutputWithContext(ctx context.Context) SnmpOutput
}

func (*Snmp) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmp)(nil)).Elem()
}

func (i *Snmp) ToSnmpOutput() SnmpOutput {
	return i.ToSnmpOutputWithContext(context.Background())
}

func (i *Snmp) ToSnmpOutputWithContext(ctx context.Context) SnmpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpOutput)
}

// SnmpArrayInput is an input type that accepts SnmpArray and SnmpArrayOutput values.
// You can construct a concrete instance of `SnmpArrayInput` via:
//
//	SnmpArray{ SnmpArgs{...} }
type SnmpArrayInput interface {
	pulumi.Input

	ToSnmpArrayOutput() SnmpArrayOutput
	ToSnmpArrayOutputWithContext(context.Context) SnmpArrayOutput
}

type SnmpArray []SnmpInput

func (SnmpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmp)(nil)).Elem()
}

func (i SnmpArray) ToSnmpArrayOutput() SnmpArrayOutput {
	return i.ToSnmpArrayOutputWithContext(context.Background())
}

func (i SnmpArray) ToSnmpArrayOutputWithContext(ctx context.Context) SnmpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpArrayOutput)
}

// SnmpMapInput is an input type that accepts SnmpMap and SnmpMapOutput values.
// You can construct a concrete instance of `SnmpMapInput` via:
//
//	SnmpMap{ "key": SnmpArgs{...} }
type SnmpMapInput interface {
	pulumi.Input

	ToSnmpMapOutput() SnmpMapOutput
	ToSnmpMapOutputWithContext(context.Context) SnmpMapOutput
}

type SnmpMap map[string]SnmpInput

func (SnmpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmp)(nil)).Elem()
}

func (i SnmpMap) ToSnmpMapOutput() SnmpMapOutput {
	return i.ToSnmpMapOutputWithContext(context.Background())
}

func (i SnmpMap) ToSnmpMapOutputWithContext(ctx context.Context) SnmpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnmpMapOutput)
}

type SnmpOutput struct{ *pulumi.OutputState }

func (SnmpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snmp)(nil)).Elem()
}

func (o SnmpOutput) ToSnmpOutput() SnmpOutput {
	return o
}

func (o SnmpOutput) ToSnmpOutputWithContext(ctx context.Context) SnmpOutput {
	return o
}

// Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
func (o SnmpOutput) Allowedaddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snmp) pulumi.StringArrayOutput { return v.Allowedaddresses }).(pulumi.StringArrayOutput)
}

// Specifies the contact information for the system administrator.
func (o SnmpOutput) SysContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snmp) pulumi.StringPtrOutput { return v.SysContact }).(pulumi.StringPtrOutput)
}

// Describes the system's physical location.
func (o SnmpOutput) SysLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snmp) pulumi.StringPtrOutput { return v.SysLocation }).(pulumi.StringPtrOutput)
}

type SnmpArrayOutput struct{ *pulumi.OutputState }

func (SnmpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snmp)(nil)).Elem()
}

func (o SnmpArrayOutput) ToSnmpArrayOutput() SnmpArrayOutput {
	return o
}

func (o SnmpArrayOutput) ToSnmpArrayOutputWithContext(ctx context.Context) SnmpArrayOutput {
	return o
}

func (o SnmpArrayOutput) Index(i pulumi.IntInput) SnmpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snmp {
		return vs[0].([]*Snmp)[vs[1].(int)]
	}).(SnmpOutput)
}

type SnmpMapOutput struct{ *pulumi.OutputState }

func (SnmpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snmp)(nil)).Elem()
}

func (o SnmpMapOutput) ToSnmpMapOutput() SnmpMapOutput {
	return o
}

func (o SnmpMapOutput) ToSnmpMapOutputWithContext(ctx context.Context) SnmpMapOutput {
	return o
}

func (o SnmpMapOutput) MapIndex(k pulumi.StringInput) SnmpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snmp {
		return vs[0].(map[string]*Snmp)[vs[1].(string)]
	}).(SnmpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpInput)(nil)).Elem(), &Snmp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpArrayInput)(nil)).Elem(), SnmpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnmpMapInput)(nil)).Elem(), SnmpMap{})
	pulumi.RegisterOutputType(SnmpOutput{})
	pulumi.RegisterOutputType(SnmpArrayOutput{})
	pulumi.RegisterOutputType(SnmpMapOutput{})
}
