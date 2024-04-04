// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.net.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SelfIpState extends com.pulumi.resources.ResourceArgs {

    public static final SelfIpState Empty = new SelfIpState();

    /**
     * The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * Name of the selfip
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the selfip
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies the port lockdown, defaults to `Allow None` if not specified.
     * 
     */
    @Import(name="portLockdowns")
    private @Nullable Output<List<String>> portLockdowns;

    /**
     * @return Specifies the port lockdown, defaults to `Allow None` if not specified.
     * 
     */
    public Optional<Output<List<String>>> portLockdowns() {
        return Optional.ofNullable(this.portLockdowns);
    }

    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     * 
     */
    @Import(name="trafficGroup")
    private @Nullable Output<String> trafficGroup;

    /**
     * @return Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     * 
     */
    public Optional<Output<String>> trafficGroup() {
        return Optional.ofNullable(this.trafficGroup);
    }

    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     * 
     */
    @Import(name="vlan")
    private @Nullable Output<String> vlan;

    /**
     * @return Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     * 
     */
    public Optional<Output<String>> vlan() {
        return Optional.ofNullable(this.vlan);
    }

    private SelfIpState() {}

    private SelfIpState(SelfIpState $) {
        this.ip = $.ip;
        this.name = $.name;
        this.portLockdowns = $.portLockdowns;
        this.trafficGroup = $.trafficGroup;
        this.vlan = $.vlan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SelfIpState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SelfIpState $;

        public Builder() {
            $ = new SelfIpState();
        }

        public Builder(SelfIpState defaults) {
            $ = new SelfIpState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param name Name of the selfip
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the selfip
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portLockdowns Specifies the port lockdown, defaults to `Allow None` if not specified.
         * 
         * @return builder
         * 
         */
        public Builder portLockdowns(@Nullable Output<List<String>> portLockdowns) {
            $.portLockdowns = portLockdowns;
            return this;
        }

        /**
         * @param portLockdowns Specifies the port lockdown, defaults to `Allow None` if not specified.
         * 
         * @return builder
         * 
         */
        public Builder portLockdowns(List<String> portLockdowns) {
            return portLockdowns(Output.of(portLockdowns));
        }

        /**
         * @param portLockdowns Specifies the port lockdown, defaults to `Allow None` if not specified.
         * 
         * @return builder
         * 
         */
        public Builder portLockdowns(String... portLockdowns) {
            return portLockdowns(List.of(portLockdowns));
        }

        /**
         * @param trafficGroup Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(@Nullable Output<String> trafficGroup) {
            $.trafficGroup = trafficGroup;
            return this;
        }

        /**
         * @param trafficGroup Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(String trafficGroup) {
            return trafficGroup(Output.of(trafficGroup));
        }

        /**
         * @param vlan Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
         * 
         * @return builder
         * 
         */
        public Builder vlan(@Nullable Output<String> vlan) {
            $.vlan = vlan;
            return this;
        }

        /**
         * @param vlan Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
         * 
         * @return builder
         * 
         */
        public Builder vlan(String vlan) {
            return vlan(Output.of(vlan));
        }

        public SelfIpState build() {
            return $;
        }
    }

}
