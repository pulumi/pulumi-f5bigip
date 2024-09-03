// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.Provision` Manage BIG-IP module provisioning. This resource will only provision at the standard levels of Dedicated, Nominal, and Minimum.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const gtm = new f5bigip.sys.Provision("gtm", {
 *     name: "gtm",
 *     cpuRatio: 0,
 *     diskRatio: 0,
 *     level: "nominal",
 *     memoryRatio: 0,
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
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    public readonly cpuRatio!: pulumi.Output<number | undefined>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    public readonly diskRatio!: pulumi.Output<number | undefined>;
    public readonly fullPath!: pulumi.Output<string>;
    /**
     * Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
     * default is `nominal`
     * possible options:
     * * nominal
     * * minimum
     * * none
     * * dedicated
     */
    public readonly level!: pulumi.Output<string | undefined>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    public readonly memoryRatio!: pulumi.Output<number | undefined>;
    /**
     * Name of module to provision in BIG-IP.
     * possible options:
     * * afm
     * * am
     * * apm
     * * cgnat
     * * asm
     * * avr
     * * dos
     * * fps
     * * gtm
     * * ilx
     * * lc
     * * ltm
     * * pem
     * * sslo
     * * swg
     * * urldb
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProvisionState | undefined;
            resourceInputs["cpuRatio"] = state ? state.cpuRatio : undefined;
            resourceInputs["diskRatio"] = state ? state.diskRatio : undefined;
            resourceInputs["fullPath"] = state ? state.fullPath : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["memoryRatio"] = state ? state.memoryRatio : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProvisionArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["cpuRatio"] = args ? args.cpuRatio : undefined;
            resourceInputs["diskRatio"] = args ? args.diskRatio : undefined;
            resourceInputs["fullPath"] = args ? args.fullPath : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["memoryRatio"] = args ? args.memoryRatio : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provision.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Provision resources.
 */
export interface ProvisionState {
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    cpuRatio?: pulumi.Input<number>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    diskRatio?: pulumi.Input<number>;
    fullPath?: pulumi.Input<string>;
    /**
     * Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
     * default is `nominal`
     * possible options:
     * * nominal
     * * minimum
     * * none
     * * dedicated
     */
    level?: pulumi.Input<string>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    memoryRatio?: pulumi.Input<number>;
    /**
     * Name of module to provision in BIG-IP.
     * possible options:
     * * afm
     * * am
     * * apm
     * * cgnat
     * * asm
     * * avr
     * * dos
     * * fps
     * * gtm
     * * ilx
     * * lc
     * * ltm
     * * pem
     * * sslo
     * * swg
     * * urldb
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Provision resource.
 */
export interface ProvisionArgs {
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    cpuRatio?: pulumi.Input<number>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    diskRatio?: pulumi.Input<number>;
    fullPath?: pulumi.Input<string>;
    /**
     * Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
     * default is `nominal`
     * possible options:
     * * nominal
     * * minimum
     * * none
     * * dedicated
     */
    level?: pulumi.Input<string>;
    /**
     * Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
     */
    memoryRatio?: pulumi.Input<number>;
    /**
     * Name of module to provision in BIG-IP.
     * possible options:
     * * afm
     * * am
     * * apm
     * * cgnat
     * * asm
     * * avr
     * * dos
     * * fps
     * * gtm
     * * ilx
     * * lc
     * * ltm
     * * pem
     * * sslo
     * * swg
     * * urldb
     */
    name: pulumi.Input<string>;
}
