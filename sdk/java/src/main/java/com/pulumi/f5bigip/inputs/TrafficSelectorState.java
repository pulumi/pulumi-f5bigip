// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TrafficSelectorState extends com.pulumi.resources.ResourceArgs {

    public static final TrafficSelectorState Empty = new TrafficSelectorState();

    /**
     * Description of the traffic selector.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the traffic selector.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
     * 
     */
    @Import(name="destinationAddress")
    private @Nullable Output<String> destinationAddress;

    /**
     * @return Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
     * 
     */
    public Optional<Output<String>> destinationAddress() {
        return Optional.ofNullable(this.destinationAddress);
    }

    /**
     * Specifies the IP port used by the application. The default value is `All Ports (0)`
     * 
     */
    @Import(name="destinationPort")
    private @Nullable Output<Integer> destinationPort;

    /**
     * @return Specifies the IP port used by the application. The default value is `All Ports (0)`
     * 
     */
    public Optional<Output<Integer>> destinationPort() {
        return Optional.ofNullable(this.destinationPort);
    }

    /**
     * Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
     * 
     */
    @Import(name="ipProtocol")
    private @Nullable Output<Integer> ipProtocol;

    /**
     * @return Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
     * 
     */
    public Optional<Output<Integer>> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    /**
     * Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
     * 
     */
    @Import(name="ipsecPolicy")
    private @Nullable Output<String> ipsecPolicy;

    /**
     * @return Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
     * 
     */
    public Optional<Output<String>> ipsecPolicy() {
        return Optional.ofNullable(this.ipsecPolicy);
    }

    /**
     * Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
     * When creating a new traffic selector, if this parameter is not specified, the default is `last`
     * 
     */
    @Import(name="order")
    private @Nullable Output<Integer> order;

    /**
     * @return Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
     * When creating a new traffic selector, if this parameter is not specified, the default is `last`
     * 
     */
    public Optional<Output<Integer>> order() {
        return Optional.ofNullable(this.order);
    }

    /**
     * Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
     * 
     */
    @Import(name="sourceAddress")
    private @Nullable Output<String> sourceAddress;

    /**
     * @return Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
     * 
     */
    public Optional<Output<String>> sourceAddress() {
        return Optional.ofNullable(this.sourceAddress);
    }

    /**
     * Specifies the IP port used by the application. The default value is `All Ports (0)`.
     * 
     */
    @Import(name="sourcePort")
    private @Nullable Output<Integer> sourcePort;

    /**
     * @return Specifies the IP port used by the application. The default value is `All Ports (0)`.
     * 
     */
    public Optional<Output<Integer>> sourcePort() {
        return Optional.ofNullable(this.sourcePort);
    }

    private TrafficSelectorState() {}

    private TrafficSelectorState(TrafficSelectorState $) {
        this.description = $.description;
        this.destinationAddress = $.destinationAddress;
        this.destinationPort = $.destinationPort;
        this.direction = $.direction;
        this.ipProtocol = $.ipProtocol;
        this.ipsecPolicy = $.ipsecPolicy;
        this.name = $.name;
        this.order = $.order;
        this.sourceAddress = $.sourceAddress;
        this.sourcePort = $.sourcePort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrafficSelectorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrafficSelectorState $;

        public Builder() {
            $ = new TrafficSelectorState();
        }

        public Builder(TrafficSelectorState defaults) {
            $ = new TrafficSelectorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the traffic selector.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the traffic selector.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param destinationAddress Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddress(@Nullable Output<String> destinationAddress) {
            $.destinationAddress = destinationAddress;
            return this;
        }

        /**
         * @param destinationAddress Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddress(String destinationAddress) {
            return destinationAddress(Output.of(destinationAddress));
        }

        /**
         * @param destinationPort Specifies the IP port used by the application. The default value is `All Ports (0)`
         * 
         * @return builder
         * 
         */
        public Builder destinationPort(@Nullable Output<Integer> destinationPort) {
            $.destinationPort = destinationPort;
            return this;
        }

        /**
         * @param destinationPort Specifies the IP port used by the application. The default value is `All Ports (0)`
         * 
         * @return builder
         * 
         */
        public Builder destinationPort(Integer destinationPort) {
            return destinationPort(Output.of(destinationPort));
        }

        /**
         * @param direction Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
         * 
         * @return builder
         * 
         */
        public Builder direction(@Nullable Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param ipProtocol Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(@Nullable Output<Integer> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param ipProtocol Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(Integer ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        /**
         * @param ipsecPolicy Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
         * 
         * @return builder
         * 
         */
        public Builder ipsecPolicy(@Nullable Output<String> ipsecPolicy) {
            $.ipsecPolicy = ipsecPolicy;
            return this;
        }

        /**
         * @param ipsecPolicy Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
         * 
         * @return builder
         * 
         */
        public Builder ipsecPolicy(String ipsecPolicy) {
            return ipsecPolicy(Output.of(ipsecPolicy));
        }

        /**
         * @param name Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the IPSec traffic-selector,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param order Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
         * When creating a new traffic selector, if this parameter is not specified, the default is `last`
         * 
         * @return builder
         * 
         */
        public Builder order(@Nullable Output<Integer> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
         * When creating a new traffic selector, if this parameter is not specified, the default is `last`
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        /**
         * @param sourceAddress Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder sourceAddress(@Nullable Output<String> sourceAddress) {
            $.sourceAddress = sourceAddress;
            return this;
        }

        /**
         * @param sourceAddress Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder sourceAddress(String sourceAddress) {
            return sourceAddress(Output.of(sourceAddress));
        }

        /**
         * @param sourcePort Specifies the IP port used by the application. The default value is `All Ports (0)`.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(@Nullable Output<Integer> sourcePort) {
            $.sourcePort = sourcePort;
            return this;
        }

        /**
         * @param sourcePort Specifies the IP port used by the application. The default value is `All Ports (0)`.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(Integer sourcePort) {
            return sourcePort(Output.of(sourcePort));
        }

        public TrafficSelectorState build() {
            return $;
        }
    }

}
