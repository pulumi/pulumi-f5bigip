// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DoState extends com.pulumi.resources.ResourceArgs {

    public static final DoState Empty = new DoState();

    /**
     * IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Import(name="bigipAddress")
    private @Nullable Output<String> bigipAddress;

    /**
     * @return IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Optional<Output<String>> bigipAddress() {
        return Optional.ofNullable(this.bigipAddress);
    }

    /**
     * Password of BIGIP host to be used for this resource
     * 
     */
    @Import(name="bigipPassword")
    private @Nullable Output<String> bigipPassword;

    /**
     * @return Password of BIGIP host to be used for this resource
     * 
     */
    public Optional<Output<String>> bigipPassword() {
        return Optional.ofNullable(this.bigipPassword);
    }

    /**
     * Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Import(name="bigipPort")
    private @Nullable Output<String> bigipPort;

    /**
     * @return Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Optional<Output<String>> bigipPort() {
        return Optional.ofNullable(this.bigipPort);
    }

    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     * 
     */
    @Import(name="bigipTokenAuth")
    private @Nullable Output<Boolean> bigipTokenAuth;

    /**
     * @return Enable to use an external authentication source (LDAP, TACACS, etc)
     * 
     */
    public Optional<Output<Boolean>> bigipTokenAuth() {
        return Optional.ofNullable(this.bigipTokenAuth);
    }

    /**
     * UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Import(name="bigipUser")
    private @Nullable Output<String> bigipUser;

    /**
     * @return UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Optional<Output<String>> bigipUser() {
        return Optional.ofNullable(this.bigipUser);
    }

    /**
     * Name of the of the Declarative DO JSON file
     * 
     */
    @Import(name="doJson")
    private @Nullable Output<String> doJson;

    /**
     * @return Name of the of the Declarative DO JSON file
     * 
     */
    public Optional<Output<String>> doJson() {
        return Optional.ofNullable(this.doJson);
    }

    /**
     * unique identifier for DO resource
     * 
     * @deprecated
     * this attribute is no longer in use
     * 
     */
    @Deprecated /* this attribute is no longer in use */
    @Import(name="tenantName")
    private @Nullable Output<String> tenantName;

    /**
     * @return unique identifier for DO resource
     * 
     * @deprecated
     * this attribute is no longer in use
     * 
     */
    @Deprecated /* this attribute is no longer in use */
    public Optional<Output<String>> tenantName() {
        return Optional.ofNullable(this.tenantName);
    }

    /**
     * DO json
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return DO json
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private DoState() {}

    private DoState(DoState $) {
        this.bigipAddress = $.bigipAddress;
        this.bigipPassword = $.bigipPassword;
        this.bigipPort = $.bigipPort;
        this.bigipTokenAuth = $.bigipTokenAuth;
        this.bigipUser = $.bigipUser;
        this.doJson = $.doJson;
        this.tenantName = $.tenantName;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DoState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DoState $;

        public Builder() {
            $ = new DoState();
        }

        public Builder(DoState defaults) {
            $ = new DoState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bigipAddress IP Address of BIGIP Host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipAddress(@Nullable Output<String> bigipAddress) {
            $.bigipAddress = bigipAddress;
            return this;
        }

        /**
         * @param bigipAddress IP Address of BIGIP Host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipAddress(String bigipAddress) {
            return bigipAddress(Output.of(bigipAddress));
        }

        /**
         * @param bigipPassword Password of BIGIP host to be used for this resource
         * 
         * @return builder
         * 
         */
        public Builder bigipPassword(@Nullable Output<String> bigipPassword) {
            $.bigipPassword = bigipPassword;
            return this;
        }

        /**
         * @param bigipPassword Password of BIGIP host to be used for this resource
         * 
         * @return builder
         * 
         */
        public Builder bigipPassword(String bigipPassword) {
            return bigipPassword(Output.of(bigipPassword));
        }

        /**
         * @param bigipPort Port number of BIGIP host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipPort(@Nullable Output<String> bigipPort) {
            $.bigipPort = bigipPort;
            return this;
        }

        /**
         * @param bigipPort Port number of BIGIP host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipPort(String bigipPort) {
            return bigipPort(Output.of(bigipPort));
        }

        /**
         * @param bigipTokenAuth Enable to use an external authentication source (LDAP, TACACS, etc)
         * 
         * @return builder
         * 
         */
        public Builder bigipTokenAuth(@Nullable Output<Boolean> bigipTokenAuth) {
            $.bigipTokenAuth = bigipTokenAuth;
            return this;
        }

        /**
         * @param bigipTokenAuth Enable to use an external authentication source (LDAP, TACACS, etc)
         * 
         * @return builder
         * 
         */
        public Builder bigipTokenAuth(Boolean bigipTokenAuth) {
            return bigipTokenAuth(Output.of(bigipTokenAuth));
        }

        /**
         * @param bigipUser UserName of BIGIP host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipUser(@Nullable Output<String> bigipUser) {
            $.bigipUser = bigipUser;
            return this;
        }

        /**
         * @param bigipUser UserName of BIGIP host to be used for this resource,this is optional parameter.
         * whenever we specify this parameter it gets overwrite provider configuration
         * 
         * @return builder
         * 
         */
        public Builder bigipUser(String bigipUser) {
            return bigipUser(Output.of(bigipUser));
        }

        /**
         * @param doJson Name of the of the Declarative DO JSON file
         * 
         * @return builder
         * 
         */
        public Builder doJson(@Nullable Output<String> doJson) {
            $.doJson = doJson;
            return this;
        }

        /**
         * @param doJson Name of the of the Declarative DO JSON file
         * 
         * @return builder
         * 
         */
        public Builder doJson(String doJson) {
            return doJson(Output.of(doJson));
        }

        /**
         * @param tenantName unique identifier for DO resource
         * 
         * @return builder
         * 
         * @deprecated
         * this attribute is no longer in use
         * 
         */
        @Deprecated /* this attribute is no longer in use */
        public Builder tenantName(@Nullable Output<String> tenantName) {
            $.tenantName = tenantName;
            return this;
        }

        /**
         * @param tenantName unique identifier for DO resource
         * 
         * @return builder
         * 
         * @deprecated
         * this attribute is no longer in use
         * 
         */
        @Deprecated /* this attribute is no longer in use */
        public Builder tenantName(String tenantName) {
            return tenantName(Output.of(tenantName));
        }

        /**
         * @param timeout DO json
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout DO json
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public DoState build() {
            return $;
        }
    }

}
