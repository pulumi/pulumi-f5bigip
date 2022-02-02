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
    /// `f5bigip.FastApplication` This resource will create and manage FAST applications on BIG-IP from provided JSON declaration.
    /// 
    /// ## Example Usage
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
    ///         var foo_app = new F5BigIP.FastApplication("foo-app", new F5BigIP.FastApplicationArgs
    ///         {
    ///             FastJson = File.ReadAllText("new_fast_app.json"),
    ///             Template = "examples/simple_http",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/fastApplication:FastApplication")]
    public partial class FastApplication : Pulumi.CustomResource
    {
        /// <summary>
        /// A FAST application name.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Output("fastJson")]
        public Output<string> FastJson { get; private set; } = null!;

        /// <summary>
        /// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
        /// </summary>
        [Output("template")]
        public Output<string?> Template { get; private set; } = null!;

        /// <summary>
        /// A FAST tenant name on which you want to manage application.
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a FastApplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FastApplication(string name, FastApplicationArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastApplication:FastApplication", name, args ?? new FastApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FastApplication(string name, Input<string> id, FastApplicationState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/fastApplication:FastApplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FastApplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FastApplication Get(string name, Input<string> id, FastApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new FastApplication(name, id, state, options);
        }
    }

    public sealed class FastApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("fastJson", required: true)]
        public Input<string> FastJson { get; set; } = null!;

        /// <summary>
        /// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public FastApplicationArgs()
        {
        }
    }

    public sealed class FastApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A FAST application name.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("fastJson")]
        public Input<string>? FastJson { get; set; }

        /// <summary>
        /// Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// A FAST tenant name on which you want to manage application.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public FastApplicationState()
        {
        }
    }
}