// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.Monitor` Configures a custom monitor for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var monitor = new F5BigIP.Ltm.Monitor("monitor", new F5BigIP.Ltm.MonitorArgs
    ///         {
    ///             Destination = "1.2.3.4:1234",
    ///             Interval = 999,
    ///             Name = "/Common/terraform_monitor",
    ///             Parent = "/Common/http",
    ///             Send = @"GET /some/path
    /// 
    /// ",
    ///             Timeout = 999,
    ///         });
    ///         var test_ftp_monitor = new F5BigIP.Ltm.Monitor("test-ftp-monitor", new F5BigIP.Ltm.MonitorArgs
    ///         {
    ///             Destination = "*:8008",
    ///             Filename = "somefile",
    ///             Interval = 5,
    ///             Name = "/Common/ftp-test",
    ///             Parent = "/Common/ftp",
    ///             TimeUntilUp = 0,
    ///             Timeout = 16,
    ///         });
    ///         var test_postgresql_monitor = new F5BigIP.Ltm.Monitor("test-postgresql-monitor", new F5BigIP.Ltm.MonitorArgs
    ///         {
    ///             Interval = 5,
    ///             Name = "/Common/test-postgresql-monitor",
    ///             Parent = "/Common/postgresql",
    ///             Password = "abcd1234",
    ///             Receive = "Test",
    ///             Send = "SELECT 'Test';",
    ///             Timeout = 16,
    ///             Username = "abcd",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Monitor : Pulumi.CustomResource
    {
        /// <summary>
        /// ftp adaptive
        /// </summary>
        [Output("adaptive")]
        public Output<string> Adaptive { get; private set; } = null!;

        /// <summary>
        /// Integer value
        /// </summary>
        [Output("adaptiveLimit")]
        public Output<int> AdaptiveLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Output("compatibility")]
        public Output<string?> Compatibility { get; private set; } = null!;

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specify an alias address for monitoring
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
        /// </summary>
        [Output("filename")]
        public Output<string?> Filename { get; private set; } = null!;

        /// <summary>
        /// Check interval in seconds
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        [Output("ipDscp")]
        public Output<int> IpDscp { get; private set; } = null!;

        [Output("manualResume")]
        public Output<string> ManualResume { get; private set; } = null!;

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Name of the monitor
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Existing LTM monitor to inherit from
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Expected response string
        /// </summary>
        [Output("receive")]
        public Output<string?> Receive { get; private set; } = null!;

        /// <summary>
        /// Expected response string.
        /// </summary>
        [Output("receiveDisable")]
        public Output<string?> ReceiveDisable { get; private set; } = null!;

        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// Request string to send
        /// </summary>
        [Output("send")]
        public Output<string> Send { get; private set; } = null!;

        /// <summary>
        /// Time in seconds
        /// </summary>
        [Output("timeUntilUp")]
        public Output<int> TimeUntilUp { get; private set; } = null!;

        /// <summary>
        /// Timeout in seconds
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        [Output("transparent")]
        public Output<string> Transparent { get; private set; } = null!;

        /// <summary>
        /// Specifies the user name if the monitored target requires authentication
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Monitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Monitor(string name, MonitorArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/monitor:Monitor", name, args ?? new MonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Monitor(string name, Input<string> id, MonitorState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/monitor:Monitor", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Monitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Monitor Get(string name, Input<string> id, MonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new Monitor(name, id, state, options);
        }
    }

    public sealed class MonitorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ftp adaptive
        /// </summary>
        [Input("adaptive")]
        public Input<string>? Adaptive { get; set; }

        /// <summary>
        /// Integer value
        /// </summary>
        [Input("adaptiveLimit")]
        public Input<int>? AdaptiveLimit { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Input("compatibility")]
        public Input<string>? Compatibility { get; set; }

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specify an alias address for monitoring
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
        /// </summary>
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// Check interval in seconds
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("ipDscp")]
        public Input<int>? IpDscp { get; set; }

        [Input("manualResume")]
        public Input<string>? ManualResume { get; set; }

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the monitor
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Existing LTM monitor to inherit from
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Expected response string
        /// </summary>
        [Input("receive")]
        public Input<string>? Receive { get; set; }

        /// <summary>
        /// Expected response string.
        /// </summary>
        [Input("receiveDisable")]
        public Input<string>? ReceiveDisable { get; set; }

        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// Request string to send
        /// </summary>
        [Input("send")]
        public Input<string>? Send { get; set; }

        /// <summary>
        /// Time in seconds
        /// </summary>
        [Input("timeUntilUp")]
        public Input<int>? TimeUntilUp { get; set; }

        /// <summary>
        /// Timeout in seconds
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the user name if the monitored target requires authentication
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public MonitorArgs()
        {
        }
    }

    public sealed class MonitorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ftp adaptive
        /// </summary>
        [Input("adaptive")]
        public Input<string>? Adaptive { get; set; }

        /// <summary>
        /// Integer value
        /// </summary>
        [Input("adaptiveLimit")]
        public Input<int>? AdaptiveLimit { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Input("compatibility")]
        public Input<string>? Compatibility { get; set; }

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specify an alias address for monitoring
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
        /// </summary>
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// Check interval in seconds
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("ipDscp")]
        public Input<int>? IpDscp { get; set; }

        [Input("manualResume")]
        public Input<string>? ManualResume { get; set; }

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the monitor
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Existing LTM monitor to inherit from
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Expected response string
        /// </summary>
        [Input("receive")]
        public Input<string>? Receive { get; set; }

        /// <summary>
        /// Expected response string.
        /// </summary>
        [Input("receiveDisable")]
        public Input<string>? ReceiveDisable { get; set; }

        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// Request string to send
        /// </summary>
        [Input("send")]
        public Input<string>? Send { get; set; }

        /// <summary>
        /// Time in seconds
        /// </summary>
        [Input("timeUntilUp")]
        public Input<int>? TimeUntilUp { get; set; }

        /// <summary>
        /// Timeout in seconds
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the user name if the monitored target requires authentication
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public MonitorState()
        {
        }
    }
}
