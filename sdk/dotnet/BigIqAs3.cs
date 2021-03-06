// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    [F5BigIPResourceType("f5bigip:index/bigIqAs3:BigIqAs3")]
    public partial class BigIqAs3 : Pulumi.CustomResource
    {
        /// <summary>
        /// AS3 json
        /// </summary>
        [Output("as3Json")]
        public Output<string> As3Json { get; private set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Output("bigiqAddress")]
        public Output<string> BigiqAddress { get; private set; } = null!;

        /// <summary>
        /// Login reference for token authentication (see BIG-IQ REST docs for details)
        /// </summary>
        [Output("bigiqLoginRef")]
        public Output<string?> BigiqLoginRef { get; private set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Output("bigiqPassword")]
        public Output<string> BigiqPassword { get; private set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Output("bigiqPort")]
        public Output<string?> BigiqPort { get; private set; } = null!;

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        [Output("bigiqTokenAuth")]
        public Output<bool?> BigiqTokenAuth { get; private set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Output("bigiqUser")]
        public Output<string> BigiqUser { get; private set; } = null!;

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Output("tenantList")]
        public Output<string> TenantList { get; private set; } = null!;


        /// <summary>
        /// Create a BigIqAs3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BigIqAs3(string name, BigIqAs3Args args, CustomResourceOptions? options = null)
            : base("f5bigip:index/bigIqAs3:BigIqAs3", name, args ?? new BigIqAs3Args(), MakeResourceOptions(options, ""))
        {
        }

        private BigIqAs3(string name, Input<string> id, BigIqAs3State? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/bigIqAs3:BigIqAs3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BigIqAs3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BigIqAs3 Get(string name, Input<string> id, BigIqAs3State? state = null, CustomResourceOptions? options = null)
        {
            return new BigIqAs3(name, id, state, options);
        }
    }

    public sealed class BigIqAs3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AS3 json
        /// </summary>
        [Input("as3Json", required: true)]
        public Input<string> As3Json { get; set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqAddress", required: true)]
        public Input<string> BigiqAddress { get; set; } = null!;

        /// <summary>
        /// Login reference for token authentication (see BIG-IQ REST docs for details)
        /// </summary>
        [Input("bigiqLoginRef")]
        public Input<string>? BigiqLoginRef { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqPassword", required: true)]
        public Input<string> BigiqPassword { get; set; } = null!;

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqPort")]
        public Input<string>? BigiqPort { get; set; }

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        [Input("bigiqTokenAuth")]
        public Input<bool>? BigiqTokenAuth { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqUser", required: true)]
        public Input<string> BigiqUser { get; set; } = null!;

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        public BigIqAs3Args()
        {
        }
    }

    public sealed class BigIqAs3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AS3 json
        /// </summary>
        [Input("as3Json")]
        public Input<string>? As3Json { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqAddress")]
        public Input<string>? BigiqAddress { get; set; }

        /// <summary>
        /// Login reference for token authentication (see BIG-IQ REST docs for details)
        /// </summary>
        [Input("bigiqLoginRef")]
        public Input<string>? BigiqLoginRef { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqPassword")]
        public Input<string>? BigiqPassword { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqPort")]
        public Input<string>? BigiqPort { get; set; }

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        [Input("bigiqTokenAuth")]
        public Input<bool>? BigiqTokenAuth { get; set; }

        /// <summary>
        /// The registration key pool to use
        /// </summary>
        [Input("bigiqUser")]
        public Input<string>? BigiqUser { get; set; }

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        public BigIqAs3State()
        {
        }
    }
}
