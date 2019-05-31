// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_ltm_profile_http2` Configures a custom profile_http2 for use by health checks.
 * 
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 * 
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const nyhttp2 = new f5bigip.ltm.ProfileHttp2("nyhttp2", {
 *     activationModes: [
 *         "alpn",
 *         "npn",
 *     ],
 *     concurrentStreamsPerConnection: 10,
 *     connectionIdleTimeout: 30,
 *     defaultsFrom: "/Common/http2",
 *     name: "/Common/NewYork_http2",
 * });
 * ```
 */
export class ProfileHttp2 extends pulumi.CustomResource {
    /**
     * Get an existing ProfileHttp2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileHttp2State, opts?: pulumi.CustomResourceOptions): ProfileHttp2 {
        return new ProfileHttp2(name, <any>state, { ...opts, id: id });
    }

    /**
     * Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
     */
    public readonly activationModes!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    public readonly concurrentStreamsPerConnection!: pulumi.Output<number | undefined>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
     */
    public readonly connectionIdleTimeout!: pulumi.Output<number | undefined>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Use the parent Http2 profile
     */
    public readonly headerTableSize!: pulumi.Output<number | undefined>;
    /**
     * Name of the profile_http2
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ProfileHttp2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileHttp2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileHttp2Args | ProfileHttp2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileHttp2State | undefined;
            inputs["activationModes"] = state ? state.activationModes : undefined;
            inputs["concurrentStreamsPerConnection"] = state ? state.concurrentStreamsPerConnection : undefined;
            inputs["connectionIdleTimeout"] = state ? state.connectionIdleTimeout : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["headerTableSize"] = state ? state.headerTableSize : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProfileHttp2Args | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["activationModes"] = args ? args.activationModes : undefined;
            inputs["concurrentStreamsPerConnection"] = args ? args.concurrentStreamsPerConnection : undefined;
            inputs["connectionIdleTimeout"] = args ? args.connectionIdleTimeout : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["headerTableSize"] = args ? args.headerTableSize : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("f5bigip:ltm/profileHttp2:ProfileHttp2", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileHttp2 resources.
 */
export interface ProfileHttp2State {
    /**
     * Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
     */
    readonly activationModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    readonly concurrentStreamsPerConnection?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
     */
    readonly connectionIdleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Use the parent Http2 profile
     */
    readonly headerTableSize?: pulumi.Input<number>;
    /**
     * Name of the profile_http2
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileHttp2 resource.
 */
export interface ProfileHttp2Args {
    /**
     * Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
     */
    readonly activationModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    readonly concurrentStreamsPerConnection?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
     */
    readonly connectionIdleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Use the parent Http2 profile
     */
    readonly headerTableSize?: pulumi.Input<number>;
    /**
     * Name of the profile_http2
     */
    readonly name: pulumi.Input<string>;
}
