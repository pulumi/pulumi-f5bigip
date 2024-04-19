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
    /// `f5bigip.ltm.ProfileFastL4` Configures a custom LTM fastL4 profile for use by health checks.
    /// 
    /// Resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
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
    ///     var profileFastl4 = new F5BigIP.Ltm.ProfileFastL4("profile_fastl4", new()
    ///     {
    ///         Name = "/Common/sjfastl4profile",
    ///         DefaultsFrom = "/Common/fastL4",
    ///         ClientTimeout = 40,
    ///         ExplicitflowMigration = "enabled",
    ///         HardwareSyncookie = "enabled",
    ///         IdleTimeout = "200",
    ///         IptosToclient = "pass-through",
    ///         IptosToserver = "pass-through",
    ///         KeepaliveInterval = "disabled",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// BIG-IP LTM fastl4 profiles can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:ltm/profileFastL4:ProfileFastL4 test-fastl4 /Common/test-fastl4
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileFastL4:ProfileFastL4")]
    public partial class ProfileFastL4 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Output("clientTimeout")]
        public Output<int> ClientTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Output("explicitflowMigration")]
        public Output<string> ExplicitflowMigration { get; private set; } = null!;

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Output("hardwareSyncookie")]
        public Output<string> HardwareSyncookie { get; private set; } = null!;

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<string> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Output("iptosToclient")]
        public Output<string> IptosToclient { get; private set; } = null!;

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Output("iptosToserver")]
        public Output<string> IptosToserver { get; private set; } = null!;

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Output("keepaliveInterval")]
        public Output<string> KeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
        /// </summary>
        [Output("lateBinding")]
        public Output<string> LateBinding { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
        /// </summary>
        [Output("looseClose")]
        public Output<string> LooseClose { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
        /// </summary>
        [Output("looseInitiation")]
        public Output<string> LooseInitiation { get; private set; } = null!;

        /// <summary>
        /// Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// name of partition
        /// </summary>
        [Output("partition")]
        public Output<string> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
        /// </summary>
        [Output("receiveWindowsize")]
        public Output<int> ReceiveWindowsize { get; private set; } = null!;

        /// <summary>
        /// Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
        /// </summary>
        [Output("tcpHandshakeTimeout")]
        public Output<string> TcpHandshakeTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileFastL4 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileFastL4(string name, ProfileFastL4Args args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastL4:ProfileFastL4", name, args ?? new ProfileFastL4Args(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileFastL4(string name, Input<string> id, ProfileFastL4State? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastL4:ProfileFastL4", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileFastL4 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileFastL4 Get(string name, Input<string> id, ProfileFastL4State? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileFastL4(name, id, state, options);
        }
    }

    public sealed class ProfileFastL4Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Input("clientTimeout")]
        public Input<int>? ClientTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Input("explicitflowMigration")]
        public Input<string>? ExplicitflowMigration { get; set; }

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Input("hardwareSyncookie")]
        public Input<string>? HardwareSyncookie { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToclient")]
        public Input<string>? IptosToclient { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToserver")]
        public Input<string>? IptosToserver { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<string>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
        /// </summary>
        [Input("lateBinding")]
        public Input<string>? LateBinding { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
        /// </summary>
        [Input("looseClose")]
        public Input<string>? LooseClose { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
        /// </summary>
        [Input("looseInitiation")]
        public Input<string>? LooseInitiation { get; set; }

        /// <summary>
        /// Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// name of partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
        /// </summary>
        [Input("receiveWindowsize")]
        public Input<int>? ReceiveWindowsize { get; set; }

        /// <summary>
        /// Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
        /// </summary>
        [Input("tcpHandshakeTimeout")]
        public Input<string>? TcpHandshakeTimeout { get; set; }

        public ProfileFastL4Args()
        {
        }
        public static new ProfileFastL4Args Empty => new ProfileFastL4Args();
    }

    public sealed class ProfileFastL4State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Input("clientTimeout")]
        public Input<int>? ClientTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Input("explicitflowMigration")]
        public Input<string>? ExplicitflowMigration { get; set; }

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Input("hardwareSyncookie")]
        public Input<string>? HardwareSyncookie { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToclient")]
        public Input<string>? IptosToclient { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToserver")]
        public Input<string>? IptosToserver { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<string>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
        /// </summary>
        [Input("lateBinding")]
        public Input<string>? LateBinding { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
        /// </summary>
        [Input("looseClose")]
        public Input<string>? LooseClose { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
        /// </summary>
        [Input("looseInitiation")]
        public Input<string>? LooseInitiation { get; set; }

        /// <summary>
        /// Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// name of partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
        /// </summary>
        [Input("receiveWindowsize")]
        public Input<int>? ReceiveWindowsize { get; set; }

        /// <summary>
        /// Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
        /// </summary>
        [Input("tcpHandshakeTimeout")]
        public Input<string>? TcpHandshakeTimeout { get; set; }

        public ProfileFastL4State()
        {
        }
        public static new ProfileFastL4State Empty => new ProfileFastL4State();
    }
}
