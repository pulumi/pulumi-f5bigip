// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.NetIkePeer` Manages a ikePeer configuration
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetIkePeerState | undefined;
            resourceInputs["appService"] = state ? state.appService : undefined;
            resourceInputs["caCertFile"] = state ? state.caCertFile : undefined;
            resourceInputs["crlFile"] = state ? state.crlFile : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dpdDelay"] = state ? state.dpdDelay : undefined;
            resourceInputs["generatePolicy"] = state ? state.generatePolicy : undefined;
            resourceInputs["lifetime"] = state ? state.lifetime : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["myCertFile"] = state ? state.myCertFile : undefined;
            resourceInputs["myCertKeyFile"] = state ? state.myCertKeyFile : undefined;
            resourceInputs["myCertKeyPassphrase"] = state ? state.myCertKeyPassphrase : undefined;
            resourceInputs["myIdType"] = state ? state.myIdType : undefined;
            resourceInputs["myIdValue"] = state ? state.myIdValue : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["natTraversal"] = state ? state.natTraversal : undefined;
            resourceInputs["passive"] = state ? state.passive : undefined;
            resourceInputs["peersCertFile"] = state ? state.peersCertFile : undefined;
            resourceInputs["peersCertType"] = state ? state.peersCertType : undefined;
            resourceInputs["peersIdType"] = state ? state.peersIdType : undefined;
            resourceInputs["peersIdValue"] = state ? state.peersIdValue : undefined;
            resourceInputs["phase1AuthMethod"] = state ? state.phase1AuthMethod : undefined;
            resourceInputs["phase1EncryptAlgorithm"] = state ? state.phase1EncryptAlgorithm : undefined;
            resourceInputs["phase1HashAlgorithm"] = state ? state.phase1HashAlgorithm : undefined;
            resourceInputs["phase1PerfectForwardSecrecy"] = state ? state.phase1PerfectForwardSecrecy : undefined;
            resourceInputs["presharedKey"] = state ? state.presharedKey : undefined;
            resourceInputs["presharedKeyEncrypted"] = state ? state.presharedKeyEncrypted : undefined;
            resourceInputs["prf"] = state ? state.prf : undefined;
            resourceInputs["proxySupport"] = state ? state.proxySupport : undefined;
            resourceInputs["remoteAddress"] = state ? state.remoteAddress : undefined;
            resourceInputs["replayWindowSize"] = state ? state.replayWindowSize : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["trafficSelectors"] = state ? state.trafficSelectors : undefined;
            resourceInputs["verifyCert"] = state ? state.verifyCert : undefined;
            resourceInputs["versions"] = state ? state.versions : undefined;
        } else {
            const args = argsOrState as NetIkePeerArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.remoteAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteAddress'");
            }
            resourceInputs["appService"] = args ? args.appService : undefined;
            resourceInputs["caCertFile"] = args ? args.caCertFile : undefined;
            resourceInputs["crlFile"] = args ? args.crlFile : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dpdDelay"] = args ? args.dpdDelay : undefined;
            resourceInputs["generatePolicy"] = args ? args.generatePolicy : undefined;
            resourceInputs["lifetime"] = args ? args.lifetime : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["myCertFile"] = args ? args.myCertFile : undefined;
            resourceInputs["myCertKeyFile"] = args ? args.myCertKeyFile : undefined;
            resourceInputs["myCertKeyPassphrase"] = args ? args.myCertKeyPassphrase : undefined;
            resourceInputs["myIdType"] = args ? args.myIdType : undefined;
            resourceInputs["myIdValue"] = args ? args.myIdValue : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["natTraversal"] = args ? args.natTraversal : undefined;
            resourceInputs["passive"] = args ? args.passive : undefined;
            resourceInputs["peersCertFile"] = args ? args.peersCertFile : undefined;
            resourceInputs["peersCertType"] = args ? args.peersCertType : undefined;
            resourceInputs["peersIdType"] = args ? args.peersIdType : undefined;
            resourceInputs["peersIdValue"] = args ? args.peersIdValue : undefined;
            resourceInputs["phase1AuthMethod"] = args ? args.phase1AuthMethod : undefined;
            resourceInputs["phase1EncryptAlgorithm"] = args ? args.phase1EncryptAlgorithm : undefined;
            resourceInputs["phase1HashAlgorithm"] = args ? args.phase1HashAlgorithm : undefined;
            resourceInputs["phase1PerfectForwardSecrecy"] = args ? args.phase1PerfectForwardSecrecy : undefined;
            resourceInputs["presharedKey"] = args ? args.presharedKey : undefined;
            resourceInputs["presharedKeyEncrypted"] = args ? args.presharedKeyEncrypted : undefined;
            resourceInputs["prf"] = args ? args.prf : undefined;
            resourceInputs["proxySupport"] = args ? args.proxySupport : undefined;
            resourceInputs["remoteAddress"] = args ? args.remoteAddress : undefined;
            resourceInputs["replayWindowSize"] = args ? args.replayWindowSize : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["trafficSelectors"] = args ? args.trafficSelectors : undefined;
            resourceInputs["verifyCert"] = args ? args.verifyCert : undefined;
            resourceInputs["versions"] = args ? args.versions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetIkePeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetIkePeer resources.
 */
