// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.BigIqAs3` provides details about bigiq as3 resource
 *
 * This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as std from "@pulumi/std";
 *
 * // Example Usage for json file
 * const exampletask = new f5bigip.BigIqAs3("exampletask", {
 *     bigiqAddress: "xx.xx.xxx.xx",
 *     bigiqUser: "xxxxx",
 *     bigiqPassword: "xxxxxxxxx",
 *     as3Json: std.file({
 *         input: "bigiq_example.json",
 *     }).then(invoke => invoke.result),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class BigIqAs3 extends pulumi.CustomResource {
    /**
     * Get an existing BigIqAs3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BigIqAs3State, opts?: pulumi.CustomResourceOptions): BigIqAs3 {
        return new BigIqAs3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/bigIqAs3:BigIqAs3';

    /**
     * Returns true if the given object is an instance of BigIqAs3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BigIqAs3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BigIqAs3.__pulumiType;
    }

    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    public readonly as3Json!: pulumi.Output<string>;
    /**
     * Address of the BIG-IQ to which your targer BIG-IP is attached
     */
    public readonly bigiqAddress!: pulumi.Output<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    public readonly bigiqLoginRef!: pulumi.Output<string | undefined>;
    /**
     * Password of the BIG-IQ to which your targer BIG-IP is attached
     */
    public readonly bigiqPassword!: pulumi.Output<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    public readonly bigiqPort!: pulumi.Output<string | undefined>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    public readonly bigiqTokenAuth!: pulumi.Output<boolean | undefined>;
    /**
     * User name  of the BIG-IQ to which your targer BIG-IP is attached
     */
    public readonly bigiqUser!: pulumi.Output<string>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to `true`
     *
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     *
     * ```json
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.7.0",
     * "id": "example-declaration-01",
     * "label": "Task1",
     * "remark": "Task 1 - HTTP Application Service",
     * "target": {
     * "address": "xx.xxx.xx.xxx"
     * },
     * "Task1": {
     * "class": "Tenant",
     * "MyWebApp1http": {
     * "class": "Application",
     * "template": "http",
     *
     *
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool"
     * },
     * "web_pool": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.0.2.33",
     * "192.0.2.13"
     * ],
     * "shareNodes": true
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     * ```
     *
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     *
     * >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
     */
    public readonly ignoreMetadata!: pulumi.Output<boolean | undefined>;
    /**
     * Name of Tenant
     */
    public readonly tenantList!: pulumi.Output<string>;

    /**
     * Create a BigIqAs3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BigIqAs3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BigIqAs3Args | BigIqAs3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BigIqAs3State | undefined;
            resourceInputs["as3Json"] = state ? state.as3Json : undefined;
            resourceInputs["bigiqAddress"] = state ? state.bigiqAddress : undefined;
            resourceInputs["bigiqLoginRef"] = state ? state.bigiqLoginRef : undefined;
            resourceInputs["bigiqPassword"] = state ? state.bigiqPassword : undefined;
            resourceInputs["bigiqPort"] = state ? state.bigiqPort : undefined;
            resourceInputs["bigiqTokenAuth"] = state ? state.bigiqTokenAuth : undefined;
            resourceInputs["bigiqUser"] = state ? state.bigiqUser : undefined;
            resourceInputs["ignoreMetadata"] = state ? state.ignoreMetadata : undefined;
            resourceInputs["tenantList"] = state ? state.tenantList : undefined;
        } else {
            const args = argsOrState as BigIqAs3Args | undefined;
            if ((!args || args.as3Json === undefined) && !opts.urn) {
                throw new Error("Missing required property 'as3Json'");
            }
            if ((!args || args.bigiqAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bigiqAddress'");
            }
            if ((!args || args.bigiqPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bigiqPassword'");
            }
            if ((!args || args.bigiqUser === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bigiqUser'");
            }
            resourceInputs["as3Json"] = args ? args.as3Json : undefined;
            resourceInputs["bigiqAddress"] = args ? args.bigiqAddress : undefined;
            resourceInputs["bigiqLoginRef"] = args?.bigiqLoginRef ? pulumi.secret(args.bigiqLoginRef) : undefined;
            resourceInputs["bigiqPassword"] = args?.bigiqPassword ? pulumi.secret(args.bigiqPassword) : undefined;
            resourceInputs["bigiqPort"] = args?.bigiqPort ? pulumi.secret(args.bigiqPort) : undefined;
            resourceInputs["bigiqTokenAuth"] = args?.bigiqTokenAuth ? pulumi.secret(args.bigiqTokenAuth) : undefined;
            resourceInputs["bigiqUser"] = args?.bigiqUser ? pulumi.secret(args.bigiqUser) : undefined;
            resourceInputs["ignoreMetadata"] = args ? args.ignoreMetadata : undefined;
            resourceInputs["tenantList"] = args ? args.tenantList : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["bigiqLoginRef", "bigiqPassword", "bigiqPort", "bigiqTokenAuth", "bigiqUser"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(BigIqAs3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BigIqAs3 resources.
 */
export interface BigIqAs3State {
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    as3Json?: pulumi.Input<string>;
    /**
     * Address of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqAddress?: pulumi.Input<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    bigiqLoginRef?: pulumi.Input<string>;
    /**
     * Password of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqPassword?: pulumi.Input<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    bigiqPort?: pulumi.Input<string>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    bigiqTokenAuth?: pulumi.Input<boolean>;
    /**
     * User name  of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqUser?: pulumi.Input<string>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to `true`
     *
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     *
     * ```json
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.7.0",
     * "id": "example-declaration-01",
     * "label": "Task1",
     * "remark": "Task 1 - HTTP Application Service",
     * "target": {
     * "address": "xx.xxx.xx.xxx"
     * },
     * "Task1": {
     * "class": "Tenant",
     * "MyWebApp1http": {
     * "class": "Application",
     * "template": "http",
     *
     *
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool"
     * },
     * "web_pool": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.0.2.33",
     * "192.0.2.13"
     * ],
     * "shareNodes": true
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     * ```
     *
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     *
     * >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
     */
    ignoreMetadata?: pulumi.Input<boolean>;
    /**
     * Name of Tenant
     */
    tenantList?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BigIqAs3 resource.
 */
export interface BigIqAs3Args {
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    as3Json: pulumi.Input<string>;
    /**
     * Address of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqAddress: pulumi.Input<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    bigiqLoginRef?: pulumi.Input<string>;
    /**
     * Password of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqPassword: pulumi.Input<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    bigiqPort?: pulumi.Input<string>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    bigiqTokenAuth?: pulumi.Input<boolean>;
    /**
     * User name  of the BIG-IQ to which your targer BIG-IP is attached
     */
    bigiqUser: pulumi.Input<string>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to `true`
     *
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     *
     * ```json
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.7.0",
     * "id": "example-declaration-01",
     * "label": "Task1",
     * "remark": "Task 1 - HTTP Application Service",
     * "target": {
     * "address": "xx.xxx.xx.xxx"
     * },
     * "Task1": {
     * "class": "Tenant",
     * "MyWebApp1http": {
     * "class": "Application",
     * "template": "http",
     *
     *
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool"
     * },
     * "web_pool": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.0.2.33",
     * "192.0.2.13"
     * ],
     * "shareNodes": true
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     * ```
     *
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     *
     * >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
     */
    ignoreMetadata?: pulumi.Input<boolean>;
    /**
     * Name of Tenant
     */
    tenantList?: pulumi.Input<string>;
}
