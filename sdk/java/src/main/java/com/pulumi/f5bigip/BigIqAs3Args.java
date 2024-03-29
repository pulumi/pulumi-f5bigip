// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BigIqAs3Args extends com.pulumi.resources.ResourceArgs {

    public static final BigIqAs3Args Empty = new BigIqAs3Args();

    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    @Import(name="as3Json", required=true)
    private Output<String> as3Json;

    /**
     * @return Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
     * 
     */
    public Output<String> as3Json() {
        return this.as3Json;
    }

    /**
     * Address of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Import(name="bigiqAddress", required=true)
    private Output<String> bigiqAddress;

    /**
     * @return Address of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqAddress() {
        return this.bigiqAddress;
    }

    /**
     * BIGIQ Login reference for token authentication
     * 
     */
    @Import(name="bigiqLoginRef")
    private @Nullable Output<String> bigiqLoginRef;

    /**
     * @return BIGIQ Login reference for token authentication
     * 
     */
    public Optional<Output<String>> bigiqLoginRef() {
        return Optional.ofNullable(this.bigiqLoginRef);
    }

    /**
     * Password of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Import(name="bigiqPassword", required=true)
    private Output<String> bigiqPassword;

    /**
     * @return Password of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqPassword() {
        return this.bigiqPassword;
    }

    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     * 
     */
    @Import(name="bigiqPort")
    private @Nullable Output<String> bigiqPort;

    /**
     * @return type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     * 
     */
    public Optional<Output<String>> bigiqPort() {
        return Optional.ofNullable(this.bigiqPort);
    }

    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     * 
     */
    @Import(name="bigiqTokenAuth")
    private @Nullable Output<Boolean> bigiqTokenAuth;

    /**
     * @return type `bool`, if set to `true` enables Token based Authentication,default is `false`
     * 
     */
    public Optional<Output<Boolean>> bigiqTokenAuth() {
        return Optional.ofNullable(this.bigiqTokenAuth);
    }

    /**
     * User name  of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    @Import(name="bigiqUser", required=true)
    private Output<String> bigiqUser;

    /**
     * @return User name  of the BIG-IQ to which your targer BIG-IP is attached
     * 
     */
    public Output<String> bigiqUser() {
        return this.bigiqUser;
    }

    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to `true`
     * 
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     * 
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     * 
     * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
     * 
     */
    @Import(name="ignoreMetadata")
    private @Nullable Output<Boolean> ignoreMetadata;

    /**
     * @return Set True if you want to ignore metadata changes during update. By default it is set to `true`
     * 
     * * `bigiq_example.json` - Example  AS3 Declarative JSON file
     * 
     * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
     * 
     * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
     * 
     */
    public Optional<Output<Boolean>> ignoreMetadata() {
        return Optional.ofNullable(this.ignoreMetadata);
    }

    /**
     * Name of Tenant
     * 
     */
    @Import(name="tenantList")
    private @Nullable Output<String> tenantList;

    /**
     * @return Name of Tenant
     * 
     */
    public Optional<Output<String>> tenantList() {
        return Optional.ofNullable(this.tenantList);
    }

    private BigIqAs3Args() {}

    private BigIqAs3Args(BigIqAs3Args $) {
        this.as3Json = $.as3Json;
        this.bigiqAddress = $.bigiqAddress;
        this.bigiqLoginRef = $.bigiqLoginRef;
        this.bigiqPassword = $.bigiqPassword;
        this.bigiqPort = $.bigiqPort;
        this.bigiqTokenAuth = $.bigiqTokenAuth;
        this.bigiqUser = $.bigiqUser;
        this.ignoreMetadata = $.ignoreMetadata;
        this.tenantList = $.tenantList;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BigIqAs3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BigIqAs3Args $;

        public Builder() {
            $ = new BigIqAs3Args();
        }

        public Builder(BigIqAs3Args defaults) {
            $ = new BigIqAs3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param as3Json Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
         * 
         * @return builder
         * 
         */
        public Builder as3Json(Output<String> as3Json) {
            $.as3Json = as3Json;
            return this;
        }

        /**
         * @param as3Json Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
         * 
         * @return builder
         * 
         */
        public Builder as3Json(String as3Json) {
            return as3Json(Output.of(as3Json));
        }

        /**
         * @param bigiqAddress Address of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqAddress(Output<String> bigiqAddress) {
            $.bigiqAddress = bigiqAddress;
            return this;
        }

        /**
         * @param bigiqAddress Address of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqAddress(String bigiqAddress) {
            return bigiqAddress(Output.of(bigiqAddress));
        }

        /**
         * @param bigiqLoginRef BIGIQ Login reference for token authentication
         * 
         * @return builder
         * 
         */
        public Builder bigiqLoginRef(@Nullable Output<String> bigiqLoginRef) {
            $.bigiqLoginRef = bigiqLoginRef;
            return this;
        }

        /**
         * @param bigiqLoginRef BIGIQ Login reference for token authentication
         * 
         * @return builder
         * 
         */
        public Builder bigiqLoginRef(String bigiqLoginRef) {
            return bigiqLoginRef(Output.of(bigiqLoginRef));
        }

        /**
         * @param bigiqPassword Password of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqPassword(Output<String> bigiqPassword) {
            $.bigiqPassword = bigiqPassword;
            return this;
        }

        /**
         * @param bigiqPassword Password of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqPassword(String bigiqPassword) {
            return bigiqPassword(Output.of(bigiqPassword));
        }

        /**
         * @param bigiqPort type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
         * 
         * @return builder
         * 
         */
        public Builder bigiqPort(@Nullable Output<String> bigiqPort) {
            $.bigiqPort = bigiqPort;
            return this;
        }

        /**
         * @param bigiqPort type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
         * 
         * @return builder
         * 
         */
        public Builder bigiqPort(String bigiqPort) {
            return bigiqPort(Output.of(bigiqPort));
        }

        /**
         * @param bigiqTokenAuth type `bool`, if set to `true` enables Token based Authentication,default is `false`
         * 
         * @return builder
         * 
         */
        public Builder bigiqTokenAuth(@Nullable Output<Boolean> bigiqTokenAuth) {
            $.bigiqTokenAuth = bigiqTokenAuth;
            return this;
        }

        /**
         * @param bigiqTokenAuth type `bool`, if set to `true` enables Token based Authentication,default is `false`
         * 
         * @return builder
         * 
         */
        public Builder bigiqTokenAuth(Boolean bigiqTokenAuth) {
            return bigiqTokenAuth(Output.of(bigiqTokenAuth));
        }

        /**
         * @param bigiqUser User name  of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqUser(Output<String> bigiqUser) {
            $.bigiqUser = bigiqUser;
            return this;
        }

        /**
         * @param bigiqUser User name  of the BIG-IQ to which your targer BIG-IP is attached
         * 
         * @return builder
         * 
         */
        public Builder bigiqUser(String bigiqUser) {
            return bigiqUser(Output.of(bigiqUser));
        }

        /**
         * @param ignoreMetadata Set True if you want to ignore metadata changes during update. By default it is set to `true`
         * 
         * * `bigiq_example.json` - Example  AS3 Declarative JSON file
         * 
         * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
         * 
         * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
         * 
         * @return builder
         * 
         */
        public Builder ignoreMetadata(@Nullable Output<Boolean> ignoreMetadata) {
            $.ignoreMetadata = ignoreMetadata;
            return this;
        }

        /**
         * @param ignoreMetadata Set True if you want to ignore metadata changes during update. By default it is set to `true`
         * 
         * * `bigiq_example.json` - Example  AS3 Declarative JSON file
         * 
         * * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
         * 
         * &gt;  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
         * 
         * @return builder
         * 
         */
        public Builder ignoreMetadata(Boolean ignoreMetadata) {
            return ignoreMetadata(Output.of(ignoreMetadata));
        }

        /**
         * @param tenantList Name of Tenant
         * 
         * @return builder
         * 
         */
        public Builder tenantList(@Nullable Output<String> tenantList) {
            $.tenantList = tenantList;
            return this;
        }

        /**
         * @param tenantList Name of Tenant
         * 
         * @return builder
         * 
         */
        public Builder tenantList(String tenantList) {
            return tenantList(Output.of(tenantList));
        }

        public BigIqAs3Args build() {
            if ($.as3Json == null) {
                throw new MissingRequiredPropertyException("BigIqAs3Args", "as3Json");
            }
            if ($.bigiqAddress == null) {
                throw new MissingRequiredPropertyException("BigIqAs3Args", "bigiqAddress");
            }
            if ($.bigiqPassword == null) {
                throw new MissingRequiredPropertyException("BigIqAs3Args", "bigiqPassword");
            }
            if ($.bigiqUser == null) {
                throw new MissingRequiredPropertyException("BigIqAs3Args", "bigiqUser");
            }
            return $;
        }
    }

}
