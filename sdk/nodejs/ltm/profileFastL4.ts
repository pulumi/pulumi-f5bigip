// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileFastL4` Configures a custom LTM fastL4 profile for use by health checks.
 *
 * Resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
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
     * Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
     */
    public readonly lateBinding!: pulumi.Output<string>;
    /**
     * Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
     */
    public readonly looseClose!: pulumi.Output<string>;
    /**
     * Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
     */
    public readonly looseInitiation!: pulumi.Output<string>;
    /**
     * Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * name of partition
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
     */
    public readonly receiveWindowsize!: pulumi.Output<number>;
    /**
     * Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
     */
    public readonly tcpHandshakeTimeout!: pulumi.Output<string>;

    /**
     * Create a ProfileFastL4 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileFastL4Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileFastL4Args | ProfileFastL4State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileFastL4State | undefined;
            resourceInputs["clientTimeout"] = state ? state.clientTimeout : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["explicitflowMigration"] = state ? state.explicitflowMigration : undefined;
            resourceInputs["hardwareSyncookie"] = state ? state.hardwareSyncookie : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["iptosToclient"] = state ? state.iptosToclient : undefined;
            resourceInputs["iptosToserver"] = state ? state.iptosToserver : undefined;
            resourceInputs["keepaliveInterval"] = state ? state.keepaliveInterval : undefined;
            resourceInputs["lateBinding"] = state ? state.lateBinding : undefined;
            resourceInputs["looseClose"] = state ? state.looseClose : undefined;
            resourceInputs["looseInitiation"] = state ? state.looseInitiation : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["receiveWindowsize"] = state ? state.receiveWindowsize : undefined;
            resourceInputs["tcpHandshakeTimeout"] = state ? state.tcpHandshakeTimeout : undefined;
        } else {
            const args = argsOrState as ProfileFastL4Args | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["clientTimeout"] = args ? args.clientTimeout : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["explicitflowMigration"] = args ? args.explicitflowMigration : undefined;
            resourceInputs["hardwareSyncookie"] = args ? args.hardwareSyncookie : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["iptosToclient"] = args ? args.iptosToclient : undefined;
            resourceInputs["iptosToserver"] = args ? args.iptosToserver : undefined;
            resourceInputs["keepaliveInterval"] = args ? args.keepaliveInterval : undefined;
            resourceInputs["lateBinding"] = args ? args.lateBinding : undefined;
            resourceInputs["looseClose"] = args ? args.looseClose : undefined;
            resourceInputs["looseInitiation"] = args ? args.looseInitiation : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["receiveWindowsize"] = args ? args.receiveWindowsize : undefined;
            resourceInputs["tcpHandshakeTimeout"] = args ? args.tcpHandshakeTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProfileFastL4.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileFastL4 resources.
 */
export interface ProfileFastL4State {
    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     */
    clientTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     */
    explicitflowMigration?: pulumi.Input<string>;
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     */
    hardwareSyncookie?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    idleTimeout?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    iptosToclient?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    iptosToserver?: pulumi.Input<string>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     */
    keepaliveInterval?: pulumi.Input<string>;
    /**
     * Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
     */
    lateBinding?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
     */
    looseClose?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
     */
    looseInitiation?: pulumi.Input<string>;
    /**
     * Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
     */
    name?: pulumi.Input<string>;
    /**
     * name of partition
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
     */
    receiveWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
     */
    tcpHandshakeTimeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileFastL4 resource.
 */
export interface ProfileFastL4Args {
    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     */
    clientTimeout?: pulumi.Input<number>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     */
    explicitflowMigration?: pulumi.Input<string>;
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     */
    hardwareSyncookie?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     */
    idleTimeout?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    iptosToclient?: pulumi.Input<string>;
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     */
    iptosToserver?: pulumi.Input<string>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     */
    keepaliveInterval?: pulumi.Input<string>;
    /**
     * Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
     */
    lateBinding?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
     */
    looseClose?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
     */
    looseInitiation?: pulumi.Input<string>;
    /**
     * Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
     */
    name: pulumi.Input<string>;
    /**
     * name of partition
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
     */
    receiveWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
     */
    tcpHandshakeTimeout?: pulumi.Input<string>;
}
