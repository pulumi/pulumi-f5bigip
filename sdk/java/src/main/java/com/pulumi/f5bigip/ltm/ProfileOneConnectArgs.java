// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProfileOneConnectArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProfileOneConnectArgs Empty = new ProfileOneConnectArgs();

    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    @Import(name="defaultsFrom")
    private @Nullable Output<String> defaultsFrom;

    /**
     * @return Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    public Optional<Output<String>> defaultsFrom() {
        return Optional.ofNullable(this.defaultsFrom);
    }

    /**
     * Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     * 
     */
    @Import(name="idleTimeoutOverride")
    private @Nullable Output<String> idleTimeoutOverride;

    /**
     * @return Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     * 
     */
    public Optional<Output<String>> idleTimeoutOverride() {
        return Optional.ofNullable(this.idleTimeoutOverride);
    }

    /**
     * Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     * 
     */
    @Import(name="limitType")
    private @Nullable Output<String> limitType;

    /**
     * @return Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     * 
     */
    public Optional<Output<String>> limitType() {
        return Optional.ofNullable(this.limitType);
    }

    /**
     * Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     * 
     */
    @Import(name="maxAge")
    private @Nullable Output<Integer> maxAge;

    /**
     * @return Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     * 
     */
    public Optional<Output<Integer>> maxAge() {
        return Optional.ofNullable(this.maxAge);
    }

    /**
     * Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     * 
     */
    @Import(name="maxReuse")
    private @Nullable Output<Integer> maxReuse;

    /**
     * @return Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     * 
     */
    public Optional<Output<Integer>> maxReuse() {
        return Optional.ofNullable(this.maxReuse);
    }

    /**
     * Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     * 
     */
    @Import(name="maxSize")
    private @Nullable Output<Integer> maxSize;

    /**
     * @return Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     * 
     */
    public Optional<Output<Integer>> maxSize() {
        return Optional.ofNullable(this.maxSize);
    }

    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Displays the administrative partition within which this profile resides
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Displays the administrative partition within which this profile resides
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    /**
     * Specify if you want to share the pool, default value is `disabled`.
     * 
     */
    @Import(name="sharePools")
    private @Nullable Output<String> sharePools;

    /**
     * @return Specify if you want to share the pool, default value is `disabled`.
     * 
     */
    public Optional<Output<String>> sharePools() {
        return Optional.ofNullable(this.sharePools);
    }

    /**
     * Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
     * 
     */
    @Import(name="sourceMask")
    private @Nullable Output<String> sourceMask;

    /**
     * @return Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
     * 
     */
    public Optional<Output<String>> sourceMask() {
        return Optional.ofNullable(this.sourceMask);
    }

    private ProfileOneConnectArgs() {}

    private ProfileOneConnectArgs(ProfileOneConnectArgs $) {
        this.defaultsFrom = $.defaultsFrom;
        this.idleTimeoutOverride = $.idleTimeoutOverride;
        this.limitType = $.limitType;
        this.maxAge = $.maxAge;
        this.maxReuse = $.maxReuse;
        this.maxSize = $.maxSize;
        this.name = $.name;
        this.partition = $.partition;
        this.sharePools = $.sharePools;
        this.sourceMask = $.sourceMask;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProfileOneConnectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProfileOneConnectArgs $;

        public Builder() {
            $ = new ProfileOneConnectArgs();
        }

        public Builder(ProfileOneConnectArgs defaults) {
            $ = new ProfileOneConnectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultsFrom Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(@Nullable Output<String> defaultsFrom) {
            $.defaultsFrom = defaultsFrom;
            return this;
        }

        /**
         * @param defaultsFrom Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(String defaultsFrom) {
            return defaultsFrom(Output.of(defaultsFrom));
        }

        /**
         * @param idleTimeoutOverride Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
         * 
         * @return builder
         * 
         */
        public Builder idleTimeoutOverride(@Nullable Output<String> idleTimeoutOverride) {
            $.idleTimeoutOverride = idleTimeoutOverride;
            return this;
        }

        /**
         * @param idleTimeoutOverride Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
         * 
         * @return builder
         * 
         */
        public Builder idleTimeoutOverride(String idleTimeoutOverride) {
            return idleTimeoutOverride(Output.of(idleTimeoutOverride));
        }

        /**
         * @param limitType Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
         * 
         * @return builder
         * 
         */
        public Builder limitType(@Nullable Output<String> limitType) {
            $.limitType = limitType;
            return this;
        }

        /**
         * @param limitType Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
         * 
         * @return builder
         * 
         */
        public Builder limitType(String limitType) {
            return limitType(Output.of(limitType));
        }

        /**
         * @param maxAge Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
         * 
         * @return builder
         * 
         */
        public Builder maxAge(@Nullable Output<Integer> maxAge) {
            $.maxAge = maxAge;
            return this;
        }

        /**
         * @param maxAge Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
         * 
         * @return builder
         * 
         */
        public Builder maxAge(Integer maxAge) {
            return maxAge(Output.of(maxAge));
        }

        /**
         * @param maxReuse Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
         * 
         * @return builder
         * 
         */
        public Builder maxReuse(@Nullable Output<Integer> maxReuse) {
            $.maxReuse = maxReuse;
            return this;
        }

        /**
         * @param maxReuse Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
         * 
         * @return builder
         * 
         */
        public Builder maxReuse(Integer maxReuse) {
            return maxReuse(Output.of(maxReuse));
        }

        /**
         * @param maxSize Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(@Nullable Output<Integer> maxSize) {
            $.maxSize = maxSize;
            return this;
        }

        /**
         * @param maxSize Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(Integer maxSize) {
            return maxSize(Output.of(maxSize));
        }

        /**
         * @param name Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition Displays the administrative partition within which this profile resides
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Displays the administrative partition within which this profile resides
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param sharePools Specify if you want to share the pool, default value is `disabled`.
         * 
         * @return builder
         * 
         */
        public Builder sharePools(@Nullable Output<String> sharePools) {
            $.sharePools = sharePools;
            return this;
        }

        /**
         * @param sharePools Specify if you want to share the pool, default value is `disabled`.
         * 
         * @return builder
         * 
         */
        public Builder sharePools(String sharePools) {
            return sharePools(Output.of(sharePools));
        }

        /**
         * @param sourceMask Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
         * 
         * @return builder
         * 
         */
        public Builder sourceMask(@Nullable Output<String> sourceMask) {
            $.sourceMask = sourceMask;
            return this;
        }

        /**
         * @param sourceMask Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
         * 
         * @return builder
         * 
         */
        public Builder sourceMask(String sourceMask) {
            return sourceMask(Output.of(sourceMask));
        }

        public ProfileOneConnectArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("ProfileOneConnectArgs", "name");
            }
            return $;
        }
    }

}
