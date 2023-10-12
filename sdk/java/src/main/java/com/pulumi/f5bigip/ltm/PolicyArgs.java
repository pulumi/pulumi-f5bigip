// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.ltm.inputs.PolicyRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyArgs Empty = new PolicyArgs();

    /**
     * Specifies the controls
     * 
     */
    @Import(name="controls")
    private @Nullable Output<List<String>> controls;

    /**
     * @return Specifies the controls
     * 
     */
    public Optional<Output<List<String>>> controls() {
        return Optional.ofNullable(this.controls);
    }

    /**
     * Specifies descriptive text that identifies the irule attached to policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Specifies descriptive text that identifies the irule attached to policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of Rule to be applied in policy.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of Rule to be applied in policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
     * 
     * @deprecated
     * This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
     * 
     */
    @Deprecated /* This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release. */
    @Import(name="publishedCopy")
    private @Nullable Output<String> publishedCopy;

    /**
     * @return If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
     * 
     * @deprecated
     * This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
     * 
     */
    @Deprecated /* This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release. */
    public Optional<Output<String>> publishedCopy() {
        return Optional.ofNullable(this.publishedCopy);
    }

    /**
     * Specifies the protocol
     * 
     */
    @Import(name="requires")
    private @Nullable Output<List<String>> requires;

    /**
     * @return Specifies the protocol
     * 
     */
    public Optional<Output<List<String>>> requires() {
        return Optional.ofNullable(this.requires);
    }

    /**
     * List of Rules can be applied using the policy. Each rule is block type with following arguments.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<PolicyRuleArgs>> rules;

    /**
     * @return List of Rules can be applied using the policy. Each rule is block type with following arguments.
     * 
     */
    public Optional<Output<List<PolicyRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Specifies the match strategy
     * 
     */
    @Import(name="strategy")
    private @Nullable Output<String> strategy;

    /**
     * @return Specifies the match strategy
     * 
     */
    public Optional<Output<String>> strategy() {
        return Optional.ofNullable(this.strategy);
    }

    private PolicyArgs() {}

    private PolicyArgs(PolicyArgs $) {
        this.controls = $.controls;
        this.description = $.description;
        this.name = $.name;
        this.publishedCopy = $.publishedCopy;
        this.requires = $.requires;
        this.rules = $.rules;
        this.strategy = $.strategy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyArgs $;

        public Builder() {
            $ = new PolicyArgs();
        }

        public Builder(PolicyArgs defaults) {
            $ = new PolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param controls Specifies the controls
         * 
         * @return builder
         * 
         */
        public Builder controls(@Nullable Output<List<String>> controls) {
            $.controls = controls;
            return this;
        }

        /**
         * @param controls Specifies the controls
         * 
         * @return builder
         * 
         */
        public Builder controls(List<String> controls) {
            return controls(Output.of(controls));
        }

        /**
         * @param controls Specifies the controls
         * 
         * @return builder
         * 
         */
        public Builder controls(String... controls) {
            return controls(List.of(controls));
        }

        /**
         * @param description Specifies descriptive text that identifies the irule attached to policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Specifies descriptive text that identifies the irule attached to policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of Rule to be applied in policy.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of Rule to be applied in policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param publishedCopy If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
         * 
         * @return builder
         * 
         * @deprecated
         * This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
         * 
         */
        @Deprecated /* This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release. */
        public Builder publishedCopy(@Nullable Output<String> publishedCopy) {
            $.publishedCopy = publishedCopy;
            return this;
        }

        /**
         * @param publishedCopy If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
         * 
         * @return builder
         * 
         * @deprecated
         * This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
         * 
         */
        @Deprecated /* This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release. */
        public Builder publishedCopy(String publishedCopy) {
            return publishedCopy(Output.of(publishedCopy));
        }

        /**
         * @param requires Specifies the protocol
         * 
         * @return builder
         * 
         */
        public Builder requires(@Nullable Output<List<String>> requires) {
            $.requires = requires;
            return this;
        }

        /**
         * @param requires Specifies the protocol
         * 
         * @return builder
         * 
         */
        public Builder requires(List<String> requires) {
            return requires(Output.of(requires));
        }

        /**
         * @param requires Specifies the protocol
         * 
         * @return builder
         * 
         */
        public Builder requires(String... requires) {
            return requires(List.of(requires));
        }

        /**
         * @param rules List of Rules can be applied using the policy. Each rule is block type with following arguments.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<PolicyRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules List of Rules can be applied using the policy. Each rule is block type with following arguments.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<PolicyRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules List of Rules can be applied using the policy. Each rule is block type with following arguments.
         * 
         * @return builder
         * 
         */
        public Builder rules(PolicyRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param strategy Specifies the match strategy
         * 
         * @return builder
         * 
         */
        public Builder strategy(@Nullable Output<String> strategy) {
            $.strategy = strategy;
            return this;
        }

        /**
         * @param strategy Specifies the match strategy
         * 
         * @return builder
         * 
         */
        public Builder strategy(String strategy) {
            return strategy(Output.of(strategy));
        }

        public PolicyArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
