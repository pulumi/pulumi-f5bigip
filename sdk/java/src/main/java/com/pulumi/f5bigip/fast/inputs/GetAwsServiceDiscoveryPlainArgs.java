// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAwsServiceDiscoveryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAwsServiceDiscoveryPlainArgs Empty = new GetAwsServiceDiscoveryPlainArgs();

    /**
     * (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    @Import(name="addressRealm")
    private @Nullable String addressRealm;

    /**
     * @return (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    public Optional<String> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }

    /**
     * (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    @Import(name="awsAccessKey")
    private @Nullable String awsAccessKey;

    /**
     * @return (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    public Optional<String> awsAccessKey() {
        return Optional.ofNullable(this.awsAccessKey);
    }

    /**
     * (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
     * 
     */
    @Import(name="awsRegion")
    private @Nullable String awsRegion;

    /**
     * @return (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
     * 
     */
    public Optional<String> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }

    /**
     * (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    @Import(name="awsSecretAccessKey")
    private @Nullable String awsSecretAccessKey;

    /**
     * @return (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
     * 
     */
    public Optional<String> awsSecretAccessKey() {
        return Optional.ofNullable(this.awsSecretAccessKey);
    }

    /**
     * (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
     * 
     */
    @Import(name="credentialUpdate")
    private @Nullable Boolean credentialUpdate;

    /**
     * @return (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
     * 
     */
    public Optional<Boolean> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }

    /**
     * (`optional`,type `string`)AWS externalID field.
     * 
     */
    @Import(name="externalId")
    private @Nullable String externalId;

    /**
     * @return (`optional`,type `string`)AWS externalID field.
     * 
     */
    public Optional<String> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    @Import(name="minimumMonitors")
    private @Nullable String minimumMonitors;

    /**
     * @return (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    public Optional<String> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }

    /**
     * (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
     * 
     */
    @Import(name="port")
    private @Nullable Integer port;

    /**
     * @return (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * (`optional`,type `string`) Assume a role (also requires the `external_id` field)
     * 
     */
    @Import(name="roleArn")
    private @Nullable String roleArn;

    /**
     * @return (`optional`,type `string`) Assume a role (also requires the `external_id` field)
     * 
     */
    public Optional<String> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * (`Required`,type `string`) The tag key associated with the node to add to this pool.
     * 
     */
    @Import(name="tagKey", required=true)
    private String tagKey;

    /**
     * @return (`Required`,type `string`) The tag key associated with the node to add to this pool.
     * 
     */
    public String tagKey() {
        return this.tagKey;
    }

    /**
     * (`Required`,type `string`) The tag value associated with the node to add to this pool.
     * 
     */
    @Import(name="tagValue", required=true)
    private String tagValue;

    /**
     * @return (`Required`,type `string`) The tag value associated with the node to add to this pool.
     * 
     */
    public String tagValue() {
        return this.tagValue;
    }

    @Import(name="type")
    private @Nullable String type;

    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
     * 
     */
    @Import(name="undetectableAction")
    private @Nullable String undetectableAction;

    /**
     * @return (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
     * 
     */
    public Optional<String> undetectableAction() {
        return Optional.ofNullable(this.undetectableAction);
    }

    /**
     * (`optional`,type `string`)Update interval for service discovery.
     * 
     */
    @Import(name="updateInterval")
    private @Nullable String updateInterval;

    /**
     * @return (`optional`,type `string`)Update interval for service discovery.
     * 
     */
    public Optional<String> updateInterval() {
        return Optional.ofNullable(this.updateInterval);
    }

    private GetAwsServiceDiscoveryPlainArgs() {}

    private GetAwsServiceDiscoveryPlainArgs(GetAwsServiceDiscoveryPlainArgs $) {
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
    public static Builder builder(GetAwsServiceDiscoveryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAwsServiceDiscoveryPlainArgs $;

        public Builder() {
            $ = new GetAwsServiceDiscoveryPlainArgs();
        }

        public Builder(GetAwsServiceDiscoveryPlainArgs defaults) {
            $ = new GetAwsServiceDiscoveryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressRealm (`optional`,type `string`)Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(@Nullable String addressRealm) {
            $.addressRealm = addressRealm;
            return this;
        }

        /**
         * @param awsAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsAccessKey(@Nullable String awsAccessKey) {
            $.awsAccessKey = awsAccessKey;
            return this;
        }

        /**
         * @param awsRegion (`optional`,type `string`) AWS region in which ADC is running,default Empty string.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(@Nullable String awsRegion) {
            $.awsRegion = awsRegion;
            return this;
        }

        /**
         * @param awsSecretAccessKey (`optional`,type `string`)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `aws_secret_access_key` field)
         * 
         * @return builder
         * 
         */
        public Builder awsSecretAccessKey(@Nullable String awsSecretAccessKey) {
            $.awsSecretAccessKey = awsSecretAccessKey;
            return this;
        }

        /**
         * @param credentialUpdate (`optional`,type `bool`) Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(@Nullable Boolean credentialUpdate) {
            $.credentialUpdate = credentialUpdate;
            return this;
        }

        /**
         * @param externalId (`optional`,type `string`)AWS externalID field.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable String externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param minimumMonitors (`optional`,type `string`)Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(@Nullable String minimumMonitors) {
            $.minimumMonitors = minimumMonitors;
            return this;
        }

        /**
         * @param port (`optional`,type `int`)Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Integer port) {
            $.port = port;
            return this;
        }

        /**
         * @param roleArn (`optional`,type `string`) Assume a role (also requires the `external_id` field)
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable String roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param tagKey (`Required`,type `string`) The tag key associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagKey(String tagKey) {
            $.tagKey = tagKey;
            return this;
        }

        /**
         * @param tagValue (`Required`,type `string`) The tag value associated with the node to add to this pool.
         * 
         * @return builder
         * 
         */
        public Builder tagValue(String tagValue) {
            $.tagValue = tagValue;
            return this;
        }

        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        /**
         * @param undetectableAction (`optional`,type `string`)Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(@Nullable String undetectableAction) {
            $.undetectableAction = undetectableAction;
            return this;
        }

        /**
         * @param updateInterval (`optional`,type `string`)Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(@Nullable String updateInterval) {
            $.updateInterval = updateInterval;
            return this;
        }

        public GetAwsServiceDiscoveryPlainArgs build() {
            $.tagKey = Objects.requireNonNull($.tagKey, "expected parameter 'tagKey' to be non-null");
            $.tagValue = Objects.requireNonNull($.tagValue, "expected parameter 'tagValue' to be non-null");
            return $;
        }
    }

}
