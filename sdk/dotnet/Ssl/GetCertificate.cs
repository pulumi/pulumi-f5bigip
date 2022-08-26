// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source (`f5bigip.ssl.Certificate`) to get the ssl-certificate details available on BIG-IP
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
        ///     var test = F5BigIP.Ssl.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "terraform_ssl_certificate",
        ///         Partition = "Common",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipSslCertificateName"] = test.Apply(getCertificateResult =&gt; getCertificateResult.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("f5bigip:ssl/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.Certificate`) to get the ssl-certificate details available on BIG-IP
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
        ///     var test = F5BigIP.Ssl.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "terraform_ssl_certificate",
        ///         Partition = "Common",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipSslCertificateName"] = test.Apply(getCertificateResult =&gt; getCertificateResult.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("f5bigip:ssl/getCertificate:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ssl_certificate
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm ssl_certificate
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ssl_certificate
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm ssl_certificate
        /// </summary>
        [Input("partition", required: true)]
        public Input<string> Partition { get; set; } = null!;

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string Certificate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of ssl_certificate configured on bigip with full path
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Bigip partition in which ssl-certificate is configured
        /// </summary>
        public readonly string Partition;

        [OutputConstructor]
        private GetCertificateResult(
            string certificate,

            string id,

            string name,

            string partition)
        {
            Certificate = certificate;
            Id = id;
            Name = name;
            Partition = partition;
        }
    }
}
