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
    /// `f5bigip.ltm.RequestLogProfile` Resource used for Configures request logging using the Request Logging profile
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var request_log_profile_tc1_child = new F5BigIP.Ltm.RequestLogProfile("request-log-profile-tc1-child", new()
    ///     {
    ///         Name = "/Common/request-log-profile-tc1-child",
    ///         DefaultsFrom = request_log_profile_tc1.Name,
    ///         RequestLogging = "disabled",
    ///         RequestlogPool = "/Common/pool2",
    ///         RequestlogErrorPool = "/Common/pool1",
    ///         RequestlogProtocol = "mds-tcp",
    ///         RequestlogErrorProtocol = "mds-tcp",
    ///         ResponselogProtocol = "mds-tcp",
    ///         ResponselogErrorProtocol = "mds-tcp",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BIG-IP LTM Request Log profiles can be imported using the `name`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:ltm/requestLogProfile:RequestLogProfile test-request-log /Common/test-request-log
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/requestLogProfile:RequestLogProfile")]
    public partial class RequestLogProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies user-defined description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Output("proxyResponse")]
        public Output<string?> ProxyResponse { get; private set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Output("proxycloseOnError")]
        public Output<string?> ProxycloseOnError { get; private set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Output("proxyrespondOnLoggingerror")]
        public Output<string?> ProxyrespondOnLoggingerror { get; private set; } = null!;

        /// <summary>
        /// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Output("requestLogging")]
        public Output<string> RequestLogging { get; private set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Output("requestlogErrorPool")]
        public Output<string?> RequestlogErrorPool { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Output("requestlogErrorProtocol")]
        public Output<string?> RequestlogErrorProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Output("requestlogErrorTemplate")]
        public Output<string?> RequestlogErrorTemplate { get; private set; } = null!;

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
        /// </summary>
        [Output("requestlogPool")]
        public Output<string?> RequestlogPool { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Output("requestlogProtocol")]
        public Output<string?> RequestlogProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on requestlog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Output("requestlogTemplate")]
        public Output<string?> RequestlogTemplate { get; private set; } = null!;

        /// <summary>
        /// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Output("responseLogging")]
        public Output<string> ResponseLogging { get; private set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging response errors. The default is `none`.
        /// </summary>
        [Output("responselogErrorPool")]
        public Output<string?> ResponselogErrorPool { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Output("responselogErrorProtocol")]
        public Output<string?> ResponselogErrorProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Output("responselogErrorTemplate")]
        public Output<string?> ResponselogErrorTemplate { get; private set; } = null!;

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
        /// </summary>
        [Output("responselogPool")]
        public Output<string?> ResponselogPool { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Output("responselogProtocol")]
        public Output<string?> ResponselogProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on responselog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Output("responselogTemplate")]
        public Output<string?> ResponselogTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a RequestLogProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RequestLogProfile(string name, RequestLogProfileArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, args ?? new RequestLogProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RequestLogProfile(string name, Input<string> id, RequestLogProfileState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RequestLogProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RequestLogProfile Get(string name, Input<string> id, RequestLogProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new RequestLogProfile(name, id, state, options);
        }
    }

    public sealed class RequestLogProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies user-defined description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxyResponse")]
        public Input<string>? ProxyResponse { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxycloseOnError")]
        public Input<string>? ProxycloseOnError { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxyrespondOnLoggingerror")]
        public Input<string>? ProxyrespondOnLoggingerror { get; set; }

        /// <summary>
        /// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Input("requestLogging")]
        public Input<string>? RequestLogging { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("requestlogErrorPool")]
        public Input<string>? RequestlogErrorPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("requestlogErrorProtocol")]
        public Input<string>? RequestlogErrorProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Input("requestlogErrorTemplate")]
        public Input<string>? RequestlogErrorTemplate { get; set; }

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
        /// </summary>
        [Input("requestlogPool")]
        public Input<string>? RequestlogPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("requestlogProtocol")]
        public Input<string>? RequestlogProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on requestlog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Input("requestlogTemplate")]
        public Input<string>? RequestlogTemplate { get; set; }

        /// <summary>
        /// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Input("responseLogging")]
        public Input<string>? ResponseLogging { get; set; }

        /// <summary>
        /// Defines the pool associated with logging response errors. The default is `none`.
        /// </summary>
        [Input("responselogErrorPool")]
        public Input<string>? ResponselogErrorPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("responselogErrorProtocol")]
        public Input<string>? ResponselogErrorProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Input("responselogErrorTemplate")]
        public Input<string>? ResponselogErrorTemplate { get; set; }

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
        /// </summary>
        [Input("responselogPool")]
        public Input<string>? ResponselogPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("responselogProtocol")]
        public Input<string>? ResponselogProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on responselog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Input("responselogTemplate")]
        public Input<string>? ResponselogTemplate { get; set; }

        public RequestLogProfileArgs()
        {
        }
        public static new RequestLogProfileArgs Empty => new RequestLogProfileArgs();
    }

    public sealed class RequestLogProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies user-defined description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxyResponse")]
        public Input<string>? ProxyResponse { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxycloseOnError")]
        public Input<string>? ProxycloseOnError { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("proxyrespondOnLoggingerror")]
        public Input<string>? ProxyrespondOnLoggingerror { get; set; }

        /// <summary>
        /// Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Input("requestLogging")]
        public Input<string>? RequestLogging { get; set; }

        /// <summary>
        /// Defines the pool associated with logging request errors. The default is None.
        /// </summary>
        [Input("requestlogErrorPool")]
        public Input<string>? RequestlogErrorPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("requestlogErrorProtocol")]
        public Input<string>? RequestlogErrorProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Input("requestlogErrorTemplate")]
        public Input<string>? RequestlogErrorTemplate { get; set; }

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
        /// </summary>
        [Input("requestlogPool")]
        public Input<string>? RequestlogPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("requestlogProtocol")]
        public Input<string>? RequestlogProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on requestlog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Input("requestlogTemplate")]
        public Input<string>? RequestlogTemplate { get; set; }

        /// <summary>
        /// Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
        /// </summary>
        [Input("responseLogging")]
        public Input<string>? ResponseLogging { get; set; }

        /// <summary>
        /// Defines the pool associated with logging response errors. The default is `none`.
        /// </summary>
        [Input("responselogErrorPool")]
        public Input<string>? ResponselogErrorPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("responselogErrorProtocol")]
        public Input<string>? ResponselogErrorProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged for request errors.
        /// </summary>
        [Input("responselogErrorTemplate")]
        public Input<string>? ResponselogErrorTemplate { get; set; }

        /// <summary>
        /// Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
        /// </summary>
        [Input("responselogPool")]
        public Input<string>? ResponselogPool { get; set; }

        /// <summary>
        /// Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
        /// </summary>
        [Input("responselogProtocol")]
        public Input<string>? ResponselogProtocol { get; set; }

        /// <summary>
        /// Specifies the directives and entries to be logged. More infor on responselog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
        /// </summary>
        [Input("responselogTemplate")]
        public Input<string>? ResponselogTemplate { get; set; }

        public RequestLogProfileState()
        {
        }
        public static new RequestLogProfileState Empty => new RequestLogProfileState();
    }
}
