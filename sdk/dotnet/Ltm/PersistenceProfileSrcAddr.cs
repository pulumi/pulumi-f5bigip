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
    /// Configures a source address persistence profile
    /// 
    /// ## Reference
    /// 
    /// `name` - (Required) Name of the virtual address
    /// 
    /// `defaults_from` - (Required) Parent cookie persistence profile
    /// 
    /// `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record
    /// 
    /// `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record
    /// 
    /// `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
    /// 
    /// `mirror` (Optional) (enabled or disabled) mirror persistence record
    /// 
    /// `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
    /// 
    /// `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
    /// 
    /// `hash_algorithm` (Optional) Specify the hash algorithm
    /// 
    /// `mask` (Optional) Identify a range of source IP addresses to manage together as a single source address affinity persistent connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
    /// 
    /// `map_proxies` (Optional) (enabled or disabled) Directs all to the same single pool member
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_persistence_profile_srcaddr.html.markdown.
    /// </summary>
    public partial class PersistenceProfileSrcAddr : Pulumi.CustomResource
    {
        [Output("appService")]
        public Output<string?> AppService { get; private set; } = null!;

        /// <summary>
        /// Inherit defaults from parent profile
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specify the hash algorithm
        /// </summary>
        [Output("hashAlgorithm")]
        public Output<string?> HashAlgorithm { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable directs all to the same single pool member
        /// </summary>
        [Output("mapProxies")]
        public Output<string?> MapProxies { get; private set; } = null!;

        /// <summary>
        /// Identify a range of source IP addresses to manage together as a single source address affinity persistent
        /// connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        /// </summary>
        [Output("mask")]
        public Output<string?> Mask { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable match across pools with given persistence record
        /// </summary>
        [Output("matchAcrossPools")]
        public Output<string?> MatchAcrossPools { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Output("matchAcrossServices")]
        public Output<string?> MatchAcrossServices { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Output("matchAcrossVirtuals")]
        public Output<string?> MatchAcrossVirtuals { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable
        /// </summary>
        [Output("mirror")]
        public Output<string?> Mirror { get; private set; } = null!;

        /// <summary>
        /// Name of the persistence profile
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual
        /// connection limits remain hard limits and are not overridden.
        /// </summary>
        [Output("overrideConnLimit")]
        public Output<string?> OverrideConnLimit { get; private set; } = null!;

        /// <summary>
        /// Timeout for persistence of the session
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a PersistenceProfileSrcAddr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PersistenceProfileSrcAddr(string name, PersistenceProfileSrcAddrArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PersistenceProfileSrcAddr(string name, Input<string> id, PersistenceProfileSrcAddrState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PersistenceProfileSrcAddr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PersistenceProfileSrcAddr Get(string name, Input<string> id, PersistenceProfileSrcAddrState? state = null, CustomResourceOptions? options = null)
        {
            return new PersistenceProfileSrcAddr(name, id, state, options);
        }
    }

    public sealed class PersistenceProfileSrcAddrArgs : Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Inherit defaults from parent profile
        /// </summary>
        [Input("defaultsFrom", required: true)]
        public Input<string> DefaultsFrom { get; set; } = null!;

        /// <summary>
        /// Specify the hash algorithm
        /// </summary>
        [Input("hashAlgorithm")]
        public Input<string>? HashAlgorithm { get; set; }

        /// <summary>
        /// To enable _ disable directs all to the same single pool member
        /// </summary>
        [Input("mapProxies")]
        public Input<string>? MapProxies { get; set; }

        /// <summary>
        /// Identify a range of source IP addresses to manage together as a single source address affinity persistent
        /// connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        /// <summary>
        /// To enable _ disable match across pools with given persistence record
        /// </summary>
        [Input("matchAcrossPools")]
        public Input<string>? MatchAcrossPools { get; set; }

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Input("matchAcrossServices")]
        public Input<string>? MatchAcrossServices { get; set; }

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Input("matchAcrossVirtuals")]
        public Input<string>? MatchAcrossVirtuals { get; set; }

        /// <summary>
        /// To enable _ disable
        /// </summary>
        [Input("mirror")]
        public Input<string>? Mirror { get; set; }

        /// <summary>
        /// Name of the persistence profile
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual
        /// connection limits remain hard limits and are not overridden.
        /// </summary>
        [Input("overrideConnLimit")]
        public Input<string>? OverrideConnLimit { get; set; }

        /// <summary>
        /// Timeout for persistence of the session
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public PersistenceProfileSrcAddrArgs()
        {
        }
    }

    public sealed class PersistenceProfileSrcAddrState : Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Inherit defaults from parent profile
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specify the hash algorithm
        /// </summary>
        [Input("hashAlgorithm")]
        public Input<string>? HashAlgorithm { get; set; }

        /// <summary>
        /// To enable _ disable directs all to the same single pool member
        /// </summary>
        [Input("mapProxies")]
        public Input<string>? MapProxies { get; set; }

        /// <summary>
        /// Identify a range of source IP addresses to manage together as a single source address affinity persistent
        /// connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        /// <summary>
        /// To enable _ disable match across pools with given persistence record
        /// </summary>
        [Input("matchAcrossPools")]
        public Input<string>? MatchAcrossPools { get; set; }

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Input("matchAcrossServices")]
        public Input<string>? MatchAcrossServices { get; set; }

        /// <summary>
        /// To enable _ disable match across services with given persistence record
        /// </summary>
        [Input("matchAcrossVirtuals")]
        public Input<string>? MatchAcrossVirtuals { get; set; }

        /// <summary>
        /// To enable _ disable
        /// </summary>
        [Input("mirror")]
        public Input<string>? Mirror { get; set; }

        /// <summary>
        /// Name of the persistence profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual
        /// connection limits remain hard limits and are not overridden.
        /// </summary>
        [Input("overrideConnLimit")]
        public Input<string>? OverrideConnLimit { get; set; }

        /// <summary>
        /// Timeout for persistence of the session
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public PersistenceProfileSrcAddrState()
        {
        }
    }
}
