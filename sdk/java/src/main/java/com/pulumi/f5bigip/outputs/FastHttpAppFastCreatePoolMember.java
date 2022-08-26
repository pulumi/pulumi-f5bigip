// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FastHttpAppFastCreatePoolMember {
    /**
     * @return List of server address to be used for FAST-Generated Pool.
     * 
     */
    private List<String> addresses;
    /**
     * @return connectionLimit value to be used for FAST-Generated Pool.
     * 
     */
    private @Nullable Integer connectionLimit;
    /**
     * @return port number of serviceport to be used for FAST-Generated Pool.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return priorityGroup value to be used for FAST-Generated Pool.
     * 
     */
    private @Nullable Integer priorityGroup;
    /**
     * @return shareNodes value to be used for FAST-Generated Pool.
     * 
     */
    private @Nullable Boolean shareNodes;

    private FastHttpAppFastCreatePoolMember() {}
    /**
     * @return List of server address to be used for FAST-Generated Pool.
     * 
     */
    public List<String> addresses() {
        return this.addresses;
    }
    /**
     * @return connectionLimit value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Integer> connectionLimit() {
        return Optional.ofNullable(this.connectionLimit);
    }
    /**
     * @return port number of serviceport to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return priorityGroup value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Integer> priorityGroup() {
        return Optional.ofNullable(this.priorityGroup);
    }
    /**
     * @return shareNodes value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Boolean> shareNodes() {
        return Optional.ofNullable(this.shareNodes);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FastHttpAppFastCreatePoolMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> addresses;
        private @Nullable Integer connectionLimit;
        private @Nullable Integer port;
        private @Nullable Integer priorityGroup;
        private @Nullable Boolean shareNodes;
        public Builder() {}
        public Builder(FastHttpAppFastCreatePoolMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addresses = defaults.addresses;
    	      this.connectionLimit = defaults.connectionLimit;
    	      this.port = defaults.port;
    	      this.priorityGroup = defaults.priorityGroup;
    	      this.shareNodes = defaults.shareNodes;
        }

        @CustomType.Setter
        public Builder addresses(List<String> addresses) {
            this.addresses = Objects.requireNonNull(addresses);
            return this;
        }
        public Builder addresses(String... addresses) {
            return addresses(List.of(addresses));
        }
        @CustomType.Setter
        public Builder connectionLimit(@Nullable Integer connectionLimit) {
            this.connectionLimit = connectionLimit;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder priorityGroup(@Nullable Integer priorityGroup) {
            this.priorityGroup = priorityGroup;
            return this;
        }
        @CustomType.Setter
        public Builder shareNodes(@Nullable Boolean shareNodes) {
            this.shareNodes = shareNodes;
            return this;
        }
        public FastHttpAppFastCreatePoolMember build() {
            final var o = new FastHttpAppFastCreatePoolMember();
            o.addresses = addresses;
            o.connectionLimit = connectionLimit;
            o.port = port;
            o.priorityGroup = priorityGroup;
            o.shareNodes = shareNodes;
            return o;
        }
    }
}
