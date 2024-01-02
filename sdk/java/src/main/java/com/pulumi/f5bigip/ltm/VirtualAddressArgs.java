// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualAddressArgs Empty = new VirtualAddressArgs();

    /**
     * Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
     * 
     */
    @Import(name="advertizeRoute")
    private @Nullable Output<String> advertizeRoute;

    /**
     * @return Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
     * 
     */
    public Optional<Output<String>> advertizeRoute() {
        return Optional.ofNullable(this.advertizeRoute);
    }

    /**
     * Enable or disable ARP for the virtual address
     * 
     */
    @Import(name="arp")
    private @Nullable Output<Boolean> arp;

    /**
     * @return Enable or disable ARP for the virtual address
     * 
     */
    public Optional<Output<Boolean>> arp() {
        return Optional.ofNullable(this.arp);
    }

    /**
     * Automatically delete the virtual address with the virtual server
     * 
     */
    @Import(name="autoDelete")
    private @Nullable Output<Boolean> autoDelete;

    /**
     * @return Automatically delete the virtual address with the virtual server
     * 
     */
    public Optional<Output<Boolean>> autoDelete() {
        return Optional.ofNullable(this.autoDelete);
    }

    /**
     * Max number of connections for virtual address
     * 
     */
    @Import(name="connLimit")
    private @Nullable Output<Integer> connLimit;

    /**
     * @return Max number of connections for virtual address
     * 
     */
    public Optional<Output<Integer>> connLimit() {
        return Optional.ofNullable(this.connLimit);
    }

    /**
     * Enable or disable the virtual address
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Enable or disable the virtual address
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     * 
     */
    @Import(name="icmpEcho")
    private @Nullable Output<String> icmpEcho;

    /**
     * @return Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
     * 
     */
    public Optional<Output<String>> icmpEcho() {
        return Optional.ofNullable(this.icmpEcho);
    }

    /**
     * Name of the virtual address
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the virtual address
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Specify the partition and traffic group
     * 
     */
    @Import(name="trafficGroup")
    private @Nullable Output<String> trafficGroup;

    /**
     * @return Specify the partition and traffic group
     * 
     */
    public Optional<Output<String>> trafficGroup() {
        return Optional.ofNullable(this.trafficGroup);
    }

    private VirtualAddressArgs() {}

    private VirtualAddressArgs(VirtualAddressArgs $) {
        this.advertizeRoute = $.advertizeRoute;
        this.arp = $.arp;
        this.autoDelete = $.autoDelete;
        this.connLimit = $.connLimit;
        this.enabled = $.enabled;
        this.icmpEcho = $.icmpEcho;
        this.name = $.name;
        this.trafficGroup = $.trafficGroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualAddressArgs $;

        public Builder() {
            $ = new VirtualAddressArgs();
        }

        public Builder(VirtualAddressArgs defaults) {
            $ = new VirtualAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param advertizeRoute Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
         * 
         * @return builder
         * 
         */
        public Builder advertizeRoute(@Nullable Output<String> advertizeRoute) {
            $.advertizeRoute = advertizeRoute;
            return this;
        }

        /**
         * @param advertizeRoute Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
         * 
         * @return builder
         * 
         */
        public Builder advertizeRoute(String advertizeRoute) {
            return advertizeRoute(Output.of(advertizeRoute));
        }

        /**
         * @param arp Enable or disable ARP for the virtual address
         * 
         * @return builder
         * 
         */
        public Builder arp(@Nullable Output<Boolean> arp) {
            $.arp = arp;
            return this;
        }

        /**
         * @param arp Enable or disable ARP for the virtual address
         * 
         * @return builder
         * 
         */
        public Builder arp(Boolean arp) {
            return arp(Output.of(arp));
        }

        /**
         * @param autoDelete Automatically delete the virtual address with the virtual server
         * 
         * @return builder
         * 
         */
        public Builder autoDelete(@Nullable Output<Boolean> autoDelete) {
            $.autoDelete = autoDelete;
            return this;
        }

        /**
         * @param autoDelete Automatically delete the virtual address with the virtual server
         * 
         * @return builder
         * 
         */
        public Builder autoDelete(Boolean autoDelete) {
            return autoDelete(Output.of(autoDelete));
        }

        /**
         * @param connLimit Max number of connections for virtual address
         * 
         * @return builder
         * 
         */
        public Builder connLimit(@Nullable Output<Integer> connLimit) {
            $.connLimit = connLimit;
            return this;
        }

        /**
         * @param connLimit Max number of connections for virtual address
         * 
         * @return builder
         * 
         */
        public Builder connLimit(Integer connLimit) {
            return connLimit(Output.of(connLimit));
        }

        /**
         * @param enabled Enable or disable the virtual address
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enable or disable the virtual address
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param icmpEcho Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
         * 
         * @return builder
         * 
         */
        public Builder icmpEcho(@Nullable Output<String> icmpEcho) {
            $.icmpEcho = icmpEcho;
            return this;
        }

        /**
         * @param icmpEcho Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
         * 
         * @return builder
         * 
         */
        public Builder icmpEcho(String icmpEcho) {
            return icmpEcho(Output.of(icmpEcho));
        }

        /**
         * @param name Name of the virtual address
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the virtual address
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param trafficGroup Specify the partition and traffic group
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(@Nullable Output<String> trafficGroup) {
            $.trafficGroup = trafficGroup;
            return this;
        }

        /**
         * @param trafficGroup Specify the partition and traffic group
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(String trafficGroup) {
            return trafficGroup(Output.of(trafficGroup));
        }

        public VirtualAddressArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("VirtualAddressArgs", "name");
            }
            return $;
        }
    }

}
