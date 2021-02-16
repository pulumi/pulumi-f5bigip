// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.As3` provides details about bigip as3 resource
 *
 * This resource is helpful to configure as3 declarative JSON on BIG-IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * from "fs";
 *
 * // Example Usage for json file
 * const as3_example1As3 = new f5bigip.As3("as3-example1As3", {as3Json: fs.readFileSync("example1.json")});
 * // Example Usage for json file with tenant filter
 * const as3_example1Index_as3As3 = new f5bigip.As3("as3-example1Index/as3As3", {
 *     as3Json: fs.readFileSync("example2.json"),
 *     tenantFilter: "Sample_03",
 * });
 * ```
 */
export class As3 extends pulumi.CustomResource {
    /**
     * Get an existing As3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: As3State, opts?: pulumi.CustomResourceOptions): As3 {
        return new As3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/as3:As3';

    /**
     * Returns true if the given object is an instance of As3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is As3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === As3.__pulumiType;
    }

    /**
     * Name of Application
     */
    public readonly applicationList!: pulumi.Output<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    public readonly as3Json!: pulumi.Output<string>;
    /**
     * If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
     */
    public readonly tenantFilter!: pulumi.Output<string | undefined>;
    /**
     * Name of Tenant
     */
    public readonly tenantList!: pulumi.Output<string>;
    /**
     * Name of Tenant
     *
     * @deprecated this attribute is no longer in use
     */
    public readonly tenantName!: pulumi.Output<string | undefined>;

    /**
     * Create a As3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: As3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: As3Args | As3State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as As3State | undefined;
            inputs["applicationList"] = state ? state.applicationList : undefined;
            inputs["as3Json"] = state ? state.as3Json : undefined;
            inputs["tenantFilter"] = state ? state.tenantFilter : undefined;
            inputs["tenantList"] = state ? state.tenantList : undefined;
            inputs["tenantName"] = state ? state.tenantName : undefined;
        } else {
            const args = argsOrState as As3Args | undefined;
            if ((!args || args.as3Json === undefined) && !opts.urn) {
                throw new Error("Missing required property 'as3Json'");
            }
            inputs["applicationList"] = args ? args.applicationList : undefined;
            inputs["as3Json"] = args ? args.as3Json : undefined;
            inputs["tenantFilter"] = args ? args.tenantFilter : undefined;
            inputs["tenantList"] = args ? args.tenantList : undefined;
            inputs["tenantName"] = args ? args.tenantName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(As3.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering As3 resources.
 */
export interface As3State {
    /**
     * Name of Application
     */
    readonly applicationList?: pulumi.Input<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    readonly as3Json?: pulumi.Input<string>;
    /**
     * If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
     */
    readonly tenantFilter?: pulumi.Input<string>;
    /**
     * Name of Tenant
     */
    readonly tenantList?: pulumi.Input<string>;
    /**
     * Name of Tenant
     *
     * @deprecated this attribute is no longer in use
     */
    readonly tenantName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a As3 resource.
 */
export interface As3Args {
    /**
     * Name of Application
     */
    readonly applicationList?: pulumi.Input<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    readonly as3Json: pulumi.Input<string>;
    /**
     * If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
     */
    readonly tenantFilter?: pulumi.Input<string>;
    /**
     * Name of Tenant
     */
    readonly tenantList?: pulumi.Input<string>;
    /**
     * Name of Tenant
     *
     * @deprecated this attribute is no longer in use
     */
    readonly tenantName?: pulumi.Input<string>;
}
