// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.NetTunnel` Manages a tunnel configuration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const example1 = new f5bigip.NetTunnel("example1", {
 *     name: "example1",
 *     localAddress: "192.16.81.240",
 *     profile: "/Common/dslite",
 * });
 * ```
 */
export class NetTunnel extends pulumi.CustomResource {
    /**
     * Get an existing NetTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetTunnelState, opts?: pulumi.CustomResourceOptions): NetTunnel {
        return new NetTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/netTunnel:NetTunnel';

    /**
     * Returns true if the given object is an instance of NetTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetTunnel.__pulumiType;
    }

    /**
     * The application service that the object belongs to
     */
    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether auto lasthop is enabled or not
     */
    public readonly autoLastHop!: pulumi.Output<string | undefined>;
    /**
     * User defined description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies an idle timeout for wildcard tunnels in seconds
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    /**
     * The key field may represent different values depending on the type of the tunnel
     */
    public readonly key!: pulumi.Output<number | undefined>;
    /**
     * Specifies a local IP address. This option is required
     */
    public readonly localAddress!: pulumi.Output<string>;
    /**
     * Specifies how the tunnel carries traffic
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * Specifies the maximum transmission unit (MTU) of the tunnel
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * Name of the tunnel
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Displays the admin-partition within which this component resides
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Specifies the profile that you want to associate with the tunnel
     */
    public readonly profile!: pulumi.Output<string>;
    /**
     * Specifies a remote IP address
     */
    public readonly remoteAddress!: pulumi.Output<string | undefined>;
    /**
     * Specifies a secondary non-floating IP address when the local-address is set to a floating address
     */
    public readonly secondaryAddress!: pulumi.Output<string | undefined>;
    /**
     * Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
     */
    public readonly tos!: pulumi.Output<string | undefined>;
    /**
     * Specifies a traffic-group for use with the tunnel
     */
    public readonly trafficGroup!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables the tunnel to be transparent
     */
    public readonly transparent!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
     */
    public readonly usePmtu!: pulumi.Output<string | undefined>;

    /**
     * Create a NetTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetTunnelArgs | NetTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetTunnelState | undefined;
            resourceInputs["appService"] = state ? state.appService : undefined;
            resourceInputs["autoLastHop"] = state ? state.autoLastHop : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["remoteAddress"] = state ? state.remoteAddress : undefined;
            resourceInputs["secondaryAddress"] = state ? state.secondaryAddress : undefined;
            resourceInputs["tos"] = state ? state.tos : undefined;
            resourceInputs["trafficGroup"] = state ? state.trafficGroup : undefined;
            resourceInputs["transparent"] = state ? state.transparent : undefined;
            resourceInputs["usePmtu"] = state ? state.usePmtu : undefined;
        } else {
            const args = argsOrState as NetTunnelArgs | undefined;
            if ((!args || args.localAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localAddress'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.profile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profile'");
            }
            resourceInputs["appService"] = args ? args.appService : undefined;
            resourceInputs["autoLastHop"] = args ? args.autoLastHop : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["remoteAddress"] = args ? args.remoteAddress : undefined;
            resourceInputs["secondaryAddress"] = args ? args.secondaryAddress : undefined;
            resourceInputs["tos"] = args ? args.tos : undefined;
            resourceInputs["trafficGroup"] = args ? args.trafficGroup : undefined;
            resourceInputs["transparent"] = args ? args.transparent : undefined;
            resourceInputs["usePmtu"] = args ? args.usePmtu : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetTunnel resources.
 */
export interface NetTunnelState {
    /**
     * The application service that the object belongs to
     */
    appService?: pulumi.Input<string>;
    /**
     * Specifies whether auto lasthop is enabled or not
     */
    autoLastHop?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout for wildcard tunnels in seconds
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * The key field may represent different values depending on the type of the tunnel
     */
    key?: pulumi.Input<number>;
    /**
     * Specifies a local IP address. This option is required
     */
    localAddress?: pulumi.Input<string>;
    /**
     * Specifies how the tunnel carries traffic
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the maximum transmission unit (MTU) of the tunnel
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the tunnel
     */
    name?: pulumi.Input<string>;
    /**
     * Displays the admin-partition within which this component resides
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to associate with the tunnel
     */
    profile?: pulumi.Input<string>;
    /**
     * Specifies a remote IP address
     */
    remoteAddress?: pulumi.Input<string>;
    /**
     * Specifies a secondary non-floating IP address when the local-address is set to a floating address
     */
    secondaryAddress?: pulumi.Input<string>;
    /**
     * Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
     */
    tos?: pulumi.Input<string>;
    /**
     * Specifies a traffic-group for use with the tunnel
     */
    trafficGroup?: pulumi.Input<string>;
    /**
     * Enables or disables the tunnel to be transparent
     */
    transparent?: pulumi.Input<string>;
    /**
     * Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
     */
    usePmtu?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetTunnel resource.
 */
export interface NetTunnelArgs {
    /**
     * The application service that the object belongs to
     */
    appService?: pulumi.Input<string>;
    /**
     * Specifies whether auto lasthop is enabled or not
     */
    autoLastHop?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an idle timeout for wildcard tunnels in seconds
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * The key field may represent different values depending on the type of the tunnel
     */
    key?: pulumi.Input<number>;
    /**
     * Specifies a local IP address. This option is required
     */
    localAddress: pulumi.Input<string>;
    /**
     * Specifies how the tunnel carries traffic
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the maximum transmission unit (MTU) of the tunnel
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the tunnel
     */
    name: pulumi.Input<string>;
    /**
     * Displays the admin-partition within which this component resides
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to associate with the tunnel
     */
    profile: pulumi.Input<string>;
    /**
     * Specifies a remote IP address
     */
    remoteAddress?: pulumi.Input<string>;
    /**
     * Specifies a secondary non-floating IP address when the local-address is set to a floating address
     */
    secondaryAddress?: pulumi.Input<string>;
    /**
     * Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
     */
    tos?: pulumi.Input<string>;
    /**
     * Specifies a traffic-group for use with the tunnel
     */
    trafficGroup?: pulumi.Input<string>;
    /**
     * Enables or disables the tunnel to be transparent
     */
    transparent?: pulumi.Input<string>;
    /**
     * Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
     */
    usePmtu?: pulumi.Input<string>;
}
