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


public final class GetConsulServiceDiscoveryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetConsulServiceDiscoveryPlainArgs Empty = new GetConsulServiceDiscoveryPlainArgs();

    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    @Import(name="addressRealm")
    private @Nullable String addressRealm;

    /**
     * @return Specifies whether to look for public or private IP addresses,default `private`.
     * 
     */
    public Optional<String> addressRealm() {
        return Optional.ofNullable(this.addressRealm);
    }

    /**
     * Specifies whether you are updating your credentials,default `false`.
     * 
     */
    @Import(name="credentialUpdate")
    private @Nullable Boolean credentialUpdate;

    /**
     * @return Specifies whether you are updating your credentials,default `false`.
     * 
     */
    public Optional<Boolean> credentialUpdate() {
        return Optional.ofNullable(this.credentialUpdate);
    }

    /**
     * Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
     * 
     */
    @Import(name="encodedToken")
    private @Nullable String encodedToken;

    /**
     * @return Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
     * 
     */
    public Optional<String> encodedToken() {
        return Optional.ofNullable(this.encodedToken);
    }

    /**
     * Custom JMESPath Query.
     * 
     */
    @Import(name="jmesPathQuery")
    private @Nullable String jmesPathQuery;

    /**
     * @return Custom JMESPath Query.
     * 
     */
    public Optional<String> jmesPathQuery() {
        return Optional.ofNullable(this.jmesPathQuery);
    }

    /**
     * Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    @Import(name="minimumMonitors")
    private @Nullable String minimumMonitors;

    /**
     * @return Member is down when fewer than minimum monitors report it healthy.
     * 
     */
    public Optional<String> minimumMonitors() {
        return Optional.ofNullable(this.minimumMonitors);
    }

    /**
     * Port to be used for AWS service discovery,default `80`.
     * 
     */
    @Import(name="port", required=true)
    private Integer port;

    /**
     * @return Port to be used for AWS service discovery,default `80`.
     * 
     */
    public Integer port() {
        return this.port;
    }

    /**
     * If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
     * 
     */
    @Import(name="rejectUnauthorized")
    private @Nullable Boolean rejectUnauthorized;

    /**
     * @return If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
     * 
     */
    public Optional<Boolean> rejectUnauthorized() {
        return Optional.ofNullable(this.rejectUnauthorized);
    }

    /**
     * CA Bundle to validate server certificates.
     * 
     */
    @Import(name="trustCa")
    private @Nullable String trustCa;

    /**
     * @return CA Bundle to validate server certificates.
     * 
     */
    public Optional<String> trustCa() {
        return Optional.ofNullable(this.trustCa);
    }

    @Import(name="type")
    private @Nullable String type;

    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Action to take when node cannot be detected,default `remove`.
     * 
     */
    @Import(name="undetectableAction")
    private @Nullable String undetectableAction;

    /**
     * @return Action to take when node cannot be detected,default `remove`.
     * 
     */
    public Optional<String> undetectableAction() {
        return Optional.ofNullable(this.undetectableAction);
    }

    /**
     * Update interval for service discovery.
     * 
     */
    @Import(name="updateInterval")
    private @Nullable String updateInterval;

    /**
     * @return Update interval for service discovery.
     * 
     */
    public Optional<String> updateInterval() {
        return Optional.ofNullable(this.updateInterval);
    }

    /**
     * The location of the node data.
     * 
     */
    @Import(name="uri", required=true)
    private String uri;

    /**
     * @return The location of the node data.
     * 
     */
    public String uri() {
        return this.uri;
    }

    private GetConsulServiceDiscoveryPlainArgs() {}

    private GetConsulServiceDiscoveryPlainArgs(GetConsulServiceDiscoveryPlainArgs $) {
        this.addressRealm = $.addressRealm;
        this.credentialUpdate = $.credentialUpdate;
        this.encodedToken = $.encodedToken;
        this.jmesPathQuery = $.jmesPathQuery;
        this.minimumMonitors = $.minimumMonitors;
        this.port = $.port;
        this.rejectUnauthorized = $.rejectUnauthorized;
        this.trustCa = $.trustCa;
        this.type = $.type;
        this.undetectableAction = $.undetectableAction;
        this.updateInterval = $.updateInterval;
        this.uri = $.uri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetConsulServiceDiscoveryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetConsulServiceDiscoveryPlainArgs $;

        public Builder() {
            $ = new GetConsulServiceDiscoveryPlainArgs();
        }

        public Builder(GetConsulServiceDiscoveryPlainArgs defaults) {
            $ = new GetConsulServiceDiscoveryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressRealm Specifies whether to look for public or private IP addresses,default `private`.
         * 
         * @return builder
         * 
         */
        public Builder addressRealm(@Nullable String addressRealm) {
            $.addressRealm = addressRealm;
            return this;
        }

        /**
         * @param credentialUpdate Specifies whether you are updating your credentials,default `false`.
         * 
         * @return builder
         * 
         */
        public Builder credentialUpdate(@Nullable Boolean credentialUpdate) {
            $.credentialUpdate = credentialUpdate;
            return this;
        }

        /**
         * @param encodedToken Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
         * 
         * @return builder
         * 
         */
        public Builder encodedToken(@Nullable String encodedToken) {
            $.encodedToken = encodedToken;
            return this;
        }

        /**
         * @param jmesPathQuery Custom JMESPath Query.
         * 
         * @return builder
         * 
         */
        public Builder jmesPathQuery(@Nullable String jmesPathQuery) {
            $.jmesPathQuery = jmesPathQuery;
            return this;
        }

        /**
         * @param minimumMonitors Member is down when fewer than minimum monitors report it healthy.
         * 
         * @return builder
         * 
         */
        public Builder minimumMonitors(@Nullable String minimumMonitors) {
            $.minimumMonitors = minimumMonitors;
            return this;
        }

        /**
         * @param port Port to be used for AWS service discovery,default `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            $.port = port;
            return this;
        }

        /**
         * @param rejectUnauthorized If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
         * 
         * @return builder
         * 
         */
        public Builder rejectUnauthorized(@Nullable Boolean rejectUnauthorized) {
            $.rejectUnauthorized = rejectUnauthorized;
            return this;
        }

        /**
         * @param trustCa CA Bundle to validate server certificates.
         * 
         * @return builder
         * 
         */
        public Builder trustCa(@Nullable String trustCa) {
            $.trustCa = trustCa;
            return this;
        }

        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        /**
         * @param undetectableAction Action to take when node cannot be detected,default `remove`.
         * 
         * @return builder
         * 
         */
        public Builder undetectableAction(@Nullable String undetectableAction) {
            $.undetectableAction = undetectableAction;
            return this;
        }

        /**
         * @param updateInterval Update interval for service discovery.
         * 
         * @return builder
         * 
         */
        public Builder updateInterval(@Nullable String updateInterval) {
            $.updateInterval = updateInterval;
            return this;
        }

        /**
         * @param uri The location of the node data.
         * 
         * @return builder
         * 
         */
        public Builder uri(String uri) {
            $.uri = uri;
            return this;
        }

        public GetConsulServiceDiscoveryPlainArgs build() {
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            $.uri = Objects.requireNonNull($.uri, "expected parameter 'uri' to be non-null");
            return $;
        }
    }

}