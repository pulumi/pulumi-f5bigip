// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.NetTunnel` Manages a tunnel configuration
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
    ///     var example1 = new F5BigIP.NetTunnel("example1", new()
    ///     {
    ///         Name = "example1",
    ///         LocalAddress = "192.16.81.240",
    ///         Profile = "/Common/dslite",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/netTunnel:NetTunnel")]
    public partial class NetTunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Output("appService")]
        public Output<string?> AppService { get; private set; } = null!;

        /// <summary>
        /// Specifies whether auto lasthop is enabled or not
        /// </summary>
        [Output("autoLastHop")]
        public Output<string?> AutoLastHop { get; private set; } = null!;

        /// <summary>
        /// User defined description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies an idle timeout for wildcard tunnels in seconds
        /// </summary>
        [Output("idleTimeout")]
        public Output<int?> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// The key field may represent different values depending on the type of the tunnel
        /// </summary>
        [Output("key")]
        public Output<int?> Key { get; private set; } = null!;

        /// <summary>
        /// Specifies a local IP address. This option is required
        /// </summary>
        [Output("localAddress")]
        public Output<string> LocalAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies how the tunnel carries traffic
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum transmission unit (MTU) of the tunnel
        /// </summary>
        [Output("mtu")]
        public Output<int?> Mtu { get; private set; } = null!;

        /// <summary>
        /// Name of the tunnel
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the admin-partition within which this component resides
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to associate with the tunnel
        /// </summary>
        [Output("profile")]
        public Output<string> Profile { get; private set; } = null!;

        /// <summary>
        /// Specifies a remote IP address
        /// </summary>
        [Output("remoteAddress")]
        public Output<string?> RemoteAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies a secondary non-floating IP address when the local-address is set to a floating address
        /// </summary>
        [Output("secondaryAddress")]
        public Output<string?> SecondaryAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        /// </summary>
        [Output("tos")]
        public Output<string?> Tos { get; private set; } = null!;

        /// <summary>
        /// Specifies a traffic-group for use with the tunnel
        /// </summary>
        [Output("trafficGroup")]
        public Output<string?> TrafficGroup { get; private set; } = null!;

        /// <summary>
        /// Enables or disables the tunnel to be transparent
        /// </summary>
        [Output("transparent")]
        public Output<string?> Transparent { get; private set; } = null!;

        /// <summary>
        /// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        /// </summary>
        [Output("usePmtu")]
        public Output<string?> UsePmtu { get; private set; } = null!;


        /// <summary>
        /// Create a NetTunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetTunnel(string name, NetTunnelArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/netTunnel:NetTunnel", name, args ?? new NetTunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetTunnel(string name, Input<string> id, NetTunnelState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/netTunnel:NetTunnel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetTunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetTunnel Get(string name, Input<string> id, NetTunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new NetTunnel(name, id, state, options);
        }
    }

    public sealed class NetTunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies whether auto lasthop is enabled or not
        /// </summary>
        [Input("autoLastHop")]
        public Input<string>? AutoLastHop { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies an idle timeout for wildcard tunnels in seconds
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// The key field may represent different values depending on the type of the tunnel
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// Specifies a local IP address. This option is required
        /// </summary>
        [Input("localAddress", required: true)]
        public Input<string> LocalAddress { get; set; } = null!;

        /// <summary>
        /// Specifies how the tunnel carries traffic
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the maximum transmission unit (MTU) of the tunnel
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the tunnel
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the admin-partition within which this component resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the profile that you want to associate with the tunnel
        /// </summary>
        [Input("profile", required: true)]
        public Input<string> Profile { get; set; } = null!;

        /// <summary>
        /// Specifies a remote IP address
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// Specifies a secondary non-floating IP address when the local-address is set to a floating address
        /// </summary>
        [Input("secondaryAddress")]
        public Input<string>? SecondaryAddress { get; set; }

        /// <summary>
        /// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        /// </summary>
        [Input("tos")]
        public Input<string>? Tos { get; set; }

        /// <summary>
        /// Specifies a traffic-group for use with the tunnel
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        /// <summary>
        /// Enables or disables the tunnel to be transparent
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        /// </summary>
        [Input("usePmtu")]
        public Input<string>? UsePmtu { get; set; }

        public NetTunnelArgs()
        {
        }
        public static new NetTunnelArgs Empty => new NetTunnelArgs();
    }

    public sealed class NetTunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies whether auto lasthop is enabled or not
        /// </summary>
        [Input("autoLastHop")]
        public Input<string>? AutoLastHop { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies an idle timeout for wildcard tunnels in seconds
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// The key field may represent different values depending on the type of the tunnel
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// Specifies a local IP address. This option is required
        /// </summary>
        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// Specifies how the tunnel carries traffic
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the maximum transmission unit (MTU) of the tunnel
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the tunnel
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the admin-partition within which this component resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the profile that you want to associate with the tunnel
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// Specifies a remote IP address
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// Specifies a secondary non-floating IP address when the local-address is set to a floating address
        /// </summary>
        [Input("secondaryAddress")]
        public Input<string>? SecondaryAddress { get; set; }

        /// <summary>
        /// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        /// </summary>
        [Input("tos")]
        public Input<string>? Tos { get; set; }

        /// <summary>
        /// Specifies a traffic-group for use with the tunnel
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        /// <summary>
        /// Enables or disables the tunnel to be transparent
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        /// </summary>
        [Input("usePmtu")]
        public Input<string>? UsePmtu { get; set; }

        public NetTunnelState()
        {
        }
        public static new NetTunnelState Empty => new NetTunnelState();
    }
}
