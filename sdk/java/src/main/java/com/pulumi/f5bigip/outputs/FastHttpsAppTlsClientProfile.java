// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FastHttpsAppTlsClientProfile {
    /**
     * @return Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
     * 
     */
    private String tlsCertName;
    /**
     * @return Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
     * 
     */
    private String tlsKeyName;

    private FastHttpsAppTlsClientProfile() {}
    /**
     * @return Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
     * 
     */
    public String tlsCertName() {
        return this.tlsCertName;
    }
    /**
     * @return Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
     * 
     */
    public String tlsKeyName() {
        return this.tlsKeyName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FastHttpsAppTlsClientProfile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String tlsCertName;
        private String tlsKeyName;
        public Builder() {}
        public Builder(FastHttpsAppTlsClientProfile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.tlsCertName = defaults.tlsCertName;
    	      this.tlsKeyName = defaults.tlsKeyName;
        }

        @CustomType.Setter
        public Builder tlsCertName(String tlsCertName) {
            this.tlsCertName = Objects.requireNonNull(tlsCertName);
            return this;
        }
        @CustomType.Setter
        public Builder tlsKeyName(String tlsKeyName) {
            this.tlsKeyName = Objects.requireNonNull(tlsKeyName);
            return this;
        }
        public FastHttpsAppTlsClientProfile build() {
            final var o = new FastHttpsAppTlsClientProfile();
            o.tlsCertName = tlsCertName;
            o.tlsKeyName = tlsKeyName;
            return o;
        }
    }
}
