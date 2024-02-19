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
    /// `f5bigip.As3` provides details about bigip as3 resource
    /// 
    /// This resource is helpful to configure as3 declarative JSON on BIG-IP.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Example Usage for json file
    ///     var as3_example1As3 = new F5BigIP.As3("as3-example1As3", new()
    ///     {
    ///         As3Json = File.ReadAllText("example1.json"),
    ///     });
    /// 
    ///     // Example Usage for json file with tenant filter
    ///     var as3_example1Index_as3As3 = new F5BigIP.As3("as3-example1Index/as3As3", new()
    ///     {
    ///         As3Json = File.ReadAllText("example2.json"),
    ///         TenantFilter = "Sample_03",
    ///     });
    /// 
    /// });
    /// ```
    /// 
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
    ///  bigip_as3.test: Importing from ID "Sample_http_01"...
    /// 
    ///  bigip_as3.test: Import prepared!
    /// 
    ///  Prepared bigip_as3 for import
    /// 
    ///  bigip_as3.test: Refreshing state... [id=Sample_http_01]
    /// 
    ///  Import successful!
    /// 
    ///  The resources that were imported are shown above. These resources are now in
    /// 
    ///  your Terraform state and will henceforth be managed by Terraform.
    /// 
    ///  $ terraform show
    /// 
    ///  bigip_as3.test:
    /// 
    ///  resource "bigip_as3" "test" {
    /// 
    ///  as3_json
    /// 
    /// = jsonencode(
    /// 
    ///  {
    /// 
    ///  action
    /// 
    /// = "deploy"
    /// 
    ///  class
    /// 
    ///  = "AS3"
    /// 
    ///  declaration = {
    /// 
    ///  Sample_http_01 = {
    /// 
    ///  A1
    /// 
    /// = {
    /// 
    ///  class
    /// 
    /// = "Application"
    /// 
    ///  jsessionid = {
    /// 
    ///  class
    /// 
    ///  = "Persist"
    /// 
    ///  cookieMethod
    /// 
    /// = "hash"
    /// 
    ///  cookieName
    /// 
    /// = "JSESSIONID"
    /// 
    ///  persistenceMethod = "cookie"
    /// 
    ///  }
    /// 
    ///  service
    /// 
    /// = {
    /// 
    ///  class
    /// 
    /// = "Service_HTTP"
    /// 
    ///  persistenceMethods = [
    /// 
    ///  {
    /// 
    ///  use = "jsessionid"
    /// 
    ///  },
    /// 
    ///  ]
    /// 
    ///  pool
    /// 
    ///  = "web_pool"
    /// 
    ///  virtualAddresses
    /// 
    ///  = [
    /// 
    ///  "10.0.2.10",
    /// 
    ///  ]
    /// 
    ///  }
    /// 
    ///  web_pool
    /// 
    ///  = {
    /// 
    ///  class
    /// 
    /// = "Pool"
    /// 
    ///  members
    /// 
    /// = [
    /// 
    ///  {
    /// 
    ///  serverAddresses = [
    /// 
    ///  "192.0.2.10",
    /// 
    ///  "192.0.2.11",
    /// 
    ///  ]
    /// 
    ///  servicePort
    /// 
    ///  = 80
    /// 
    ///  },
    /// 
    ///  ]
    /// 
    ///  monitors = [
    /// 
    ///  "http",
    /// 
    ///  ]
    /// 
    ///  }
    /// 
    ///  }
    /// 
    ///  class = "Tenant"
    /// 
    ///  }
    /// 
    ///  class
    /// 
    /// = "ADC"
    /// 
    ///  id
    /// 
    ///  = "UDP_DNS_Sample"
    /// 
    ///  label
    /// 
    /// = "UDP_DNS_Sample"
    /// 
    ///  remark
    /// 
    ///  = "Sample of a UDP DNS Load Balancer Service"
    /// 
    ///  schemaVersion
    /// 
    /// = "3.0.0"
    /// 
    ///  }
    /// 
    ///  persist
    /// 
    ///  = true
    /// 
    ///  }
    /// 
    ///  )
    /// 
    ///  id
    /// 
    /// = "Sample_http_01"
    /// 
    ///  tenant_filter = "Sample_http_01"
    /// 
    ///  tenant_list
    /// 
    ///  = "Sample_http_01"
    /// 
    ///  }
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:index/as3:As3 test Sample_http_01,Sample_non_http_01
    /// ```
    /// 
    ///  bigip_as3.test: Importing from ID "Sample_http_01,Sample_non_http_01"...
    /// 
    ///  bigip_as3.test: Import prepared!
    /// 
    ///  Prepared bigip_as3 for import
    /// 
    ///  bigip_as3.test: Refreshing state... [id=Sample_http_01,Sample_non_http_01]
    /// 
    ///  Import successful!
    /// 
    ///  The resources that were imported are shown above. These resources are now in
    /// 
    ///  your Terraform state and will henceforth be managed by Terraform.
    /// 
    ///  $ terraform show
    /// 
    ///  bigip_as3.test:
    /// 
    ///  resource "bigip_as3" "test" {
    /// 
    ///  as3_json
    /// 
    /// = jsonencode(
    /// 
    ///  {
    /// 
    ///  action
    /// 
    /// = "deploy"
    /// 
    ///  class
    /// 
    ///  = "AS3"
    /// 
    ///  declaration = {
    /// 
    ///  Sample_http_01
    /// 
    ///  = {
    /// 
    ///  A1
    /// 
    /// = {
    /// 
    ///  class
    /// 
    /// = "Application"
    /// 
    ///  jsessionid = {
    /// 
    ///  class
    /// 
    ///  = "Persist"
    /// 
    ///  cookieMethod
    /// 
    /// = "hash"
    /// 
    ///  cookieName
    /// 
    /// = "JSESSIONID"
    /// 
    ///  persistenceMethod = "cookie"
    /// 
    ///  }
    /// 
    ///  service
    /// 
    /// = {
    /// 
    ///  class
    /// 
    /// = "Service_HTTP"
    /// 
    ///  persistenceMethods = [
    /// 
    ///  {
    /// 
    ///  use = "jsessionid"
    /// 
    ///  },
    /// 
    ///  ]
    /// 
    ///  pool
    /// 
    ///  = "web_pool"
    /// 
    ///  virtualAddresses
    /// 
    ///  = [
    /// 
    ///  "10.0.2.10",
    /// 
    ///  ]
    /// 
    ///  }
    /// 
    ///  web_pool
    /// 
    ///  = {
    /// 
    ///  class
    /// 
    /// = "Pool"
    /// 
    ///  members
    /// 
    /// = [
    /// 
    ///  {
    /// 
    ///  serverAddresses = [
    /// 
    ///  "192.0.2.10",
    /// 
    ///  "192.0.2.11",
    /// 
    ///  ]
    /// 
    ///  servicePort
    /// 
    ///  = 80
    /// 
    ///  },
    /// 
    ///  ]
    /// 
    ///  monitors = [
    /// 
    ///  "http",
    /// 
    ///  ]
    /// 
    ///  }
    /// 
    ///  }
    /// 
    ///  class = "Tenant"
    /// 
    ///  }
    /// 
    ///  Sample_non_http_01 = {
    /// 
    ///  DNS_Service = {
    /// 
    ///  Pool1
    /// 
    ///  = {
    /// 
    ///  class
    /// 
    /// = "Pool"
    /// 
    ///  members
    /// 
    /// = [
    /// 
    ///  {
    /// 
    ///  serverAddresses = [
    /// 
    ///  "10.1.10.100",
    /// 
    ///  ]
    /// 
    ///  servicePort
    /// 
    ///  = 53
    /// 
    ///  },
    /// 
    ///  {
    /// 
    ///  serverAddresses = [
    /// 
    ///  "10.1.10.101",
    /// 
    ///  ]
    /// 
    ///  servicePort
    /// 
    ///  = 53
    /// 
    ///  },
    /// 
    ///  ]
    /// 
    ///  monitors = [
    /// 
    ///  "icmp",
    /// 
    ///  ]
    /// 
    ///  }
    /// 
    ///  class
    /// 
    ///  = "Application"
    /// 
    ///  service = {
    /// 
    ///  class
    /// 
    /// = "Service_UDP"
    /// 
    ///  pool
    /// 
    ///  = "Pool1"
    /// 
    ///  virtualAddresses = [
    /// 
    ///  "10.1.20.121",
    /// 
    ///  ]
    /// 
    ///  virtualPort
    /// 
    /// = 53
    /// 
    ///  }
    /// 
    ///  }
    /// 
    ///  class
    /// 
    ///  = "Tenant"
    /// 
    ///  }
    /// 
    ///  class
    /// 
    /// = "ADC"
    /// 
    ///  id
    /// 
    ///  = "UDP_DNS_Sample"
    /// 
    ///  label
    /// 
    /// = "UDP_DNS_Sample"
    /// 
    ///  remark
    /// 
    ///  = "Sample of a UDP DNS Load Balancer Service"
    /// 
    ///  schemaVersion
    /// 
    /// = "3.0.0"
    /// 
    ///  }
    /// 
    ///  persist
    /// 
    ///  = true
    /// 
    ///  }
    /// 
    ///  )
    /// 
    ///  id
    /// 
    /// = "Sample_http_01,Sample_non_http_01"
    /// 
    ///  tenant_filter = "Sample_http_01,Sample_non_http_01"
    /// 
    ///  tenant_list
    /// 
    ///  = "Sample_http_01,Sample_non_http_01"
    /// 
    ///  }
    /// 
    ///  * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/as3:As3")]
    public partial class As3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application deployed through AS3 Declaration
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
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        [Output("ignoreMetadata")]
        public Output<bool?> IgnoreMetadata { get; private set; } = null!;

        /// <summary>
        /// Will define Perapp mode enabled on BIG-IP or not
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
        public Output<string?> TenantFilter { get; private set; } = null!;

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Output("tenantList")]
        public Output<string> TenantList { get; private set; } = null!;

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Output("tenantName")]
        public Output<string?> TenantName { get; private set; } = null!;


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
        /// Application deployed through AS3 Declaration
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
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
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
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        /// <summary>
        /// Name of Tenant
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
        /// Application deployed through AS3 Declaration
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
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        [Input("ignoreMetadata")]
        public Input<bool>? IgnoreMetadata { get; set; }

        /// <summary>
        /// Will define Perapp mode enabled on BIG-IP or not
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
        /// Name of Tenant
        /// </summary>
        [Input("tenantList")]
        public Input<string>? TenantList { get; set; }

        /// <summary>
        /// Name of Tenant
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        public As3State()
        {
        }
        public static new As3State Empty => new As3State();
    }
}
