// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileServerSsl` Manages server SSL profiles on a BIG-IP
 *
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_ServerSsl = new f5bigip.ltm.ProfileServerSsl("test-ServerSsl", {
 *     authenticate: "always",
 *     ciphers: "DEFAULT",
 *     defaultsFrom: "/Common/serverssl",
 *     name: "/Common/test-ServerSsl",
 *     partition: "Common",
 * });
 * ```
 */
export class ProfileServerSsl extends pulumi.CustomResource {
    /**
     * Get an existing ProfileServerSsl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileServerSslState, opts?: pulumi.CustomResourceOptions): ProfileServerSsl {
        return new ProfileServerSsl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileServerSsl:ProfileServerSsl';

    /**
     * Returns true if the given object is an instance of ProfileServerSsl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileServerSsl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileServerSsl.__pulumiType;
    }

    /**
     * Alert time out
     */
    public readonly alertTimeout!: pulumi.Output<string>;
    /**
     * Server authentication once / always (default is once).
     */
    public readonly authenticate!: pulumi.Output<string>;
    /**
     * Client certificate chain traversal depth. Default 9.
     */
    public readonly authenticateDepth!: pulumi.Output<number>;
    /**
     * Client certificate file path. Default None.
     */
    public readonly caFile!: pulumi.Output<string>;
    /**
     * Cache size (sessions).
     */
    public readonly cacheSize!: pulumi.Output<number>;
    /**
     * Cache time out
     */
    public readonly cacheTimeout!: pulumi.Output<number>;
    /**
     * Specifies the name of the certificate that the system uses for server-side SSL processing.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Specifies the certificates-key chain to associate with the SSL profile
     */
    public readonly chain!: pulumi.Output<string>;
    /**
     * Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
     */
    public readonly ciphers!: pulumi.Output<string>;
    /**
     * The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Response if the cert is expired (drop / ignore).
     */
    public readonly expireCertResponseControl!: pulumi.Output<string>;
    /**
     * full path of the profile
     */
    public readonly fullPath!: pulumi.Output<string>;
    /**
     * generation
     */
    public readonly generation!: pulumi.Output<number>;
    /**
     * Generic alerts enabled / disabled.
     */
    public readonly genericAlert!: pulumi.Output<string>;
    /**
     * Handshake time out (seconds)
     */
    public readonly handshakeTimeout!: pulumi.Output<string>;
    /**
     * Specifies the file name of the SSL key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    public readonly modSslMethods!: pulumi.Output<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Specifies the name of the profile. (type `string`)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Device partition to manage resources on.
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * Client Certificate Constrained Delegation CA passphrase
     */
    public readonly passphrase!: pulumi.Output<string>;
    /**
     * Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
     */
    public readonly peerCertMode!: pulumi.Output<string>;
    /**
     * Proxy SSL enabled / disabled. Default is disabled.
     */
    public readonly proxySsl!: pulumi.Output<string>;
    /**
     * Renogotiate Period (seconds)
     */
    public readonly renegotiatePeriod!: pulumi.Output<string>;
    /**
     * Renogotiate Size
     */
    public readonly renegotiateSize!: pulumi.Output<string>;
    /**
     * Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
     */
    public readonly renegotiation!: pulumi.Output<string>;
    /**
     * When `true`, client certificate is retained in SSL session.
     */
    public readonly retainCertificate!: pulumi.Output<string>;
    /**
     * Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
     * When `request` is set the system request secure renegotation of SSL connections.
     * `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
     * The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
     */
    public readonly secureRenegotiation!: pulumi.Output<string>;
    /**
     * Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Session Mirroring (enabled / disabled)
     */
    public readonly sessionMirroring!: pulumi.Output<string>;
    /**
     * Session Ticket (enabled / disabled)
     */
    public readonly sessionTicket!: pulumi.Output<string>;
    /**
     * Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
     * There can be only one SSL profile with this setting enabled.
     */
    public readonly sniDefault!: pulumi.Output<string>;
    /**
     * Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
     */
    public readonly sniRequire!: pulumi.Output<string>;
    /**
     * SSL forward Proxy (enabled / disabled)
     */
    public readonly sslForwardProxy!: pulumi.Output<string>;
    /**
     * SSL forward Proxy Bypass (enabled / disabled)
     */
    public readonly sslForwardProxyBypass!: pulumi.Output<string>;
    /**
     * SSL sign hash (any, sha1, sha256, sha384)
     */
    public readonly sslSignHash!: pulumi.Output<string>;
    /**
     * Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
     */
    public readonly strictResume!: pulumi.Output<string>;
    public readonly tmOptions!: pulumi.Output<string[]>;
    /**
     * Unclean Shutdown (enabled / disabled)
     */
    public readonly uncleanShutdown!: pulumi.Output<string>;
    /**
     * Unclean Shutdown (drop / ignore)
     */
    public readonly untrustedCertResponseControl!: pulumi.Output<string>;

    /**
     * Create a ProfileServerSsl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileServerSslArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileServerSslArgs | ProfileServerSslState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileServerSslState | undefined;
            inputs["alertTimeout"] = state ? state.alertTimeout : undefined;
            inputs["authenticate"] = state ? state.authenticate : undefined;
            inputs["authenticateDepth"] = state ? state.authenticateDepth : undefined;
            inputs["caFile"] = state ? state.caFile : undefined;
            inputs["cacheSize"] = state ? state.cacheSize : undefined;
            inputs["cacheTimeout"] = state ? state.cacheTimeout : undefined;
            inputs["cert"] = state ? state.cert : undefined;
            inputs["chain"] = state ? state.chain : undefined;
            inputs["ciphers"] = state ? state.ciphers : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["expireCertResponseControl"] = state ? state.expireCertResponseControl : undefined;
            inputs["fullPath"] = state ? state.fullPath : undefined;
            inputs["generation"] = state ? state.generation : undefined;
            inputs["genericAlert"] = state ? state.genericAlert : undefined;
            inputs["handshakeTimeout"] = state ? state.handshakeTimeout : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["modSslMethods"] = state ? state.modSslMethods : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
            inputs["passphrase"] = state ? state.passphrase : undefined;
            inputs["peerCertMode"] = state ? state.peerCertMode : undefined;
            inputs["proxySsl"] = state ? state.proxySsl : undefined;
            inputs["renegotiatePeriod"] = state ? state.renegotiatePeriod : undefined;
            inputs["renegotiateSize"] = state ? state.renegotiateSize : undefined;
            inputs["renegotiation"] = state ? state.renegotiation : undefined;
            inputs["retainCertificate"] = state ? state.retainCertificate : undefined;
            inputs["secureRenegotiation"] = state ? state.secureRenegotiation : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["sessionMirroring"] = state ? state.sessionMirroring : undefined;
            inputs["sessionTicket"] = state ? state.sessionTicket : undefined;
            inputs["sniDefault"] = state ? state.sniDefault : undefined;
            inputs["sniRequire"] = state ? state.sniRequire : undefined;
            inputs["sslForwardProxy"] = state ? state.sslForwardProxy : undefined;
            inputs["sslForwardProxyBypass"] = state ? state.sslForwardProxyBypass : undefined;
            inputs["sslSignHash"] = state ? state.sslSignHash : undefined;
            inputs["strictResume"] = state ? state.strictResume : undefined;
            inputs["tmOptions"] = state ? state.tmOptions : undefined;
            inputs["uncleanShutdown"] = state ? state.uncleanShutdown : undefined;
            inputs["untrustedCertResponseControl"] = state ? state.untrustedCertResponseControl : undefined;
        } else {
            const args = argsOrState as ProfileServerSslArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["alertTimeout"] = args ? args.alertTimeout : undefined;
            inputs["authenticate"] = args ? args.authenticate : undefined;
            inputs["authenticateDepth"] = args ? args.authenticateDepth : undefined;
            inputs["caFile"] = args ? args.caFile : undefined;
            inputs["cacheSize"] = args ? args.cacheSize : undefined;
            inputs["cacheTimeout"] = args ? args.cacheTimeout : undefined;
            inputs["cert"] = args ? args.cert : undefined;
            inputs["chain"] = args ? args.chain : undefined;
            inputs["ciphers"] = args ? args.ciphers : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["expireCertResponseControl"] = args ? args.expireCertResponseControl : undefined;
            inputs["fullPath"] = args ? args.fullPath : undefined;
            inputs["generation"] = args ? args.generation : undefined;
            inputs["genericAlert"] = args ? args.genericAlert : undefined;
            inputs["handshakeTimeout"] = args ? args.handshakeTimeout : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["modSslMethods"] = args ? args.modSslMethods : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
            inputs["passphrase"] = args ? args.passphrase : undefined;
            inputs["peerCertMode"] = args ? args.peerCertMode : undefined;
            inputs["proxySsl"] = args ? args.proxySsl : undefined;
            inputs["renegotiatePeriod"] = args ? args.renegotiatePeriod : undefined;
            inputs["renegotiateSize"] = args ? args.renegotiateSize : undefined;
            inputs["renegotiation"] = args ? args.renegotiation : undefined;
            inputs["retainCertificate"] = args ? args.retainCertificate : undefined;
            inputs["secureRenegotiation"] = args ? args.secureRenegotiation : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sessionMirroring"] = args ? args.sessionMirroring : undefined;
            inputs["sessionTicket"] = args ? args.sessionTicket : undefined;
            inputs["sniDefault"] = args ? args.sniDefault : undefined;
            inputs["sniRequire"] = args ? args.sniRequire : undefined;
            inputs["sslForwardProxy"] = args ? args.sslForwardProxy : undefined;
            inputs["sslForwardProxyBypass"] = args ? args.sslForwardProxyBypass : undefined;
            inputs["sslSignHash"] = args ? args.sslSignHash : undefined;
            inputs["strictResume"] = args ? args.strictResume : undefined;
            inputs["tmOptions"] = args ? args.tmOptions : undefined;
            inputs["uncleanShutdown"] = args ? args.uncleanShutdown : undefined;
            inputs["untrustedCertResponseControl"] = args ? args.untrustedCertResponseControl : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileServerSsl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileServerSsl resources.
 */
