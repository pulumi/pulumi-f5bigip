// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.IRuleArgs;
import com.pulumi.f5bigip.ltm.inputs.IRuleState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.IRule` Creates iRule on BIG-IP F5 device
 * 
 * For resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
 * import com.pulumi.f5bigip.ltm.IRule;
 * import com.pulumi.f5bigip.ltm.IRuleArgs;
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
 *         // Loading from a file is the preferred method
 *         var rule = new IRule("rule", IRuleArgs.builder()        
 *             .name("/Common/terraform_irule")
 *             .irule(StdFunctions.file(FileArgs.builder()
 *                 .input("myirule.tcl")
 *                 .build()).result())
 *             .build());
 * 
 *         var rule2 = new IRule("rule2", IRuleArgs.builder()        
 *             .name("/Common/terraform_irule2")
 *             .irule("""
 * when CLIENT_ACCEPTED {
 *      log local0. "test"
 *    }
 *             """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ##myirule.tcl
 * 
 */
@ResourceType(type="f5bigip:ltm/iRule:IRule")
public class IRule extends com.pulumi.resources.CustomResource {
    /**
     * Body of the iRule
     * 
     */
    @Export(name="irule", refs={String.class}, tree="[0]")
    private Output<String> irule;

    /**
     * @return Body of the iRule
     * 
     */
    public Output<String> irule() {
        return this.irule;
    }
    /**
     * Name of the iRule
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the iRule
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IRule(String name) {
        this(name, IRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IRule(String name, IRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IRule(String name, IRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/iRule:IRule", name, args == null ? IRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IRule(String name, Output<String> id, @Nullable IRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/iRule:IRule", name, state, makeResourceOptions(options, id));
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
    public static IRule get(String name, Output<String> id, @Nullable IRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IRule(name, id, state, options);
    }
}
