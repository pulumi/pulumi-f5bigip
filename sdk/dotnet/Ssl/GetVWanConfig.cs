// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetVWanConfig
    {
        public static Task<GetVWanConfigResult> InvokeAsync(GetVWanConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVWanConfigResult>("f5bigip:ssl/getVWanConfig:getVWanConfig", args ?? new GetVWanConfigArgs(), options.WithVersion());
    }


    public sealed class GetVWanConfigArgs : Pulumi.InvokeArgs
    {
        [Input("azureVwanName", required: true)]
        public string AzureVwanName { get; set; } = null!;

        [Input("azureVwanResourcegroup", required: true)]
        public string AzureVwanResourcegroup { get; set; } = null!;

        [Input("azureVwanVpnsite", required: true)]
        public string AzureVwanVpnsite { get; set; } = null!;

        public GetVWanConfigArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVWanConfigResult
    {
        public readonly string AzureVwanName;
        public readonly string AzureVwanResourcegroup;
        public readonly string AzureVwanVpnsite;
        public readonly string BigipGwIp;
        public readonly string HubAddressSpace;
        public readonly ImmutableArray<string> HubConnectedSubnets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string PresharedKey;
        public readonly ImmutableArray<string> VwanGwAddresses;

        [OutputConstructor]
        private GetVWanConfigResult(
            string azureVwanName,

            string azureVwanResourcegroup,

            string azureVwanVpnsite,

            string bigipGwIp,

            string hubAddressSpace,

            ImmutableArray<string> hubConnectedSubnets,

            string id,

            string presharedKey,

            ImmutableArray<string> vwanGwAddresses)
        {
            AzureVwanName = azureVwanName;
            AzureVwanResourcegroup = azureVwanResourcegroup;
            AzureVwanVpnsite = azureVwanVpnsite;
            BigipGwIp = bigipGwIp;
            HubAddressSpace = hubAddressSpace;
            HubConnectedSubnets = hubConnectedSubnets;
            Id = id;
            PresharedKey = presharedKey;
            VwanGwAddresses = vwanGwAddresses;
        }
    }
}
