// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Fast
{
    public static class GetAzureServiceDiscovery
    {
        /// <summary>
        /// Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
        /// </summary>
        public static Task<GetAzureServiceDiscoveryResult> InvokeAsync(GetAzureServiceDiscoveryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAzureServiceDiscoveryResult>("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", args ?? new GetAzureServiceDiscoveryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
        /// </summary>
        public static Output<GetAzureServiceDiscoveryResult> Invoke(GetAzureServiceDiscoveryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureServiceDiscoveryResult>("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", args ?? new GetAzureServiceDiscoveryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAzureServiceDiscoveryArgs : global::Pulumi.InvokeArgs
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
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public string? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for Azure service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public int? Port { get; set; }

        /// <summary>
        /// Azure Resource Group name.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public string ResourceGroup { get; set; } = null!;

        /// <summary>
        /// Azure subscription ID.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey")]
        public string? TagKey { get; set; }

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue")]
        public string? TagValue { get; set; }

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

        public GetAzureServiceDiscoveryArgs()
        {
        }
        public static new GetAzureServiceDiscoveryArgs Empty => new GetAzureServiceDiscoveryArgs();
    }

    public sealed class GetAzureServiceDiscoveryInvokeArgs : global::Pulumi.InvokeArgs
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
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public Input<string>? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for Azure service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Azure Resource Group name.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        /// <summary>
        /// Azure subscription ID.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

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

        public GetAzureServiceDiscoveryInvokeArgs()
        {
        }
        public static new GetAzureServiceDiscoveryInvokeArgs Empty => new GetAzureServiceDiscoveryInvokeArgs();
    }


    [OutputType]
    public sealed class GetAzureServiceDiscoveryResult
    {
        public readonly string? AddressRealm;
        /// <summary>
        /// The JSON for Azure service discovery block.
        /// </summary>
        public readonly string AzureSdJson;
        public readonly bool? CredentialUpdate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MinimumMonitors;
        public readonly int? Port;
        public readonly string ResourceGroup;
        public readonly string SubscriptionId;
        public readonly string? TagKey;
        public readonly string? TagValue;
        public readonly string? Type;
        public readonly string? UndetectableAction;
        public readonly string? UpdateInterval;

        [OutputConstructor]
        private GetAzureServiceDiscoveryResult(
            string? addressRealm,

            string azureSdJson,

            bool? credentialUpdate,

            string id,

            string? minimumMonitors,

            int? port,

            string resourceGroup,

            string subscriptionId,

            string? tagKey,

            string? tagValue,

            string? type,

            string? undetectableAction,

            string? updateInterval)
        {
            AddressRealm = addressRealm;
            AzureSdJson = azureSdJson;
            CredentialUpdate = credentialUpdate;
            Id = id;
            MinimumMonitors = minimumMonitors;
            Port = port;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
            TagKey = tagKey;
            TagValue = tagValue;
            Type = type;
            UndetectableAction = undetectableAction;
            UpdateInterval = updateInterval;
        }
    }
}
