// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.VirtualServer` Configures Virtual Server
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_virtual_server.html.markdown.
    /// </summary>
    public partial class VirtualServer : Pulumi.CustomResource
    {
        /// <summary>
        /// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
        /// </summary>
        [Output("clientProfiles")]
        public Output<ImmutableArray<string>> ClientProfiles { get; private set; } = null!;

        [Output("defaultPersistenceProfile")]
        public Output<string?> DefaultPersistenceProfile { get; private set; } = null!;

        /// <summary>
        /// Description of Virtual server
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Destination IP
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
        /// </summary>
        [Output("fallbackPersistenceProfile")]
        public Output<string> FallbackPersistenceProfile { get; private set; } = null!;

        /// <summary>
        /// all, tcp, udp
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
        /// </summary>
        [Output("irules")]
        public Output<ImmutableArray<string>> Irules { get; private set; } = null!;

        /// <summary>
        /// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
        /// </summary>
        [Output("mask")]
        public Output<string?> Mask { get; private set; } = null!;

        /// <summary>
        /// Name of the virtual server
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of persistence profiles associated with the Virtual Server.
        /// </summary>
        [Output("persistenceProfiles")]
        public Output<ImmutableArray<string>> PersistenceProfiles { get; private set; } = null!;

        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Default pool name
        /// </summary>
        [Output("pool")]
        public Output<string?> Pool { get; private set; } = null!;

        /// <summary>
        /// Listen port for the virtual server
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
        /// </summary>
        [Output("profiles")]
        public Output<ImmutableArray<string>> Profiles { get; private set; } = null!;

        /// <summary>
        /// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
        /// </summary>
        [Output("serverProfiles")]
        public Output<ImmutableArray<string>> ServerProfiles { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
        /// </summary>
        [Output("snatpool")]
        public Output<string> Snatpool { get; private set; } = null!;

        /// <summary>
        /// Specifies an IP address or network from which the virtual server will accept traffic.
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;

        /// <summary>
        /// Can be either omitted for none or the values automap or snat
        /// </summary>
        [Output("sourceAddressTranslation")]
        public Output<string> SourceAddressTranslation { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the virtual server and its resources are available for load balancing. The default is
        /// Enabled
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
        /// </summary>
        [Output("translateAddress")]
        public Output<string> TranslateAddress { get; private set; } = null!;

        /// <summary>
        /// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
        /// </summary>
        [Output("translatePort")]
        public Output<string> TranslatePort { get; private set; } = null!;

        /// <summary>
        /// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
        /// </summary>
        [Output("vlans")]
        public Output<ImmutableArray<string>> Vlans { get; private set; } = null!;

        /// <summary>
        /// Enables the virtual server on the VLANs specified by the VLANs option.
        /// </summary>
        [Output("vlansEnabled")]
        public Output<bool> VlansEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualServer(string name, VirtualServerArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/virtualServer:VirtualServer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VirtualServer(string name, Input<string> id, VirtualServerState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/virtualServer:VirtualServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualServer Get(string name, Input<string> id, VirtualServerState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualServer(name, id, state, options);
        }
    }

    public sealed class VirtualServerArgs : Pulumi.ResourceArgs
    {
        [Input("clientProfiles")]
        private InputList<string>? _clientProfiles;

        /// <summary>
        /// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
        /// </summary>
        public InputList<string> ClientProfiles
        {
            get => _clientProfiles ?? (_clientProfiles = new InputList<string>());
            set => _clientProfiles = value;
        }

        [Input("defaultPersistenceProfile")]
        public Input<string>? DefaultPersistenceProfile { get; set; }

        /// <summary>
        /// Description of Virtual server
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination IP
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
        /// </summary>
        [Input("fallbackPersistenceProfile")]
        public Input<string>? FallbackPersistenceProfile { get; set; }

        /// <summary>
        /// all, tcp, udp
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("irules")]
        private InputList<string>? _irules;

        /// <summary>
        /// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
        /// </summary>
        public InputList<string> Irules
        {
            get => _irules ?? (_irules = new InputList<string>());
            set => _irules = value;
        }

        /// <summary>
        /// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        /// <summary>
        /// Name of the virtual server
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("persistenceProfiles")]
        private InputList<string>? _persistenceProfiles;

        /// <summary>
        /// List of persistence profiles associated with the Virtual Server.
        /// </summary>
        public InputList<string> PersistenceProfiles
        {
            get => _persistenceProfiles ?? (_persistenceProfiles = new InputList<string>());
            set => _persistenceProfiles = value;
        }

        [Input("policies")]
        private InputList<string>? _policies;
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Default pool name
        /// </summary>
        [Input("pool")]
        public Input<string>? Pool { get; set; }

        /// <summary>
        /// Listen port for the virtual server
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("profiles")]
        private InputList<string>? _profiles;

        /// <summary>
        /// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
        /// </summary>
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        [Input("serverProfiles")]
        private InputList<string>? _serverProfiles;

        /// <summary>
        /// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
        /// </summary>
        public InputList<string> ServerProfiles
        {
            get => _serverProfiles ?? (_serverProfiles = new InputList<string>());
            set => _serverProfiles = value;
        }

        /// <summary>
        /// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
        /// </summary>
        [Input("snatpool")]
        public Input<string>? Snatpool { get; set; }

        /// <summary>
        /// Specifies an IP address or network from which the virtual server will accept traffic.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Can be either omitted for none or the values automap or snat
        /// </summary>
        [Input("sourceAddressTranslation")]
        public Input<string>? SourceAddressTranslation { get; set; }

        /// <summary>
        /// Specifies whether the virtual server and its resources are available for load balancing. The default is
        /// Enabled
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
        /// </summary>
        [Input("translateAddress")]
        public Input<string>? TranslateAddress { get; set; }

        /// <summary>
        /// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
        /// </summary>
        [Input("translatePort")]
        public Input<string>? TranslatePort { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        /// <summary>
        /// Enables the virtual server on the VLANs specified by the VLANs option.
        /// </summary>
        [Input("vlansEnabled")]
        public Input<bool>? VlansEnabled { get; set; }

        public VirtualServerArgs()
        {
        }
    }

    public sealed class VirtualServerState : Pulumi.ResourceArgs
    {
        [Input("clientProfiles")]
        private InputList<string>? _clientProfiles;

        /// <summary>
        /// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
        /// </summary>
        public InputList<string> ClientProfiles
        {
            get => _clientProfiles ?? (_clientProfiles = new InputList<string>());
            set => _clientProfiles = value;
        }

        [Input("defaultPersistenceProfile")]
        public Input<string>? DefaultPersistenceProfile { get; set; }

        /// <summary>
        /// Description of Virtual server
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination IP
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
        /// </summary>
        [Input("fallbackPersistenceProfile")]
        public Input<string>? FallbackPersistenceProfile { get; set; }

        /// <summary>
        /// all, tcp, udp
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("irules")]
        private InputList<string>? _irules;

        /// <summary>
        /// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
        /// </summary>
        public InputList<string> Irules
        {
            get => _irules ?? (_irules = new InputList<string>());
            set => _irules = value;
        }

        /// <summary>
        /// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        /// <summary>
        /// Name of the virtual server
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("persistenceProfiles")]
        private InputList<string>? _persistenceProfiles;

        /// <summary>
        /// List of persistence profiles associated with the Virtual Server.
        /// </summary>
        public InputList<string> PersistenceProfiles
        {
            get => _persistenceProfiles ?? (_persistenceProfiles = new InputList<string>());
            set => _persistenceProfiles = value;
        }

        [Input("policies")]
        private InputList<string>? _policies;
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Default pool name
        /// </summary>
        [Input("pool")]
        public Input<string>? Pool { get; set; }

        /// <summary>
        /// Listen port for the virtual server
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("profiles")]
        private InputList<string>? _profiles;

        /// <summary>
        /// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
        /// </summary>
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        [Input("serverProfiles")]
        private InputList<string>? _serverProfiles;

        /// <summary>
        /// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
        /// </summary>
        public InputList<string> ServerProfiles
        {
            get => _serverProfiles ?? (_serverProfiles = new InputList<string>());
            set => _serverProfiles = value;
        }

        /// <summary>
        /// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
        /// </summary>
        [Input("snatpool")]
        public Input<string>? Snatpool { get; set; }

        /// <summary>
        /// Specifies an IP address or network from which the virtual server will accept traffic.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Can be either omitted for none or the values automap or snat
        /// </summary>
        [Input("sourceAddressTranslation")]
        public Input<string>? SourceAddressTranslation { get; set; }

        /// <summary>
        /// Specifies whether the virtual server and its resources are available for load balancing. The default is
        /// Enabled
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
        /// </summary>
        [Input("translateAddress")]
        public Input<string>? TranslateAddress { get; set; }

        /// <summary>
        /// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
        /// </summary>
        [Input("translatePort")]
        public Input<string>? TranslatePort { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        /// <summary>
        /// Enables the virtual server on the VLANs specified by the VLANs option.
        /// </summary>
        [Input("vlansEnabled")]
        public Input<bool>? VlansEnabled { get; set; }

        public VirtualServerState()
        {
        }
    }
}