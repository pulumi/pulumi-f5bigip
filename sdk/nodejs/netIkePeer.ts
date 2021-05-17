// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.NetIkePeer` Manages a ikePeer configuration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const example1 = new f5bigip.NetIkePeer("example1", {
 *     localAddress: "192.16.81.240",
 *     name: "example1",
 *     profile: "/Common/dslite",
 * });
 * ```
 */
export class NetIkePeer extends pulumi.CustomResource {
    /**
     * Get an existing NetIkePeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetIkePeerState, opts?: pulumi.CustomResourceOptions): NetIkePeer {
        return new NetIkePeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/netIkePeer:NetIkePeer';

    /**
     * Returns true if the given object is an instance of NetIkePeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetIkePeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetIkePeer.__pulumiType;
    }

    /**
     * The application service that the object belongs to
     */
    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * the trusted root and intermediate certificate authorities
     */
    public readonly caCertFile!: pulumi.Output<string>;
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     */
    public readonly crlFile!: pulumi.Output<string>;
    /**
     * User defined description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     */
    public readonly dpdDelay!: pulumi.Output<number>;
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     */
    public readonly generatePolicy!: pulumi.Output<string>;
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     */
    public readonly lifetime!: pulumi.Output<number>;
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Specifies the name of the certificate file object
     */
    public readonly myCertFile!: pulumi.Output<string>;
    /**
     * Specifies the name of the certificate key file object
     */
    public readonly myCertKeyFile!: pulumi.Output<string>;
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     */
    public readonly myCertKeyPassphrase!: pulumi.Output<string>;
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     */
    public readonly myIdType!: pulumi.Output<string>;
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     */
    public readonly myIdValue!: pulumi.Output<string>;
    /**
     * Name of the ike_peer
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enables use of the NAT-Traversal IPsec extension
     */
    public readonly natTraversal!: pulumi.Output<string>;
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     */
    public readonly passive!: pulumi.Output<string>;
    /**
     * Specifies the peer’s certificate for authentication
     */
    public readonly peersCertFile!: pulumi.Output<string>;
    /**
     * Specifies that the only peers-cert-type supported is certfile
     */
    public readonly peersCertType!: pulumi.Output<string>;
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     */
    public readonly peersIdType!: pulumi.Output<string>;
    /**
     * Specifies the peer’s identifier to be received
     */
    public readonly peersIdValue!: pulumi.Output<string>;
    /**
     * Specifies the authentication method used for phase 1 negotiation
     */
    public readonly phase1AuthMethod!: pulumi.Output<string>;
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     */
    public readonly phase1EncryptAlgorithm!: pulumi.Output<string>;
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     */
    public readonly phase1HashAlgorithm!: pulumi.Output<string>;
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     */
    public readonly phase1PerfectForwardSecrecy!: pulumi.Output<string>;
    /**
     * Specifies the preshared key for ISAKMP SAs
     */
    public readonly presharedKey!: pulumi.Output<string | undefined>;
    /**
     * Display the encrypted preshared-key for the IKE remote node
     */
    public readonly presharedKeyEncrypted!: pulumi.Output<string>;
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     */
    public readonly prf!: pulumi.Output<string>;
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     */
    public readonly proxySupport!: pulumi.Output<string>;
    /**
     * Specifies the IP address of the IKE remote node
     */
    public readonly remoteAddress!: pulumi.Output<string>;
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     */
    public readonly replayWindowSize!: pulumi.Output<number>;
    /**
     * Enables or disables this IKE remote node
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     */
    public readonly trafficSelectors!: pulumi.Output<string[]>;
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     */
    public readonly verifyCert!: pulumi.Output<string>;
    /**
     * Specifies which version of IKE to be used
     */
    public readonly versions!: pulumi.Output<string[]>;

    /**
     * Create a NetIkePeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetIkePeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetIkePeerArgs | NetIkePeerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetIkePeerState | undefined;
            inputs["appService"] = state ? state.appService : undefined;
            inputs["caCertFile"] = state ? state.caCertFile : undefined;
            inputs["crlFile"] = state ? state.crlFile : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dpdDelay"] = state ? state.dpdDelay : undefined;
            inputs["generatePolicy"] = state ? state.generatePolicy : undefined;
            inputs["lifetime"] = state ? state.lifetime : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["myCertFile"] = state ? state.myCertFile : undefined;
            inputs["myCertKeyFile"] = state ? state.myCertKeyFile : undefined;
            inputs["myCertKeyPassphrase"] = state ? state.myCertKeyPassphrase : undefined;
            inputs["myIdType"] = state ? state.myIdType : undefined;
            inputs["myIdValue"] = state ? state.myIdValue : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natTraversal"] = state ? state.natTraversal : undefined;
            inputs["passive"] = state ? state.passive : undefined;
            inputs["peersCertFile"] = state ? state.peersCertFile : undefined;
            inputs["peersCertType"] = state ? state.peersCertType : undefined;
            inputs["peersIdType"] = state ? state.peersIdType : undefined;
            inputs["peersIdValue"] = state ? state.peersIdValue : undefined;
            inputs["phase1AuthMethod"] = state ? state.phase1AuthMethod : undefined;
            inputs["phase1EncryptAlgorithm"] = state ? state.phase1EncryptAlgorithm : undefined;
            inputs["phase1HashAlgorithm"] = state ? state.phase1HashAlgorithm : undefined;
            inputs["phase1PerfectForwardSecrecy"] = state ? state.phase1PerfectForwardSecrecy : undefined;
            inputs["presharedKey"] = state ? state.presharedKey : undefined;
            inputs["presharedKeyEncrypted"] = state ? state.presharedKeyEncrypted : undefined;
            inputs["prf"] = state ? state.prf : undefined;
            inputs["proxySupport"] = state ? state.proxySupport : undefined;
            inputs["remoteAddress"] = state ? state.remoteAddress : undefined;
            inputs["replayWindowSize"] = state ? state.replayWindowSize : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["trafficSelectors"] = state ? state.trafficSelectors : undefined;
            inputs["verifyCert"] = state ? state.verifyCert : undefined;
            inputs["versions"] = state ? state.versions : undefined;
        } else {
            const args = argsOrState as NetIkePeerArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.remoteAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteAddress'");
            }
            inputs["appService"] = args ? args.appService : undefined;
            inputs["caCertFile"] = args ? args.caCertFile : undefined;
            inputs["crlFile"] = args ? args.crlFile : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dpdDelay"] = args ? args.dpdDelay : undefined;
            inputs["generatePolicy"] = args ? args.generatePolicy : undefined;
            inputs["lifetime"] = args ? args.lifetime : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["myCertFile"] = args ? args.myCertFile : undefined;
            inputs["myCertKeyFile"] = args ? args.myCertKeyFile : undefined;
            inputs["myCertKeyPassphrase"] = args ? args.myCertKeyPassphrase : undefined;
            inputs["myIdType"] = args ? args.myIdType : undefined;
            inputs["myIdValue"] = args ? args.myIdValue : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natTraversal"] = args ? args.natTraversal : undefined;
            inputs["passive"] = args ? args.passive : undefined;
            inputs["peersCertFile"] = args ? args.peersCertFile : undefined;
            inputs["peersCertType"] = args ? args.peersCertType : undefined;
            inputs["peersIdType"] = args ? args.peersIdType : undefined;
            inputs["peersIdValue"] = args ? args.peersIdValue : undefined;
            inputs["phase1AuthMethod"] = args ? args.phase1AuthMethod : undefined;
            inputs["phase1EncryptAlgorithm"] = args ? args.phase1EncryptAlgorithm : undefined;
            inputs["phase1HashAlgorithm"] = args ? args.phase1HashAlgorithm : undefined;
            inputs["phase1PerfectForwardSecrecy"] = args ? args.phase1PerfectForwardSecrecy : undefined;
            inputs["presharedKey"] = args ? args.presharedKey : undefined;
            inputs["presharedKeyEncrypted"] = args ? args.presharedKeyEncrypted : undefined;
            inputs["prf"] = args ? args.prf : undefined;
            inputs["proxySupport"] = args ? args.proxySupport : undefined;
            inputs["remoteAddress"] = args ? args.remoteAddress : undefined;
            inputs["replayWindowSize"] = args ? args.replayWindowSize : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["trafficSelectors"] = args ? args.trafficSelectors : undefined;
            inputs["verifyCert"] = args ? args.verifyCert : undefined;
            inputs["versions"] = args ? args.versions : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetIkePeer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetIkePeer resources.
 */
export interface NetIkePeerState {
    /**
     * The application service that the object belongs to
     */
    readonly appService?: pulumi.Input<string>;
    /**
     * the trusted root and intermediate certificate authorities
     */
    readonly caCertFile?: pulumi.Input<string>;
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     */
    readonly crlFile?: pulumi.Input<string>;
    /**
     * User defined description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     */
    readonly dpdDelay?: pulumi.Input<number>;
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     */
    readonly generatePolicy?: pulumi.Input<string>;
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     */
    readonly lifetime?: pulumi.Input<number>;
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate file object
     */
    readonly myCertFile?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate key file object
     */
    readonly myCertKeyFile?: pulumi.Input<string>;
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     */
    readonly myCertKeyPassphrase?: pulumi.Input<string>;
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     */
    readonly myIdType?: pulumi.Input<string>;
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     */
    readonly myIdValue?: pulumi.Input<string>;
    /**
     * Name of the ike_peer
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Enables use of the NAT-Traversal IPsec extension
     */
    readonly natTraversal?: pulumi.Input<string>;
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     */
    readonly passive?: pulumi.Input<string>;
    /**
     * Specifies the peer’s certificate for authentication
     */
    readonly peersCertFile?: pulumi.Input<string>;
    /**
     * Specifies that the only peers-cert-type supported is certfile
     */
    readonly peersCertType?: pulumi.Input<string>;
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     */
    readonly peersIdType?: pulumi.Input<string>;
    /**
     * Specifies the peer’s identifier to be received
     */
    readonly peersIdValue?: pulumi.Input<string>;
    /**
     * Specifies the authentication method used for phase 1 negotiation
     */
    readonly phase1AuthMethod?: pulumi.Input<string>;
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     */
    readonly phase1EncryptAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     */
    readonly phase1HashAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     */
    readonly phase1PerfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the preshared key for ISAKMP SAs
     */
    readonly presharedKey?: pulumi.Input<string>;
    /**
     * Display the encrypted preshared-key for the IKE remote node
     */
    readonly presharedKeyEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     */
    readonly prf?: pulumi.Input<string>;
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     */
    readonly proxySupport?: pulumi.Input<string>;
    /**
     * Specifies the IP address of the IKE remote node
     */
    readonly remoteAddress?: pulumi.Input<string>;
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     */
    readonly replayWindowSize?: pulumi.Input<number>;
    /**
     * Enables or disables this IKE remote node
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     */
    readonly trafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     */
    readonly verifyCert?: pulumi.Input<string>;
    /**
     * Specifies which version of IKE to be used
     */
    readonly versions?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NetIkePeer resource.
 */
export interface NetIkePeerArgs {
    /**
     * The application service that the object belongs to
     */
    readonly appService?: pulumi.Input<string>;
    /**
     * the trusted root and intermediate certificate authorities
     */
    readonly caCertFile?: pulumi.Input<string>;
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     */
    readonly crlFile?: pulumi.Input<string>;
    /**
     * User defined description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     */
    readonly dpdDelay?: pulumi.Input<number>;
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     */
    readonly generatePolicy?: pulumi.Input<string>;
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     */
    readonly lifetime?: pulumi.Input<number>;
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate file object
     */
    readonly myCertFile?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate key file object
     */
    readonly myCertKeyFile?: pulumi.Input<string>;
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     */
    readonly myCertKeyPassphrase?: pulumi.Input<string>;
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     */
    readonly myIdType?: pulumi.Input<string>;
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     */
    readonly myIdValue?: pulumi.Input<string>;
    /**
     * Name of the ike_peer
     */
    readonly name: pulumi.Input<string>;
    /**
     * Enables use of the NAT-Traversal IPsec extension
     */
    readonly natTraversal?: pulumi.Input<string>;
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     */
    readonly passive?: pulumi.Input<string>;
    /**
     * Specifies the peer’s certificate for authentication
     */
    readonly peersCertFile?: pulumi.Input<string>;
    /**
     * Specifies that the only peers-cert-type supported is certfile
     */
    readonly peersCertType?: pulumi.Input<string>;
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     */
    readonly peersIdType?: pulumi.Input<string>;
    /**
     * Specifies the peer’s identifier to be received
     */
    readonly peersIdValue?: pulumi.Input<string>;
    /**
     * Specifies the authentication method used for phase 1 negotiation
     */
    readonly phase1AuthMethod?: pulumi.Input<string>;
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     */
    readonly phase1EncryptAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     */
    readonly phase1HashAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     */
    readonly phase1PerfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the preshared key for ISAKMP SAs
     */
    readonly presharedKey?: pulumi.Input<string>;
    /**
     * Display the encrypted preshared-key for the IKE remote node
     */
    readonly presharedKeyEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     */
    readonly prf?: pulumi.Input<string>;
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     */
    readonly proxySupport?: pulumi.Input<string>;
    /**
     * Specifies the IP address of the IKE remote node
     */
    readonly remoteAddress: pulumi.Input<string>;
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     */
    readonly replayWindowSize?: pulumi.Input<number>;
    /**
     * Enables or disables this IKE remote node
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     */
    readonly trafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     */
    readonly verifyCert?: pulumi.Input<string>;
    /**
     * Specifies which version of IKE to be used
     */
    readonly versions?: pulumi.Input<pulumi.Input<string>[]>;
}
