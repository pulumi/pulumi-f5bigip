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
    /// `f5bigip.BigIqAs3` provides details about bigiq as3 resource
    /// 
    /// This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Example Usage for json file
    ///     var exampletask = new F5BigIP.BigIqAs3("exampletask", new()
    ///     {
    ///         As3Json = File.ReadAllText("bigiq_example.json"),
    ///         BigiqAddress = "xx.xx.xxx.xx",
    ///         BigiqPassword = "xxxxxxxxx",
    ///         BigiqUser = "xxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/bigIqAs3:BigIqAs3")]
    public partial class BigIqAs3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Output("as3Json")]
        public Output<string> As3Json { get; private set; } = null!;

        /// <summary>
        /// Address of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        [Output("bigiqAddress")]
        public Output<string> BigiqAddress { get; private set; } = null!;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        [Output("bigiqLoginRef")]
        public Output<string?> BigiqLoginRef { get; private set; } = null!;

        /// <summary>
        /// Password of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        [Output("bigiqPassword")]
        public Output<string> BigiqPassword { get; private set; } = null!;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        [Output("bigiqPort")]
        public Output<string?> BigiqPort { get; private set; } = null!;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        [Output("bigiqTokenAuth")]
        public Output<bool?> BigiqTokenAuth { get; private set; } = null!;

        /// <summary>
        /// User name  of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        [Output("bigiqUser")]
        public Output<string> BigiqUser { get; private set; } = null!;

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to `true`
        /// 
        /// * `bigiq_example.json` - Example  AS3 Declarative JSON file
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// 
        /// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
        /// 
        /// &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        /// </summary>
        [Output("ignoreMetadata")]
        public Output<bool?> IgnoreMetadata { get; private set; } = null!;

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
                AdditionalSecretOutputs =
                {
                    "bigiqLoginRef",
                    "bigiqPassword",
                    "bigiqPort",
                    "bigiqTokenAuth",
                    "bigiqUser",
                },
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

    public sealed class BigIqAs3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("as3Json", required: true)]
        public Input<string> As3Json { get; set; } = null!;

        /// <summary>
        /// Address of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        [Input("bigiqAddress", required: true)]
        public Input<string> BigiqAddress { get; set; } = null!;

        [Input("bigiqLoginRef")]
        private Input<string>? _bigiqLoginRef;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        public Input<string>? BigiqLoginRef
        {
            get => _bigiqLoginRef;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqLoginRef = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPassword", required: true)]
        private Input<string>? _bigiqPassword;

        /// <summary>
        /// Password of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        public Input<string>? BigiqPassword
        {
            get => _bigiqPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPort")]
        private Input<string>? _bigiqPort;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        public Input<string>? BigiqPort
        {
            get => _bigiqPort;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPort = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqTokenAuth")]
        private Input<bool>? _bigiqTokenAuth;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        public Input<bool>? BigiqTokenAuth
        {
            get => _bigiqTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqUser", required: true)]
        private Input<string>? _bigiqUser;

        /// <summary>
        /// User name  of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        public Input<string>? BigiqUser
        {
            get => _bigiqUser;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqUser = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to `true`
        /// 
        /// * `bigiq_example.json` - Example  AS3 Declarative JSON file
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// 
        /// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
        /// 
        /// &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        /// </summary>
        [Input("ignoreMetadata")]
        public Input<bool>? IgnoreMetadata { get; set; }

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        public BigIqAs3Args()
        {
        }
        public static new BigIqAs3Args Empty => new BigIqAs3Args();
    }

    public sealed class BigIqAs3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("as3Json")]
        public Input<string>? As3Json { get; set; }

        /// <summary>
        /// Address of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        [Input("bigiqAddress")]
        public Input<string>? BigiqAddress { get; set; }

        [Input("bigiqLoginRef")]
        private Input<string>? _bigiqLoginRef;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        public Input<string>? BigiqLoginRef
        {
            get => _bigiqLoginRef;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqLoginRef = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPassword")]
        private Input<string>? _bigiqPassword;

        /// <summary>
        /// Password of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        public Input<string>? BigiqPassword
        {
            get => _bigiqPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPort")]
        private Input<string>? _bigiqPort;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        public Input<string>? BigiqPort
        {
            get => _bigiqPort;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPort = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqTokenAuth")]
        private Input<bool>? _bigiqTokenAuth;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        public Input<bool>? BigiqTokenAuth
        {
            get => _bigiqTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqUser")]
        private Input<string>? _bigiqUser;

        /// <summary>
        /// User name  of the BIG-IQ to which your targer BIG-IP is attached
        /// </summary>
        public Input<string>? BigiqUser
        {
            get => _bigiqUser;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqUser = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to `true`
        /// 
        /// * `bigiq_example.json` - Example  AS3 Declarative JSON file
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// 
        /// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
        /// 
        /// &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        /// </summary>
        [Input("ignoreMetadata")]
        public Input<bool>? IgnoreMetadata { get; set; }

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        public BigIqAs3State()
        {
        }
        public static new BigIqAs3State Empty => new BigIqAs3State();
    }
}
