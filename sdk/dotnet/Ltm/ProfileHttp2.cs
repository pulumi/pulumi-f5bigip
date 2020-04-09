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
    /// `f5bigip.ltm.ProfileHttp2` Configures a custom profile_http2 for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_profile_http2.html.markdown.
    /// </summary>
    public partial class ProfileHttp2 : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        /// </summary>
        [Output("activationModes")]
        public Output<ImmutableArray<string>> ActivationModes { get; private set; } = null!;

        /// <summary>
        /// Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        /// </summary>
        [Output("concurrentStreamsPerConnection")]
        public Output<int?> ConcurrentStreamsPerConnection { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        /// </summary>
        [Output("connectionIdleTimeout")]
        public Output<int?> ConnectionIdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Use the parent Http2 profile
        /// </summary>
        [Output("headerTableSize")]
        public Output<int?> HeaderTableSize { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_http2
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileHttp2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileHttp2(string name, ProfileHttp2Args args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttp2:ProfileHttp2", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProfileHttp2(string name, Input<string> id, ProfileHttp2State? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttp2:ProfileHttp2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileHttp2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileHttp2 Get(string name, Input<string> id, ProfileHttp2State? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileHttp2(name, id, state, options);
        }
    }

    public sealed class ProfileHttp2Args : Pulumi.ResourceArgs
    {
        [Input("activationModes")]
        private InputList<string>? _activationModes;

        /// <summary>
        /// Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        /// </summary>
        public InputList<string> ActivationModes
        {
            get => _activationModes ?? (_activationModes = new InputList<string>());
            set => _activationModes = value;
        }

        /// <summary>
        /// Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        /// </summary>
        [Input("concurrentStreamsPerConnection")]
        public Input<int>? ConcurrentStreamsPerConnection { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        /// </summary>
        [Input("connectionIdleTimeout")]
        public Input<int>? ConnectionIdleTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Use the parent Http2 profile
        /// </summary>
        [Input("headerTableSize")]
        public Input<int>? HeaderTableSize { get; set; }

        /// <summary>
        /// Name of the profile_http2
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ProfileHttp2Args()
        {
        }
    }

    public sealed class ProfileHttp2State : Pulumi.ResourceArgs
    {
        [Input("activationModes")]
        private InputList<string>? _activationModes;

        /// <summary>
        /// Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        /// </summary>
        public InputList<string> ActivationModes
        {
            get => _activationModes ?? (_activationModes = new InputList<string>());
            set => _activationModes = value;
        }

        /// <summary>
        /// Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        /// </summary>
        [Input("concurrentStreamsPerConnection")]
        public Input<int>? ConcurrentStreamsPerConnection { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        /// </summary>
        [Input("connectionIdleTimeout")]
        public Input<int>? ConnectionIdleTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Use the parent Http2 profile
        /// </summary>
        [Input("headerTableSize")]
        public Input<int>? HeaderTableSize { get; set; }

        /// <summary>
        /// Name of the profile_http2
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileHttp2State()
        {
        }
    }
}
