// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    [F5BigIPResourceType("f5bigip:index/do:Do")]
    public partial class Do : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP Address of BIGIP Host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Output("bigipAddress")]
        public Output<string?> BigipAddress { get; private set; } = null!;

        /// <summary>
        /// Password of BIGIP host to be used for this resource
        /// </summary>
        [Output("bigipPassword")]
        public Output<string?> BigipPassword { get; private set; } = null!;

        /// <summary>
        /// Port number of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Output("bigipPort")]
        public Output<string?> BigipPort { get; private set; } = null!;

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        [Output("bigipTokenAuth")]
        public Output<bool?> BigipTokenAuth { get; private set; } = null!;

        /// <summary>
        /// UserName of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Output("bigipUser")]
        public Output<string?> BigipUser { get; private set; } = null!;

        /// <summary>
        /// Name of the of the Declarative DO JSON file
        /// </summary>
        [Output("doJson")]
        public Output<string> DoJson { get; private set; } = null!;

        /// <summary>
        /// unique identifier for DO resource
        /// </summary>
        [Output("tenantName")]
        public Output<string?> TenantName { get; private set; } = null!;

        /// <summary>
        /// DO json
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a Do resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Do(string name, DoArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/do:Do", name, args ?? new DoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Do(string name, Input<string> id, DoState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/do:Do", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "bigipPassword",
                    "bigipTokenAuth",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Do resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Do Get(string name, Input<string> id, DoState? state = null, CustomResourceOptions? options = null)
        {
            return new Do(name, id, state, options);
        }
    }

    public sealed class DoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address of BIGIP Host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipAddress")]
        public Input<string>? BigipAddress { get; set; }

        [Input("bigipPassword")]
        private Input<string>? _bigipPassword;

        /// <summary>
        /// Password of BIGIP host to be used for this resource
        /// </summary>
        public Input<string>? BigipPassword
        {
            get => _bigipPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigipPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipPort")]
        public Input<string>? BigipPort { get; set; }

        [Input("bigipTokenAuth")]
        private Input<bool>? _bigipTokenAuth;

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        public Input<bool>? BigipTokenAuth
        {
            get => _bigipTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigipTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// UserName of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipUser")]
        public Input<string>? BigipUser { get; set; }

        /// <summary>
        /// Name of the of the Declarative DO JSON file
        /// </summary>
        [Input("doJson", required: true)]
        public Input<string> DoJson { get; set; } = null!;

        /// <summary>
        /// unique identifier for DO resource
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        /// <summary>
        /// DO json
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public DoArgs()
        {
        }
        public static new DoArgs Empty => new DoArgs();
    }

    public sealed class DoState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address of BIGIP Host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipAddress")]
        public Input<string>? BigipAddress { get; set; }

        [Input("bigipPassword")]
        private Input<string>? _bigipPassword;

        /// <summary>
        /// Password of BIGIP host to be used for this resource
        /// </summary>
        public Input<string>? BigipPassword
        {
            get => _bigipPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigipPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipPort")]
        public Input<string>? BigipPort { get; set; }

        [Input("bigipTokenAuth")]
        private Input<bool>? _bigipTokenAuth;

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        public Input<bool>? BigipTokenAuth
        {
            get => _bigipTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigipTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// UserName of BIGIP host to be used for this resource,this is optional parameter.
        /// whenever we specify this parameter it gets overwrite provider configuration
        /// </summary>
        [Input("bigipUser")]
        public Input<string>? BigipUser { get; set; }

        /// <summary>
        /// Name of the of the Declarative DO JSON file
        /// </summary>
        [Input("doJson")]
        public Input<string>? DoJson { get; set; }

        /// <summary>
        /// unique identifier for DO resource
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        /// <summary>
        /// DO json
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public DoState()
        {
        }
        public static new DoState Empty => new DoState();
    }
}
