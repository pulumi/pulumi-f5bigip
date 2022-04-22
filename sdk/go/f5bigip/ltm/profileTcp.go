// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.ProfileTcp` Configures a custom profileTcp for use by health checks.
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewProfileTcp(ctx, "sanjose-tcp-lan-profile", &ltm.ProfileTcpArgs{
// 			CloseWaitTimeout:  pulumi.Int(5),
// 			DeferredAccept:    pulumi.String("enabled"),
// 			FastOpen:          pulumi.String("enabled"),
// 			Finwait2timeout:   pulumi.Int(5),
// 			FinwaitTimeout:    pulumi.Int(300),
// 			IdleTimeout:       pulumi.Int(200),
// 			KeepaliveInterval: pulumi.Int(1700),
// 			Name:              pulumi.String("/Common/sanjose-tcp-lan-profile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProfileTcp struct {
	pulumi.CustomResourceState

	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout pulumi.IntOutput `pulumi:"closeWaitTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept pulumi.StringOutput `pulumi:"deferredAccept"`
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen pulumi.StringOutput `pulumi:"fastOpen"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout pulumi.IntOutput `pulumi:"finwait2timeout"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout pulumi.IntOutput `pulumi:"finwaitTimeout"`
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval pulumi.IntOutput `pulumi:"keepaliveInterval"`
	// Name of the profile_tcp
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
}

// NewProfileTcp registers a new resource with the given unique name, arguments, and options.
func NewProfileTcp(ctx *pulumi.Context,
	name string, args *ProfileTcpArgs, opts ...pulumi.ResourceOption) (*ProfileTcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource ProfileTcp
	err := ctx.RegisterResource("f5bigip:ltm/profileTcp:ProfileTcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileTcp gets an existing ProfileTcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileTcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileTcpState, opts ...pulumi.ResourceOption) (*ProfileTcp, error) {
	var resource ProfileTcp
	err := ctx.ReadResource("f5bigip:ltm/profileTcp:ProfileTcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileTcp resources.
type profileTcpState struct {
	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout *int `pulumi:"closeWaitTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept *string `pulumi:"deferredAccept"`
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen *string `pulumi:"fastOpen"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout *int `pulumi:"finwait2timeout"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout *int `pulumi:"finwaitTimeout"`
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval *int `pulumi:"keepaliveInterval"`
	// Name of the profile_tcp
	Name *string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
}

type ProfileTcpState struct {
	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept pulumi.StringPtrInput
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout pulumi.IntPtrInput
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout pulumi.IntPtrInput
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval pulumi.IntPtrInput
	// Name of the profile_tcp
	Name pulumi.StringPtrInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
}

func (ProfileTcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileTcpState)(nil)).Elem()
}

type profileTcpArgs struct {
	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout *int `pulumi:"closeWaitTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept *string `pulumi:"deferredAccept"`
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen *string `pulumi:"fastOpen"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout *int `pulumi:"finwait2timeout"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout *int `pulumi:"finwaitTimeout"`
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval *int `pulumi:"keepaliveInterval"`
	// Name of the profile_tcp
	Name string `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition *string `pulumi:"partition"`
}

// The set of arguments for constructing a ProfileTcp resource.
type ProfileTcpArgs struct {
	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept pulumi.StringPtrInput
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen pulumi.StringPtrInput
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout pulumi.IntPtrInput
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout pulumi.IntPtrInput
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval pulumi.IntPtrInput
	// Name of the profile_tcp
	Name pulumi.StringInput
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrInput
}

func (ProfileTcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileTcpArgs)(nil)).Elem()
}

type ProfileTcpInput interface {
	pulumi.Input

	ToProfileTcpOutput() ProfileTcpOutput
	ToProfileTcpOutputWithContext(ctx context.Context) ProfileTcpOutput
}

func (*ProfileTcp) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileTcp)(nil)).Elem()
}

func (i *ProfileTcp) ToProfileTcpOutput() ProfileTcpOutput {
	return i.ToProfileTcpOutputWithContext(context.Background())
}

func (i *ProfileTcp) ToProfileTcpOutputWithContext(ctx context.Context) ProfileTcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileTcpOutput)
}

