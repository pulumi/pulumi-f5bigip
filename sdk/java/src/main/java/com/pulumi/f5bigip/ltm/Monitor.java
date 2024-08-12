// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.MonitorArgs;
import com.pulumi.f5bigip.ltm.inputs.MonitorState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.Monitor` Configures a custom monitor for use by health checks.
 * 
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-monitor`.
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
 * import com.pulumi.f5bigip.ltm.Monitor;
 * import com.pulumi.f5bigip.ltm.MonitorArgs;
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
 *         var monitor = new Monitor("monitor", MonitorArgs.builder()
 *             .name("/Common/terraform_monitor")
 *             .parent("/Common/http")
 *             .send("""
 * GET /some/path
 *             """)
 *             .timeout("999")
 *             .interval("998")
 *             .destination("1.2.3.4:1234")
 *             .build());
 * 
 *         var test_https_monitor = new Monitor("test-https-monitor", MonitorArgs.builder()
 *             .name("/Common/terraform_monitor")
 *             .parent("/Common/http")
 *             .sslProfile("/Common/serverssl")
 *             .send("""
 * GET /some/path
 *             """)
 *             .interval("999")
 *             .timeout("1000")
 *             .build());
 * 
 *         var test_ftp_monitor = new Monitor("test-ftp-monitor", MonitorArgs.builder()
 *             .name("/Common/ftp-test")
 *             .parent("/Common/ftp")
 *             .interval(5)
 *             .timeUntilUp(0)
 *             .timeout(16)
 *             .destination("*:8008")
 *             .filename("somefile")
 *             .build());
 * 
 *         var test_postgresql_monitor = new Monitor("test-postgresql-monitor", MonitorArgs.builder()
 *             .name("/Common/test-postgresql-monitor")
 *             .parent("/Common/postgresql")
 *             .send("SELECT 'Test';")
 *             .receive("Test")
 *             .interval(5)
 *             .timeout(16)
 *             .username("abcd")
 *             .password("abcd1234")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;      
 * 
 * ## Importing
 * 
 * An existing monitor can be imported into this resource by supplying monitor Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_monitor.monitor /Common/terraform_monitor
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/monitor:Monitor")
public class Monitor extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
     * 
     */
    @Export(name="adaptive", refs={String.class}, tree="[0]")
    private Output<String> adaptive;

    /**
     * @return Specifies whether adaptive response time monitoring is enabled for this monitor. The default is `disabled`.
     * 
     */
    public Output<String> adaptive() {
        return this.adaptive;
    }
    /**
     * Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
     * 
     */
    @Export(name="adaptiveLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> adaptiveLimit;

    /**
     * @return Specifies the absolute number of milliseconds that may not be exceeded by a monitor probe, regardless of Allowed Divergence.
     * 
     */
    public Output<Integer> adaptiveLimit() {
        return this.adaptiveLimit;
    }
    /**
     * Specifies the location in the LDAP tree from which the monitor starts the health check
     * 
     */
    @Export(name="base", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> base;

    /**
     * @return Specifies the location in the LDAP tree from which the monitor starts the health check
     * 
     */
    public Output<Optional<String>> base() {
        return Codegen.optional(this.base);
    }
    /**
     * Specifies whether the system will query the LDAP servers pointed to by any referrals in the query results.
     * 
     */
    @Export(name="chaseReferrals", refs={String.class}, tree="[0]")
    private Output<String> chaseReferrals;

    /**
     * @return Specifies whether the system will query the LDAP servers pointed to by any referrals in the query results.
     * 
     */
    public Output<String> chaseReferrals() {
        return this.chaseReferrals;
    }
    /**
     * Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts &#39;enabled&#39; or &#39;disabled&#39; values, the default value is &#39;enabled&#39;.
     * 
     */
    @Export(name="compatibility", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> compatibility;

    /**
     * @return Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts &#39;enabled&#39; or &#39;disabled&#39; values, the default value is &#39;enabled&#39;.
     * 
     */
    public Output<Optional<String>> compatibility() {
        return Codegen.optional(this.compatibility);
    }
    /**
     * Custom parent monitor for the system to use for setting initial values for the new monitor.
     * 
     */
    @Export(name="customParent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customParent;

    /**
     * @return Custom parent monitor for the system to use for setting initial values for the new monitor.
     * 
     */
    public Output<Optional<String>> customParent() {
        return Codegen.optional(this.customParent);
    }
    /**
     * Specifies the database in which the user is created
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> database;

    /**
     * @return Specifies the database in which the user is created
     * 
     */
    public Output<Optional<String>> database() {
        return Codegen.optional(this.database);
    }
    /**
     * Specify an alias address for monitoring
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return Specify an alias address for monitoring
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     * 
     */
    @Export(name="filename", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> filename;

    /**
     * @return Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
     * 
     */
    public Output<Optional<String>> filename() {
        return Codegen.optional(this.filename);
    }
    /**
     * Specifies an LDAP key for which the monitor searches
     * 
     */
    @Export(name="filter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> filter;

    /**
     * @return Specifies an LDAP key for which the monitor searches
     * 
     */
    public Output<Optional<String>> filter() {
        return Codegen.optional(this.filter);
    }
    /**
     * Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
     * 
     */
    @Export(name="interval", refs={Integer.class}, tree="[0]")
    private Output<Integer> interval;

    /**
     * @return Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown,value of `interval` should be always less than `timeout`. Default is `5`.
     * 
     */
    public Output<Integer> interval() {
        return this.interval;
    }
    /**
     * Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
     * 
     */
    @Export(name="ipDscp", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipDscp;

    /**
     * @return Displays the differentiated services code point (DSCP).The default is `0 (zero)`.
     * 
     */
    public Output<Integer> ipDscp() {
        return this.ipDscp;
    }
    /**
     * Specifies whether the target must include attributes in its response to be considered up. The options are no (Specifies
     * that the system performs only a one-level search (based on the Filter setting), and does not require that the target
     * returns any attributes.) and yes (Specifies that the system performs a sub-tree search, and if the target returns no
     * attributes, the target is considered down.)
     * 
     */
    @Export(name="mandatoryAttributes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mandatoryAttributes;

    /**
     * @return Specifies whether the target must include attributes in its response to be considered up. The options are no (Specifies
     * that the system performs only a one-level search (based on the Filter setting), and does not require that the target
     * returns any attributes.) and yes (Specifies that the system performs a sub-tree search, and if the target returns no
     * attributes, the target is considered down.)
     * 
     */
    public Output<Optional<String>> mandatoryAttributes() {
        return Codegen.optional(this.mandatoryAttributes);
    }
    /**
     * Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     * 
     */
    @Export(name="manualResume", refs={String.class}, tree="[0]")
    private Output<String> manualResume;

    /**
     * @return Specifies whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     * 
     */
    public Output<String> manualResume() {
        return this.manualResume;
    }
    /**
     * Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output<String> mode;

    /**
     * @return Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the Name of the LTM Monitor.Name of Monitor should be full path,full path is the combination of the `partition + monitor name`,For ex:`/Common/test-ltm-monitor`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Parent monitor for the system to use for setting initial values for the new monitor.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return Parent monitor for the system to use for setting initial values for the new monitor.
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Specifies the password if the monitored target requires authentication
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Specifies the password if the monitored target requires authentication
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
     * 
     */
    @Export(name="receive", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> receive;

    /**
     * @return Specifies the regular expression representing the text string that the monitor looks for in the returned resource.
     * 
     */
    public Output<Optional<String>> receive() {
        return Codegen.optional(this.receive);
    }
    /**
     * The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
     * 
     */
    @Export(name="receiveDisable", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> receiveDisable;

    /**
     * @return The system marks the node or pool member disabled when its response matches Receive Disable String but not Receive String.
     * 
     */
    public Output<Optional<String>> receiveDisable() {
        return Codegen.optional(this.receiveDisable);
    }
    /**
     * Instructs the system to mark the target resource down when the test is successful.
     * 
     */
    @Export(name="reverse", refs={String.class}, tree="[0]")
    private Output<String> reverse;

    /**
     * @return Instructs the system to mark the target resource down when the test is successful.
     * 
     */
    public Output<String> reverse() {
        return this.reverse;
    }
    /**
     * Specifies the secure communications protocol that the monitor uses to communicate with the target. The options are none
     * (Specifies that the system does not use a security protocol for communications with the target.), ssl (Specifies that
     * the system uses the SSL protocol for communications with the target.), and tls (Specifies that the system uses the TLS
     * protocol for communications with the target.)
     * 
     */
    @Export(name="security", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> security;

    /**
     * @return Specifies the secure communications protocol that the monitor uses to communicate with the target. The options are none
     * (Specifies that the system does not use a security protocol for communications with the target.), ssl (Specifies that
     * the system uses the SSL protocol for communications with the target.), and tls (Specifies that the system uses the TLS
     * protocol for communications with the target.)
     * 
     */
    public Output<Optional<String>> security() {
        return Codegen.optional(this.security);
    }
    /**
     * Specifies the text string that the monitor sends to the target object.
     * 
     */
    @Export(name="send", refs={String.class}, tree="[0]")
    private Output<String> send;

    /**
     * @return Specifies the text string that the monitor sends to the target object.
     * 
     */
    public Output<String> send() {
        return this.send;
    }
    /**
     * Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
     * 
     */
    @Export(name="sslProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslProfile;

    /**
     * @return Specifies the ssl profile for the monitor. It only makes sense when the parent is `/Common/https`
     * 
     */
    public Output<Optional<String>> sslProfile() {
        return Codegen.optional(this.sslProfile);
    }
    /**
     * Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
     * 
     */
    @Export(name="timeUntilUp", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeUntilUp;

    /**
     * @return Specifies the number of seconds to wait after a resource first responds correctly to the monitor before setting the resource to up.
     * 
     */
    public Output<Integer> timeUntilUp() {
        return this.timeUntilUp;
    }
    /**
     * Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeout;

    /**
     * @return Specifies the number of seconds the target has in which to respond to the monitor request. The default is `16` seconds
     * 
     */
    public Output<Integer> timeout() {
        return this.timeout;
    }
    /**
     * Specifies whether the monitor operates in transparent mode.
     * 
     */
    @Export(name="transparent", refs={String.class}, tree="[0]")
    private Output<String> transparent;

    /**
     * @return Specifies whether the monitor operates in transparent mode.
     * 
     */
    public Output<String> transparent() {
        return this.transparent;
    }
    /**
     * Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
     * 
     */
    @Export(name="upInterval", refs={Integer.class}, tree="[0]")
    private Output<Integer> upInterval;

    /**
     * @return Specifies the interval for the system to use to perform the health check when a resource is up. The default is `0(Disabled)`
     * 
     */
    public Output<Integer> upInterval() {
        return this.upInterval;
    }
    /**
     * Specifies the user name if the monitored target requires authentication
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return Specifies the user name if the monitored target requires authentication
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Monitor(java.lang.String name) {
        this(name, MonitorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Monitor(java.lang.String name, MonitorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Monitor(java.lang.String name, MonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/monitor:Monitor", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Monitor(java.lang.String name, Output<java.lang.String> id, @Nullable MonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/monitor:Monitor", name, state, makeResourceOptions(options, id), false);
    }

    private static MonitorArgs makeArgs(MonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MonitorArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static Monitor get(java.lang.String name, Output<java.lang.String> id, @Nullable MonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Monitor(name, id, state, options);
    }
}
