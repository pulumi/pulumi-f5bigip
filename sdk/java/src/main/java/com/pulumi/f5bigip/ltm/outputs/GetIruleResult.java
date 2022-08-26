// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

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
    private String irule;
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
    public String irule() {
        return this.irule;
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
        private String irule;
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
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder irule(String irule) {
            this.irule = Objects.requireNonNull(irule);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder partition(String partition) {
            this.partition = Objects.requireNonNull(partition);
            return this;
        }
        public GetIruleResult build() {
            final var o = new GetIruleResult();
            o.id = id;
            o.irule = irule;
            o.name = name;
            o.partition = partition;
            return o;
        }
    }
}
