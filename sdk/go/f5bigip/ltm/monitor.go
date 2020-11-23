// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = ltm.NewMonitor(ctx, "test_ftp_monitor", &ltm.MonitorArgs{
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
// 		_, err = ltm.NewMonitor(ctx, "test_postgresql_monitor", &ltm.MonitorArgs{
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

	// ftp adaptive
	Adaptive pulumi.StringOutput `pulumi:"adaptive"`
	// Integer value
	AdaptiveLimit pulumi.IntOutput `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrOutput `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Specify an alias address for monitoring
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// Check interval in seconds
	Interval     pulumi.IntOutput    `pulumi:"interval"`
	IpDscp       pulumi.IntOutput    `pulumi:"ipDscp"`
	ManualResume pulumi.StringOutput `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Name of the monitor
	Name pulumi.StringOutput `pulumi:"name"`
	// Existing LTM monitor to inherit from
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Expected response string
	Receive pulumi.StringPtrOutput `pulumi:"receive"`
	// Expected response string.
	ReceiveDisable pulumi.StringPtrOutput `pulumi:"receiveDisable"`
	Reverse        pulumi.StringOutput    `pulumi:"reverse"`
	// Request string to send
	Send pulumi.StringOutput `pulumi:"send"`
	// Time in seconds
	TimeUntilUp pulumi.IntOutput `pulumi:"timeUntilUp"`
	// Timeout in seconds
	Timeout     pulumi.IntOutput    `pulumi:"timeout"`
	Transparent pulumi.StringOutput `pulumi:"transparent"`
	// Specifies the user name if the monitored target requires authentication
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil {
		args = &MonitorArgs{}
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
	// ftp adaptive
	Adaptive *string `pulumi:"adaptive"`
	// Integer value
	AdaptiveLimit *int `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility *string `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database *string `pulumi:"database"`
	// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specify an alias address for monitoring
	Destination *string `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename *string `pulumi:"filename"`
	// Check interval in seconds
	Interval     *int    `pulumi:"interval"`
	IpDscp       *int    `pulumi:"ipDscp"`
	ManualResume *string `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode *string `pulumi:"mode"`
	// Name of the monitor
	Name *string `pulumi:"name"`
	// Existing LTM monitor to inherit from
	Parent *string `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password *string `pulumi:"password"`
	// Expected response string
	Receive *string `pulumi:"receive"`
	// Expected response string.
	ReceiveDisable *string `pulumi:"receiveDisable"`
	Reverse        *string `pulumi:"reverse"`
	// Request string to send
	Send *string `pulumi:"send"`
	// Time in seconds
	TimeUntilUp *int `pulumi:"timeUntilUp"`
	// Timeout in seconds
	Timeout     *int    `pulumi:"timeout"`
	Transparent *string `pulumi:"transparent"`
	// Specifies the user name if the monitored target requires authentication
	Username *string `pulumi:"username"`
}

type MonitorState struct {
	// ftp adaptive
	Adaptive pulumi.StringPtrInput
	// Integer value
	AdaptiveLimit pulumi.IntPtrInput
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrInput
	// Specifies the database in which the user is created
	Database pulumi.StringPtrInput
	// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
	DefaultsFrom pulumi.StringPtrInput
	// Specify an alias address for monitoring
	Destination pulumi.StringPtrInput
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrInput
	// Check interval in seconds
	Interval     pulumi.IntPtrInput
	IpDscp       pulumi.IntPtrInput
	ManualResume pulumi.StringPtrInput
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringPtrInput
	// Name of the monitor
	Name pulumi.StringPtrInput
	// Existing LTM monitor to inherit from
	Parent pulumi.StringPtrInput
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrInput
	// Expected response string
	Receive pulumi.StringPtrInput
	// Expected response string.
	ReceiveDisable pulumi.StringPtrInput
	Reverse        pulumi.StringPtrInput
	// Request string to send
	Send pulumi.StringPtrInput
	// Time in seconds
	TimeUntilUp pulumi.IntPtrInput
	// Timeout in seconds
	Timeout     pulumi.IntPtrInput
	Transparent pulumi.StringPtrInput
	// Specifies the user name if the monitored target requires authentication
	Username pulumi.StringPtrInput
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// ftp adaptive
	Adaptive *string `pulumi:"adaptive"`
	// Integer value
	AdaptiveLimit *int `pulumi:"adaptiveLimit"`
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility *string `pulumi:"compatibility"`
	// Specifies the database in which the user is created
	Database *string `pulumi:"database"`
	// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specify an alias address for monitoring
	Destination *string `pulumi:"destination"`
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename *string `pulumi:"filename"`
	// Check interval in seconds
	Interval     *int    `pulumi:"interval"`
	IpDscp       *int    `pulumi:"ipDscp"`
	ManualResume *string `pulumi:"manualResume"`
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode *string `pulumi:"mode"`
	// Name of the monitor
	Name string `pulumi:"name"`
	// Existing LTM monitor to inherit from
	Parent string `pulumi:"parent"`
	// Specifies the password if the monitored target requires authentication
	Password *string `pulumi:"password"`
	// Expected response string
	Receive *string `pulumi:"receive"`
	// Expected response string.
	ReceiveDisable *string `pulumi:"receiveDisable"`
	Reverse        *string `pulumi:"reverse"`
	// Request string to send
	Send *string `pulumi:"send"`
	// Time in seconds
	TimeUntilUp *int `pulumi:"timeUntilUp"`
	// Timeout in seconds
	Timeout     *int    `pulumi:"timeout"`
	Transparent *string `pulumi:"transparent"`
	// Specifies the user name if the monitored target requires authentication
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// ftp adaptive
	Adaptive pulumi.StringPtrInput
	// Integer value
	AdaptiveLimit pulumi.IntPtrInput
	// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
	Compatibility pulumi.StringPtrInput
	// Specifies the database in which the user is created
	Database pulumi.StringPtrInput
	// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
	DefaultsFrom pulumi.StringPtrInput
	// Specify an alias address for monitoring
	Destination pulumi.StringPtrInput
	// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
	Filename pulumi.StringPtrInput
	// Check interval in seconds
	Interval     pulumi.IntPtrInput
	IpDscp       pulumi.IntPtrInput
	ManualResume pulumi.StringPtrInput
	// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
	Mode pulumi.StringPtrInput
	// Name of the monitor
	Name pulumi.StringInput
	// Existing LTM monitor to inherit from
	Parent pulumi.StringInput
	// Specifies the password if the monitored target requires authentication
	Password pulumi.StringPtrInput
	// Expected response string
	Receive pulumi.StringPtrInput
	// Expected response string.
	ReceiveDisable pulumi.StringPtrInput
	Reverse        pulumi.StringPtrInput
	// Request string to send
	Send pulumi.StringPtrInput
	// Time in seconds
	TimeUntilUp pulumi.IntPtrInput
	// Timeout in seconds
	Timeout     pulumi.IntPtrInput
	Transparent pulumi.StringPtrInput
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

func (Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((*Monitor)(nil)).Elem()
}

func (i Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

type MonitorOutput struct {
	*pulumi.OutputState
}

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorOutput)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
}
