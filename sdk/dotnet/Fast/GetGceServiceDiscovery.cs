// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Fast
{
    public static class GetGceServiceDiscovery
    {
        /// <summary>
        /// Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var tC3 = F5BigIP.Fast.GetGceServiceDiscovery.Invoke(new()
        ///     {
        ///         Region = "testgceregion",
        ///         TagKey = "testgcetag",
        ///         TagValue = "testgcevalue",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGceServiceDiscoveryResult> InvokeAsync(GetGceServiceDiscoveryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGceServiceDiscoveryResult>("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", args ?? new GetGceServiceDiscoveryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var tC3 = F5BigIP.Fast.GetGceServiceDiscovery.Invoke(new()
        ///     {
        ///         Region = "testgceregion",
        ///         TagKey = "testgcetag",
        ///         TagValue = "testgcevalue",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGceServiceDiscoveryResult> Invoke(GetGceServiceDiscoveryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGceServiceDiscoveryResult>("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", args ?? new GetGceServiceDiscoveryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGceServiceDiscoveryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public string? AddressRealm { get; set; }

        /// <summary>
        /// (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public bool? CredentialUpdate { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Base 64 encoded service account credentials JSON.
        /// </summary>
        [Input("encodedCredentials")]
        public string? EncodedCredentials { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public string? MinimumMonitors { get; set; }

        /// <summary>
        /// (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public int? Port { get; set; }

        /// <summary>
        /// (`optional`,type `string`)For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// (`Required`,type `string`) GCE region in which ADC is running.
        /// </summary>
        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        /// <summary>
        /// (`Required`,type `string`) The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey", required: true)]
        public string TagKey { get; set; } = null!;

        /// <summary>
        /// (`Required`,type `string`) The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue", required: true)]
        public string TagValue { get; set; } = null!;

        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
        /// </summary>
        [Input("undetectableAction")]
        public string? UndetectableAction { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Update interval for service discovery.
        /// </summary>
        [Input("updateInterval")]
        public string? UpdateInterval { get; set; }

        public GetGceServiceDiscoveryArgs()
        {
        }
        public static new GetGceServiceDiscoveryArgs Empty => new GetGceServiceDiscoveryArgs();
    }

    public sealed class GetGceServiceDiscoveryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public Input<string>? AddressRealm { get; set; }

        /// <summary>
        /// (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public Input<bool>? CredentialUpdate { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Base 64 encoded service account credentials JSON.
        /// </summary>
        [Input("encodedCredentials")]
        public Input<string>? EncodedCredentials { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public Input<string>? MinimumMonitors { get; set; }

        /// <summary>
        /// (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// (`optional`,type `string`)For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (`Required`,type `string`) GCE region in which ADC is running.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// (`Required`,type `string`) The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// (`Required`,type `string`) The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
        /// </summary>
        [Input("undetectableAction")]
        public Input<string>? UndetectableAction { get; set; }

        /// <summary>
        /// (`optional`,type `string`)Update interval for service discovery.
        /// </summary>
        [Input("updateInterval")]
        public Input<string>? UpdateInterval { get; set; }

        public GetGceServiceDiscoveryInvokeArgs()
        {
        }
        public static new GetGceServiceDiscoveryInvokeArgs Empty => new GetGceServiceDiscoveryInvokeArgs();
    }


    [OutputType]
    public sealed class GetGceServiceDiscoveryResult
    {
        public readonly string? AddressRealm;
        public readonly bool? CredentialUpdate;
        public readonly string? EncodedCredentials;
        /// <summary>
        /// The JSON for GCE service discovery block.
        /// </summary>
        public readonly string GceSdJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MinimumMonitors;
        public readonly int? Port;
        public readonly string? ProjectId;
        public readonly string Region;
        public readonly string TagKey;
        public readonly string TagValue;
        public readonly string? Type;
        public readonly string? UndetectableAction;
        public readonly string? UpdateInterval;

        [OutputConstructor]
        private GetGceServiceDiscoveryResult(
            string? addressRealm,

            bool? credentialUpdate,

            string? encodedCredentials,

            string gceSdJson,

            string id,

            string? minimumMonitors,

            int? port,

            string? projectId,

            string region,

            string tagKey,

            string tagValue,

            string? type,

            string? undetectableAction,

            string? updateInterval)
        {
            AddressRealm = addressRealm;
            CredentialUpdate = credentialUpdate;
            EncodedCredentials = encodedCredentials;
            GceSdJson = gceSdJson;
            Id = id;
            MinimumMonitors = minimumMonitors;
            Port = port;
            ProjectId = projectId;
            Region = region;
            TagKey = tagKey;
            TagValue = tagValue;
            Type = type;
            UndetectableAction = undetectableAction;
            UpdateInterval = updateInterval;
        }
    }
}
