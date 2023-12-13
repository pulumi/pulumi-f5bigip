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
public final class GetGceServiceDiscoveryResult {
    private @Nullable String addressRealm;
    private @Nullable Boolean credentialUpdate;
    private @Nullable String encodedCredentials;
    /**
     * @return The JSON for GCE service discovery block.
     * 
     */
    private String gceSdJson;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String minimumMonitors;
    private @Nullable Integer port;
    private @Nullable String projectId;
    private String region;
    private String tagKey;
    private String tagValue;
    private @Nullable String type;
    private @Nullable String undetectableAction;
    private @Nullable String updateInterval;

    private GetGceServiceDiscoveryResult() {}
    public Optional<String> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }
    public Optional<Boolean> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }
    public Optional<String> encodedCredentials() {
        return Optional.ofNullable(this.encodedCredentials);
    }
    /**
     * @return The JSON for GCE service discovery block.
     * 
     */
    public String gceSdJson() {
        return this.gceSdJson;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public String region() {
        return this.region;
    }
    public String tagKey() {
        return this.tagKey;
    }
    public String tagValue() {
        return this.tagValue;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGceServiceDiscoveryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String addressRealm;
        private @Nullable Boolean credentialUpdate;
        private @Nullable String encodedCredentials;
        private String gceSdJson;
        private String id;
        private @Nullable String minimumMonitors;
        private @Nullable Integer port;
        private @Nullable String projectId;
        private String region;
        private String tagKey;
        private String tagValue;
        private @Nullable String type;
        private @Nullable String undetectableAction;
        private @Nullable String updateInterval;
        public Builder() {}
        public Builder(GetGceServiceDiscoveryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressRealm = defaults.addressRealm;
    	      this.credentialUpdate = defaults.credentialUpdate;
    	      this.encodedCredentials = defaults.encodedCredentials;
    	      this.gceSdJson = defaults.gceSdJson;
    	      this.id = defaults.id;
    	      this.minimumMonitors = defaults.minimumMonitors;
    	      this.port = defaults.port;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.tagKey = defaults.tagKey;
    	      this.tagValue = defaults.tagValue;
    	      this.type = defaults.type;
    	      this.undetectableAction = defaults.undetectableAction;
    	      this.updateInterval = defaults.updateInterval;
        }

        @CustomType.Setter
        public Builder addressRealm(@Nullable String addressRealm) {
            this.addressRealm = addressRealm;
            return this;
        }
        @CustomType.Setter
        public Builder credentialUpdate(@Nullable Boolean credentialUpdate) {
            this.credentialUpdate = credentialUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder encodedCredentials(@Nullable String encodedCredentials) {
            this.encodedCredentials = encodedCredentials;
            return this;
        }
        @CustomType.Setter
        public Builder gceSdJson(String gceSdJson) {
            this.gceSdJson = Objects.requireNonNull(gceSdJson);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder minimumMonitors(@Nullable String minimumMonitors) {
            this.minimumMonitors = minimumMonitors;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder tagKey(String tagKey) {
            this.tagKey = Objects.requireNonNull(tagKey);
            return this;
        }
        @CustomType.Setter
        public Builder tagValue(String tagValue) {
            this.tagValue = Objects.requireNonNull(tagValue);
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
        public GetGceServiceDiscoveryResult build() {
            final var _resultValue = new GetGceServiceDiscoveryResult();
            _resultValue.addressRealm = addressRealm;
            _resultValue.credentialUpdate = credentialUpdate;
            _resultValue.encodedCredentials = encodedCredentials;
            _resultValue.gceSdJson = gceSdJson;
            _resultValue.id = id;
            _resultValue.minimumMonitors = minimumMonitors;
            _resultValue.port = port;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.tagKey = tagKey;
            _resultValue.tagValue = tagValue;
            _resultValue.type = type;
            _resultValue.undetectableAction = undetectableAction;
            _resultValue.updateInterval = updateInterval;
            return _resultValue;
        }
    }
}
