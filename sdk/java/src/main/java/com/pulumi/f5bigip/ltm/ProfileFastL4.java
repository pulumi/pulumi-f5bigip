// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileFastL4Args;
import com.pulumi.f5bigip.ltm.inputs.ProfileFastL4State;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileFastL4` Configures a custom LTM fastL4 profile for use by health checks.
 * 
 * Resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.ProfileFastL4;
 * import com.pulumi.f5bigip.ltm.ProfileFastL4Args;
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
 *         var profileFastl4 = new ProfileFastL4(&#34;profileFastl4&#34;, ProfileFastL4Args.builder()        
 *             .clientTimeout(40)
 *             .defaultsFrom(&#34;/Common/fastL4&#34;)
 *             .explicitflowMigration(&#34;enabled&#34;)
 *             .hardwareSyncookie(&#34;enabled&#34;)
 *             .idleTimeout(&#34;200&#34;)
 *             .iptosToclient(&#34;pass-through&#34;)
 *             .iptosToserver(&#34;pass-through&#34;)
 *             .keepaliveInterval(&#34;disabled&#34;)
 *             .name(&#34;/Common/sjfastl4profile&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * BIG-IP LTM fastl4 profiles can be imported using the `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import f5bigip:ltm/profileFastL4:ProfileFastL4 test-fastl4 /Common/test-fastl4
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileFastL4:ProfileFastL4")
public class ProfileFastL4 extends com.pulumi.resources.CustomResource {
    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     * 
     */
    @Export(name="clientTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> clientTimeout;

    /**
     * @return Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     * 
     */
    public Output<Integer> clientTimeout() {
        return this.clientTimeout;
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
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     * 
     */
    @Export(name="explicitflowMigration", refs={String.class}, tree="[0]")
    private Output<String> explicitflowMigration;

    /**
     * @return Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     * 
     */
    public Output<String> explicitflowMigration() {
        return this.explicitflowMigration;
    }
    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     * 
     */
    @Export(name="hardwareSyncookie", refs={String.class}, tree="[0]")
    private Output<String> hardwareSyncookie;

    /**
     * @return Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     * 
     */
    public Output<String> hardwareSyncookie() {
        return this.hardwareSyncookie;
    }
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    @Export(name="idleTimeout", refs={String.class}, tree="[0]")
    private Output<String> idleTimeout;

    /**
     * @return Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    public Output<String> idleTimeout() {
        return this.idleTimeout;
    }
    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    @Export(name="iptosToclient", refs={String.class}, tree="[0]")
    private Output<String> iptosToclient;

    /**
     * @return Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    public Output<String> iptosToclient() {
        return this.iptosToclient;
    }
    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    @Export(name="iptosToserver", refs={String.class}, tree="[0]")
    private Output<String> iptosToserver;

    /**
     * @return Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    public Output<String> iptosToserver() {
        return this.iptosToserver;
    }
    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     * 
     */
    @Export(name="keepaliveInterval", refs={String.class}, tree="[0]")
    private Output<String> keepaliveInterval;

    /**
     * @return Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     * 
     */
    public Output<String> keepaliveInterval() {
        return this.keepaliveInterval;
    }
    /**
     * Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
     * 
     */
    @Export(name="lateBinding", refs={String.class}, tree="[0]")
    private Output<String> lateBinding;

    /**
     * @return Enables intelligent selection of a back-end server or pool, using an iRule to make the selection. The default is `disabled`.
     * 
     */
    public Output<String> lateBinding() {
        return this.lateBinding;
    }
    /**
     * Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
     * 
     */
    @Export(name="looseClose", refs={String.class}, tree="[0]")
    private Output<String> looseClose;

    /**
     * @return Specifies, when checked (enabled), that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default is disabled.
     * 
     */
    public Output<String> looseClose() {
        return this.looseClose;
    }
    /**
     * Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
     * 
     */
    @Export(name="looseInitiation", refs={String.class}, tree="[0]")
    private Output<String> looseInitiation;

    /**
     * @return Specifies, when checked (enabled), that the system initializes a connection when it receives any TCP packet, rather that requiring a SYN packet for connection initiation. The default is disabled. We recommend that if you enable the Loose Initiation option, you also enable the Loose Close option.
     * 
     */
    public Output<String> looseInitiation() {
        return this.looseInitiation;
    }
    /**
     * Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the LTM fastL4 Profile.The full path is the combination of the `partition + name` of the resource (For example `/Common/my-fastl4profile`) or  `partition + directory + name` of the resource  (example: `/Common/test/my-fastl4profile`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * name of partition
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output<String> partition;

    /**
     * @return name of partition
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }
    /**
     * Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
     * 
     */
    @Export(name="receiveWindowsize", refs={Integer.class}, tree="[0]")
    private Output<Integer> receiveWindowsize;

    /**
     * @return Specifies the amount of data the BIG-IP system can accept without acknowledging the server. The default is 0 (zero).
     * 
     */
    public Output<Integer> receiveWindowsize() {
        return this.receiveWindowsize;
    }
    /**
     * Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
     * 
     */
    @Export(name="tcpHandshakeTimeout", refs={String.class}, tree="[0]")
    private Output<String> tcpHandshakeTimeout;

    /**
     * @return Specifies the acceptable duration for a TCP handshake, that is, the maximum idle time between a client synchronization (SYN) and a client acknowledgment (ACK).The default is `5 seconds`.
     * 
     */
    public Output<String> tcpHandshakeTimeout() {
        return this.tcpHandshakeTimeout;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileFastL4(String name) {
        this(name, ProfileFastL4Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileFastL4(String name, ProfileFastL4Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileFastL4(String name, ProfileFastL4Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFastL4:ProfileFastL4", name, args == null ? ProfileFastL4Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileFastL4(String name, Output<String> id, @Nullable ProfileFastL4State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFastL4:ProfileFastL4", name, state, makeResourceOptions(options, id));
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
    public static ProfileFastL4 get(String name, Output<String> id, @Nullable ProfileFastL4State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileFastL4(name, id, state, options);
    }
}
