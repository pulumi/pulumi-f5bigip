// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WafPolicyIpExceptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final WafPolicyIpExceptionArgs Empty = new WafPolicyIpExceptionArgs();

    /**
     * Specifies how the system responds to blocking requests sent from this IP address. Possible options [`always`, `never`, `policy-default`].
     * 
     */
    @Import(name="blockRequests")
    private @Nullable Output<String> blockRequests;

    /**
     * @return Specifies how the system responds to blocking requests sent from this IP address. Possible options [`always`, `never`, `policy-default`].
     * 
     */
    public Optional<Output<String>> blockRequests() {
        return Optional.ofNullable(this.blockRequests);
    }

    /**
     * Specifies the description of the policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Specifies the description of the policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies when enabled that the system considers this IP address legitimate and does not take it into account when performing brute force prevention.
     * 
     */
    @Import(name="ignoreAnomalies")
    private @Nullable Output<Boolean> ignoreAnomalies;

    /**
     * @return Specifies when enabled that the system considers this IP address legitimate and does not take it into account when performing brute force prevention.
     * 
     */
    public Optional<Output<Boolean>> ignoreAnomalies() {
        return Optional.ofNullable(this.ignoreAnomalies);
    }

    /**
     * Specifies when enabled that the system considers this IP address legitimate even if it is found in the IP Intelligence database (a database of questionable IP addresses).
     * 
     */
    @Import(name="ignoreIpreputation")
    private @Nullable Output<Boolean> ignoreIpreputation;

    /**
     * @return Specifies when enabled that the system considers this IP address legitimate even if it is found in the IP Intelligence database (a database of questionable IP addresses).
     * 
     */
    public Optional<Output<Boolean>> ignoreIpreputation() {
        return Optional.ofNullable(this.ignoreIpreputation);
    }

    /**
     * Specifies the IP address that you want the system to trust.
     * 
     */
    @Import(name="ipAddress", required=true)
    private Output<String> ipAddress;

    /**
     * @return Specifies the IP address that you want the system to trust.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }

    /**
     * Specifies the netmask of the exceptional IP address. This is an optional field.
     * 
     */
    @Import(name="ipMask", required=true)
    private Output<String> ipMask;

    /**
     * @return Specifies the netmask of the exceptional IP address. This is an optional field.
     * 
     */
    public Output<String> ipMask() {
        return this.ipMask;
    }

    /**
     * Specifies when enabled the Policy Builder considers traffic from this IP address as being safe.
     * 
     */
    @Import(name="trustedbyPolicybuilder")
    private @Nullable Output<Boolean> trustedbyPolicybuilder;

    /**
     * @return Specifies when enabled the Policy Builder considers traffic from this IP address as being safe.
     * 
     */
    public Optional<Output<Boolean>> trustedbyPolicybuilder() {
        return Optional.ofNullable(this.trustedbyPolicybuilder);
    }

    private WafPolicyIpExceptionArgs() {}

    private WafPolicyIpExceptionArgs(WafPolicyIpExceptionArgs $) {
        this.blockRequests = $.blockRequests;
        this.description = $.description;
        this.ignoreAnomalies = $.ignoreAnomalies;
        this.ignoreIpreputation = $.ignoreIpreputation;
        this.ipAddress = $.ipAddress;
        this.ipMask = $.ipMask;
        this.trustedbyPolicybuilder = $.trustedbyPolicybuilder;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WafPolicyIpExceptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WafPolicyIpExceptionArgs $;

        public Builder() {
            $ = new WafPolicyIpExceptionArgs();
        }

        public Builder(WafPolicyIpExceptionArgs defaults) {
            $ = new WafPolicyIpExceptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param blockRequests Specifies how the system responds to blocking requests sent from this IP address. Possible options [`always`, `never`, `policy-default`].
         * 
         * @return builder
         * 
         */
        public Builder blockRequests(@Nullable Output<String> blockRequests) {
            $.blockRequests = blockRequests;
            return this;
        }

        /**
         * @param blockRequests Specifies how the system responds to blocking requests sent from this IP address. Possible options [`always`, `never`, `policy-default`].
         * 
         * @return builder
         * 
         */
        public Builder blockRequests(String blockRequests) {
            return blockRequests(Output.of(blockRequests));
        }

        /**
         * @param description Specifies the description of the policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Specifies the description of the policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ignoreAnomalies Specifies when enabled that the system considers this IP address legitimate and does not take it into account when performing brute force prevention.
         * 
         * @return builder
         * 
         */
        public Builder ignoreAnomalies(@Nullable Output<Boolean> ignoreAnomalies) {
            $.ignoreAnomalies = ignoreAnomalies;
            return this;
        }

        /**
         * @param ignoreAnomalies Specifies when enabled that the system considers this IP address legitimate and does not take it into account when performing brute force prevention.
         * 
         * @return builder
         * 
         */
        public Builder ignoreAnomalies(Boolean ignoreAnomalies) {
            return ignoreAnomalies(Output.of(ignoreAnomalies));
        }

        /**
         * @param ignoreIpreputation Specifies when enabled that the system considers this IP address legitimate even if it is found in the IP Intelligence database (a database of questionable IP addresses).
         * 
         * @return builder
         * 
         */
        public Builder ignoreIpreputation(@Nullable Output<Boolean> ignoreIpreputation) {
            $.ignoreIpreputation = ignoreIpreputation;
            return this;
        }

        /**
         * @param ignoreIpreputation Specifies when enabled that the system considers this IP address legitimate even if it is found in the IP Intelligence database (a database of questionable IP addresses).
         * 
         * @return builder
         * 
         */
        public Builder ignoreIpreputation(Boolean ignoreIpreputation) {
            return ignoreIpreputation(Output.of(ignoreIpreputation));
        }

        /**
         * @param ipAddress Specifies the IP address that you want the system to trust.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress Specifies the IP address that you want the system to trust.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param ipMask Specifies the netmask of the exceptional IP address. This is an optional field.
         * 
         * @return builder
         * 
         */
        public Builder ipMask(Output<String> ipMask) {
            $.ipMask = ipMask;
            return this;
        }

        /**
         * @param ipMask Specifies the netmask of the exceptional IP address. This is an optional field.
         * 
         * @return builder
         * 
         */
        public Builder ipMask(String ipMask) {
            return ipMask(Output.of(ipMask));
        }

        /**
         * @param trustedbyPolicybuilder Specifies when enabled the Policy Builder considers traffic from this IP address as being safe.
         * 
         * @return builder
         * 
         */
        public Builder trustedbyPolicybuilder(@Nullable Output<Boolean> trustedbyPolicybuilder) {
            $.trustedbyPolicybuilder = trustedbyPolicybuilder;
            return this;
        }

        /**
         * @param trustedbyPolicybuilder Specifies when enabled the Policy Builder considers traffic from this IP address as being safe.
         * 
         * @return builder
         * 
         */
        public Builder trustedbyPolicybuilder(Boolean trustedbyPolicybuilder) {
            return trustedbyPolicybuilder(Output.of(trustedbyPolicybuilder));
        }

        public WafPolicyIpExceptionArgs build() {
            $.ipAddress = Objects.requireNonNull($.ipAddress, "expected parameter 'ipAddress' to be non-null");
            $.ipMask = Objects.requireNonNull($.ipMask, "expected parameter 'ipMask' to be non-null");
            return $;
        }
    }

}