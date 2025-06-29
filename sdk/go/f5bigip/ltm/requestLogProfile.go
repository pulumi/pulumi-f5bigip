// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.RequestLogProfile` Resource used for Configures request logging using the Request Logging profile
//
// ## Example Usage
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
//			_, err := ltm.NewRequestLogProfile(ctx, "request-log-profile-tc1-child", &ltm.RequestLogProfileArgs{
//				Name:                     pulumi.String("/Common/request-log-profile-tc1-child"),
//				DefaultsFrom:             pulumi.Any(request_log_profile_tc1.Name),
//				RequestLogging:           pulumi.String("disabled"),
//				RequestlogPool:           pulumi.String("/Common/pool2"),
//				RequestlogErrorPool:      pulumi.String("/Common/pool1"),
//				RequestlogProtocol:       pulumi.String("mds-tcp"),
//				RequestlogErrorProtocol:  pulumi.String("mds-tcp"),
//				ResponselogProtocol:      pulumi.String("mds-tcp"),
//				ResponselogErrorProtocol: pulumi.String("mds-tcp"),
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
// ## Import
//
// BIG-IP LTM Request Log profiles can be imported using the `name`, e.g.
//
// bash
//
// ```sh
// $ pulumi import f5bigip:ltm/requestLogProfile:RequestLogProfile test-request-log /Common/test-request-log
// ```
type RequestLogProfile struct {
	pulumi.CustomResourceState

	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Specifies user-defined description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyResponse pulumi.StringPtrOutput `pulumi:"proxyResponse"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxycloseOnError pulumi.StringPtrOutput `pulumi:"proxycloseOnError"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyrespondOnLoggingerror pulumi.StringPtrOutput `pulumi:"proxyrespondOnLoggingerror"`
	// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	RequestLogging pulumi.StringOutput `pulumi:"requestLogging"`
	// Defines the pool associated with logging request errors. The default is None.
	RequestlogErrorPool pulumi.StringPtrOutput `pulumi:"requestlogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogErrorProtocol pulumi.StringPtrOutput `pulumi:"requestlogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	RequestlogErrorTemplate pulumi.StringPtrOutput `pulumi:"requestlogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
	RequestlogPool pulumi.StringPtrOutput `pulumi:"requestlogPool"`
	// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogProtocol pulumi.StringPtrOutput `pulumi:"requestlogProtocol"`
	// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	RequestlogTemplate pulumi.StringPtrOutput `pulumi:"requestlogTemplate"`
	// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	ResponseLogging pulumi.StringOutput `pulumi:"responseLogging"`
	// Defines the pool associated with logging response errors. The default is `none`.
	ResponselogErrorPool pulumi.StringPtrOutput `pulumi:"responselogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogErrorProtocol pulumi.StringPtrOutput `pulumi:"responselogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	ResponselogErrorTemplate pulumi.StringPtrOutput `pulumi:"responselogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
	ResponselogPool pulumi.StringPtrOutput `pulumi:"responselogPool"`
	// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogProtocol pulumi.StringPtrOutput `pulumi:"responselogProtocol"`
	// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	ResponselogTemplate pulumi.StringPtrOutput `pulumi:"responselogTemplate"`
}

