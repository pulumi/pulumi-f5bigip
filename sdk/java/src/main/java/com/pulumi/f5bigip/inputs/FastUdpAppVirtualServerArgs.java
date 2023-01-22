// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class FastUdpAppVirtualServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final FastUdpAppVirtualServerArgs Empty = new FastUdpAppVirtualServerArgs();

    /**
     * IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * Port number to used for accessing virtual server/application
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return Port number to used for accessing virtual server/application
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    private FastUdpAppVirtualServerArgs() {}

    private FastUdpAppVirtualServerArgs(FastUdpAppVirtualServerArgs $) {
        this.ip = $.ip;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FastUdpAppVirtualServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FastUdpAppVirtualServerArgs $;

        public Builder() {
            $ = new FastUdpAppVirtualServerArgs();
        }

        public Builder(FastUdpAppVirtualServerArgs defaults) {
            $ = new FastUdpAppVirtualServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param port Port number to used for accessing virtual server/application
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port number to used for accessing virtual server/application
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public FastUdpAppVirtualServerArgs build() {
            $.ip = Objects.requireNonNull($.ip, "expected parameter 'ip' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            return $;
        }
    }

}
