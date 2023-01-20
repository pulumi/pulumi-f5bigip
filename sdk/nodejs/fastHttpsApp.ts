// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `f5bigip.FastHttpsApp` This resource will create and manage FAST HTTPS applications on BIG-IP
 *
 * [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const fastHttpsApp = new f5bigip.FastHttpsApp("fastHttpsApp", {
 *     application: "fasthttpsapp",
 *     tenant: "fasthttpstenant",
 *     virtualServer: {
 *         ip: "10.30.40.44",
 *         port: 443,
 *     },
 * });
 * ```
 */
export class FastHttpsApp extends pulumi.CustomResource {
    /**
     * Get an existing FastHttpsApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FastHttpsAppState, opts?: pulumi.CustomResourceOptions): FastHttpsApp {
        return new FastHttpsApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/fastHttpsApp:FastHttpsApp';

    /**
     * Returns true if the given object is an instance of FastHttpsApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FastHttpsApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FastHttpsApp.__pulumiType;
    }

    /**
     * Name of the FAST HTTPS application.
     */
    public readonly application!: pulumi.Output<string>;
    /**
     * List of LTM Policies to be applied FAST HTTPS Application.
     */
    public readonly endpointLtmPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    public readonly existingMonitor!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing BIG-IP pool.
     */
    public readonly existingPool!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    public readonly existingSnatPool!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing TLS client profile.
     */
    public readonly existingTlsClientProfile!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing TLS server profile.
     */
    public readonly existingTlsServerProfile!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing WAF Security policy.
     */
    public readonly existingWafSecurityPolicy!: pulumi.Output<string | undefined>;
    /**
     * Json payload for FAST HTTPS application.
     */
    public /*out*/ readonly fastHttpsJson!: pulumi.Output<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    public readonly loadBalancingMode!: pulumi.Output<string | undefined>;
    /**
     * `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    public readonly monitor!: pulumi.Output<outputs.FastHttpsAppMonitor | undefined>;
    /**
     * `poolMembers` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    public readonly poolMembers!: pulumi.Output<outputs.FastHttpsAppPoolMember[] | undefined>;
    /**
     * List of security log profiles to be used for FAST application
     */
    public readonly securityLogProfiles!: pulumi.Output<string[] | undefined>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    public readonly slowRampTime!: pulumi.Output<number | undefined>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    public readonly snatPoolAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the FAST HTTPS application tenant.
     */
    public readonly tenant!: pulumi.Output<string>;
    /**
     * `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     */
    public readonly tlsClientProfile!: pulumi.Output<outputs.FastHttpsAppTlsClientProfile | undefined>;
    /**
     * `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     */
    public readonly tlsServerProfile!: pulumi.Output<outputs.FastHttpsAppTlsServerProfile | undefined>;
    /**
     * `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    public readonly virtualServer!: pulumi.Output<outputs.FastHttpsAppVirtualServer | undefined>;
    /**
     * `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     */
    public readonly wafSecurityPolicy!: pulumi.Output<outputs.FastHttpsAppWafSecurityPolicy | undefined>;

    /**
     * Create a FastHttpsApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FastHttpsAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FastHttpsAppArgs | FastHttpsAppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FastHttpsAppState | undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["endpointLtmPolicies"] = state ? state.endpointLtmPolicies : undefined;
            resourceInputs["existingMonitor"] = state ? state.existingMonitor : undefined;
            resourceInputs["existingPool"] = state ? state.existingPool : undefined;
            resourceInputs["existingSnatPool"] = state ? state.existingSnatPool : undefined;
            resourceInputs["existingTlsClientProfile"] = state ? state.existingTlsClientProfile : undefined;
            resourceInputs["existingTlsServerProfile"] = state ? state.existingTlsServerProfile : undefined;
            resourceInputs["existingWafSecurityPolicy"] = state ? state.existingWafSecurityPolicy : undefined;
            resourceInputs["fastHttpsJson"] = state ? state.fastHttpsJson : undefined;
            resourceInputs["loadBalancingMode"] = state ? state.loadBalancingMode : undefined;
            resourceInputs["monitor"] = state ? state.monitor : undefined;
            resourceInputs["poolMembers"] = state ? state.poolMembers : undefined;
            resourceInputs["securityLogProfiles"] = state ? state.securityLogProfiles : undefined;
            resourceInputs["slowRampTime"] = state ? state.slowRampTime : undefined;
            resourceInputs["snatPoolAddresses"] = state ? state.snatPoolAddresses : undefined;
            resourceInputs["tenant"] = state ? state.tenant : undefined;
            resourceInputs["tlsClientProfile"] = state ? state.tlsClientProfile : undefined;
            resourceInputs["tlsServerProfile"] = state ? state.tlsServerProfile : undefined;
            resourceInputs["virtualServer"] = state ? state.virtualServer : undefined;
            resourceInputs["wafSecurityPolicy"] = state ? state.wafSecurityPolicy : undefined;
        } else {
            const args = argsOrState as FastHttpsAppArgs | undefined;
            if ((!args || args.application === undefined) && !opts.urn) {
                throw new Error("Missing required property 'application'");
            }
            if ((!args || args.tenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenant'");
            }
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["endpointLtmPolicies"] = args ? args.endpointLtmPolicies : undefined;
            resourceInputs["existingMonitor"] = args ? args.existingMonitor : undefined;
            resourceInputs["existingPool"] = args ? args.existingPool : undefined;
            resourceInputs["existingSnatPool"] = args ? args.existingSnatPool : undefined;
            resourceInputs["existingTlsClientProfile"] = args ? args.existingTlsClientProfile : undefined;
            resourceInputs["existingTlsServerProfile"] = args ? args.existingTlsServerProfile : undefined;
            resourceInputs["existingWafSecurityPolicy"] = args ? args.existingWafSecurityPolicy : undefined;
            resourceInputs["loadBalancingMode"] = args ? args.loadBalancingMode : undefined;
            resourceInputs["monitor"] = args ? args.monitor : undefined;
            resourceInputs["poolMembers"] = args ? args.poolMembers : undefined;
            resourceInputs["securityLogProfiles"] = args ? args.securityLogProfiles : undefined;
            resourceInputs["slowRampTime"] = args ? args.slowRampTime : undefined;
            resourceInputs["snatPoolAddresses"] = args ? args.snatPoolAddresses : undefined;
            resourceInputs["tenant"] = args ? args.tenant : undefined;
            resourceInputs["tlsClientProfile"] = args ? args.tlsClientProfile : undefined;
            resourceInputs["tlsServerProfile"] = args ? args.tlsServerProfile : undefined;
            resourceInputs["virtualServer"] = args ? args.virtualServer : undefined;
            resourceInputs["wafSecurityPolicy"] = args ? args.wafSecurityPolicy : undefined;
            resourceInputs["fastHttpsJson"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FastHttpsApp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FastHttpsApp resources.
 */
export interface FastHttpsAppState {
    /**
     * Name of the FAST HTTPS application.
     */
    application?: pulumi.Input<string>;
    /**
     * List of LTM Policies to be applied FAST HTTPS Application.
     */
    endpointLtmPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    existingMonitor?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP pool.
     */
    existingPool?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    existingSnatPool?: pulumi.Input<string>;
    /**
     * Name of an existing TLS client profile.
     */
    existingTlsClientProfile?: pulumi.Input<string>;
    /**
     * Name of an existing TLS server profile.
     */
    existingTlsServerProfile?: pulumi.Input<string>;
    /**
     * Name of an existing WAF Security policy.
     */
    existingWafSecurityPolicy?: pulumi.Input<string>;
    /**
     * Json payload for FAST HTTPS application.
     */
    fastHttpsJson?: pulumi.Input<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    loadBalancingMode?: pulumi.Input<string>;
    /**
     * `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    monitor?: pulumi.Input<inputs.FastHttpsAppMonitor>;
    /**
     * `poolMembers` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    poolMembers?: pulumi.Input<pulumi.Input<inputs.FastHttpsAppPoolMember>[]>;
    /**
     * List of security log profiles to be used for FAST application
     */
    securityLogProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    slowRampTime?: pulumi.Input<number>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    snatPoolAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the FAST HTTPS application tenant.
     */
    tenant?: pulumi.Input<string>;
    /**
     * `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     */
    tlsClientProfile?: pulumi.Input<inputs.FastHttpsAppTlsClientProfile>;
    /**
     * `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     */
    tlsServerProfile?: pulumi.Input<inputs.FastHttpsAppTlsServerProfile>;
    /**
     * `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    virtualServer?: pulumi.Input<inputs.FastHttpsAppVirtualServer>;
    /**
     * `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     */
    wafSecurityPolicy?: pulumi.Input<inputs.FastHttpsAppWafSecurityPolicy>;
}

/**
 * The set of arguments for constructing a FastHttpsApp resource.
 */
export interface FastHttpsAppArgs {
    /**
     * Name of the FAST HTTPS application.
     */
    application: pulumi.Input<string>;
    /**
     * List of LTM Policies to be applied FAST HTTPS Application.
     */
    endpointLtmPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    existingMonitor?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP pool.
     */
    existingPool?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    existingSnatPool?: pulumi.Input<string>;
    /**
     * Name of an existing TLS client profile.
     */
    existingTlsClientProfile?: pulumi.Input<string>;
    /**
     * Name of an existing TLS server profile.
     */
    existingTlsServerProfile?: pulumi.Input<string>;
    /**
     * Name of an existing WAF Security policy.
     */
    existingWafSecurityPolicy?: pulumi.Input<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    loadBalancingMode?: pulumi.Input<string>;
    /**
     * `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    monitor?: pulumi.Input<inputs.FastHttpsAppMonitor>;
    /**
     * `poolMembers` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    poolMembers?: pulumi.Input<pulumi.Input<inputs.FastHttpsAppPoolMember>[]>;
    /**
     * List of security log profiles to be used for FAST application
     */
    securityLogProfiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    slowRampTime?: pulumi.Input<number>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    snatPoolAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the FAST HTTPS application tenant.
     */
    tenant: pulumi.Input<string>;
    /**
     * `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     */
    tlsClientProfile?: pulumi.Input<inputs.FastHttpsAppTlsClientProfile>;
    /**
     * `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     */
    tlsServerProfile?: pulumi.Input<inputs.FastHttpsAppTlsServerProfile>;
    /**
     * `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    virtualServer?: pulumi.Input<inputs.FastHttpsAppVirtualServer>;
    /**
     * `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     */
    wafSecurityPolicy?: pulumi.Input<inputs.FastHttpsAppWafSecurityPolicy>;
}
