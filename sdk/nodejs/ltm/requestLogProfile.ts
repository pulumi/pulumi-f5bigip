// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.RequestLogProfile` Resource used for Configures request logging using the Request Logging profile
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const request_log_profile_tc1_child = new f5bigip.ltm.RequestLogProfile("request-log-profile-tc1-child", {
 *     name: "/Common/request-log-profile-tc1-child",
 *     defaultsFrom: request_log_profile_tc1.name,
 *     requestLogging: "disabled",
 *     requestlogPool: "/Common/pool2",
 *     requestlogErrorPool: "/Common/pool1",
 *     requestlogProtocol: "mds-tcp",
 *     requestlogErrorProtocol: "mds-tcp",
 *     responselogProtocol: "mds-tcp",
 *     responselogErrorProtocol: "mds-tcp",
 * });
 * ```
 *
 * ## Import
 *
 * BIG-IP LTM Request Log profiles can be imported using the `name`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import f5bigip:ltm/requestLogProfile:RequestLogProfile test-request-log /Common/test-request-log
 * ```
 */
export class RequestLogProfile extends pulumi.CustomResource {
    /**
     * Get an existing RequestLogProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RequestLogProfileState, opts?: pulumi.CustomResourceOptions): RequestLogProfile {
        return new RequestLogProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/requestLogProfile:RequestLogProfile';

    /**
     * Returns true if the given object is an instance of RequestLogProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RequestLogProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RequestLogProfile.__pulumiType;
    }

    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Specifies user-defined description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    public readonly proxyResponse!: pulumi.Output<string | undefined>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    public readonly proxycloseOnError!: pulumi.Output<string | undefined>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    public readonly proxyrespondOnLoggingerror!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    public readonly requestLogging!: pulumi.Output<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    public readonly requestlogErrorPool!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    public readonly requestlogErrorProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    public readonly requestlogErrorTemplate!: pulumi.Output<string | undefined>;
    /**
     * Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
     */
    public readonly requestlogPool!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    public readonly requestlogProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    public readonly requestlogTemplate!: pulumi.Output<string | undefined>;
    /**
     * Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    public readonly responseLogging!: pulumi.Output<string>;
    /**
     * Defines the pool associated with logging response errors. The default is `none`.
     */
    public readonly responselogErrorPool!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    public readonly responselogErrorProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    public readonly responselogErrorTemplate!: pulumi.Output<string | undefined>;
    /**
     * Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
     */
    public readonly responselogPool!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    public readonly responselogProtocol!: pulumi.Output<string | undefined>;
    /**
     * Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    public readonly responselogTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a RequestLogProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RequestLogProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RequestLogProfileArgs | RequestLogProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RequestLogProfileState | undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["proxyResponse"] = state ? state.proxyResponse : undefined;
            resourceInputs["proxycloseOnError"] = state ? state.proxycloseOnError : undefined;
            resourceInputs["proxyrespondOnLoggingerror"] = state ? state.proxyrespondOnLoggingerror : undefined;
            resourceInputs["requestLogging"] = state ? state.requestLogging : undefined;
            resourceInputs["requestlogErrorPool"] = state ? state.requestlogErrorPool : undefined;
            resourceInputs["requestlogErrorProtocol"] = state ? state.requestlogErrorProtocol : undefined;
            resourceInputs["requestlogErrorTemplate"] = state ? state.requestlogErrorTemplate : undefined;
            resourceInputs["requestlogPool"] = state ? state.requestlogPool : undefined;
            resourceInputs["requestlogProtocol"] = state ? state.requestlogProtocol : undefined;
            resourceInputs["requestlogTemplate"] = state ? state.requestlogTemplate : undefined;
            resourceInputs["responseLogging"] = state ? state.responseLogging : undefined;
            resourceInputs["responselogErrorPool"] = state ? state.responselogErrorPool : undefined;
            resourceInputs["responselogErrorProtocol"] = state ? state.responselogErrorProtocol : undefined;
            resourceInputs["responselogErrorTemplate"] = state ? state.responselogErrorTemplate : undefined;
            resourceInputs["responselogPool"] = state ? state.responselogPool : undefined;
            resourceInputs["responselogProtocol"] = state ? state.responselogProtocol : undefined;
            resourceInputs["responselogTemplate"] = state ? state.responselogTemplate : undefined;
        } else {
            const args = argsOrState as RequestLogProfileArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["proxyResponse"] = args ? args.proxyResponse : undefined;
            resourceInputs["proxycloseOnError"] = args ? args.proxycloseOnError : undefined;
            resourceInputs["proxyrespondOnLoggingerror"] = args ? args.proxyrespondOnLoggingerror : undefined;
            resourceInputs["requestLogging"] = args ? args.requestLogging : undefined;
            resourceInputs["requestlogErrorPool"] = args ? args.requestlogErrorPool : undefined;
            resourceInputs["requestlogErrorProtocol"] = args ? args.requestlogErrorProtocol : undefined;
            resourceInputs["requestlogErrorTemplate"] = args ? args.requestlogErrorTemplate : undefined;
            resourceInputs["requestlogPool"] = args ? args.requestlogPool : undefined;
            resourceInputs["requestlogProtocol"] = args ? args.requestlogProtocol : undefined;
            resourceInputs["requestlogTemplate"] = args ? args.requestlogTemplate : undefined;
            resourceInputs["responseLogging"] = args ? args.responseLogging : undefined;
            resourceInputs["responselogErrorPool"] = args ? args.responselogErrorPool : undefined;
            resourceInputs["responselogErrorProtocol"] = args ? args.responselogErrorProtocol : undefined;
            resourceInputs["responselogErrorTemplate"] = args ? args.responselogErrorTemplate : undefined;
            resourceInputs["responselogPool"] = args ? args.responselogPool : undefined;
            resourceInputs["responselogProtocol"] = args ? args.responselogProtocol : undefined;
            resourceInputs["responselogTemplate"] = args ? args.responselogTemplate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RequestLogProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RequestLogProfile resources.
 */
export interface RequestLogProfileState {
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies user-defined description.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
     */
    name?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxyResponse?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxycloseOnError?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxyrespondOnLoggingerror?: pulumi.Input<string>;
    /**
     * Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    requestLogging?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    requestlogErrorPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    requestlogErrorProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    requestlogErrorTemplate?: pulumi.Input<string>;
    /**
     * Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
     */
    requestlogPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    requestlogProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    requestlogTemplate?: pulumi.Input<string>;
    /**
     * Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    responseLogging?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging response errors. The default is `none`.
     */
    responselogErrorPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    responselogErrorProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    responselogErrorTemplate?: pulumi.Input<string>;
    /**
     * Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
     */
    responselogPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    responselogProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    responselogTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RequestLogProfile resource.
 */
export interface RequestLogProfileArgs {
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies user-defined description.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
     */
    name: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxyResponse?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxycloseOnError?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    proxyrespondOnLoggingerror?: pulumi.Input<string>;
    /**
     * Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    requestLogging?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging request errors. The default is None.
     */
    requestlogErrorPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    requestlogErrorProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    requestlogErrorTemplate?: pulumi.Input<string>;
    /**
     * Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
     */
    requestlogPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    requestlogProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged. More infor on requestlogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    requestlogTemplate?: pulumi.Input<string>;
    /**
     * Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     */
    responseLogging?: pulumi.Input<string>;
    /**
     * Defines the pool associated with logging response errors. The default is `none`.
     */
    responselogErrorPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    responselogErrorProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged for request errors.
     */
    responselogErrorTemplate?: pulumi.Input<string>;
    /**
     * Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
     */
    responselogPool?: pulumi.Input<string>;
    /**
     * Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     */
    responselogProtocol?: pulumi.Input<string>;
    /**
     * Specifies the directives and entries to be logged. More infor on responselogTemplate can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     */
    responselogTemplate?: pulumi.Input<string>;
}
