// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileFastL4` Configures a custom profileFastl4 for use by health checks.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const profileFastl4 = new f5bigip.ltm.ProfileFastL4("profile_fastl4", {
 *     clientTimeout: 40,
 *     defaultsFrom: "/Common/fastL4",
 *     explicitflowMigration: "enabled",
 *     hardwareSyncookie: "enabled",
 *     idleTimeout: "200",
 *     iptosToclient: "pass-through",
 *     iptosToserver: "pass-through",
 *     keepaliveInterval: "disabled", //This cannot take enabled
 *     name: "/Common/sjfastl4profile",
 *     partition: "Common",
 * });
 * ```
 *
 * ## Import
 *
 * BIG-IP LTM fastl4 profiles can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import f5bigip:ltm/profileFastL4:ProfileFastL4 test-fastl4 /Common/test-fastl4
 * ```
 */
export class ProfileFastL4 extends pulumi.CustomResource {
    /**
     * Get an existing ProfileFastL4 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileFastL4State, opts?: pulumi.CustomResourceOptions): ProfileFastL4 {
        return new ProfileFastL4(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileFastL4:ProfileFastL4';

    /**
     * Returns true if the given object is an instance of ProfileFastL4.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileFastL4 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileFastL4.__pulumiType;
    }

    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     */
    public readonly clientTimeout!: pulumi.Output<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     */
    public readonly explicitflowMigration!: pulumi.Output<string>;
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     */
    public readonly hardwareSyncookie!: pulumi.Output<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    public readonly idleTimeout!: pulumi.Output<string>;
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    public readonly iptosToclient!: pulumi.Output<string>;
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    public readonly iptosToserver!: pulumi.Output<string>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     */
    public readonly keepaliveInterval!: pulumi.Output<string>;
    /**
     * Name of the profile_fastl4
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    public readonly partition!: pulumi.Output<string>;

    /**
     * Create a ProfileFastL4 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileFastL4Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileFastL4Args | ProfileFastL4State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileFastL4State | undefined;
            inputs["clientTimeout"] = state ? state.clientTimeout : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["explicitflowMigration"] = state ? state.explicitflowMigration : undefined;
            inputs["hardwareSyncookie"] = state ? state.hardwareSyncookie : undefined;
            inputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            inputs["iptosToclient"] = state ? state.iptosToclient : undefined;
            inputs["iptosToserver"] = state ? state.iptosToserver : undefined;
            inputs["keepaliveInterval"] = state ? state.keepaliveInterval : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
        } else {
            const args = argsOrState as ProfileFastL4Args | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["clientTimeout"] = args ? args.clientTimeout : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["explicitflowMigration"] = args ? args.explicitflowMigration : undefined;
            inputs["hardwareSyncookie"] = args ? args.hardwareSyncookie : undefined;
            inputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            inputs["iptosToclient"] = args ? args.iptosToclient : undefined;
            inputs["iptosToserver"] = args ? args.iptosToserver : undefined;
            inputs["keepaliveInterval"] = args ? args.keepaliveInterval : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileFastL4.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileFastL4 resources.
 */
export interface ProfileFastL4State {
    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     */
    readonly clientTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     */
    readonly explicitflowMigration?: pulumi.Input<string>;
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     */
    readonly hardwareSyncookie?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    readonly idleTimeout?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    readonly iptosToclient?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    readonly iptosToserver?: pulumi.Input<string>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     */
    readonly keepaliveInterval?: pulumi.Input<string>;
    /**
     * Name of the profile_fastl4
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileFastL4 resource.
 */
export interface ProfileFastL4Args {
    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     */
    readonly clientTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     */
    readonly explicitflowMigration?: pulumi.Input<string>;
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     */
    readonly hardwareSyncookie?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    readonly idleTimeout?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    readonly iptosToclient?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    readonly iptosToserver?: pulumi.Input<string>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     */
    readonly keepaliveInterval?: pulumi.Input<string>;
    /**
     * Name of the profile_fastl4
     */
    readonly name: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    readonly partition?: pulumi.Input<string>;
}
