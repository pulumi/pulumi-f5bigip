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
    /// `f5bigip.NetIkePeer` Manages a ike_peer configuration
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/netIkePeer:NetIkePeer")]
    public partial class NetIkePeer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Output("appService")]
        public Output<string?> AppService { get; private set; } = null!;

        /// <summary>
        /// the trusted root and intermediate certificate authorities
        /// </summary>
        [Output("caCertFile")]
        public Output<string> CaCertFile { get; private set; } = null!;

        /// <summary>
        /// Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
        /// </summary>
        [Output("crlFile")]
        public Output<string> CrlFile { get; private set; } = null!;

        /// <summary>
        /// User defined description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds between Dead Peer Detection messages
        /// </summary>
        [Output("dpdDelay")]
        public Output<int> DpdDelay { get; private set; } = null!;

        /// <summary>
        /// Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
        /// </summary>
        [Output("generatePolicy")]
        public Output<string> GeneratePolicy { get; private set; } = null!;

        /// <summary>
        /// Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
        /// </summary>
        [Output("lifetime")]
        public Output<int> Lifetime { get; private set; } = null!;

        /// <summary>
        /// Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the certificate file object
        /// </summary>
        [Output("myCertFile")]
        public Output<string> MyCertFile { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the certificate key file object
        /// </summary>
        [Output("myCertKeyFile")]
        public Output<string> MyCertKeyFile { get; private set; } = null!;

        /// <summary>
        /// Specifies the passphrase of the key used for my-cert-key-file
        /// </summary>
        [Output("myCertKeyPassphrase")]
        public Output<string> MyCertKeyPassphrase { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
        /// </summary>
        [Output("myIdType")]
        public Output<string> MyIdType { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier value sent to the remote host in the phase 1 negotiation
        /// </summary>
        [Output("myIdValue")]
        public Output<string> MyIdValue { get; private set; } = null!;

        /// <summary>
        /// Name of the ike_peer
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enables use of the NAT-Traversal IPsec extension
        /// </summary>
        [Output("natTraversal")]
        public Output<string> NatTraversal { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
        /// </summary>
        [Output("passive")]
        public Output<string> Passive { get; private set; } = null!;

        /// <summary>
        /// Specifies the peer’s certificate for authentication
        /// </summary>
        [Output("peersCertFile")]
        public Output<string> PeersCertFile { get; private set; } = null!;

        /// <summary>
        /// Specifies that the only peers-cert-type supported is certfile
        /// </summary>
        [Output("peersCertType")]
        public Output<string> PeersCertType { get; private set; } = null!;

        /// <summary>
        /// Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
        /// </summary>
        [Output("peersIdType")]
        public Output<string> PeersIdType { get; private set; } = null!;

        /// <summary>
        /// Specifies the peer’s identifier to be received
        /// </summary>
        [Output("peersIdValue")]
        public Output<string> PeersIdValue { get; private set; } = null!;

        /// <summary>
        /// Specifies the authentication method used for phase 1 negotiation
        /// </summary>
        [Output("phase1AuthMethod")]
        public Output<string> Phase1AuthMethod { get; private set; } = null!;

        /// <summary>
        /// Specifies the encryption algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Output("phase1EncryptAlgorithm")]
        public Output<string> Phase1EncryptAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Defines the hash algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Output("phase1HashAlgorithm")]
        public Output<string> Phase1HashAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
        /// </summary>
        [Output("phase1PerfectForwardSecrecy")]
        public Output<string> Phase1PerfectForwardSecrecy { get; private set; } = null!;

        /// <summary>
        /// Specifies the preshared key for ISAKMP SAs
        /// </summary>
        [Output("presharedKey")]
        public Output<string?> PresharedKey { get; private set; } = null!;

        /// <summary>
        /// Display the encrypted preshared-key for the IKE remote node
        /// </summary>
        [Output("presharedKeyEncrypted")]
        public Output<string> PresharedKeyEncrypted { get; private set; } = null!;

        /// <summary>
        /// Specifies the pseudo-random function used to derive keying material for all cryptographic operations
        /// </summary>
        [Output("prf")]
        public Output<string> Prf { get; private set; } = null!;

        /// <summary>
        /// If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
        /// </summary>
        [Output("proxySupport")]
        public Output<string> ProxySupport { get; private set; } = null!;

        /// <summary>
        /// Specifies the IP address of the IKE remote node
        /// </summary>
        [Output("remoteAddress")]
        public Output<string> RemoteAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
        /// </summary>
        [Output("replayWindowSize")]
        public Output<int> ReplayWindowSize { get; private set; } = null!;

        /// <summary>
        /// Enables or disables this IKE remote node
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Specifies the names of the traffic-selector objects associated with this ike-peer
        /// </summary>
        [Output("trafficSelectors")]
        public Output<ImmutableArray<string>> TrafficSelectors { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
        /// </summary>
        [Output("verifyCert")]
        public Output<string> VerifyCert { get; private set; } = null!;

        /// <summary>
        /// Specifies which version of IKE to be used
        /// </summary>
        [Output("versions")]
        public Output<ImmutableArray<string>> Versions { get; private set; } = null!;


        /// <summary>
        /// Create a NetIkePeer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetIkePeer(string name, NetIkePeerArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/netIkePeer:NetIkePeer", name, args ?? new NetIkePeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetIkePeer(string name, Input<string> id, NetIkePeerState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/netIkePeer:NetIkePeer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetIkePeer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetIkePeer Get(string name, Input<string> id, NetIkePeerState? state = null, CustomResourceOptions? options = null)
        {
            return new NetIkePeer(name, id, state, options);
        }
    }

    public sealed class NetIkePeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// the trusted root and intermediate certificate authorities
        /// </summary>
        [Input("caCertFile")]
        public Input<string>? CaCertFile { get; set; }

        /// <summary>
        /// Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
        /// </summary>
        [Input("crlFile")]
        public Input<string>? CrlFile { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the number of seconds between Dead Peer Detection messages
        /// </summary>
        [Input("dpdDelay")]
        public Input<int>? DpdDelay { get; set; }

        /// <summary>
        /// Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
        /// </summary>
        [Input("generatePolicy")]
        public Input<string>? GeneratePolicy { get; set; }

        /// <summary>
        /// Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
        /// </summary>
        [Input("lifetime")]
        public Input<int>? Lifetime { get; set; }

        /// <summary>
        /// Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the name of the certificate file object
        /// </summary>
        [Input("myCertFile")]
        public Input<string>? MyCertFile { get; set; }

        /// <summary>
        /// Specifies the name of the certificate key file object
        /// </summary>
        [Input("myCertKeyFile")]
        public Input<string>? MyCertKeyFile { get; set; }

        /// <summary>
        /// Specifies the passphrase of the key used for my-cert-key-file
        /// </summary>
        [Input("myCertKeyPassphrase")]
        public Input<string>? MyCertKeyPassphrase { get; set; }

        /// <summary>
        /// Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
        /// </summary>
        [Input("myIdType")]
        public Input<string>? MyIdType { get; set; }

        /// <summary>
        /// Specifies the identifier value sent to the remote host in the phase 1 negotiation
        /// </summary>
        [Input("myIdValue")]
        public Input<string>? MyIdValue { get; set; }

        /// <summary>
        /// Name of the ike_peer
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Enables use of the NAT-Traversal IPsec extension
        /// </summary>
        [Input("natTraversal")]
        public Input<string>? NatTraversal { get; set; }

        /// <summary>
        /// Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
        /// </summary>
        [Input("passive")]
        public Input<string>? Passive { get; set; }

        /// <summary>
        /// Specifies the peer’s certificate for authentication
        /// </summary>
        [Input("peersCertFile")]
        public Input<string>? PeersCertFile { get; set; }

        /// <summary>
        /// Specifies that the only peers-cert-type supported is certfile
        /// </summary>
        [Input("peersCertType")]
        public Input<string>? PeersCertType { get; set; }

        /// <summary>
        /// Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
        /// </summary>
        [Input("peersIdType")]
        public Input<string>? PeersIdType { get; set; }

        /// <summary>
        /// Specifies the peer’s identifier to be received
        /// </summary>
        [Input("peersIdValue")]
        public Input<string>? PeersIdValue { get; set; }

        /// <summary>
        /// Specifies the authentication method used for phase 1 negotiation
        /// </summary>
        [Input("phase1AuthMethod")]
        public Input<string>? Phase1AuthMethod { get; set; }

        /// <summary>
        /// Specifies the encryption algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Input("phase1EncryptAlgorithm")]
        public Input<string>? Phase1EncryptAlgorithm { get; set; }

        /// <summary>
        /// Defines the hash algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Input("phase1HashAlgorithm")]
        public Input<string>? Phase1HashAlgorithm { get; set; }

        /// <summary>
        /// Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
        /// </summary>
        [Input("phase1PerfectForwardSecrecy")]
        public Input<string>? Phase1PerfectForwardSecrecy { get; set; }

        /// <summary>
        /// Specifies the preshared key for ISAKMP SAs
        /// </summary>
        [Input("presharedKey")]
        public Input<string>? PresharedKey { get; set; }

        /// <summary>
        /// Display the encrypted preshared-key for the IKE remote node
        /// </summary>
        [Input("presharedKeyEncrypted")]
        public Input<string>? PresharedKeyEncrypted { get; set; }

        /// <summary>
        /// Specifies the pseudo-random function used to derive keying material for all cryptographic operations
        /// </summary>
        [Input("prf")]
        public Input<string>? Prf { get; set; }

        /// <summary>
        /// If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
        /// </summary>
        [Input("proxySupport")]
        public Input<string>? ProxySupport { get; set; }

        /// <summary>
        /// Specifies the IP address of the IKE remote node
        /// </summary>
        [Input("remoteAddress", required: true)]
        public Input<string> RemoteAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
        /// </summary>
        [Input("replayWindowSize")]
        public Input<int>? ReplayWindowSize { get; set; }

        /// <summary>
        /// Enables or disables this IKE remote node
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("trafficSelectors")]
        private InputList<string>? _trafficSelectors;

        /// <summary>
        /// Specifies the names of the traffic-selector objects associated with this ike-peer
        /// </summary>
        public InputList<string> TrafficSelectors
        {
            get => _trafficSelectors ?? (_trafficSelectors = new InputList<string>());
            set => _trafficSelectors = value;
        }

        /// <summary>
        /// Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
        /// </summary>
        [Input("verifyCert")]
        public Input<string>? VerifyCert { get; set; }

        [Input("versions")]
        private InputList<string>? _versions;

        /// <summary>
        /// Specifies which version of IKE to be used
        /// </summary>
        public InputList<string> Versions
        {
            get => _versions ?? (_versions = new InputList<string>());
            set => _versions = value;
        }

        public NetIkePeerArgs()
        {
        }
        public static new NetIkePeerArgs Empty => new NetIkePeerArgs();
    }

    public sealed class NetIkePeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application service that the object belongs to
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// the trusted root and intermediate certificate authorities
        /// </summary>
        [Input("caCertFile")]
        public Input<string>? CaCertFile { get; set; }

        /// <summary>
        /// Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
        /// </summary>
        [Input("crlFile")]
        public Input<string>? CrlFile { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the number of seconds between Dead Peer Detection messages
        /// </summary>
        [Input("dpdDelay")]
        public Input<int>? DpdDelay { get; set; }

        /// <summary>
        /// Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
        /// </summary>
        [Input("generatePolicy")]
        public Input<string>? GeneratePolicy { get; set; }

        /// <summary>
        /// Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
        /// </summary>
        [Input("lifetime")]
        public Input<int>? Lifetime { get; set; }

        /// <summary>
        /// Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the name of the certificate file object
        /// </summary>
        [Input("myCertFile")]
        public Input<string>? MyCertFile { get; set; }

        /// <summary>
        /// Specifies the name of the certificate key file object
        /// </summary>
        [Input("myCertKeyFile")]
        public Input<string>? MyCertKeyFile { get; set; }

        /// <summary>
        /// Specifies the passphrase of the key used for my-cert-key-file
        /// </summary>
        [Input("myCertKeyPassphrase")]
        public Input<string>? MyCertKeyPassphrase { get; set; }

        /// <summary>
        /// Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
        /// </summary>
        [Input("myIdType")]
        public Input<string>? MyIdType { get; set; }

        /// <summary>
        /// Specifies the identifier value sent to the remote host in the phase 1 negotiation
        /// </summary>
        [Input("myIdValue")]
        public Input<string>? MyIdValue { get; set; }

        /// <summary>
        /// Name of the ike_peer
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enables use of the NAT-Traversal IPsec extension
        /// </summary>
        [Input("natTraversal")]
        public Input<string>? NatTraversal { get; set; }

        /// <summary>
        /// Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
        /// </summary>
        [Input("passive")]
        public Input<string>? Passive { get; set; }

        /// <summary>
        /// Specifies the peer’s certificate for authentication
        /// </summary>
        [Input("peersCertFile")]
        public Input<string>? PeersCertFile { get; set; }

        /// <summary>
        /// Specifies that the only peers-cert-type supported is certfile
        /// </summary>
        [Input("peersCertType")]
        public Input<string>? PeersCertType { get; set; }

        /// <summary>
        /// Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
        /// </summary>
        [Input("peersIdType")]
        public Input<string>? PeersIdType { get; set; }

        /// <summary>
        /// Specifies the peer’s identifier to be received
        /// </summary>
        [Input("peersIdValue")]
        public Input<string>? PeersIdValue { get; set; }

        /// <summary>
        /// Specifies the authentication method used for phase 1 negotiation
        /// </summary>
        [Input("phase1AuthMethod")]
        public Input<string>? Phase1AuthMethod { get; set; }

        /// <summary>
        /// Specifies the encryption algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Input("phase1EncryptAlgorithm")]
        public Input<string>? Phase1EncryptAlgorithm { get; set; }

        /// <summary>
        /// Defines the hash algorithm used for the isakmp phase 1 negotiation
        /// </summary>
        [Input("phase1HashAlgorithm")]
        public Input<string>? Phase1HashAlgorithm { get; set; }

        /// <summary>
        /// Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
        /// </summary>
        [Input("phase1PerfectForwardSecrecy")]
        public Input<string>? Phase1PerfectForwardSecrecy { get; set; }

        /// <summary>
        /// Specifies the preshared key for ISAKMP SAs
        /// </summary>
        [Input("presharedKey")]
        public Input<string>? PresharedKey { get; set; }

        /// <summary>
        /// Display the encrypted preshared-key for the IKE remote node
        /// </summary>
        [Input("presharedKeyEncrypted")]
        public Input<string>? PresharedKeyEncrypted { get; set; }

        /// <summary>
        /// Specifies the pseudo-random function used to derive keying material for all cryptographic operations
        /// </summary>
        [Input("prf")]
        public Input<string>? Prf { get; set; }

        /// <summary>
        /// If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
        /// </summary>
        [Input("proxySupport")]
        public Input<string>? ProxySupport { get; set; }

        /// <summary>
        /// Specifies the IP address of the IKE remote node
        /// </summary>
        [Input("remoteAddress")]
        public Input<string>? RemoteAddress { get; set; }

        /// <summary>
        /// Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
        /// </summary>
        [Input("replayWindowSize")]
        public Input<int>? ReplayWindowSize { get; set; }

        /// <summary>
        /// Enables or disables this IKE remote node
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("trafficSelectors")]
        private InputList<string>? _trafficSelectors;

        /// <summary>
        /// Specifies the names of the traffic-selector objects associated with this ike-peer
        /// </summary>
        public InputList<string> TrafficSelectors
        {
            get => _trafficSelectors ?? (_trafficSelectors = new InputList<string>());
            set => _trafficSelectors = value;
        }

        /// <summary>
        /// Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
        /// </summary>
        [Input("verifyCert")]
        public Input<string>? VerifyCert { get; set; }

        [Input("versions")]
        private InputList<string>? _versions;

        /// <summary>
        /// Specifies which version of IKE to be used
        /// </summary>
        public InputList<string> Versions
        {
            get => _versions ?? (_versions = new InputList<string>());
            set => _versions = value;
        }

        public NetIkePeerState()
        {
        }
        public static new NetIkePeerState Empty => new NetIkePeerState();
    }
}
