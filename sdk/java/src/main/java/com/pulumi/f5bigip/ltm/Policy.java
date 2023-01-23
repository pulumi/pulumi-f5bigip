// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PolicyArgs;
import com.pulumi.f5bigip.ltm.inputs.PolicyState;
import com.pulumi.f5bigip.ltm.outputs.PolicyRule;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
 * 
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-policy`.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.Pool;
 * import com.pulumi.f5bigip.ltm.PoolArgs;
 * import com.pulumi.f5bigip.ltm.Policy;
 * import com.pulumi.f5bigip.ltm.PolicyArgs;
 * import com.pulumi.f5bigip.ltm.inputs.PolicyRuleArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var mypool = new Pool(&#34;mypool&#34;, PoolArgs.builder()        
 *             .name(&#34;/Common/test-pool&#34;)
 *             .allowNat(&#34;yes&#34;)
 *             .allowSnat(&#34;yes&#34;)
 *             .loadBalancingMode(&#34;round-robin&#34;)
 *             .build());
 * 
 *         var test_policy = new Policy(&#34;test-policy&#34;, PolicyArgs.builder()        
 *             .name(&#34;/Common/test-policy&#34;)
 *             .strategy(&#34;first-match&#34;)
 *             .requires(&#34;http&#34;)
 *             .controls(&#34;forwarding&#34;)
 *             .rules(PolicyRuleArgs.builder()
 *                 .name(&#34;rule6&#34;)
 *                 .actions(PolicyRuleActionArgs.builder()
 *                     .forward(true)
 *                     .connection(false)
 *                     .pool(mypool.name())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(mypool)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the controls
     * 
     */
    @Export(name="controls", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> controls;

    /**
     * @return Specifies the controls
     * 
     */
    public Output<Optional<List<String>>> controls() {
        return Codegen.optional(this.controls);
    }
    /**
     * Name of Rule to be applied in policy.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of Rule to be applied in policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode.
     * 
     */
    @Export(name="publishedCopy", type=String.class, parameters={})
    private Output</* @Nullable */ String> publishedCopy;

    /**
     * @return If you want to publish the policy else it will be deployed in Drafts mode.
     * 
     */
    public Output<Optional<String>> publishedCopy() {
        return Codegen.optional(this.publishedCopy);
    }
    /**
     * Specifies the protocol
     * 
     */
    @Export(name="requires", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> requires;

    /**
     * @return Specifies the protocol
     * 
     */
    public Output<Optional<List<String>>> requires() {
        return Codegen.optional(this.requires);
    }
    /**
     * List of Rules can be applied using the policy. Each rule is block type with following arguments.
     * 
     */
    @Export(name="rules", type=List.class, parameters={PolicyRule.class})
    private Output</* @Nullable */ List<PolicyRule>> rules;

    /**
     * @return List of Rules can be applied using the policy. Each rule is block type with following arguments.
     * 
     */
    public Output<Optional<List<PolicyRule>>> rules() {
        return Codegen.optional(this.rules);
    }
    /**
     * Specifies the match strategy
     * 
     */
    @Export(name="strategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> strategy;

    /**
     * @return Specifies the match strategy
     * 
     */
    public Output<Optional<String>> strategy() {
        return Codegen.optional(this.strategy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/policy:Policy", name, state, makeResourceOptions(options, id));
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
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}