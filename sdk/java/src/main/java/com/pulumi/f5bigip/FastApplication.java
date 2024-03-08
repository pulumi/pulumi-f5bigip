// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.FastApplicationArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.FastApplicationState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.FastApplication` This resource will create and manage FAST applications on BIG-IP from provided JSON declaration.
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
 * import com.pulumi.f5bigip.FastApplication;
 * import com.pulumi.f5bigip.FastApplicationArgs;
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
 *         var foo_app = new FastApplication(&#34;foo-app&#34;, FastApplicationArgs.builder()        
 *             .fastJson(Files.readString(Paths.get(&#34;new_fast_app.json&#34;)))
 *             .template(&#34;examples/simple_http&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/fastApplication:FastApplication")
public class FastApplication extends com.pulumi.resources.CustomResource {
    /**
     * A FAST application name.
     * 
     * * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
     * 
     */
    @Export(name="application", refs={String.class}, tree="[0]")
    private Output<String> application;

    /**
     * @return A FAST application name.
     * 
     * * `FAST documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/
     * 
     */
    public Output<String> application() {
        return this.application;
    }
    /**
     * Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
     * 
     */
    @Export(name="fastJson", refs={String.class}, tree="[0]")
    private Output<String> fastJson;

    /**
     * @return Path/Filename of Declarative FAST JSON which is a json file used with builtin ```file``` function
     * 
     */
    public Output<String> fastJson() {
        return this.fastJson;
    }
    /**
     * Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
     * 
     */
    @Export(name="template", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> template;

    /**
     * @return Name of installed FAST template used to create FAST application. This parameter is required when creating new resource.
     * 
     */
    public Output<Optional<String>> template() {
        return Codegen.optional(this.template);
    }
    /**
     * A FAST tenant name on which you want to manage application.
     * 
     */
    @Export(name="tenant", refs={String.class}, tree="[0]")
    private Output<String> tenant;

    /**
     * @return A FAST tenant name on which you want to manage application.
     * 
     */
    public Output<String> tenant() {
        return this.tenant;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FastApplication(String name) {
        this(name, FastApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FastApplication(String name, FastApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FastApplication(String name, FastApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/fastApplication:FastApplication", name, args == null ? FastApplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FastApplication(String name, Output<String> id, @Nullable FastApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/fastApplication:FastApplication", name, state, makeResourceOptions(options, id));
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
    public static FastApplication get(String name, Output<String> id, @Nullable FastApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FastApplication(name, id, state, options);
    }
}
