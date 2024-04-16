// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileRewriteArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileRewriteState;
import com.pulumi.f5bigip.ltm.outputs.ProfileRewriteCookieRule;
import com.pulumi.f5bigip.ltm.outputs.ProfileRewriteRequest;
import com.pulumi.f5bigip.ltm.outputs.ProfileRewriteResponse;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `bigip_ltm_rewrite_profile` Configures ltm policies to manage traffic assigned to a virtual server
 * 
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-profile`.
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
 * import com.pulumi.f5bigip.ltm.ProfileRewrite;
 * import com.pulumi.f5bigip.ltm.ProfileRewriteArgs;
 * import com.pulumi.f5bigip.ltm.inputs.ProfileRewriteRequestArgs;
 * import com.pulumi.f5bigip.ltm.inputs.ProfileRewriteResponseArgs;
 * import com.pulumi.f5bigip.ltm.inputs.ProfileRewriteCookieRuleArgs;
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
 *         var test_profile = new ProfileRewrite(&#34;test-profile&#34;, ProfileRewriteArgs.builder()        
 *             .name(&#34;/Common/tf_profile&#34;)
 *             .defaultsFrom(&#34;/Common/rewrite&#34;)
 *             .bypassLists(&#34;http://notouch.com&#34;)
 *             .rewriteLists(&#34;http://some.com&#34;)
 *             .rewriteMode(&#34;portal&#34;)
 *             .cacheType(&#34;cache-img-css-js&#34;)
 *             .caFile(&#34;/Common/ca-bundle.crt&#34;)
 *             .crlFile(&#34;none&#34;)
 *             .signingCert(&#34;/Common/default.crt&#34;)
 *             .signingKey(&#34;/Common/default.key&#34;)
 *             .splitTunneling(&#34;true&#34;)
 *             .build());
 * 
 *         var test_profile2 = new ProfileRewrite(&#34;test-profile2&#34;, ProfileRewriteArgs.builder()        
 *             .name(&#34;/Common/tf_profile_translate&#34;)
 *             .defaultsFrom(&#34;/Common/rewrite&#34;)
 *             .rewriteMode(&#34;uri-translation&#34;)
 *             .requests(ProfileRewriteRequestArgs.builder()
 *                 .insertXfwdFor(&#34;enabled&#34;)
 *                 .insertXfwdHost(&#34;disabled&#34;)
 *                 .insertXfwdProtocol(&#34;enabled&#34;)
 *                 .rewriteHeaders(&#34;disabled&#34;)
 *                 .build())
 *             .responses(ProfileRewriteResponseArgs.builder()
 *                 .rewriteContent(&#34;enabled&#34;)
 *                 .rewriteHeaders(&#34;disabled&#34;)
 *                 .build())
 *             .cookieRules(            
 *                 ProfileRewriteCookieRuleArgs.builder()
 *                     .ruleName(&#34;cookie1&#34;)
 *                     .clientDomain(&#34;wrong.com&#34;)
 *                     .clientPath(&#34;/this/&#34;)
 *                     .serverDomain(&#34;wrong.com&#34;)
 *                     .serverPath(&#34;/this/&#34;)
 *                     .build(),
 *                 ProfileRewriteCookieRuleArgs.builder()
 *                     .ruleName(&#34;cookie2&#34;)
 *                     .clientDomain(&#34;incorrect.com&#34;)
 *                     .clientPath(&#34;/this/&#34;)
 *                     .serverDomain(&#34;absolute.com&#34;)
 *                     .serverPath(&#34;/this/&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:ltm/profileRewrite:ProfileRewrite")
public class ProfileRewrite extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
     * 
     */
    @Export(name="bypassLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> bypassLists;

    /**
     * @return Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
     * 
     */
    public Output<Optional<List<String>>> bypassLists() {
        return Codegen.optional(this.bypassLists);
    }
    /**
     * Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
     * 
     */
    @Export(name="caFile", refs={String.class}, tree="[0]")
    private Output<String> caFile;

    /**
     * @return Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
     * 
     */
    public Output<String> caFile() {
        return this.caFile;
    }
    /**
     * Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
     * 
     */
    @Export(name="cacheType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cacheType;

    /**
     * @return Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
     * 
     */
    public Output<Optional<String>> cacheType() {
        return Codegen.optional(this.cacheType);
    }
    /**
     * Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
     * 
     */
    @Export(name="cookieRules", refs={List.class,ProfileRewriteCookieRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ProfileRewriteCookieRule>> cookieRules;

    /**
     * @return Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
     * 
     */
    public Output<Optional<List<ProfileRewriteCookieRule>>> cookieRules() {
        return Codegen.optional(this.cookieRules);
    }
    /**
     * Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
     * 
     */
    @Export(name="crlFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> crlFile;

    /**
     * @return Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
     * 
     */
    public Output<Optional<String>> crlFile() {
        return Codegen.optional(this.crlFile);
    }
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultsFrom;

    /**
     * @return Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
     * 
     */
    public Output<Optional<String>> defaultsFrom() {
        return Codegen.optional(this.defaultsFrom);
    }
    /**
     * Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Block type. Each request is block type with following arguments.
     * 
     */
    @Export(name="requests", refs={List.class,ProfileRewriteRequest.class}, tree="[0,1]")
    private Output<List<ProfileRewriteRequest>> requests;

    /**
     * @return Block type. Each request is block type with following arguments.
     * 
     */
    public Output<List<ProfileRewriteRequest>> requests() {
        return this.requests;
    }
    /**
     * Block type. Each request is block type with following arguments.
     * 
     */
    @Export(name="responses", refs={List.class,ProfileRewriteResponse.class}, tree="[0,1]")
    private Output<List<ProfileRewriteResponse>> responses;

    /**
     * @return Block type. Each request is block type with following arguments.
     * 
     */
    public Output<List<ProfileRewriteResponse>> responses() {
        return this.responses;
    }
    /**
     * Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
     * 
     */
    @Export(name="rewriteLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> rewriteLists;

    /**
     * @return Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
     * 
     */
    public Output<Optional<List<String>>> rewriteLists() {
        return Codegen.optional(this.rewriteLists);
    }
    /**
     * Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
     * 
     */
    @Export(name="rewriteMode", refs={String.class}, tree="[0]")
    private Output<String> rewriteMode;

    /**
     * @return Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
     * 
     */
    public Output<String> rewriteMode() {
        return this.rewriteMode;
    }
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
     * 
     */
    @Export(name="signingCert", refs={String.class}, tree="[0]")
    private Output<String> signingCert;

    /**
     * @return Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
     * 
     */
    public Output<String> signingCert() {
        return this.signingCert;
    }
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
     * 
     */
    @Export(name="signingKey", refs={String.class}, tree="[0]")
    private Output<String> signingKey;

    /**
     * @return Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
     * 
     */
    public Output<String> signingKey() {
        return this.signingKey;
    }
    /**
     * Specifies a pass phrase to use for encrypting the private signing key. Since it&#39;s a sensitive entity idempotency will fail in the update call.
     * 
     */
    @Export(name="signingKeyPassword", refs={String.class}, tree="[0]")
    private Output<String> signingKeyPassword;

    /**
     * @return Specifies a pass phrase to use for encrypting the private signing key. Since it&#39;s a sensitive entity idempotency will fail in the update call.
     * 
     */
    public Output<String> signingKeyPassword() {
        return this.signingKeyPassword;
    }
    /**
     * Specifies the type of Client caching. Valid choices are: `true, false`
     * 
     */
    @Export(name="splitTunneling", refs={String.class}, tree="[0]")
    private Output<String> splitTunneling;

    /**
     * @return Specifies the type of Client caching. Valid choices are: `true, false`
     * 
     */
    public Output<String> splitTunneling() {
        return this.splitTunneling;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileRewrite(String name) {
        this(name, ProfileRewriteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileRewrite(String name, ProfileRewriteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileRewrite(String name, ProfileRewriteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileRewrite:ProfileRewrite", name, args == null ? ProfileRewriteArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileRewrite(String name, Output<String> id, @Nullable ProfileRewriteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileRewrite:ProfileRewrite", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "signingKeyPassword"
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
    public static ProfileRewrite get(String name, Output<String> id, @Nullable ProfileRewriteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileRewrite(name, id, state, options);
    }
}
