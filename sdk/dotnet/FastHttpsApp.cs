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
    /// `f5bigip.FastHttpsApp` This resource will create and manage FAST HTTPS applications on BIG-IP
    /// 
    /// [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
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
    ///     var fastHttpsApp = new F5BigIP.FastHttpsApp("fastHttpsApp", new()
    ///     {
    ///         Application = "fasthttpsapp",
    ///         Tenant = "fasthttpstenant",
    ///         VirtualServer = new F5BigIP.Inputs.FastHttpsAppVirtualServerArgs
    ///         {
    ///             Ip = "10.30.40.44",
    ///             Port = 443,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With Service Discovery
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var tC3AzureServiceDiscovery = F5BigIP.Fast.GetAzureServiceDiscovery.Invoke(new()
    ///     {
    ///         ResourceGroup = "testazurerg",
    ///         SubscriptionId = "testazuresid",
    ///         TagKey = "testazuretag",
    ///         TagValue = "testazurevalue",
    ///     });
    /// 
    ///     var tC3GceServiceDiscovery = F5BigIP.Fast.GetGceServiceDiscovery.Invoke(new()
    ///     {
    ///         TagKey = "testgcetag",
    ///         TagValue = "testgcevalue",
    ///         Region = "testgceregion",
    ///     });
    /// 
    ///     var fastHttpsApp = new F5BigIP.FastHttpsApp("fastHttpsApp", new()
    ///     {
    ///         Tenant = "fasthttpstenant",
    ///         Application = "fasthttpsapp",
    ///         VirtualServer = new F5BigIP.Inputs.FastHttpsAppVirtualServerArgs
    ///         {
    ///             Ip = "10.30.40.44",
    ///             Port = 443,
    ///         },
    ///         PoolMembers = new[]
    ///         {
    ///             new F5BigIP.Inputs.FastHttpsAppPoolMemberArgs
    ///             {
    ///                 Addresses = new[]
    ///                 {
    ///                     "10.11.40.120",
    ///                     "10.11.30.121",
    ///                     "10.11.30.122",
    ///                 },
    ///                 Port = 80,
    ///             },
    ///         },
    ///         ServiceDiscoveries = new[]
    ///         {
    ///             tC3GceServiceDiscovery.Apply(getGceServiceDiscoveryResult =&gt; getGceServiceDiscoveryResult.GceSdJson),
    ///             tC3AzureServiceDiscovery.Apply(getAzureServiceDiscoveryResult =&gt; getAzureServiceDiscoveryResult.AzureSdJson),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/fastHttpsApp:FastHttpsApp")]
    public partial class FastHttpsApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTPS Application.
        /// </summary>
        [Output("endpointLtmPolicies")]
        public Output<ImmutableArray<string>> EndpointLtmPolicies { get; private set; } = null!;

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
        /// Name of an existing TLS client profile.
        /// </summary>
        [Output("existingTlsClientProfile")]
        public Output<string?> ExistingTlsClientProfile { get; private set; } = null!;

        /// <summary>
        /// Name of an existing TLS server profile.
        /// </summary>
        [Output("existingTlsServerProfile")]
        public Output<string?> ExistingTlsServerProfile { get; private set; } = null!;

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Output("existingWafSecurityPolicy")]
        public Output<string?> ExistingWafSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// Type of fallback persistence record to be created for each new client connection.
        /// </summary>
        [Output("fallbackPersistence")]
        public Output<string?> FallbackPersistence { get; private set; } = null!;

        /// <summary>
        /// Json payload for FAST HTTPS application.
        /// </summary>
        [Output("fastHttpsJson")]
        public Output<string> FastHttpsJson { get; private set; } = null!;

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
        public Output<Outputs.FastHttpsAppMonitor?> Monitor { get; private set; } = null!;

        /// <summary>
        /// Name of an existing BIG-IP persistence profile to be used.
        /// </summary>
        [Output("persistenceProfile")]
        public Output<string?> PersistenceProfile { get; private set; } = null!;

        /// <summary>
        /// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
        /// </summary>
        [Output("persistenceType")]
        public Output<string?> PersistenceType { get; private set; } = null!;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        [Output("poolMembers")]
        public Output<ImmutableArray<Outputs.FastHttpsAppPoolMember>> PoolMembers { get; private set; } = null!;

        /// <summary>
        /// List of security log profiles to be used for FAST application
        /// </summary>
        [Output("securityLogProfiles")]
        public Output<ImmutableArray<string>> SecurityLogProfiles { get; private set; } = null!;

        /// <summary>
        /// List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
        /// </summary>
        [Output("serviceDiscoveries")]
        public Output<ImmutableArray<string>> ServiceDiscoveries { get; private set; } = null!;

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
        /// `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
        /// See TLS Client Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_client_profile` or `tls_client_profile` used for encrypt server-side connections.
        /// </summary>
        [Output("tlsClientProfile")]
        public Output<Outputs.FastHttpsAppTlsClientProfile?> TlsClientProfile { get; private set; } = null!;

        /// <summary>
        /// `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
        /// See TLS Server Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_server_profile` or `tls_server_profile` used for decrypt client-side connections.
        /// </summary>
        [Output("tlsServerProfile")]
        public Output<Outputs.FastHttpsAppTlsServerProfile?> TlsServerProfile { get; private set; } = null!;

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Output("virtualServer")]
        public Output<Outputs.FastHttpsAppVirtualServer?> VirtualServer { get; private set; } = null!;

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Output("wafSecurityPolicy")]
        public Output<Outputs.FastHttpsAppWafSecurityPolicy?> WafSecurityPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a FastHttpsApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FastHttpsApp(string name, FastHttpsAppArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastHttpsApp:FastHttpsApp", name, args ?? new FastHttpsAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FastHttpsApp(string name, Input<string> id, FastHttpsAppState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastHttpsApp:FastHttpsApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FastHttpsApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FastHttpsApp Get(string name, Input<string> id, FastHttpsAppState? state = null, CustomResourceOptions? options = null)
        {
            return new FastHttpsApp(name, id, state, options);
        }
    }

    public sealed class FastHttpsAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        [Input("endpointLtmPolicies")]
        private InputList<string>? _endpointLtmPolicies;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTPS Application.
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
        /// Name of an existing TLS client profile.
        /// </summary>
        [Input("existingTlsClientProfile")]
        public Input<string>? ExistingTlsClientProfile { get; set; }

        /// <summary>
        /// Name of an existing TLS server profile.
        /// </summary>
        [Input("existingTlsServerProfile")]
        public Input<string>? ExistingTlsServerProfile { get; set; }

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Input("existingWafSecurityPolicy")]
        public Input<string>? ExistingWafSecurityPolicy { get; set; }

        /// <summary>
        /// Type of fallback persistence record to be created for each new client connection.
        /// </summary>
        [Input("fallbackPersistence")]
        public Input<string>? FallbackPersistence { get; set; }

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
        public Input<Inputs.FastHttpsAppMonitorArgs>? Monitor { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP persistence profile to be used.
        /// </summary>
        [Input("persistenceProfile")]
        public Input<string>? PersistenceProfile { get; set; }

        /// <summary>
        /// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
        /// </summary>
        [Input("persistenceType")]
        public Input<string>? PersistenceType { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastHttpsAppPoolMemberArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastHttpsAppPoolMemberArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastHttpsAppPoolMemberArgs>());
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

        [Input("serviceDiscoveries")]
        private InputList<string>? _serviceDiscoveries;

        /// <summary>
        /// List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
        /// </summary>
        public InputList<string> ServiceDiscoveries
        {
            get => _serviceDiscoveries ?? (_serviceDiscoveries = new InputList<string>());
            set => _serviceDiscoveries = value;
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
        /// `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
        /// See TLS Client Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_client_profile` or `tls_client_profile` used for encrypt server-side connections.
        /// </summary>
        [Input("tlsClientProfile")]
        public Input<Inputs.FastHttpsAppTlsClientProfileArgs>? TlsClientProfile { get; set; }

        /// <summary>
        /// `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
        /// See TLS Server Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_server_profile` or `tls_server_profile` used for decrypt client-side connections.
        /// </summary>
        [Input("tlsServerProfile")]
        public Input<Inputs.FastHttpsAppTlsServerProfileArgs>? TlsServerProfile { get; set; }

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastHttpsAppVirtualServerArgs>? VirtualServer { get; set; }

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Input("wafSecurityPolicy")]
        public Input<Inputs.FastHttpsAppWafSecurityPolicyArgs>? WafSecurityPolicy { get; set; }

        public FastHttpsAppArgs()
        {
        }
        public static new FastHttpsAppArgs Empty => new FastHttpsAppArgs();
    }

    public sealed class FastHttpsAppState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the FAST HTTPS application.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        [Input("endpointLtmPolicies")]
        private InputList<string>? _endpointLtmPolicies;

        /// <summary>
        /// List of LTM Policies to be applied FAST HTTPS Application.
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
        /// Name of an existing TLS client profile.
        /// </summary>
        [Input("existingTlsClientProfile")]
        public Input<string>? ExistingTlsClientProfile { get; set; }

        /// <summary>
        /// Name of an existing TLS server profile.
        /// </summary>
        [Input("existingTlsServerProfile")]
        public Input<string>? ExistingTlsServerProfile { get; set; }

        /// <summary>
        /// Name of an existing WAF Security policy.
        /// </summary>
        [Input("existingWafSecurityPolicy")]
        public Input<string>? ExistingWafSecurityPolicy { get; set; }

        /// <summary>
        /// Type of fallback persistence record to be created for each new client connection.
        /// </summary>
        [Input("fallbackPersistence")]
        public Input<string>? FallbackPersistence { get; set; }

        /// <summary>
        /// Json payload for FAST HTTPS application.
        /// </summary>
        [Input("fastHttpsJson")]
        public Input<string>? FastHttpsJson { get; set; }

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
        public Input<Inputs.FastHttpsAppMonitorGetArgs>? Monitor { get; set; }

        /// <summary>
        /// Name of an existing BIG-IP persistence profile to be used.
        /// </summary>
        [Input("persistenceProfile")]
        public Input<string>? PersistenceProfile { get; set; }

        /// <summary>
        /// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
        /// </summary>
        [Input("persistenceType")]
        public Input<string>? PersistenceType { get; set; }

        [Input("poolMembers")]
        private InputList<Inputs.FastHttpsAppPoolMemberGetArgs>? _poolMembers;

        /// <summary>
        /// `pool_members` block takes input for FAST-Generated Pool.
        /// See Pool Members below for more details.
        /// </summary>
        public InputList<Inputs.FastHttpsAppPoolMemberGetArgs> PoolMembers
        {
            get => _poolMembers ?? (_poolMembers = new InputList<Inputs.FastHttpsAppPoolMemberGetArgs>());
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

        [Input("serviceDiscoveries")]
        private InputList<string>? _serviceDiscoveries;

        /// <summary>
        /// List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
        /// </summary>
        public InputList<string> ServiceDiscoveries
        {
            get => _serviceDiscoveries ?? (_serviceDiscoveries = new InputList<string>());
            set => _serviceDiscoveries = value;
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
        /// `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
        /// See TLS Client Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_client_profile` or `tls_client_profile` used for encrypt server-side connections.
        /// </summary>
        [Input("tlsClientProfile")]
        public Input<Inputs.FastHttpsAppTlsClientProfileGetArgs>? TlsClientProfile { get; set; }

        /// <summary>
        /// `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
        /// See TLS Server Profile below for more details.
        /// 
        /// &gt; **NOTE** Profile provided by `existing_tls_server_profile` or `tls_server_profile` used for decrypt client-side connections.
        /// </summary>
        [Input("tlsServerProfile")]
        public Input<Inputs.FastHttpsAppTlsServerProfileGetArgs>? TlsServerProfile { get; set; }

        /// <summary>
        /// `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
        /// See virtual server below for more details.
        /// </summary>
        [Input("virtualServer")]
        public Input<Inputs.FastHttpsAppVirtualServerGetArgs>? VirtualServer { get; set; }

        /// <summary>
        /// `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
        /// See WAF Security Policy below for more details.
        /// </summary>
        [Input("wafSecurityPolicy")]
        public Input<Inputs.FastHttpsAppWafSecurityPolicyGetArgs>? WafSecurityPolicy { get; set; }

        public FastHttpsAppState()
        {
        }
        public static new FastHttpsAppState Empty => new FastHttpsAppState();
    }
}