// ProfileTcpArrayInput is an input type that accepts ProfileTcpArray and ProfileTcpArrayOutput values.
// You can construct a concrete instance of `ProfileTcpArrayInput` via:
//
//          ProfileTcpArray{ ProfileTcpArgs{...} }
type ProfileTcpArrayInput interface {
	pulumi.Input

	ToProfileTcpArrayOutput() ProfileTcpArrayOutput
	ToProfileTcpArrayOutputWithContext(context.Context) ProfileTcpArrayOutput
}

type ProfileTcpArray []ProfileTcpInput

func (ProfileTcpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileTcp)(nil)).Elem()
}

func (i ProfileTcpArray) ToProfileTcpArrayOutput() ProfileTcpArrayOutput {
	return i.ToProfileTcpArrayOutputWithContext(context.Background())
}

func (i ProfileTcpArray) ToProfileTcpArrayOutputWithContext(ctx context.Context) ProfileTcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileTcpArrayOutput)
}

// ProfileTcpMapInput is an input type that accepts ProfileTcpMap and ProfileTcpMapOutput values.
// You can construct a concrete instance of `ProfileTcpMapInput` via:
//
//          ProfileTcpMap{ "key": ProfileTcpArgs{...} }
type ProfileTcpMapInput interface {
	pulumi.Input

	ToProfileTcpMapOutput() ProfileTcpMapOutput
	ToProfileTcpMapOutputWithContext(context.Context) ProfileTcpMapOutput
}

type ProfileTcpMap map[string]ProfileTcpInput

func (ProfileTcpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileTcp)(nil)).Elem()
}

func (i ProfileTcpMap) ToProfileTcpMapOutput() ProfileTcpMapOutput {
	return i.ToProfileTcpMapOutputWithContext(context.Background())
}

func (i ProfileTcpMap) ToProfileTcpMapOutputWithContext(ctx context.Context) ProfileTcpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileTcpMapOutput)
}

type ProfileTcpOutput struct{ *pulumi.OutputState }

func (ProfileTcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileTcp)(nil)).Elem()
}

func (o ProfileTcpOutput) ToProfileTcpOutput() ProfileTcpOutput {
	return o
}

func (o ProfileTcpOutput) ToProfileTcpOutputWithContext(ctx context.Context) ProfileTcpOutput {
	return o
}

type ProfileTcpArrayOutput struct{ *pulumi.OutputState }

func (ProfileTcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileTcp)(nil)).Elem()
}

func (o ProfileTcpArrayOutput) ToProfileTcpArrayOutput() ProfileTcpArrayOutput {
	return o
}

func (o ProfileTcpArrayOutput) ToProfileTcpArrayOutputWithContext(ctx context.Context) ProfileTcpArrayOutput {
	return o
}

func (o ProfileTcpArrayOutput) Index(i pulumi.IntInput) ProfileTcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProfileTcp {
		return vs[0].([]*ProfileTcp)[vs[1].(int)]
	}).(ProfileTcpOutput)
}

type ProfileTcpMapOutput struct{ *pulumi.OutputState }

func (ProfileTcpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileTcp)(nil)).Elem()
}

func (o ProfileTcpMapOutput) ToProfileTcpMapOutput() ProfileTcpMapOutput {
	return o
}

func (o ProfileTcpMapOutput) ToProfileTcpMapOutputWithContext(ctx context.Context) ProfileTcpMapOutput {
	return o
}

func (o ProfileTcpMapOutput) MapIndex(k pulumi.StringInput) ProfileTcpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProfileTcp {
		return vs[0].(map[string]*ProfileTcp)[vs[1].(string)]
	}).(ProfileTcpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileTcpInput)(nil)).Elem(), &ProfileTcp{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileTcpArrayInput)(nil)).Elem(), ProfileTcpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileTcpMapInput)(nil)).Elem(), ProfileTcpMap{})
	pulumi.RegisterOutputType(ProfileTcpOutput{})
	pulumi.RegisterOutputType(ProfileTcpArrayOutput{})
	pulumi.RegisterOutputType(ProfileTcpMapOutput{})
}
