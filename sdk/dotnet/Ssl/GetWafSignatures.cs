// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetWafSignatures
    {
        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafSignatures`) to get the details of attack signatures available on BIG-IP WAF
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var wAFSIG1 = F5BigIP.Ssl.GetWafSignatures.Invoke(new()
        ///     {
        ///         SignatureId = 200104004,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWafSignaturesResult> InvokeAsync(GetWafSignaturesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWafSignaturesResult>("f5bigip:ssl/getWafSignatures:getWafSignatures", args ?? new GetWafSignaturesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafSignatures`) to get the details of attack signatures available on BIG-IP WAF
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var wAFSIG1 = F5BigIP.Ssl.GetWafSignatures.Invoke(new()
        ///     {
        ///         SignatureId = 200104004,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWafSignaturesResult> Invoke(GetWafSignaturesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWafSignaturesResult>("f5bigip:ssl/getWafSignatures:getWafSignatures", args ?? new GetWafSignaturesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWafSignaturesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The relative detection accuracy of the signature.
        /// </summary>
        [Input("accuracy")]
        public string? Accuracy { get; set; }

        /// <summary>
        /// Description of the signature.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        [Input("enabled")]
        public bool? Enabled { get; set; }

        /// <summary>
        /// Name of the signature as configured on the system.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("performStaging")]
        public bool? PerformStaging { get; set; }

        /// <summary>
        /// The relative risk level of the attack that matches this signature.
        /// </summary>
        [Input("risk")]
        public string? Risk { get; set; }

        /// <summary>
        /// ID of the signature in the BIG-IP WAF database.
        /// </summary>
        [Input("signatureId", required: true)]
        public int SignatureId { get; set; }

        /// <summary>
        /// System generated ID of the signature.
        /// </summary>
        [Input("systemSignatureId")]
        public string? SystemSignatureId { get; set; }

        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// Type of the signature.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetWafSignaturesArgs()
        {
        }
        public static new GetWafSignaturesArgs Empty => new GetWafSignaturesArgs();
    }

    public sealed class GetWafSignaturesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The relative detection accuracy of the signature.
        /// </summary>
        [Input("accuracy")]
        public Input<string>? Accuracy { get; set; }

        /// <summary>
        /// Description of the signature.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the signature as configured on the system.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("performStaging")]
        public Input<bool>? PerformStaging { get; set; }

        /// <summary>
        /// The relative risk level of the attack that matches this signature.
        /// </summary>
        [Input("risk")]
        public Input<string>? Risk { get; set; }

        /// <summary>
        /// ID of the signature in the BIG-IP WAF database.
        /// </summary>
        [Input("signatureId", required: true)]
        public Input<int> SignatureId { get; set; } = null!;

        /// <summary>
        /// System generated ID of the signature.
        /// </summary>
        [Input("systemSignatureId")]
        public Input<string>? SystemSignatureId { get; set; }

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Type of the signature.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetWafSignaturesInvokeArgs()
        {
        }
        public static new GetWafSignaturesInvokeArgs Empty => new GetWafSignaturesInvokeArgs();
    }


    [OutputType]
    public sealed class GetWafSignaturesResult
    {
        /// <summary>
        /// The relative detection accuracy of the signature.
        /// </summary>
        public readonly string Accuracy;
        /// <summary>
        /// Description of the signature.
        /// </summary>
        public readonly string? Description;
        public readonly bool? Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        /// <summary>
        /// Name of the signature as configured on the system.
        /// </summary>
        public readonly string Name;
        public readonly bool? PerformStaging;
        /// <summary>
        /// The relative risk level of the attack that matches this signature.
        /// </summary>
        public readonly string Risk;
        /// <summary>
        /// ID of the signature in the database.
        /// </summary>
        public readonly int SignatureId;
        /// <summary>
        /// System generated ID of the signature.
        /// </summary>
        public readonly string SystemSignatureId;
        public readonly string Tag;
        /// <summary>
        /// Type of the signature.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWafSignaturesResult(
            string accuracy,

            string? description,

            bool? enabled,

            string id,

            string json,

            string name,

            bool? performStaging,

            string risk,

            int signatureId,

            string systemSignatureId,

            string tag,

            string type)
        {
            Accuracy = accuracy;
            Description = description;
            Enabled = enabled;
            Id = id;
            Json = json;
            Name = name;
            PerformStaging = performStaging;
            Risk = risk;
            SignatureId = signatureId;
            SystemSignatureId = systemSignatureId;
            Tag = tag;
            Type = type;
        }
    }
}
