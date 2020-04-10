// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.ProfileTcp` Configures a custom profileTcp for use by health checks.
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
type ProfileTcp struct {
	pulumi.CustomResourceState

	// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
	CloseWaitTimeout pulumi.IntPtrOutput `pulumi:"closeWaitTimeout"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
	DeferredAccept pulumi.StringPtrOutput `pulumi:"deferredAccept"`
	// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
	FastOpen pulumi.StringPtrOutput `pulumi:"fastOpen"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
	Finwait2timeout pulumi.IntPtrOutput `pulumi:"finwait2timeout"`
	// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
	FinwaitTimeout pulumi.IntPtrOutput `pulumi:"finwaitTimeout"`
	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrOutput `pulumi:"idleTimeout"`
	// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
	KeepaliveInterval pulumi.IntPtrOutput `pulumi:"keepaliveInterval"`
	// Name of the profile_tcp
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which this profile resides
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
}

// NewProfileTcp registers a new resource with the given unique name, arguments, and options.
func NewProfileTcp(ctx *pulumi.Context,
	name string, args *ProfileTcpArgs, opts ...pulumi.ResourceOption) (*ProfileTcp, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileTcpArgs{}
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
