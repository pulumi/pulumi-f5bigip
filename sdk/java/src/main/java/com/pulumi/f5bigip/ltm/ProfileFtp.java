// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileFtpArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileFtpState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileFtp` Configures a custom profile_ftp.
 * 
 * Resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
 * 
 * ## Example Usage
 * 
 * ### For Bigip versions (14.x - 16.x)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.ProfileFtp;
 * import com.pulumi.f5bigip.ltm.ProfileFtpArgs;
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
 *         var sanjose_ftp_profile = new ProfileFtp(&#34;sanjose-ftp-profile&#34;, ProfileFtpArgs.builder()        
 *             .name(&#34;/Common/sanjose-ftp-profile&#34;)
 *             .defaultsFrom(&#34;/Common/ftp&#34;)
 *             .port(2020)
 *             .description(&#34;test-tftp-profile&#34;)
 *             .ftpsMode(&#34;allow&#34;)
 *             .enforceTlssessionReuse(&#34;enabled&#34;)
 *             .allowActiveMode(&#34;enabled&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;      
 * 
 * ### For Bigip versions (12.x - 13.x)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.ProfileFtp;
 * import com.pulumi.f5bigip.ltm.ProfileFtpArgs;
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
 *         var sanjose_ftp_profile = new ProfileFtp(&#34;sanjose-ftp-profile&#34;, ProfileFtpArgs.builder()        
 *             .name(&#34;/Common/sanjose-ftp-profile&#34;)
 *             .defaultsFrom(&#34;/Common/ftp&#34;)
 *             .port(2020)
 *             .description(&#34;test-tftp-profile&#34;)
 *             .allowFtps(&#34;enabled&#34;)
 *             .translateExtended(&#34;enabled&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Common Arguments for all versions
 * 
 * * `security` - (Optional)Specifies, when checked (enabled), that the system inspects FTP traffic for security vulnerabilities using an FTP security profile. This option is available only on systems licensed for BIG-IP ASM.
 * 
 * * `port` - (Optional)Allows you to configure the FTP service to run on an alternate port. The default is 20.
 * 
 * * `log_profile` - (Optional)Configures the ALG log profile that controls logging
 * 
 * * `log_publisher` - (Optional)Configures the log publisher that handles events logging for this profile
 * 
 * *  `inherit_parent_profile` - (Optional)Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses FastL4 only.
 * 
 * * `description` - (Optional)User defined description for FTP profile
 * 
 */
@ResourceType(type="f5bigip:ltm/profileFtp:ProfileFtp")
public class ProfileFtp extends com.pulumi.resources.CustomResource {
    /**
     * Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled.
     * 
     */
    @Export(name="allowActiveMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowActiveMode;

    /**
     * @return Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled.
     * 
     */
    public Output<Optional<String>> allowActiveMode() {
        return Codegen.optional(this.allowActiveMode);
    }
    /**
     * Allows explicit FTPS negotiation
     * 
     */
    @Export(name="allowFtps", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowFtps;

    /**
     * @return Allows explicit FTPS negotiation
     * 
     */
    public Output<Optional<String>> allowFtps() {
        return Codegen.optional(this.allowFtps);
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
     * User defined description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User defined description
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default
     * value is unchecked (disabled).
     * 
     */
    @Export(name="enforceTlssessionReuse", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enforceTlssessionReuse;

    /**
     * @return Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default
     * value is unchecked (disabled).
     * 
     */
    public Output<Optional<String>> enforceTlssessionReuse() {
        return Codegen.optional(this.enforceTlssessionReuse);
    }
    /**
     * Allows explicit FTPS negotiation
     * 
     */
    @Export(name="ftpsMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ftpsMode;

    /**
     * @return Allows explicit FTPS negotiation
     * 
     */
    public Output<Optional<String>> ftpsMode() {
        return Codegen.optional(this.ftpsMode);
    }
    /**
     * Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     * 
     */
    @Export(name="inheritParentProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inheritParentProfile;

    /**
     * @return Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     * 
     */
    public Output<Optional<String>> inheritParentProfile() {
        return Codegen.optional(this.inheritParentProfile);
    }
    /**
     * inherent vlan list
     * 
     */
    @Export(name="inheritVlanList", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inheritVlanList;

    /**
     * @return inherent vlan list
     * 
     */
    public Output<Optional<String>> inheritVlanList() {
        return Codegen.optional(this.inheritVlanList);
    }
    /**
     * Configures the ALG log profile that controls logging
     * 
     */
    @Export(name="logProfile", refs={String.class}, tree="[0]")
    private Output<String> logProfile;

    /**
     * @return Configures the ALG log profile that controls logging
     * 
     */
    public Output<String> logProfile() {
        return this.logProfile;
    }
    /**
     * Configures the log publisher that handles events logging for this profile
     * 
     */
    @Export(name="logPublisher", refs={String.class}, tree="[0]")
    private Output<String> logPublisher;

    /**
     * @return Configures the log publisher that handles events logging for this profile
     * 
     */
    public Output<String> logPublisher() {
        return this.logPublisher;
    }
    /**
     * Name of the profile_ftp
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the profile_ftp
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Displays the administrative partition within which this profile resides
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output<String> partition;

    /**
     * @return Displays the administrative partition within which this profile resides
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }
    /**
     * Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> port;

    /**
     * @return Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
    }
    /**
     * Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     * 
     */
    @Export(name="security", refs={String.class}, tree="[0]")
    private Output<String> security;

    /**
     * @return Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     * 
     */
    public Output<String> security() {
        return this.security;
    }
    /**
     * This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV
     * and PORT when communicating with IPv4 servers.
     * 
     */
    @Export(name="translateExtended", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> translateExtended;

    /**
     * @return This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV
     * and PORT when communicating with IPv4 servers.
     * 
     */
    public Output<Optional<String>> translateExtended() {
        return Codegen.optional(this.translateExtended);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileFtp(String name) {
        this(name, ProfileFtpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileFtp(String name, ProfileFtpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileFtp(String name, ProfileFtpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFtp:ProfileFtp", name, args == null ? ProfileFtpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileFtp(String name, Output<String> id, @Nullable ProfileFtpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFtp:ProfileFtp", name, state, makeResourceOptions(options, id));
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
    public static ProfileFtp get(String name, Output<String> id, @Nullable ProfileFtpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileFtp(name, id, state, options);
    }
}
