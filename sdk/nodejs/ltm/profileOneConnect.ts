// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileOneConnect` Configures a custom profileOneconnect for use by health checks.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_oneconnect = new f5bigip.ltm.ProfileOneConnect("test-oneconnect", {
 *     name: "/Common/test-oneconnect",
 * });
 * ```
 *
 * ## Import
 *
 * BIG-IP LTM oneconnect profiles can be imported using the `name` , e.g.
 *
 * ```sh
 *  $ pulumi import f5bigip:ltm/profileOneConnect:ProfileOneConnect test-oneconnect /Common/test-oneconnect
 * ```
 */
export class ProfileOneConnect extends pulumi.CustomResource {
    /**
     * Get an existing ProfileOneConnect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileOneConnectState, opts?: pulumi.CustomResourceOptions): ProfileOneConnect {
        return new ProfileOneConnect(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileOneConnect:ProfileOneConnect';

    /**
     * Returns true if the given object is an instance of ProfileOneConnect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileOneConnect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileOneConnect.__pulumiType;
    }

    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     */
    public readonly idleTimeoutOverride!: pulumi.Output<string>;
    /**
     * Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     */
    public readonly limitType!: pulumi.Output<string>;
    /**
     * Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     */
    public readonly maxAge!: pulumi.Output<number>;
    /**
     * Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     */
    public readonly maxReuse!: pulumi.Output<number>;
    /**
     * Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     */
    public readonly maxSize!: pulumi.Output<number>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * Specify if you want to share the pool, default value is `disabled`.
     */
    public readonly sharePools!: pulumi.Output<string>;
    /**
     * Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
     */
    public readonly sourceMask!: pulumi.Output<string>;

    /**
     * Create a ProfileOneConnect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileOneConnectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileOneConnectArgs | ProfileOneConnectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileOneConnectState | undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["idleTimeoutOverride"] = state ? state.idleTimeoutOverride : undefined;
            inputs["limitType"] = state ? state.limitType : undefined;
            inputs["maxAge"] = state ? state.maxAge : undefined;
            inputs["maxReuse"] = state ? state.maxReuse : undefined;
            inputs["maxSize"] = state ? state.maxSize : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
            inputs["sharePools"] = state ? state.sharePools : undefined;
            inputs["sourceMask"] = state ? state.sourceMask : undefined;
        } else {
            const args = argsOrState as ProfileOneConnectArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["idleTimeoutOverride"] = args ? args.idleTimeoutOverride : undefined;
            inputs["limitType"] = args ? args.limitType : undefined;
            inputs["maxAge"] = args ? args.maxAge : undefined;
            inputs["maxReuse"] = args ? args.maxReuse : undefined;
            inputs["maxSize"] = args ? args.maxSize : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
            inputs["sharePools"] = args ? args.sharePools : undefined;
            inputs["sourceMask"] = args ? args.sourceMask : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProfileOneConnect.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileOneConnect resources.
 */
export interface ProfileOneConnectState {
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     */
    readonly idleTimeoutOverride?: pulumi.Input<string>;
    /**
     * Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     */
    readonly limitType?: pulumi.Input<string>;
    /**
     * Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     */
    readonly maxAge?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     */
    readonly maxReuse?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     */
    readonly maxSize?: pulumi.Input<number>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specify if you want to share the pool, default value is `disabled`.
     */
    readonly sharePools?: pulumi.Input<string>;
    /**
     * Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
     */
    readonly sourceMask?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileOneConnect resource.
 */
export interface ProfileOneConnectArgs {
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     */
    readonly idleTimeoutOverride?: pulumi.Input<string>;
    /**
     * Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     */
    readonly limitType?: pulumi.Input<string>;
    /**
     * Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     */
    readonly maxAge?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     */
    readonly maxReuse?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     */
    readonly maxSize?: pulumi.Input<number>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profileName`,For example `/Common/test-oneconnect-profile`.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specify if you want to share the pool, default value is `disabled`.
     */
    readonly sharePools?: pulumi.Input<string>;
    /**
     * Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
     */
    readonly sourceMask?: pulumi.Input<string>;
}
