// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.VirtualAddressArgs;
import com.pulumi.f5bigip.ltm.inputs.VirtualAddressState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.VirtualAddress` Configures Virtual Server
 * 
 * For resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource. For example /Common/virtual_server.
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
 * import com.pulumi.f5bigip.ltm.VirtualAddress;
 * import com.pulumi.f5bigip.ltm.VirtualAddressArgs;
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
 *         var vsVa = new VirtualAddress("vsVa", VirtualAddressArgs.builder()
 *             .name("/Common/xxxxx")
 *             .advertizeRoute("enabled")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:ltm/virtualAddress:VirtualAddress")
public class VirtualAddress extends com.pulumi.resources.CustomResource {
    /**
     * Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article &lt;https://support.f5.com/csp/article/K85543242&gt; )
     * 
     */
    @Export(name="advertizeRoute", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> advertizeRoute;

    /**
     * @return Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article &lt;https://support.f5.com/csp/article/K85543242&gt; )
     * 
     */
    public Output<Optional<String>> advertizeRoute() {
        return Codegen.optional(this.advertizeRoute);
    }
    /**
     * Enable or disable ARP for the virtual address
     * 
     */
    @Export(name="arp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> arp;

    /**
     * @return Enable or disable ARP for the virtual address
     * 
     */
    public Output<Optional<Boolean>> arp() {
        return Codegen.optional(this.arp);
    }
    /**
     * Automatically delete the virtual address with the virtual server
     * 
     */
    @Export(name="autoDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoDelete;

    /**
     * @return Automatically delete the virtual address with the virtual server
     * 
     */
    public Output<Optional<Boolean>> autoDelete() {
        return Codegen.optional(this.autoDelete);
    }
    /**
     * Max number of connections for virtual address
     * 
     */
    @Export(name="connLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> connLimit;

    /**
     * @return Max number of connections for virtual address
     * 
     */
    public Output<Optional<Integer>> connLimit() {
        return Codegen.optional(this.connLimit);
    }
    /**
     * Enable or disable the virtual address
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Enable or disable the virtual address
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     * 
     */
    @Export(name="icmpEcho", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icmpEcho;

    /**
     * @return Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     * 
     */
    public Output<Optional<String>> icmpEcho() {
        return Codegen.optional(this.icmpEcho);
    }
    /**
     * Name of the virtual address
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the virtual address
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specify the partition and traffic group
     * 
     */
    @Export(name="trafficGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trafficGroup;

    /**
     * @return Specify the partition and traffic group
     * 
     */
    public Output<Optional<String>> trafficGroup() {
        return Codegen.optional(this.trafficGroup);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualAddress(java.lang.String name) {
        this(name, VirtualAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualAddress(java.lang.String name, VirtualAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualAddress(java.lang.String name, VirtualAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/virtualAddress:VirtualAddress", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VirtualAddress(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/virtualAddress:VirtualAddress", name, state, makeResourceOptions(options, id), false);
    }

    private static VirtualAddressArgs makeArgs(VirtualAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VirtualAddressArgs.Empty : args;
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
    public static VirtualAddress get(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualAddress(name, id, state, options);
    }
}
