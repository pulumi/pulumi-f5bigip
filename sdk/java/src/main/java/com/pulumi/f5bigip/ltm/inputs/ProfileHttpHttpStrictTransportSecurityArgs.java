// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProfileHttpHttpStrictTransportSecurityArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProfileHttpHttpStrictTransportSecurityArgs Empty = new ProfileHttpHttpStrictTransportSecurityArgs();

    /**
     * Specifies whether to include the includeSubdomains directive in the HSTS header. The default is enabled. If no string is specified, then default value will be assigned.
     * 
     */
    @Import(name="includeSubdomains")
    private @Nullable Output<String> includeSubdomains;

    /**
     * @return Specifies whether to include the includeSubdomains directive in the HSTS header. The default is enabled. If no string is specified, then default value will be assigned.
     * 
     */
    public Optional<Output<String>> includeSubdomains() {
        return Optional.ofNullable(this.includeSubdomains);
    }

    /**
     * Specifies the maximum age to assume the connection should remain secure. The default is 16070400 seconds. If no value is specified, then default value will be assigned.
     * 
     */
    @Import(name="maximumAge")
    private @Nullable Output<Integer> maximumAge;

    /**
     * @return Specifies the maximum age to assume the connection should remain secure. The default is 16070400 seconds. If no value is specified, then default value will be assigned.
     * 
     */
    public Optional<Output<Integer>> maximumAge() {
        return Optional.ofNullable(this.maximumAge);
    }

    /**
     * Specifies whether to include the HSTS response header. The default is disabled.If no string is specified, then default value will be assigned.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Specifies whether to include the HSTS response header. The default is disabled.If no string is specified, then default value will be assigned.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Specifies whether to include the preload directive in the HSTS header. The default is disabled. If no string is specified, then default value will be assigned.
     * 
     */
    @Import(name="preload")
    private @Nullable Output<String> preload;

    /**
     * @return Specifies whether to include the preload directive in the HSTS header. The default is disabled. If no string is specified, then default value will be assigned.
     * 
     */
    public Optional<Output<String>> preload() {
        return Optional.ofNullable(this.preload);
    }

    private ProfileHttpHttpStrictTransportSecurityArgs() {}

    private ProfileHttpHttpStrictTransportSecurityArgs(ProfileHttpHttpStrictTransportSecurityArgs $) {
        this.includeSubdomains = $.includeSubdomains;
        this.maximumAge = $.maximumAge;
        this.mode = $.mode;
        this.preload = $.preload;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProfileHttpHttpStrictTransportSecurityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProfileHttpHttpStrictTransportSecurityArgs $;

        public Builder() {
            $ = new ProfileHttpHttpStrictTransportSecurityArgs();
        }

        public Builder(ProfileHttpHttpStrictTransportSecurityArgs defaults) {
            $ = new ProfileHttpHttpStrictTransportSecurityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param includeSubdomains Specifies whether to include the includeSubdomains directive in the HSTS header. The default is enabled. If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder includeSubdomains(@Nullable Output<String> includeSubdomains) {
            $.includeSubdomains = includeSubdomains;
            return this;
        }

        /**
         * @param includeSubdomains Specifies whether to include the includeSubdomains directive in the HSTS header. The default is enabled. If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder includeSubdomains(String includeSubdomains) {
            return includeSubdomains(Output.of(includeSubdomains));
        }

        /**
         * @param maximumAge Specifies the maximum age to assume the connection should remain secure. The default is 16070400 seconds. If no value is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder maximumAge(@Nullable Output<Integer> maximumAge) {
            $.maximumAge = maximumAge;
            return this;
        }

        /**
         * @param maximumAge Specifies the maximum age to assume the connection should remain secure. The default is 16070400 seconds. If no value is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder maximumAge(Integer maximumAge) {
            return maximumAge(Output.of(maximumAge));
        }

        /**
         * @param mode Specifies whether to include the HSTS response header. The default is disabled.If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Specifies whether to include the HSTS response header. The default is disabled.If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param preload Specifies whether to include the preload directive in the HSTS header. The default is disabled. If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder preload(@Nullable Output<String> preload) {
            $.preload = preload;
            return this;
        }

        /**
         * @param preload Specifies whether to include the preload directive in the HSTS header. The default is disabled. If no string is specified, then default value will be assigned.
         * 
         * @return builder
         * 
         */
        public Builder preload(String preload) {
            return preload(Output.of(preload));
        }

        public ProfileHttpHttpStrictTransportSecurityArgs build() {
            return $;
        }
    }

}
