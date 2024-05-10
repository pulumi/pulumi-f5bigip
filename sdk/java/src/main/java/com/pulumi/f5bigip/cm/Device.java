// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.cm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.cm.DeviceArgs;
import com.pulumi.f5bigip.cm.inputs.DeviceState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.cm.Device` provides details about a specific bigip
 * 
 * This resource is helpful when configuring the BIG-IP device in cluster or in HA mode.
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
 * import com.pulumi.f5bigip.cm.Device;
 * import com.pulumi.f5bigip.cm.DeviceArgs;
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
 *         var myNewDevice = new Device("myNewDevice", DeviceArgs.builder()        
 *             .name("bigip300.f5.com")
 *             .configsyncIp("2.2.2.2")
 *             .mirrorIp("10.10.10.10")
 *             .mirrorSecondaryIp("11.11.11.11")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:cm/device:Device")
public class Device extends com.pulumi.resources.CustomResource {
    /**
     * IP address used for config sync
     * 
     */
    @Export(name="configsyncIp", refs={String.class}, tree="[0]")
    private Output<String> configsyncIp;

    /**
     * @return IP address used for config sync
     * 
     */
    public Output<String> configsyncIp() {
        return this.configsyncIp;
    }
    /**
     * IP address used for state mirroring
     * 
     */
    @Export(name="mirrorIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mirrorIp;

    /**
     * @return IP address used for state mirroring
     * 
     */
    public Output<Optional<String>> mirrorIp() {
        return Codegen.optional(this.mirrorIp);
    }
    /**
     * Secondary IP address used for state mirroring
     * 
     */
    @Export(name="mirrorSecondaryIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mirrorSecondaryIp;

    /**
     * @return Secondary IP address used for state mirroring
     * 
     */
    public Output<Optional<String>> mirrorSecondaryIp() {
        return Codegen.optional(this.mirrorSecondaryIp);
    }
    /**
     * Address of the Device which needs to be Deviceensed
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Address of the Device which needs to be Deviceensed
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Device(String name) {
        this(name, DeviceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Device(String name, DeviceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Device(String name, DeviceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:cm/device:Device", name, args == null ? DeviceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Device(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:cm/device:Device", name, state, makeResourceOptions(options, id));
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
    public static Device get(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Device(name, id, state, options);
    }
}