export interface NetIkePeerState {
    /**
     * The application service that the object belongs to
     */
    appService?: pulumi.Input<string>;
    /**
     * the trusted root and intermediate certificate authorities
     */
    caCertFile?: pulumi.Input<string>;
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     */
    crlFile?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     */
    dpdDelay?: pulumi.Input<number>;
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     */
    generatePolicy?: pulumi.Input<string>;
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     */
    lifetime?: pulumi.Input<number>;
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate file object
     */
    myCertFile?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate key file object
     */
    myCertKeyFile?: pulumi.Input<string>;
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     */
    myCertKeyPassphrase?: pulumi.Input<string>;
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     */
    myIdType?: pulumi.Input<string>;
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     */
    myIdValue?: pulumi.Input<string>;
    /**
     * Name of the ike_peer
     */
    name?: pulumi.Input<string>;
    /**
     * Enables use of the NAT-Traversal IPsec extension
     */
    natTraversal?: pulumi.Input<string>;
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     */
    passive?: pulumi.Input<string>;
    /**
     * Specifies the peer’s certificate for authentication
     */
    peersCertFile?: pulumi.Input<string>;
    /**
     * Specifies that the only peers-cert-type supported is certfile
     */
    peersCertType?: pulumi.Input<string>;
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     */
    peersIdType?: pulumi.Input<string>;
    /**
     * Specifies the peer’s identifier to be received
     */
    peersIdValue?: pulumi.Input<string>;
    /**
     * Specifies the authentication method used for phase 1 negotiation
     */
    phase1AuthMethod?: pulumi.Input<string>;
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     */
    phase1EncryptAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     */
    phase1HashAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     */
    phase1PerfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the preshared key for ISAKMP SAs
     */
    presharedKey?: pulumi.Input<string>;
    /**
     * Display the encrypted preshared-key for the IKE remote node
     */
    presharedKeyEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     */
    prf?: pulumi.Input<string>;
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     */
    proxySupport?: pulumi.Input<string>;
    /**
     * Specifies the IP address of the IKE remote node
     */
    remoteAddress?: pulumi.Input<string>;
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     */
    replayWindowSize?: pulumi.Input<number>;
    /**
     * Enables or disables this IKE remote node
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     */
    trafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     */
    verifyCert?: pulumi.Input<string>;
    /**
     * Specifies which version of IKE to be used
     */
    versions?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NetIkePeer resource.
 */
export interface NetIkePeerArgs {
    /**
     * The application service that the object belongs to
     */
    appService?: pulumi.Input<string>;
    /**
     * the trusted root and intermediate certificate authorities
     */
    caCertFile?: pulumi.Input<string>;
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     */
    crlFile?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     */
    dpdDelay?: pulumi.Input<number>;
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     */
    generatePolicy?: pulumi.Input<string>;
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     */
    lifetime?: pulumi.Input<number>;
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate file object
     */
    myCertFile?: pulumi.Input<string>;
    /**
     * Specifies the name of the certificate key file object
     */
    myCertKeyFile?: pulumi.Input<string>;
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     */
    myCertKeyPassphrase?: pulumi.Input<string>;
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     */
    myIdType?: pulumi.Input<string>;
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     */
    myIdValue?: pulumi.Input<string>;
    /**
     * Name of the ike_peer
     */
    name: pulumi.Input<string>;
    /**
     * Enables use of the NAT-Traversal IPsec extension
     */
    natTraversal?: pulumi.Input<string>;
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     */
    passive?: pulumi.Input<string>;
    /**
     * Specifies the peer’s certificate for authentication
     */
    peersCertFile?: pulumi.Input<string>;
    /**
     * Specifies that the only peers-cert-type supported is certfile
     */
    peersCertType?: pulumi.Input<string>;
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     */
    peersIdType?: pulumi.Input<string>;
    /**
     * Specifies the peer’s identifier to be received
     */
    peersIdValue?: pulumi.Input<string>;
    /**
     * Specifies the authentication method used for phase 1 negotiation
     */
    phase1AuthMethod?: pulumi.Input<string>;
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     */
    phase1EncryptAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     */
    phase1HashAlgorithm?: pulumi.Input<string>;
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     */
    phase1PerfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the preshared key for ISAKMP SAs
     */
    presharedKey?: pulumi.Input<string>;
    /**
     * Display the encrypted preshared-key for the IKE remote node
     */
    presharedKeyEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     */
    prf?: pulumi.Input<string>;
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     */
    proxySupport?: pulumi.Input<string>;
    /**
     * Specifies the IP address of the IKE remote node
     */
    remoteAddress: pulumi.Input<string>;
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     */
    replayWindowSize?: pulumi.Input<number>;
    /**
     * Enables or disables this IKE remote node
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     */
    trafficSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     */
    verifyCert?: pulumi.Input<string>;
    /**
     * Specifies which version of IKE to be used
     */
    versions?: pulumi.Input<pulumi.Input<string>[]>;
}
