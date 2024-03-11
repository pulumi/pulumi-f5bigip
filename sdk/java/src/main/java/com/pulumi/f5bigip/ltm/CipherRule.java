// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.CipherRuleArgs;
import com.pulumi.f5bigip.ltm.inputs.CipherRuleState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.CipherRule` Manages F5 BIG-IP LTM cipher rule using iControl REST.
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
 * import com.pulumi.f5bigip.ltm.CipherRule;
 * import com.pulumi.f5bigip.ltm.CipherRuleArgs;
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
 *         var testCipherRule = new CipherRule(&#34;testCipherRule&#34;, CipherRuleArgs.builder()        
 *             .cipher(&#34;TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384&#34;)
 *             .dhGroups(&#34;P256:P384:FFDHE2048:FFDHE3072:FFDHE4096&#34;)
 *             .name(&#34;/Common/test_cipher_rule&#34;)
 *             .signatureAlgorithms(&#34;DEFAULT&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Importing
 * 
 * An existing cipher rule can be imported into this resource by supplying the cipher rule full path name  ex : `/partition/name`
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_cipher_rule.test_cipher_rule /Common/test_cipher_rule
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/cipherRule:CipherRule")
public class CipherRule extends com.pulumi.resources.CustomResource {
    /**
     * Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
     * 
     */
    @Export(name="cipher", refs={String.class}, tree="[0]")
    private Output<String> cipher;

    /**
     * @return Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
     * 
     */
    public Output<String> cipher() {
        return this.cipher;
    }
    /**
     * The Partition in which the Cipher Rule will be created.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The Partition in which the Cipher Rule will be created.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies the DH Groups algorithms, separated by colons (:).
     * 
     */
    @Export(name="dhGroups", refs={String.class}, tree="[0]")
    private Output<String> dhGroups;

    /**
     * @return Specifies the DH Groups algorithms, separated by colons (:).
     * 
     */
    public Output<String> dhGroups() {
        return this.dhGroups;
    }
    /**
     * Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the Signature Algorithms, separated by colons (:).
     * 
     */
    @Export(name="signatureAlgorithms", refs={String.class}, tree="[0]")
    private Output<String> signatureAlgorithms;

    /**
     * @return Specifies the Signature Algorithms, separated by colons (:).
     * 
     */
    public Output<String> signatureAlgorithms() {
        return this.signatureAlgorithms;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CipherRule(String name) {
        this(name, CipherRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CipherRule(String name, CipherRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CipherRule(String name, CipherRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/cipherRule:CipherRule", name, args == null ? CipherRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CipherRule(String name, Output<String> id, @Nullable CipherRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/cipherRule:CipherRule", name, state, makeResourceOptions(options, id));
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
    public static CipherRule get(String name, Output<String> id, @Nullable CipherRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CipherRule(name, id, state, options);
    }
}
