// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_sys_provision` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const provision_ilx = new f5bigip.sys.Provision("provision-ilx", {
 *     cpuRatio: 0,
 *     diskRatio: 0,
 *     fullPath: "ilx",
 *     level: "nominal",
 *     memoryRatio: 0,
 *     name: "/Common/ilx",
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProvisionState, opts?: pulumi.CustomResourceOptions): Provision {
        return new Provision(name, <any>state, { ...opts, id: id });
    }

    /**
     * cpu Ratio
     */
    public readonly cpuRatio: pulumi.Output<number | undefined>;
    /**
     * disk Ratio
     */
    public readonly diskRatio: pulumi.Output<number | undefined>;
    /**
     * path
     */
    public readonly fullPath: pulumi.Output<string | undefined>;
    /**
     * what level nominal or dedicated
     */
    public readonly level: pulumi.Output<string | undefined>;
    /**
     * memory Ratio
     */
    public readonly memoryRatio: pulumi.Output<number | undefined>;
    /**
     * Name of the module to be provisioned
     */
    public readonly name: pulumi.Output<string>;

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
        if (opts && opts.id) {
            const state: ProvisionState = argsOrState as ProvisionState | undefined;
            inputs["cpuRatio"] = state ? state.cpuRatio : undefined;
            inputs["diskRatio"] = state ? state.diskRatio : undefined;
            inputs["fullPath"] = state ? state.fullPath : undefined;
            inputs["level"] = state ? state.level : undefined;
            inputs["memoryRatio"] = state ? state.memoryRatio : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProvisionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["cpuRatio"] = args ? args.cpuRatio : undefined;
            inputs["diskRatio"] = args ? args.diskRatio : undefined;
            inputs["fullPath"] = args ? args.fullPath : undefined;
            inputs["level"] = args ? args.level : undefined;
            inputs["memoryRatio"] = args ? args.memoryRatio : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("f5bigip:sys/provision:Provision", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Provision resources.
 */
export interface ProvisionState {
    /**
     * cpu Ratio
     */
    readonly cpuRatio?: pulumi.Input<number>;
    /**
     * disk Ratio
     */
    readonly diskRatio?: pulumi.Input<number>;
    /**
     * path
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * what level nominal or dedicated
     */
    readonly level?: pulumi.Input<string>;
    /**
     * memory Ratio
     */
    readonly memoryRatio?: pulumi.Input<number>;
    /**
     * Name of the module to be provisioned
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Provision resource.
 */
export interface ProvisionArgs {
    /**
     * cpu Ratio
     */
    readonly cpuRatio?: pulumi.Input<number>;
    /**
     * disk Ratio
     */
    readonly diskRatio?: pulumi.Input<number>;
    /**
     * path
     */
    readonly fullPath?: pulumi.Input<string>;
    /**
     * what level nominal or dedicated
     */
    readonly level?: pulumi.Input<string>;
    /**
     * memory Ratio
     */
    readonly memoryRatio?: pulumi.Input<number>;
    /**
     * Name of the module to be provisioned
     */
    readonly name: pulumi.Input<string>;
}
