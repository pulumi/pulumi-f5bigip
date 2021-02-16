// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class BigIpLicense extends pulumi.CustomResource {
    /**
     * Get an existing BigIpLicense resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BigIpLicenseState, opts?: pulumi.CustomResourceOptions): BigIpLicense {
        return new BigIpLicense(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/bigIpLicense:BigIpLicense';

    /**
     * Returns true if the given object is an instance of BigIpLicense.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BigIpLicense {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BigIpLicense.__pulumiType;
    }

    /**
     * Tmsh command to execute tmsh commands like install
     */
    public readonly command!: pulumi.Output<string>;
    /**
     * A unique Key F5 provides for Licensing BIG-IP
     */
    public readonly registrationKey!: pulumi.Output<string>;

    /**
     * Create a BigIpLicense resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BigIpLicenseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BigIpLicenseArgs | BigIpLicenseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BigIpLicenseState | undefined;
            inputs["command"] = state ? state.command : undefined;
            inputs["registrationKey"] = state ? state.registrationKey : undefined;
        } else {
            const args = argsOrState as BigIpLicenseArgs | undefined;
            if ((!args || args.command === undefined) && !opts.urn) {
                throw new Error("Missing required property 'command'");
            }
            if ((!args || args.registrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registrationKey'");
            }
            inputs["command"] = args ? args.command : undefined;
            inputs["registrationKey"] = args ? args.registrationKey : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BigIpLicense.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BigIpLicense resources.
 */
export interface BigIpLicenseState {
    /**
     * Tmsh command to execute tmsh commands like install
     */
    readonly command?: pulumi.Input<string>;
    /**
     * A unique Key F5 provides for Licensing BIG-IP
     */
    readonly registrationKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BigIpLicense resource.
 */
export interface BigIpLicenseArgs {
    /**
     * Tmsh command to execute tmsh commands like install
     */
    readonly command: pulumi.Input<string>;
    /**
     * A unique Key F5 provides for Licensing BIG-IP
     */
    readonly registrationKey: pulumi.Input<string>;
}
