// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ltm.outputs.GetNodeFqdn;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetNodeResult {
    /**
     * @return The address of the node.
     * 
     */
    private @Nullable String address;
    /**
     * @return Node connection limit.
     * 
     */
    private Integer connectionLimit;
    /**
     * @return User defined description of the node.
     * 
     */
    private @Nullable String description;
    /**
     * @return The dynamic ratio number for the node.
     * 
     */
    private Integer dynamicRatio;
    private GetNodeFqdn fqdn;
    /**
     * @return Full path of the node (partition and name)
     * 
     */
    private @Nullable String fullPath;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Specifies the health monitors the system currently uses to monitor this node.
     * 
     */
    private String monitor;
    private String name;
    private String partition;
    /**
     * @return Node rate limit.
     * 
     */
    private String rateLimit;
    /**
     * @return Node ratio weight.
     * 
     */
    private Integer ratio;
    private String session;
    /**
     * @return The current state of the node.
     * 
     */
    private String state;

    private GetNodeResult() {}
    /**
     * @return The address of the node.
     * 
     */
    public Optional<String> address() {
        return Optional.ofNullable(this.address);
    }
    /**
     * @return Node connection limit.
     * 
     */
    public Integer connectionLimit() {
        return this.connectionLimit;
    }
    /**
     * @return User defined description of the node.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The dynamic ratio number for the node.
     * 
     */
    public Integer dynamicRatio() {
        return this.dynamicRatio;
    }
    public GetNodeFqdn fqdn() {
        return this.fqdn;
    }
    /**
     * @return Full path of the node (partition and name)
     * 
     */
    public Optional<String> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Specifies the health monitors the system currently uses to monitor this node.
     * 
     */
    public String monitor() {
        return this.monitor;
    }
    public String name() {
        return this.name;
    }
    public String partition() {
        return this.partition;
    }
    /**
     * @return Node rate limit.
     * 
     */
    public String rateLimit() {
        return this.rateLimit;
    }
    /**
     * @return Node ratio weight.
     * 
     */
    public Integer ratio() {
        return this.ratio;
    }
    public String session() {
        return this.session;
    }
    /**
     * @return The current state of the node.
     * 
     */
    public String state() {
        return this.state;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNodeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String address;
        private Integer connectionLimit;
        private @Nullable String description;
        private Integer dynamicRatio;
        private GetNodeFqdn fqdn;
        private @Nullable String fullPath;
        private String id;
        private String monitor;
        private String name;
        private String partition;
        private String rateLimit;
        private Integer ratio;
        private String session;
        private String state;
        public Builder() {}
        public Builder(GetNodeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.connectionLimit = defaults.connectionLimit;
    	      this.description = defaults.description;
    	      this.dynamicRatio = defaults.dynamicRatio;
    	      this.fqdn = defaults.fqdn;
    	      this.fullPath = defaults.fullPath;
    	      this.id = defaults.id;
    	      this.monitor = defaults.monitor;
    	      this.name = defaults.name;
    	      this.partition = defaults.partition;
    	      this.rateLimit = defaults.rateLimit;
    	      this.ratio = defaults.ratio;
    	      this.session = defaults.session;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder address(@Nullable String address) {

            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder connectionLimit(Integer connectionLimit) {
            if (connectionLimit == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "connectionLimit");
            }
            this.connectionLimit = connectionLimit;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder dynamicRatio(Integer dynamicRatio) {
            if (dynamicRatio == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "dynamicRatio");
            }
            this.dynamicRatio = dynamicRatio;
            return this;
        }
        @CustomType.Setter
        public Builder fqdn(GetNodeFqdn fqdn) {
            if (fqdn == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "fqdn");
            }
            this.fqdn = fqdn;
            return this;
        }
        @CustomType.Setter
        public Builder fullPath(@Nullable String fullPath) {

            this.fullPath = fullPath;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder monitor(String monitor) {
            if (monitor == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "monitor");
            }
            this.monitor = monitor;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder partition(String partition) {
            if (partition == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "partition");
            }
            this.partition = partition;
            return this;
        }
        @CustomType.Setter
        public Builder rateLimit(String rateLimit) {
            if (rateLimit == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "rateLimit");
            }
            this.rateLimit = rateLimit;
            return this;
        }
        @CustomType.Setter
        public Builder ratio(Integer ratio) {
            if (ratio == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "ratio");
            }
            this.ratio = ratio;
            return this;
        }
        @CustomType.Setter
        public Builder session(String session) {
            if (session == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "session");
            }
            this.session = session;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetNodeResult", "state");
            }
            this.state = state;
            return this;
        }
        public GetNodeResult build() {
            final var _resultValue = new GetNodeResult();
            _resultValue.address = address;
            _resultValue.connectionLimit = connectionLimit;
            _resultValue.description = description;
            _resultValue.dynamicRatio = dynamicRatio;
            _resultValue.fqdn = fqdn;
            _resultValue.fullPath = fullPath;
            _resultValue.id = id;
            _resultValue.monitor = monitor;
            _resultValue.name = name;
            _resultValue.partition = partition;
            _resultValue.rateLimit = rateLimit;
            _resultValue.ratio = ratio;
            _resultValue.session = session;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}
