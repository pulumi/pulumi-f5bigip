// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys
{
    /// <summary>
    /// `f5bigip.sys.Ocsp` Manages F5 BIG-IP OCSP responder using iControl REST.
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
    ///     var test_ocsp = new F5BigIP.Sys.Ocsp("test-ocsp", new()
    ///     {
    ///         Name = "/Uncommon/test-ocsp",
    ///         Passphrase = "testabcdef",
    ///         ProxyServerPool = "/Common/test-poolxyz",
    ///         SignerCert = "/Common/le-ssl",
    ///         SignerKey = "/Common/le-ssl",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Importing
    /// 
    /// An existing OCSP can be imported into this resource by supplying the full path name  ex : `/partition/name`
    /// An example is below:
    /// ```sh
    /// $ terraform import bigip_sys_ocsp.test-ocsp /Common/test-ocsp
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:sys/ocsp:Ocsp")]
    public partial class Ocsp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
        /// </summary>
        [Output("cacheErrorTimeout")]
        public Output<int?> CacheErrorTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
        /// </summary>
        [Output("cacheTimeout")]
        public Output<string?> CacheTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
        /// </summary>
        [Output("clockSkew")]
        public Output<int?> ClockSkew { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
        /// </summary>
        [Output("concurrentConnectionsLimit")]
        public Output<int?> ConcurrentConnectionsLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
        /// </summary>
        [Output("connectionTimeout")]
        public Output<int?> ConnectionTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Output("dnsResolver")]
        public Output<string?> DnsResolver { get; private set; } = null!;

        /// <summary>
        /// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies a passphrase used to sign an OCSP request.
        /// </summary>
        [Output("passphrase")]
        public Output<string?> Passphrase { get; private set; } = null!;

        /// <summary>
        /// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Output("proxyServerPool")]
        public Output<string?> ProxyServerPool { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL of the OCSP responder.
        /// </summary>
        [Output("responderUrl")]
        public Output<string?> ResponderUrl { get; private set; } = null!;

        /// <summary>
        /// Specifies the route domain for the OCSP responder.
        /// </summary>
        [Output("routeDomain")]
        public Output<string?> RouteDomain { get; private set; } = null!;

        /// <summary>
        /// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
        /// </summary>
        [Output("signHash")]
        public Output<string?> SignHash { get; private set; } = null!;

        /// <summary>
        /// Specifies the certificate used to sign the OCSP request.
        /// </summary>
        [Output("signerCert")]
        public Output<string?> SignerCert { get; private set; } = null!;

        /// <summary>
        /// Specifies the key used to sign the OCSP request.
        /// </summary>
        [Output("signerKey")]
        public Output<string?> SignerKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
        /// </summary>
        [Output("statusAge")]
        public Output<int?> StatusAge { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
        /// </summary>
        [Output("strictRespCertCheck")]
        public Output<string?> StrictRespCertCheck { get; private set; } = null!;

        /// <summary>
        /// Specifies the certificates used for validating the OCSP response.
        /// </summary>
        [Output("trustedResponders")]
        public Output<string?> TrustedResponders { get; private set; } = null!;


        /// <summary>
        /// Create a Ocsp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ocsp(string name, OcspArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:sys/ocsp:Ocsp", name, args ?? new OcspArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ocsp(string name, Input<string> id, OcspState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/ocsp:Ocsp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "passphrase",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Ocsp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ocsp Get(string name, Input<string> id, OcspState? state = null, CustomResourceOptions? options = null)
        {
            return new Ocsp(name, id, state, options);
        }
    }

    public sealed class OcspArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
        /// </summary>
        [Input("cacheErrorTimeout")]
        public Input<int>? CacheErrorTimeout { get; set; }

        /// <summary>
        /// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
        /// </summary>
        [Input("cacheTimeout")]
        public Input<string>? CacheTimeout { get; set; }

        /// <summary>
        /// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
        /// </summary>
        [Input("clockSkew")]
        public Input<int>? ClockSkew { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
        /// </summary>
        [Input("concurrentConnectionsLimit")]
        public Input<int>? ConcurrentConnectionsLimit { get; set; }

        /// <summary>
        /// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
        /// </summary>
        [Input("connectionTimeout")]
        public Input<int>? ConnectionTimeout { get; set; }

        /// <summary>
        /// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Input("dnsResolver")]
        public Input<string>? DnsResolver { get; set; }

        /// <summary>
        /// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// Specifies a passphrase used to sign an OCSP request.
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Input("proxyServerPool")]
        public Input<string>? ProxyServerPool { get; set; }

        /// <summary>
        /// Specifies the URL of the OCSP responder.
        /// </summary>
        [Input("responderUrl")]
        public Input<string>? ResponderUrl { get; set; }

        /// <summary>
        /// Specifies the route domain for the OCSP responder.
        /// </summary>
        [Input("routeDomain")]
        public Input<string>? RouteDomain { get; set; }

        /// <summary>
        /// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
        /// </summary>
        [Input("signHash")]
        public Input<string>? SignHash { get; set; }

        /// <summary>
        /// Specifies the certificate used to sign the OCSP request.
        /// </summary>
        [Input("signerCert")]
        public Input<string>? SignerCert { get; set; }

        /// <summary>
        /// Specifies the key used to sign the OCSP request.
        /// </summary>
        [Input("signerKey")]
        public Input<string>? SignerKey { get; set; }

        /// <summary>
        /// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
        /// </summary>
        [Input("statusAge")]
        public Input<int>? StatusAge { get; set; }

        /// <summary>
        /// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
        /// </summary>
        [Input("strictRespCertCheck")]
        public Input<string>? StrictRespCertCheck { get; set; }

        /// <summary>
        /// Specifies the certificates used for validating the OCSP response.
        /// </summary>
        [Input("trustedResponders")]
        public Input<string>? TrustedResponders { get; set; }

        public OcspArgs()
        {
        }
        public static new OcspArgs Empty => new OcspArgs();
    }

    public sealed class OcspState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
        /// </summary>
        [Input("cacheErrorTimeout")]
        public Input<int>? CacheErrorTimeout { get; set; }

        /// <summary>
        /// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
        /// </summary>
        [Input("cacheTimeout")]
        public Input<string>? CacheTimeout { get; set; }

        /// <summary>
        /// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
        /// </summary>
        [Input("clockSkew")]
        public Input<int>? ClockSkew { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
        /// </summary>
        [Input("concurrentConnectionsLimit")]
        public Input<int>? ConcurrentConnectionsLimit { get; set; }

        /// <summary>
        /// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
        /// </summary>
        [Input("connectionTimeout")]
        public Input<int>? ConnectionTimeout { get; set; }

        /// <summary>
        /// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Input("dnsResolver")]
        public Input<string>? DnsResolver { get; set; }

        /// <summary>
        /// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// Specifies a passphrase used to sign an OCSP request.
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
        /// </summary>
        [Input("proxyServerPool")]
        public Input<string>? ProxyServerPool { get; set; }

        /// <summary>
        /// Specifies the URL of the OCSP responder.
        /// </summary>
        [Input("responderUrl")]
        public Input<string>? ResponderUrl { get; set; }

        /// <summary>
        /// Specifies the route domain for the OCSP responder.
        /// </summary>
        [Input("routeDomain")]
        public Input<string>? RouteDomain { get; set; }

        /// <summary>
        /// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
        /// </summary>
        [Input("signHash")]
        public Input<string>? SignHash { get; set; }

        /// <summary>
        /// Specifies the certificate used to sign the OCSP request.
        /// </summary>
        [Input("signerCert")]
        public Input<string>? SignerCert { get; set; }

        /// <summary>
        /// Specifies the key used to sign the OCSP request.
        /// </summary>
        [Input("signerKey")]
        public Input<string>? SignerKey { get; set; }

        /// <summary>
        /// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
        /// </summary>
        [Input("statusAge")]
        public Input<int>? StatusAge { get; set; }

        /// <summary>
        /// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
        /// </summary>
        [Input("strictRespCertCheck")]
        public Input<string>? StrictRespCertCheck { get; set; }

        /// <summary>
        /// Specifies the certificates used for validating the OCSP response.
        /// </summary>
        [Input("trustedResponders")]
        public Input<string>? TrustedResponders { get; set; }

        public OcspState()
        {
        }
        public static new OcspState Empty => new OcspState();
    }
}
