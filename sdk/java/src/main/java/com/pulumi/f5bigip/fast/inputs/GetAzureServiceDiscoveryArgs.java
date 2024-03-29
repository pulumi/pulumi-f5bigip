// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAzureServiceDiscoveryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAzureServiceDiscoveryArgs Empty = new GetAzureServiceDiscoveryArgs();

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
     * Port to be used for Azure service discovery,default `80`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port to be used for Azure service discovery,default `80`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Azure Resource Group name.
     * 
     */
    @Import(name="resourceGroup", required=true)
    private Output<String> resourceGroup;

    /**
     * @return Azure Resource Group name.
     * 
     */
    public Output<String> resourceGroup() {
        return this.resourceGroup;
    }

    /**
     * Azure subscription ID.
     * 
     */
    @Import(name="subscriptionId", required=true)
    private Output<String> subscriptionId;

    /**
     * @return Azure subscription ID.
     * 
     */
    public Output<String> subscriptionId() {
        return this.subscriptionId;
    }

    /**
     * The tag key associated with the node to add to this pool.
     * 
     */
    @Import(name="tagKey")
    private @Nullable Output<String> tagKey;

    /**
     * @return The tag key associated with the node to add to this pool.
     * 
     */
    public Optional<Output<String>> tagKey() {
        return Optional.ofNullable(this.tagKey);
    }

    /**
     * The tag value associated with the node to add to this pool.
     * 
     */
    @Import(name="tagValue")
    private @Nullable Output<String> tagValue;

    /**
     * @return The tag value associated with the node to add to this pool.
     * 
     */
    public Optional<Output<String>> tagValue() {
        return Optional.ofNullable(this.tagValue);
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

    private GetAzureServiceDiscoveryArgs() {}

    private GetAzureServiceDiscoveryArgs(GetAzureServiceDiscoveryArgs $) {
        this.addressRealm = $.addressRealm;
        this.credentialUpdate = $.credentialUpdate;
        this.minimumMonitors = $.minimumMonitors;
        this.port = $.port;
        this.resourceGroup = $.resourceGroup;
        this.subscriptionId = $.subscriptionId;
        this.tagKey = $.tagKey;
        this.tagValue = $.tagValue;
        this.type = $.type;
        this.undetectableAction = $.undetectableAction;
        this.updateInterval = $.updateInterval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAzureServiceDiscoveryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAzureServiceDiscoveryArgs $;

        public Builder() {
            $ = new GetAzureServiceDiscoveryArgs();
        }

        public Builder(GetAzureServiceDiscoveryArgs defaults) {
            $ = new GetAzureServiceDiscoveryArgs(Objects.requireNonNull(defaults));
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
         * @param port Port to be used for Azure service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port to be used for Azure service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param resourceGroup Azure Resource Group name.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(Output<String> resourceGroup) {
            $.resourceGroup = resourceGroup;
            return this;
        }

        /**
         * @param resourceGroup Azure Resource Group name.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroup(String resourceGroup) {
            return resourceGroup(Output.of(resourceGroup));
        }

        /**
         * @param subscriptionId Azure subscription ID.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(Output<String> subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        /**
         * @param subscriptionId Azure subscription ID.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(String subscriptionId) {
            return subscriptionId(Output.of(subscriptionId));
        }

        /**
         * @param tagKey The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(@Nullable Output<String> tagKey) {
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
        public Builder tagValue(@Nullable Output<String> tagValue) {
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

        public GetAzureServiceDiscoveryArgs build() {
            if ($.resourceGroup == null) {
                throw new MissingRequiredPropertyException("GetAzureServiceDiscoveryArgs", "resourceGroup");
            }
            if ($.subscriptionId == null) {
                throw new MissingRequiredPropertyException("GetAzureServiceDiscoveryArgs", "subscriptionId");
            }
            return $;
        }
    }

}
