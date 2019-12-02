// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5bigip.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.ProfileTcp` Configures a custom profile_tcp for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_tcp.html.markdown.
    /// </summary>
    public partial class ProfileTcp : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Output("closeWaitTimeout")]
        public Output<int?> CloseWaitTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        /// </summary>
        [Output("deferredAccept")]
        public Output<string?> DeferredAccept { get; private set; } = null!;

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
        /// </summary>
        [Output("fastOpen")]
        public Output<string?> FastOpen { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        /// </summary>
        [Output("finwait2timeout")]
        public Output<int?> Finwait2timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        /// </summary>
        [Output("finwaitTimeout")]
        public Output<int?> FinwaitTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int?> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Output("keepaliveInterval")]
        public Output<int?> KeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_tcp
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileTcp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileTcp(string name, ProfileTcpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileTcp:ProfileTcp", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProfileTcp(string name, Input<string> id, ProfileTcpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileTcp:ProfileTcp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileTcp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileTcp Get(string name, Input<string> id, ProfileTcpState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileTcp(name, id, state, options);
        }
    }

    public sealed class ProfileTcpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Input("closeWaitTimeout")]
        public Input<int>? CloseWaitTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        /// </summary>
        [Input("deferredAccept")]
        public Input<string>? DeferredAccept { get; set; }

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
        /// </summary>
        [Input("fastOpen")]
        public Input<string>? FastOpen { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        /// </summary>
        [Input("finwait2timeout")]
        public Input<int>? Finwait2timeout { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        /// </summary>
        [Input("finwaitTimeout")]
        public Input<int>? FinwaitTimeout { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Name of the profile_tcp
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public ProfileTcpArgs()
        {
        }
    }

    public sealed class ProfileTcpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Input("closeWaitTimeout")]
        public Input<int>? CloseWaitTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        /// </summary>
        [Input("deferredAccept")]
        public Input<string>? DeferredAccept { get; set; }

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
        /// </summary>
        [Input("fastOpen")]
        public Input<string>? FastOpen { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        /// </summary>
        [Input("finwait2timeout")]
        public Input<int>? Finwait2timeout { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        /// </summary>
        [Input("finwaitTimeout")]
        public Input<int>? FinwaitTimeout { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Name of the profile_tcp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public ProfileTcpState()
        {
        }
    }
}
