// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Monitor` Configures a custom monitor for use by health checks.
 *
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-monitor`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const monitor = new f5bigip.ltm.Monitor("monitor", {
 *     destination: "1.2.3.4:1234",
 *     interval: 998,
 *     name: "/Common/terraform_monitor",
 *     parent: "/Common/http",
 *     send: `GET /some/path
 *
 * `,
 *     timeout: 999,
 * });
 * const test_https_monitor = new f5bigip.ltm.Monitor("test-https-monitor", {
 *     interval: 999,
 *     name: "/Common/terraform_monitor",
 *     parent: "/Common/http",
 *     send: `GET /some/path
 *
 * `,
 *     sslProfile: "/Common/serverssl",
 *     timeout: 1000,
 * });
 * const test_ftp_monitor = new f5bigip.ltm.Monitor("test-ftp-monitor", {
 *     destination: "*:8008",
 *     filename: "somefile",
 *     interval: 5,
 *     name: "/Common/ftp-test",
 *     parent: "/Common/ftp",
 *     timeUntilUp: 0,
 *     timeout: 16,
 * });
 * const test_postgresql_monitor = new f5bigip.ltm.Monitor("test-postgresql-monitor", {
 *     interval: 5,
 *     name: "/Common/test-postgresql-monitor",
 *     parent: "/Common/postgresql",
 *     password: "abcd1234",
 *     receive: "Test",
 *     send: "SELECT 'Test';",
 *     timeout: 16,
 *     username: "abcd",
 * });
 * ```
 * <!--End PulumiCodeChooser -->      
 *
 * ## Importing
 *
 * An existing monitor can be imported into this resource by supplying monitor Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_monitor.monitor /Common/terraform_monitor
 * ```
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorState, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/monitor:Monitor';

    /**
     * Returns true if the given object is an instance of Monitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Monitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Monitor.__pulumiType;
    }

    /**
     * Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
     */
    public readonly adaptive!: pulumi.Output<string>;
    /**
     * Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
     */
    public readonly adaptiveLimit!: pulumi.Output<number>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    public readonly compatibility!: pulumi.Output<string | undefined>;
    /**
     * Custom parent monitor for the system to use for setting initial values for the new monitor.
     */
    public readonly customParent!: pulumi.Output<string | undefined>;
    /**
     * Specifies the database in which the user is created
     */
    public readonly database!: pulumi.Output<string | undefined>;
    /**
     * Specify an alias address for monitoring
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    public readonly filename!: pulumi.Output<string | undefined>;
    /**
     * Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
     */
    public readonly interval!: pulumi.Output<number>;
    /**
     * Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
     */
    public readonly ipDscp!: pulumi.Output<number>;
    /**
     * Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     */
    public readonly manualResume!: pulumi.Output<string>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Parent monitor for the system to use for setting initial values for the new monitor.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
     */
    public readonly receive!: pulumi.Output<string | undefined>;
    /**
     * The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
     */
    public readonly receiveDisable!: pulumi.Output<string | undefined>;
    /**
     * Instructs the system to mark the target resource down when the test is successful.
     */
    public readonly reverse!: pulumi.Output<string>;
    /**
     * Specifies the text string that the monitor sends to the target object.
     */
    public readonly send!: pulumi.Output<string>;
    /**
     * Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
     */
    public readonly sslProfile!: pulumi.Output<string | undefined>;
    /**
     * Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
     */
    public readonly timeUntilUp!: pulumi.Output<number>;
    /**
     * Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Specifies whether the monitor operates in transparent mode.
     */
    public readonly transparent!: pulumi.Output<string>;
    /**
     * Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
     */
    public readonly upInterval!: pulumi.Output<number>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorArgs | MonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitorState | undefined;
            resourceInputs["adaptive"] = state ? state.adaptive : undefined;
            resourceInputs["adaptiveLimit"] = state ? state.adaptiveLimit : undefined;
            resourceInputs["compatibility"] = state ? state.compatibility : undefined;
            resourceInputs["customParent"] = state ? state.customParent : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["ipDscp"] = state ? state.ipDscp : undefined;
            resourceInputs["manualResume"] = state ? state.manualResume : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["receive"] = state ? state.receive : undefined;
            resourceInputs["receiveDisable"] = state ? state.receiveDisable : undefined;
            resourceInputs["reverse"] = state ? state.reverse : undefined;
            resourceInputs["send"] = state ? state.send : undefined;
            resourceInputs["sslProfile"] = state ? state.sslProfile : undefined;
            resourceInputs["timeUntilUp"] = state ? state.timeUntilUp : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["transparent"] = state ? state.transparent : undefined;
            resourceInputs["upInterval"] = state ? state.upInterval : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as MonitorArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            resourceInputs["adaptive"] = args ? args.adaptive : undefined;
            resourceInputs["adaptiveLimit"] = args ? args.adaptiveLimit : undefined;
            resourceInputs["compatibility"] = args ? args.compatibility : undefined;
            resourceInputs["customParent"] = args ? args.customParent : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["ipDscp"] = args ? args.ipDscp : undefined;
            resourceInputs["manualResume"] = args ? args.manualResume : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["receive"] = args ? args.receive : undefined;
            resourceInputs["receiveDisable"] = args ? args.receiveDisable : undefined;
            resourceInputs["reverse"] = args ? args.reverse : undefined;
            resourceInputs["send"] = args ? args.send : undefined;
            resourceInputs["sslProfile"] = args ? args.sslProfile : undefined;
            resourceInputs["timeUntilUp"] = args ? args.timeUntilUp : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["transparent"] = args ? args.transparent : undefined;
            resourceInputs["upInterval"] = args ? args.upInterval : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Monitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Monitor resources.
 */
export interface MonitorState {
    /**
     * Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
     */
    adaptive?: pulumi.Input<string>;
    /**
     * Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
     */
    adaptiveLimit?: pulumi.Input<number>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    compatibility?: pulumi.Input<string>;
    /**
     * Custom parent monitor for the system to use for setting initial values for the new monitor.
     */
    customParent?: pulumi.Input<string>;
    /**
     * Specifies the database in which the user is created
     */
    database?: pulumi.Input<string>;
    /**
     * Specify an alias address for monitoring
     */
    destination?: pulumi.Input<string>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    filename?: pulumi.Input<string>;
    /**
     * Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
     */
    interval?: pulumi.Input<number>;
    /**
     * Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
     */
    ipDscp?: pulumi.Input<number>;
    /**
     * Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     */
    manualResume?: pulumi.Input<string>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
     */
    name?: pulumi.Input<string>;
    /**
     * Parent monitor for the system to use for setting initial values for the new monitor.
     */
    parent?: pulumi.Input<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
     */
    receive?: pulumi.Input<string>;
    /**
     * The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
     */
    receiveDisable?: pulumi.Input<string>;
    /**
     * Instructs the system to mark the target resource down when the test is successful.
     */
    reverse?: pulumi.Input<string>;
    /**
     * Specifies the text string that the monitor sends to the target object.
     */
    send?: pulumi.Input<string>;
    /**
     * Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
     */
    sslProfile?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
     */
    timeUntilUp?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies whether the monitor operates in transparent mode.
     */
    transparent?: pulumi.Input<string>;
    /**
     * Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
     */
    upInterval?: pulumi.Input<number>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
     */
    adaptive?: pulumi.Input<string>;
    /**
     * Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
     */
    adaptiveLimit?: pulumi.Input<number>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    compatibility?: pulumi.Input<string>;
    /**
     * Custom parent monitor for the system to use for setting initial values for the new monitor.
     */
    customParent?: pulumi.Input<string>;
    /**
     * Specifies the database in which the user is created
     */
    database?: pulumi.Input<string>;
    /**
     * Specify an alias address for monitoring
     */
    destination?: pulumi.Input<string>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    filename?: pulumi.Input<string>;
    /**
     * Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
     */
    interval?: pulumi.Input<number>;
    /**
     * Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
     */
    ipDscp?: pulumi.Input<number>;
    /**
     * Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     */
    manualResume?: pulumi.Input<string>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
     */
    name: pulumi.Input<string>;
    /**
     * Parent monitor for the system to use for setting initial values for the new monitor.
     */
    parent: pulumi.Input<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
     */
    receive?: pulumi.Input<string>;
    /**
     * The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
     */
    receiveDisable?: pulumi.Input<string>;
    /**
     * Instructs the system to mark the target resource down when the test is successful.
     */
    reverse?: pulumi.Input<string>;
    /**
     * Specifies the text string that the monitor sends to the target object.
     */
    send?: pulumi.Input<string>;
    /**
     * Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
     */
    sslProfile?: pulumi.Input<string>;
    /**
     * Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
     */
    timeUntilUp?: pulumi.Input<number>;
    /**
     * Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies whether the monitor operates in transparent mode.
     */
    transparent?: pulumi.Input<string>;
    /**
     * Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
     */
    upInterval?: pulumi.Input<number>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    username?: pulumi.Input<string>;
}
