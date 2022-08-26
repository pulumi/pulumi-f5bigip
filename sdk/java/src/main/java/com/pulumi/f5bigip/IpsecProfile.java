// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.IpsecProfileArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.IpsecProfileState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.IpsecProfile` Manage IPSec Profiles on a BIG-IP
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.IpsecProfile;
 * import com.pulumi.f5bigip.IpsecProfileArgs;
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
 *         var azurevWANProfile = new IpsecProfile(&#34;azurevWANProfile&#34;, IpsecProfileArgs.builder()        
 *             .description(&#34;mytestipsecprofile&#34;)
 *             .name(&#34;/Common/Mytestipsecprofile&#34;)
 *             .trafficSelector(&#34;test-trafficselector&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:index/ipsecProfile:IpsecProfile")
public class IpsecProfile extends com.pulumi.resources.CustomResource {
    /**
     * Specifies descriptive text that identifies the IPsec interface tunnel profile.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return Specifies descriptive text that identifies the IPsec interface tunnel profile.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     * 
     */
    @Export(name="parentProfile", type=String.class, parameters={})
    private Output</* @Nullable */ String> parentProfile;

    /**
     * @return Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     * 
     */
    public Output<Optional<String>> parentProfile() {
        return Codegen.optional(this.parentProfile);
    }
    /**
     * Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     * 
     */
    @Export(name="trafficSelector", type=String.class, parameters={})
    private Output<String> trafficSelector;

    /**
     * @return Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     * 
     */
    public Output<String> trafficSelector() {
        return this.trafficSelector;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpsecProfile(String name) {
        this(name, IpsecProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpsecProfile(String name, IpsecProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpsecProfile(String name, IpsecProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/ipsecProfile:IpsecProfile", name, args == null ? IpsecProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpsecProfile(String name, Output<String> id, @Nullable IpsecProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/ipsecProfile:IpsecProfile", name, state, makeResourceOptions(options, id));
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
    public static IpsecProfile get(String name, Output<String> id, @Nullable IpsecProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpsecProfile(name, id, state, options);
    }
}
