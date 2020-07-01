// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configures a cookie persistence profile
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const dstaddr = new f5bigip.ltm.PersistenceProfileDstAddr("dstaddr", {
 *     defaultsFrom: "/Common/dest_addr",
 *     hashAlgorithm: "carp",
 *     mask: "255.255.255.255",
 *     matchAcrossPools: "enabled",
 *     matchAcrossServices: "enabled",
 *     matchAcrossVirtuals: "enabled",
 *     mirror: "enabled",
 *     name: "/Common/terraform_ppdstaddr",
 *     overrideConnLimit: "enabled",
 *     timeout: 3600,
 * });
 * ```
 *
 * ## Reference
 *
 * `name` - (Required) Name of the virtual address
 *
 * `defaultsFrom` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.
 *
 * `matchAcrossPools` (Optional) (enabled or disabled) match across pools with given persistence record
 *
 * `matchAcrossServices` (Optional) (enabled or disabled) match across services with given persistence record
 *
 * `matchAcrossVirtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
 *
 * `mirror` (Optional) (enabled or disabled) mirror persistence record
 *
 * `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
 *
 * `overrideConnLimit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
 */
export class PersistenceProfileDstAddr extends pulumi.CustomResource {
    /**
     * Get an existing PersistenceProfileDstAddr resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PersistenceProfileDstAddrState, opts?: pulumi.CustomResourceOptions): PersistenceProfileDstAddr {
        return new PersistenceProfileDstAddr(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr';

    /**
     * Returns true if the given object is an instance of PersistenceProfileDstAddr.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PersistenceProfileDstAddr {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PersistenceProfileDstAddr.__pulumiType;
    }

    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * Inherit defaults from parent profile
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Specify the hash algorithm
     */
    public readonly hashAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     */
    public readonly mask!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    public readonly matchAcrossPools!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    public readonly matchAcrossServices!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    public readonly matchAcrossVirtuals!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable
     */
    public readonly mirror!: pulumi.Output<string | undefined>;
    /**
     * Name of the persistence profile
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    public readonly overrideConnLimit!: pulumi.Output<string | undefined>;
    /**
     * Timeout for persistence of the session
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a PersistenceProfileDstAddr resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PersistenceProfileDstAddrArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PersistenceProfileDstAddrArgs | PersistenceProfileDstAddrState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PersistenceProfileDstAddrState | undefined;
            inputs["appService"] = state ? state.appService : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["hashAlgorithm"] = state ? state.hashAlgorithm : undefined;
            inputs["mask"] = state ? state.mask : undefined;
            inputs["matchAcrossPools"] = state ? state.matchAcrossPools : undefined;
            inputs["matchAcrossServices"] = state ? state.matchAcrossServices : undefined;
            inputs["matchAcrossVirtuals"] = state ? state.matchAcrossVirtuals : undefined;
            inputs["mirror"] = state ? state.mirror : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overrideConnLimit"] = state ? state.overrideConnLimit : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as PersistenceProfileDstAddrArgs | undefined;
            if (!args || args.defaultsFrom === undefined) {
                throw new Error("Missing required property 'defaultsFrom'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["appService"] = args ? args.appService : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["hashAlgorithm"] = args ? args.hashAlgorithm : undefined;
            inputs["mask"] = args ? args.mask : undefined;
            inputs["matchAcrossPools"] = args ? args.matchAcrossPools : undefined;
            inputs["matchAcrossServices"] = args ? args.matchAcrossServices : undefined;
            inputs["matchAcrossVirtuals"] = args ? args.matchAcrossVirtuals : undefined;
            inputs["mirror"] = args ? args.mirror : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overrideConnLimit"] = args ? args.overrideConnLimit : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PersistenceProfileDstAddr.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PersistenceProfileDstAddr resources.
 */
export interface PersistenceProfileDstAddrState {
    readonly appService?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specify the hash algorithm
     */
    readonly hashAlgorithm?: pulumi.Input<string>;
    /**
     * Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     */
    readonly mask?: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossVirtuals?: pulumi.Input<string>;
    /**
     * To enable _ disable
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the persistence profile
     */
    readonly name?: pulumi.Input<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    readonly overrideConnLimit?: pulumi.Input<string>;
    /**
     * Timeout for persistence of the session
     */
    readonly timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PersistenceProfileDstAddr resource.
 */
export interface PersistenceProfileDstAddrArgs {
    readonly appService?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom: pulumi.Input<string>;
    /**
     * Specify the hash algorithm
     */
    readonly hashAlgorithm?: pulumi.Input<string>;
    /**
     * Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     */
    readonly mask?: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossVirtuals?: pulumi.Input<string>;
    /**
     * To enable _ disable
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the persistence profile
     */
    readonly name: pulumi.Input<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    readonly overrideConnLimit?: pulumi.Input<string>;
    /**
     * Timeout for persistence of the session
     */
    readonly timeout?: pulumi.Input<number>;
}