// NewRequestLogProfile registers a new resource with the given unique name, arguments, and options.
func NewRequestLogProfile(ctx *pulumi.Context,
	name string, args *RequestLogProfileArgs, opts ...pulumi.ResourceOption) (*RequestLogProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RequestLogProfile
	err := ctx.RegisterResource("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestLogProfile gets an existing RequestLogProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestLogProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestLogProfileState, opts ...pulumi.ResourceOption) (*RequestLogProfile, error) {
	var resource RequestLogProfile
	err := ctx.ReadResource("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestLogProfile resources.
type requestLogProfileState struct {
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies user-defined description.
	Description *string `pulumi:"description"`
	// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
	Name *string `pulumi:"name"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyResponse *string `pulumi:"proxyResponse"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxycloseOnError *string `pulumi:"proxycloseOnError"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyrespondOnLoggingerror *string `pulumi:"proxyrespondOnLoggingerror"`
	// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	RequestLogging *string `pulumi:"requestLogging"`
	// Defines the pool associated with logging request errors. The default is None.
	RequestlogErrorPool *string `pulumi:"requestlogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogErrorProtocol *string `pulumi:"requestlogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	RequestlogErrorTemplate *string `pulumi:"requestlogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
	RequestlogPool *string `pulumi:"requestlogPool"`
	// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogProtocol *string `pulumi:"requestlogProtocol"`
	// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	RequestlogTemplate *string `pulumi:"requestlogTemplate"`
	// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	ResponseLogging *string `pulumi:"responseLogging"`
	// Defines the pool associated with logging response errors. The default is `none`.
	ResponselogErrorPool *string `pulumi:"responselogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogErrorProtocol *string `pulumi:"responselogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	ResponselogErrorTemplate *string `pulumi:"responselogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
	ResponselogPool *string `pulumi:"responselogPool"`
	// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogProtocol *string `pulumi:"responselogProtocol"`
	// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	ResponselogTemplate *string `pulumi:"responselogTemplate"`
}

type RequestLogProfileState struct {
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies user-defined description.
	Description pulumi.StringPtrInput
	// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
	Name pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxyResponse pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxycloseOnError pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxyrespondOnLoggingerror pulumi.StringPtrInput
	// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	RequestLogging pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	RequestlogErrorPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogErrorProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged for request errors.
	RequestlogErrorTemplate pulumi.StringPtrInput
	// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
	RequestlogPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	RequestlogTemplate pulumi.StringPtrInput
	// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	ResponseLogging pulumi.StringPtrInput
	// Defines the pool associated with logging response errors. The default is `none`.
	ResponselogErrorPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogErrorProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged for request errors.
	ResponselogErrorTemplate pulumi.StringPtrInput
	// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
	ResponselogPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	ResponselogTemplate pulumi.StringPtrInput
}

func (RequestLogProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestLogProfileState)(nil)).Elem()
}

type requestLogProfileArgs struct {
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specifies user-defined description.
	Description *string `pulumi:"description"`
	// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
	Name string `pulumi:"name"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyResponse *string `pulumi:"proxyResponse"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxycloseOnError *string `pulumi:"proxycloseOnError"`
	// Defines the pool associated with logging request errors. The default is None.
	ProxyrespondOnLoggingerror *string `pulumi:"proxyrespondOnLoggingerror"`
	// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	RequestLogging *string `pulumi:"requestLogging"`
	// Defines the pool associated with logging request errors. The default is None.
	RequestlogErrorPool *string `pulumi:"requestlogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogErrorProtocol *string `pulumi:"requestlogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	RequestlogErrorTemplate *string `pulumi:"requestlogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
	RequestlogPool *string `pulumi:"requestlogPool"`
	// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogProtocol *string `pulumi:"requestlogProtocol"`
	// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	RequestlogTemplate *string `pulumi:"requestlogTemplate"`
	// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	ResponseLogging *string `pulumi:"responseLogging"`
	// Defines the pool associated with logging response errors. The default is `none`.
	ResponselogErrorPool *string `pulumi:"responselogErrorPool"`
	// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogErrorProtocol *string `pulumi:"responselogErrorProtocol"`
	// Specifies the directives and entries to be logged for request errors.
	ResponselogErrorTemplate *string `pulumi:"responselogErrorTemplate"`
	// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
	ResponselogPool *string `pulumi:"responselogPool"`
	// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogProtocol *string `pulumi:"responselogProtocol"`
	// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	ResponselogTemplate *string `pulumi:"responselogTemplate"`
}

// The set of arguments for constructing a RequestLogProfile resource.
type RequestLogProfileArgs struct {
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
	DefaultsFrom pulumi.StringPtrInput
	// Specifies user-defined description.
	Description pulumi.StringPtrInput
	// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
	Name pulumi.StringInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxyResponse pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxycloseOnError pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	ProxyrespondOnLoggingerror pulumi.StringPtrInput
	// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	RequestLogging pulumi.StringPtrInput
	// Defines the pool associated with logging request errors. The default is None.
	RequestlogErrorPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogErrorProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged for request errors.
	RequestlogErrorTemplate pulumi.StringPtrInput
	// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
	RequestlogPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	RequestlogProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	RequestlogTemplate pulumi.StringPtrInput
	// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
	ResponseLogging pulumi.StringPtrInput
	// Defines the pool associated with logging response errors. The default is `none`.
	ResponselogErrorPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogErrorProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged for request errors.
	ResponselogErrorTemplate pulumi.StringPtrInput
	// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
	ResponselogPool pulumi.StringPtrInput
	// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
	ResponselogProtocol pulumi.StringPtrInput
	// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
	ResponselogTemplate pulumi.StringPtrInput
}

func (RequestLogProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestLogProfileArgs)(nil)).Elem()
}

type RequestLogProfileInput interface {
	pulumi.Input

	ToRequestLogProfileOutput() RequestLogProfileOutput
	ToRequestLogProfileOutputWithContext(ctx context.Context) RequestLogProfileOutput
}

func (*RequestLogProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestLogProfile)(nil)).Elem()
}

func (i *RequestLogProfile) ToRequestLogProfileOutput() RequestLogProfileOutput {
	return i.ToRequestLogProfileOutputWithContext(context.Background())
}

func (i *RequestLogProfile) ToRequestLogProfileOutputWithContext(ctx context.Context) RequestLogProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestLogProfileOutput)
}

// RequestLogProfileArrayInput is an input type that accepts RequestLogProfileArray and RequestLogProfileArrayOutput values.
// You can construct a concrete instance of `RequestLogProfileArrayInput` via:
//
//	RequestLogProfileArray{ RequestLogProfileArgs{...} }
type RequestLogProfileArrayInput interface {
	pulumi.Input

	ToRequestLogProfileArrayOutput() RequestLogProfileArrayOutput
	ToRequestLogProfileArrayOutputWithContext(context.Context) RequestLogProfileArrayOutput
}

type RequestLogProfileArray []RequestLogProfileInput

func (RequestLogProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequestLogProfile)(nil)).Elem()
}

func (i RequestLogProfileArray) ToRequestLogProfileArrayOutput() RequestLogProfileArrayOutput {
	return i.ToRequestLogProfileArrayOutputWithContext(context.Background())
}

func (i RequestLogProfileArray) ToRequestLogProfileArrayOutputWithContext(ctx context.Context) RequestLogProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestLogProfileArrayOutput)
}

// RequestLogProfileMapInput is an input type that accepts RequestLogProfileMap and RequestLogProfileMapOutput values.
// You can construct a concrete instance of `RequestLogProfileMapInput` via:
//
//	RequestLogProfileMap{ "key": RequestLogProfileArgs{...} }
type RequestLogProfileMapInput interface {
	pulumi.Input

	ToRequestLogProfileMapOutput() RequestLogProfileMapOutput
	ToRequestLogProfileMapOutputWithContext(context.Context) RequestLogProfileMapOutput
}

type RequestLogProfileMap map[string]RequestLogProfileInput

func (RequestLogProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequestLogProfile)(nil)).Elem()
}

func (i RequestLogProfileMap) ToRequestLogProfileMapOutput() RequestLogProfileMapOutput {
	return i.ToRequestLogProfileMapOutputWithContext(context.Background())
}

func (i RequestLogProfileMap) ToRequestLogProfileMapOutputWithContext(ctx context.Context) RequestLogProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestLogProfileMapOutput)
}

type RequestLogProfileOutput struct{ *pulumi.OutputState }

func (RequestLogProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestLogProfile)(nil)).Elem()
}

func (o RequestLogProfileOutput) ToRequestLogProfileOutput() RequestLogProfileOutput {
	return o
}

func (o RequestLogProfileOutput) ToRequestLogProfileOutputWithContext(ctx context.Context) RequestLogProfileOutput {
	return o
}

// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
func (o RequestLogProfileOutput) DefaultsFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.DefaultsFrom }).(pulumi.StringPtrOutput)
}

// Specifies user-defined description.
func (o RequestLogProfileOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
func (o RequestLogProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the pool associated with logging request errors. The default is None.
func (o RequestLogProfileOutput) ProxyResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ProxyResponse }).(pulumi.StringPtrOutput)
}

// Defines the pool associated with logging request errors. The default is None.
func (o RequestLogProfileOutput) ProxycloseOnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ProxycloseOnError }).(pulumi.StringPtrOutput)
}

// Defines the pool associated with logging request errors. The default is None.
func (o RequestLogProfileOutput) ProxyrespondOnLoggingerror() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ProxyrespondOnLoggingerror }).(pulumi.StringPtrOutput)
}

// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
func (o RequestLogProfileOutput) RequestLogging() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringOutput { return v.RequestLogging }).(pulumi.StringOutput)
}

// Defines the pool associated with logging request errors. The default is None.
func (o RequestLogProfileOutput) RequestlogErrorPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogErrorPool }).(pulumi.StringPtrOutput)
}

// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
func (o RequestLogProfileOutput) RequestlogErrorProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogErrorProtocol }).(pulumi.StringPtrOutput)
}

// Specifies the directives and entries to be logged for request errors.
func (o RequestLogProfileOutput) RequestlogErrorTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogErrorTemplate }).(pulumi.StringPtrOutput)
}

// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
func (o RequestLogProfileOutput) RequestlogPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogPool }).(pulumi.StringPtrOutput)
}

// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
func (o RequestLogProfileOutput) RequestlogProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogProtocol }).(pulumi.StringPtrOutput)
}

// Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
func (o RequestLogProfileOutput) RequestlogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.RequestlogTemplate }).(pulumi.StringPtrOutput)
}

// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
func (o RequestLogProfileOutput) ResponseLogging() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringOutput { return v.ResponseLogging }).(pulumi.StringOutput)
}

// Defines the pool associated with logging response errors. The default is `none`.
func (o RequestLogProfileOutput) ResponselogErrorPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogErrorPool }).(pulumi.StringPtrOutput)
}

// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
func (o RequestLogProfileOutput) ResponselogErrorProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogErrorProtocol }).(pulumi.StringPtrOutput)
}

// Specifies the directives and entries to be logged for request errors.
func (o RequestLogProfileOutput) ResponselogErrorTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogErrorTemplate }).(pulumi.StringPtrOutput)
}

// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
func (o RequestLogProfileOutput) ResponselogPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogPool }).(pulumi.StringPtrOutput)
}

// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
func (o RequestLogProfileOutput) ResponselogProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogProtocol }).(pulumi.StringPtrOutput)
}

// Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
func (o RequestLogProfileOutput) ResponselogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestLogProfile) pulumi.StringPtrOutput { return v.ResponselogTemplate }).(pulumi.StringPtrOutput)
}

type RequestLogProfileArrayOutput struct{ *pulumi.OutputState }

func (RequestLogProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequestLogProfile)(nil)).Elem()
}

func (o RequestLogProfileArrayOutput) ToRequestLogProfileArrayOutput() RequestLogProfileArrayOutput {
	return o
}

func (o RequestLogProfileArrayOutput) ToRequestLogProfileArrayOutputWithContext(ctx context.Context) RequestLogProfileArrayOutput {
	return o
}

func (o RequestLogProfileArrayOutput) Index(i pulumi.IntInput) RequestLogProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RequestLogProfile {
		return vs[0].([]*RequestLogProfile)[vs[1].(int)]
	}).(RequestLogProfileOutput)
}

type RequestLogProfileMapOutput struct{ *pulumi.OutputState }

func (RequestLogProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequestLogProfile)(nil)).Elem()
}

func (o RequestLogProfileMapOutput) ToRequestLogProfileMapOutput() RequestLogProfileMapOutput {
	return o
}

func (o RequestLogProfileMapOutput) ToRequestLogProfileMapOutputWithContext(ctx context.Context) RequestLogProfileMapOutput {
	return o
}

func (o RequestLogProfileMapOutput) MapIndex(k pulumi.StringInput) RequestLogProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RequestLogProfile {
		return vs[0].(map[string]*RequestLogProfile)[vs[1].(string)]
	}).(RequestLogProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequestLogProfileInput)(nil)).Elem(), &RequestLogProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequestLogProfileArrayInput)(nil)).Elem(), RequestLogProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequestLogProfileMapInput)(nil)).Elem(), RequestLogProfileMap{})
	pulumi.RegisterOutputType(RequestLogProfileOutput{})
	pulumi.RegisterOutputType(RequestLogProfileArrayOutput{})
	pulumi.RegisterOutputType(RequestLogProfileMapOutput{})
}
