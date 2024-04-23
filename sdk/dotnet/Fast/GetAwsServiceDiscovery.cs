// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Fast
{
    public static class GetAwsServiceDiscovery
    {
        /// <summary>
        /// Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
        ///     var TC2 = F5BigIP.Fast.GetAwsServiceDiscovery.Invoke(new()
        ///     {
        ///         TagKey = "testawstagkey",
        ///         TagValue = "testawstagvalue",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAwsServiceDiscoveryResult> InvokeAsync(GetAwsServiceDiscoveryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAwsServiceDiscoveryResult>("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", args ?? new GetAwsServiceDiscoveryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
        ///     var TC2 = F5BigIP.Fast.GetAwsServiceDiscovery.Invoke(new()
        ///     {
        ///         TagKey = "testawstagkey",
        ///         TagValue = "testawstagvalue",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAwsServiceDiscoveryResult> Invoke(GetAwsServiceDiscoveryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsServiceDiscoveryResult>("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", args ?? new GetAwsServiceDiscoveryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsServiceDiscoveryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public string? AddressRealm { get; set; }

        [Input("awsAccessKey")]
        private string? _awsAccessKey;

        /// <summary>
        /// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
        /// </summary>
        public string? AwsAccessKey
        {
            get => _awsAccessKey;
            set => _awsAccessKey = value;
        }

        /// <summary>
        /// AWS region in which ADC is running,default Empty string.
        /// </summary>
        [Input("awsRegion")]
        public string? AwsRegion { get; set; }

        [Input("awsSecretAccessKey")]
        private string? _awsSecretAccessKey;

        /// <summary>
        /// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
        /// </summary>
        public string? AwsSecretAccessKey
        {
            get => _awsSecretAccessKey;
            set => _awsSecretAccessKey = value;
        }

        /// <summary>
        /// Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public bool? CredentialUpdate { get; set; }

        /// <summary>
        /// AWS externalID field.
        /// </summary>
        [Input("externalId")]
        public string? ExternalId { get; set; }

        /// <summary>
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public string? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public int? Port { get; set; }

        /// <summary>
        /// Assume a role (also requires the `external_id` field)
        /// </summary>
        [Input("roleArn")]
        public string? RoleArn { get; set; }

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey", required: true)]
        public string TagKey { get; set; } = null!;

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue", required: true)]
        public string TagValue { get; set; } = null!;

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

        public GetAwsServiceDiscoveryArgs()
        {
        }
        public static new GetAwsServiceDiscoveryArgs Empty => new GetAwsServiceDiscoveryArgs();
    }

    public sealed class GetAwsServiceDiscoveryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies whether to look for public or private IP addresses,default `private`.
        /// </summary>
        [Input("addressRealm")]
        public Input<string>? AddressRealm { get; set; }

        [Input("awsAccessKey")]
        private Input<string>? _awsAccessKey;

        /// <summary>
        /// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
        /// </summary>
        public Input<string>? AwsAccessKey
        {
            get => _awsAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _awsAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// AWS region in which ADC is running,default Empty string.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        [Input("awsSecretAccessKey")]
        private Input<string>? _awsSecretAccessKey;

        /// <summary>
        /// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
        /// </summary>
        public Input<string>? AwsSecretAccessKey
        {
            get => _awsSecretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _awsSecretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies whether you are updating your credentials,default `false`.
        /// </summary>
        [Input("credentialUpdate")]
        public Input<bool>? CredentialUpdate { get; set; }

        /// <summary>
        /// AWS externalID field.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Member is down when fewer than minimum monitors report it healthy.
        /// </summary>
        [Input("minimumMonitors")]
        public Input<string>? MinimumMonitors { get; set; }

        /// <summary>
        /// Port to be used for AWS service discovery,default `80`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Assume a role (also requires the `external_id` field)
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The tag key associated with the node to add to this pool.
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// The tag value associated with the node to add to this pool.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

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

        public GetAwsServiceDiscoveryInvokeArgs()
        {
        }
        public static new GetAwsServiceDiscoveryInvokeArgs Empty => new GetAwsServiceDiscoveryInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsServiceDiscoveryResult
    {
        public readonly string? AddressRealm;
        public readonly string? AwsAccessKey;
        public readonly string AwsRegion;
        /// <summary>
        /// The JSON for AWS service discovery block.
        /// </summary>
        public readonly string AwsSdJson;
        public readonly string? AwsSecretAccessKey;
        public readonly bool? CredentialUpdate;
        public readonly string ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MinimumMonitors;
        public readonly int? Port;
        public readonly string RoleArn;
        public readonly string TagKey;
        public readonly string TagValue;
        public readonly string? Type;
        public readonly string? UndetectableAction;
        public readonly string? UpdateInterval;

        [OutputConstructor]
        private GetAwsServiceDiscoveryResult(
            string? addressRealm,

            string? awsAccessKey,

            string awsRegion,

            string awsSdJson,

            string? awsSecretAccessKey,

            bool? credentialUpdate,

            string externalId,

            string id,

            string? minimumMonitors,

            int? port,

            string roleArn,

            string tagKey,

            string tagValue,

            string? type,

            string? undetectableAction,

            string? updateInterval)
        {
            AddressRealm = addressRealm;
            AwsAccessKey = awsAccessKey;
            AwsRegion = awsRegion;
            AwsSdJson = awsSdJson;
            AwsSecretAccessKey = awsSecretAccessKey;
            CredentialUpdate = credentialUpdate;
            ExternalId = externalId;
            Id = id;
            MinimumMonitors = minimumMonitors;
            Port = port;
            RoleArn = roleArn;
            TagKey = tagKey;
            TagValue = tagValue;
            Type = type;
            UndetectableAction = undetectableAction;
            UpdateInterval = updateInterval;
        }
    }
}
