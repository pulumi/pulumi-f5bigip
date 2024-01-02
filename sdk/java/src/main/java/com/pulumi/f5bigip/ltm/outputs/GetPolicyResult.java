// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ltm.outputs.GetPolicyRule;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPolicyResult {
    /**
     * @return Specifies the controls.
     * 
     */
    private @Nullable List<String> controls;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the policy.
     * 
     */
    private String name;
    private @Nullable String publishedCopy;
    /**
     * @return Specifies the protocol.
     * 
     */
    private @Nullable List<String> requires;
    /**
     * @return Rules defined in the policy.
     * 
     */
    private @Nullable List<GetPolicyRule> rules;
    /**
     * @return Specifies the match strategy.
     * 
     */
    private @Nullable String strategy;

    private GetPolicyResult() {}
    /**
     * @return Specifies the controls.
     * 
     */
    public List<String> controls() {
        return this.controls == null ? List.of() : this.controls;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the policy.
     * 
     */
    public String name() {
        return this.name;
    }
    public Optional<String> publishedCopy() {
        return Optional.ofNullable(this.publishedCopy);
    }
    /**
     * @return Specifies the protocol.
     * 
     */
    public List<String> requires() {
        return this.requires == null ? List.of() : this.requires;
    }
    /**
     * @return Rules defined in the policy.
     * 
     */
    public List<GetPolicyRule> rules() {
        return this.rules == null ? List.of() : this.rules;
    }
    /**
     * @return Specifies the match strategy.
     * 
     */
    public Optional<String> strategy() {
        return Optional.ofNullable(this.strategy);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPolicyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> controls;
        private String id;
        private String name;
        private @Nullable String publishedCopy;
        private @Nullable List<String> requires;
        private @Nullable List<GetPolicyRule> rules;
        private @Nullable String strategy;
        public Builder() {}
        public Builder(GetPolicyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.controls = defaults.controls;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.publishedCopy = defaults.publishedCopy;
    	      this.requires = defaults.requires;
    	      this.rules = defaults.rules;
    	      this.strategy = defaults.strategy;
        }

        @CustomType.Setter
        public Builder controls(@Nullable List<String> controls) {

            this.controls = controls;
            return this;
        }
        public Builder controls(String... controls) {
            return controls(List.of(controls));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPolicyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPolicyResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder publishedCopy(@Nullable String publishedCopy) {

            this.publishedCopy = publishedCopy;
            return this;
        }
        @CustomType.Setter
        public Builder requires(@Nullable List<String> requires) {

            this.requires = requires;
            return this;
        }
        public Builder requires(String... requires) {
            return requires(List.of(requires));
        }
        @CustomType.Setter
        public Builder rules(@Nullable List<GetPolicyRule> rules) {

            this.rules = rules;
            return this;
        }
        public Builder rules(GetPolicyRule... rules) {
            return rules(List.of(rules));
        }
        @CustomType.Setter
        public Builder strategy(@Nullable String strategy) {

            this.strategy = strategy;
            return this;
        }
        public GetPolicyResult build() {
            final var _resultValue = new GetPolicyResult();
            _resultValue.controls = controls;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.publishedCopy = publishedCopy;
            _resultValue.requires = requires;
            _resultValue.rules = rules;
            _resultValue.strategy = strategy;
            return _resultValue;
        }
    }
}
