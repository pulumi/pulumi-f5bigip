// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.ProfileFastL4` Configures a custom profileFastl4 for use by health checks.
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
// 		_, err := ltm.NewProfileFastL4(ctx, "profileFastl4", &ltm.ProfileFastL4Args{
// 			ClientTimeout:         pulumi.Int(40),
// 			DefaultsFrom:          pulumi.String("/Common/fastL4"),
// 			ExplicitflowMigration: pulumi.String("enabled"),
// 			HardwareSyncookie:     pulumi.String("enabled"),
// 			IdleTimeout:           pulumi.String("200"),
// 			IptosToclient:         pulumi.String("pass-through"),
// 			IptosToserver:         pulumi.String("pass-through"),
// 			KeepaliveInterval:     pulumi.String("disabled"),
// 			Name:                  pulumi.String("/Common/sjfastl4profile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// BIG-IP LTM fastl4 profiles can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import f5bigip:ltm/profileFastL4:ProfileFastL4 test-fastl4 /Common/test-fastl4
// ```
type ProfileFastL4 struct {
	pulumi.CustomResourceState

	// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
	ClientTimeout pulumi.IntOutput `pulumi:"clientTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
	ExplicitflowMigration pulumi.StringOutput `pulumi:"explicitflowMigration"`
	// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
	HardwareSyncookie pulumi.StringOutput `pulumi:"hardwareSyncookie"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.StringOutput `pulumi:"idleTimeout"`
	// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToclient pulumi.StringOutput `pulumi:"iptosToclient"`
	// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToserver pulumi.StringOutput `pulumi:"iptosToserver"`
	// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
	KeepaliveInterval pulumi.StringOutput `pulumi:"keepaliveInterval"`
	// Name of the profile_fastl4
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringOutput `pulumi:"partition"`
}

// NewProfileFastL4 registers a new resource with the given unique name, arguments, and options.
func NewProfileFastL4(ctx *pulumi.Context,
	name string, args *ProfileFastL4Args, opts ...pulumi.ResourceOption) (*ProfileFastL4, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileFastL4Args{}
	}
	var resource ProfileFastL4
	err := ctx.RegisterResource("f5bigip:ltm/profileFastL4:ProfileFastL4", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileFastL4 gets an existing ProfileFastL4 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileFastL4(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileFastL4State, opts ...pulumi.ResourceOption) (*ProfileFastL4, error) {
	var resource ProfileFastL4
	err := ctx.ReadResource("f5bigip:ltm/profileFastL4:ProfileFastL4", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileFastL4 resources.
type profileFastL4State struct {
	// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
	ClientTimeout *int `pulumi:"clientTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
	ExplicitflowMigration *string `pulumi:"explicitflowMigration"`
	// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
	HardwareSyncookie *string `pulumi:"hardwareSyncookie"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout *string `pulumi:"idleTimeout"`
	// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToclient *string `pulumi:"iptosToclient"`
	// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToserver *string `pulumi:"iptosToserver"`
	// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
	KeepaliveInterval *string `pulumi:"keepaliveInterval"`
	// Name of the profile_fastl4
	Name *string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
}

type ProfileFastL4State struct {
	// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
	ClientTimeout pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
	ExplicitflowMigration pulumi.StringPtrInput
	// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
	HardwareSyncookie pulumi.StringPtrInput
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.StringPtrInput
	// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToclient pulumi.StringPtrInput
	// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToserver pulumi.StringPtrInput
	// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
	KeepaliveInterval pulumi.StringPtrInput
	// Name of the profile_fastl4
	Name pulumi.StringPtrInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
}

func (ProfileFastL4State) ElementType() reflect.Type {
	return reflect.TypeOf((*profileFastL4State)(nil)).Elem()
}

type profileFastL4Args struct {
	// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
	ClientTimeout *int `pulumi:"clientTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
	ExplicitflowMigration *string `pulumi:"explicitflowMigration"`
	// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
	HardwareSyncookie *string `pulumi:"hardwareSyncookie"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout *string `pulumi:"idleTimeout"`
	// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToclient *string `pulumi:"iptosToclient"`
	// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToserver *string `pulumi:"iptosToserver"`
	// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
	KeepaliveInterval *string `pulumi:"keepaliveInterval"`
	// Name of the profile_fastl4
	Name string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
}

// The set of arguments for constructing a ProfileFastL4 resource.
type ProfileFastL4Args struct {
	// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
	ClientTimeout pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
	ExplicitflowMigration pulumi.StringPtrInput
	// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
	HardwareSyncookie pulumi.StringPtrInput
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.StringPtrInput
	// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToclient pulumi.StringPtrInput
	// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
	IptosToserver pulumi.StringPtrInput
	// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
	KeepaliveInterval pulumi.StringPtrInput
	// Name of the profile_fastl4
	Name pulumi.StringInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
}

func (ProfileFastL4Args) ElementType() reflect.Type {
	return reflect.TypeOf((*profileFastL4Args)(nil)).Elem()
}

type ProfileFastL4Input interface {
	pulumi.Input

	ToProfileFastL4Output() ProfileFastL4Output
	ToProfileFastL4OutputWithContext(ctx context.Context) ProfileFastL4Output
}

func (ProfileFastL4) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFastL4)(nil)).Elem()
}

func (i ProfileFastL4) ToProfileFastL4Output() ProfileFastL4Output {
	return i.ToProfileFastL4OutputWithContext(context.Background())
}

func (i ProfileFastL4) ToProfileFastL4OutputWithContext(ctx context.Context) ProfileFastL4Output {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFastL4Output)
}

type ProfileFastL4Output struct {
	*pulumi.OutputState
}

func (ProfileFastL4Output) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFastL4Output)(nil)).Elem()
}

func (o ProfileFastL4Output) ToProfileFastL4Output() ProfileFastL4Output {
	return o
}

func (o ProfileFastL4Output) ToProfileFastL4OutputWithContext(ctx context.Context) ProfileFastL4Output {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfileFastL4Output{})
}
