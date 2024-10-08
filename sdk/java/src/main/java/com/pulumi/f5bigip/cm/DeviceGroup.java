// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.cm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.cm.DeviceGroupArgs;
import com.pulumi.f5bigip.cm.inputs.DeviceGroupState;
import com.pulumi.f5bigip.cm.outputs.DeviceGroupDevice;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.cm.DeviceGroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
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
 * import com.pulumi.f5bigip.cm.DeviceGroup;
 * import com.pulumi.f5bigip.cm.DeviceGroupArgs;
 * import com.pulumi.f5bigip.cm.inputs.DeviceGroupDeviceArgs;
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
 *         var myNewDevicegroup = new DeviceGroup("myNewDevicegroup", DeviceGroupArgs.builder()
 *             .name("sanjose_devicegroup")
 *             .autoSync("enabled")
 *             .fullLoadOnSync("true")
 *             .type("sync-only")
 *             .devices(            
 *                 DeviceGroupDeviceArgs.builder()
 *                     .name("bigip1.cisco.com")
 *                     .build(),
 *                 DeviceGroupDeviceArgs.builder()
 *                     .name("bigip200.f5.com")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:cm/deviceGroup:DeviceGroup")
public class DeviceGroup extends com.pulumi.resources.CustomResource {
    /**
     * Specifies if the device-group will automatically sync configuration data to its members
     * 
     */
    @Export(name="autoSync", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoSync;

    /**
     * @return Specifies if the device-group will automatically sync configuration data to its members
     * 
     */
    public Output<Optional<String>> autoSync() {
        return Codegen.optional(this.autoSync);
    }
    /**
     * Description of Device group
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of Device group
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the device to be included in device group, this need to be configured before using devicegroup resource
     * 
     */
    @Export(name="devices", refs={List.class,DeviceGroupDevice.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DeviceGroupDevice>> devices;

    /**
     * @return Name of the device to be included in device group, this need to be configured before using devicegroup resource
     * 
     */
    public Output<Optional<List<DeviceGroupDevice>>> devices() {
        return Codegen.optional(this.devices);
    }
    /**
     * Specifies if the device-group will perform a full-load upon sync
     * 
     */
    @Export(name="fullLoadOnSync", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fullLoadOnSync;

    /**
     * @return Specifies if the device-group will perform a full-load upon sync
     * 
     */
    public Output<Optional<String>> fullLoadOnSync() {
        return Codegen.optional(this.fullLoadOnSync);
    }
    /**
     * Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
     * 
     */
    @Export(name="incrementalConfig", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> incrementalConfig;

    /**
     * @return Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
     * 
     */
    public Output<Optional<Integer>> incrementalConfig() {
        return Codegen.optional(this.incrementalConfig);
    }
    /**
     * Is the name of the device Group
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> name;

    /**
     * @return Is the name of the device Group
     * 
     */
    public Output<Optional<String>> name() {
        return Codegen.optional(this.name);
    }
    /**
     * Specifies if the device-group will use a network connection for failover
     * 
     */
    @Export(name="networkFailover", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> networkFailover;

    /**
     * @return Specifies if the device-group will use a network connection for failover
     * 
     */
    public Output<Optional<String>> networkFailover() {
        return Codegen.optional(this.networkFailover);
    }
    /**
     * Device administrative partition
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Device administrative partition
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * Specifies whether the configuration should be saved upon auto-sync.
     * 
     */
    @Export(name="saveOnAutoSync", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> saveOnAutoSync;

    /**
     * @return Specifies whether the configuration should be saved upon auto-sync.
     * 
     */
    public Output<Optional<String>> saveOnAutoSync() {
        return Codegen.optional(this.saveOnAutoSync);
    }
    /**
     * Specifies if the device-group will be used for failover or resource syncing
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Specifies if the device-group will be used for failover or resource syncing
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeviceGroup(java.lang.String name) {
        this(name, DeviceGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeviceGroup(java.lang.String name, @Nullable DeviceGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeviceGroup(java.lang.String name, @Nullable DeviceGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:cm/deviceGroup:DeviceGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DeviceGroup(java.lang.String name, Output<java.lang.String> id, @Nullable DeviceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:cm/deviceGroup:DeviceGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static DeviceGroupArgs makeArgs(@Nullable DeviceGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DeviceGroupArgs.Empty : args;
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
    public static DeviceGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable DeviceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeviceGroup(name, id, state, options);
    }
}
