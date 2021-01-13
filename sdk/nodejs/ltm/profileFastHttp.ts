// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileFastHttp` Configures a custom profileFasthttp for use by health checks.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sjfasthttpprofile = new f5bigip.ltm.ProfileFastHttp("sjfasthttpprofile", {
 *     connpoolMaxreuse: 2,
 *     connpoolMaxsize: 2048,
 *     connpoolMinsize: 0,
 *     connpoolReplenish: "enabled",
 *     connpoolStep: 4,
 *     connpoolidleTimeoutoverride: 0,
 *     defaultsFrom: "/Common/fasthttp",
 *     forcehttp10response: "disabled",
 *     idleTimeout: 300,
 *     maxheaderSize: 32768,
 *     name: "/Common/sjfasthttpprofile",
 * });
 * ```
 */
export class ProfileFastHttp extends pulumi.CustomResource {
    /**
     * Get an existing ProfileFastHttp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileFastHttpState, opts?: pulumi.CustomResourceOptions): ProfileFastHttp {
        return new ProfileFastHttp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileFastHttp:ProfileFastHttp';

    /**
     * Returns true if the given object is an instance of ProfileFastHttp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileFastHttp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileFastHttp.__pulumiType;
    }

    /**
     * Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
     */
    public readonly connpoolMaxreuse!: pulumi.Output<number | undefined>;
    /**
     * Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
     */
    public readonly connpoolMaxsize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
     */
    public readonly connpoolMinsize!: pulumi.Output<number | undefined>;
    /**
     * The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
     */
    public readonly connpoolReplenish!: pulumi.Output<string | undefined>;
    /**
     * Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
     */
    public readonly connpoolStep!: pulumi.Output<number | undefined>;
    /**
     * Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
     */
    public readonly connpoolidleTimeoutoverride!: pulumi.Output<number | undefined>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
     */
    public readonly forcehttp10response!: pulumi.Output<string | undefined>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    /**
     * Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
     */
    public readonly maxheaderSize!: pulumi.Output<number | undefined>;
    /**
     * Name of the profile_fasthttp
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ProfileFastHttp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileFastHttpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileFastHttpArgs | ProfileFastHttpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileFastHttpState | undefined;
            inputs["connpoolMaxreuse"] = state ? state.connpoolMaxreuse : undefined;
            inputs["connpoolMaxsize"] = state ? state.connpoolMaxsize : undefined;
            inputs["connpoolMinsize"] = state ? state.connpoolMinsize : undefined;
            inputs["connpoolReplenish"] = state ? state.connpoolReplenish : undefined;
            inputs["connpoolStep"] = state ? state.connpoolStep : undefined;
            inputs["connpoolidleTimeoutoverride"] = state ? state.connpoolidleTimeoutoverride : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["forcehttp10response"] = state ? state.forcehttp10response : undefined;
            inputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            inputs["maxheaderSize"] = state ? state.maxheaderSize : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProfileFastHttpArgs | undefined;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["connpoolMaxreuse"] = args ? args.connpoolMaxreuse : undefined;
            inputs["connpoolMaxsize"] = args ? args.connpoolMaxsize : undefined;
            inputs["connpoolMinsize"] = args ? args.connpoolMinsize : undefined;
            inputs["connpoolReplenish"] = args ? args.connpoolReplenish : undefined;
            inputs["connpoolStep"] = args ? args.connpoolStep : undefined;
            inputs["connpoolidleTimeoutoverride"] = args ? args.connpoolidleTimeoutoverride : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["forcehttp10response"] = args ? args.forcehttp10response : undefined;
            inputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            inputs["maxheaderSize"] = args ? args.maxheaderSize : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileFastHttp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileFastHttp resources.
 */
export interface ProfileFastHttpState {
    /**
     * Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
     */
    readonly connpoolMaxreuse?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
     */
    readonly connpoolMaxsize?: pulumi.Input<number>;
    /**
     * Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
     */
    readonly connpoolMinsize?: pulumi.Input<number>;
    /**
     * The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
     */
    readonly connpoolReplenish?: pulumi.Input<string>;
    /**
     * Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
     */
    readonly connpoolStep?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
     */
    readonly connpoolidleTimeoutoverride?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
     */
    readonly forcehttp10response?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
     */
    readonly maxheaderSize?: pulumi.Input<number>;
    /**
     * Name of the profile_fasthttp
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileFastHttp resource.
 */
export interface ProfileFastHttpArgs {
    /**
     * Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
     */
    readonly connpoolMaxreuse?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
     */
    readonly connpoolMaxsize?: pulumi.Input<number>;
    /**
     * Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
     */
    readonly connpoolMinsize?: pulumi.Input<number>;
    /**
     * The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
     */
    readonly connpoolReplenish?: pulumi.Input<string>;
    /**
     * Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
     */
    readonly connpoolStep?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
     */
    readonly connpoolidleTimeoutoverride?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
     */
    readonly forcehttp10response?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
     */
    readonly maxheaderSize?: pulumi.Input<number>;
    /**
     * Name of the profile_fasthttp
     */
    readonly name: pulumi.Input<string>;
}
