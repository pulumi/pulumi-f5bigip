// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetWafEntityUrl
    {
        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to create JSON for WAF URL to later use with an existing WAF policy.
        /// 
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
        ///     var WAFURL1 = F5BigIP.Ssl.GetWafEntityUrl.Invoke(new()
        ///     {
        ///         Name = "/foobar",
        ///         Description = "this is a test",
        ///         Type = "explicit",
        ///         Protocol = "HTTP",
        ///         PerformStaging = true,
        ///         SignatureOverridesDisables = new[]
        ///         {
        ///             12345678,
        ///             87654321,
        ///         },
        ///         MethodOverrides = new[]
        ///         {
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlMethodOverrideInputArgs
        ///             {
        ///                 Allow = false,
        ///                 Method = "BCOPY",
        ///             },
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlMethodOverrideInputArgs
        ///             {
        ///                 Allow = true,
        ///                 Method = "BDELETE",
        ///             },
        ///         },
        ///         CrossOriginRequestsEnforcements = new[]
        ///         {
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs
        ///             {
        ///                 IncludeSubdomains = true,
        ///                 OriginName = "app1.com",
        ///                 OriginPort = "80",
        ///                 OriginProtocol = "http",
        ///             },
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs
        ///             {
        ///                 IncludeSubdomains = true,
        ///                 OriginName = "app2.com",
        ///                 OriginPort = "443",
        ///                 OriginProtocol = "http",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetWafEntityUrlResult> InvokeAsync(GetWafEntityUrlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWafEntityUrlResult>("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", args ?? new GetWafEntityUrlArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to create JSON for WAF URL to later use with an existing WAF policy.
        /// 
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
        ///     var WAFURL1 = F5BigIP.Ssl.GetWafEntityUrl.Invoke(new()
        ///     {
        ///         Name = "/foobar",
        ///         Description = "this is a test",
        ///         Type = "explicit",
        ///         Protocol = "HTTP",
        ///         PerformStaging = true,
        ///         SignatureOverridesDisables = new[]
        ///         {
        ///             12345678,
        ///             87654321,
        ///         },
        ///         MethodOverrides = new[]
        ///         {
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlMethodOverrideInputArgs
        ///             {
        ///                 Allow = false,
        ///                 Method = "BCOPY",
        ///             },
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlMethodOverrideInputArgs
        ///             {
        ///                 Allow = true,
        ///                 Method = "BDELETE",
        ///             },
        ///         },
        ///         CrossOriginRequestsEnforcements = new[]
        ///         {
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs
        ///             {
        ///                 IncludeSubdomains = true,
        ///                 OriginName = "app1.com",
        ///                 OriginPort = "80",
        ///                 OriginProtocol = "http",
        ///             },
        ///             new F5BigIP.Ssl.Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs
        ///             {
        ///                 IncludeSubdomains = true,
        ///                 OriginName = "app2.com",
        ///                 OriginPort = "443",
        ///                 OriginProtocol = "http",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWafEntityUrlResult> Invoke(GetWafEntityUrlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafEntityUrlResult>("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", args ?? new GetWafEntityUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWafEntityUrlArgs : global::Pulumi.InvokeArgs
    {
        [Input("crossOriginRequestsEnforcements")]
        private List<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementArgs>? _crossOriginRequestsEnforcements;

        /// <summary>
        /// A list of options that enables your web-application to share data with a website hosted on a
        /// different domain.
        /// </summary>
        public List<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementArgs> CrossOriginRequestsEnforcements
        {
            get => _crossOriginRequestsEnforcements ?? (_crossOriginRequestsEnforcements = new List<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementArgs>());
            set => _crossOriginRequestsEnforcements = value;
        }

        /// <summary>
        /// A description of the URL.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Select a Method for the URL to create an API endpoint. Default is : *.
        /// </summary>
        [Input("method")]
        public string? Method { get; set; }

        [Input("methodOverrides")]
        private List<Inputs.GetWafEntityUrlMethodOverrideArgs>? _methodOverrides;

        /// <summary>
        /// A list of methods that are allowed or disallowed for a specific URL.
        /// </summary>
        public List<Inputs.GetWafEntityUrlMethodOverrideArgs> MethodOverrides
        {
            get => _methodOverrides ?? (_methodOverrides = new List<Inputs.GetWafEntityUrlMethodOverrideArgs>());
            set => _methodOverrides = value;
        }

        /// <summary>
        /// WAF entity URL name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
        /// </summary>
        [Input("performStaging")]
        public bool? PerformStaging { get; set; }

        /// <summary>
        /// Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        [Input("signatureOverridesDisables")]
        private List<int>? _signatureOverridesDisables;

        /// <summary>
        /// List of Attack Signature Ids which are disabled for this particular URL.
        /// </summary>
        public List<int> SignatureOverridesDisables
        {
            get => _signatureOverridesDisables ?? (_signatureOverridesDisables = new List<int>());
            set => _signatureOverridesDisables = value;
        }

        /// <summary>
        /// Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetWafEntityUrlArgs()
        {
        }
        public static new GetWafEntityUrlArgs Empty => new GetWafEntityUrlArgs();
    }

    public sealed class GetWafEntityUrlInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("crossOriginRequestsEnforcements")]
        private InputList<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs>? _crossOriginRequestsEnforcements;

        /// <summary>
        /// A list of options that enables your web-application to share data with a website hosted on a
        /// different domain.
        /// </summary>
        public InputList<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs> CrossOriginRequestsEnforcements
        {
            get => _crossOriginRequestsEnforcements ?? (_crossOriginRequestsEnforcements = new InputList<Inputs.GetWafEntityUrlCrossOriginRequestsEnforcementInputArgs>());
            set => _crossOriginRequestsEnforcements = value;
        }

        /// <summary>
        /// A description of the URL.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Select a Method for the URL to create an API endpoint. Default is : *.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("methodOverrides")]
        private InputList<Inputs.GetWafEntityUrlMethodOverrideInputArgs>? _methodOverrides;

        /// <summary>
        /// A list of methods that are allowed or disallowed for a specific URL.
        /// </summary>
        public InputList<Inputs.GetWafEntityUrlMethodOverrideInputArgs> MethodOverrides
        {
            get => _methodOverrides ?? (_methodOverrides = new InputList<Inputs.GetWafEntityUrlMethodOverrideInputArgs>());
            set => _methodOverrides = value;
        }

        /// <summary>
        /// WAF entity URL name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
        /// </summary>
        [Input("performStaging")]
        public Input<bool>? PerformStaging { get; set; }

        /// <summary>
        /// Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("signatureOverridesDisables")]
        private InputList<int>? _signatureOverridesDisables;

        /// <summary>
        /// List of Attack Signature Ids which are disabled for this particular URL.
        /// </summary>
        public InputList<int> SignatureOverridesDisables
        {
            get => _signatureOverridesDisables ?? (_signatureOverridesDisables = new InputList<int>());
            set => _signatureOverridesDisables = value;
        }

        /// <summary>
        /// Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetWafEntityUrlInvokeArgs()
        {
        }
        public static new GetWafEntityUrlInvokeArgs Empty => new GetWafEntityUrlInvokeArgs();
    }


    [OutputType]
    public sealed class GetWafEntityUrlResult
    {
        public readonly ImmutableArray<Outputs.GetWafEntityUrlCrossOriginRequestsEnforcementResult> CrossOriginRequestsEnforcements;
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Json string representing created WAF entity URL declaration in JSON format
        /// </summary>
        public readonly string Json;
        public readonly string? Method;
        public readonly ImmutableArray<Outputs.GetWafEntityUrlMethodOverrideResult> MethodOverrides;
        public readonly string Name;
        public readonly bool? PerformStaging;
        public readonly string? Protocol;
        public readonly ImmutableArray<int> SignatureOverridesDisables;
        public readonly string? Type;

        [OutputConstructor]
        private GetWafEntityUrlResult(
            ImmutableArray<Outputs.GetWafEntityUrlCrossOriginRequestsEnforcementResult> crossOriginRequestsEnforcements,

            string? description,

            string id,

            string json,

            string? method,

            ImmutableArray<Outputs.GetWafEntityUrlMethodOverrideResult> methodOverrides,

            string name,

            bool? performStaging,

            string? protocol,

            ImmutableArray<int> signatureOverridesDisables,

            string? type)
        {
            CrossOriginRequestsEnforcements = crossOriginRequestsEnforcements;
            Description = description;
            Id = id;
            Json = json;
            Method = method;
            MethodOverrides = methodOverrides;
            Name = name;
            PerformStaging = performStaging;
            Protocol = protocol;
            SignatureOverridesDisables = signatureOverridesDisables;
            Type = type;
        }
    }
}
