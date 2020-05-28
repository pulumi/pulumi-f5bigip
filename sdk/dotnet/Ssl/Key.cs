// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    /// <summary>
    /// `f5bigip.ssl.Key` This resource will import SSL certificate key on BIG-IP LTM. 
    /// Certificate key can be imported from certificate key files on the local disk, in PEM format
    /// 
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test_key = new F5BigIP.Ssl.Key("test-key", new F5BigIP.Ssl.KeyArgs
    ///         {
    ///             Name = "serverkey.key",
    ///             Content = File.ReadAllText("serverkey.key"),
    ///             Partition = "Common",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Key : Pulumi.CustomResource
    {
        /// <summary>
        /// Content of SSL certificate key present on local Disk
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Name of the SSL Certificate key to be Imported on to BIGIP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Partition on to SSL Certificate key to be imported
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;


        /// <summary>
        /// Create a Key resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Key(string name, KeyArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ssl/key:Key", name, args ?? new KeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Key(string name, Input<string> id, KeyState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ssl/key:Key", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Key resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Key Get(string name, Input<string> id, KeyState? state = null, CustomResourceOptions? options = null)
        {
            return new Key(name, id, state, options);
        }
    }

    public sealed class KeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content of SSL certificate key present on local Disk
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Name of the SSL Certificate key to be Imported on to BIGIP
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Partition on to SSL Certificate key to be imported
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public KeyArgs()
        {
        }
    }

    public sealed class KeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content of SSL certificate key present on local Disk
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Name of the SSL Certificate key to be Imported on to BIGIP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Partition on to SSL Certificate key to be imported
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public KeyState()
        {
        }
    }
}