export interface ProfileServerSslState {
    /**
     * Alert time out
     */
    readonly alertTimeout?: pulumi.Input<string>;
    /**
     * Server authentication once / always (default is once).
     */
    readonly authenticate?: pulumi.Input<string>;
    /**
     * Client certificate chain traversal depth. Default 9.
     */
    readonly authenticateDepth?: pulumi.Input<number>;
    /**
     * Client certificate file path. Default None.
     */
    readonly caFile?: pulumi.Input<string>;
    /**
     * Cache size (sessions).
     */
    readonly cacheSize?: pulumi.Input<number>;
    /**
     * Cache time out
     */
    readonly cacheTimeout?: pulumi.Input<number>;
    /**
     * Specifies the name of the certificate that the system uses for server-side SSL processing.
     */
    readonly cert?: pulumi.Input<string>;
    /**
     * Specifies the certificates-key chain to associate with the SSL profile
     */
    readonly chain?: pulumi.Input<string>;
    /**
     * Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
     */
    readonly ciphers?: pulumi.Input<string>;
    /**
     * The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Response if the cert is expired (drop / ignore).
     */
    readonly expireCertResponseControl?: pulumi.Input<string>;
    /**
     * full path of the profile
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * generation
     */
    readonly generation?: pulumi.Input<number>;
    /**
     * Generic alerts enabled / disabled.
     */
    readonly genericAlert?: pulumi.Input<string>;
    /**
     * Handshake time out (seconds)
     */
    readonly handshakeTimeout?: pulumi.Input<string>;
    /**
     * Specifies the file name of the SSL key.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    readonly modSslMethods?: pulumi.Input<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the profile. (type `string`)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Device partition to manage resources on.
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Client Certificate Constrained Delegation CA passphrase
     */
    readonly passphrase?: pulumi.Input<string>;
    /**
     * Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
     */
    readonly peerCertMode?: pulumi.Input<string>;
    /**
     * Proxy SSL enabled / disabled. Default is disabled.
     */
    readonly proxySsl?: pulumi.Input<string>;
    /**
     * Renogotiate Period (seconds)
     */
    readonly renegotiatePeriod?: pulumi.Input<string>;
    /**
     * Renogotiate Size
     */
    readonly renegotiateSize?: pulumi.Input<string>;
    /**
     * Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
     */
    readonly renegotiation?: pulumi.Input<string>;
    /**
     * When `true`, client certificate is retained in SSL session.
     */
    readonly retainCertificate?: pulumi.Input<string>;
    /**
     * Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
     * When `request` is set the system request secure renegotation of SSL connections.
     * `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
     * The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
     */
    readonly secureRenegotiation?: pulumi.Input<string>;
    /**
     * Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * Session Mirroring (enabled / disabled)
     */
    readonly sessionMirroring?: pulumi.Input<string>;
    /**
     * Session Ticket (enabled / disabled)
     */
    readonly sessionTicket?: pulumi.Input<string>;
    /**
     * Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
     * There can be only one SSL profile with this setting enabled.
     */
    readonly sniDefault?: pulumi.Input<string>;
    /**
     * Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
     */
    readonly sniRequire?: pulumi.Input<string>;
    /**
     * SSL forward Proxy (enabled / disabled)
     */
    readonly sslForwardProxy?: pulumi.Input<string>;
    /**
     * SSL forward Proxy Bypass (enabled / disabled)
     */
    readonly sslForwardProxyBypass?: pulumi.Input<string>;
    /**
     * SSL sign hash (any, sha1, sha256, sha384)
     */
    readonly sslSignHash?: pulumi.Input<string>;
    /**
     * Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
     */
    readonly strictResume?: pulumi.Input<string>;
    readonly tmOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unclean Shutdown (enabled / disabled)
     */
    readonly uncleanShutdown?: pulumi.Input<string>;
    /**
     * Unclean Shutdown (drop / ignore)
     */
    readonly untrustedCertResponseControl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileServerSsl resource.
 */
export interface ProfileServerSslArgs {
    /**
     * Alert time out
     */
    readonly alertTimeout?: pulumi.Input<string>;
    /**
     * Server authentication once / always (default is once).
     */
    readonly authenticate?: pulumi.Input<string>;
    /**
     * Client certificate chain traversal depth. Default 9.
     */
    readonly authenticateDepth?: pulumi.Input<number>;
    /**
     * Client certificate file path. Default None.
     */
    readonly caFile?: pulumi.Input<string>;
    /**
     * Cache size (sessions).
     */
    readonly cacheSize?: pulumi.Input<number>;
    /**
     * Cache time out
     */
    readonly cacheTimeout?: pulumi.Input<number>;
    /**
     * Specifies the name of the certificate that the system uses for server-side SSL processing.
     */
    readonly cert?: pulumi.Input<string>;
    /**
     * Specifies the certificates-key chain to associate with the SSL profile
     */
    readonly chain?: pulumi.Input<string>;
    /**
     * Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
     */
    readonly ciphers?: pulumi.Input<string>;
    /**
     * The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Response if the cert is expired (drop / ignore).
     */
    readonly expireCertResponseControl?: pulumi.Input<string>;
    /**
     * full path of the profile
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * generation
     */
    readonly generation?: pulumi.Input<number>;
    /**
     * Generic alerts enabled / disabled.
     */
    readonly genericAlert?: pulumi.Input<string>;
    /**
     * Handshake time out (seconds)
     */
    readonly handshakeTimeout?: pulumi.Input<string>;
    /**
     * Specifies the file name of the SSL key.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    readonly modSslMethods?: pulumi.Input<string>;
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Specifies the name of the profile. (type `string`)
     */
    readonly name: pulumi.Input<string>;
    /**
     * Device partition to manage resources on.
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Client Certificate Constrained Delegation CA passphrase
     */
    readonly passphrase?: pulumi.Input<string>;
    /**
     * Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
     */
    readonly peerCertMode?: pulumi.Input<string>;
    /**
     * Proxy SSL enabled / disabled. Default is disabled.
     */
    readonly proxySsl?: pulumi.Input<string>;
    /**
     * Renogotiate Period (seconds)
     */
    readonly renegotiatePeriod?: pulumi.Input<string>;
    /**
     * Renogotiate Size
     */
    readonly renegotiateSize?: pulumi.Input<string>;
    /**
     * Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
     */
    readonly renegotiation?: pulumi.Input<string>;
    /**
     * When `true`, client certificate is retained in SSL session.
     */
    readonly retainCertificate?: pulumi.Input<string>;
    /**
     * Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
     * When `request` is set the system request secure renegotation of SSL connections.
     * `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
     * The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
     */
    readonly secureRenegotiation?: pulumi.Input<string>;
    /**
     * Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * Session Mirroring (enabled / disabled)
     */
    readonly sessionMirroring?: pulumi.Input<string>;
    /**
     * Session Ticket (enabled / disabled)
     */
    readonly sessionTicket?: pulumi.Input<string>;
    /**
     * Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
     * There can be only one SSL profile with this setting enabled.
     */
    readonly sniDefault?: pulumi.Input<string>;
    /**
     * Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
     */
    readonly sniRequire?: pulumi.Input<string>;
    /**
     * SSL forward Proxy (enabled / disabled)
     */
    readonly sslForwardProxy?: pulumi.Input<string>;
    /**
     * SSL forward Proxy Bypass (enabled / disabled)
     */
    readonly sslForwardProxyBypass?: pulumi.Input<string>;
    /**
     * SSL sign hash (any, sha1, sha256, sha384)
     */
    readonly sslSignHash?: pulumi.Input<string>;
    /**
     * Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
     */
    readonly strictResume?: pulumi.Input<string>;
    readonly tmOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unclean Shutdown (enabled / disabled)
     */
    readonly uncleanShutdown?: pulumi.Input<string>;
    /**
     * Unclean Shutdown (drop / ignore)
     */
    readonly untrustedCertResponseControl?: pulumi.Input<string>;
}
