// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys
{
    /// <summary>
    /// `f5bigip.sys.Dns` Configures DNS server on F5 BIG-IP
    /// </summary>
    public partial class Dns : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Dns Servers
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name or IP address of the DNS server
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// Configures the number of dots needed in a name before an initial absolute query will be made.
        /// </summary>
        [Output("numberOfDots")]
        public Output<int?> NumberOfDots { get; private set; } = null!;

        /// <summary>
        /// Specify what domains you want to search
        /// </summary>
        [Output("searches")]
        public Output<ImmutableArray<string>> Searches { get; private set; } = null!;


        /// <summary>
        /// Create a Dns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dns(string name, DnsArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:sys/dns:Dns", name, args ?? new DnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dns(string name, Input<string> id, DnsState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/dns:Dns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dns Get(string name, Input<string> id, DnsState? state = null, CustomResourceOptions? options = null)
        {
            return new Dns(name, id, state, options);
        }
    }

    public sealed class DnsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Dns Servers
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// Name or IP address of the DNS server
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        /// <summary>
        /// Configures the number of dots needed in a name before an initial absolute query will be made.
        /// </summary>
        [Input("numberOfDots")]
        public Input<int>? NumberOfDots { get; set; }

        [Input("searches")]
        private InputList<string>? _searches;

        /// <summary>
        /// Specify what domains you want to search
        /// </summary>
        public InputList<string> Searches
        {
            get => _searches ?? (_searches = new InputList<string>());
            set => _searches = value;
        }

        public DnsArgs()
        {
        }
    }

    public sealed class DnsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Dns Servers
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// Name or IP address of the DNS server
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        /// <summary>
        /// Configures the number of dots needed in a name before an initial absolute query will be made.
        /// </summary>
        [Input("numberOfDots")]
        public Input<int>? NumberOfDots { get; set; }

        [Input("searches")]
        private InputList<string>? _searches;

        /// <summary>
        /// Specify what domains you want to search
        /// </summary>
        public InputList<string> Searches
        {
            get => _searches ?? (_searches = new InputList<string>());
            set => _searches = value;
        }

        public DnsState()
        {
        }
    }
}
