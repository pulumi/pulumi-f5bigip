// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.sys.SnmpTrapsArgs;
import com.pulumi.f5bigip.sys.inputs.SnmpTrapsState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP
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
 * import com.pulumi.f5bigip.sys.SnmpTraps;
 * import com.pulumi.f5bigip.sys.SnmpTrapsArgs;
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
 *         var snmpTraps = new SnmpTraps("snmpTraps", SnmpTrapsArgs.builder()
 *             .name("snmptraps")
 *             .community("f5community")
 *             .host("195.10.10.1")
 *             .description("Setup snmp traps")
 *             .port(111)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:sys/snmpTraps:SnmpTraps")
public class SnmpTraps extends com.pulumi.resources.CustomResource {
    /**
     * Encrypted password
     * 
     */
    @Export(name="authPasswordencrypted", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authPasswordencrypted;

    /**
     * @return Encrypted password
     * 
     */
    public Output<Optional<String>> authPasswordencrypted() {
        return Codegen.optional(this.authPasswordencrypted);
    }
    /**
     * Specifies the protocol used to authenticate the user.
     * 
     */
    @Export(name="authProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authProtocol;

    /**
     * @return Specifies the protocol used to authenticate the user.
     * 
     */
    public Output<Optional<String>> authProtocol() {
        return Codegen.optional(this.authProtocol);
    }
    /**
     * Specifies the community string used for this trap.
     * 
     */
    @Export(name="community", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> community;

    /**
     * @return Specifies the community string used for this trap.
     * 
     */
    public Output<Optional<String>> community() {
        return Codegen.optional(this.community);
    }
    /**
     * The port that the trap will be sent to.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The port that the trap will be sent to.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies the authoritative security engine for SNMPv3.
     * 
     */
    @Export(name="engineId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> engineId;

    /**
     * @return Specifies the authoritative security engine for SNMPv3.
     * 
     */
    public Output<Optional<String>> engineId() {
        return Codegen.optional(this.engineId);
    }
    /**
     * The host the trap will be sent to.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> host;

    /**
     * @return The host the trap will be sent to.
     * 
     */
    public Output<Optional<String>> host() {
        return Codegen.optional(this.host);
    }
    /**
     * Name of the snmp trap.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> name;

    /**
     * @return Name of the snmp trap.
     * 
     */
    public Output<Optional<String>> name() {
        return Codegen.optional(this.name);
    }
    /**
     * User defined description.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> port;

    /**
     * @return User defined description.
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
    }
    /**
     * Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     * 
     */
    @Export(name="privacyPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privacyPassword;

    /**
     * @return Specifies the clear text password used to encrypt traffic. This field will not be displayed.
     * 
     */
    public Output<Optional<String>> privacyPassword() {
        return Codegen.optional(this.privacyPassword);
    }
    /**
     * Specifies the encrypted password used to encrypt traffic.
     * 
     */
    @Export(name="privacyPasswordEncrypted", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privacyPasswordEncrypted;

    /**
     * @return Specifies the encrypted password used to encrypt traffic.
     * 
     */
    public Output<Optional<String>> privacyPasswordEncrypted() {
        return Codegen.optional(this.privacyPasswordEncrypted);
    }
    /**
     * Specifies the protocol used to encrypt traffic.
     * 
     */
    @Export(name="privacyProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privacyProtocol;

    /**
     * @return Specifies the protocol used to encrypt traffic.
     * 
     */
    public Output<Optional<String>> privacyProtocol() {
        return Codegen.optional(this.privacyProtocol);
    }
    /**
     * Specifies whether or not traffic is encrypted and whether or not authentication is required.
     * 
     */
    @Export(name="securityLevel", refs={String.class}, tree="[0]")
    private Output<String> securityLevel;

    /**
     * @return Specifies whether or not traffic is encrypted and whether or not authentication is required.
     * 
     */
    public Output<String> securityLevel() {
        return this.securityLevel;
    }
    /**
     * Security name used in conjunction with SNMPv3.
     * 
     */
    @Export(name="securityName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityName;

    /**
     * @return Security name used in conjunction with SNMPv3.
     * 
     */
    public Output<Optional<String>> securityName() {
        return Codegen.optional(this.securityName);
    }
    /**
     * SNMP version used for sending the trap.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return SNMP version used for sending the trap.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SnmpTraps(java.lang.String name) {
        this(name, SnmpTrapsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SnmpTraps(java.lang.String name, @Nullable SnmpTrapsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SnmpTraps(java.lang.String name, @Nullable SnmpTrapsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/snmpTraps:SnmpTraps", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SnmpTraps(java.lang.String name, Output<java.lang.String> id, @Nullable SnmpTrapsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/snmpTraps:SnmpTraps", name, state, makeResourceOptions(options, id), false);
    }

    private static SnmpTrapsArgs makeArgs(@Nullable SnmpTrapsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SnmpTrapsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static SnmpTraps get(java.lang.String name, Output<java.lang.String> id, @Nullable SnmpTrapsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SnmpTraps(name, id, state, options);
    }
}
