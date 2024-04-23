// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.FastApplication` This resource will create and manage FAST applications on BIG-IP from provided JSON declaration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as std from "@pulumi/std";
 *
 * const foo_app = new f5bigip.FastApplication("foo-app", {
 *     template: "examples/simple_http",
 *     fastJson: std.file({
 *         input: "new_fast_app.json",
 *     }).then(invoke => invoke.result),
 * });
 * ```
 */
export class FastApplication extends pulumi.CustomResource {
    /**
     * Get an existing FastApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FastApplicationState, opts?: pulumi.CustomResourceOptions): FastApplication {
        return new FastApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/fastApplication:FastApplication';

    /**
     * Returns true if the given object is an instance of FastApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FastApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FastApplication.__pulumiType;
    }

    /**
     * A FAST application name.
     *
     *
     *
     * * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
     */
    public /*out*/ readonly application!: pulumi.Output<string>;
    /**
     * Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
     */
    public readonly fastJson!: pulumi.Output<string>;
    /**
     * Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * A FAST tenant name on which you want to manage application.
     */
    public /*out*/ readonly tenant!: pulumi.Output<string>;

    /**
     * Create a FastApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FastApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FastApplicationArgs | FastApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FastApplicationState | undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["fastJson"] = state ? state.fastJson : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["tenant"] = state ? state.tenant : undefined;
        } else {
            const args = argsOrState as FastApplicationArgs | undefined;
            if ((!args || args.fastJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fastJson'");
            }
            resourceInputs["fastJson"] = args ? args.fastJson : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["application"] = undefined /*out*/;
            resourceInputs["tenant"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FastApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FastApplication resources.
 */
export interface FastApplicationState {
    /**
     * A FAST application name.
     *
     *
     *
     * * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
     */
    application?: pulumi.Input<string>;
    /**
     * Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
     */
    fastJson?: pulumi.Input<string>;
    /**
     * Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
     */
    template?: pulumi.Input<string>;
    /**
     * A FAST tenant name on which you want to manage application.
     */
    tenant?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FastApplication resource.
 */
export interface FastApplicationArgs {
    /**
     * Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
     */
    fastJson: pulumi.Input<string>;
    /**
     * Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
     */
    template?: pulumi.Input<string>;
}
