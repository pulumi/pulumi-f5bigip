// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.Monitor` Configures a custom monitor for use by health checks.
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewMonitor(ctx, "monitor", &ltm.MonitorArgs{
// 			Destination: pulumi.String("1.2.3.4:1234"),
// 			Interval:    pulumi.Int(999),
// 			Name:        pulumi.String("/Common/terraform_monitor"),
// 			Parent:      pulumi.String("/Common/http"),
// 			Send: pulumi.String(fmt.Sprintf("%v%v", "GET /some/path\n", "\n")),
// 			Timeout: pulumi.Int(999),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewMonitor(ctx, "test-ftp-monitor", &ltm.MonitorArgs{
// 			Destination: pulumi.String("*:8008"),
// 			Filename:    pulumi.String("somefile"),
// 			Interval:    pulumi.Int(5),
// 			Name:        pulumi.String("/Common/ftp-test"),
// 			Parent:      pulumi.String("/Common/ftp"),
// 			TimeUntilUp: pulumi.Int(0),
// 			Timeout:     pulumi.Int(16),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewMonitor(ctx, "test-postgresql-monitor", &ltm.MonitorArgs{
// 			Interval: pulumi.Int(5),
// 			Name:     pulumi.String("/Common/test-postgresql-monitor"),
// 			Parent:   pulumi.String("/Common/postgresql"),
// 			Password: pulumi.String("abcd1234"),
// 			Receive:  pulumi.String("Test"),
// 			Send:     pulumi.String("SELECT 'Test';"),
// 			Timeout:  pulumi.Int(16),
// 			Username: pulumi.String("abcd"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Monitor struct {
	pulumi.CustomResourceState

	// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
	Adaptive pulumi.StringOutput `pulumi:"adaptive"`
	// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
	AdaptiveLimit pulumi.IntOutput `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrOutput `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// Specify an alias address for monitoring
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default is `5`
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
	IpDscp pulumi.IntOutput `pulumi:"ipDscp"`
	// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
	ManualResume pulumi.StringOutput `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parent monitor for the system to use for setting initial values for the new monitor.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
	Receive pulumi.StringPtrOutput `pulumi:"receive"`
	// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
	ReceiveDisable pulumi.StringPtrOutput `pulumi:"receiveDisable"`
	// Instructs the system to mark the target resource down when the test is successful.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// Specifies the text string that the monitor sends to the target object.
	Send pulumi.StringOutput `pulumi:"send"`
	// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
	TimeUntilUp pulumi.IntOutput `pulumi:"timeUntilUp"`
	// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// Specifies whether the monitor operates in transparent mode.
	Transparent pulumi.StringOutput `pulumi:"transparent"`
	// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
	UpInterval pulumi.IntOutput `pulumi:"upInterval"`
	// Specifies the user name if the monitored target requires authentication
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource Monitor
	err := ctx.RegisterResource("f5bigip:ltm/monitor:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("f5bigip:ltm/monitor:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
	// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
	Adaptive *string `pulumi:"adaptive"`
	// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
	AdaptiveLimit *int `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility *string `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database *string `pulumi:"database"`
	// Specify an alias address for monitoring
	Destination *string `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename *string `pulumi:"filename"`
	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default is `5`
	Interval *int `pulumi:"interval"`
	// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
	IpDscp *int `pulumi:"ipDscp"`
	// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
	ManualResume *string `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode *string `pulumi:"mode"`
	// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
	Name *string `pulumi:"name"`
	// Parent monitor for the system to use for setting initial values for the new monitor.
	Parent *string `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password *string `pulumi:"password"`
	// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
	Receive *string `pulumi:"receive"`
	// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
	ReceiveDisable *string `pulumi:"receiveDisable"`
	// Instructs the system to mark the target resource down when the test is successful.
	Reverse *string `pulumi:"reverse"`
	// Specifies the text string that the monitor sends to the target object.
	Send *string `pulumi:"send"`
	// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
	TimeUntilUp *int `pulumi:"timeUntilUp"`
	// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
	Timeout *int `pulumi:"timeout"`
	// Specifies whether the monitor operates in transparent mode.
	Transparent *string `pulumi:"transparent"`
	// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
	UpInterval *int `pulumi:"upInterval"`
	// Specifies the user name if the monitored target requires authentication
	Username *string `pulumi:"username"`
}

type MonitorState struct {
	// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
	Adaptive pulumi.StringPtrInput
	// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
	AdaptiveLimit pulumi.IntPtrInput
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrInput
	// Specifies the database in which the user is created
	Database pulumi.StringPtrInput
	// Specify an alias address for monitoring
	Destination pulumi.StringPtrInput
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrInput
	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default is `5`
	Interval pulumi.IntPtrInput
	// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
	IpDscp pulumi.IntPtrInput
	// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
	ManualResume pulumi.StringPtrInput
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringPtrInput
	// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
	Name pulumi.StringPtrInput
	// Parent monitor for the system to use for setting initial values for the new monitor.
	Parent pulumi.StringPtrInput
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrInput
	// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
	Receive pulumi.StringPtrInput
	// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
	ReceiveDisable pulumi.StringPtrInput
	// Instructs the system to mark the target resource down when the test is successful.
	Reverse pulumi.StringPtrInput
	// Specifies the text string that the monitor sends to the target object.
	Send pulumi.StringPtrInput
	// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
	TimeUntilUp pulumi.IntPtrInput
	// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
	Timeout pulumi.IntPtrInput
	// Specifies whether the monitor operates in transparent mode.
	Transparent pulumi.StringPtrInput
	// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
	UpInterval pulumi.IntPtrInput
	// Specifies the user name if the monitored target requires authentication
	Username pulumi.StringPtrInput
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
	Adaptive *string `pulumi:"adaptive"`
	// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
	AdaptiveLimit *int `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility *string `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database *string `pulumi:"database"`
	// Specify an alias address for monitoring
	Destination *string `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename *string `pulumi:"filename"`
	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default is `5`
	Interval *int `pulumi:"interval"`
	// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
	IpDscp *int `pulumi:"ipDscp"`
	// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
	ManualResume *string `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode *string `pulumi:"mode"`
	// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
	Name string `pulumi:"name"`
	// Parent monitor for the system to use for setting initial values for the new monitor.
	Parent string `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password *string `pulumi:"password"`
	// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
	Receive *string `pulumi:"receive"`
	// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
	ReceiveDisable *string `pulumi:"receiveDisable"`
	// Instructs the system to mark the target resource down when the test is successful.
	Reverse *string `pulumi:"reverse"`
	// Specifies the text string that the monitor sends to the target object.
	Send *string `pulumi:"send"`
	// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
	TimeUntilUp *int `pulumi:"timeUntilUp"`
	// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
	Timeout *int `pulumi:"timeout"`
	// Specifies whether the monitor operates in transparent mode.
	Transparent *string `pulumi:"transparent"`
	// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
	UpInterval *int `pulumi:"upInterval"`
	// Specifies the user name if the monitored target requires authentication
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
	Adaptive pulumi.StringPtrInput
	// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
	AdaptiveLimit pulumi.IntPtrInput
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrInput
	// Specifies the database in which the user is created
	Database pulumi.StringPtrInput
	// Specify an alias address for monitoring
	Destination pulumi.StringPtrInput
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrInput
	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default is `5`
	Interval pulumi.IntPtrInput
	// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
	IpDscp pulumi.IntPtrInput
	// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
	ManualResume pulumi.StringPtrInput
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringPtrInput
	// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
	Name pulumi.StringInput
	// Parent monitor for the system to use for setting initial values for the new monitor.
	Parent pulumi.StringInput
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrInput
	// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
	Receive pulumi.StringPtrInput
	// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
	ReceiveDisable pulumi.StringPtrInput
	// Instructs the system to mark the target resource down when the test is successful.
	Reverse pulumi.StringPtrInput
	// Specifies the text string that the monitor sends to the target object.
	Send pulumi.StringPtrInput
	// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
	TimeUntilUp pulumi.IntPtrInput
	// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
	Timeout pulumi.IntPtrInput
	// Specifies whether the monitor operates in transparent mode.
	Transparent pulumi.StringPtrInput
	// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
	UpInterval pulumi.IntPtrInput
	// Specifies the user name if the monitored target requires authentication
	Username pulumi.StringPtrInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

// MonitorArrayInput is an input type that accepts MonitorArray and MonitorArrayOutput values.
// You can construct a concrete instance of `MonitorArrayInput` via:
//
//          MonitorArray{ MonitorArgs{...} }
type MonitorArrayInput interface {
	pulumi.Input

	ToMonitorArrayOutput() MonitorArrayOutput
	ToMonitorArrayOutputWithContext(context.Context) MonitorArrayOutput
}

type MonitorArray []MonitorInput

func (MonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (i MonitorArray) ToMonitorArrayOutput() MonitorArrayOutput {
	return i.ToMonitorArrayOutputWithContext(context.Background())
}

func (i MonitorArray) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorArrayOutput)
}

// MonitorMapInput is an input type that accepts MonitorMap and MonitorMapOutput values.
// You can construct a concrete instance of `MonitorMapInput` via:
//
//          MonitorMap{ "key": MonitorArgs{...} }
type MonitorMapInput interface {
	pulumi.Input

	ToMonitorMapOutput() MonitorMapOutput
	ToMonitorMapOutputWithContext(context.Context) MonitorMapOutput
}

type MonitorMap map[string]MonitorInput

func (MonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (i MonitorMap) ToMonitorMapOutput() MonitorMapOutput {
	return i.ToMonitorMapOutputWithContext(context.Background())
}

func (i MonitorMap) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorMapOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

type MonitorArrayOutput struct{ *pulumi.OutputState }

func (MonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (o MonitorArrayOutput) ToMonitorArrayOutput() MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) Index(i pulumi.IntInput) MonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].([]*Monitor)[vs[1].(int)]
	}).(MonitorOutput)
}

type MonitorMapOutput struct{ *pulumi.OutputState }

func (MonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (o MonitorMapOutput) ToMonitorMapOutput() MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) MapIndex(k pulumi.StringInput) MonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].(map[string]*Monitor)[vs[1].(string)]
	}).(MonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorInput)(nil)).Elem(), &Monitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorArrayInput)(nil)).Elem(), MonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorMapInput)(nil)).Elem(), MonitorMap{})
	pulumi.RegisterOutputType(MonitorOutput{})
	pulumi.RegisterOutputType(MonitorArrayOutput{})
	pulumi.RegisterOutputType(MonitorMapOutput{})
}
