// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_ltm_monitor` Configures a custom monitor for use by health checks.
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
 * const monitor = new f5bigip.ltm.Monitor("monitor", {
 *     destination: "1.2.3.4:1234",
 *     interval: 999,
 *     name: "/Common/terraform_monitor",
 *     parent: "/Common/http",
 *     send: "GET /some/path\n",
 *     timeout: 999,
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorState, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, <any>state, { ...opts, id: id });
    }

    /**
     * ftp adaptive
     */
    public readonly adaptive: pulumi.Output<string | undefined>;
    /**
     * Integer value
     */
    public readonly adaptiveLimit: pulumi.Output<number | undefined>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    public readonly compatibility: pulumi.Output<string | undefined>;
    /**
     * Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
     */
    public readonly defaultsFrom: pulumi.Output<string | undefined>;
    /**
     * Specify an alias address for monitoring
     */
    public readonly destination: pulumi.Output<string | undefined>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    public readonly filename: pulumi.Output<string | undefined>;
    /**
     * Check interval in seconds
     */
    public readonly interval: pulumi.Output<number | undefined>;
    public readonly ipDscp: pulumi.Output<number | undefined>;
    public readonly manualResume: pulumi.Output<string | undefined>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    public readonly mode: pulumi.Output<string | undefined>;
    /**
     * Name of the monitor
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Existing LTM monitor to inherit from
     */
    public readonly parent: pulumi.Output<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    public readonly password: pulumi.Output<string | undefined>;
    /**
     * Expected response string
     */
    public readonly receive: pulumi.Output<string | undefined>;
    /**
     * Expected response string.
     */
    public readonly receiveDisable: pulumi.Output<string | undefined>;
    public readonly reverse: pulumi.Output<string | undefined>;
    /**
     * Request string to send
     */
    public readonly send: pulumi.Output<string | undefined>;
    /**
     * Time in seconds
     */
    public readonly timeUntilUp: pulumi.Output<number | undefined>;
    /**
     * Timeout in seconds
     */
    public readonly timeout: pulumi.Output<number | undefined>;
    public readonly transparent: pulumi.Output<string | undefined>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    public readonly username: pulumi.Output<string | undefined>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorArgs | MonitorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: MonitorState = argsOrState as MonitorState | undefined;
            inputs["adaptive"] = state ? state.adaptive : undefined;
            inputs["adaptiveLimit"] = state ? state.adaptiveLimit : undefined;
            inputs["compatibility"] = state ? state.compatibility : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["interval"] = state ? state.interval : undefined;
            inputs["ipDscp"] = state ? state.ipDscp : undefined;
            inputs["manualResume"] = state ? state.manualResume : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["receive"] = state ? state.receive : undefined;
            inputs["receiveDisable"] = state ? state.receiveDisable : undefined;
            inputs["reverse"] = state ? state.reverse : undefined;
            inputs["send"] = state ? state.send : undefined;
            inputs["timeUntilUp"] = state ? state.timeUntilUp : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["transparent"] = state ? state.transparent : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as MonitorArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["adaptive"] = args ? args.adaptive : undefined;
            inputs["adaptiveLimit"] = args ? args.adaptiveLimit : undefined;
            inputs["compatibility"] = args ? args.compatibility : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["interval"] = args ? args.interval : undefined;
            inputs["ipDscp"] = args ? args.ipDscp : undefined;
            inputs["manualResume"] = args ? args.manualResume : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["receive"] = args ? args.receive : undefined;
            inputs["receiveDisable"] = args ? args.receiveDisable : undefined;
            inputs["reverse"] = args ? args.reverse : undefined;
            inputs["send"] = args ? args.send : undefined;
            inputs["timeUntilUp"] = args ? args.timeUntilUp : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["transparent"] = args ? args.transparent : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        super("f5bigip:ltm/monitor:Monitor", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Monitor resources.
 */
export interface MonitorState {
    /**
     * ftp adaptive
     */
    readonly adaptive?: pulumi.Input<string>;
    /**
     * Integer value
     */
    readonly adaptiveLimit?: pulumi.Input<number>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    readonly compatibility?: pulumi.Input<string>;
    /**
     * Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specify an alias address for monitoring
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * Check interval in seconds
     */
    readonly interval?: pulumi.Input<number>;
    readonly ipDscp?: pulumi.Input<number>;
    readonly manualResume?: pulumi.Input<string>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Name of the monitor
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Existing LTM monitor to inherit from
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Expected response string
     */
    readonly receive?: pulumi.Input<string>;
    /**
     * Expected response string.
     */
    readonly receiveDisable?: pulumi.Input<string>;
    readonly reverse?: pulumi.Input<string>;
    /**
     * Request string to send
     */
    readonly send?: pulumi.Input<string>;
    /**
     * Time in seconds
     */
    readonly timeUntilUp?: pulumi.Input<number>;
    /**
     * Timeout in seconds
     */
    readonly timeout?: pulumi.Input<number>;
    readonly transparent?: pulumi.Input<string>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * ftp adaptive
     */
    readonly adaptive?: pulumi.Input<string>;
    /**
     * Integer value
     */
    readonly adaptiveLimit?: pulumi.Input<number>;
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
     */
    readonly compatibility?: pulumi.Input<string>;
    /**
     * Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Specify an alias address for monitoring
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     */
    readonly filename?: pulumi.Input<string>;
    /**
     * Check interval in seconds
     */
    readonly interval?: pulumi.Input<number>;
    readonly ipDscp?: pulumi.Input<number>;
    readonly manualResume?: pulumi.Input<string>;
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Name of the monitor
     */
    readonly name: pulumi.Input<string>;
    /**
     * Existing LTM monitor to inherit from
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Specifies the password if the monitored target requires authentication
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Expected response string
     */
    readonly receive?: pulumi.Input<string>;
    /**
     * Expected response string.
     */
    readonly receiveDisable?: pulumi.Input<string>;
    readonly reverse?: pulumi.Input<string>;
    /**
     * Request string to send
     */
    readonly send?: pulumi.Input<string>;
    /**
     * Time in seconds
     */
    readonly timeUntilUp?: pulumi.Input<number>;
    /**
     * Timeout in seconds
     */
    readonly timeout?: pulumi.Input<number>;
    readonly transparent?: pulumi.Input<string>;
    /**
     * Specifies the user name if the monitored target requires authentication
     */
    readonly username?: pulumi.Input<string>;
}
