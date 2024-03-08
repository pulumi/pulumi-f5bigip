// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.CipherGroupArgs;
import com.pulumi.f5bigip.ltm.inputs.CipherGroupState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.CipherGroup` Manages F5 BIG-IP LTM cipher group using iControl REST.
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
 * import com.pulumi.f5bigip.ltm.CipherGroup;
 * import com.pulumi.f5bigip.ltm.CipherGroupArgs;
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
 *         var test_cipher_group = new CipherGroup(&#34;test-cipher-group&#34;, CipherGroupArgs.builder()        
 *             .allows(&#34;/Common/f5-aes&#34;)
 *             .name(&#34;/Common/test-cipher-group-01&#34;)
 *             .ordering(&#34;speed&#34;)
 *             .requires(&#34;/Common/f5-quic&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Importing
 * 
 * An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_cipher_group.test_cipher_group /Common/test_cipher_group
 * 
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/cipherGroup:CipherGroup")
public class CipherGroup extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     * 
     */
    @Export(name="allows", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allows;

    /**
     * @return Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     * 
     */
    public Output<Optional<List<String>>> allows() {
        return Codegen.optional(this.allows);
    }
    /**
     * Specifies descriptive text that identifies the cipher rule
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Specifies descriptive text that identifies the cipher rule
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
     * 
     */
    @Export(name="ordering", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ordering;

    /**
     * @return Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
     * 
     */
    public Output<Optional<String>> ordering() {
        return Codegen.optional(this.ordering);
    }
    /**
     * Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     * 
     */
    @Export(name="requires", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> requires;

    /**
     * @return Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     * 
     */
    public Output<Optional<List<String>>> requires() {
        return Codegen.optional(this.requires);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CipherGroup(String name) {
        this(name, CipherGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CipherGroup(String name, CipherGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CipherGroup(String name, CipherGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/cipherGroup:CipherGroup", name, args == null ? CipherGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CipherGroup(String name, Output<String> id, @Nullable CipherGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/cipherGroup:CipherGroup", name, state, makeResourceOptions(options, id));
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
    public static CipherGroup get(String name, Output<String> id, @Nullable CipherGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CipherGroup(name, id, state, options);
    }
}
