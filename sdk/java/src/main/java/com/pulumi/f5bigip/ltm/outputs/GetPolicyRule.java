// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.f5bigip.ltm.outputs.GetPolicyRuleAction;
import com.pulumi.f5bigip.ltm.outputs.GetPolicyRuleCondition;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetPolicyRule {
    private @Nullable List<GetPolicyRuleAction> actions;
    private @Nullable List<GetPolicyRuleCondition> conditions;
    /**
     * @return Name of the policy which includes partion ( /partition/policy-name )
     * 
     */
    private String name;

    private GetPolicyRule() {}
    public List<GetPolicyRuleAction> actions() {
        return this.actions == null ? List.of() : this.actions;
    }
    public List<GetPolicyRuleCondition> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    /**
     * @return Name of the policy which includes partion ( /partition/policy-name )
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPolicyRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetPolicyRuleAction> actions;
        private @Nullable List<GetPolicyRuleCondition> conditions;
        private String name;
        public Builder() {}
        public Builder(GetPolicyRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.conditions = defaults.conditions;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder actions(@Nullable List<GetPolicyRuleAction> actions) {
            this.actions = actions;
            return this;
        }
        public Builder actions(GetPolicyRuleAction... actions) {
            return actions(List.of(actions));
        }
        @CustomType.Setter
        public Builder conditions(@Nullable List<GetPolicyRuleCondition> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(GetPolicyRuleCondition... conditions) {
            return conditions(List.of(conditions));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetPolicyRule build() {
            final var o = new GetPolicyRule();
            o.actions = actions;
            o.conditions = conditions;
            o.name = name;
            return o;
        }
    }
}