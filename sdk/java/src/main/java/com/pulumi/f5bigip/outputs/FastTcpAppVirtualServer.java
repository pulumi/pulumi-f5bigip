// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FastTcpAppVirtualServer {
    /**
     * @return IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     * 
     */
    private String ip;
    /**
     * @return Port number to used for accessing virtual server/application
     * 
     */
    private Integer port;

    private FastTcpAppVirtualServer() {}
    /**
     * @return IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return Port number to used for accessing virtual server/application
     * 
     */
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FastTcpAppVirtualServer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private Integer port;
        public Builder() {}
        public Builder(FastTcpAppVirtualServer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            if (ip == null) {
              throw new MissingRequiredPropertyException("FastTcpAppVirtualServer", "ip");
            }
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            if (port == null) {
              throw new MissingRequiredPropertyException("FastTcpAppVirtualServer", "port");
            }
            this.port = port;
            return this;
        }
        public FastTcpAppVirtualServer build() {
            final var _resultValue = new FastTcpAppVirtualServer();
            _resultValue.ip = ip;
            _resultValue.port = port;
            return _resultValue;
        }
    }
}
