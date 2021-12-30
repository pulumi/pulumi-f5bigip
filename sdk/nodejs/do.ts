// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.Do` provides details about bigip do resource
 *
 * This resource is helpful to configure do declarative JSON on BIG-IP.
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as fs from "fs";
 *
 * const do_example = new f5bigip.Do("do-example", {
 *     doJson: fs.readFileSync("example.json", "utf-8"),
 *     timeout: 15,
 * });
 * ```
 */
export class Do extends pulumi.CustomResource {
    /**
     * Get an existing Do resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DoState, opts?: pulumi.CustomResourceOptions): Do {
        return new Do(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/do:Do';

    /**
     * Returns true if the given object is an instance of Do.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Do {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Do.__pulumiType;
    }

    /**
     * IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    public readonly bigipAddress!: pulumi.Output<string | undefined>;
    /**
     * Password of  BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    public readonly bigipPassword!: pulumi.Output<string | undefined>;
    /**
     * Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    public readonly bigipPort!: pulumi.Output<string | undefined>;
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     */
    public readonly bigipTokenAuth!: pulumi.Output<boolean | undefined>;
    /**
     * UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    public readonly bigipUser!: pulumi.Output<string | undefined>;
    /**
     * Name of the of the Declarative DO JSON file
     */
    public readonly doJson!: pulumi.Output<string>;
    /**
     * unique identifier for DO resource
     *
     * @deprecated this attribute is no longer in use
     */
    public readonly tenantName!: pulumi.Output<string | undefined>;
    /**
     * DO json
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a Do resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DoArgs | DoState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DoState | undefined;
            inputs["bigipAddress"] = state ? state.bigipAddress : undefined;
            inputs["bigipPassword"] = state ? state.bigipPassword : undefined;
            inputs["bigipPort"] = state ? state.bigipPort : undefined;
            inputs["bigipTokenAuth"] = state ? state.bigipTokenAuth : undefined;
            inputs["bigipUser"] = state ? state.bigipUser : undefined;
            inputs["doJson"] = state ? state.doJson : undefined;
            inputs["tenantName"] = state ? state.tenantName : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as DoArgs | undefined;
            if ((!args || args.doJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'doJson'");
            }
            inputs["bigipAddress"] = args ? args.bigipAddress : undefined;
            inputs["bigipPassword"] = args ? args.bigipPassword : undefined;
            inputs["bigipPort"] = args ? args.bigipPort : undefined;
            inputs["bigipTokenAuth"] = args ? args.bigipTokenAuth : undefined;
            inputs["bigipUser"] = args ? args.bigipUser : undefined;
            inputs["doJson"] = args ? args.doJson : undefined;
            inputs["tenantName"] = args ? args.tenantName : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Do.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Do resources.
 */
export interface DoState {
    /**
     * IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipAddress?: pulumi.Input<string>;
    /**
     * Password of  BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipPassword?: pulumi.Input<string>;
    /**
     * Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipPort?: pulumi.Input<string>;
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     */
    bigipTokenAuth?: pulumi.Input<boolean>;
    /**
     * UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipUser?: pulumi.Input<string>;
    /**
     * Name of the of the Declarative DO JSON file
     */
    doJson?: pulumi.Input<string>;
    /**
     * unique identifier for DO resource
     *
     * @deprecated this attribute is no longer in use
     */
    tenantName?: pulumi.Input<string>;
    /**
     * DO json
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Do resource.
 */
export interface DoArgs {
    /**
     * IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipAddress?: pulumi.Input<string>;
    /**
     * Password of  BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipPassword?: pulumi.Input<string>;
    /**
     * Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipPort?: pulumi.Input<string>;
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     */
    bigipTokenAuth?: pulumi.Input<boolean>;
    /**
     * UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     */
    bigipUser?: pulumi.Input<string>;
    /**
     * Name of the of the Declarative DO JSON file
     */
    doJson: pulumi.Input<string>;
    /**
     * unique identifier for DO resource
     *
     * @deprecated this attribute is no longer in use
     */
    tenantName?: pulumi.Input<string>;
    /**
     * DO json
     */
    timeout?: pulumi.Input<number>;
}
