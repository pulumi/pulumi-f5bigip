// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
            if (tlsCertName == null) {
              throw new MissingRequiredPropertyException("FastHttpsAppTlsClientProfile", "tlsCertName");
            }
            this.tlsCertName = tlsCertName;
            return this;
        }
        @CustomType.Setter
        public Builder tlsKeyName(String tlsKeyName) {
            if (tlsKeyName == null) {
              throw new MissingRequiredPropertyException("FastHttpsAppTlsClientProfile", "tlsKeyName");
            }
            this.tlsKeyName = tlsKeyName;
            return this;
        }
        public FastHttpsAppTlsClientProfile build() {
            final var _resultValue = new FastHttpsAppTlsClientProfile();
            _resultValue.tlsCertName = tlsCertName;
            _resultValue.tlsKeyName = tlsKeyName;
            return _resultValue;
        }
    }
}