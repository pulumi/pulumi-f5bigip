// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.SnmpTraps` provides details bout how to enable snmpTraps resource on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const snmpTraps = new f5bigip.sys.SnmpTraps("snmp_traps", {
 *     name: "snmptraps",
 *     community: "f5community",
 *     host: "195.10.10.1",
 *     description: "Setup snmp traps",
 *     port: 111,
 * });
 * ```
 */
export class SnmpTraps extends pulumi.CustomResource {
    /**
     * Get an existing SnmpTraps resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnmpTrapsState, opts?: pulumi.CustomResourceOptions): SnmpTraps {
        return new SnmpTraps(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/snmpTraps:SnmpTraps';

    /**
     * Returns true if the given object is an instance of SnmpTraps.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnmpTraps {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnmpTraps.__pulumiType;
    }

    /**
     * Encrypted password
     */
    public readonly authPasswordencrypted!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    public readonly authProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the community string used for this trap.
     */
    public readonly community!: pulumi.Output<string | undefined>;
    /**
     * The port that the trap will be sent to.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    public readonly engineId!: pulumi.Output<string | undefined>;
    /**
     * The host the trap will be sent to.
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * Name of the snmp trap.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * User defined description.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    public readonly privacyPassword!: pulumi.Output<string | undefined>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    public readonly privacyPasswordEncrypted!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    public readonly privacyProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    public readonly securityLevel!: pulumi.Output<string>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    public readonly securityName!: pulumi.Output<string | undefined>;
    /**
     * SNMP version used for sending the trap.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a SnmpTraps resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnmpTrapsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnmpTrapsArgs | SnmpTrapsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnmpTrapsState | undefined;
            resourceInputs["authPasswordencrypted"] = state ? state.authPasswordencrypted : undefined;
            resourceInputs["authProtocol"] = state ? state.authProtocol : undefined;
            resourceInputs["community"] = state ? state.community : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engineId"] = state ? state.engineId : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privacyPassword"] = state ? state.privacyPassword : undefined;
            resourceInputs["privacyPasswordEncrypted"] = state ? state.privacyPasswordEncrypted : undefined;
            resourceInputs["privacyProtocol"] = state ? state.privacyProtocol : undefined;
            resourceInputs["securityLevel"] = state ? state.securityLevel : undefined;
            resourceInputs["securityName"] = state ? state.securityName : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SnmpTrapsArgs | undefined;
            resourceInputs["authPasswordencrypted"] = args ? args.authPasswordencrypted : undefined;
            resourceInputs["authProtocol"] = args ? args.authProtocol : undefined;
            resourceInputs["community"] = args ? args.community : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineId"] = args ? args.engineId : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["privacyPassword"] = args ? args.privacyPassword : undefined;
            resourceInputs["privacyPasswordEncrypted"] = args ? args.privacyPasswordEncrypted : undefined;
            resourceInputs["privacyProtocol"] = args ? args.privacyProtocol : undefined;
            resourceInputs["securityLevel"] = args ? args.securityLevel : undefined;
            resourceInputs["securityName"] = args ? args.securityName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnmpTraps.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnmpTraps resources.
 */
export interface SnmpTrapsState {
    /**
     * Encrypted password
     */
    authPasswordencrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    authProtocol?: pulumi.Input<string>;
    /**
     * Specifies the community string used for this trap.
     */
    community?: pulumi.Input<string>;
    /**
     * The port that the trap will be sent to.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    engineId?: pulumi.Input<string>;
    /**
     * The host the trap will be sent to.
     */
    host?: pulumi.Input<string>;
    /**
     * Name of the snmp trap.
     */
    name?: pulumi.Input<string>;
    /**
     * User defined description.
     */
    port?: pulumi.Input<number>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    privacyPassword?: pulumi.Input<string>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    privacyPasswordEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    privacyProtocol?: pulumi.Input<string>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    securityLevel?: pulumi.Input<string>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    securityName?: pulumi.Input<string>;
    /**
     * SNMP version used for sending the trap.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnmpTraps resource.
 */
export interface SnmpTrapsArgs {
    /**
     * Encrypted password
     */
    authPasswordencrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    authProtocol?: pulumi.Input<string>;
    /**
     * Specifies the community string used for this trap.
     */
    community?: pulumi.Input<string>;
    /**
     * The port that the trap will be sent to.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    engineId?: pulumi.Input<string>;
    /**
     * The host the trap will be sent to.
     */
    host?: pulumi.Input<string>;
    /**
     * Name of the snmp trap.
     */
    name?: pulumi.Input<string>;
    /**
     * User defined description.
     */
    port?: pulumi.Input<number>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    privacyPassword?: pulumi.Input<string>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    privacyPasswordEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    privacyProtocol?: pulumi.Input<string>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    securityLevel?: pulumi.Input<string>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    securityName?: pulumi.Input<string>;
    /**
     * SNMP version used for sending the trap.
     */
    version?: pulumi.Input<string>;
}
