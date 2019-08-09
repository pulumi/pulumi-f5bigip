// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.VirtualAddress` Configures Virtual Server
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
 * const vsVa = new f5bigip.ltm.VirtualAddress("vsVa", {
 *     advertizeRoute: true,
 *     name: "/Common/vs_va",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_virtual_address.html.markdown.
 */
export class VirtualAddress extends pulumi.CustomResource {
    /**
     * Get an existing VirtualAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualAddressState, opts?: pulumi.CustomResourceOptions): VirtualAddress {
        return new VirtualAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/virtualAddress:VirtualAddress';

    /**
     * Returns true if the given object is an instance of VirtualAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualAddress.__pulumiType;
    }

    /**
     * Enabled dynamic routing of the address
     */
    public readonly advertizeRoute!: pulumi.Output<boolean | undefined>;
    /**
     * Enable or disable ARP for the virtual address
     */
    public readonly arp!: pulumi.Output<boolean | undefined>;
    /**
     * Automatically delete the virtual address with the virtual server
     */
    public readonly autoDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Max number of connections for virtual address
     */
    public readonly connLimit!: pulumi.Output<number | undefined>;
    /**
     * Enable or disable the virtual address
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable/Disable ICMP response to the virtual address
     */
    public readonly icmpEcho!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the virtual address
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specify the partition and traffic group
     */
    public readonly trafficGroup!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualAddressArgs | VirtualAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualAddressState | undefined;
            inputs["advertizeRoute"] = state ? state.advertizeRoute : undefined;
            inputs["arp"] = state ? state.arp : undefined;
            inputs["autoDelete"] = state ? state.autoDelete : undefined;
            inputs["connLimit"] = state ? state.connLimit : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["icmpEcho"] = state ? state.icmpEcho : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["trafficGroup"] = state ? state.trafficGroup : undefined;
        } else {
            const args = argsOrState as VirtualAddressArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["advertizeRoute"] = args ? args.advertizeRoute : undefined;
            inputs["arp"] = args ? args.arp : undefined;
            inputs["autoDelete"] = args ? args.autoDelete : undefined;
            inputs["connLimit"] = args ? args.connLimit : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["icmpEcho"] = args ? args.icmpEcho : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["trafficGroup"] = args ? args.trafficGroup : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualAddress resources.
 */
export interface VirtualAddressState {
    /**
     * Enabled dynamic routing of the address
     */
    readonly advertizeRoute?: pulumi.Input<boolean>;
    /**
     * Enable or disable ARP for the virtual address
     */
    readonly arp?: pulumi.Input<boolean>;
    /**
     * Automatically delete the virtual address with the virtual server
     */
    readonly autoDelete?: pulumi.Input<boolean>;
    /**
     * Max number of connections for virtual address
     */
    readonly connLimit?: pulumi.Input<number>;
    /**
     * Enable or disable the virtual address
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Enable/Disable ICMP response to the virtual address
     */
    readonly icmpEcho?: pulumi.Input<boolean>;
    /**
     * Name of the virtual address
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specify the partition and traffic group
     */
    readonly trafficGroup?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualAddress resource.
 */
export interface VirtualAddressArgs {
    /**
     * Enabled dynamic routing of the address
     */
    readonly advertizeRoute?: pulumi.Input<boolean>;
    /**
     * Enable or disable ARP for the virtual address
     */
    readonly arp?: pulumi.Input<boolean>;
    /**
     * Automatically delete the virtual address with the virtual server
     */
    readonly autoDelete?: pulumi.Input<boolean>;
    /**
     * Max number of connections for virtual address
     */
    readonly connLimit?: pulumi.Input<number>;
    /**
     * Enable or disable the virtual address
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Enable/Disable ICMP response to the virtual address
     */
    readonly icmpEcho?: pulumi.Input<boolean>;
    /**
     * Name of the virtual address
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specify the partition and traffic group
     */
    readonly trafficGroup?: pulumi.Input<string>;
}
