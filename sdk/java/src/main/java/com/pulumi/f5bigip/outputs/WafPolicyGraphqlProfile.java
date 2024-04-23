// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.outputs.WafPolicyGraphqlProfileDefenseAttribute;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WafPolicyGraphqlProfile {
    /**
     * @return Specifies when checked (enabled) that you want attack signatures and threat campaigns to be detected on this GraphQL profile and possibly override the security policy settings of an attack signature or threat campaign specifically for this GraphQL profile. After you enable this setting, the system displays a list of attack signatures and and threat campaigns. The default is enabled
     * 
     */
    private @Nullable Boolean attackSignaturesCheck;
    /**
     * @return defense_attributes settings for policy
     * 
     */
    private @Nullable List<WafPolicyGraphqlProfileDefenseAttribute> defenseAttributes;
    /**
     * @return Specifies when checked (enabled) that the system enforces the security policy settings of a meta character for the GraphQL profile. After you enable this setting, the system displays a list of meta characters. The default is enabled
     * 
     */
    private @Nullable Boolean metacharElementcheck;
    /**
     * @return The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     * 
     */
    private String name;

    private WafPolicyGraphqlProfile() {}
    /**
     * @return Specifies when checked (enabled) that you want attack signatures and threat campaigns to be detected on this GraphQL profile and possibly override the security policy settings of an attack signature or threat campaign specifically for this GraphQL profile. After you enable this setting, the system displays a list of attack signatures and and threat campaigns. The default is enabled
     * 
     */
    public Optional<Boolean> attackSignaturesCheck() {
        return Optional.ofNullable(this.attackSignaturesCheck);
    }
    /**
     * @return defense_attributes settings for policy
     * 
     */
    public List<WafPolicyGraphqlProfileDefenseAttribute> defenseAttributes() {
        return this.defenseAttributes == null ? List.of() : this.defenseAttributes;
    }
    /**
     * @return Specifies when checked (enabled) that the system enforces the security policy settings of a meta character for the GraphQL profile. After you enable this setting, the system displays a list of meta characters. The default is enabled
     * 
     */
    public Optional<Boolean> metacharElementcheck() {
        return Optional.ofNullable(this.metacharElementcheck);
    }
    /**
     * @return The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WafPolicyGraphqlProfile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean attackSignaturesCheck;
        private @Nullable List<WafPolicyGraphqlProfileDefenseAttribute> defenseAttributes;
        private @Nullable Boolean metacharElementcheck;
        private String name;
        public Builder() {}
        public Builder(WafPolicyGraphqlProfile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attackSignaturesCheck = defaults.attackSignaturesCheck;
    	      this.defenseAttributes = defaults.defenseAttributes;
    	      this.metacharElementcheck = defaults.metacharElementcheck;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder attackSignaturesCheck(@Nullable Boolean attackSignaturesCheck) {

            this.attackSignaturesCheck = attackSignaturesCheck;
            return this;
        }
        @CustomType.Setter
        public Builder defenseAttributes(@Nullable List<WafPolicyGraphqlProfileDefenseAttribute> defenseAttributes) {

            this.defenseAttributes = defenseAttributes;
            return this;
        }
        public Builder defenseAttributes(WafPolicyGraphqlProfileDefenseAttribute... defenseAttributes) {
            return defenseAttributes(List.of(defenseAttributes));
        }
        @CustomType.Setter
        public Builder metacharElementcheck(@Nullable Boolean metacharElementcheck) {

            this.metacharElementcheck = metacharElementcheck;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("WafPolicyGraphqlProfile", "name");
            }
            this.name = name;
            return this;
        }
        public WafPolicyGraphqlProfile build() {
            final var _resultValue = new WafPolicyGraphqlProfile();
            _resultValue.attackSignaturesCheck = attackSignaturesCheck;
            _resultValue.defenseAttributes = defenseAttributes;
            _resultValue.metacharElementcheck = metacharElementcheck;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
