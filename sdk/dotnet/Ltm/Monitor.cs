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
    /// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-monitor`.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var monitor = new F5BigIP.Ltm.Monitor("monitor", new()
    ///     {
    ///         Destination = "1.2.3.4:1234",
    ///         Interval = 998,
    ///         Name = "/Common/terraform_monitor",
    ///         Parent = "/Common/http",
    ///         Send = @"GET /some/path
    /// 
    /// ",
    ///         Timeout = 999,
    ///     });
    /// 
    ///     var test_https_monitor = new F5BigIP.Ltm.Monitor("test-https-monitor", new()
    ///     {
    ///         Interval = 999,
    ///         Name = "/Common/terraform_monitor",
    ///         Parent = "/Common/http",
    ///         Send = @"GET /some/path
    /// 
    /// ",
    ///         SslProfile = "/Common/serverssl",
    ///         Timeout = 1000,
    ///     });
    /// 
    ///     var test_ftp_monitor = new F5BigIP.Ltm.Monitor("test-ftp-monitor", new()
    ///     {
    ///         Destination = "*:8008",
    ///         Filename = "somefile",
    ///         Interval = 5,
    ///         Name = "/Common/ftp-test",
    ///         Parent = "/Common/ftp",
    ///         TimeUntilUp = 0,
    ///         Timeout = 16,
    ///     });
    /// 
    ///     var test_postgresql_monitor = new F5BigIP.Ltm.Monitor("test-postgresql-monitor", new()
    ///     {
    ///         Interval = 5,
    ///         Name = "/Common/test-postgresql-monitor",
    ///         Parent = "/Common/postgresql",
    ///         Password = "abcd1234",
    ///         Receive = "Test",
    ///         Send = "SELECT 'Test';",
    ///         Timeout = 16,
    ///         Username = "abcd",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;      
    /// 
    /// ## Importing
    /// 
    /// An existing monitor can be imported into this resource by supplying monitor Name in `full path` as `id`.
    /// An example is below:
    /// ```sh
    /// $ terraform import bigip_ltm_monitor.monitor /Common/terraform_monitor
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/monitor:Monitor")]
    public partial class Monitor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
        /// </summary>
        [Output("adaptive")]
        public Output<string> Adaptive { get; private set; } = null!;

        /// <summary>
        /// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
        /// </summary>
        [Output("adaptiveLimit")]
        public Output<int> AdaptiveLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Output("compatibility")]
        public Output<string?> Compatibility { get; private set; } = null!;

        /// <summary>
        /// Custom parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Output("customParent")]
        public Output<string?> CustomParent { get; private set; } = null!;

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

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
        /// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
        /// </summary>
        [Output("ipDscp")]
        public Output<int> IpDscp { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
        /// </summary>
        [Output("manualResume")]
        public Output<string> ManualResume { get; private set; } = null!;

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
        /// </summary>
        [Output("receive")]
        public Output<string?> Receive { get; private set; } = null!;

        /// <summary>
        /// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
        /// </summary>
        [Output("receiveDisable")]
        public Output<string?> ReceiveDisable { get; private set; } = null!;

        /// <summary>
        /// Instructs the system to mark the target resource down when the test is successful.
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// Specifies the text string that the monitor sends to the target object.
        /// </summary>
        [Output("send")]
        public Output<string> Send { get; private set; } = null!;

        /// <summary>
        /// Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
        /// </summary>
        [Output("sslProfile")]
        public Output<string?> SslProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
        /// </summary>
        [Output("timeUntilUp")]
        public Output<int> TimeUntilUp { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the monitor operates in transparent mode.
        /// </summary>
        [Output("transparent")]
        public Output<string> Transparent { get; private set; } = null!;

        /// <summary>
        /// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
        /// </summary>
        [Output("upInterval")]
        public Output<int> UpInterval { get; private set; } = null!;

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
                AdditionalSecretOutputs =
                {
                    "password",
                },
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

    public sealed class MonitorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
        /// </summary>
        [Input("adaptive")]
        public Input<string>? Adaptive { get; set; }

        /// <summary>
        /// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
        /// </summary>
        [Input("adaptiveLimit")]
        public Input<int>? AdaptiveLimit { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Input("compatibility")]
        public Input<string>? Compatibility { get; set; }

        /// <summary>
        /// Custom parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Input("customParent")]
        public Input<string>? CustomParent { get; set; }

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

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
        /// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
        /// </summary>
        [Input("ipDscp")]
        public Input<int>? IpDscp { get; set; }

        /// <summary>
        /// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
        /// </summary>
        [Input("manualResume")]
        public Input<string>? ManualResume { get; set; }

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
        /// </summary>
        [Input("receive")]
        public Input<string>? Receive { get; set; }

        /// <summary>
        /// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
        /// </summary>
        [Input("receiveDisable")]
        public Input<string>? ReceiveDisable { get; set; }

        /// <summary>
        /// Instructs the system to mark the target resource down when the test is successful.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// Specifies the text string that the monitor sends to the target object.
        /// </summary>
        [Input("send")]
        public Input<string>? Send { get; set; }

        /// <summary>
        /// Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
        /// </summary>
        [Input("sslProfile")]
        public Input<string>? SslProfile { get; set; }

        /// <summary>
        /// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
        /// </summary>
        [Input("timeUntilUp")]
        public Input<int>? TimeUntilUp { get; set; }

        /// <summary>
        /// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies whether the monitor operates in transparent mode.
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
        /// </summary>
        [Input("upInterval")]
        public Input<int>? UpInterval { get; set; }

        /// <summary>
        /// Specifies the user name if the monitored target requires authentication
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public MonitorArgs()
        {
        }
        public static new MonitorArgs Empty => new MonitorArgs();
    }

    public sealed class MonitorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
        /// </summary>
        [Input("adaptive")]
        public Input<string>? Adaptive { get; set; }

        /// <summary>
        /// Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
        /// </summary>
        [Input("adaptiveLimit")]
        public Input<int>? AdaptiveLimit { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        /// </summary>
        [Input("compatibility")]
        public Input<string>? Compatibility { get; set; }

        /// <summary>
        /// Custom parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Input("customParent")]
        public Input<string>? CustomParent { get; set; }

        /// <summary>
        /// Specifies the database in which the user is created
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

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
        /// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
        /// </summary>
        [Input("ipDscp")]
        public Input<int>? IpDscp { get; set; }

        /// <summary>
        /// Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
        /// </summary>
        [Input("manualResume")]
        public Input<string>? ManualResume { get; set; }

        /// <summary>
        /// Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parent monitor for the system to use for setting initial values for the new monitor.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Specifies the password if the monitored target requires authentication
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
        /// </summary>
        [Input("receive")]
        public Input<string>? Receive { get; set; }

        /// <summary>
        /// The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
        /// </summary>
        [Input("receiveDisable")]
        public Input<string>? ReceiveDisable { get; set; }

        /// <summary>
        /// Instructs the system to mark the target resource down when the test is successful.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// Specifies the text string that the monitor sends to the target object.
        /// </summary>
        [Input("send")]
        public Input<string>? Send { get; set; }

        /// <summary>
        /// Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
        /// </summary>
        [Input("sslProfile")]
        public Input<string>? SslProfile { get; set; }

        /// <summary>
        /// Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
        /// </summary>
        [Input("timeUntilUp")]
        public Input<int>? TimeUntilUp { get; set; }

        /// <summary>
        /// Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Specifies whether the monitor operates in transparent mode.
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
        /// </summary>
        [Input("upInterval")]
        public Input<int>? UpInterval { get; set; }

        /// <summary>
        /// Specifies the user name if the monitored target requires authentication
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public MonitorState()
        {
        }
        public static new MonitorState Empty => new MonitorState();
    }
}
