// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.As3Args;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.As3State;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * bigip_as3.test: Importing from ID &#34;Sample_http_01&#34;...
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
 * resource &#34;bigip_as3&#34; &#34;test&#34; {
 * 
 *     as3_json      = jsonencode(
 *     
 *         {
 *     
 *             action      = &#34;deploy&#34;
 *     
 *             class       = &#34;AS3&#34;
 *     
 *             declaration = {
 *     
 *                 Sample_http_01 = {
 *     
 *                     A1    = {
 *     
 *                         class      = &#34;Application&#34;
 *     
 *                         jsessionid = {
 *     
 *                             class             = &#34;Persist&#34;
 *     
 *                             cookieMethod      = &#34;hash&#34;
 *     
 *                             cookieName        = &#34;JSESSIONID&#34;
 *     
 *                             persistenceMethod = &#34;cookie&#34;
 *     
 *                         }
 *     
 *                         service    = {
 *     
 *                             class              = &#34;Service_HTTP&#34;
 *     
 *                             persistenceMethods = [
 *     
 *                                 {
 *     
 *                                     use = &#34;jsessionid&#34;
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             pool               = &#34;web_pool&#34;
 *     
 *                             virtualAddresses   = [
 *     
 *                                 &#34;10.0.2.10&#34;,
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         web_pool   = {
 *     
 *                             class    = &#34;Pool&#34;
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         &#34;192.0.2.10&#34;,
 *     
 *                                         &#34;192.0.2.11&#34;,
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
 *                                 &#34;http&#34;,
 *     
 *                             ]
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class = &#34;Tenant&#34;
 *     
 *                 }
 *     
 *                 class          = &#34;ADC&#34;
 *     
 *                 id             = &#34;UDP_DNS_Sample&#34;
 *     
 *                 label          = &#34;UDP_DNS_Sample&#34;
 *     
 *                 remark         = &#34;Sample of a UDP DNS Load Balancer Service&#34;
 *     
 *                 schemaVersion  = &#34;3.0.0&#34;
 *     
 *             }
 *     
 *             persist     = true
 *     
 *         }
 *     
 *     )
 *     
 *     id            = &#34;Sample_http_01&#34;
 *     
 *     tenant_filter = &#34;Sample_http_01&#34;
 *     
 *     tenant_list   = &#34;Sample_http_01&#34;
 * 
 * }
 * 
 * ```sh
 * $ pulumi import f5bigip:index/as3:As3 test Sample_http_01,Sample_non_http_01
 * ```
 * 
 * bigip_as3.test: Importing from ID &#34;Sample_http_01,Sample_non_http_01&#34;...
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
 * resource &#34;bigip_as3&#34; &#34;test&#34; {
 * 
 *     as3_json      = jsonencode(
 *     
 *         {
 *     
 *             action      = &#34;deploy&#34;
 *     
 *             class       = &#34;AS3&#34;
 *     
 *             declaration = {
 *     
 *                 Sample_http_01     = {
 *     
 *                     A1    = {
 *     
 *                         class      = &#34;Application&#34;
 *     
 *                         jsessionid = {
 *     
 *                             class             = &#34;Persist&#34;
 *     
 *                             cookieMethod      = &#34;hash&#34;
 *     
 *                             cookieName        = &#34;JSESSIONID&#34;
 *     
 *                             persistenceMethod = &#34;cookie&#34;
 *     
 *                         }
 *     
 *                         service    = {
 *     
 *                             class              = &#34;Service_HTTP&#34;
 *     
 *                             persistenceMethods = [
 *     
 *                                 {
 *     
 *                                     use = &#34;jsessionid&#34;
 *     
 *                                 },
 *     
 *                             ]
 *     
 *                             pool               = &#34;web_pool&#34;
 *     
 *                             virtualAddresses   = [
 *     
 *                                 &#34;10.0.2.10&#34;,
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         web_pool   = {
 *     
 *                             class    = &#34;Pool&#34;
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         &#34;192.0.2.10&#34;,
 *     
 *                                         &#34;192.0.2.11&#34;,
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
 *                                 &#34;http&#34;,
 *     
 *                             ]
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class = &#34;Tenant&#34;
 *     
 *                 }
 *     
 *                 Sample_non_http_01 = {
 *     
 *                     DNS_Service = {
 *     
 *                         Pool1   = {
 *     
 *                             class    = &#34;Pool&#34;
 *     
 *                             members  = [
 *     
 *                                 {
 *     
 *                                     serverAddresses = [
 *     
 *                                         &#34;10.1.10.100&#34;,
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
 *                                         &#34;10.1.10.101&#34;,
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
 *                                 &#34;icmp&#34;,
 *     
 *                             ]
 *     
 *                         }
 *     
 *                         class   = &#34;Application&#34;
 *     
 *                         service = {
 *     
 *                             class            = &#34;Service_UDP&#34;
 *     
 *                             pool             = &#34;Pool1&#34;
 *     
 *                             virtualAddresses = [
 *     
 *                                 &#34;10.1.20.121&#34;,
 *     
 *                             ]
 *     
 *                             virtualPort      = 53
 *     
 *                         }
 *     
 *                     }
 *     
 *                     class       = &#34;Tenant&#34;
 *     
 *                 }
 *     
 *                 class              = &#34;ADC&#34;
 *     
 *                 id                 = &#34;UDP_DNS_Sample&#34;
 *     
 *                 label              = &#34;UDP_DNS_Sample&#34;
 *     
 *                 remark             = &#34;Sample of a UDP DNS Load Balancer Service&#34;
 *     
 *                 schemaVersion      = &#34;3.0.0&#34;
 *     
 *             }
 *     
 *             persist     = true
 *     
 *         }
 *     
 *     )
 *     
 *     id            = &#34;Sample_http_01,Sample_non_http_01&#34;
 *     
 *     tenant_filter = &#34;Sample_http_01,Sample_non_http_01&#34;
 *     
 *     tenant_list   = &#34;Sample_http_01,Sample_non_http_01&#34;
 * 
 * }
 * 
 * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html
 * 
 */
@ResourceType(type="f5bigip:index/as3:As3")
public class As3 extends com.pulumi.resources.CustomResource {
    /**
     * List of applications currently deployed on the Big-Ip
     * 
     */
    @Export(name="applicationList", refs={String.class}, tree="[0]")
    private Output<String> applicationList;

    /**
     * @return List of applications currently deployed on the Big-Ip
     * 
     */
    public Output<String> applicationList() {
        return this.applicationList;
    }
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    @Export(name="as3Json", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> as3Json;

    /**
     * @return Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    public Output<Optional<String>> as3Json() {
        return Codegen.optional(this.as3Json);
    }
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to false
     * 
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     * 
     * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
     * 
     */
    @Export(name="ignoreMetadata", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreMetadata;

    /**
     * @return Set True if you want to ignore metadata changes during update. By default it is set to false
     * 
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     * 
     * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
     * 
     */
    public Output<Optional<Boolean>> ignoreMetadata() {
        return Codegen.optional(this.ignoreMetadata);
    }
    /**
     * Will specify whether is deployment is done via Per-Application Way or Traditional Way
     * 
     */
    @Export(name="perAppMode", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> perAppMode;

    /**
     * @return Will specify whether is deployment is done via Per-Application Way or Traditional Way
     * 
     */
    public Output<Boolean> perAppMode() {
        return this.perAppMode;
    }
    /**
     * ID of AS3 post declaration async task
     * 
     */
    @Export(name="taskId", refs={String.class}, tree="[0]")
    private Output<String> taskId;

    /**
     * @return ID of AS3 post declaration async task
     * 
     */
    public Output<String> taskId() {
        return this.taskId;
    }
    /**
     * If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     * 
     */
    @Export(name="tenantFilter", refs={String.class}, tree="[0]")
    private Output<String> tenantFilter;

    /**
     * @return If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     * 
     */
    public Output<String> tenantFilter() {
        return this.tenantFilter;
    }
    /**
     * List of tenants currently deployed on the Big-Ip
     * 
     */
    @Export(name="tenantList", refs={String.class}, tree="[0]")
    private Output<String> tenantList;

    /**
     * @return List of tenants currently deployed on the Big-Ip
     * 
     */
    public Output<String> tenantList() {
        return this.tenantList;
    }
    /**
     * Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
     * 
     */
    @Export(name="tenantName", refs={String.class}, tree="[0]")
    private Output<String> tenantName;

    /**
     * @return Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
     * 
     */
    public Output<String> tenantName() {
        return this.tenantName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public As3(String name) {
        this(name, As3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public As3(String name, @Nullable As3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public As3(String name, @Nullable As3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/as3:As3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private As3(String name, Output<String> id, @Nullable As3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/as3:As3", name, state, makeResourceOptions(options, id));
    }

    private static As3Args makeArgs(@Nullable As3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? As3Args.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static As3 get(String name, Output<String> id, @Nullable As3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new As3(name, id, state, options);
    }
}
