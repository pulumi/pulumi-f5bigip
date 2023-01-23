// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpsecProfileArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpsecProfileArgs Empty = new IpsecProfileArgs();

    /**
     * Specifies descriptive text that identifies the IPsec interface tunnel profile.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Specifies descriptive text that identifies the IPsec interface tunnel profile.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     * 
     */
    @Import(name="parentProfile")
    private @Nullable Output<String> parentProfile;

    /**
     * @return Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     * 
     */
    public Optional<Output<String>> parentProfile() {
        return Optional.ofNullable(this.parentProfile);
    }

    /**
     * Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     * 
     */
    @Import(name="trafficSelector")
    private @Nullable Output<String> trafficSelector;

    /**
     * @return Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     * 
     */
    public Optional<Output<String>> trafficSelector() {
        return Optional.ofNullable(this.trafficSelector);
    }

    private IpsecProfileArgs() {}

    private IpsecProfileArgs(IpsecProfileArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.parentProfile = $.parentProfile;
        this.trafficSelector = $.trafficSelector;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpsecProfileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpsecProfileArgs $;

        public Builder() {
            $ = new IpsecProfileArgs();
        }

        public Builder(IpsecProfileArgs defaults) {
            $ = new IpsecProfileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Specifies descriptive text that identifies the IPsec interface tunnel profile.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Specifies descriptive text that identifies the IPsec interface tunnel profile.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Displays the name of the IPsec interface tunnel profile,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentProfile Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
         * 
         * @return builder
         * 
         */
        public Builder parentProfile(@Nullable Output<String> parentProfile) {
            $.parentProfile = parentProfile;
            return this;
        }

        /**
         * @param parentProfile Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
         * 
         * @return builder
         * 
         */
        public Builder parentProfile(String parentProfile) {
            return parentProfile(Output.of(parentProfile));
        }

        /**
         * @param trafficSelector Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
         * 
         * @return builder
         * 
         */
        public Builder trafficSelector(@Nullable Output<String> trafficSelector) {
            $.trafficSelector = trafficSelector;
            return this;
        }

        /**
         * @param trafficSelector Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
         * 
         * @return builder
         * 
         */
        public Builder trafficSelector(String trafficSelector) {
            return trafficSelector(Output.of(trafficSelector));
        }

        public IpsecProfileArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}