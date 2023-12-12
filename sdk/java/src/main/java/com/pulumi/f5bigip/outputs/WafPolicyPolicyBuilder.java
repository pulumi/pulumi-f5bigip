// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WafPolicyPolicyBuilder {
    /**
     * @return learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
     * 
     */
    private @Nullable String learningMode;

    private WafPolicyPolicyBuilder() {}
    /**
     * @return learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
     * 
     */
    public Optional<String> learningMode() {
        return Optional.ofNullable(this.learningMode);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WafPolicyPolicyBuilder defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String learningMode;
        public Builder() {}
        public Builder(WafPolicyPolicyBuilder defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.learningMode = defaults.learningMode;
        }

        @CustomType.Setter
        public Builder learningMode(@Nullable String learningMode) {
            this.learningMode = learningMode;
            return this;
        }
        public WafPolicyPolicyBuilder build() {
            final var _resultValue = new WafPolicyPolicyBuilder();
            _resultValue.learningMode = learningMode;
            return _resultValue;
        }
    }
}
