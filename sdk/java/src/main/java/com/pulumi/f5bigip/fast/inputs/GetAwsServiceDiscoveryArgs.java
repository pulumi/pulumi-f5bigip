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


public final class GetAwsServiceDiscoveryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAwsServiceDiscoveryArgs Empty = new GetAwsServiceDiscoveryArgs();

    /**
     * (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    @Import(name="addressRealm")
    private @Nullable Output<String> addressRealm;

    /**
     * @return (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    public Optional<Output<String>> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }

    /**
     * (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    @Import(name="awsAccessKey")
    private @Nullable Output<String> awsAccessKey;

    /**
     * @return (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    public Optional<Output<String>> awsAccessKey() {
        return Optional.ofNullable(this.awsAccessKey);
    }

    /**
     * (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
     * 
     */
    @Import(name="awsRegion")
    private @Nullable Output<String> awsRegion;

    /**
     * @return (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
     * 
     */
    public Optional<Output<String>> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }

    /**
     * (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    @Import(name="awsSecretAccessKey")
    private @Nullable Output<String> awsSecretAccessKey;

    /**
     * @return (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    public Optional<Output<String>> awsSecretAccessKey() {
        return Optional.ofNullable(this.awsSecretAccessKey);
    }

    /**
     * (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
     * 
     */
    @Import(name="credentialUpdate")
    private @Nullable Output<Boolean> credentialUpdate;

    /**
     * @return (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
     * 
     */
    public Optional<Output<Boolean>> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }

    /**
     * (`optional`,type `string`)AWS externalID field.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return (`optional`,type `string`)AWS externalID field.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    @Import(name="minimumMonitors")
    private @Nullable Output<String> minimumMonitors;

    /**
     * @return (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    public Optional<Output<String>> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }

    /**
     * (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * (`optional`,type `string`) Assume a role (also requires the `external_id` field)
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return (`optional`,type `string`) Assume a role (also requires the `external_id` field)
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * (`Required`,type `string`) The tag key associated with the node to add to this pool.
     * 
     */
    @Import(name="tagKey", required=true)
    private Output<String> tagKey;

    /**
     * @return (`Required`,type `string`) The tag key associated with the node to add to this pool.
     * 
     */
    public Output<String> tagKey() {
        return this.tagKey;
    }

    /**
     * (`Required`,type `string`) The tag value associated with the node to add to this pool.
     * 
     */
    @Import(name="tagValue", required=true)
    private Output<String> tagValue;

    /**
     * @return (`Required`,type `string`) The tag value associated with the node to add to this pool.
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
     * (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
     * 
     */
    @Import(name="undetectableAction")
    private @Nullable Output<String> undetectableAction;

    /**
     * @return (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
     * 
     */
    public Optional<Output<String>> undetectableAction() {
        return Optional.ofNullable(this.undetectableAction);
    }

    /**
     * (`optional`,type `string`)Update interval for service discovery.
     * 
     */
    @Import(name="updateInterval")
    private @Nullable Output<String> updateInterval;

    /**
     * @return (`optional`,type `string`)Update interval for service discovery.
     * 
     */
    public Optional<Output<String>> updateInterval() {
        return Optional.ofNullable(this.updateInterval);
    }

    private GetAwsServiceDiscoveryArgs() {}

    private GetAwsServiceDiscoveryArgs(GetAwsServiceDiscoveryArgs $) {
        this.addressRealm = $.addressRealm;
        this.awsAccessKey = $.awsAccessKey;
        this.awsRegion = $.awsRegion;
        this.awsSecretAccessKey = $.awsSecretAccessKey;
        this.credentialUpdate = $.credentialUpdate;
        this.externalId = $.externalId;
        this.minimumMonitors = $.minimumMonitors;
        this.port = $.port;
        this.roleArn = $.roleArn;
        this.tagKey = $.tagKey;
        this.tagValue = $.tagValue;
        this.type = $.type;
        this.undetectableAction = $.undetectableAction;
        this.updateInterval = $.updateInterval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAwsServiceDiscoveryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAwsServiceDiscoveryArgs $;

        public Builder() {
            $ = new GetAwsServiceDiscoveryArgs();
        }

        public Builder(GetAwsServiceDiscoveryArgs defaults) {
            $ = new GetAwsServiceDiscoveryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressRealm (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(@Nullable Output<String> addressRealm) {
            $.addressRealm = addressRealm;
            return this;
        }

        /**
         * @param addressRealm (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(String addressRealm) {
            return addressRealm(Output.of(addressRealm));
        }

        /**
         * @param awsAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsAccessKey(@Nullable Output<String> awsAccessKey) {
            $.awsAccessKey = awsAccessKey;
            return this;
        }

        /**
         * @param awsAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsAccessKey(String awsAccessKey) {
            return awsAccessKey(Output.of(awsAccessKey));
        }

        /**
         * @param awsRegion (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(@Nullable Output<String> awsRegion) {
            $.awsRegion = awsRegion;
            return this;
        }

        /**
         * @param awsRegion (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(String awsRegion) {
            return awsRegion(Output.of(awsRegion));
        }

        /**
         * @param awsSecretAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsSecretAccessKey(@Nullable Output<String> awsSecretAccessKey) {
            $.awsSecretAccessKey = awsSecretAccessKey;
            return this;
        }

        /**
         * @param awsSecretAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsSecretAccessKey(String awsSecretAccessKey) {
            return awsSecretAccessKey(Output.of(awsSecretAccessKey));
        }

        /**
         * @param credentialUpdate (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(@Nullable Output<Boolean> credentialUpdate) {
            $.credentialUpdate = credentialUpdate;
            return this;
        }

        /**
         * @param credentialUpdate (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(Boolean credentialUpdate) {
            return credentialUpdate(Output.of(credentialUpdate));
        }

        /**
         * @param externalId (`optional`,type `string`)AWS externalID field.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId (`optional`,type `string`)AWS externalID field.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param minimumMonitors (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(@Nullable Output<String> minimumMonitors) {
            $.minimumMonitors = minimumMonitors;
            return this;
        }

        /**
         * @param minimumMonitors (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(String minimumMonitors) {
            return minimumMonitors(Output.of(minimumMonitors));
        }

        /**
         * @param port (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param roleArn (`optional`,type `string`) Assume a role (also requires the `external_id` field)
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn (`optional`,type `string`) Assume a role (also requires the `external_id` field)
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tagKey (`Required`,type `string`) The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(Output<String> tagKey) {
            $.tagKey = tagKey;
            return this;
        }

        /**
         * @param tagKey (`Required`,type `string`) The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(String tagKey) {
            return tagKey(Output.of(tagKey));
        }

        /**
         * @param tagValue (`Required`,type `string`) The tag value associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagValue(Output<String> tagValue) {
            $.tagValue = tagValue;
            return this;
        }

        /**
         * @param tagValue (`Required`,type `string`) The tag value associated with the node to add to this pool.
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
         * @param undetectableAction (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(@Nullable Output<String> undetectableAction) {
            $.undetectableAction = undetectableAction;
            return this;
        }

        /**
         * @param undetectableAction (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(String undetectableAction) {
            return undetectableAction(Output.of(undetectableAction));
        }

        /**
         * @param updateInterval (`optional`,type `string`)Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(@Nullable Output<String> updateInterval) {
            $.updateInterval = updateInterval;
            return this;
        }

        /**
         * @param updateInterval (`optional`,type `string`)Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(String updateInterval) {
            return updateInterval(Output.of(updateInterval));
        }

        public GetAwsServiceDiscoveryArgs build() {
            $.tagKey = Objects.requireNonNull($.tagKey, "expected parameter 'tagKey' to be non-null");
            $.tagValue = Objects.requireNonNull($.tagValue, "expected parameter 'tagValue' to be non-null");
            return $;
        }
    }

}
