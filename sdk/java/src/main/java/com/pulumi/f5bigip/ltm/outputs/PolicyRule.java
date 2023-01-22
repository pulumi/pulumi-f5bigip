// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.f5bigip.ltm.outputs.PolicyRuleAction;
import com.pulumi.f5bigip.ltm.outputs.PolicyRuleCondition;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class PolicyRule {
    /**
     * @return Block type. See action block for more details.
     * 
     */
    private @Nullable List<PolicyRuleAction> actions;
    /**
     * @return Block type. See condition block for more details.
     * 
     */
    private @Nullable List<PolicyRuleCondition> conditions;
    /**
     * @return Name of Rule to be applied in policy.
     * 
     */
    private String name;

    private PolicyRule() {}
    /**
     * @return Block type. See action block for more details.
     * 
     */
    public List<PolicyRuleAction> actions() {
        return this.actions == null ? List.of() : this.actions;
    }
    /**
     * @return Block type. See condition block for more details.
     * 
     */
    public List<PolicyRuleCondition> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    /**
     * @return Name of Rule to be applied in policy.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicyRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<PolicyRuleAction> actions;
        private @Nullable List<PolicyRuleCondition> conditions;
        private String name;
        public Builder() {}
        public Builder(PolicyRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.conditions = defaults.conditions;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder actions(@Nullable List<PolicyRuleAction> actions) {
            this.actions = actions;
            return this;
        }
        public Builder actions(PolicyRuleAction... actions) {
            return actions(List.of(actions));
        }
        @CustomType.Setter
        public Builder conditions(@Nullable List<PolicyRuleCondition> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(PolicyRuleCondition... conditions) {
            return conditions(List.of(conditions));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public PolicyRule build() {
            final var o = new PolicyRule();
            o.actions = actions;
            o.conditions = conditions;
            o.name = name;
            return o;
        }
    }
}
