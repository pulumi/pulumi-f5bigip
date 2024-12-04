// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PoolArgs;
import com.pulumi.f5bigip.ltm.inputs.PoolState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.Pool` Manages F5 BIG-IP LTM pools via iControl REST API.
 * 
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource or  `partition + directory + name`.
 * For example `/Common/my-pool`.
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
 * import com.pulumi.f5bigip.ltm.Monitor;
 * import com.pulumi.f5bigip.ltm.MonitorArgs;
 * import com.pulumi.f5bigip.ltm.Pool;
 * import com.pulumi.f5bigip.ltm.PoolArgs;
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
 *         var monitor = new Monitor("monitor", MonitorArgs.builder()
 *             .name("/Common/terraform_monitor")
 *             .parent("/Common/http")
 *             .build());
 * 
 *         var pool = new Pool("pool", PoolArgs.builder()
 *             .name("/Common/Axiom_Environment_APP1_Pool")
 *             .loadBalancingMode("round-robin")
 *             .minimumActiveMembers(1)
 *             .monitors(monitor.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;      
 * 
 * ## Importing
 * 
 * An existing pool can be imported into this resource by supplying pool Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_pool.k8s_prod_import /Common/k8prod_Pool
 * 
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/pool:Pool")
public class Pool extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
     * 
     */
    @Export(name="allowNat", refs={String.class}, tree="[0]")
    private Output<String> allowNat;

    /**
     * @return Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
     * 
     */
    public Output<String> allowNat() {
        return this.allowNat;
    }
    /**
     * Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
     * 
     */
    @Export(name="allowSnat", refs={String.class}, tree="[0]")
    private Output<String> allowSnat;

    /**
     * @return Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
     * 
     */
    public Output<String> allowSnat() {
        return this.allowSnat;
    }
    /**
     * Specifies descriptive text that identifies the pool.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Specifies descriptive text that identifies the pool.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies the load balancing method. The default is `round-robin`. Possible options: [`dynamic-ratio-member`,`dynamic-ratio-node`, `fastest-app-response`,`fastest-node`, `least-connections-members`,`least-connections-node`,`least-sessions`,`observed-member`,`observed-node`,`predictive-member`,`predictive-node`,`ratio-least-connections-member`,`ratio-least-connections-node`,`ratio-member`,`ratio-node`,`ratio-session`,`round-robin`,`weighted-least-connections-member`,`weighted-least-connections-node`]
     * 
     */
    @Export(name="loadBalancingMode", refs={String.class}, tree="[0]")
    private Output<String> loadBalancingMode;

    /**
     * @return Specifies the load balancing method. The default is `round-robin`. Possible options: [`dynamic-ratio-member`,`dynamic-ratio-node`, `fastest-app-response`,`fastest-node`, `least-connections-members`,`least-connections-node`,`least-sessions`,`observed-member`,`observed-node`,`predictive-member`,`predictive-node`,`ratio-least-connections-member`,`ratio-least-connections-node`,`ratio-member`,`ratio-node`,`ratio-session`,`round-robin`,`weighted-least-connections-member`,`weighted-least-connections-node`]
     * 
     */
    public Output<String> loadBalancingMode() {
        return this.loadBalancingMode;
    }
    /**
     * Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
     * 
     */
    @Export(name="minimumActiveMembers", refs={Integer.class}, tree="[0]")
    private Output<Integer> minimumActiveMembers;

    /**
     * @return Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
     * 
     */
    public Output<Integer> minimumActiveMembers() {
        return this.minimumActiveMembers;
    }
    /**
     * List of monitor names to associate with the pool
     * 
     */
    @Export(name="monitors", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> monitors;

    /**
     * @return List of monitor names to associate with the pool
     * 
     */
    public Output<List<String>> monitors() {
        return this.monitors;
    }
    /**
     * Name of the pool,it should be `full path`.The full path is the combination of the `partition + name` of the pool.(For example `/Common/my-pool`)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the pool,it should be `full path`.The full path is the combination of the `partition + name` of the pool.(For example `/Common/my-pool`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the number of times the system tries to contact a new pool member after a passive failure.
     * 
     */
    @Export(name="reselectTries", refs={Integer.class}, tree="[0]")
    private Output<Integer> reselectTries;

    /**
     * @return Specifies the number of times the system tries to contact a new pool member after a passive failure.
     * 
     */
    public Output<Integer> reselectTries() {
        return this.reselectTries;
    }
    /**
     * Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
     * 
     */
    @Export(name="serviceDownAction", refs={String.class}, tree="[0]")
    private Output<String> serviceDownAction;

    /**
     * @return Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
     * 
     */
    public Output<String> serviceDownAction() {
        return this.serviceDownAction;
    }
    /**
     * Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
     * 
     */
    @Export(name="slowRampTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> slowRampTime;

    /**
     * @return Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
     * 
     */
    public Output<Integer> slowRampTime() {
        return this.slowRampTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Pool(java.lang.String name) {
        this(name, PoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Pool(java.lang.String name, PoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Pool(java.lang.String name, PoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/pool:Pool", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Pool(java.lang.String name, Output<java.lang.String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/pool:Pool", name, state, makeResourceOptions(options, id), false);
    }

    private static PoolArgs makeArgs(PoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PoolArgs.Empty : args;
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
    public static Pool get(java.lang.String name, Output<java.lang.String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Pool(name, id, state, options);
    }
}
