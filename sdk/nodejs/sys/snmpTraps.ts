// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_sys_snmp_traps` provides details bout how to enable snmp_traps resource on BIG-IP
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const bigip_sys_snmp_traps_snmp_traps = new f5bigip.sys.SnmpTraps("snmp_traps", {
 *     community: "f5community",
 *     description: "Setup snmp traps",
 *     host: "195.10.10.1",
 *     name: "snmptraps",
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnmpTrapsState, opts?: pulumi.CustomResourceOptions): SnmpTraps {
        return new SnmpTraps(name, <any>state, { ...opts, id: id });
    }

    /**
     * Encrypted password
     */
    public readonly authPasswordencrypted: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    public readonly authProtocol: pulumi.Output<string | undefined>;
    /**
     * Specifies the community string used for this trap.
     */
    public readonly community: pulumi.Output<string | undefined>;
    /**
     * The port that the trap will be sent to.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    public readonly engineId: pulumi.Output<string | undefined>;
    /**
     * The host the trap will be sent to.
     */
    public readonly host: pulumi.Output<string | undefined>;
    /**
     * Name of the snmp trap.
     */
    public readonly name: pulumi.Output<string | undefined>;
    /**
     * User defined description.
     */
    public readonly port: pulumi.Output<number | undefined>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    public readonly privacyPassword: pulumi.Output<string | undefined>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    public readonly privacyPasswordEncrypted: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    public readonly privacyProtocol: pulumi.Output<string | undefined>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    public readonly securityLevel: pulumi.Output<string | undefined>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    public readonly securityName: pulumi.Output<string | undefined>;
    /**
     * SNMP version used for sending the trap.
     */
    public readonly version: pulumi.Output<string | undefined>;

    /**
     * Create a SnmpTraps resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnmpTrapsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnmpTrapsArgs | SnmpTrapsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SnmpTrapsState = argsOrState as SnmpTrapsState | undefined;
            inputs["authPasswordencrypted"] = state ? state.authPasswordencrypted : undefined;
            inputs["authProtocol"] = state ? state.authProtocol : undefined;
            inputs["community"] = state ? state.community : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["engineId"] = state ? state.engineId : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["privacyPassword"] = state ? state.privacyPassword : undefined;
            inputs["privacyPasswordEncrypted"] = state ? state.privacyPasswordEncrypted : undefined;
            inputs["privacyProtocol"] = state ? state.privacyProtocol : undefined;
            inputs["securityLevel"] = state ? state.securityLevel : undefined;
            inputs["securityName"] = state ? state.securityName : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SnmpTrapsArgs | undefined;
            inputs["authPasswordencrypted"] = args ? args.authPasswordencrypted : undefined;
            inputs["authProtocol"] = args ? args.authProtocol : undefined;
            inputs["community"] = args ? args.community : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["engineId"] = args ? args.engineId : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["privacyPassword"] = args ? args.privacyPassword : undefined;
            inputs["privacyPasswordEncrypted"] = args ? args.privacyPasswordEncrypted : undefined;
            inputs["privacyProtocol"] = args ? args.privacyProtocol : undefined;
            inputs["securityLevel"] = args ? args.securityLevel : undefined;
            inputs["securityName"] = args ? args.securityName : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        super("f5bigip:sys/snmpTraps:SnmpTraps", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnmpTraps resources.
 */
export interface SnmpTrapsState {
    /**
     * Encrypted password
     */
    readonly authPasswordencrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    readonly authProtocol?: pulumi.Input<string>;
    /**
     * Specifies the community string used for this trap.
     */
    readonly community?: pulumi.Input<string>;
    /**
     * The port that the trap will be sent to.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    readonly engineId?: pulumi.Input<string>;
    /**
     * The host the trap will be sent to.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Name of the snmp trap.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User defined description.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    readonly privacyPassword?: pulumi.Input<string>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    readonly privacyPasswordEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    readonly privacyProtocol?: pulumi.Input<string>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    readonly securityLevel?: pulumi.Input<string>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    readonly securityName?: pulumi.Input<string>;
    /**
     * SNMP version used for sending the trap.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnmpTraps resource.
 */
export interface SnmpTrapsArgs {
    /**
     * Encrypted password
     */
    readonly authPasswordencrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to authenticate the user.
     */
    readonly authProtocol?: pulumi.Input<string>;
    /**
     * Specifies the community string used for this trap.
     */
    readonly community?: pulumi.Input<string>;
    /**
     * The port that the trap will be sent to.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the authoritative security engine for SNMPv3.
     */
    readonly engineId?: pulumi.Input<string>;
    /**
     * The host the trap will be sent to.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Name of the snmp trap.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User defined description.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     */
    readonly privacyPassword?: pulumi.Input<string>;
    /**
     * Specifies the encrypted password used to encrypt traffic.
     */
    readonly privacyPasswordEncrypted?: pulumi.Input<string>;
    /**
     * Specifies the protocol used to encrypt traffic.
     */
    readonly privacyProtocol?: pulumi.Input<string>;
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     */
    readonly securityLevel?: pulumi.Input<string>;
    /**
     * Security name used in conjunction with SNMPv3.
     */
    readonly securityName?: pulumi.Input<string>;
    /**
     * SNMP version used for sending the trap.
     */
    readonly version?: pulumi.Input<string>;
}
