// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.BigIqAs3Args;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.BigIqAs3State;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.BigIqAs3` provides details about bigiq as3 resource
 * 
 * This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.
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
 * import com.pulumi.f5bigip.BigIqAs3;
 * import com.pulumi.f5bigip.BigIqAs3Args;
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
 *         // Example Usage for json file
 *         var exampletask = new BigIqAs3("exampletask", BigIqAs3Args.builder()
 *             .bigiqAddress("xx.xx.xxx.xx")
 *             .bigiqUser("xxxxx")
 *             .bigiqPassword("xxxxxxxxx")
 *             .as3Json(StdFunctions.file(FileArgs.builder()
 *                 .input("bigiq_example.json")
 *                 .build()).result())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/bigIqAs3:BigIqAs3")
public class BigIqAs3 extends com.pulumi.resources.CustomResource {
    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    @Export(name="as3Json", refs={String.class}, tree="[0]")
    private Output<String> as3Json;

    /**
     * @return Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    public Output<String> as3Json() {
        return this.as3Json;
    }
    /**
     * Address of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Export(name="bigiqAddress", refs={String.class}, tree="[0]")
    private Output<String> bigiqAddress;

    /**
     * @return Address of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqAddress() {
        return this.bigiqAddress;
    }
    /**
     * BIGIQ Login reference for token authentication
     * 
     */
    @Export(name="bigiqLoginRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigiqLoginRef;

    /**
     * @return BIGIQ Login reference for token authentication
     * 
     */
    public Output<Optional<String>> bigiqLoginRef() {
        return Codegen.optional(this.bigiqLoginRef);
    }
    /**
     * Password of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Export(name="bigiqPassword", refs={String.class}, tree="[0]")
    private Output<String> bigiqPassword;

    /**
     * @return Password of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqPassword() {
        return this.bigiqPassword;
    }
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     * 
     */
    @Export(name="bigiqPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigiqPort;

    /**
     * @return type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     * 
     */
    public Output<Optional<String>> bigiqPort() {
        return Codegen.optional(this.bigiqPort);
    }
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     * 
     */
    @Export(name="bigiqTokenAuth", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bigiqTokenAuth;

    /**
     * @return type `bool`, if set to `true` enables Token based Authentication,default is `false`
     * 
     */
    public Output<Optional<Boolean>> bigiqTokenAuth() {
        return Codegen.optional(this.bigiqTokenAuth);
    }
    /**
     * User name  of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Export(name="bigiqUser", refs={String.class}, tree="[0]")
    private Output<String> bigiqUser;

    /**
     * @return User name  of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqUser() {
        return this.bigiqUser;
    }
    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to `true`
     * 
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     * 
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     * 
     * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
     * 
     */
    @Export(name="ignoreMetadata", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreMetadata;

    /**
     * @return Set True if you want to ignore metadata changes during update. By default it is set to `true`
     * 
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     * 
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     * 
     * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
     * 
     */
    public Output<Optional<Boolean>> ignoreMetadata() {
        return Codegen.optional(this.ignoreMetadata);
    }
    /**
     * Name of Tenant
     * 
     */
    @Export(name="tenantList", refs={String.class}, tree="[0]")
    private Output<String> tenantList;

    /**
     * @return Name of Tenant
     * 
     */
    public Output<String> tenantList() {
        return this.tenantList;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BigIqAs3(String name) {
        this(name, BigIqAs3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BigIqAs3(String name, BigIqAs3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BigIqAs3(String name, BigIqAs3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/bigIqAs3:BigIqAs3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private BigIqAs3(String name, Output<String> id, @Nullable BigIqAs3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/bigIqAs3:BigIqAs3", name, state, makeResourceOptions(options, id));
    }

    private static BigIqAs3Args makeArgs(BigIqAs3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BigIqAs3Args.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "bigiqLoginRef",
                "bigiqPassword",
                "bigiqPort",
                "bigiqTokenAuth",
                "bigiqUser"
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
    public static BigIqAs3 get(String name, Output<String> id, @Nullable BigIqAs3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BigIqAs3(name, id, state, options);
    }
}
