// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.VirtualServer` Configures Virtual Server
 *
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (example: `/Common/test-virtualserver` ) or `partition + directory + name` of the resource (example: `/Common/test/test-virtualserver` ).
 * When including directory in `fullpath` we have to make sure it is created in the given partition before using it.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const http = new f5bigip.ltm.VirtualServer("http", {
 *     name: "/Common/terraform_vs_http",
 *     destination: "10.12.12.12",
 *     port: 80,
 *     pool: "/Common/the-default-pool",
 * });
 * // A Virtual server with SSL enabled
 * const httpsVirtualServer = new f5bigip.ltm.VirtualServer("httpsVirtualServer", {
 *     name: "/Common/terraform_vs_https",
 *     destination: _var.vip_ip,
 *     description: "VirtualServer-test",
 *     port: 443,
 *     pool: _var.pool,
 *     profiles: [
 *         "/Common/tcp",
 *         "/Common/my-awesome-ssl-cert",
 *         "/Common/http",
 *     ],
 *     sourceAddressTranslation: "automap",
 *     translateAddress: "enabled",
 *     translatePort: "enabled",
 * });
 * // A Virtual server with separate client and server profiles
 * const httpsLtm_virtualServerVirtualServer = new f5bigip.ltm.VirtualServer("httpsLtm/virtualServerVirtualServer", {
 *     name: "/Common/terraform_vs_https",
 *     destination: "10.255.255.254",
 *     description: "VirtualServer-test",
 *     port: 443,
 *     clientProfiles: ["/Common/clientssl"],
 *     serverProfiles: ["/Common/serverssl"],
 *     securityLogProfiles: ["/Common/global-network"],
 *     sourceAddressTranslation: "automap",
 * });
 * ```
 */
export class VirtualServer extends pulumi.CustomResource {
    /**
     * Get an existing VirtualServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualServerState, opts?: pulumi.CustomResourceOptions): VirtualServer {
        return new VirtualServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/virtualServer:VirtualServer';

    /**
     * Returns true if the given object is an instance of VirtualServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualServer.__pulumiType;
    }

    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     */
    public readonly clientProfiles!: pulumi.Output<string[] | undefined>;
    public readonly defaultPersistenceProfile!: pulumi.Output<string | undefined>;
    /**
     * Description of Virtual server
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Destination IP
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    public readonly fallbackPersistenceProfile!: pulumi.Output<string>;
    /**
     * Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     */
    public readonly irules!: pulumi.Output<string[] | undefined>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    public readonly mask!: pulumi.Output<string>;
    /**
     * Name of the virtual server
     */
    public readonly name!: pulumi.Output<string>;
    public readonly perFlowRequestAccessPolicy!: pulumi.Output<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    public readonly persistenceProfiles!: pulumi.Output<string[] | undefined>;
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * Default pool name
     */
    public readonly pool!: pulumi.Output<string | undefined>;
    /**
     * Listen port for the virtual server
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     */
    public readonly profiles!: pulumi.Output<string[]>;
    /**
     * Specifies the log profile applied to the virtual server.
     */
    public readonly securityLogProfiles!: pulumi.Output<string[] | undefined>;
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    public readonly serverProfiles!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    public readonly snatpool!: pulumi.Output<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    public readonly sourceAddressTranslation!: pulumi.Output<string>;
    /**
     * Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    public readonly translateAddress!: pulumi.Output<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    public readonly translatePort!: pulumi.Output<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
     */
    public readonly vlans!: pulumi.Output<string[] | undefined>;
    /**
     * Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     */
    public readonly vlansEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VirtualServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualServerArgs | VirtualServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualServerState | undefined;
            resourceInputs["clientProfiles"] = state ? state.clientProfiles : undefined;
            resourceInputs["defaultPersistenceProfile"] = state ? state.defaultPersistenceProfile : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["fallbackPersistenceProfile"] = state ? state.fallbackPersistenceProfile : undefined;
            resourceInputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            resourceInputs["irules"] = state ? state.irules : undefined;
            resourceInputs["mask"] = state ? state.mask : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["perFlowRequestAccessPolicy"] = state ? state.perFlowRequestAccessPolicy : undefined;
            resourceInputs["persistenceProfiles"] = state ? state.persistenceProfiles : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["pool"] = state ? state.pool : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["profiles"] = state ? state.profiles : undefined;
            resourceInputs["securityLogProfiles"] = state ? state.securityLogProfiles : undefined;
            resourceInputs["serverProfiles"] = state ? state.serverProfiles : undefined;
            resourceInputs["snatpool"] = state ? state.snatpool : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceAddressTranslation"] = state ? state.sourceAddressTranslation : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["translateAddress"] = state ? state.translateAddress : undefined;
            resourceInputs["translatePort"] = state ? state.translatePort : undefined;
            resourceInputs["vlans"] = state ? state.vlans : undefined;
            resourceInputs["vlansEnabled"] = state ? state.vlansEnabled : undefined;
        } else {
            const args = argsOrState as VirtualServerArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            resourceInputs["clientProfiles"] = args ? args.clientProfiles : undefined;
            resourceInputs["defaultPersistenceProfile"] = args ? args.defaultPersistenceProfile : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["fallbackPersistenceProfile"] = args ? args.fallbackPersistenceProfile : undefined;
            resourceInputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            resourceInputs["irules"] = args ? args.irules : undefined;
            resourceInputs["mask"] = args ? args.mask : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["perFlowRequestAccessPolicy"] = args ? args.perFlowRequestAccessPolicy : undefined;
            resourceInputs["persistenceProfiles"] = args ? args.persistenceProfiles : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["pool"] = args ? args.pool : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["profiles"] = args ? args.profiles : undefined;
            resourceInputs["securityLogProfiles"] = args ? args.securityLogProfiles : undefined;
            resourceInputs["serverProfiles"] = args ? args.serverProfiles : undefined;
            resourceInputs["snatpool"] = args ? args.snatpool : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceAddressTranslation"] = args ? args.sourceAddressTranslation : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["translateAddress"] = args ? args.translateAddress : undefined;
            resourceInputs["translatePort"] = args ? args.translatePort : undefined;
            resourceInputs["vlans"] = args ? args.vlans : undefined;
            resourceInputs["vlansEnabled"] = args ? args.vlansEnabled : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualServer resources.
 */
export interface VirtualServerState {
    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     */
    clientProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    defaultPersistenceProfile?: pulumi.Input<string>;
    /**
     * Description of Virtual server
     */
    description?: pulumi.Input<string>;
    /**
     * Destination IP
     */
    destination?: pulumi.Input<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    fallbackPersistenceProfile?: pulumi.Input<string>;
    /**
     * Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     */
    irules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    mask?: pulumi.Input<string>;
    /**
     * Name of the virtual server
     */
    name?: pulumi.Input<string>;
    perFlowRequestAccessPolicy?: pulumi.Input<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    persistenceProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default pool name
     */
    pool?: pulumi.Input<string>;
    /**
     * Listen port for the virtual server
     */
    port?: pulumi.Input<number>;
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     */
    profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the log profile applied to the virtual server.
     */
    securityLogProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    serverProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    snatpool?: pulumi.Input<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    source?: pulumi.Input<string>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    sourceAddressTranslation?: pulumi.Input<string>;
    /**
     * Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     */
    state?: pulumi.Input<string>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    translateAddress?: pulumi.Input<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    translatePort?: pulumi.Input<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
     */
    vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     */
    vlansEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VirtualServer resource.
 */
export interface VirtualServerArgs {
    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     */
    clientProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    defaultPersistenceProfile?: pulumi.Input<string>;
    /**
     * Description of Virtual server
     */
    description?: pulumi.Input<string>;
    /**
     * Destination IP
     */
    destination: pulumi.Input<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    fallbackPersistenceProfile?: pulumi.Input<string>;
    /**
     * Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     */
    irules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    mask?: pulumi.Input<string>;
    /**
     * Name of the virtual server
     */
    name: pulumi.Input<string>;
    perFlowRequestAccessPolicy?: pulumi.Input<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    persistenceProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default pool name
     */
    pool?: pulumi.Input<string>;
    /**
     * Listen port for the virtual server
     */
    port: pulumi.Input<number>;
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     */
    profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the log profile applied to the virtual server.
     */
    securityLogProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    serverProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    snatpool?: pulumi.Input<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    source?: pulumi.Input<string>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    sourceAddressTranslation?: pulumi.Input<string>;
    /**
     * Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     */
    state?: pulumi.Input<string>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    translateAddress?: pulumi.Input<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    translatePort?: pulumi.Input<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
     */
    vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     */
    vlansEnabled?: pulumi.Input<boolean>;
}
