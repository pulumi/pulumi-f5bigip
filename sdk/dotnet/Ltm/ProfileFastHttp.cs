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
    /// `f5bigip.ltm.ProfileFastHttp` Configures a custom profile_fasthttp for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
    ///     var sjfasthttpprofile = new F5BigIP.Ltm.ProfileFastHttp("sjfasthttpprofile", new()
    ///     {
    ///         Name = "/Common/sjfasthttpprofile",
    ///         DefaultsFrom = "/Common/fasthttp",
    ///         IdleTimeout = 300,
    ///         ConnpoolidleTimeoutoverride = 0,
    ///         ConnpoolMaxreuse = 2,
    ///         ConnpoolMaxsize = 2048,
    ///         ConnpoolMinsize = 0,
    ///         ConnpoolReplenish = "enabled",
    ///         ConnpoolStep = 4,
    ///         Forcehttp10response = "disabled",
    ///         MaxheaderSize = 32768,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileFastHttp:ProfileFastHttp")]
    public partial class ProfileFastHttp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        /// </summary>
        [Output("connpoolMaxreuse")]
        public Output<int> ConnpoolMaxreuse { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        /// </summary>
        [Output("connpoolMaxsize")]
        public Output<int> ConnpoolMaxsize { get; private set; } = null!;

        /// <summary>
        /// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        /// </summary>
        [Output("connpoolMinsize")]
        public Output<int> ConnpoolMinsize { get; private set; } = null!;

        /// <summary>
        /// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        /// </summary>
        [Output("connpoolReplenish")]
        public Output<string> ConnpoolReplenish { get; private set; } = null!;

        /// <summary>
        /// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        /// </summary>
        [Output("connpoolStep")]
        public Output<int> ConnpoolStep { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        /// </summary>
        [Output("connpoolidleTimeoutoverride")]
        public Output<int> ConnpoolidleTimeoutoverride { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        /// </summary>
        [Output("forcehttp10response")]
        public Output<string> Forcehttp10response { get; private set; } = null!;

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        /// </summary>
        [Output("maxheaderSize")]
        public Output<int> MaxheaderSize { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_fasthttp
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileFastHttp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileFastHttp(string name, ProfileFastHttpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, args ?? new ProfileFastHttpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileFastHttp(string name, Input<string> id, ProfileFastHttpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileFastHttp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileFastHttp Get(string name, Input<string> id, ProfileFastHttpState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileFastHttp(name, id, state, options);
        }
    }

    public sealed class ProfileFastHttpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        /// </summary>
        [Input("connpoolMaxreuse")]
        public Input<int>? ConnpoolMaxreuse { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        /// </summary>
        [Input("connpoolMaxsize")]
        public Input<int>? ConnpoolMaxsize { get; set; }

        /// <summary>
        /// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        /// </summary>
        [Input("connpoolMinsize")]
        public Input<int>? ConnpoolMinsize { get; set; }

        /// <summary>
        /// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        /// </summary>
        [Input("connpoolReplenish")]
        public Input<string>? ConnpoolReplenish { get; set; }

        /// <summary>
        /// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        /// </summary>
        [Input("connpoolStep")]
        public Input<int>? ConnpoolStep { get; set; }

        /// <summary>
        /// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        /// </summary>
        [Input("connpoolidleTimeoutoverride")]
        public Input<int>? ConnpoolidleTimeoutoverride { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        /// </summary>
        [Input("forcehttp10response")]
        public Input<string>? Forcehttp10response { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        /// </summary>
        [Input("maxheaderSize")]
        public Input<int>? MaxheaderSize { get; set; }

        /// <summary>
        /// Name of the profile_fasthttp
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ProfileFastHttpArgs()
        {
        }
        public static new ProfileFastHttpArgs Empty => new ProfileFastHttpArgs();
    }

    public sealed class ProfileFastHttpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        /// </summary>
        [Input("connpoolMaxreuse")]
        public Input<int>? ConnpoolMaxreuse { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        /// </summary>
        [Input("connpoolMaxsize")]
        public Input<int>? ConnpoolMaxsize { get; set; }

        /// <summary>
        /// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        /// </summary>
        [Input("connpoolMinsize")]
        public Input<int>? ConnpoolMinsize { get; set; }

        /// <summary>
        /// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        /// </summary>
        [Input("connpoolReplenish")]
        public Input<string>? ConnpoolReplenish { get; set; }

        /// <summary>
        /// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        /// </summary>
        [Input("connpoolStep")]
        public Input<int>? ConnpoolStep { get; set; }

        /// <summary>
        /// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        /// </summary>
        [Input("connpoolidleTimeoutoverride")]
        public Input<int>? ConnpoolidleTimeoutoverride { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        /// </summary>
        [Input("forcehttp10response")]
        public Input<string>? Forcehttp10response { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        /// </summary>
        [Input("maxheaderSize")]
        public Input<int>? MaxheaderSize { get; set; }

        /// <summary>
        /// Name of the profile_fasthttp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileFastHttpState()
        {
        }
        public static new ProfileFastHttpState Empty => new ProfileFastHttpState();
    }
}
