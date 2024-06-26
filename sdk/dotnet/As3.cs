// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// ## Import
    /// 
    /// As3 resources can be imported using the partition name, e.g., ( use comma separated partition names if there are multiple partitions in as3 deployments )
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01,Sample_non_http_01
    /// ```
    /// 
    /// #### Import examples ( single and multiple partitions )
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:index/as3:As3 test Sample_http_01
    /// ```
    /// 
    /// bigip_as3.test: Importing from ID "Sample_http_01"...
    /// 
    /// bigip_as3.test: Import prepared!
    /// 
    ///   Prepared bigip_as3 for import
    /// 
    /// bigip_as3.test: Refreshing state... [id=Sample_http_01]
    /// 
    /// Import successful!
    /// 
    /// The resources that were imported are shown above. These resources are now in
    /// 
    /// your Terraform state and will henceforth be managed by Terraform.
    /// 
    /// $ terraform show
    /// 
    /// bigip_as3.test:
    /// 
    /// resource "bigip_as3" "test" {
    /// 
    ///     as3_json      = jsonencode(
    ///     
    ///         {
    ///     
    ///             action      = "deploy"
    ///     
    ///             class       = "AS3"
    ///     
    ///             declaration = {
    ///     
    ///                 Sample_http_01 = {
    ///     
    ///                     A1    = {
    ///     
    ///                         class      = "Application"
    ///     
    ///                         jsessionid = {
    ///     
    ///                             class             = "Persist"
    ///     
    ///                             cookieMethod      = "hash"
    ///     
    ///                             cookieName        = "JSESSIONID"
    ///     
    ///                             persistenceMethod = "cookie"
    ///     
    ///                         }
    ///     
    ///                         service    = {
    ///     
    ///                             class              = "Service_HTTP"
    ///     
    ///                             persistenceMethods = [
    ///     
    ///                                 {
    ///     
    ///                                     use = "jsessionid"
    ///     
    ///                                 },
    ///     
    ///                             ]
    ///     
    ///                             pool               = "web_pool"
    ///     
    ///                             virtualAddresses   = [
    ///     
    ///                                 "10.0.2.10",
    ///     
    ///                             ]
    ///     
    ///                         }
    ///     
    ///                         web_pool   = {
    ///     
    ///                             class    = "Pool"
    ///     
    ///                             members  = [
    ///     
    ///                                 {
    ///     
    ///                                     serverAddresses = [
    ///     
    ///                                         "192.0.2.10",
    ///     
    ///                                         "192.0.2.11",
    ///     
    ///                                     ]
    ///     
    ///                                     servicePort     = 80
    ///     
    ///                                 },
    ///     
    ///                             ]
    ///     
    ///                             monitors = [
    ///     
    ///                                 "http",
    ///     
    ///                             ]
    ///     
    ///                         }
    ///     
    ///                     }
    ///     
    ///                     class = "Tenant"
    ///     
    ///                 }
    ///     
    ///                 class          = "ADC"
    ///     
    ///                 id             = "UDP_DNS_Sample"
    ///     
    ///                 label          = "UDP_DNS_Sample"
    ///     
    ///                 remark         = "Sample of a UDP DNS Load Balancer Service"
    ///     
    ///                 schemaVersion  = "3.0.0"
    ///     
    ///             }
    ///     
    ///             persist     = true
    ///     
    ///         }
    ///     
    ///     )
    ///     
    ///     id            = "Sample_http_01"
    ///     
    ///     tenant_filter = "Sample_http_01"
    ///     
    ///     tenant_list   = "Sample_http_01"
    /// 
    /// }
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:index/as3:As3 test Sample_http_01,Sample_non_http_01
    /// ```
    /// 
    /// bigip_as3.test: Importing from ID "Sample_http_01,Sample_non_http_01"...
    /// 
    /// bigip_as3.test: Import prepared!
    /// 
    ///   Prepared bigip_as3 for import
    /// 
    /// bigip_as3.test: Refreshing state... [id=Sample_http_01,Sample_non_http_01]
    /// 
    /// Import successful!
    /// 
    /// The resources that were imported are shown above. These resources are now in
    /// 
    /// your Terraform state and will henceforth be managed by Terraform.
    /// 
    /// $ terraform show
    /// 
    /// bigip_as3.test:
    /// 
    /// resource "bigip_as3" "test" {
    /// 
    ///     as3_json      = jsonencode(
    ///     
    ///         {
    ///     
    ///             action      = "deploy"
    ///     
    ///             class       = "AS3"
    ///     
    ///             declaration = {
    ///     
    ///                 Sample_http_01     = {
    ///     
    ///                     A1    = {
    ///     
    ///                         class      = "Application"
    ///     
    ///                         jsessionid = {
    ///     
    ///                             class             = "Persist"
    ///     
    ///                             cookieMethod      = "hash"
    ///     
    ///                             cookieName        = "JSESSIONID"
    ///     
    ///                             persistenceMethod = "cookie"
    ///     
    ///                         }
    ///     
    ///                         service    = {
    ///     
    ///                             class              = "Service_HTTP"
    ///     
    ///                             persistenceMethods = [
    ///     
    ///                                 {
    ///     
    ///                                     use = "jsessionid"
    ///     
    ///                                 },
    ///     
    ///                             ]
    ///     
    ///                             pool               = "web_pool"
    ///     
    ///                             virtualAddresses   = [
    ///     
    ///                                 "10.0.2.10",
    ///     
    ///                             ]
    ///     
    ///                         }
    ///     
    ///                         web_pool   = {
    ///     
    ///                             class    = "Pool"
    ///     
    ///                             members  = [
    ///     
    ///                                 {
    ///     
    ///                                     serverAddresses = [
    ///     
    ///                                         "192.0.2.10",
    ///     
    ///                                         "192.0.2.11",
    ///     
    ///                                     ]
    ///     
    ///                                     servicePort     = 80
    ///     
    ///                                 },
    ///     
    ///                             ]
    ///     
    ///                             monitors = [
    ///     
    ///                                 "http",
    ///     
    ///                             ]
    ///     
    ///                         }
    ///     
    ///                     }
    ///     
    ///                     class = "Tenant"
    ///     
    ///                 }
    ///     
    ///                 Sample_non_http_01 = {
    ///     
    ///                     DNS_Service = {
    ///     
    ///                         Pool1   = {
    ///     
    ///                             class    = "Pool"
    ///     
    ///                             members  = [
    ///     
    ///                                 {
    ///     
    ///                                     serverAddresses = [
    ///     
    ///                                         "10.1.10.100",
    ///     
    ///                                     ]
    ///     
    ///                                     servicePort     = 53
    ///     
    ///                                 },
    ///     
    ///                                 {
    ///     
    ///                                     serverAddresses = [
    ///     
    ///                                         "10.1.10.101",
    ///     
    ///                                     ]
    ///     
    ///                                     servicePort     = 53
    ///     
    ///                                 },
    ///     
    ///                             ]
    ///     
    ///                             monitors = [
    ///     
    ///                                 "icmp",
    ///     
    ///                             ]
    ///     
    ///                         }
    ///     
    ///                         class   = "Application"
    ///     
    ///                         service = {
    ///     
    ///                             class            = "Service_UDP"
    ///     
    ///                             pool             = "Pool1"
    ///     
    ///                             virtualAddresses = [
    ///     
    ///                                 "10.1.20.121",
    ///     
    ///                             ]
    ///     
    ///                             virtualPort      = 53
    ///     
    ///                         }
    ///     
    ///                     }
    ///     
    ///                     class       = "Tenant"
    ///     
    ///                 }
    ///     
    ///                 class              = "ADC"
    ///     
    ///                 id                 = "UDP_DNS_Sample"
    ///     
    ///                 label              = "UDP_DNS_Sample"
    ///     
    ///                 remark             = "Sample of a UDP DNS Load Balancer Service"
    ///     
    ///                 schemaVersion      = "3.0.0"
    ///     
    ///             }
    ///     
    ///             persist     = true
    ///     
    ///         }
    ///     
    ///     )
    ///     
    ///     id            = "Sample_http_01,Sample_non_http_01"
    ///     
    ///     tenant_filter = "Sample_http_01,Sample_non_http_01"
    ///     
    ///     tenant_list   = "Sample_http_01,Sample_non_http_01"
    /// 
    /// }
    /// 
    /// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/as3:As3")]
    public partial class As3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of applications currently deployed on the Big-Ip
        /// </summary>
        [Output("applicationList")]
        public Output<string> ApplicationList { get; private set; } = null!;

        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Output("as3Json")]
        public Output<string?> As3Json { get; private set; } = null!;

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to false
        /// 
        /// * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_01": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_1": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.0.2.10"
        /// ],
        /// "pool": "web_pool"
        /// },
        /// "web_pool": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.1.100",
        /// "192.0.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_02": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_2": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.2.2.10"
        /// ],
        /// "pool": "web_pool2"
        /// },
        /// "web_pool2": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.2.1.100",
        /// "192.2.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// },
        /// "Sample_03": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_3": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.1.2.10"
        /// ],
        /// "pool": "web_pool3"
        /// },
        /// "web_pool3": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.3.1.100",
        /// "192.3.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// 
        /// * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
        /// 
        /// ```json
        /// {
        /// "Application1": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.2.1"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.2.10",
        /// "192.0.2.20"
        /// ]
        /// }
        /// ]
        /// }
        /// },
        /// "Application2": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.3.2"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.3.30",
        /// "192.0.3.40"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// ```
        /// </summary>
        [Output("ignoreMetadata")]
        public Output<bool?> IgnoreMetadata { get; private set; } = null!;

        /// <summary>
        /// Will specify whether is deployment is done via Per-Application Way or Traditional Way
        /// </summary>
        [Output("perAppMode")]
        public Output<bool> PerAppMode { get; private set; } = null!;

        /// <summary>
        /// ID of AS3 post declaration async task
        /// </summary>
        [Output("taskId")]
        public Output<string> TaskId { get; private set; } = null!;

        /// <summary>
        /// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
        /// </summary>
        [Output("tenantFilter")]
        public Output<string> TenantFilter { get; private set; } = null!;

        /// <summary>
        /// List of tenants currently deployed on the Big-Ip
        /// </summary>
        [Output("tenantList")]
        public Output<string> TenantList { get; private set; } = null!;

        /// <summary>
        /// Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
        /// </summary>
        [Output("tenantName")]
        public Output<string> TenantName { get; private set; } = null!;


        /// <summary>
        /// Create a As3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public As3(string name, As3Args? args = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/as3:As3", name, args ?? new As3Args(), MakeResourceOptions(options, ""))
        {
        }

        private As3(string name, Input<string> id, As3State? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/as3:As3", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing As3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static As3 Get(string name, Input<string> id, As3State? state = null, CustomResourceOptions? options = null)
        {
            return new As3(name, id, state, options);
        }
    }

    public sealed class As3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of applications currently deployed on the Big-Ip
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("as3Json")]
        public Input<string>? As3Json { get; set; }

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to false
        /// 
        /// * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_01": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_1": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.0.2.10"
        /// ],
        /// "pool": "web_pool"
        /// },
        /// "web_pool": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.1.100",
        /// "192.0.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_02": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_2": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.2.2.10"
        /// ],
        /// "pool": "web_pool2"
        /// },
        /// "web_pool2": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.2.1.100",
        /// "192.2.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// },
        /// "Sample_03": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_3": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.1.2.10"
        /// ],
        /// "pool": "web_pool3"
        /// },
        /// "web_pool3": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.3.1.100",
        /// "192.3.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// 
        /// * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
        /// 
        /// ```json
        /// {
        /// "Application1": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.2.1"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.2.10",
        /// "192.0.2.20"
        /// ]
        /// }
        /// ]
        /// }
        /// },
        /// "Application2": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.3.2"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.3.30",
        /// "192.0.3.40"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// ```
        /// </summary>
        [Input("ignoreMetadata")]
        public Input<bool>? IgnoreMetadata { get; set; }

        /// <summary>
        /// ID of AS3 post declaration async task
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
        /// </summary>
        [Input("tenantFilter")]
        public Input<string>? TenantFilter { get; set; }

        /// <summary>
        /// List of tenants currently deployed on the Big-Ip
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        /// <summary>
        /// Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        public As3Args()
        {
        }
        public static new As3Args Empty => new As3Args();
    }

    public sealed class As3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of applications currently deployed on the Big-Ip
        /// </summary>
        [Input("applicationList")]
        public Input<string>? ApplicationList { get; set; }

        /// <summary>
        /// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        /// </summary>
        [Input("as3Json")]
        public Input<string>? As3Json { get; set; }

        /// <summary>
        /// Set True if you want to ignore metadata changes during update. By default it is set to false
        /// 
        /// * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_01": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_1": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.0.2.10"
        /// ],
        /// "pool": "web_pool"
        /// },
        /// "web_pool": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.1.100",
        /// "192.0.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```json
        /// 
        /// {
        /// "class": "AS3",
        /// "action": "deploy",
        /// "persist": true,
        /// "declaration": {
        /// "class": "ADC",
        /// "schemaVersion": "3.0.0",
        /// "id": "example-declaration-01",
        /// "label": "Sample 1",
        /// "remark": "Simple HTTP application with round robin pool",
        /// "Sample_02": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_2": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.2.2.10"
        /// ],
        /// "pool": "web_pool2"
        /// },
        /// "web_pool2": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.2.1.100",
        /// "192.2.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// },
        /// "Sample_03": {
        /// "class": "Tenant",
        /// "defaultRouteDomain": 0,
        /// "Application_3": {
        /// "class": "Application",
        /// "template": "http",
        /// "serviceMain": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "10.1.2.10"
        /// ],
        /// "pool": "web_pool3"
        /// },
        /// "web_pool3": {
        /// "class": "Pool",
        /// "monitors": [
        /// "http"
        /// ],
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.3.1.100",
        /// "192.3.1.110"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// }
        /// 
        /// ```
        /// 
        /// * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
        /// 
        /// ```json
        /// {
        /// "Application1": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.2.1"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.2.10",
        /// "192.0.2.20"
        /// ]
        /// }
        /// ]
        /// }
        /// },
        /// "Application2": {
        /// "class": "Application",
        /// "service": {
        /// "class": "Service_HTTP",
        /// "virtualAddresses": [
        /// "192.0.3.2"
        /// ],
        /// "pool": "pool"
        /// },
        /// "pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 80,
        /// "serverAddresses": [
        /// "192.0.3.30",
        /// "192.0.3.40"
        /// ]
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// ```
        /// </summary>
        [Input("ignoreMetadata")]
        public Input<bool>? IgnoreMetadata { get; set; }

        /// <summary>
        /// Will specify whether is deployment is done via Per-Application Way or Traditional Way
        /// </summary>
        [Input("perAppMode")]
        public Input<bool>? PerAppMode { get; set; }

        /// <summary>
        /// ID of AS3 post declaration async task
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
        /// </summary>
        [Input("tenantFilter")]
        public Input<string>? TenantFilter { get; set; }

        /// <summary>
        /// List of tenants currently deployed on the Big-Ip
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        /// <summary>
        /// Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        public As3State()
        {
        }
        public static new As3State Empty => new As3State();
    }
}
