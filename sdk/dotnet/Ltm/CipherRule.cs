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
    /// `f5bigip.ltm.CipherRule` Manages F5 BIG-IP LTM cipher rule using iControl REST.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testCipherRule = new F5BigIP.Ltm.CipherRule("test_cipher_rule", new()
    ///     {
    ///         Name = "/Common/test_cipher_rule",
    ///         Cipher = "TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384",
    ///         DhGroups = "P256:P384:FFDHE2048:FFDHE3072:FFDHE4096",
    ///         SignatureAlgorithms = "DEFAULT",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Importing
    /// 
    /// An existing cipher rule can be imported into this resource by supplying the cipher rule full path name  ex : `/partition/name`
    /// An example is below:
    /// ```sh
    /// $ terraform import bigip_ltm_cipher_rule.test_cipher_rule /Common/test_cipher_rule
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/cipherRule:CipherRule")]
    public partial class CipherRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        /// </summary>
        [Output("cipher")]
        public Output<string> Cipher { get; private set; } = null!;

        /// <summary>
        /// The Partition in which the Cipher Rule will be created.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the DH Groups algorithms, separated by colons (:).
        /// </summary>
        [Output("dhGroups")]
        public Output<string> DhGroups { get; private set; } = null!;

        /// <summary>
        /// Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Signature Algorithms, separated by colons (:).
        /// </summary>
        [Output("signatureAlgorithms")]
        public Output<string> SignatureAlgorithms { get; private set; } = null!;


        /// <summary>
        /// Create a CipherRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CipherRule(string name, CipherRuleArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/cipherRule:CipherRule", name, args ?? new CipherRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CipherRule(string name, Input<string> id, CipherRuleState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/cipherRule:CipherRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CipherRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CipherRule Get(string name, Input<string> id, CipherRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new CipherRule(name, id, state, options);
        }
    }

    public sealed class CipherRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        /// </summary>
        [Input("cipher", required: true)]
        public Input<string> Cipher { get; set; } = null!;

        /// <summary>
        /// The Partition in which the Cipher Rule will be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the DH Groups algorithms, separated by colons (:).
        /// </summary>
        [Input("dhGroups")]
        public Input<string>? DhGroups { get; set; }

        /// <summary>
        /// Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Signature Algorithms, separated by colons (:).
        /// </summary>
        [Input("signatureAlgorithms")]
        public Input<string>? SignatureAlgorithms { get; set; }

        public CipherRuleArgs()
        {
        }
        public static new CipherRuleArgs Empty => new CipherRuleArgs();
    }

    public sealed class CipherRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        /// </summary>
        [Input("cipher")]
        public Input<string>? Cipher { get; set; }

        /// <summary>
        /// The Partition in which the Cipher Rule will be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the DH Groups algorithms, separated by colons (:).
        /// </summary>
        [Input("dhGroups")]
        public Input<string>? DhGroups { get; set; }

        /// <summary>
        /// Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Signature Algorithms, separated by colons (:).
        /// </summary>
        [Input("signatureAlgorithms")]
        public Input<string>? SignatureAlgorithms { get; set; }

        public CipherRuleState()
        {
        }
        public static new CipherRuleState Empty => new CipherRuleState();
    }
}
