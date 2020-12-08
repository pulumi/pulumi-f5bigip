// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.ProfileFastHttp` Configures a custom profileFasthttp for use by health checks.
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
// 		_, err := ltm.NewProfileFastHttp(ctx, "sjfasthttpprofile", &ltm.ProfileFastHttpArgs{
// 			ConnpoolMaxreuse:            pulumi.Int(2),
// 			ConnpoolMaxsize:             pulumi.Int(2048),
// 			ConnpoolMinsize:             pulumi.Int(0),
// 			ConnpoolReplenish:           pulumi.String("enabled"),
// 			ConnpoolStep:                pulumi.Int(4),
// 			ConnpoolidleTimeoutoverride: pulumi.Int(0),
// 			DefaultsFrom:                pulumi.String("/Common/fasthttp"),
// 			Forcehttp10response:         pulumi.String("disabled"),
// 			IdleTimeout:                 pulumi.Int(300),
// 			MaxheaderSize:               pulumi.Int(32768),
// 			Name:                        pulumi.String("/Common/sjfasthttpprofile"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProfileFastHttp struct {
	pulumi.CustomResourceState

	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxreuse pulumi.IntPtrOutput `pulumi:"connpoolMaxreuse"`
	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxsize pulumi.IntPtrOutput `pulumi:"connpoolMaxsize"`
	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinsize pulumi.IntPtrOutput `pulumi:"connpoolMinsize"`
	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
	ConnpoolReplenish pulumi.StringPtrOutput `pulumi:"connpoolReplenish"`
	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep pulumi.IntPtrOutput `pulumi:"connpoolStep"`
	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolidleTimeoutoverride pulumi.IntPtrOutput `pulumi:"connpoolidleTimeoutoverride"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	Forcehttp10response pulumi.StringPtrOutput `pulumi:"forcehttp10response"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrOutput `pulumi:"idleTimeout"`
	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxheaderSize pulumi.IntPtrOutput `pulumi:"maxheaderSize"`
	// Name of the profile_fasthttp
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewProfileFastHttp registers a new resource with the given unique name, arguments, and options.
func NewProfileFastHttp(ctx *pulumi.Context,
	name string, args *ProfileFastHttpArgs, opts ...pulumi.ResourceOption) (*ProfileFastHttp, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileFastHttpArgs{}
	}
	var resource ProfileFastHttp
	err := ctx.RegisterResource("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileFastHttp gets an existing ProfileFastHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileFastHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileFastHttpState, opts ...pulumi.ResourceOption) (*ProfileFastHttp, error) {
	var resource ProfileFastHttp
	err := ctx.ReadResource("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileFastHttp resources.
type profileFastHttpState struct {
	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxreuse *int `pulumi:"connpoolMaxreuse"`
	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxsize *int `pulumi:"connpoolMaxsize"`
	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinsize *int `pulumi:"connpoolMinsize"`
	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
	ConnpoolReplenish *string `pulumi:"connpoolReplenish"`
	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep *int `pulumi:"connpoolStep"`
	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolidleTimeoutoverride *int `pulumi:"connpoolidleTimeoutoverride"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	Forcehttp10response *string `pulumi:"forcehttp10response"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxheaderSize *int `pulumi:"maxheaderSize"`
	// Name of the profile_fasthttp
	Name *string `pulumi:"name"`
}

type ProfileFastHttpState struct {
	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxreuse pulumi.IntPtrInput
	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxsize pulumi.IntPtrInput
	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinsize pulumi.IntPtrInput
	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
	ConnpoolReplenish pulumi.StringPtrInput
	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep pulumi.IntPtrInput
	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolidleTimeoutoverride pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	Forcehttp10response pulumi.StringPtrInput
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxheaderSize pulumi.IntPtrInput
	// Name of the profile_fasthttp
	Name pulumi.StringPtrInput
}

func (ProfileFastHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileFastHttpState)(nil)).Elem()
}

type profileFastHttpArgs struct {
	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxreuse *int `pulumi:"connpoolMaxreuse"`
	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxsize *int `pulumi:"connpoolMaxsize"`
	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinsize *int `pulumi:"connpoolMinsize"`
	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
	ConnpoolReplenish *string `pulumi:"connpoolReplenish"`
	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep *int `pulumi:"connpoolStep"`
	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolidleTimeoutoverride *int `pulumi:"connpoolidleTimeoutoverride"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	Forcehttp10response *string `pulumi:"forcehttp10response"`
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxheaderSize *int `pulumi:"maxheaderSize"`
	// Name of the profile_fasthttp
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a ProfileFastHttp resource.
type ProfileFastHttpArgs struct {
	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxreuse pulumi.IntPtrInput
	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxsize pulumi.IntPtrInput
	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinsize pulumi.IntPtrInput
	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
	ConnpoolReplenish pulumi.StringPtrInput
	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep pulumi.IntPtrInput
	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolidleTimeoutoverride pulumi.IntPtrInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	Forcehttp10response pulumi.StringPtrInput
	// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
	IdleTimeout pulumi.IntPtrInput
	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxheaderSize pulumi.IntPtrInput
	// Name of the profile_fasthttp
	Name pulumi.StringInput
}

func (ProfileFastHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileFastHttpArgs)(nil)).Elem()
}

type ProfileFastHttpInput interface {
	pulumi.Input

	ToProfileFastHttpOutput() ProfileFastHttpOutput
	ToProfileFastHttpOutputWithContext(ctx context.Context) ProfileFastHttpOutput
}

func (ProfileFastHttp) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFastHttp)(nil)).Elem()
}

func (i ProfileFastHttp) ToProfileFastHttpOutput() ProfileFastHttpOutput {
	return i.ToProfileFastHttpOutputWithContext(context.Background())
}

func (i ProfileFastHttp) ToProfileFastHttpOutputWithContext(ctx context.Context) ProfileFastHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileFastHttpOutput)
}

type ProfileFastHttpOutput struct {
	*pulumi.OutputState
}

func (ProfileFastHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileFastHttpOutput)(nil)).Elem()
}

func (o ProfileFastHttpOutput) ToProfileFastHttpOutput() ProfileFastHttpOutput {
	return o
}

func (o ProfileFastHttpOutput) ToProfileFastHttpOutputWithContext(ctx context.Context) ProfileFastHttpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfileFastHttpOutput{})
}
