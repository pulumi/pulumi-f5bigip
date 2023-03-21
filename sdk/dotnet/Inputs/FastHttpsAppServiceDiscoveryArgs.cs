// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class FastHttpsAppServiceDiscoveryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to look for public or private IP addresses. Default :`private`
        /// </summary>
        [Input("sdAddressRealm")]
        public Input<string>? SdAddressRealm { get; set; }

        /// <summary>
        /// Information for discovering AWS nodes that are not in the same region as your BIG-IP.
        /// </summary>
        [Input("sdAwsAccessKey")]
        public Input<string>? SdAwsAccessKey { get; set; }

        /// <summary>
        /// Empty string (default) means region in which ADC is running.
        /// </summary>
        [Input("sdAwsRegion")]
        public Input<string>? SdAwsRegion { get; set; }

        /// <summary>
        /// Will be stored in the declaration as an encrypted string.
        /// </summary>
        [Input("sdAwsSecretAccessKey")]
        public Input<string>? SdAwsSecretAccessKey { get; set; }

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("sdAwsTagKey")]
        public Input<string>? SdAwsTagKey { get; set; }

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("sdAwsTagVal")]
        public Input<string>? SdAwsTagVal { get; set; }

        /// <summary>
        /// Azure Active Directory ID (AKA tenant ID).
        /// </summary>
        [Input("sdAzureDirectoryId")]
        public Input<string>? SdAzureDirectoryId { get; set; }

        /// <summary>
        /// Azure Resource Group name.
        /// </summary>
        [Input("sdAzureResourceGroup")]
        public Input<string>? SdAzureResourceGroup { get; set; }

        /// <summary>
        /// ID of resource to find nodes by.
        /// </summary>
        [Input("sdAzureResourceId")]
        public Input<string>? SdAzureResourceId { get; set; }

        /// <summary>
        /// Azure subscription ID.
        /// </summary>
        [Input("sdAzureSubscriptionId")]
        public Input<string>? SdAzureSubscriptionId { get; set; }

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("sdAzureTagKey")]
        public Input<string>? SdAzureTagKey { get; set; }

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("sdAzureTagVal")]
        public Input<string>? SdAzureTagVal { get; set; }

        /// <summary>
        /// Empty string (default) means region in which ADC is running.
        /// </summary>
        [Input("sdGceRegion")]
        public Input<string>? SdGceRegion { get; set; }

        /// <summary>
        /// The tag key associated with the node to add to this pool
        /// </summary>
        [Input("sdGceTagKey")]
        public Input<string>? SdGceTagKey { get; set; }

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("sdGceTagVal")]
        public Input<string>? SdGceTagVal { get; set; }

        /// <summary>
        /// port number of serviceport to be used for FAST-Generated Pool.
        /// </summary>
        [Input("sdPort", required: true)]
        public Input<int> SdPort { get; set; } = null!;

        /// <summary>
        /// service discovery account type, options [`aws`,`azure`,`gce`]
        /// </summary>
        [Input("sdType", required: true)]
        public Input<string> SdType { get; set; } = null!;

        /// <summary>
        /// Action to take when node cannot be detected. Default `remove`.
        /// </summary>
        [Input("sdUndetectableAction")]
        public Input<string>? SdUndetectableAction { get; set; }

        public FastHttpsAppServiceDiscoveryArgs()
        {
        }
        public static new FastHttpsAppServiceDiscoveryArgs Empty => new FastHttpsAppServiceDiscoveryArgs();
    }
}
