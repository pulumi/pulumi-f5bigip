// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.SnatPoolArgs;
import com.pulumi.f5bigip.ltm.inputs.SnatPoolState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.SnatPool` Collections of SNAT translation addresses
 * 
 * Resource should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.SnatPool;
 * import com.pulumi.f5bigip.ltm.SnatPoolArgs;
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
 *         var snatpoolSanjose = new SnatPool(&#34;snatpoolSanjose&#34;, SnatPoolArgs.builder()        
 *             .name(&#34;/Common/snatpool_sanjose&#34;)
 *             .members(            
 *                 &#34;191.1.1.1&#34;,
 *                 &#34;194.2.2.2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:ltm/snatPool:SnatPool")
public class SnatPool extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * Name of the snatpool
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the snatpool
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SnatPool(String name) {
        this(name, SnatPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SnatPool(String name, SnatPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SnatPool(String name, SnatPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/snatPool:SnatPool", name, args == null ? SnatPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SnatPool(String name, Output<String> id, @Nullable SnatPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/snatPool:SnatPool", name, state, makeResourceOptions(options, id));
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
    public static SnatPool get(String name, Output<String> id, @Nullable SnatPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SnatPool(name, id, state, options);
    }
}
