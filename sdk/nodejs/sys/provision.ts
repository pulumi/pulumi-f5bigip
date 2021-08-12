// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.Provision` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_provision = new f5bigip.sys.Provision("test-provision", {
 *     cpuRatio: 0,
 *     diskRatio: 0,
 *     fullPath: "asm",
 *     level: "none",
 *     memoryRatio: 0,
 *     name: "TEST_ASM_PROVISION_NAME",
 * });
 * ```
 */
export class Provision extends pulumi.CustomResource {
    /**
     * Get an existing Provision resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProvisionState, opts?: pulumi.CustomResourceOptions): Provision {
        return new Provision(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/provision:Provision';

    /**
     * Returns true if the given object is an instance of Provision.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provision {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provision.__pulumiType;
    }

    /**
     * cpu Ratio
     */
    public readonly cpuRatio!: pulumi.Output<number | undefined>;
    /**
     * disk Ratio
     */
    public readonly diskRatio!: pulumi.Output<number | undefined>;
    /**
     * path
     */
    public readonly fullPath!: pulumi.Output<string>;
    /**
     * what level nominal or dedicated
     */
    public readonly level!: pulumi.Output<string | undefined>;
    /**
     * memory Ratio
     */
    public readonly memoryRatio!: pulumi.Output<number | undefined>;
    /**
     * Name of the module to be provisioned
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Provision resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProvisionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProvisionArgs | ProvisionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProvisionState | undefined;
            inputs["cpuRatio"] = state ? state.cpuRatio : undefined;
            inputs["diskRatio"] = state ? state.diskRatio : undefined;
            inputs["fullPath"] = state ? state.fullPath : undefined;
            inputs["level"] = state ? state.level : undefined;
            inputs["memoryRatio"] = state ? state.memoryRatio : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProvisionArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["cpuRatio"] = args ? args.cpuRatio : undefined;
            inputs["diskRatio"] = args ? args.diskRatio : undefined;
            inputs["fullPath"] = args ? args.fullPath : undefined;
            inputs["level"] = args ? args.level : undefined;
            inputs["memoryRatio"] = args ? args.memoryRatio : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provision.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Provision resources.
 */
export interface ProvisionState {
    /**
     * cpu Ratio
     */
    cpuRatio?: pulumi.Input<number>;
    /**
     * disk Ratio
     */
    diskRatio?: pulumi.Input<number>;
    /**
     * path
     */
    fullPath?: pulumi.Input<string>;
    /**
     * what level nominal or dedicated
     */
    level?: pulumi.Input<string>;
    /**
     * memory Ratio
     */
    memoryRatio?: pulumi.Input<number>;
    /**
     * Name of the module to be provisioned
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Provision resource.
 */
export interface ProvisionArgs {
    /**
     * cpu Ratio
     */
    cpuRatio?: pulumi.Input<number>;
    /**
     * disk Ratio
     */
    diskRatio?: pulumi.Input<number>;
    /**
     * path
     */
    fullPath?: pulumi.Input<string>;
    /**
     * what level nominal or dedicated
     */
    level?: pulumi.Input<string>;
    /**
     * memory Ratio
     */
    memoryRatio?: pulumi.Input<number>;
    /**
     * Name of the module to be provisioned
     */
    name: pulumi.Input<string>;
}
