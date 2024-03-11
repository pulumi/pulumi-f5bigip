// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.net;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.net.VlanArgs;
import com.pulumi.f5bigip.net.inputs.VlanState;
import com.pulumi.f5bigip.net.outputs.VlanInterface;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.net.Vlan` Manages a vlan configuration
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
 * import com.pulumi.f5bigip.net.Vlan;
 * import com.pulumi.f5bigip.net.VlanArgs;
 * import com.pulumi.f5bigip.net.inputs.VlanInterfaceArgs;
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
 *         var vlan1 = new Vlan(&#34;vlan1&#34;, VlanArgs.builder()        
 *             .interfaces(VlanInterfaceArgs.builder()
 *                 .tagged(false)
 *                 .vlanport(1.2)
 *                 .build())
 *             .name(&#34;/Common/Internal&#34;)
 *             .tag(101)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:net/vlan:Vlan")
public class Vlan extends com.pulumi.resources.CustomResource {
    /**
     * Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [`default`, `src-ip`, `dst-ip`]
     * 
     */
    @Export(name="cmpHash", refs={String.class}, tree="[0]")
    private Output<String> cmpHash;

    /**
     * @return Specifies how the traffic on the VLAN will be disaggregated. The value selected determines the traffic disaggregation method. possible options: [`default`, `src-ip`, `dst-ip`]
     * 
     */
    public Output<String> cmpHash() {
        return this.cmpHash;
    }
    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     * 
     */
    @Export(name="interfaces", refs={List.class,VlanInterface.class}, tree="[0,1]")
    private Output</* @Nullable */ List<VlanInterface>> interfaces;

    /**
     * @return Specifies which interfaces you want this VLAN to use for traffic management.
     * 
     */
    public Output<Optional<List<VlanInterface>>> interfaces() {
        return Codegen.optional(this.interfaces);
    }
    /**
     * Specifies the maximum transmission unit (MTU) for traffic on this VLAN. The default value is `1500`.
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> mtu;

    /**
     * @return Specifies the maximum transmission unit (MTU) for traffic on this VLAN. The default value is `1500`.
     * 
     */
    public Output<Optional<Integer>> mtu() {
        return Codegen.optional(this.mtu);
    }
    /**
     * Name of the vlan
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the vlan
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     * 
     */
    @Export(name="tag", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tag;

    /**
     * @return Specifies a number that the system adds into the header of any frame passing through the VLAN.
     * 
     */
    public Output<Optional<Integer>> tag() {
        return Codegen.optional(this.tag);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vlan(String name) {
        this(name, VlanArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vlan(String name, VlanArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vlan(String name, VlanArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:net/vlan:Vlan", name, args == null ? VlanArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vlan(String name, Output<String> id, @Nullable VlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:net/vlan:Vlan", name, state, makeResourceOptions(options, id));
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
    public static Vlan get(String name, Output<String> id, @Nullable VlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vlan(name, id, state, options);
    }
}
