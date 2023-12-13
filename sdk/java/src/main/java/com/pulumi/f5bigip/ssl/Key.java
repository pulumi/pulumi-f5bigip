// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ssl.KeyArgs;
import com.pulumi.f5bigip.ssl.inputs.KeyState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ssl.Key` This resource will import SSL certificate key on BIG-IP LTM.
 * Certificate key can be imported from certificate key files on the local disk, in PEM format
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ssl.Key;
 * import com.pulumi.f5bigip.ssl.KeyArgs;
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
 *         var test_key = new Key(&#34;test-key&#34;, KeyArgs.builder()        
 *             .name(&#34;serverkey.key&#34;)
 *             .content(Files.readString(Paths.get(&#34;serverkey.key&#34;)))
 *             .partition(&#34;Common&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:ssl/key:Key")
public class Key extends com.pulumi.resources.CustomResource {
    /**
     * Content of SSL certificate key present on local Disk
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return Content of SSL certificate key present on local Disk
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * Full Path Name of ssl key
     * 
     */
    @Export(name="fullPath", refs={String.class}, tree="[0]")
    private Output<String> fullPath;

    /**
     * @return Full Path Name of ssl key
     * 
     */
    public Output<String> fullPath() {
        return this.fullPath;
    }
    /**
     * Name of the SSL Certificate key to be Imported on to BIGIP
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the SSL Certificate key to be Imported on to BIGIP
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Partition of ssl certificate key
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Partition of ssl certificate key
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * Passphrase on key.
     * 
     */
    @Export(name="passphrase", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> passphrase;

    /**
     * @return Passphrase on key.
     * 
     */
    public Output<Optional<String>> passphrase() {
        return Codegen.optional(this.passphrase);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Key(String name) {
        this(name, KeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Key(String name, KeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Key(String name, KeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ssl/key:Key", name, args == null ? KeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Key(String name, Output<String> id, @Nullable KeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ssl/key:Key", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "content",
                "passphrase"
            ))
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
    public static Key get(String name, Output<String> id, @Nullable KeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Key(name, id, state, options);
    }
}
