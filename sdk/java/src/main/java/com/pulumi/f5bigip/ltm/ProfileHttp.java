// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileHttpArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileHttpState;
import com.pulumi.f5bigip.ltm.outputs.ProfileHttpEnforcement;
import com.pulumi.f5bigip.ltm.outputs.ProfileHttpHttpStrictTransportSecurity;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileHttp` Configures a custom profile_http for use by health checks.
 * 
 * For resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
 * import com.pulumi.f5bigip.ltm.ProfileHttp;
 * import com.pulumi.f5bigip.ltm.ProfileHttpArgs;
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
 *         var sanjose_http = new ProfileHttp(&#34;sanjose-http&#34;, ProfileHttpArgs.builder()        
 *             .name(&#34;/Common/sanjose-http&#34;)
 *             .defaultsFrom(&#34;/Common/http&#34;)
 *             .fallbackHost(&#34;titanic&#34;)
 *             .fallbackStatusCodes(            
 *                 &#34;400&#34;,
 *                 &#34;500&#34;,
 *                 &#34;300&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * BIG-IP LTM http profiles can be imported using the `name`, e.g.
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import f5bigip:ltm/profileHttp:ProfileHttp test-http /Common/test-http
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileHttp:ProfileHttp")
public class ProfileHttp extends com.pulumi.resources.CustomResource {
    /**
     * Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request&#39;s XFF (X-forwarded-for) headers, if they exist.
     * 
     */
    @Export(name="acceptXff", refs={String.class}, tree="[0]")
    private Output<String> acceptXff;

    /**
     * @return Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request&#39;s XFF (X-forwarded-for) headers, if they exist.
     * 
     */
    public Output<String> acceptXff() {
        return this.acceptXff;
    }
    /**
     * The application service to which the object belongs.
     * 
     */
    @Export(name="appService", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> appService;

    /**
     * @return The application service to which the object belongs.
     * 
     */
    public Output<Optional<String>> appService() {
        return Codegen.optional(this.appService);
    }
    /**
     * Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is `none`
     * 
     */
    @Export(name="basicAuthRealm", refs={String.class}, tree="[0]")
    private Output<String> basicAuthRealm;

    /**
     * @return Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is `none`
     * 
     */
    public Output<String> basicAuthRealm() {
        return this.basicAuthRealm;
    }
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
    private Output<String> defaultsFrom;

    /**
     * @return Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    public Output<String> defaultsFrom() {
        return this.defaultsFrom;
    }
    /**
     * Specifies user-defined description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Specifies user-defined description.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Type a passphrase for cookie encryption.
     * 
     */
    @Export(name="encryptCookieSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptCookieSecret;

    /**
     * @return Type a passphrase for cookie encryption.
     * 
     */
    public Output<Optional<String>> encryptCookieSecret() {
        return Codegen.optional(this.encryptCookieSecret);
    }
    /**
     * Type the cookie names for the system to encrypt.
     * 
     */
    @Export(name="encryptCookies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> encryptCookies;

    /**
     * @return Type the cookie names for the system to encrypt.
     * 
     */
    public Output<Optional<List<String>>> encryptCookies() {
        return Codegen.optional(this.encryptCookies);
    }
    /**
     * See Enforcement below for more details.
     * 
     */
    @Export(name="enforcements", refs={List.class,ProfileHttpEnforcement.class}, tree="[0,1]")
    private Output<List<ProfileHttpEnforcement>> enforcements;

    /**
     * @return See Enforcement below for more details.
     * 
     */
    public Output<List<ProfileHttpEnforcement>> enforcements() {
        return this.enforcements;
    }
    /**
     * Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
     * 
     */
    @Export(name="fallbackHost", refs={String.class}, tree="[0]")
    private Output<String> fallbackHost;

    /**
     * @return Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
     * 
     */
    public Output<String> fallbackHost() {
        return this.fallbackHost;
    }
    /**
     * Specifies one or more three-digit status codes that can be returned by an HTTP server,that should trigger a redirection to the fallback host.
     * 
     */
    @Export(name="fallbackStatusCodes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> fallbackStatusCodes;

    /**
     * @return Specifies one or more three-digit status codes that can be returned by an HTTP server,that should trigger a redirection to the fallback host.
     * 
     */
    public Output<Optional<List<String>>> fallbackStatusCodes() {
        return Codegen.optional(this.fallbackStatusCodes);
    }
    /**
     * Specifies the header string that you want to erase from an HTTP request. Default is `none`.
     * 
     */
    @Export(name="headErase", refs={String.class}, tree="[0]")
    private Output<String> headErase;

    /**
     * @return Specifies the header string that you want to erase from an HTTP request. Default is `none`.
     * 
     */
    public Output<String> headErase() {
        return this.headErase;
    }
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request.Default is `none`.
     * 
     */
    @Export(name="headInsert", refs={String.class}, tree="[0]")
    private Output<String> headInsert;

    /**
     * @return Specifies a quoted header string that you want to insert into an HTTP request.Default is `none`.
     * 
     */
    public Output<String> headInsert() {
        return this.headInsert;
    }
    /**
     * See Http_Strict_Transport_Security below for more details.
     * 
     */
    @Export(name="httpStrictTransportSecurities", refs={List.class,ProfileHttpHttpStrictTransportSecurity.class}, tree="[0,1]")
    private Output<List<ProfileHttpHttpStrictTransportSecurity>> httpStrictTransportSecurities;

    /**
     * @return See Http_Strict_Transport_Security below for more details.
     * 
     */
    public Output<List<ProfileHttpHttpStrictTransportSecurity>> httpStrictTransportSecurities() {
        return this.httpStrictTransportSecurities;
    }
    /**
     * Specifies, when enabled, that the system inserts an X-Forwarded-For header in an HTTP request with the client IP address, to use with connection pooling. The default is `Disabled`.
     * 
     */
    @Export(name="insertXforwardedFor", refs={String.class}, tree="[0]")
    private Output<String> insertXforwardedFor;

    /**
     * @return Specifies, when enabled, that the system inserts an X-Forwarded-For header in an HTTP request with the client IP address, to use with connection pooling. The default is `Disabled`.
     * 
     */
    public Output<String> insertXforwardedFor() {
        return this.insertXforwardedFor;
    }
    /**
     * Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you
     * specify in the LWS Maximum Columns setting.
     * 
     */
    @Export(name="lwsSeparator", refs={String.class}, tree="[0]")
    private Output<String> lwsSeparator;

    /**
     * @return Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you
     * specify in the LWS Maximum Columns setting.
     * 
     */
    public Output<String> lwsSeparator() {
        return this.lwsSeparator;
    }
    /**
     * Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you specify in the LWS Maximum Columns setting.
     * 
     */
    @Export(name="lwsWidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> lwsWidth;

    /**
     * @return Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you specify in the LWS Maximum Columns setting.
     * 
     */
    public Output<Integer> lwsWidth() {
        return this.lwsWidth;
    }
    /**
     * Specifies the name of the http profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/test-http-profile`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the http profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/test-http-profile`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
     * 
     */
    @Export(name="oneconnectTransformations", refs={String.class}, tree="[0]")
    private Output<String> oneconnectTransformations;

    /**
     * @return Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
     * 
     */
    public Output<String> oneconnectTransformations() {
        return this.oneconnectTransformations;
    }
    /**
     * Specifies the proxy mode for this profile: reverse, explicit, or transparent. The default is `reverse`.
     * 
     */
    @Export(name="proxyType", refs={String.class}, tree="[0]")
    private Output<String> proxyType;

    /**
     * @return Specifies the proxy mode for this profile: reverse, explicit, or transparent. The default is `reverse`.
     * 
     */
    public Output<String> proxyType() {
        return this.proxyType;
    }
    /**
     * Specifies whether the system rewrites the URIs that are part of HTTP redirect (3XX) responses. The default is `none`.
     * 
     */
    @Export(name="redirectRewrite", refs={String.class}, tree="[0]")
    private Output<String> redirectRewrite;

    /**
     * @return Specifies whether the system rewrites the URIs that are part of HTTP redirect (3XX) responses. The default is `none`.
     * 
     */
    public Output<String> redirectRewrite() {
        return this.redirectRewrite;
    }
    /**
     * Specifies how the system handles HTTP content that is chunked by a client. The default is `preserve`.
     * 
     */
    @Export(name="requestChunking", refs={String.class}, tree="[0]")
    private Output<String> requestChunking;

    /**
     * @return Specifies how the system handles HTTP content that is chunked by a client. The default is `preserve`.
     * 
     */
    public Output<String> requestChunking() {
        return this.requestChunking;
    }
    /**
     * Specifies how the system handles HTTP content that is chunked by a server. The default is `selective`.
     * 
     */
    @Export(name="responseChunking", refs={String.class}, tree="[0]")
    private Output<String> responseChunking;

    /**
     * @return Specifies how the system handles HTTP content that is chunked by a server. The default is `selective`.
     * 
     */
    public Output<String> responseChunking() {
        return this.responseChunking;
    }
    /**
     * Specifies headers that the BIG-IP system allows in an HTTP response.If you are specifying more than one header, separate the headers with a blank space.
     * 
     */
    @Export(name="responseHeadersPermitteds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> responseHeadersPermitteds;

    /**
     * @return Specifies headers that the BIG-IP system allows in an HTTP response.If you are specifying more than one header, separate the headers with a blank space.
     * 
     */
    public Output<List<String>> responseHeadersPermitteds() {
        return this.responseHeadersPermitteds;
    }
    /**
     * Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. In order to remove it, &#34;none&#34; string is to be passed. If server_agent_name is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;BigIP&#34; explicitly.
     * 
     */
    @Export(name="serverAgentName", refs={String.class}, tree="[0]")
    private Output<String> serverAgentName;

    /**
     * @return Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. In order to remove it, &#34;none&#34; string is to be passed. If server_agent_name is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;BigIP&#34; explicitly.
     * 
     */
    public Output<String> serverAgentName() {
        return this.serverAgentName;
    }
    /**
     * Displays the administrative partition within which this profile resides.
     * 
     */
    @Export(name="tmPartition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tmPartition;

    /**
     * @return Displays the administrative partition within which this profile resides.
     * 
     */
    public Output<Optional<String>> tmPartition() {
        return Codegen.optional(this.tmPartition);
    }
    /**
     * Specifies the hostname to include into Via header
     * 
     */
    @Export(name="viaHostName", refs={String.class}, tree="[0]")
    private Output<String> viaHostName;

    /**
     * @return Specifies the hostname to include into Via header
     * 
     */
    public Output<String> viaHostName() {
        return this.viaHostName;
    }
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     * 
     */
    @Export(name="viaRequest", refs={String.class}, tree="[0]")
    private Output<String> viaRequest;

    /**
     * @return Specifies whether to append, remove, or preserve a Via header in an HTTP request
     * 
     */
    public Output<String> viaRequest() {
        return this.viaRequest;
    }
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     * 
     */
    @Export(name="viaResponse", refs={String.class}, tree="[0]")
    private Output<String> viaResponse;

    /**
     * @return Specifies whether to append, remove, or preserve a Via header in an HTTP request
     * 
     */
    public Output<String> viaResponse() {
        return this.viaResponse;
    }
    /**
     * Specifies alternative XFF headers instead of the default X-forwarded-for header.
     * 
     */
    @Export(name="xffAlternativeNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> xffAlternativeNames;

    /**
     * @return Specifies alternative XFF headers instead of the default X-forwarded-for header.
     * 
     */
    public Output<List<String>> xffAlternativeNames() {
        return this.xffAlternativeNames;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileHttp(String name) {
        this(name, ProfileHttpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileHttp(String name, ProfileHttpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileHttp(String name, ProfileHttpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileHttp:ProfileHttp", name, args == null ? ProfileHttpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileHttp(String name, Output<String> id, @Nullable ProfileHttpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileHttp:ProfileHttp", name, state, makeResourceOptions(options, id));
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
    public static ProfileHttp get(String name, Output<String> id, @Nullable ProfileHttpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileHttp(name, id, state, options);
    }
}
