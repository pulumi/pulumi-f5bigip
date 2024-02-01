// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EventServiceDiscoveryNode {
    /**
     * @return name of node
     * 
     */
    private @Nullable String id;
    /**
     * @return ip of nonde
     * 
     */
    private @Nullable String ip;
    /**
     * @return port
     * 
     */
    private @Nullable Integer port;

    private EventServiceDiscoveryNode() {}
    /**
     * @return name of node
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return ip of nonde
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return port
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EventServiceDiscoveryNode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        private @Nullable String ip;
        private @Nullable Integer port;
        public Builder() {}
        public Builder(EventServiceDiscoveryNode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {

            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {

            this.port = port;
            return this;
        }
        public EventServiceDiscoveryNode build() {
            final var _resultValue = new EventServiceDiscoveryNode();
            _resultValue.id = id;
            _resultValue.ip = ip;
            _resultValue.port = port;
            return _resultValue;
        }
    }
}
