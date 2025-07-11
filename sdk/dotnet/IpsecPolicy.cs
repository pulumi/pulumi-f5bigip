// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.IpsecPolicy` Manage IPSec policies on a BIG-IP
    /// 
    /// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-policy)
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
    ///     var test_policy = new F5BigIP.IpsecPolicy("test-policy", new()
    ///     {
    ///         Name = "/Common/test-policy",
    ///         Description = "created by terraform provider",
    ///         Protocol = "esp",
    ///         Mode = "tunnel",
    ///         TunnelLocalAddress = "192.168.1.1",
    ///         TunnelRemoteAddress = "10.10.1.1",
    ///         AuthAlgorithm = "sha1",
    ///         EncryptAlgorithm = "3des",
    ///         Lifetime = 3,
    ///         Ipcomp = "deflate",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/ipsecPolicy:IpsecPolicy")]
    public partial class IpsecPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        /// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        /// </summary>
        [Output("authAlgorithm")]
        public Output<string> AuthAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Description of the IPSec policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        /// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        /// </summary>
        [Output("encryptAlgorithm")]
        public Output<string> EncryptAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        /// </summary>
        [Output("ipcomp")]
        public Output<string> Ipcomp { get; private set; } = null!;

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in kilobytes.
        /// </summary>
        [Output("kbLifetime")]
        public Output<int> KbLifetime { get; private set; } = null!;

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in minutes.
        /// </summary>
        [Output("lifetime")]
        public Output<int> Lifetime { get; private set; } = null!;

        /// <summary>
        /// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        /// modp4096, modp6144, modp8192`
        /// </summary>
        [Output("perfectForwardSecrecy")]
        public Output<string> PerfectForwardSecrecy { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPsec protocol. Valid choices are: `ah, esp`
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Output("tunnelLocalAddress")]
        public Output<string> TunnelLocalAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Output("tunnelRemoteAddress")]
        public Output<string> TunnelRemoteAddress { get; private set; } = null!;


        /// <summary>
        /// Create a IpsecPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpsecPolicy(string name, IpsecPolicyArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/ipsecPolicy:IpsecPolicy", name, args ?? new IpsecPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpsecPolicy(string name, Input<string> id, IpsecPolicyState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/ipsecPolicy:IpsecPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpsecPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpsecPolicy Get(string name, Input<string> id, IpsecPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IpsecPolicy(name, id, state, options);
        }
    }

    public sealed class IpsecPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        /// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        /// </summary>
        [Input("authAlgorithm")]
        public Input<string>? AuthAlgorithm { get; set; }

        /// <summary>
        /// Description of the IPSec policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        /// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        /// </summary>
        [Input("encryptAlgorithm")]
        public Input<string>? EncryptAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        /// </summary>
        [Input("ipcomp")]
        public Input<string>? Ipcomp { get; set; }

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in kilobytes.
        /// </summary>
        [Input("kbLifetime")]
        public Input<int>? KbLifetime { get; set; }

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in minutes.
        /// </summary>
        [Input("lifetime")]
        public Input<int>? Lifetime { get; set; }

        /// <summary>
        /// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        /// modp4096, modp6144, modp8192`
        /// </summary>
        [Input("perfectForwardSecrecy")]
        public Input<string>? PerfectForwardSecrecy { get; set; }

        /// <summary>
        /// Specifies the IPsec protocol. Valid choices are: `ah, esp`
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Input("tunnelLocalAddress")]
        public Input<string>? TunnelLocalAddress { get; set; }

        /// <summary>
        /// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Input("tunnelRemoteAddress")]
        public Input<string>? TunnelRemoteAddress { get; set; }

        public IpsecPolicyArgs()
        {
        }
        public static new IpsecPolicyArgs Empty => new IpsecPolicyArgs();
    }

    public sealed class IpsecPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        /// aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        /// </summary>
        [Input("authAlgorithm")]
        public Input<string>? AuthAlgorithm { get; set; }

        /// <summary>
        /// Description of the IPSec policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        /// aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        /// </summary>
        [Input("encryptAlgorithm")]
        public Input<string>? EncryptAlgorithm { get; set; }

        /// <summary>
        /// Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        /// </summary>
        [Input("ipcomp")]
        public Input<string>? Ipcomp { get; set; }

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in kilobytes.
        /// </summary>
        [Input("kbLifetime")]
        public Input<int>? KbLifetime { get; set; }

        /// <summary>
        /// Specifies the length of time before the IKE security association expires, in minutes.
        /// </summary>
        [Input("lifetime")]
        public Input<int>? Lifetime { get; set; }

        /// <summary>
        /// Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        /// modp4096, modp6144, modp8192`
        /// </summary>
        [Input("perfectForwardSecrecy")]
        public Input<string>? PerfectForwardSecrecy { get; set; }

        /// <summary>
        /// Specifies the IPsec protocol. Valid choices are: `ah, esp`
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Input("tunnelLocalAddress")]
        public Input<string>? TunnelLocalAddress { get; set; }

        /// <summary>
        /// Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        /// </summary>
        [Input("tunnelRemoteAddress")]
        public Input<string>? TunnelRemoteAddress { get; set; }

        public IpsecPolicyState()
        {
        }
        public static new IpsecPolicyState Empty => new IpsecPolicyState();
    }
}
