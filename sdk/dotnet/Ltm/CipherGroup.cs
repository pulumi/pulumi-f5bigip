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
    /// `f5bigip.ltm.CipherGroup` Manages F5 BIG-IP LTM cipher group using iControl REST.
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
    ///     var test_cipher_group = new F5BigIP.Ltm.CipherGroup("test-cipher-group", new()
    ///     {
    ///         Allows = new[]
    ///         {
    ///             "/Common/f5-aes",
    ///         },
    ///         Name = "/Common/test-cipher-group-01",
    ///         Ordering = "speed",
    ///         Requires = new[]
    ///         {
    ///             "/Common/f5-quic",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/cipherGroup:CipherGroup")]
    public partial class CipherGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        [Output("allows")]
        public Output<ImmutableArray<string>> Allows { get; private set; } = null!;

        /// <summary>
        /// Specifies descriptive text that identifies the cipher rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        /// </summary>
        [Output("ordering")]
        public Output<string?> Ordering { get; private set; } = null!;

        /// <summary>
        /// Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        [Output("requires")]
        public Output<ImmutableArray<string>> Requires { get; private set; } = null!;


        /// <summary>
        /// Create a CipherGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CipherGroup(string name, CipherGroupArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/cipherGroup:CipherGroup", name, args ?? new CipherGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CipherGroup(string name, Input<string> id, CipherGroupState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/cipherGroup:CipherGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CipherGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CipherGroup Get(string name, Input<string> id, CipherGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new CipherGroup(name, id, state, options);
        }
    }

    public sealed class CipherGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<string>? _allows;

        /// <summary>
        /// Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        public InputList<string> Allows
        {
            get => _allows ?? (_allows = new InputList<string>());
            set => _allows = value;
        }

        /// <summary>
        /// Specifies descriptive text that identifies the cipher rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        /// </summary>
        [Input("ordering")]
        public Input<string>? Ordering { get; set; }

        [Input("requires")]
        private InputList<string>? _requires;

        /// <summary>
        /// Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        public InputList<string> Requires
        {
            get => _requires ?? (_requires = new InputList<string>());
            set => _requires = value;
        }

        public CipherGroupArgs()
        {
        }
        public static new CipherGroupArgs Empty => new CipherGroupArgs();
    }

    public sealed class CipherGroupState : global::Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<string>? _allows;

        /// <summary>
        /// Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        public InputList<string> Allows
        {
            get => _allows ?? (_allows = new InputList<string>());
            set => _allows = value;
        }

        /// <summary>
        /// Specifies descriptive text that identifies the cipher rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        /// </summary>
        [Input("ordering")]
        public Input<string>? Ordering { get; set; }

        [Input("requires")]
        private InputList<string>? _requires;

        /// <summary>
        /// Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        /// </summary>
        public InputList<string> Requires
        {
            get => _requires ?? (_requires = new InputList<string>());
            set => _requires = value;
        }

        public CipherGroupState()
        {
        }
        public static new CipherGroupState Empty => new CipherGroupState();
    }
}
