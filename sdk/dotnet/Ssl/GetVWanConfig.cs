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
        /// <summary>
        /// Use this data source (`f5bigip.ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
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
        ///     var vwanconfig = F5BigIP.Ssl.GetVWanConfig.Invoke(new()
        ///     {
        ///         AzureVwanResourcegroup = "azurevwan-bigip-rg-9c8d",
        ///         AzureVwanName = "azurevwan-bigip-vwan-9c8d",
        ///         AzureVwanVpnsite = "azurevwan-bigip-vsite-9c8d",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Pre-required Environment Settings:
        /// 
        /// * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
        /// 
        /// * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
        /// 
        /// * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
        /// 
        /// * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
        /// 
        /// * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
        /// 
        /// * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
        /// </summary>
        public static Task<GetVWanConfigResult> InvokeAsync(GetVWanConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVWanConfigResult>("f5bigip:ssl/getVWanConfig:getVWanConfig", args ?? new GetVWanConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
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
        ///     var vwanconfig = F5BigIP.Ssl.GetVWanConfig.Invoke(new()
        ///     {
        ///         AzureVwanResourcegroup = "azurevwan-bigip-rg-9c8d",
        ///         AzureVwanName = "azurevwan-bigip-vwan-9c8d",
        ///         AzureVwanVpnsite = "azurevwan-bigip-vsite-9c8d",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Pre-required Environment Settings:
        /// 
        /// * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
        /// 
        /// * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
        /// 
        /// * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
        /// 
        /// * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
        /// 
        /// * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
        /// 
        /// * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
        /// </summary>
        public static Output<GetVWanConfigResult> Invoke(GetVWanConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVWanConfigResult>("f5bigip:ssl/getVWanConfig:getVWanConfig", args ?? new GetVWanConfigInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
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
        ///     var vwanconfig = F5BigIP.Ssl.GetVWanConfig.Invoke(new()
        ///     {
        ///         AzureVwanResourcegroup = "azurevwan-bigip-rg-9c8d",
        ///         AzureVwanName = "azurevwan-bigip-vwan-9c8d",
        ///         AzureVwanVpnsite = "azurevwan-bigip-vsite-9c8d",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Pre-required Environment Settings:
        /// 
        /// * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
        /// 
        /// * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
        /// 
        /// * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
        /// 
        /// * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
        /// 
        /// * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
        /// 
        /// * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
        /// </summary>
        public static Output<GetVWanConfigResult> Invoke(GetVWanConfigInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVWanConfigResult>("f5bigip:ssl/getVWanConfig:getVWanConfig", args ?? new GetVWanConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVWanConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure vWAN Name
        /// </summary>
        [Input("azureVwanName", required: true)]
        public string AzureVwanName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure vWAN resource group
        /// </summary>
        [Input("azureVwanResourcegroup", required: true)]
        public string AzureVwanResourcegroup { get; set; } = null!;

        /// <summary>
        /// Name of the Azure vWAN VPN site from which configuration to be download
        /// </summary>
        [Input("azureVwanVpnsite", required: true)]
        public string AzureVwanVpnsite { get; set; } = null!;

        public GetVWanConfigArgs()
        {
        }
        public static new GetVWanConfigArgs Empty => new GetVWanConfigArgs();
    }

    public sealed class GetVWanConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure vWAN Name
        /// </summary>
        [Input("azureVwanName", required: true)]
        public Input<string> AzureVwanName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure vWAN resource group
        /// </summary>
        [Input("azureVwanResourcegroup", required: true)]
        public Input<string> AzureVwanResourcegroup { get; set; } = null!;

        /// <summary>
        /// Name of the Azure vWAN VPN site from which configuration to be download
        /// </summary>
        [Input("azureVwanVpnsite", required: true)]
        public Input<string> AzureVwanVpnsite { get; set; } = null!;

        public GetVWanConfigInvokeArgs()
        {
        }
        public static new GetVWanConfigInvokeArgs Empty => new GetVWanConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetVWanConfigResult
    {
        public readonly string AzureVwanName;
        public readonly string AzureVwanResourcegroup;
        public readonly string AzureVwanVpnsite;
        /// <summary>
        /// (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
        /// </summary>
        public readonly string BigipGwIp;
        /// <summary>
        /// (type `string`) Provides IP Address space used on vWAN Hub.
        /// </summary>
        public readonly string HubAddressSpace;
        /// <summary>
        /// (type `list`) Provides Subnets connected to vWAN Hub.
        /// </summary>
        public readonly ImmutableArray<string> HubConnectedSubnets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
        /// </summary>
        public readonly string PresharedKey;
        /// <summary>
        /// (type `list`) Provides vWAN Gateway Address for IPSec End point
        /// </summary>
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
