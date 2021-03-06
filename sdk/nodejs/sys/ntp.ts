// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.Ntp` provides details about a specific bigip
 *
 * This resource is helpful when configuring NTP server on the BIG-IP.
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const ntp1 = new f5bigip.sys.Ntp("ntp1", {
 *     description: "/Common/NTP1",
 *     servers: ["time.facebook.com"],
 *     timezone: "America/Los_Angeles",
 * });
 * ```
 */
export class Ntp extends pulumi.CustomResource {
    /**
     * Get an existing Ntp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NtpState, opts?: pulumi.CustomResourceOptions): Ntp {
        return new Ntp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/ntp:Ntp';

    /**
     * Returns true if the given object is an instance of Ntp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ntp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ntp.__pulumiType;
    }

    /**
     * Name of the ntp Servers
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Adds NTP servers to or deletes NTP servers from the BIG-IP system.
     */
    public readonly servers!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the time zone that you want to use for the system time.
     */
    public readonly timezone!: pulumi.Output<string | undefined>;

    /**
     * Create a Ntp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NtpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NtpArgs | NtpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NtpState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["servers"] = state ? state.servers : undefined;
            inputs["timezone"] = state ? state.timezone : undefined;
        } else {
            const args = argsOrState as NtpArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["servers"] = args ? args.servers : undefined;
            inputs["timezone"] = args ? args.timezone : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ntp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ntp resources.
 */
export interface NtpState {
    /**
     * Name of the ntp Servers
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Adds NTP servers to or deletes NTP servers from the BIG-IP system.
     */
    readonly servers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the time zone that you want to use for the system time.
     */
    readonly timezone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ntp resource.
 */
export interface NtpArgs {
    /**
     * Name of the ntp Servers
     */
    readonly description: pulumi.Input<string>;
    /**
     * Adds NTP servers to or deletes NTP servers from the BIG-IP system.
     */
    readonly servers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the time zone that you want to use for the system time.
     */
    readonly timezone?: pulumi.Input<string>;
}
