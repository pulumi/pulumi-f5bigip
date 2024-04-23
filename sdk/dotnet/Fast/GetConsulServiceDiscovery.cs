// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Fast
{
    public static class GetConsulServiceDiscovery
    {
        /// <summary>
        /// Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
        ///     var TC2 = F5BigIP.Fast.GetConsulServiceDiscovery.Invoke(new()
        ///     {
        ///         Uri = "https://192.0.2.100:8500/v1/catalog/nodes",
        ///         Port = 8080,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetConsulServiceDiscoveryResult> InvokeAsync(GetConsulServiceDiscoveryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConsulServiceDiscoveryResult>("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", args ?? new GetConsulServiceDiscoveryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
        ///     var TC2 = F5BigIP.Fast.GetConsulServiceDiscovery.Invoke(new()
        ///     {
        ///         Uri = "https://192.0.2.100:8500/v1/catalog/nodes",
        ///         Port = 8080,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetConsulServiceDiscoveryResult> Invoke(GetConsulServiceDiscoveryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConsulServiceDiscoveryResult>("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", args ?? new GetConsulServiceDiscoveryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConsulServiceDiscoveryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public string? AddressRealm { get; set; }

        /// <summary>
        /// Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public bool? CredentialUpdate { get; set; }

        /// <summary>
        /// Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
        /// </summary>
        [Input("encodedToken")]
        public string? EncodedToken { get; set; }

        /// <summary>
        /// Custom JMESPath Query.
        /// </summary>
        [Input("jmesPathQuery")]
        public string? JmesPathQuery { get; set; }

        /// <summary>
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public string? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port", required: true)]
        public int Port { get; set; }

        /// <summary>
        /// If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
        /// </summary>
        [Input("rejectUnauthorized")]
        public bool? RejectUnauthorized { get; set; }

        /// <summary>
        /// CA Bundle to validate server certificates.
        /// </summary>
        [Input("trustCa")]
        public string? TrustCa { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Action to take when node cannot be detected,default `remove`.
        /// </summary>
        [Input("undetectableAction")]
        public string? UndetectableAction { get; set; }

        /// <summary>
        /// Update interval for service discovery.
        /// </summary>
        [Input("updateInterval")]
        public string? UpdateInterval { get; set; }

        /// <summary>
        /// The location of the node data.
        /// </summary>
        [Input("uri", required: true)]
        public string Uri { get; set; } = null!;

        public GetConsulServiceDiscoveryArgs()
        {
        }
        public static new GetConsulServiceDiscoveryArgs Empty => new GetConsulServiceDiscoveryArgs();
    }

    public sealed class GetConsulServiceDiscoveryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public Input<string>? AddressRealm { get; set; }

        /// <summary>
        /// Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public Input<bool>? CredentialUpdate { get; set; }

        /// <summary>
        /// Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
        /// </summary>
        [Input("encodedToken")]
        public Input<string>? EncodedToken { get; set; }

        /// <summary>
        /// Custom JMESPath Query.
        /// </summary>
        [Input("jmesPathQuery")]
        public Input<string>? JmesPathQuery { get; set; }

        /// <summary>
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public Input<string>? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
        /// </summary>
        [Input("rejectUnauthorized")]
        public Input<bool>? RejectUnauthorized { get; set; }

        /// <summary>
        /// CA Bundle to validate server certificates.
        /// </summary>
        [Input("trustCa")]
        public Input<string>? TrustCa { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Action to take when node cannot be detected,default `remove`.
        /// </summary>
        [Input("undetectableAction")]
        public Input<string>? UndetectableAction { get; set; }

        /// <summary>
        /// Update interval for service discovery.
        /// </summary>
        [Input("updateInterval")]
        public Input<string>? UpdateInterval { get; set; }

        /// <summary>
        /// The location of the node data.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public GetConsulServiceDiscoveryInvokeArgs()
        {
        }
        public static new GetConsulServiceDiscoveryInvokeArgs Empty => new GetConsulServiceDiscoveryInvokeArgs();
    }


    [OutputType]
    public sealed class GetConsulServiceDiscoveryResult
    {
        public readonly string? AddressRealm;
        public readonly string ConsulSdJson;
        public readonly bool? CredentialUpdate;
        public readonly string? EncodedToken;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? JmesPathQuery;
        public readonly string? MinimumMonitors;
        public readonly int Port;
        public readonly bool? RejectUnauthorized;
        public readonly string? TrustCa;
        public readonly string? Type;
        public readonly string? UndetectableAction;
        public readonly string? UpdateInterval;
        public readonly string Uri;

        [OutputConstructor]
        private GetConsulServiceDiscoveryResult(
            string? addressRealm,

            string consulSdJson,

            bool? credentialUpdate,

            string? encodedToken,

            string id,

            string? jmesPathQuery,

            string? minimumMonitors,

            int port,

            bool? rejectUnauthorized,

            string? trustCa,

            string? type,

            string? undetectableAction,

            string? updateInterval,

            string uri)
        {
            AddressRealm = addressRealm;
            ConsulSdJson = consulSdJson;
            CredentialUpdate = credentialUpdate;
            EncodedToken = encodedToken;
            Id = id;
            JmesPathQuery = jmesPathQuery;
            MinimumMonitors = minimumMonitors;
            Port = port;
            RejectUnauthorized = rejectUnauthorized;
            TrustCa = trustCa;
            Type = type;
            UndetectableAction = undetectableAction;
            UpdateInterval = updateInterval;
            Uri = uri;
        }
    }
}
