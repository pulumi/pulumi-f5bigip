// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.FastTcpApp` This resource will create and manage FAST TCP applications on BIG-IP from provided JSON declaration.
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
    ///     var fast_tcp_app = new F5BigIP.FastTcpApp("fast-tcp-app", new()
    ///     {
    ///         Application = "tcp_app_2",
    ///         PoolMembers = new[]
    ///         {
    ///             new F5BigIP.Inputs.FastTcpAppPoolMemberArgs
    ///             {
    ///                 Addresses = new[]
    ///                 {
    ///                     "10.11.34.65",
    ///                     "56.43.23.76",
    ///                 },
    ///                 ConnectionLimit = 4,
    ///                 Port = 443,
    ///                 PriorityGroup = 1,
    ///                 ShareNodes = true,
    ///             },
    ///         },
    ///         Tenant = "tcp_app_tenant",
    ///         VirtualServer = new F5BigIP.Inputs.FastTcpAppVirtualServerArgs
    ///         {
    ///             Ip = "11.12.16.30",
    ///             Port = 443,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/fastTcpApp:FastTcpApp")]
    public partial class FastTcpApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the FAST TCP application.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Output("existingMonitor")]
        public Output<string?> ExistingMonitor { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP pool.
        /// </summary>
        [Output("existingPool")]
        public Output<string?> ExistingPool { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Output("existingSnatPool")]
        public Output<string?> ExistingSnatPool { get; private set; } = null!;

        /// <summary>
        /// Json payload for FAST TCP application.
        /// </summary>
        [Output("fastTcpJson")]
        public Output<string> FastTcpJson { get; private set; } = null!;

        /// <summary>
        /// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
        /// </summary>
        [Output("loadBalancingMode")]
        public Output<string?> LoadBalancingMode { get; private set; } = null!;

        /// <summary>
        /// `monitor` block takes input for FAST-Generated Pool Monitor.
        /// See Pool Monitor below for more details.
        /// </summary>
        [Output("monitor")]
        public Output<Outputs.FastTcpAppMonitor?> Monitor { get; private set; } = null!;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        [Output("poolMembers")]
        public Output<ImmutableArray<Outputs.FastTcpAppPoolMember>> PoolMembers { get; private set; } = null!;

        /// <summary>
        /// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
        /// </summary>
        [Output("slowRampTime")]
        public Output<int?> SlowRampTime { get; private set; } = null!;

        /// <summary>
        /// List of address to be used for FAST-Generated SNAT Pool.
        /// </summary>
        [Output("snatPoolAddresses")]
        public Output<ImmutableArray<string>> SnatPoolAddresses { get; private set; } = null!;

        /// <summary>
        /// Name of the FAST TCP application tenant.
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Output("virtualServer")]
        public Output<Outputs.FastTcpAppVirtualServer?> VirtualServer { get; private set; } = null!;


        /// <summary>
        /// Create a FastTcpApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FastTcpApp(string name, FastTcpAppArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastTcpApp:FastTcpApp", name, args ?? new FastTcpAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FastTcpApp(string name, Input<string> id, FastTcpAppState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastTcpApp:FastTcpApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FastTcpApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FastTcpApp Get(string name, Input<string> id, FastTcpAppState? state = null, CustomResourceOptions? options = null)
        {
            return new FastTcpApp(name, id, state, options);
        }
    }

    public sealed class FastTcpAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST TCP application.
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Input("existingMonitor")]
        public Input<string>? ExistingMonitor { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP pool.
        /// </summary>
        [Input("existingPool")]
        public Input<string>? ExistingPool { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Input("existingSnatPool")]
        public Input<string>? ExistingSnatPool { get; set; }

        /// <summary>
        /// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
        /// </summary>
        [Input("loadBalancingMode")]
        public Input<string>? LoadBalancingMode { get; set; }

        /// <summary>
        /// `monitor` block takes input for FAST-Generated Pool Monitor.
        /// See Pool Monitor below for more details.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.FastTcpAppMonitorArgs>? Monitor { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastTcpAppPoolMemberArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastTcpAppPoolMemberArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastTcpAppPoolMemberArgs>());
            set => _poolMembers = value;
        }

        /// <summary>
        /// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
        /// </summary>
        [Input("slowRampTime")]
        public Input<int>? SlowRampTime { get; set; }

        [Input("snatPoolAddresses")]
        private InputList<string>? _snatPoolAddresses;

        /// <summary>
        /// List of address to be used for FAST-Generated SNAT Pool.
        /// </summary>
        public InputList<string> SnatPoolAddresses
        {
            get => _snatPoolAddresses ?? (_snatPoolAddresses = new InputList<string>());
            set => _snatPoolAddresses = value;
        }

        /// <summary>
        /// Name of the FAST TCP application tenant.
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastTcpAppVirtualServerArgs>? VirtualServer { get; set; }

        public FastTcpAppArgs()
        {
        }
        public static new FastTcpAppArgs Empty => new FastTcpAppArgs();
    }

    public sealed class FastTcpAppState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST TCP application.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Input("existingMonitor")]
        public Input<string>? ExistingMonitor { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP pool.
        /// </summary>
        [Input("existingPool")]
        public Input<string>? ExistingPool { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Input("existingSnatPool")]
        public Input<string>? ExistingSnatPool { get; set; }

        /// <summary>
        /// Json payload for FAST TCP application.
        /// </summary>
        [Input("fastTcpJson")]
        public Input<string>? FastTcpJson { get; set; }

        /// <summary>
        /// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
        /// </summary>
        [Input("loadBalancingMode")]
        public Input<string>? LoadBalancingMode { get; set; }

        /// <summary>
        /// `monitor` block takes input for FAST-Generated Pool Monitor.
        /// See Pool Monitor below for more details.
        /// </summary>
        [Input("monitor")]
        public Input<Inputs.FastTcpAppMonitorGetArgs>? Monitor { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastTcpAppPoolMemberGetArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastTcpAppPoolMemberGetArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastTcpAppPoolMemberGetArgs>());
            set => _poolMembers = value;
        }

        /// <summary>
        /// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
        /// </summary>
        [Input("slowRampTime")]
        public Input<int>? SlowRampTime { get; set; }

        [Input("snatPoolAddresses")]
        private InputList<string>? _snatPoolAddresses;

        /// <summary>
        /// List of address to be used for FAST-Generated SNAT Pool.
        /// </summary>
        public InputList<string> SnatPoolAddresses
        {
            get => _snatPoolAddresses ?? (_snatPoolAddresses = new InputList<string>());
            set => _snatPoolAddresses = value;
        }

        /// <summary>
        /// Name of the FAST TCP application tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastTcpAppVirtualServerGetArgs>? VirtualServer { get; set; }

        public FastTcpAppState()
        {
        }
        public static new FastTcpAppState Empty => new FastTcpAppState();
    }
}
