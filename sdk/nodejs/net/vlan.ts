// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.net.Vlan` Manages a vlan configuration
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const vlan1 = new f5bigip.net.Vlan("vlan1", {
 *     name: "/Common/Internal",
 *     tag: 101,
 *     interfaces: [{
 *         vlanport: "1.2",
 *         tagged: false,
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Vlan extends pulumi.CustomResource {
    /**
     * Get an existing Vlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VlanState, opts?: pulumi.CustomResourceOptions): Vlan {
        return new Vlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:net/vlan:Vlan';

    /**
     * Returns true if the given object is an instance of Vlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vlan.__pulumiType;
    }

    /**
     * Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [`default`, `src-ip`, `dst-ip`]
     */
    public readonly cmpHash!: pulumi.Output<string>;
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    public readonly interfaces!: pulumi.Output<outputs.net.VlanInterface[] | undefined>;
    /**
     * Specifies the maximum transmission unit (MTU) for traffic on this VLAN. The default value is `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * Name of the vlan
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     */
    public readonly tag!: pulumi.Output<number | undefined>;

    /**
     * Create a Vlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VlanArgs | VlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VlanState | undefined;
            resourceInputs["cmpHash"] = state ? state.cmpHash : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
        } else {
            const args = argsOrState as VlanArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["cmpHash"] = args ? args.cmpHash : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vlan resources.
 */
export interface VlanState {
    /**
     * Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [`default`, `src-ip`, `dst-ip`]
     */
    cmpHash?: pulumi.Input<string>;
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.net.VlanInterface>[]>;
    /**
     * Specifies the maximum transmission unit (MTU) for traffic on this VLAN. The default value is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the vlan
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     */
    tag?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Vlan resource.
 */
export interface VlanArgs {
    /**
     * Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [`default`, `src-ip`, `dst-ip`]
     */
    cmpHash?: pulumi.Input<string>;
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.net.VlanInterface>[]>;
    /**
     * Specifies the maximum transmission unit (MTU) for traffic on this VLAN. The default value is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the vlan
     */
    name: pulumi.Input<string>;
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     */
    tag?: pulumi.Input<number>;
}
