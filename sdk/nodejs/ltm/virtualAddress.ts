// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.VirtualAddress` Configures Virtual Server
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/virtual_server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const vsVa = new f5bigip.ltm.VirtualAddress("vs_va", {
 *     name: "/Common/xxxxx",
 *     advertizeRoute: "enabled",
 * });
 * ```
 */
export class VirtualAddress extends pulumi.CustomResource {
    /**
     * Get an existing VirtualAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
     */
    public readonly advertizeRoute!: pulumi.Output<string | undefined>;
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
     * Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     */
    public readonly icmpEcho!: pulumi.Output<string | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualAddressState | undefined;
            resourceInputs["advertizeRoute"] = state ? state.advertizeRoute : undefined;
            resourceInputs["arp"] = state ? state.arp : undefined;
            resourceInputs["autoDelete"] = state ? state.autoDelete : undefined;
            resourceInputs["connLimit"] = state ? state.connLimit : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["icmpEcho"] = state ? state.icmpEcho : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["trafficGroup"] = state ? state.trafficGroup : undefined;
        } else {
            const args = argsOrState as VirtualAddressArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["advertizeRoute"] = args ? args.advertizeRoute : undefined;
            resourceInputs["arp"] = args ? args.arp : undefined;
            resourceInputs["autoDelete"] = args ? args.autoDelete : undefined;
            resourceInputs["connLimit"] = args ? args.connLimit : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["icmpEcho"] = args ? args.icmpEcho : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["trafficGroup"] = args ? args.trafficGroup : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualAddress resources.
 */
export interface VirtualAddressState {
    /**
     * Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
     */
    advertizeRoute?: pulumi.Input<string>;
    /**
     * Enable or disable ARP for the virtual address
     */
    arp?: pulumi.Input<boolean>;
    /**
     * Automatically delete the virtual address with the virtual server
     */
    autoDelete?: pulumi.Input<boolean>;
    /**
     * Max number of connections for virtual address
     */
    connLimit?: pulumi.Input<number>;
    /**
     * Enable or disable the virtual address
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     */
    icmpEcho?: pulumi.Input<string>;
    /**
     * Name of the virtual address
     */
    name?: pulumi.Input<string>;
    /**
     * Specify the partition and traffic group
     */
    trafficGroup?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualAddress resource.
 */
export interface VirtualAddressArgs {
    /**
     * Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
     */
    advertizeRoute?: pulumi.Input<string>;
    /**
     * Enable or disable ARP for the virtual address
     */
    arp?: pulumi.Input<boolean>;
    /**
     * Automatically delete the virtual address with the virtual server
     */
    autoDelete?: pulumi.Input<boolean>;
    /**
     * Max number of connections for virtual address
     */
    connLimit?: pulumi.Input<number>;
    /**
     * Enable or disable the virtual address
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     */
    icmpEcho?: pulumi.Input<string>;
    /**
     * Name of the virtual address
     */
    name: pulumi.Input<string>;
    /**
     * Specify the partition and traffic group
     */
    trafficGroup?: pulumi.Input<string>;
}
