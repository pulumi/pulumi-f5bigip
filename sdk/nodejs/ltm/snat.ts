// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Snat` Manages a snat configuration
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
 * const testSnat = new f5bigip.ltm.Snat("test-snat", {
 *     autolasthop: "default",
 *     fullPath: "/Common/test-snat",
 *     mirror: "disabled",
 *     name: "TEST_SNAT_NAME",
 *     origins: [
 *         {
 *             name: "2.2.2.2",
 *         },
 *         {
 *             name: "3.3.3.3",
 *         },
 *     ],
 *     partition: "Common",
 *     translation: "/Common/136.1.1.1",
 *     vlansdisabled: true,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_snat.html.markdown.
 */
export class Snat extends pulumi.CustomResource {
    /**
     * Get an existing Snat resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatState, opts?: pulumi.CustomResourceOptions): Snat {
        return new Snat(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/snat:Snat';

    /**
     * Returns true if the given object is an instance of Snat.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snat {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snat.__pulumiType;
    }

    /**
     * Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
     */
    public readonly autolasthop!: pulumi.Output<string | undefined>;
    /**
     * Fullpath
     */
    public readonly fullPath!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables mirroring of SNAT connections.
     */
    public readonly mirror!: pulumi.Output<string | undefined>;
    /**
     * Name of the snat
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IP or hostname of the snat
     */
    public readonly origins!: pulumi.Output<outputs.ltm.SnatOrigin[]>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
     */
    public readonly snatpool!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
     */
    public readonly sourceport!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
     */
    public readonly translation!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
     */
    public readonly vlans!: pulumi.Output<string[] | undefined>;
    /**
     * Disables the SNAT on all VLANs.
     */
    public readonly vlansdisabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Snat resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnatArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnatArgs | SnatState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SnatState | undefined;
            inputs["autolasthop"] = state ? state.autolasthop : undefined;
            inputs["fullPath"] = state ? state.fullPath : undefined;
            inputs["mirror"] = state ? state.mirror : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["origins"] = state ? state.origins : undefined;
            inputs["partition"] = state ? state.partition : undefined;
            inputs["snatpool"] = state ? state.snatpool : undefined;
            inputs["sourceport"] = state ? state.sourceport : undefined;
            inputs["translation"] = state ? state.translation : undefined;
            inputs["vlans"] = state ? state.vlans : undefined;
            inputs["vlansdisabled"] = state ? state.vlansdisabled : undefined;
        } else {
            const args = argsOrState as SnatArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.origins === undefined) {
                throw new Error("Missing required property 'origins'");
            }
            inputs["autolasthop"] = args ? args.autolasthop : undefined;
            inputs["fullPath"] = args ? args.fullPath : undefined;
            inputs["mirror"] = args ? args.mirror : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["origins"] = args ? args.origins : undefined;
            inputs["partition"] = args ? args.partition : undefined;
            inputs["snatpool"] = args ? args.snatpool : undefined;
            inputs["sourceport"] = args ? args.sourceport : undefined;
            inputs["translation"] = args ? args.translation : undefined;
            inputs["vlans"] = args ? args.vlans : undefined;
            inputs["vlansdisabled"] = args ? args.vlansdisabled : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Snat.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snat resources.
 */
export interface SnatState {
    /**
     * Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
     */
    readonly autolasthop?: pulumi.Input<string>;
    /**
     * Fullpath
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * Enables or disables mirroring of SNAT connections.
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the snat
     */
    readonly name?: pulumi.Input<string>;
    /**
     * IP or hostname of the snat
     */
    readonly origins?: pulumi.Input<pulumi.Input<inputs.ltm.SnatOrigin>[]>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
     */
    readonly snatpool?: pulumi.Input<string>;
    /**
     * Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
     */
    readonly sourceport?: pulumi.Input<string>;
    /**
     * Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
     */
    readonly translation?: pulumi.Input<string>;
    /**
     * Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
     */
    readonly vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Disables the SNAT on all VLANs.
     */
    readonly vlansdisabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Snat resource.
 */
export interface SnatArgs {
    /**
     * Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
     */
    readonly autolasthop?: pulumi.Input<string>;
    /**
     * Fullpath
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * Enables or disables mirroring of SNAT connections.
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the snat
     */
    readonly name: pulumi.Input<string>;
    /**
     * IP or hostname of the snat
     */
    readonly origins: pulumi.Input<pulumi.Input<inputs.ltm.SnatOrigin>[]>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
     */
    readonly snatpool?: pulumi.Input<string>;
    /**
     * Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
     */
    readonly sourceport?: pulumi.Input<string>;
    /**
     * Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
     */
    readonly translation?: pulumi.Input<string>;
    /**
     * Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
     */
    readonly vlans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Disables the SNAT on all VLANs.
     */
    readonly vlansdisabled?: pulumi.Input<boolean>;
}
