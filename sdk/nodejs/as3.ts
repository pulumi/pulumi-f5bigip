// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.As3` provides details about bigip as3 resource
 *
 * This resource is helpful to configure as3 declarative JSON on BIG-IP.
 *
 * ## Import
 *
 * As3 resources can be imported using the partition name, e.g., ( use comma separated partition names if there are multiple partitions in as3 deployments )
 *
 * ```sh
 * $ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01
 * ```
 *
 * ```sh
 * $ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01,Sample_non_http_01
 * ```
 *
 * #### Import examples ( single and multiple partitions )
 *
 * ```sh
 * $ pulumi import f5bigip:index/as3:As3 test Sample_http_01
 * ```
 *
 * bigip_as3.test: Importing from ID "Sample_http_01"...
 *
 * bigip_as3.test: Import prepared!
 *
 *   Prepared bigip_as3 for import
 *
 * bigip_as3.test: Refreshing state... [id=Sample_http_01]
 *
 * Import successful!
 *
 * The resources that were imported are shown above. These resources are now in
 *
 * your Terraform state and will henceforth be managed by Terraform.
 *
 * $ terraform show
 *
 * bigip_as3.test:
 *
 * resource "bigip_as3" "test" {
 *
 *     as3_json      = jsonencode(
 *     
 *         {
 *     
 *             action      = "deploy"
 *     
 *             class       = "AS3"
 *     
 *             declaration = {
 *     
 *                 Sample_http_01 = {
 *     
 *                     A1    = {
 *     
 *                         class      = "Application"
 *     
 *                         jsessionid = {
 *     
 *                             class             = "Persist"
 *     
 *                             cookieMethod      = "hash"
 *     
 *                             cookieName        = "JSESSIONID"
 *     
 *                             persistenceMethod = "cookie"
 *     
 *                         }
 *     
 *                         service    = {
 *     
 *                             class              = "Service_HTTP"
 *     
 *                             persistenceMethods = [
 *     
 *                                 {
 *     
 *                                     use = "jsessionid"
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             pool               = "web_pool"
 *     
 *                             virtualAddresses   = [
 *     
 *                                 "10.0.2.10",
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         web_pool   = {
 *     
 *                             class    = "Pool"
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         "192.0.2.10",
 *     
 *                                         "192.0.2.11",
 *     
 *                                     ]
 *     
 *                                     servicePort     = 80
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             monitors = [
 *     
 *                                 "http",
 *     
 *                             ]
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class = "Tenant"
 *     
 *                 }
 *     
 *                 class          = "ADC"
 *     
 *                 id             = "UDP_DNS_Sample"
 *     
 *                 label          = "UDP_DNS_Sample"
 *     
 *                 remark         = "Sample of a UDP DNS Load Balancer Service"
 *     
 *                 schemaVersion  = "3.0.0"
 *     
 *             }
 *     
 *             persist     = true
 *     
 *         }
 *     
 *     )
 *     
 *     id            = "Sample_http_01"
 *     
 *     tenant_filter = "Sample_http_01"
 *     
 *     tenant_list   = "Sample_http_01"
 *
 * }
 *
 * ```sh
 * $ pulumi import f5bigip:index/as3:As3 test Sample_http_01,Sample_non_http_01
 * ```
 *
 * bigip_as3.test: Importing from ID "Sample_http_01,Sample_non_http_01"...
 *
 * bigip_as3.test: Import prepared!
 *
 *   Prepared bigip_as3 for import
 *
 * bigip_as3.test: Refreshing state... [id=Sample_http_01,Sample_non_http_01]
 *
 * Import successful!
 *
 * The resources that were imported are shown above. These resources are now in
 *
 * your Terraform state and will henceforth be managed by Terraform.
 *
 * $ terraform show
 *
 * bigip_as3.test:
 *
 * resource "bigip_as3" "test" {
 *
 *     as3_json      = jsonencode(
 *     
 *         {
 *     
 *             action      = "deploy"
 *     
 *             class       = "AS3"
 *     
 *             declaration = {
 *     
 *                 Sample_http_01     = {
 *     
 *                     A1    = {
 *     
 *                         class      = "Application"
 *     
 *                         jsessionid = {
 *     
 *                             class             = "Persist"
 *     
 *                             cookieMethod      = "hash"
 *     
 *                             cookieName        = "JSESSIONID"
 *     
 *                             persistenceMethod = "cookie"
 *     
 *                         }
 *     
 *                         service    = {
 *     
 *                             class              = "Service_HTTP"
 *     
 *                             persistenceMethods = [
 *     
 *                                 {
 *     
 *                                     use = "jsessionid"
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             pool               = "web_pool"
 *     
 *                             virtualAddresses   = [
 *     
 *                                 "10.0.2.10",
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         web_pool   = {
 *     
 *                             class    = "Pool"
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         "192.0.2.10",
 *     
 *                                         "192.0.2.11",
 *     
 *                                     ]
 *     
 *                                     servicePort     = 80
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             monitors = [
 *     
 *                                 "http",
 *     
 *                             ]
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class = "Tenant"
 *     
 *                 }
 *     
 *                 Sample_non_http_01 = {
 *     
 *                     DNS_Service = {
 *     
 *                         Pool1   = {
 *     
 *                             class    = "Pool"
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         "10.1.10.100",
 *     
 *                                     ]
 *     
 *                                     servicePort     = 53
 *     
 *                                 },
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         "10.1.10.101",
 *     
 *                                     ]
 *     
 *                                     servicePort     = 53
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             monitors = [
 *     
 *                                 "icmp",
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         class   = "Application"
 *     
 *                         service = {
 *     
 *                             class            = "Service_UDP"
 *     
 *                             pool             = "Pool1"
 *     
 *                             virtualAddresses = [
 *     
 *                                 "10.1.20.121",
 *     
 *                             ]
 *     
 *                             virtualPort      = 53
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class       = "Tenant"
 *     
 *                 }
 *     
 *                 class              = "ADC"
 *     
 *                 id                 = "UDP_DNS_Sample"
 *     
 *                 label              = "UDP_DNS_Sample"
 *     
 *                 remark             = "Sample of a UDP DNS Load Balancer Service"
 *     
 *                 schemaVersion      = "3.0.0"
 *     
 *             }
 *     
 *             persist     = true
 *     
 *         }
 *     
 *     )
 *     
 *     id            = "Sample_http_01,Sample_non_http_01"
 *     
 *     tenant_filter = "Sample_http_01,Sample_non_http_01"
 *     
 *     tenant_list   = "Sample_http_01,Sample_non_http_01"
 *
 * }
 *
 * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html
 */
export class As3 extends pulumi.CustomResource {
    /**
     * Get an existing As3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: As3State, opts?: pulumi.CustomResourceOptions): As3 {
        return new As3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/as3:As3';

    /**
     * Returns true if the given object is an instance of As3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is As3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === As3.__pulumiType;
    }

    /**
     * Application deployed through AS3 Declaration
     */
    public readonly applicationList!: pulumi.Output<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    public readonly as3Json!: pulumi.Output<string | undefined>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to false
     *
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_01": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_1": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.0.2.10"
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
     * "192.0.1.100",
     * "192.0.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_02": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_2": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.2.2.10"
     * ],
     * "pool": "web_pool2"
     * },
     * "web_pool2": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.2.1.100",
     * "192.2.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * },
     * "Sample_03": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_3": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool3"
     * },
     * "web_pool3": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.3.1.100",
     * "192.3.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     */
    public readonly ignoreMetadata!: pulumi.Output<boolean | undefined>;
    /**
     * Will define Perapp mode enabled on BIG-IP or not
     */
    public /*out*/ readonly perAppMode!: pulumi.Output<boolean>;
    /**
     * ID of AS3 post declaration async task
     */
    public readonly taskId!: pulumi.Output<string>;
    /**
     * If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     */
    public readonly tenantFilter!: pulumi.Output<string | undefined>;
    /**
     * Name of Tenant
     */
    public readonly tenantList!: pulumi.Output<string>;
    /**
     * Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random
     * name would be generated.
     */
    public readonly tenantName!: pulumi.Output<string | undefined>;

    /**
     * Create a As3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: As3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: As3Args | As3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as As3State | undefined;
            resourceInputs["applicationList"] = state ? state.applicationList : undefined;
            resourceInputs["as3Json"] = state ? state.as3Json : undefined;
            resourceInputs["ignoreMetadata"] = state ? state.ignoreMetadata : undefined;
            resourceInputs["perAppMode"] = state ? state.perAppMode : undefined;
            resourceInputs["taskId"] = state ? state.taskId : undefined;
            resourceInputs["tenantFilter"] = state ? state.tenantFilter : undefined;
            resourceInputs["tenantList"] = state ? state.tenantList : undefined;
            resourceInputs["tenantName"] = state ? state.tenantName : undefined;
        } else {
            const args = argsOrState as As3Args | undefined;
            resourceInputs["applicationList"] = args ? args.applicationList : undefined;
            resourceInputs["as3Json"] = args ? args.as3Json : undefined;
            resourceInputs["ignoreMetadata"] = args ? args.ignoreMetadata : undefined;
            resourceInputs["taskId"] = args ? args.taskId : undefined;
            resourceInputs["tenantFilter"] = args ? args.tenantFilter : undefined;
            resourceInputs["tenantList"] = args ? args.tenantList : undefined;
            resourceInputs["tenantName"] = args ? args.tenantName : undefined;
            resourceInputs["perAppMode"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(As3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering As3 resources.
 */
export interface As3State {
    /**
     * Application deployed through AS3 Declaration
     */
    applicationList?: pulumi.Input<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    as3Json?: pulumi.Input<string>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to false
     *
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_01": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_1": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.0.2.10"
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
     * "192.0.1.100",
     * "192.0.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_02": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_2": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.2.2.10"
     * ],
     * "pool": "web_pool2"
     * },
     * "web_pool2": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.2.1.100",
     * "192.2.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * },
     * "Sample_03": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_3": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool3"
     * },
     * "web_pool3": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.3.1.100",
     * "192.3.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     */
    ignoreMetadata?: pulumi.Input<boolean>;
    /**
     * Will define Perapp mode enabled on BIG-IP or not
     */
    perAppMode?: pulumi.Input<boolean>;
    /**
     * ID of AS3 post declaration async task
     */
    taskId?: pulumi.Input<string>;
    /**
     * If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     */
    tenantFilter?: pulumi.Input<string>;
    /**
     * Name of Tenant
     */
    tenantList?: pulumi.Input<string>;
    /**
     * Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random
     * name would be generated.
     */
    tenantName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a As3 resource.
 */
export interface As3Args {
    /**
     * Application deployed through AS3 Declaration
     */
    applicationList?: pulumi.Input<string>;
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     */
    as3Json?: pulumi.Input<string>;
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to false
     *
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_01": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_1": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.0.2.10"
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
     * "192.0.1.100",
     * "192.0.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     *
     * ```json
     *
     * {
     * "class": "AS3",
     * "action": "deploy",
     * "persist": true,
     * "declaration": {
     * "class": "ADC",
     * "schemaVersion": "3.0.0",
     * "id": "example-declaration-01",
     * "label": "Sample 1",
     * "remark": "Simple HTTP application with round robin pool",
     * "Sample_02": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_2": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.2.2.10"
     * ],
     * "pool": "web_pool2"
     * },
     * "web_pool2": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.2.1.100",
     * "192.2.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * },
     * "Sample_03": {
     * "class": "Tenant",
     * "defaultRouteDomain": 0,
     * "Application_3": {
     * "class": "Application",
     * "template": "http",
     * "serviceMain": {
     * "class": "Service_HTTP",
     * "virtualAddresses": [
     * "10.1.2.10"
     * ],
     * "pool": "web_pool3"
     * },
     * "web_pool3": {
     * "class": "Pool",
     * "monitors": [
     * "http"
     * ],
     * "members": [
     * {
     * "servicePort": 80,
     * "serverAddresses": [
     * "192.3.1.100",
     * "192.3.1.110"
     * ]
     * }
     * ]
     * }
     * }
     * }
     * }
     * }
     *
     * ```
     */
    ignoreMetadata?: pulumi.Input<boolean>;
    /**
     * ID of AS3 post declaration async task
     */
    taskId?: pulumi.Input<string>;
    /**
     * If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     */
    tenantFilter?: pulumi.Input<string>;
    /**
     * Name of Tenant
     */
    tenantList?: pulumi.Input<string>;
    /**
     * Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random
     * name would be generated.
     */
    tenantName?: pulumi.Input<string>;
}
