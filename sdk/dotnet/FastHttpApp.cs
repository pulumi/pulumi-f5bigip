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
    /// `f5bigip.FastHttpApp` This resource will create and manage FAST HTTP applications on BIG-IP
    /// 
    /// [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fastHttpApp = new F5BigIP.FastHttpApp("fastHttpApp", new()
    ///     {
    ///         Application = "fasthttpapp",
    ///         Tenant = "fasthttptenant",
    ///         VirtualServer = new F5BigIP.Inputs.FastHttpAppVirtualServerArgs
    ///         {
    ///             Ip = "10.30.30.44",
    ///             Port = 443,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/fastHttpApp:FastHttpApp")]
    public partial class FastHttpApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTP Application.
        /// </summary>
        [Output("endpointLtmPolicies")]
        public Output<ImmutableArray<string>> EndpointLtmPolicies { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Output("existingMonitor")]
        public Output<string?> ExistingMonitor { get; private set; } = null!;

        /// <summary>
        /// Select an existing BIG-IP Pool
        /// </summary>
        [Output("existingPool")]
        public Output<string?> ExistingPool { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Output("existingSnatPool")]
        public Output<string?> ExistingSnatPool { get; private set; } = null!;

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Output("existingWafSecurityPolicy")]
        public Output<string?> ExistingWafSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// Json payload for FAST HTTP application.
        /// </summary>
        [Output("fastHttpJson")]
        public Output<string> FastHttpJson { get; private set; } = null!;

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
        public Output<Outputs.FastHttpAppMonitor?> Monitor { get; private set; } = null!;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        [Output("poolMembers")]
        public Output<ImmutableArray<Outputs.FastHttpAppPoolMember>> PoolMembers { get; private set; } = null!;

        /// <summary>
        /// List of security log profiles to be used for FAST application
        /// </summary>
        [Output("securityLogProfiles")]
        public Output<ImmutableArray<string>> SecurityLogProfiles { get; private set; } = null!;

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
        /// Name of the FAST HTTPS application tenant.
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Output("virtualServer")]
        public Output<Outputs.FastHttpAppVirtualServer?> VirtualServer { get; private set; } = null!;

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Output("wafSecurityPolicy")]
        public Output<Outputs.FastHttpAppWafSecurityPolicy?> WafSecurityPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a FastHttpApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FastHttpApp(string name, FastHttpAppArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastHttpApp:FastHttpApp", name, args ?? new FastHttpAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FastHttpApp(string name, Input<string> id, FastHttpAppState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastHttpApp:FastHttpApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FastHttpApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FastHttpApp Get(string name, Input<string> id, FastHttpAppState? state = null, CustomResourceOptions? options = null)
        {
            return new FastHttpApp(name, id, state, options);
        }
    }

    public sealed class FastHttpAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        [Input("endpointLtmPolicies")]
        private InputList<string>? _endpointLtmPolicies;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTP Application.
        /// </summary>
        public InputList<string> EndpointLtmPolicies
        {
            get => _endpointLtmPolicies ?? (_endpointLtmPolicies = new InputList<string>());
            set => _endpointLtmPolicies = value;
        }

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Input("existingMonitor")]
        public Input<string>? ExistingMonitor { get; set; }

        /// <summary>
        /// Select an existing BIG-IP Pool
        /// </summary>
        [Input("existingPool")]
        public Input<string>? ExistingPool { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Input("existingSnatPool")]
        public Input<string>? ExistingSnatPool { get; set; }

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Input("existingWafSecurityPolicy")]
        public Input<string>? ExistingWafSecurityPolicy { get; set; }

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
        public Input<Inputs.FastHttpAppMonitorArgs>? Monitor { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastHttpAppPoolMemberArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastHttpAppPoolMemberArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastHttpAppPoolMemberArgs>());
            set => _poolMembers = value;
        }

        [Input("securityLogProfiles")]
        private InputList<string>? _securityLogProfiles;

        /// <summary>
        /// List of security log profiles to be used for FAST application
        /// </summary>
        public InputList<string> SecurityLogProfiles
        {
            get => _securityLogProfiles ?? (_securityLogProfiles = new InputList<string>());
            set => _securityLogProfiles = value;
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
        /// Name of the FAST HTTPS application tenant.
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastHttpAppVirtualServerArgs>? VirtualServer { get; set; }

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Input("wafSecurityPolicy")]
        public Input<Inputs.FastHttpAppWafSecurityPolicyArgs>? WafSecurityPolicy { get; set; }

        public FastHttpAppArgs()
        {
        }
        public static new FastHttpAppArgs Empty => new FastHttpAppArgs();
    }

    public sealed class FastHttpAppState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        [Input("endpointLtmPolicies")]
        private InputList<string>? _endpointLtmPolicies;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTP Application.
        /// </summary>
        public InputList<string> EndpointLtmPolicies
        {
            get => _endpointLtmPolicies ?? (_endpointLtmPolicies = new InputList<string>());
            set => _endpointLtmPolicies = value;
        }

        /// <summary>
        /// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
        /// </summary>
        [Input("existingMonitor")]
        public Input<string>? ExistingMonitor { get; set; }

        /// <summary>
        /// Select an existing BIG-IP Pool
        /// </summary>
        [Input("existingPool")]
        public Input<string>? ExistingPool { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP SNAT pool.
        /// </summary>
        [Input("existingSnatPool")]
        public Input<string>? ExistingSnatPool { get; set; }

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Input("existingWafSecurityPolicy")]
        public Input<string>? ExistingWafSecurityPolicy { get; set; }

        /// <summary>
        /// Json payload for FAST HTTP application.
        /// </summary>
        [Input("fastHttpJson")]
        public Input<string>? FastHttpJson { get; set; }

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
        public Input<Inputs.FastHttpAppMonitorGetArgs>? Monitor { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastHttpAppPoolMemberGetArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastHttpAppPoolMemberGetArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastHttpAppPoolMemberGetArgs>());
            set => _poolMembers = value;
        }

        [Input("securityLogProfiles")]
        private InputList<string>? _securityLogProfiles;

        /// <summary>
        /// List of security log profiles to be used for FAST application
        /// </summary>
        public InputList<string> SecurityLogProfiles
        {
            get => _securityLogProfiles ?? (_securityLogProfiles = new InputList<string>());
            set => _securityLogProfiles = value;
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
        /// Name of the FAST HTTPS application tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastHttpAppVirtualServerGetArgs>? VirtualServer { get; set; }

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Input("wafSecurityPolicy")]
        public Input<Inputs.FastHttpAppWafSecurityPolicyGetArgs>? WafSecurityPolicy { get; set; }

        public FastHttpAppState()
        {
        }
        public static new FastHttpAppState Empty => new FastHttpAppState();
    }
}