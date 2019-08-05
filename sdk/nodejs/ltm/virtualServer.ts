// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_virtual_server.html.markdown.
 */
export class VirtualServer extends pulumi.CustomResource {
    /**
     * Get an existing VirtualServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
    public readonly clientProfiles!: pulumi.Output<string[]>;
    /**
     * Destination IP
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    public readonly fallbackPersistenceProfile!: pulumi.Output<string>;
    /**
     * all, tcp, udp
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    public readonly irules!: pulumi.Output<string[] | undefined>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    public readonly mask!: pulumi.Output<string | undefined>;
    /**
     * Name of the virtual server
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    public readonly persistenceProfiles!: pulumi.Output<string[]>;
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
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    public readonly serverProfiles!: pulumi.Output<string[]>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    public readonly snatpool!: pulumi.Output<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    public readonly sourceAddressTranslation!: pulumi.Output<string>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    public readonly translateAddress!: pulumi.Output<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    public readonly translatePort!: pulumi.Output<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
     */
    public readonly vlans!: pulumi.Output<string[] | undefined>;
    /**
     * Enables the virtual server on the VLANs specified by the VLANs option.
     */
    public readonly vlansEnabled!: pulumi.Output<boolean>;

    /**
     * Create a VirtualServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualServerArgs | VirtualServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualServerState | undefined;
            inputs["clientProfiles"] = state ? state.clientProfiles : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["fallbackPersistenceProfile"] = state ? state.fallbackPersistenceProfile : undefined;
            inputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            inputs["irules"] = state ? state.irules : undefined;
            inputs["mask"] = state ? state.mask : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["persistenceProfiles"] = state ? state.persistenceProfiles : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["pool"] = state ? state.pool : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["profiles"] = state ? state.profiles : undefined;
            inputs["serverProfiles"] = state ? state.serverProfiles : undefined;
            inputs["snatpool"] = state ? state.snatpool : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourceAddressTranslation"] = state ? state.sourceAddressTranslation : undefined;
            inputs["translateAddress"] = state ? state.translateAddress : undefined;
            inputs["translatePort"] = state ? state.translatePort : undefined;
            inputs["vlans"] = state ? state.vlans : undefined;
            inputs["vlansEnabled"] = state ? state.vlansEnabled : undefined;
        } else {
            const args = argsOrState as VirtualServerArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            inputs["clientProfiles"] = args ? args.clientProfiles : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["fallbackPersistenceProfile"] = args ? args.fallbackPersistenceProfile : undefined;
            inputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            inputs["irules"] = args ? args.irules : undefined;
            inputs["mask"] = args ? args.mask : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["persistenceProfiles"] = args ? args.persistenceProfiles : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["pool"] = args ? args.pool : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["profiles"] = args ? args.profiles : undefined;
            inputs["serverProfiles"] = args ? args.serverProfiles : undefined;
            inputs["snatpool"] = args ? args.snatpool : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourceAddressTranslation"] = args ? args.sourceAddressTranslation : undefined;
            inputs["translateAddress"] = args ? args.translateAddress : undefined;
            inputs["translatePort"] = args ? args.translatePort : undefined;
            inputs["vlans"] = args ? args.vlans : undefined;
            inputs["vlansEnabled"] = args ? args.vlansEnabled : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualServer resources.
 */
export interface VirtualServerState {
    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     */
    readonly clientProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination IP
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    readonly fallbackPersistenceProfile?: pulumi.Input<string>;
    /**
     * all, tcp, udp
     */
    readonly ipProtocol?: pulumi.Input<string>;
    readonly irules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    readonly mask?: pulumi.Input<string>;
    /**
     * Name of the virtual server
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    readonly persistenceProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default pool name
     */
    readonly pool?: pulumi.Input<string>;
    /**
     * Listen port for the virtual server
     */
    readonly port?: pulumi.Input<number>;
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     */
    readonly profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    readonly serverProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    readonly snatpool?: pulumi.Input<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    readonly sourceAddressTranslation?: pulumi.Input<string>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    readonly translateAddress?: pulumi.Input<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    readonly translatePort?: pulumi.Input<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
     */
    readonly vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables the virtual server on the VLANs specified by the VLANs option.
     */
    readonly vlansEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VirtualServer resource.
 */
export interface VirtualServerArgs {
    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     */
    readonly clientProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination IP
     */
    readonly destination: pulumi.Input<string>;
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     */
    readonly fallbackPersistenceProfile?: pulumi.Input<string>;
    /**
     * all, tcp, udp
     */
    readonly ipProtocol?: pulumi.Input<string>;
    readonly irules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     */
    readonly mask?: pulumi.Input<string>;
    /**
     * Name of the virtual server
     */
    readonly name: pulumi.Input<string>;
    /**
     * List of persistence profiles associated with the Virtual Server.
     */
    readonly persistenceProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default pool name
     */
    readonly pool?: pulumi.Input<string>;
    /**
     * Listen port for the virtual server
     */
    readonly port: pulumi.Input<number>;
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     */
    readonly profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     */
    readonly serverProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
     */
    readonly snatpool?: pulumi.Input<string>;
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * Can be either omitted for none or the values automap or snat
     */
    readonly sourceAddressTranslation?: pulumi.Input<string>;
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     */
    readonly translateAddress?: pulumi.Input<string>;
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     */
    readonly translatePort?: pulumi.Input<string>;
    /**
     * The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
     */
    readonly vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables the virtual server on the VLANs specified by the VLANs option.
     */
    readonly vlansEnabled?: pulumi.Input<boolean>;
}
