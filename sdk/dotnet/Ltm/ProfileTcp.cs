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
    /// `f5bigip.ltm.ProfileTcp` Configures a custom TCP LTM Profile for use by health checks.
    /// 
    /// Resources should be named with their `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
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
    ///     var sanjose_tcp_lan_profile = new F5BigIP.Ltm.ProfileTcp("sanjose-tcp-lan-profile", new()
    ///     {
    ///         CloseWaitTimeout = 5,
    ///         DeferredAccept = "enabled",
    ///         FastOpen = "enabled",
    ///         Finwait2timeout = 5,
    ///         FinwaitTimeout = 300,
    ///         IdleTimeout = 200,
    ///         KeepaliveInterval = 1700,
    ///         Name = "/Common/sanjose-tcp-lan-profile",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Importing
    /// 
    /// An existing tcp profile can be imported into this resource by supplying tcp profile Name in `full path` as `id`.
    /// An example is below:
    /// ```sh
    /// $ terraform import bigip_ltm_profile_tcp.tcp-lan-profile-import /Common/test-tcp-lan-profile
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileTcp:ProfileTcp")]
    public partial class ProfileTcp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Output("closeWaitTimeout")]
        public Output<int> CloseWaitTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
        /// </summary>
        [Output("congestionControl")]
        public Output<string?> CongestionControl { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        /// </summary>
        [Output("deferredAccept")]
        public Output<string> DeferredAccept { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
        /// </summary>
        [Output("delayedAcks")]
        public Output<string?> DelayedAcks { get; private set; } = null!;

        /// <summary>
        /// Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
        /// </summary>
        [Output("earlyRetransmit")]
        public Output<string?> EarlyRetransmit { get; private set; } = null!;

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fast_open` set to `enabled`, argument `verified_accept` can't be set to `enabled`.
        /// </summary>
        [Output("fastOpen")]
        public Output<string> FastOpen { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        /// </summary>
        [Output("finwait2timeout")]
        public Output<int> Finwait2timeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        /// </summary>
        [Output("finwaitTimeout")]
        public Output<int> FinwaitTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
        /// </summary>
        [Output("initialCongestionWindowsize")]
        public Output<int?> InitialCongestionWindowsize { get; private set; } = null!;

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Output("keepaliveInterval")]
        public Output<int> KeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
        /// </summary>
        [Output("nagle")]
        public Output<string?> Nagle { get; private set; } = null!;

        /// <summary>
        /// Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// name of partition
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies the proxy buffer level, in bytes, at which the receive window is closed.
        /// </summary>
        [Output("proxybufferHigh")]
        public Output<int?> ProxybufferHigh { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
        /// </summary>
        [Output("receiveWindowsize")]
        public Output<int?> ReceiveWindowsize { get; private set; } = null!;

        /// <summary>
        /// Specifies the SEND window size. The default is 131072 bytes.
        /// </summary>
        [Output("sendBuffersize")]
        public Output<int?> SendBuffersize { get; private set; } = null!;

        /// <summary>
        /// Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
        /// </summary>
        [Output("taillossProbe")]
        public Output<string?> TaillossProbe { get; private set; } = null!;

        /// <summary>
        /// Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
        /// </summary>
        [Output("timewaitRecycle")]
        public Output<string?> TimewaitRecycle { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
        /// </summary>
        [Output("verifiedAccept")]
        public Output<string?> VerifiedAccept { get; private set; } = null!;

        /// <summary>
        /// Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
        /// </summary>
        [Output("zerowindowTimeout")]
        public Output<int?> ZerowindowTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileTcp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileTcp(string name, ProfileTcpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileTcp:ProfileTcp", name, args ?? new ProfileTcpArgs(), MakeResourceOptions(options, ""))
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

    public sealed class ProfileTcpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Input("closeWaitTimeout")]
        public Input<int>? CloseWaitTimeout { get; set; }

        /// <summary>
        /// Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
        /// </summary>
        [Input("congestionControl")]
        public Input<string>? CongestionControl { get; set; }

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
        /// Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
        /// </summary>
        [Input("delayedAcks")]
        public Input<string>? DelayedAcks { get; set; }

        /// <summary>
        /// Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
        /// </summary>
        [Input("earlyRetransmit")]
        public Input<string>? EarlyRetransmit { get; set; }

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fast_open` set to `enabled`, argument `verified_accept` can't be set to `enabled`.
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
        /// Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
        /// </summary>
        [Input("initialCongestionWindowsize")]
        public Input<int>? InitialCongestionWindowsize { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
        /// </summary>
        [Input("nagle")]
        public Input<string>? Nagle { get; set; }

        /// <summary>
        /// Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// name of partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the proxy buffer level, in bytes, at which the receive window is closed.
        /// </summary>
        [Input("proxybufferHigh")]
        public Input<int>? ProxybufferHigh { get; set; }

        /// <summary>
        /// Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
        /// </summary>
        [Input("receiveWindowsize")]
        public Input<int>? ReceiveWindowsize { get; set; }

        /// <summary>
        /// Specifies the SEND window size. The default is 131072 bytes.
        /// </summary>
        [Input("sendBuffersize")]
        public Input<int>? SendBuffersize { get; set; }

        /// <summary>
        /// Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
        /// </summary>
        [Input("taillossProbe")]
        public Input<string>? TaillossProbe { get; set; }

        /// <summary>
        /// Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
        /// </summary>
        [Input("timewaitRecycle")]
        public Input<string>? TimewaitRecycle { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
        /// </summary>
        [Input("verifiedAccept")]
        public Input<string>? VerifiedAccept { get; set; }

        /// <summary>
        /// Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
        /// </summary>
        [Input("zerowindowTimeout")]
        public Input<int>? ZerowindowTimeout { get; set; }

        public ProfileTcpArgs()
        {
        }
        public static new ProfileTcpArgs Empty => new ProfileTcpArgs();
    }

    public sealed class ProfileTcpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        /// </summary>
        [Input("closeWaitTimeout")]
        public Input<int>? CloseWaitTimeout { get; set; }

        /// <summary>
        /// Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
        /// </summary>
        [Input("congestionControl")]
        public Input<string>? CongestionControl { get; set; }

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
        /// Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
        /// </summary>
        [Input("delayedAcks")]
        public Input<string>? DelayedAcks { get; set; }

        /// <summary>
        /// Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
        /// </summary>
        [Input("earlyRetransmit")]
        public Input<string>? EarlyRetransmit { get; set; }

        /// <summary>
        /// When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fast_open` set to `enabled`, argument `verified_accept` can't be set to `enabled`.
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
        /// Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
        /// </summary>
        [Input("initialCongestionWindowsize")]
        public Input<int>? InitialCongestionWindowsize { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
        /// </summary>
        [Input("nagle")]
        public Input<string>? Nagle { get; set; }

        /// <summary>
        /// Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// name of partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the proxy buffer level, in bytes, at which the receive window is closed.
        /// </summary>
        [Input("proxybufferHigh")]
        public Input<int>? ProxybufferHigh { get; set; }

        /// <summary>
        /// Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
        /// </summary>
        [Input("receiveWindowsize")]
        public Input<int>? ReceiveWindowsize { get; set; }

        /// <summary>
        /// Specifies the SEND window size. The default is 131072 bytes.
        /// </summary>
        [Input("sendBuffersize")]
        public Input<int>? SendBuffersize { get; set; }

        /// <summary>
        /// Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
        /// </summary>
        [Input("taillossProbe")]
        public Input<string>? TaillossProbe { get; set; }

        /// <summary>
        /// Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
        /// </summary>
        [Input("timewaitRecycle")]
        public Input<string>? TimewaitRecycle { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
        /// </summary>
        [Input("verifiedAccept")]
        public Input<string>? VerifiedAccept { get; set; }

        /// <summary>
        /// Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
        /// </summary>
        [Input("zerowindowTimeout")]
        public Input<int>? ZerowindowTimeout { get; set; }

        public ProfileTcpState()
        {
        }
        public static new ProfileTcpState Empty => new ProfileTcpState();
    }
}
