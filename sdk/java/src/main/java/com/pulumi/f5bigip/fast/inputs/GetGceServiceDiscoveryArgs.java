// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGceServiceDiscoveryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGceServiceDiscoveryArgs Empty = new GetGceServiceDiscoveryArgs();

    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    @Import(name="addressRealm")
    private @Nullable Output<String> addressRealm;

    /**
     * @return Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    public Optional<Output<String>> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }

    /**
     * Specifies whether you are updating your credentials,default `false`.
     * 
     */
    @Import(name="credentialUpdate")
    private @Nullable Output<Boolean> credentialUpdate;

    /**
     * @return Specifies whether you are updating your credentials,default `false`.
     * 
     */
    public Optional<Output<Boolean>> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }

    /**
     * Base 64 encoded service account credentials JSON.
     * 
     */
    @Import(name="encodedCredentials")
    private @Nullable Output<String> encodedCredentials;

    /**
     * @return Base 64 encoded service account credentials JSON.
     * 
     */
    public Optional<Output<String>> encodedCredentials() {
        return Optional.ofNullable(this.encodedCredentials);
    }

    /**
     * Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    @Import(name="minimumMonitors")
    private @Nullable Output<String> minimumMonitors;

    /**
     * @return Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    public Optional<Output<String>> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }

    /**
     * Port to be used for AWS service discovery,default `80`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port to be used for AWS service discovery,default `80`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * GCE region in which ADC is running.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return GCE region in which ADC is running.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * The tag key associated with the node to add to this pool.
     * 
     */
    @Import(name="tagKey", required=true)
    private Output<String> tagKey;

    /**
     * @return The tag key associated with the node to add to this pool.
     * 
     */
    public Output<String> tagKey() {
        return this.tagKey;
    }

    /**
     * The tag value associated with the node to add to this pool.
     * 
     */
    @Import(name="tagValue", required=true)
    private Output<String> tagValue;

    /**
     * @return The tag value associated with the node to add to this pool.
     * 
     */
    public Output<String> tagValue() {
        return this.tagValue;
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Action to take when node cannot be detected,default `remove`.
     * 
     */
    @Import(name="undetectableAction")
    private @Nullable Output<String> undetectableAction;

    /**
     * @return Action to take when node cannot be detected,default `remove`.
     * 
     */
    public Optional<Output<String>> undetectableAction() {
        return Optional.ofNullable(this.undetectableAction);
    }

    /**
     * Update interval for service discovery.
     * 
     */
    @Import(name="updateInterval")
    private @Nullable Output<String> updateInterval;

    /**
     * @return Update interval for service discovery.
     * 
     */
    public Optional<Output<String>> updateInterval() {
        return Optional.ofNullable(this.updateInterval);
    }

    private GetGceServiceDiscoveryArgs() {}

    private GetGceServiceDiscoveryArgs(GetGceServiceDiscoveryArgs $) {
        this.addressRealm = $.addressRealm;
        this.credentialUpdate = $.credentialUpdate;
        this.encodedCredentials = $.encodedCredentials;
        this.minimumMonitors = $.minimumMonitors;
        this.port = $.port;
        this.projectId = $.projectId;
        this.region = $.region;
        this.tagKey = $.tagKey;
        this.tagValue = $.tagValue;
        this.type = $.type;
        this.undetectableAction = $.undetectableAction;
        this.updateInterval = $.updateInterval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGceServiceDiscoveryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGceServiceDiscoveryArgs $;

        public Builder() {
            $ = new GetGceServiceDiscoveryArgs();
        }

        public Builder(GetGceServiceDiscoveryArgs defaults) {
            $ = new GetGceServiceDiscoveryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressRealm Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(@Nullable Output<String> addressRealm) {
            $.addressRealm = addressRealm;
            return this;
        }

        /**
         * @param addressRealm Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(String addressRealm) {
            return addressRealm(Output.of(addressRealm));
        }

        /**
         * @param credentialUpdate Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(@Nullable Output<Boolean> credentialUpdate) {
            $.credentialUpdate = credentialUpdate;
            return this;
        }

        /**
         * @param credentialUpdate Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(Boolean credentialUpdate) {
            return credentialUpdate(Output.of(credentialUpdate));
        }

        /**
         * @param encodedCredentials Base 64 encoded service account credentials JSON.
         * 
         * @return builder
         * 
         */
        public Builder encodedCredentials(@Nullable Output<String> encodedCredentials) {
            $.encodedCredentials = encodedCredentials;
            return this;
        }

        /**
         * @param encodedCredentials Base 64 encoded service account credentials JSON.
         * 
         * @return builder
         * 
         */
        public Builder encodedCredentials(String encodedCredentials) {
            return encodedCredentials(Output.of(encodedCredentials));
        }

        /**
         * @param minimumMonitors Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(@Nullable Output<String> minimumMonitors) {
            $.minimumMonitors = minimumMonitors;
            return this;
        }

        /**
         * @param minimumMonitors Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(String minimumMonitors) {
            return minimumMonitors(Output.of(minimumMonitors));
        }

        /**
         * @param port Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param projectId For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region GCE region in which ADC is running.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region GCE region in which ADC is running.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tagKey The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(Output<String> tagKey) {
            $.tagKey = tagKey;
            return this;
        }

        /**
         * @param tagKey The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(String tagKey) {
            return tagKey(Output.of(tagKey));
        }

        /**
         * @param tagValue The tag value associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagValue(Output<String> tagValue) {
            $.tagValue = tagValue;
            return this;
        }

        /**
         * @param tagValue The tag value associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagValue(String tagValue) {
            return tagValue(Output.of(tagValue));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param undetectableAction Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(@Nullable Output<String> undetectableAction) {
            $.undetectableAction = undetectableAction;
            return this;
        }

        /**
         * @param undetectableAction Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(String undetectableAction) {
            return undetectableAction(Output.of(undetectableAction));
        }

        /**
         * @param updateInterval Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(@Nullable Output<String> updateInterval) {
            $.updateInterval = updateInterval;
            return this;
        }

        /**
         * @param updateInterval Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(String updateInterval) {
            return updateInterval(Output.of(updateInterval));
        }

        public GetGceServiceDiscoveryArgs build() {
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            $.tagKey = Objects.requireNonNull($.tagKey, "expected parameter 'tagKey' to be non-null");
            $.tagValue = Objects.requireNonNull($.tagValue, "expected parameter 'tagValue' to be non-null");
            return $;
        }
    }

}
