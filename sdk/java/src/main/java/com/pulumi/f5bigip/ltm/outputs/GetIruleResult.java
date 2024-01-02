// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIruleResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Irule configured on bigip
     * 
     */
    private @Nullable String irule;
    /**
     * @return Name of irule configured on bigip with full path
     * 
     */
    private String name;
    /**
     * @return Bigip partition in which rule is configured
     * 
     */
    private String partition;

    private GetIruleResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Irule configured on bigip
     * 
     */
    public Optional<String> irule() {
        return Optional.ofNullable(this.irule);
    }
    /**
     * @return Name of irule configured on bigip with full path
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Bigip partition in which rule is configured
     * 
     */
    public String partition() {
        return this.partition;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIruleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String irule;
        private String name;
        private String partition;
        public Builder() {}
        public Builder(GetIruleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.irule = defaults.irule;
    	      this.name = defaults.name;
    	      this.partition = defaults.partition;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIruleResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder irule(@Nullable String irule) {

            this.irule = irule;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIruleResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder partition(String partition) {
            if (partition == null) {
              throw new MissingRequiredPropertyException("GetIruleResult", "partition");
            }
            this.partition = partition;
            return this;
        }
        public GetIruleResult build() {
            final var _resultValue = new GetIruleResult();
            _resultValue.id = id;
            _resultValue.irule = irule;
            _resultValue.name = name;
            _resultValue.partition = partition;
            return _resultValue;
        }
    }
}
