// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.IpsecPolicyArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.IpsecPolicyState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `f5bigip.IpsecPolicy` Manage IPSec policies on a BIG-IP
 * 
 * Resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name (example: /Common/test-policy)
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
 * import com.pulumi.f5bigip.IpsecPolicy;
 * import com.pulumi.f5bigip.IpsecPolicyArgs;
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
 *         var test_policy = new IpsecPolicy("test-policy", IpsecPolicyArgs.builder()
 *             .name("/Common/test-policy")
 *             .description("created by terraform provider")
 *             .protocol("esp")
 *             .mode("tunnel")
 *             .tunnelLocalAddress("192.168.1.1")
 *             .tunnelRemoteAddress("10.10.1.1")
 *             .authAlgorithm("sha1")
 *             .encryptAlgorithm("3des")
 *             .lifetime(3)
 *             .ipcomp("deflate")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/ipsecPolicy:IpsecPolicy")
public class IpsecPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     * 
     */
    @Export(name="authAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> authAlgorithm;

    /**
     * @return Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     * 
     */
    public Output<String> authAlgorithm() {
        return this.authAlgorithm;
    }
    /**
     * Description of the IPSec policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the IPSec policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     * 
     */
    @Export(name="encryptAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> encryptAlgorithm;

    /**
     * @return Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     * 
     */
    public Output<String> encryptAlgorithm() {
        return this.encryptAlgorithm;
    }
    /**
     * Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
     * 
     */
    @Export(name="ipcomp", refs={String.class}, tree="[0]")
    private Output<String> ipcomp;

    /**
     * @return Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
     * 
     */
    public Output<String> ipcomp() {
        return this.ipcomp;
    }
    /**
     * Specifies the length of time before the IKE security association expires, in kilobytes.
     * 
     */
    @Export(name="kbLifetime", refs={Integer.class}, tree="[0]")
    private Output<Integer> kbLifetime;

    /**
     * @return Specifies the length of time before the IKE security association expires, in kilobytes.
     * 
     */
    public Output<Integer> kbLifetime() {
        return this.kbLifetime;
    }
    /**
     * Specifies the length of time before the IKE security association expires, in minutes.
     * 
     */
    @Export(name="lifetime", refs={Integer.class}, tree="[0]")
    private Output<Integer> lifetime;

    /**
     * @return Specifies the length of time before the IKE security association expires, in minutes.
     * 
     */
    public Output<Integer> lifetime() {
        return this.lifetime;
    }
    /**
     * Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output<String> mode;

    /**
     * @return Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     * 
     */
    @Export(name="perfectForwardSecrecy", refs={String.class}, tree="[0]")
    private Output<String> perfectForwardSecrecy;

    /**
     * @return Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     * 
     */
    public Output<String> perfectForwardSecrecy() {
        return this.perfectForwardSecrecy;
    }
    /**
     * Specifies the IPsec protocol. Valid choices are: `ah, esp`
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return Specifies the IPsec protocol. Valid choices are: `ah, esp`
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    @Export(name="tunnelLocalAddress", refs={String.class}, tree="[0]")
    private Output<String> tunnelLocalAddress;

    /**
     * @return Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    public Output<String> tunnelLocalAddress() {
        return this.tunnelLocalAddress;
    }
    /**
     * Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    @Export(name="tunnelRemoteAddress", refs={String.class}, tree="[0]")
    private Output<String> tunnelRemoteAddress;

    /**
     * @return Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    public Output<String> tunnelRemoteAddress() {
        return this.tunnelRemoteAddress;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpsecPolicy(java.lang.String name) {
        this(name, IpsecPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpsecPolicy(java.lang.String name, IpsecPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpsecPolicy(java.lang.String name, IpsecPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/ipsecPolicy:IpsecPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IpsecPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable IpsecPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/ipsecPolicy:IpsecPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static IpsecPolicyArgs makeArgs(IpsecPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IpsecPolicyArgs.Empty : args;
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
    public static IpsecPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable IpsecPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpsecPolicy(name, id, state, options);
    }
}
