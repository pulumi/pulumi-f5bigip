// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.TrafficSelectorArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.TrafficSelectorState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `f5bigip.TrafficSelector` Manage IPSec Traffic Selectors on BIG-IP
 * 
 * Resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name (example: /Common/test-selector)
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
 * import com.pulumi.f5bigip.TrafficSelector;
 * import com.pulumi.f5bigip.TrafficSelectorArgs;
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
 *         var test_selector = new TrafficSelector(&#34;test-selector&#34;, TrafficSelectorArgs.builder()        
 *             .destinationAddress(&#34;3.10.11.2/32&#34;)
 *             .name(&#34;/Common/test-selector&#34;)
 *             .sourceAddress(&#34;2.10.11.12/32&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/trafficSelector:TrafficSelector")
public class TrafficSelector extends com.pulumi.resources.CustomResource {
    /**
     * Description of the traffic selector.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the traffic selector.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
     * 
     */
    @Export(name="destinationAddress", refs={String.class}, tree="[0]")
    private Output<String> destinationAddress;

    /**
     * @return Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
     * 
     */
    public Output<String> destinationAddress() {
        return this.destinationAddress;
    }
    /**
     * Specifies the IP port used by the application. The default value is `All Ports (0)`
     * 
     */
    @Export(name="destinationPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> destinationPort;

    /**
     * @return Specifies the IP port used by the application. The default value is `All Ports (0)`
     * 
     */
    public Output<Integer> destinationPort() {
        return this.destinationPort;
    }
    /**
     * Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
     * 
     */
    @Export(name="ipProtocol", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipProtocol;

    /**
     * @return Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
     * 
     */
    public Output<Integer> ipProtocol() {
        return this.ipProtocol;
    }
    /**
     * Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
     * 
     */
    @Export(name="ipsecPolicy", refs={String.class}, tree="[0]")
    private Output<String> ipsecPolicy;

    /**
     * @return Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
     * 
     */
    public Output<String> ipsecPolicy() {
        return this.ipsecPolicy;
    }
    /**
     * Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
     * When creating a new traffic selector, if this parameter is not specified, the default is `last`
     * 
     */
    @Export(name="order", refs={Integer.class}, tree="[0]")
    private Output<Integer> order;

    /**
     * @return Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
     * When creating a new traffic selector, if this parameter is not specified, the default is `last`
     * 
     */
    public Output<Integer> order() {
        return this.order;
    }
    /**
     * Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
     * 
     */
    @Export(name="sourceAddress", refs={String.class}, tree="[0]")
    private Output<String> sourceAddress;

    /**
     * @return Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
     * 
     */
    public Output<String> sourceAddress() {
        return this.sourceAddress;
    }
    /**
     * Specifies the IP port used by the application. The default value is `All Ports (0)`.
     * 
     */
    @Export(name="sourcePort", refs={Integer.class}, tree="[0]")
    private Output<Integer> sourcePort;

    /**
     * @return Specifies the IP port used by the application. The default value is `All Ports (0)`.
     * 
     */
    public Output<Integer> sourcePort() {
        return this.sourcePort;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TrafficSelector(String name) {
        this(name, TrafficSelectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TrafficSelector(String name, TrafficSelectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TrafficSelector(String name, TrafficSelectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/trafficSelector:TrafficSelector", name, args == null ? TrafficSelectorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TrafficSelector(String name, Output<String> id, @Nullable TrafficSelectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/trafficSelector:TrafficSelector", name, state, makeResourceOptions(options, id));
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
    public static TrafficSelector get(String name, Output<String> id, @Nullable TrafficSelectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TrafficSelector(name, id, state, options);
    }
}
