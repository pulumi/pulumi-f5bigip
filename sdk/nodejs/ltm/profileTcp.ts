// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileTcp` Configures a custom TCP LTM Profile for use by health checks.
 *
 * Resources should be named with their `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sanjose_tcp_lan_profile = new f5bigip.ltm.ProfileTcp("sanjose-tcp-lan-profile", {
 *     closeWaitTimeout: 5,
 *     deferredAccept: "enabled",
 *     fastOpen: "enabled",
 *     finwait2timeout: 5,
 *     finwaitTimeout: 300,
 *     idleTimeout: 200,
 *     keepaliveInterval: 1700,
 *     name: "/Common/sanjose-tcp-lan-profile",
 * });
 * ```
 */
export class ProfileTcp extends pulumi.CustomResource {
    /**
     * Get an existing ProfileTcp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileTcpState, opts?: pulumi.CustomResourceOptions): ProfileTcp {
        return new ProfileTcp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileTcp:ProfileTcp';

    /**
     * Returns true if the given object is an instance of ProfileTcp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileTcp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileTcp.__pulumiType;
    }

    /**
     * Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
     */
    public readonly closeWaitTimeout!: pulumi.Output<number>;
    /**
     * Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
     */
    public readonly congestionControl!: pulumi.Output<string | undefined>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
     */
    public readonly deferredAccept!: pulumi.Output<string>;
    /**
     * Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
     */
    public readonly delayedAcks!: pulumi.Output<string | undefined>;
    /**
     * Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
     */
    public readonly earlyRetransmit!: pulumi.Output<string | undefined>;
    /**
     * When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fastOpen` set to `enabled`, argument `verifiedAccept` can't be set to `enabled`.
     */
    public readonly fastOpen!: pulumi.Output<string>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
     */
    public readonly finwait2timeout!: pulumi.Output<number>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
     */
    public readonly finwaitTimeout!: pulumi.Output<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
     */
    public readonly idleTimeout!: pulumi.Output<number>;
    /**
     * Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
     */
    public readonly initialCongestionWindowsize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
     */
    public readonly keepaliveInterval!: pulumi.Output<number>;
    /**
     * Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
     */
    public readonly nagle!: pulumi.Output<string | undefined>;
    /**
     * Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * name of partition
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Specifies the proxy buffer level, in bytes, at which the receive window is closed.
     */
    public readonly proxybufferHigh!: pulumi.Output<number | undefined>;
    /**
     * Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
     */
    public readonly receiveWindowsize!: pulumi.Output<number | undefined>;
    /**
     * Specifies the SEND window size. The default is 131072 bytes.
     */
    public readonly sendBuffersize!: pulumi.Output<number | undefined>;
    /**
     * Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
     */
    public readonly taillossProbe!: pulumi.Output<string | undefined>;
    /**
     * Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
     */
    public readonly timewaitRecycle!: pulumi.Output<string | undefined>;
    /**
     * Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
     */
    public readonly verifiedAccept!: pulumi.Output<string | undefined>;
    /**
     * Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
     */
    public readonly zerowindowTimeout!: pulumi.Output<number | undefined>;

