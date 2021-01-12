// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileHttp2` Configures a custom profileHttp2 for use by health checks.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const nyhttp2 = new f5bigip.ltm.ProfileHttp2("nyhttp2", {
 *     name: "/Common/test-profile-http2",
 *     frameSize: 2021,
 *     receiveWindow: 31,
 *     writeSize: 16380,
 *     headerTableSize: 4092,
 *     includeContentLength: "enabled",
 *     enforceTlsRequirements: "enabled",
 *     insertHeader: "disabled",
 *     concurrentStreamsPerConnection: 30,
 *     connectionIdleTimeout: 100,
 *     activationModes: ["always"],
 * });
 * //Child Profile which inherits parent http2 profile
 * const nyhttp2_child = new f5bigip.ltm.ProfileHttp2("nyhttp2-child", {
 *     name: "/Common/test-profile-http2-child",
 *     defaultsFrom: nyhttp2.name,
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileHttp2State, opts?: pulumi.CustomResourceOptions): ProfileHttp2 {
        return new ProfileHttp2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileHttp2:ProfileHttp2';

    /**
     * Returns true if the given object is an instance of ProfileHttp2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileHttp2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileHttp2.__pulumiType;
    }

    /**
     * This setting specifies the condition that will cause the BIG-IP system to handle an incoming connection as an HTTP/2 connection, Allowed values : `[“alpn”]` (or) `[“always”]`.
     */
    public readonly activationModes!: pulumi.Output<string[]>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    public readonly concurrentStreamsPerConnection!: pulumi.Output<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion.
     */
    public readonly connectionIdleTimeout!: pulumi.Output<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Enable or disable enforcement of TLS requirements,Allowed Values : `"enabled"/"disabled"` [Default:`"enabled"`].
     */
    public readonly enforceTlsRequirements!: pulumi.Output<string>;
    /**
     * The size of the data frames, in bytes, that the HTTP/2 protocol sends to the client. `Default: 2048`.
     */
    public readonly frameSize!: pulumi.Output<number>;
    /**
     * The size of the header table, in KB, for the HTTP headers that the HTTP/2 protocol compresses to save bandwidth.
     */
    public readonly headerTableSize!: pulumi.Output<number>;
    /**
     * Enable to include content-length in HTTP/2 headers,Default : disabled
     */
    public readonly includeContentLength!: pulumi.Output<string>;
    /**
     * This setting specifies whether the BIG-IP system should add an HTTP header to the HTTP request to show that the request was received over HTTP/2, Allowed Values : `"enabled"/"disabled"` [ Default: `"disabled"`].
     */
    public readonly insertHeader!: pulumi.Output<string>;
    /**
     * This setting specifies the name of the header that the BIG-IP system will add to the HTTP request when the Insert Header is enabled.
     */
    public readonly insertHeaderName!: pulumi.Output<string>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-http2-profile`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The flow-control size for upload streams, in KB. `Default: 32`.
     */
    public readonly receiveWindow!: pulumi.Output<number>;
    /**
     * The total size of combined data frames, in bytes, that the HTTP/2 protocol sends in a single write function. `Default: 16384`".
     */
    public readonly writeSize!: pulumi.Output<number>;

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
            inputs["enforceTlsRequirements"] = state ? state.enforceTlsRequirements : undefined;
            inputs["frameSize"] = state ? state.frameSize : undefined;
            inputs["headerTableSize"] = state ? state.headerTableSize : undefined;
            inputs["includeContentLength"] = state ? state.includeContentLength : undefined;
            inputs["insertHeader"] = state ? state.insertHeader : undefined;
            inputs["insertHeaderName"] = state ? state.insertHeaderName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["receiveWindow"] = state ? state.receiveWindow : undefined;
            inputs["writeSize"] = state ? state.writeSize : undefined;
        } else {
            const args = argsOrState as ProfileHttp2Args | undefined;
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            inputs["activationModes"] = args ? args.activationModes : undefined;
            inputs["concurrentStreamsPerConnection"] = args ? args.concurrentStreamsPerConnection : undefined;
            inputs["connectionIdleTimeout"] = args ? args.connectionIdleTimeout : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["enforceTlsRequirements"] = args ? args.enforceTlsRequirements : undefined;
            inputs["frameSize"] = args ? args.frameSize : undefined;
            inputs["headerTableSize"] = args ? args.headerTableSize : undefined;
            inputs["includeContentLength"] = args ? args.includeContentLength : undefined;
            inputs["insertHeader"] = args ? args.insertHeader : undefined;
            inputs["insertHeaderName"] = args ? args.insertHeaderName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["receiveWindow"] = args ? args.receiveWindow : undefined;
            inputs["writeSize"] = args ? args.writeSize : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileHttp2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileHttp2 resources.
 */
export interface ProfileHttp2State {
    /**
     * This setting specifies the condition that will cause the BIG-IP system to handle an incoming connection as an HTTP/2 connection, Allowed values : `[“alpn”]` (or) `[“always”]`.
     */
    readonly activationModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    readonly concurrentStreamsPerConnection?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion.
     */
    readonly connectionIdleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Enable or disable enforcement of TLS requirements,Allowed Values : `"enabled"/"disabled"` [Default:`"enabled"`].
     */
    readonly enforceTlsRequirements?: pulumi.Input<string>;
    /**
     * The size of the data frames, in bytes, that the HTTP/2 protocol sends to the client. `Default: 2048`.
     */
    readonly frameSize?: pulumi.Input<number>;
    /**
     * The size of the header table, in KB, for the HTTP headers that the HTTP/2 protocol compresses to save bandwidth.
     */
    readonly headerTableSize?: pulumi.Input<number>;
    /**
     * Enable to include content-length in HTTP/2 headers,Default : disabled
     */
    readonly includeContentLength?: pulumi.Input<string>;
    /**
     * This setting specifies whether the BIG-IP system should add an HTTP header to the HTTP request to show that the request was received over HTTP/2, Allowed Values : `"enabled"/"disabled"` [ Default: `"disabled"`].
     */
    readonly insertHeader?: pulumi.Input<string>;
    /**
     * This setting specifies the name of the header that the BIG-IP system will add to the HTTP request when the Insert Header is enabled.
     */
    readonly insertHeaderName?: pulumi.Input<string>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-http2-profile`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The flow-control size for upload streams, in KB. `Default: 32`.
     */
    readonly receiveWindow?: pulumi.Input<number>;
    /**
     * The total size of combined data frames, in bytes, that the HTTP/2 protocol sends in a single write function. `Default: 16384`".
     */
    readonly writeSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProfileHttp2 resource.
 */
export interface ProfileHttp2Args {
    /**
     * This setting specifies the condition that will cause the BIG-IP system to handle an incoming connection as an HTTP/2 connection, Allowed values : `[“alpn”]` (or) `[“always”]`.
     */
    readonly activationModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
     */
    readonly concurrentStreamsPerConnection?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion.
     */
    readonly connectionIdleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Enable or disable enforcement of TLS requirements,Allowed Values : `"enabled"/"disabled"` [Default:`"enabled"`].
     */
    readonly enforceTlsRequirements?: pulumi.Input<string>;
    /**
     * The size of the data frames, in bytes, that the HTTP/2 protocol sends to the client. `Default: 2048`.
     */
    readonly frameSize?: pulumi.Input<number>;
    /**
     * The size of the header table, in KB, for the HTTP headers that the HTTP/2 protocol compresses to save bandwidth.
     */
    readonly headerTableSize?: pulumi.Input<number>;
    /**
     * Enable to include content-length in HTTP/2 headers,Default : disabled
     */
    readonly includeContentLength?: pulumi.Input<string>;
    /**
     * This setting specifies whether the BIG-IP system should add an HTTP header to the HTTP request to show that the request was received over HTTP/2, Allowed Values : `"enabled"/"disabled"` [ Default: `"disabled"`].
     */
    readonly insertHeader?: pulumi.Input<string>;
    /**
     * This setting specifies the name of the header that the BIG-IP system will add to the HTTP request when the Insert Header is enabled.
     */
    readonly insertHeaderName?: pulumi.Input<string>;
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-http2-profile`.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The flow-control size for upload streams, in KB. `Default: 32`.
     */
    readonly receiveWindow?: pulumi.Input<number>;
    /**
     * The total size of combined data frames, in bytes, that the HTTP/2 protocol sends in a single write function. `Default: 16384`".
     */
    readonly writeSize?: pulumi.Input<number>;
}
