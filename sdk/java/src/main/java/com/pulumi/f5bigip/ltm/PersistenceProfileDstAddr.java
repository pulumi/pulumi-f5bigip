// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PersistenceProfileDstAddrArgs;
import com.pulumi.f5bigip.ltm.inputs.PersistenceProfileDstAddrState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Configures a cookie persistence profile
 * 
 * ## Example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileDstAddr;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileDstAddrArgs;
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
 *         var dstaddr = new PersistenceProfileDstAddr("dstaddr", PersistenceProfileDstAddrArgs.builder()
 *             .name("/Common/terraform_ppdstaddr")
 *             .defaultsFrom("/Common/dest_addr")
 *             .matchAcrossPools("enabled")
 *             .matchAcrossServices("enabled")
 *             .matchAcrossVirtuals("enabled")
 *             .mirror("enabled")
 *             .timeout(3600)
 *             .overrideConnLimit("enabled")
 *             .hashAlgorithm("carp")
 *             .mask("255.255.255.255")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Reference
 * 
 * `name` - (Required) Name of the virtual address
 * 
 * `defaults_from` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.
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
 * ## Importing
 * 
 * An dest-addr persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_persistence_profile_dstaddr.dstaddr &#34;/Common/terraform_ppdstaddr&#34;
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr")
public class PersistenceProfileDstAddr extends com.pulumi.resources.CustomResource {
    @Export(name="appService", refs={String.class}, tree="[0]")
    private Output<String> appService;

    public Output<String> appService() {
        return this.appService;
    }
    /**
     * Inherit defaults from parent profile
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
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
    @Export(name="hashAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> hashAlgorithm;

    /**
     * @return Specify the hash algorithm
     * 
     */
    public Output<String> hashAlgorithm() {
        return this.hashAlgorithm;
    }
    /**
     * Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
     * when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
     * 
     */
    @Export(name="mask", refs={String.class}, tree="[0]")
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
    @Export(name="matchAcrossPools", refs={String.class}, tree="[0]")
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
    @Export(name="matchAcrossServices", refs={String.class}, tree="[0]")
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
    @Export(name="matchAcrossVirtuals", refs={String.class}, tree="[0]")
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
    @Export(name="mirror", refs={String.class}, tree="[0]")
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
    @Export(name="name", refs={String.class}, tree="[0]")
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
    @Export(name="overrideConnLimit", refs={String.class}, tree="[0]")
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
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
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
    public PersistenceProfileDstAddr(String name) {
        this(name, PersistenceProfileDstAddrArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PersistenceProfileDstAddr(String name, PersistenceProfileDstAddrArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PersistenceProfileDstAddr(String name, PersistenceProfileDstAddrArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, args == null ? PersistenceProfileDstAddrArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PersistenceProfileDstAddr(String name, Output<String> id, @Nullable PersistenceProfileDstAddrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, state, makeResourceOptions(options, id));
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
    public static PersistenceProfileDstAddr get(String name, Output<String> id, @Nullable PersistenceProfileDstAddrState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PersistenceProfileDstAddr(name, id, state, options);
    }
}
