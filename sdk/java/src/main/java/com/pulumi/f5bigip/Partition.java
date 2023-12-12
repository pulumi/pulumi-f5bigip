// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.PartitionArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.PartitionState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.Partition` Manages F5 BIG-IP partitions
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.Partition;
 * import com.pulumi.f5bigip.PartitionArgs;
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
 *         var test_partition = new Partition(&#34;test-partition&#34;, PartitionArgs.builder()        
 *             .description(&#34;created by terraform&#34;)
 *             .name(&#34;test-partition&#34;)
 *             .routeDomainId(2)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:index/partition:Partition")
public class Partition extends com.pulumi.resources.CustomResource {
    /**
     * Description of the partition.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the partition.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the partition.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the partition.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Route domain id of the partition.
     * 
     */
    @Export(name="routeDomainId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> routeDomainId;

    /**
     * @return Route domain id of the partition.
     * 
     */
    public Output<Optional<Integer>> routeDomainId() {
        return Codegen.optional(this.routeDomainId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Partition(String name) {
        this(name, PartitionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Partition(String name, PartitionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Partition(String name, PartitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/partition:Partition", name, args == null ? PartitionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Partition(String name, Output<String> id, @Nullable PartitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/partition:Partition", name, state, makeResourceOptions(options, id));
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
    public static Partition get(String name, Output<String> id, @Nullable PartitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Partition(name, id, state, options);
    }
}
