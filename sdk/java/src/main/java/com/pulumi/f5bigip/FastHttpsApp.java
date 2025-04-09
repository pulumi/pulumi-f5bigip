// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.FastHttpsAppArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.FastHttpsAppState;
import com.pulumi.f5bigip.outputs.FastHttpsAppMonitor;
import com.pulumi.f5bigip.outputs.FastHttpsAppPoolMember;
import com.pulumi.f5bigip.outputs.FastHttpsAppTlsClientProfile;
import com.pulumi.f5bigip.outputs.FastHttpsAppTlsServerProfile;
import com.pulumi.f5bigip.outputs.FastHttpsAppVirtualServer;
import com.pulumi.f5bigip.outputs.FastHttpsAppWafSecurityPolicy;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.FastHttpsApp` This resource will create and manage FAST HTTPS applications on BIG-IP
 * 
 * [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.FastHttpsApp;
 * import com.pulumi.f5bigip.FastHttpsAppArgs;
 * import com.pulumi.f5bigip.inputs.FastHttpsAppVirtualServerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var fastHttpsApp = new FastHttpsApp("fastHttpsApp", FastHttpsAppArgs.builder()
 *             .tenant("fasthttpstenant")
 *             .application("fasthttpsapp")
 *             .virtualServer(FastHttpsAppVirtualServerArgs.builder()
 *                 .ip("10.30.40.44")
 *                 .port(443)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### With Service Discovery
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.fast.FastFunctions;
 * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
 * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
 * import com.pulumi.f5bigip.FastHttpsApp;
 * import com.pulumi.f5bigip.FastHttpsAppArgs;
 * import com.pulumi.f5bigip.inputs.FastHttpsAppVirtualServerArgs;
 * import com.pulumi.f5bigip.inputs.FastHttpsAppPoolMemberArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
 *             .resourceGroup("testazurerg")
 *             .subscriptionId("testazuresid")
 *             .tagKey("testazuretag")
 *             .tagValue("testazurevalue")
 *             .build());
 * 
 *         final var TC3GetGceServiceDiscovery = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
 *             .tagKey("testgcetag")
 *             .tagValue("testgcevalue")
 *             .region("testgceregion")
 *             .build());
 * 
 *         var fastHttpsApp = new FastHttpsApp("fastHttpsApp", FastHttpsAppArgs.builder()
 *             .tenant("fasthttpstenant")
 *             .application("fasthttpsapp")
 *             .virtualServer(FastHttpsAppVirtualServerArgs.builder()
 *                 .ip("10.30.40.44")
 *                 .port(443)
 *                 .build())
 *             .poolMembers(FastHttpsAppPoolMemberArgs.builder()
 *                 .addresses(                
 *                     "10.11.40.120",
 *                     "10.11.30.121",
 *                     "10.11.30.122")
 *                 .port(80)
 *                 .build())
 *             .serviceDiscoveries(            
 *                 TC3GetGceServiceDiscovery.gceSdJson(),
 *                 TC3.azureSdJson())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/fastHttpsApp:FastHttpsApp")
public class FastHttpsApp extends com.pulumi.resources.CustomResource {
    /**
     * Name of the FAST HTTPS application.
     * 
     */
    @Export(name="application", refs={String.class}, tree="[0]")
    private Output<String> application;

    /**
     * @return Name of the FAST HTTPS application.
     * 
     */
    public Output<String> application() {
        return this.application;
    }
    /**
     * List of LTM Policies to be applied FAST HTTPS Application.
     * 
     */
    @Export(name="endpointLtmPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> endpointLtmPolicies;

    /**
     * @return List of LTM Policies to be applied FAST HTTPS Application.
     * 
     */
    public Output<Optional<List<String>>> endpointLtmPolicies() {
        return Codegen.optional(this.endpointLtmPolicies);
    }
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     * 
     */
    @Export(name="existingMonitor", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingMonitor;

    /**
     * @return Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     * 
     */
    public Output<Optional<String>> existingMonitor() {
        return Codegen.optional(this.existingMonitor);
    }
    /**
     * Name of an existing BIG-IP pool.
     * 
     */
    @Export(name="existingPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingPool;

    /**
     * @return Name of an existing BIG-IP pool.
     * 
     */
    public Output<Optional<String>> existingPool() {
        return Codegen.optional(this.existingPool);
    }
    /**
     * Name of an existing BIG-IP SNAT pool.
     * 
     */
    @Export(name="existingSnatPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingSnatPool;

    /**
     * @return Name of an existing BIG-IP SNAT pool.
     * 
     */
    public Output<Optional<String>> existingSnatPool() {
        return Codegen.optional(this.existingSnatPool);
    }
    /**
     * Name of an existing TLS client profile.
     * 
     */
    @Export(name="existingTlsClientProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingTlsClientProfile;

    /**
     * @return Name of an existing TLS client profile.
     * 
     */
    public Output<Optional<String>> existingTlsClientProfile() {
        return Codegen.optional(this.existingTlsClientProfile);
    }
    /**
     * Name of an existing TLS server profile.
     * 
     */
    @Export(name="existingTlsServerProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingTlsServerProfile;

    /**
     * @return Name of an existing TLS server profile.
     * 
     */
    public Output<Optional<String>> existingTlsServerProfile() {
        return Codegen.optional(this.existingTlsServerProfile);
    }
    /**
     * Name of an existing WAF Security policy.
     * 
     */
    @Export(name="existingWafSecurityPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> existingWafSecurityPolicy;

    /**
     * @return Name of an existing WAF Security policy.
     * 
     */
    public Output<Optional<String>> existingWafSecurityPolicy() {
        return Codegen.optional(this.existingWafSecurityPolicy);
    }
    /**
     * Type of fallback persistence record to be created for each new client connection.
     * 
     */
    @Export(name="fallbackPersistence", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fallbackPersistence;

    /**
     * @return Type of fallback persistence record to be created for each new client connection.
     * 
     */
    public Output<Optional<String>> fallbackPersistence() {
        return Codegen.optional(this.fallbackPersistence);
    }
    /**
     * Json payload for FAST HTTPS application.
     * 
     */
    @Export(name="fastHttpsJson", refs={String.class}, tree="[0]")
    private Output<String> fastHttpsJson;

    /**
     * @return Json payload for FAST HTTPS application.
     * 
     */
    public Output<String> fastHttpsJson() {
        return this.fastHttpsJson;
    }
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     * 
     */
    @Export(name="loadBalancingMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loadBalancingMode;

    /**
     * @return A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     * 
     */
    public Output<Optional<String>> loadBalancingMode() {
        return Codegen.optional(this.loadBalancingMode);
    }
    /**
     * `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    @Export(name="monitor", refs={FastHttpsAppMonitor.class}, tree="[0]")
    private Output</* @Nullable */ FastHttpsAppMonitor> monitor;

    /**
     * @return `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    public Output<Optional<FastHttpsAppMonitor>> monitor() {
        return Codegen.optional(this.monitor);
    }
    /**
     * Name of an existing BIG-IP persistence profile to be used.
     * 
     */
    @Export(name="persistenceProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> persistenceProfile;

    /**
     * @return Name of an existing BIG-IP persistence profile to be used.
     * 
     */
    public Output<Optional<String>> persistenceProfile() {
        return Codegen.optional(this.persistenceProfile);
    }
    /**
     * Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
     * 
     */
    @Export(name="persistenceType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> persistenceType;

    /**
     * @return Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
     * 
     */
    public Output<Optional<String>> persistenceType() {
        return Codegen.optional(this.persistenceType);
    }
    /**
     * `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    @Export(name="poolMembers", refs={List.class,FastHttpsAppPoolMember.class}, tree="[0,1]")
    private Output<List<FastHttpsAppPoolMember>> poolMembers;

    /**
     * @return `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    public Output<List<FastHttpsAppPoolMember>> poolMembers() {
        return this.poolMembers;
    }
    /**
     * List of security log profiles to be used for FAST application
     * 
     */
    @Export(name="securityLogProfiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityLogProfiles;

    /**
     * @return List of security log profiles to be used for FAST application
     * 
     */
    public Output<Optional<List<String>>> securityLogProfiles() {
        return Codegen.optional(this.securityLogProfiles);
    }
    /**
     * List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
     * 
     */
    @Export(name="serviceDiscoveries", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> serviceDiscoveries;

    /**
     * @return List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
     * 
     */
    public Output<Optional<List<String>>> serviceDiscoveries() {
        return Codegen.optional(this.serviceDiscoveries);
    }
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     * 
     */
    @Export(name="slowRampTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> slowRampTime;

    /**
     * @return Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     * 
     */
    public Output<Optional<Integer>> slowRampTime() {
        return Codegen.optional(this.slowRampTime);
    }
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    @Export(name="snatPoolAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> snatPoolAddresses;

    /**
     * @return List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    public Output<Optional<List<String>>> snatPoolAddresses() {
        return Codegen.optional(this.snatPoolAddresses);
    }
    /**
     * Name of the FAST HTTPS application tenant.
     * 
     */
    @Export(name="tenant", refs={String.class}, tree="[0]")
    private Output<String> tenant;

    /**
     * @return Name of the FAST HTTPS application tenant.
     * 
     */
    public Output<String> tenant() {
        return this.tenant;
    }
    /**
     * `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     * 
     * &gt; **NOTE** Profile provided by `existing_tls_client_profile` or `tls_client_profile` used for encrypt server-side connections.
     * 
     */
    @Export(name="tlsClientProfile", refs={FastHttpsAppTlsClientProfile.class}, tree="[0]")
    private Output</* @Nullable */ FastHttpsAppTlsClientProfile> tlsClientProfile;

    /**
     * @return `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     * 
     * &gt; **NOTE** Profile provided by `existing_tls_client_profile` or `tls_client_profile` used for encrypt server-side connections.
     * 
     */
    public Output<Optional<FastHttpsAppTlsClientProfile>> tlsClientProfile() {
        return Codegen.optional(this.tlsClientProfile);
    }
    /**
     * `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     * &gt; **NOTE** Profile provided by `existing_tls_server_profile` or `tls_server_profile` used for decrypt client-side connections.
     * 
     */
    @Export(name="tlsServerProfile", refs={FastHttpsAppTlsServerProfile.class}, tree="[0]")
    private Output</* @Nullable */ FastHttpsAppTlsServerProfile> tlsServerProfile;

    /**
     * @return `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     * &gt; **NOTE** Profile provided by `existing_tls_server_profile` or `tls_server_profile` used for decrypt client-side connections.
     * 
     */
    public Output<Optional<FastHttpsAppTlsServerProfile>> tlsServerProfile() {
        return Codegen.optional(this.tlsServerProfile);
    }
    /**
     * `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    @Export(name="virtualServer", refs={FastHttpsAppVirtualServer.class}, tree="[0]")
    private Output</* @Nullable */ FastHttpsAppVirtualServer> virtualServer;

    /**
     * @return `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    public Output<Optional<FastHttpsAppVirtualServer>> virtualServer() {
        return Codegen.optional(this.virtualServer);
    }
    /**
     * `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    @Export(name="wafSecurityPolicy", refs={FastHttpsAppWafSecurityPolicy.class}, tree="[0]")
    private Output</* @Nullable */ FastHttpsAppWafSecurityPolicy> wafSecurityPolicy;

    /**
     * @return `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    public Output<Optional<FastHttpsAppWafSecurityPolicy>> wafSecurityPolicy() {
        return Codegen.optional(this.wafSecurityPolicy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FastHttpsApp(java.lang.String name) {
        this(name, FastHttpsAppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FastHttpsApp(java.lang.String name, FastHttpsAppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FastHttpsApp(java.lang.String name, FastHttpsAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/fastHttpsApp:FastHttpsApp", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FastHttpsApp(java.lang.String name, Output<java.lang.String> id, @Nullable FastHttpsAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/fastHttpsApp:FastHttpsApp", name, state, makeResourceOptions(options, id), false);
    }

    private static FastHttpsAppArgs makeArgs(FastHttpsAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FastHttpsAppArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static FastHttpsApp get(java.lang.String name, Output<java.lang.String> id, @Nullable FastHttpsAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FastHttpsApp(name, id, state, options);
    }
}
