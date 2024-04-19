// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test = new f5bigip.EventServiceDiscovery("test", {
 *     taskid: "~Sample_event_sd~My_app~My_pool",
 *     nodes: [
 *         {
 *             id: "newNode1",
 *             ip: "192.168.2.3",
 *             port: 8080,
 *         },
 *         {
 *             id: "newNode2",
 *             ip: "192.168.2.4",
 *             port: 8080,
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class EventServiceDiscovery extends pulumi.CustomResource {
    /**
     * Get an existing EventServiceDiscovery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventServiceDiscoveryState, opts?: pulumi.CustomResourceOptions): EventServiceDiscovery {
        return new EventServiceDiscovery(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/eventServiceDiscovery:EventServiceDiscovery';

    /**
     * Returns true if the given object is an instance of EventServiceDiscovery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventServiceDiscovery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventServiceDiscovery.__pulumiType;
    }

    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     *
     * For more information, please refer below document
     * https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
     *
     * Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
     *
     * With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
     *
     * When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
     *
     * For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
     *
     * Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
     * First we show the initial declaration to POST to the BIG-IP system.
     *
     * {
     * "class": "ADC",
     * "schemaVersion": "3.9.0",
     * "id": "Pool",
     * "Sample_event_sd": {
     * "class": "Tenant",
     * "My_app": {
     * "class": "Application",
     * "My_pool": {
     * "class": "Pool",
     * "members": [
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "static",
     * "serverAddresses": [
     * "192.0.2.2"
     * ]
     * },
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "event"
     * }
     * ]
     * }
     * }
     * }
     * }
     *
     *
     * Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
     */
    public readonly nodes!: pulumi.Output<outputs.EventServiceDiscoveryNode[] | undefined>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    public readonly taskid!: pulumi.Output<string>;

    /**
     * Create a EventServiceDiscovery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventServiceDiscoveryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventServiceDiscoveryArgs | EventServiceDiscoveryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventServiceDiscoveryState | undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["taskid"] = state ? state.taskid : undefined;
        } else {
            const args = argsOrState as EventServiceDiscoveryArgs | undefined;
            if ((!args || args.taskid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskid'");
            }
            resourceInputs["nodes"] = args ? args.nodes : undefined;
            resourceInputs["taskid"] = args ? args.taskid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventServiceDiscovery.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventServiceDiscovery resources.
 */
export interface EventServiceDiscoveryState {
    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     *
     * For more information, please refer below document
     * https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
     *
     * Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
     *
     * With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
     *
     * When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
     *
     * For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
     *
     * Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
     * First we show the initial declaration to POST to the BIG-IP system.
     *
     * {
     * "class": "ADC",
     * "schemaVersion": "3.9.0",
     * "id": "Pool",
     * "Sample_event_sd": {
     * "class": "Tenant",
     * "My_app": {
     * "class": "Application",
     * "My_pool": {
     * "class": "Pool",
     * "members": [
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "static",
     * "serverAddresses": [
     * "192.0.2.2"
     * ]
     * },
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "event"
     * }
     * ]
     * }
     * }
     * }
     * }
     *
     *
     * Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.EventServiceDiscoveryNode>[]>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    taskid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventServiceDiscovery resource.
 */
export interface EventServiceDiscoveryArgs {
    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     *
     * For more information, please refer below document
     * https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
     *
     * Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
     *
     * With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
     *
     * When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
     *
     * For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
     *
     * Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
     * First we show the initial declaration to POST to the BIG-IP system.
     *
     * {
     * "class": "ADC",
     * "schemaVersion": "3.9.0",
     * "id": "Pool",
     * "Sample_event_sd": {
     * "class": "Tenant",
     * "My_app": {
     * "class": "Application",
     * "My_pool": {
     * "class": "Pool",
     * "members": [
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "static",
     * "serverAddresses": [
     * "192.0.2.2"
     * ]
     * },
     * {
     * "servicePort": 8080,
     * "addressDiscovery": "event"
     * }
     * ]
     * }
     * }
     * }
     * }
     *
     *
     * Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.EventServiceDiscoveryNode>[]>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    taskid: pulumi.Input<string>;
}
