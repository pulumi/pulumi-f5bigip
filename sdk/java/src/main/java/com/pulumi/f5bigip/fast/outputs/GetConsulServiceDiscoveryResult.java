// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetConsulServiceDiscoveryResult {
    private @Nullable String addressRealm;
    /**
     * @return The JSON for Hashicorp Consul service discovery block.
     * 
     */
    private String consulSdJson;
    private @Nullable Boolean credentialUpdate;
    private @Nullable String encodedToken;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String jmesPathQuery;
    private @Nullable String minimumMonitors;
    private Integer port;
    private @Nullable Boolean rejectUnauthorized;
    private @Nullable String trustCa;
    private @Nullable String type;
    private @Nullable String undetectableAction;
    private @Nullable String updateInterval;
    private String uri;

    private GetConsulServiceDiscoveryResult() {}
    public Optional<String> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }
    /**
     * @return The JSON for Hashicorp Consul service discovery block.
     * 
     */
    public String consulSdJson() {
        return this.consulSdJson;
    }
    public Optional<Boolean> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }
    public Optional<String> encodedToken() {
        return Optional.ofNullable(this.encodedToken);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> jmesPathQuery() {
        return Optional.ofNullable(this.jmesPathQuery);
    }
    public Optional<String> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }
    public Integer port() {
        return this.port;
    }
    public Optional<Boolean> rejectUnauthorized() {
        return Optional.ofNullable(this.rejectUnauthorized);
    }
    public Optional<String> trustCa() {
        return Optional.ofNullable(this.trustCa);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    public Optional<String> undetectableAction() {
        return Optional.ofNullable(this.undetectableAction);
    }
    public Optional<String> updateInterval() {
        return Optional.ofNullable(this.updateInterval);
    }
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConsulServiceDiscoveryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String addressRealm;
        private String consulSdJson;
        private @Nullable Boolean credentialUpdate;
        private @Nullable String encodedToken;
        private String id;
        private @Nullable String jmesPathQuery;
        private @Nullable String minimumMonitors;
        private Integer port;
        private @Nullable Boolean rejectUnauthorized;
        private @Nullable String trustCa;
        private @Nullable String type;
        private @Nullable String undetectableAction;
        private @Nullable String updateInterval;
        private String uri;
        public Builder() {}
        public Builder(GetConsulServiceDiscoveryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressRealm = defaults.addressRealm;
    	      this.consulSdJson = defaults.consulSdJson;
    	      this.credentialUpdate = defaults.credentialUpdate;
    	      this.encodedToken = defaults.encodedToken;
    	      this.id = defaults.id;
    	      this.jmesPathQuery = defaults.jmesPathQuery;
    	      this.minimumMonitors = defaults.minimumMonitors;
    	      this.port = defaults.port;
    	      this.rejectUnauthorized = defaults.rejectUnauthorized;
    	      this.trustCa = defaults.trustCa;
    	      this.type = defaults.type;
    	      this.undetectableAction = defaults.undetectableAction;
    	      this.updateInterval = defaults.updateInterval;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder addressRealm(@Nullable String addressRealm) {
            this.addressRealm = addressRealm;
            return this;
        }
        @CustomType.Setter
        public Builder consulSdJson(String consulSdJson) {
            this.consulSdJson = Objects.requireNonNull(consulSdJson);
            return this;
        }
        @CustomType.Setter
        public Builder credentialUpdate(@Nullable Boolean credentialUpdate) {
            this.credentialUpdate = credentialUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder encodedToken(@Nullable String encodedToken) {
            this.encodedToken = encodedToken;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder jmesPathQuery(@Nullable String jmesPathQuery) {
            this.jmesPathQuery = jmesPathQuery;
            return this;
        }
        @CustomType.Setter
        public Builder minimumMonitors(@Nullable String minimumMonitors) {
            this.minimumMonitors = minimumMonitors;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder rejectUnauthorized(@Nullable Boolean rejectUnauthorized) {
            this.rejectUnauthorized = rejectUnauthorized;
            return this;
        }
        @CustomType.Setter
        public Builder trustCa(@Nullable String trustCa) {
            this.trustCa = trustCa;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder undetectableAction(@Nullable String undetectableAction) {
            this.undetectableAction = undetectableAction;
            return this;
        }
        @CustomType.Setter
        public Builder updateInterval(@Nullable String updateInterval) {
            this.updateInterval = updateInterval;
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public GetConsulServiceDiscoveryResult build() {
            final var o = new GetConsulServiceDiscoveryResult();
            o.addressRealm = addressRealm;
            o.consulSdJson = consulSdJson;
            o.credentialUpdate = credentialUpdate;
            o.encodedToken = encodedToken;
            o.id = id;
            o.jmesPathQuery = jmesPathQuery;
            o.minimumMonitors = minimumMonitors;
            o.port = port;
            o.rejectUnauthorized = rejectUnauthorized;
            o.trustCa = trustCa;
            o.type = type;
            o.undetectableAction = undetectableAction;
            o.updateInterval = updateInterval;
            o.uri = uri;
            return o;
        }
    }
}