    /**
     * Create a ProfileTcp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileTcpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileTcpArgs | ProfileTcpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileTcpState | undefined;
            resourceInputs["closeWaitTimeout"] = state ? state.closeWaitTimeout : undefined;
            resourceInputs["congestionControl"] = state ? state.congestionControl : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["deferredAccept"] = state ? state.deferredAccept : undefined;
            resourceInputs["delayedAcks"] = state ? state.delayedAcks : undefined;
            resourceInputs["earlyRetransmit"] = state ? state.earlyRetransmit : undefined;
            resourceInputs["fastOpen"] = state ? state.fastOpen : undefined;
            resourceInputs["finwait2timeout"] = state ? state.finwait2timeout : undefined;
            resourceInputs["finwaitTimeout"] = state ? state.finwaitTimeout : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["initialCongestionWindowsize"] = state ? state.initialCongestionWindowsize : undefined;
            resourceInputs["keepaliveInterval"] = state ? state.keepaliveInterval : undefined;
            resourceInputs["nagle"] = state ? state.nagle : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["proxybufferHigh"] = state ? state.proxybufferHigh : undefined;
            resourceInputs["receiveWindowsize"] = state ? state.receiveWindowsize : undefined;
            resourceInputs["sendBuffersize"] = state ? state.sendBuffersize : undefined;
            resourceInputs["taillossProbe"] = state ? state.taillossProbe : undefined;
            resourceInputs["timewaitRecycle"] = state ? state.timewaitRecycle : undefined;
            resourceInputs["verifiedAccept"] = state ? state.verifiedAccept : undefined;
            resourceInputs["zerowindowTimeout"] = state ? state.zerowindowTimeout : undefined;
        } else {
            const args = argsOrState as ProfileTcpArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["closeWaitTimeout"] = args ? args.closeWaitTimeout : undefined;
            resourceInputs["congestionControl"] = args ? args.congestionControl : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["deferredAccept"] = args ? args.deferredAccept : undefined;
            resourceInputs["delayedAcks"] = args ? args.delayedAcks : undefined;
            resourceInputs["earlyRetransmit"] = args ? args.earlyRetransmit : undefined;
            resourceInputs["fastOpen"] = args ? args.fastOpen : undefined;
            resourceInputs["finwait2timeout"] = args ? args.finwait2timeout : undefined;
            resourceInputs["finwaitTimeout"] = args ? args.finwaitTimeout : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["initialCongestionWindowsize"] = args ? args.initialCongestionWindowsize : undefined;
            resourceInputs["keepaliveInterval"] = args ? args.keepaliveInterval : undefined;
            resourceInputs["nagle"] = args ? args.nagle : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["proxybufferHigh"] = args ? args.proxybufferHigh : undefined;
            resourceInputs["receiveWindowsize"] = args ? args.receiveWindowsize : undefined;
            resourceInputs["sendBuffersize"] = args ? args.sendBuffersize : undefined;
            resourceInputs["taillossProbe"] = args ? args.taillossProbe : undefined;
            resourceInputs["timewaitRecycle"] = args ? args.timewaitRecycle : undefined;
            resourceInputs["verifiedAccept"] = args ? args.verifiedAccept : undefined;
            resourceInputs["zerowindowTimeout"] = args ? args.zerowindowTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProfileTcp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileTcp resources.
 */
export interface ProfileTcpState {
    /**
     * Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
     */
    closeWaitTimeout?: pulumi.Input<number>;
    /**
     * Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
     */
    congestionControl?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
     */
    deferredAccept?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
     */
    delayedAcks?: pulumi.Input<string>;
    /**
     * Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
     */
    earlyRetransmit?: pulumi.Input<string>;
    /**
     * When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fastOpen` set to `enabled`, argument `verifiedAccept` can't be set to `enabled`.
     */
    fastOpen?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
     */
    finwait2timeout?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
     */
    finwaitTimeout?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
     */
    initialCongestionWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
     */
    keepaliveInterval?: pulumi.Input<number>;
    /**
     * Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
     */
    nagle?: pulumi.Input<string>;
    /**
     * Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
     */
    name?: pulumi.Input<string>;
    /**
     * name of partition
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the proxy buffer level, in bytes, at which the receive window is closed.
     */
    proxybufferHigh?: pulumi.Input<number>;
    /**
     * Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
     */
    receiveWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the SEND window size. The default is 131072 bytes.
     */
    sendBuffersize?: pulumi.Input<number>;
    /**
     * Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
     */
    taillossProbe?: pulumi.Input<string>;
    /**
     * Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
     */
    timewaitRecycle?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
     */
    verifiedAccept?: pulumi.Input<string>;
    /**
     * Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
     */
    zerowindowTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProfileTcp resource.
 */
export interface ProfileTcpArgs {
    /**
     * Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
     */
    closeWaitTimeout?: pulumi.Input<number>;
    /**
     * Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default is High Speed.
     */
    congestionControl?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
     */
    deferredAccept?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system can send fewer than one ACK (acknowledgment) segment per data segment received. By default, this setting is enabled.
     */
    delayedAcks?: pulumi.Input<string>;
    /**
     * Enabling this setting allows TCP to assume a packet is lost after fewer than the standard number of duplicate ACKs, if there is no way to send new data and generate more duplicate ACKs.
     */
    earlyRetransmit?: pulumi.Input<string>;
    /**
     * When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. Default is `enabled`. If `fastOpen` set to `enabled`, argument `verifiedAccept` can't be set to `enabled`.
     */
    fastOpen?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
     */
    finwait2timeout?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
     */
    finwaitTimeout?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * Specifies the initial congestion window size for connections to this destination. Actual window size is this value multiplied by the MSS (Maximum Segment Size) for the same connection. The default is 10. Valid values range from 0 to 64.
     */
    initialCongestionWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
     */
    keepaliveInterval?: pulumi.Input<number>;
    /**
     * Specifies whether the system applies Nagle's algorithm to reduce the number of short segments on the network.If you select Auto, the system determines whether to use Nagle's algorithm based on network conditions. By default, this setting is disabled.
     */
    nagle?: pulumi.Input<string>;
    /**
     * Name of the LTM TCP Profile,name should be `full path`. The full path is the combination of the `partition + name` (example: /Common/my-pool ) or  `partition + directory + name` of the resource  (example: /Common/test/my-pool )
     */
    name: pulumi.Input<string>;
    /**
     * name of partition
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the proxy buffer level, in bytes, at which the receive window is closed.
     */
    proxybufferHigh?: pulumi.Input<number>;
    /**
     * Specifies the maximum advertised RECEIVE window size. This value represents the maximum number of bytes to which the RECEIVE window can scale. The default is 65535 bytes.
     */
    receiveWindowsize?: pulumi.Input<number>;
    /**
     * Specifies the SEND window size. The default is 131072 bytes.
     */
    sendBuffersize?: pulumi.Input<number>;
    /**
     * Enabling this setting allows TCP to send a probe segment to trigger fast recovery instead of recovering a loss via a retransmission timeout,By default, this setting is enabled.
     */
    taillossProbe?: pulumi.Input<string>;
    /**
     * Using this setting enabled, the system can recycle a wait-state connection immediately upon receipt of a new connection request instead of having to wait until the connection times out of the wait state. By default, this setting is enabled.
     */
    timewaitRecycle?: pulumi.Input<string>;
    /**
     * Specifies, when checked (enabled), that the system can actually communicate with the server before establishing a client connection. To determine this, the system sends the server a SYN packet before responding to the client's SYN with a SYN-ACK. When unchecked, the system accepts the client connection before selecting a server to talk to. By default, this setting is `disabled`.
     */
    verifiedAccept?: pulumi.Input<string>;
    /**
     * Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window.
     */
    zerowindowTimeout?: pulumi.Input<number>;
}
