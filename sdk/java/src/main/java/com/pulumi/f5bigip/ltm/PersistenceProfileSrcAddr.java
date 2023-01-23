// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PersistenceProfileSrcAddrArgs;
import com.pulumi.f5bigip.ltm.inputs.PersistenceProfileSrcAddrState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Configures a source address persistence profile
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileSrcAddr;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileSrcAddrArgs;
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
 *         var srcaddr = new PersistenceProfileSrcAddr(&#34;srcaddr&#34;, PersistenceProfileSrcAddrArgs.builder()        
 *             .defaultsFrom(&#34;/Common/source_addr&#34;)
 *             .hashAlgorithm(&#34;carp&#34;)
 *             .mapProxies(&#34;enabled&#34;)
 *             .mask(&#34;255.255.255.255&#34;)
 *             .matchAcrossPools(&#34;enabled&#34;)
 *             .matchAcrossServices(&#34;enabled&#34;)
 *             .matchAcrossVirtuals(&#34;enabled&#34;)
 *             .mirror(&#34;enabled&#34;)
 *             .name(&#34;/Common/terraform_srcaddr&#34;)
 *             .overrideConnLimit(&#34;enabled&#34;)
 *             .timeout(3600)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Reference
 * 
 * `name` - (Required) Name of the virtual address
 * 
 * `defaults_from` - (Required) Parent cookie persistence profile
 * 
 * `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record
 * 
 * `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record
 * 
 * `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
 * 
 * `mirror` (Optional) (enabled or disabled) mirror persistence record
 * 
 * `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
 * 
 * `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
 * 
 * `hash_algorithm` (Optional) Specify the hash algorithm
 * 
 * `mask` (Optional) Identify a range of source IP addresses to manage together as a single source address affinity persistent connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
 * 
 * `map_proxies` (Optional) (enabled or disabled) Directs all to the same single pool member
 * 
 */
@ResourceType(type="f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr")
public class PersistenceProfileSrcAddr extends com.pulumi.resources.CustomResource {
    @Export(name="appService", type=String.class, parameters={})
    private Output<String> appService;

    public Output<String> appService() {
        return this.appService;
    }
    /**
     * Inherit defaults from parent profile
     * 
     */
    @Export(name="defaultsFrom", type=String.class, parameters={})
    private Output<String> defaultsFrom;

    /**
     * @return Inherit defaults from parent profile
     * 
     */
    public Output<String> defaultsFrom() {
        return this.defaultsFrom;
    }
    /**
     * Specify the hash algorithm
     * 
     */
    @Export(name="hashAlgorithm", type=String.class, parameters={})
    private Output<String> hashAlgorithm;

    /**
     * @return Specify the hash algorithm
     * 
     */
    public Output<String> hashAlgorithm() {
        return this.hashAlgorithm;
    }
    /**
     * To enable _ disable directs all to the same single pool member
     * 
     */
    @Export(name="mapProxies", type=String.class, parameters={})
    private Output<String> mapProxies;

    /**
     * @return To enable _ disable directs all to the same single pool member
     * 
     */
    public Output<String> mapProxies() {
        return this.mapProxies;
    }
    /**
     * Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     * 
     */
    @Export(name="mask", type=String.class, parameters={})
    private Output<String> mask;

    /**
     * @return Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     * 
     */
    public Output<String> mask() {
        return this.mask;
    }
    /**
     * To enable _ disable match across pools with given persistence record
     * 
     */
    @Export(name="matchAcrossPools", type=String.class, parameters={})
    private Output<String> matchAcrossPools;

    /**
     * @return To enable _ disable match across pools with given persistence record
     * 
     */
    public Output<String> matchAcrossPools() {
        return this.matchAcrossPools;
    }
    /**
     * To enable _ disable match across services with given persistence record
     * 
     */
    @Export(name="matchAcrossServices", type=String.class, parameters={})
    private Output<String> matchAcrossServices;

    /**
     * @return To enable _ disable match across services with given persistence record
     * 
     */
    public Output<String> matchAcrossServices() {
        return this.matchAcrossServices;
    }
    /**
     * To enable _ disable match across services with given persistence record
     * 
     */
    @Export(name="matchAcrossVirtuals", type=String.class, parameters={})
    private Output<String> matchAcrossVirtuals;

    /**
     * @return To enable _ disable match across services with given persistence record
     * 
     */
    public Output<String> matchAcrossVirtuals() {
        return this.matchAcrossVirtuals;
    }
    /**
     * To enable _ disable
     * 
     */
    @Export(name="mirror", type=String.class, parameters={})
    private Output<String> mirror;

    /**
     * @return To enable _ disable
     * 
     */
    public Output<String> mirror() {
        return this.mirror;
    }
    /**
     * Name of the persistence profile
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the persistence profile
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     * 
     */
    @Export(name="overrideConnLimit", type=String.class, parameters={})
    private Output<String> overrideConnLimit;

    /**
     * @return To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     * 
     */
    public Output<String> overrideConnLimit() {
        return this.overrideConnLimit;
    }
    /**
     * Timeout for persistence of the session
     * 
     */
    @Export(name="timeout", type=Integer.class, parameters={})
    private Output<Integer> timeout;

    /**
     * @return Timeout for persistence of the session
     * 
     */
    public Output<Integer> timeout() {
        return this.timeout;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PersistenceProfileSrcAddr(String name) {
        this(name, PersistenceProfileSrcAddrArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PersistenceProfileSrcAddr(String name, PersistenceProfileSrcAddrArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PersistenceProfileSrcAddr(String name, PersistenceProfileSrcAddrArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, args == null ? PersistenceProfileSrcAddrArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PersistenceProfileSrcAddr(String name, Output<String> id, @Nullable PersistenceProfileSrcAddrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, state, makeResourceOptions(options, id));
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
    public static PersistenceProfileSrcAddr get(String name, Output<String> id, @Nullable PersistenceProfileSrcAddrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PersistenceProfileSrcAddr(name, id, state, options);
    }
}