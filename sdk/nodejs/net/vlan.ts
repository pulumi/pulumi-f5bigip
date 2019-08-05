// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_net_vlan` Manages a vlan configuration
 * 
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 * 
 * 
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const vlan1 = new f5bigip.net.Vlan("vlan1", {
 *     interfaces: [{
 *         tagged: false,
 *         vlanport: "1.2",
 *     }],
 *     name: "/Common/Internal",
 *     tag: 101,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/net_vlan.html.markdown.
 */
export class Vlan extends pulumi.CustomResource {
    /**
     * Get an existing Vlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    public readonly interfaces!: pulumi.Output<{ tagged?: boolean, vlanport?: string }[] | undefined>;
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VlanState | undefined;
            inputs["interfaces"] = state ? state.interfaces : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tag"] = state ? state.tag : undefined;
        } else {
            const args = argsOrState as VlanArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["interfaces"] = args ? args.interfaces : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tag"] = args ? args.tag : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Vlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vlan resources.
 */
export interface VlanState {
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    readonly interfaces?: pulumi.Input<pulumi.Input<{ tagged?: pulumi.Input<boolean>, vlanport?: pulumi.Input<string> }>[]>;
    /**
     * Name of the vlan
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     */
    readonly tag?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Vlan resource.
 */
export interface VlanArgs {
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     */
    readonly interfaces?: pulumi.Input<pulumi.Input<{ tagged?: pulumi.Input<boolean>, vlanport?: pulumi.Input<string> }>[]>;
    /**
     * Name of the vlan
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     */
    readonly tag?: pulumi.Input<number>;
}
