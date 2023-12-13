// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PersistenceProfileCookieArgs;
import com.pulumi.f5bigip.ltm.inputs.PersistenceProfileCookieState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Configures a cookie persistence profile
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileCookie;
 * import com.pulumi.f5bigip.ltm.PersistenceProfileCookieArgs;
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
 *         var testPpcookie = new PersistenceProfileCookie(&#34;testPpcookie&#34;, PersistenceProfileCookieArgs.builder()        
 *             .name(&#34;/Common/terraform_cookie&#34;)
 *             .defaultsFrom(&#34;/Common/cookie&#34;)
 *             .matchAcrossPools(&#34;enabled&#34;)
 *             .matchAcrossServices(&#34;enabled&#34;)
 *             .matchAcrossVirtuals(&#34;enabled&#34;)
 *             .timeout(3600)
 *             .overrideConnLimit(&#34;enabled&#34;)
 *             .alwaysSend(&#34;enabled&#34;)
 *             .cookieEncryption(&#34;required&#34;)
 *             .cookieEncryptionPassphrase(&#34;iam&#34;)
 *             .cookieName(&#34;ham&#34;)
 *             .expiration(&#34;1:0:0&#34;)
 *             .hashLength(0)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Reference
 * 
 * `name` - (Required) Name of the virtual address
 * 
 * `defaults_from` - (Required) Parent cookie persistence profile
 * 
 * `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record
 * 
 * `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record
 * 
 * `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
 * 
 * `method` (Optional) Specifies the type of cookie processing that the system uses. The default value is insert
 * 
 * `mirror` (Optional) (enabled or disabled) mirror persistence record
 * 
 * `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
 * 
 * `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
 * 
 * `always_send` (Optional) (enabled or disabled) always send cookies
 * 
 * `cookie_encryption` (Optional) (required, preferred, or disabled) To required, preferred, or disabled policy for cookie encryption
 * 
 * `cookie_encryption_passphrase` (Optional) (required, preferred, or disabled) Passphrase for encrypted cookies. The field is encrypted on the server and will always return differently then set.
 * If this is configured specify `ignore_changes` under the `lifecycle` block to ignore returned encrypted value.
 * 
 * `cookie_name` (Optional) Name of the cookie to track persistence
 * 
 * `expiration` (Optional) Expiration TTL for cookie specified in DAY:HOUR:MIN:SECONDS (Examples: 1:0:0:0 one day, 1:0:0 one hour, 30:0 thirty minutes)
 * 
 * `hash_length` (Optional) (Integer) Length of hash to apply to cookie
 * 
 * `hash_offset` (Optional) (Integer) Number of characters to skip in the cookie for the hash
 * 
 * `httponly` (Optional) (enabled or disabled) Sending only over http
 * 
 */
@ResourceType(type="f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie")
public class PersistenceProfileCookie extends com.pulumi.resources.CustomResource {
    /**
     * To enable _ disable always sending cookies
     * 
     */
    @Export(name="alwaysSend", refs={String.class}, tree="[0]")
    private Output<String> alwaysSend;

    /**
     * @return To enable _ disable always sending cookies
     * 
     */
    public Output<String> alwaysSend() {
        return this.alwaysSend;
    }
    @Export(name="appService", refs={String.class}, tree="[0]")
    private Output<String> appService;

    public Output<String> appService() {
        return this.appService;
    }
    /**
     * To required, preferred, or disabled policy for cookie encryption
     * 
     */
    @Export(name="cookieEncryption", refs={String.class}, tree="[0]")
    private Output<String> cookieEncryption;

    /**
     * @return To required, preferred, or disabled policy for cookie encryption
     * 
     */
    public Output<String> cookieEncryption() {
        return this.cookieEncryption;
    }
    /**
     * Passphrase for encrypted cookies
     * 
     */
    @Export(name="cookieEncryptionPassphrase", refs={String.class}, tree="[0]")
    private Output<String> cookieEncryptionPassphrase;

    /**
     * @return Passphrase for encrypted cookies
     * 
     */
    public Output<String> cookieEncryptionPassphrase() {
        return this.cookieEncryptionPassphrase;
    }
    /**
     * Name of the cookie to track persistence
     * 
     */
    @Export(name="cookieName", refs={String.class}, tree="[0]")
    private Output<String> cookieName;

    /**
     * @return Name of the cookie to track persistence
     * 
     */
    public Output<String> cookieName() {
        return this.cookieName;
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
     * Expiration TTL for cookie specified in D:H:M:S or in seconds
     * 
     */
    @Export(name="expiration", refs={String.class}, tree="[0]")
    private Output<String> expiration;

    /**
     * @return Expiration TTL for cookie specified in D:H:M:S or in seconds
     * 
     */
    public Output<String> expiration() {
        return this.expiration;
    }
    /**
     * Length of hash to apply to cookie
     * 
     */
    @Export(name="hashLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> hashLength;

    /**
     * @return Length of hash to apply to cookie
     * 
     */
    public Output<Integer> hashLength() {
        return this.hashLength;
    }
    /**
     * Number of characters to skip in the cookie for the hash
     * 
     */
    @Export(name="hashOffset", refs={Integer.class}, tree="[0]")
    private Output<Integer> hashOffset;

    /**
     * @return Number of characters to skip in the cookie for the hash
     * 
     */
    public Output<Integer> hashOffset() {
        return this.hashOffset;
    }
    /**
     * To enable _ disable sending only over http
     * 
     */
    @Export(name="httponly", refs={String.class}, tree="[0]")
    private Output<String> httponly;

    /**
     * @return To enable _ disable sending only over http
     * 
     */
    public Output<String> httponly() {
        return this.httponly;
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
     * To enable _ disable match across virtual servers with given persistence record
     * 
     */
    @Export(name="matchAcrossVirtuals", refs={String.class}, tree="[0]")
    private Output<String> matchAcrossVirtuals;

    /**
     * @return To enable _ disable match across virtual servers with given persistence record
     * 
     */
    public Output<String> matchAcrossVirtuals() {
        return this.matchAcrossVirtuals;
    }
    /**
     * Specifies the type of cookie processing that the system uses
     * 
     */
    @Export(name="method", refs={String.class}, tree="[0]")
    private Output<String> method;

    /**
     * @return Specifies the type of cookie processing that the system uses
     * 
     */
    public Output<String> method() {
        return this.method;
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
    public PersistenceProfileCookie(String name) {
        this(name, PersistenceProfileCookieArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PersistenceProfileCookie(String name, PersistenceProfileCookieArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PersistenceProfileCookie(String name, PersistenceProfileCookieArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie", name, args == null ? PersistenceProfileCookieArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PersistenceProfileCookie(String name, Output<String> id, @Nullable PersistenceProfileCookieState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie", name, state, makeResourceOptions(options, id));
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
    public static PersistenceProfileCookie get(String name, Output<String> id, @Nullable PersistenceProfileCookieState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PersistenceProfileCookie(name, id, state, options);
    }
}
