// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.sys.DnsArgs;
import com.pulumi.f5bigip.sys.inputs.DnsState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.sys.Dns` Configures DNS Name server on F5 BIG-IP
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
 * import com.pulumi.f5bigip.sys.Dns;
 * import com.pulumi.f5bigip.sys.DnsArgs;
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
 *         var dns1 = new Dns("dns1", DnsArgs.builder()
 *             .description("/Common/DNS1")
 *             .nameServers("1.1.1.1")
 *             .searches("f5.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:sys/dns:Dns")
public class Dns extends com.pulumi.resources.CustomResource {
    /**
     * Provide description for your DNS server
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Provide description for your DNS server
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Specifies the name servers that the system uses to validate DNS lookups, and resolve host names.
     * 
     */
    @Export(name="nameServers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nameServers;

    /**
     * @return Specifies the name servers that the system uses to validate DNS lookups, and resolve host names.
     * 
     */
    public Output<List<String>> nameServers() {
        return this.nameServers;
    }
    /**
     * Configures the number of dots needed in a name before an initial absolute query will be made.
     * 
     */
    @Export(name="numberOfDots", refs={Integer.class}, tree="[0]")
    private Output<Integer> numberOfDots;

    /**
     * @return Configures the number of dots needed in a name before an initial absolute query will be made.
     * 
     */
    public Output<Integer> numberOfDots() {
        return this.numberOfDots;
    }
    /**
     * Specifies the domains that the system searches for local domain lookups, to resolve local host names.
     * 
     */
    @Export(name="searches", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> searches;

    /**
     * @return Specifies the domains that the system searches for local domain lookups, to resolve local host names.
     * 
     */
    public Output<Optional<List<String>>> searches() {
        return Codegen.optional(this.searches);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dns(String name) {
        this(name, DnsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dns(String name, DnsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dns(String name, DnsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/dns:Dns", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Dns(String name, Output<String> id, @Nullable DnsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/dns:Dns", name, state, makeResourceOptions(options, id));
    }

    private static DnsArgs makeArgs(DnsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DnsArgs.Empty : args;
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
    public static Dns get(String name, Output<String> id, @Nullable DnsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dns(name, id, state, options);
    }
}
