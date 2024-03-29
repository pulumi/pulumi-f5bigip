// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVWanConfigResult {
    private String azureVwanName;
    private String azureVwanResourcegroup;
    private String azureVwanVpnsite;
    /**
     * @return (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
     * 
     */
    private String bigipGwIp;
    /**
     * @return (type `string`) Provides IP Address space used on vWAN Hub.
     * 
     */
    private String hubAddressSpace;
    /**
     * @return (type `list`) Provides Subnets connected to vWAN Hub.
     * 
     */
    private List<String> hubConnectedSubnets;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
     * 
     */
    private String presharedKey;
    /**
     * @return (type `list`) Provides vWAN Gateway Address for IPSec End point
     * 
     */
    private List<String> vwanGwAddresses;

    private GetVWanConfigResult() {}
    public String azureVwanName() {
        return this.azureVwanName;
    }
    public String azureVwanResourcegroup() {
        return this.azureVwanResourcegroup;
    }
    public String azureVwanVpnsite() {
        return this.azureVwanVpnsite;
    }
    /**
     * @return (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
     * 
     */
    public String bigipGwIp() {
        return this.bigipGwIp;
    }
    /**
     * @return (type `string`) Provides IP Address space used on vWAN Hub.
     * 
     */
    public String hubAddressSpace() {
        return this.hubAddressSpace;
    }
    /**
     * @return (type `list`) Provides Subnets connected to vWAN Hub.
     * 
     */
    public List<String> hubConnectedSubnets() {
        return this.hubConnectedSubnets;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
     * 
     */
    public String presharedKey() {
        return this.presharedKey;
    }
    /**
     * @return (type `list`) Provides vWAN Gateway Address for IPSec End point
     * 
     */
    public List<String> vwanGwAddresses() {
        return this.vwanGwAddresses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVWanConfigResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String azureVwanName;
        private String azureVwanResourcegroup;
        private String azureVwanVpnsite;
        private String bigipGwIp;
        private String hubAddressSpace;
        private List<String> hubConnectedSubnets;
        private String id;
        private String presharedKey;
        private List<String> vwanGwAddresses;
        public Builder() {}
        public Builder(GetVWanConfigResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.azureVwanName = defaults.azureVwanName;
    	      this.azureVwanResourcegroup = defaults.azureVwanResourcegroup;
    	      this.azureVwanVpnsite = defaults.azureVwanVpnsite;
    	      this.bigipGwIp = defaults.bigipGwIp;
    	      this.hubAddressSpace = defaults.hubAddressSpace;
    	      this.hubConnectedSubnets = defaults.hubConnectedSubnets;
    	      this.id = defaults.id;
    	      this.presharedKey = defaults.presharedKey;
    	      this.vwanGwAddresses = defaults.vwanGwAddresses;
        }

        @CustomType.Setter
        public Builder azureVwanName(String azureVwanName) {
            if (azureVwanName == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "azureVwanName");
            }
            this.azureVwanName = azureVwanName;
            return this;
        }
        @CustomType.Setter
        public Builder azureVwanResourcegroup(String azureVwanResourcegroup) {
            if (azureVwanResourcegroup == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "azureVwanResourcegroup");
            }
            this.azureVwanResourcegroup = azureVwanResourcegroup;
            return this;
        }
        @CustomType.Setter
        public Builder azureVwanVpnsite(String azureVwanVpnsite) {
            if (azureVwanVpnsite == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "azureVwanVpnsite");
            }
            this.azureVwanVpnsite = azureVwanVpnsite;
            return this;
        }
        @CustomType.Setter
        public Builder bigipGwIp(String bigipGwIp) {
            if (bigipGwIp == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "bigipGwIp");
            }
            this.bigipGwIp = bigipGwIp;
            return this;
        }
        @CustomType.Setter
        public Builder hubAddressSpace(String hubAddressSpace) {
            if (hubAddressSpace == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "hubAddressSpace");
            }
            this.hubAddressSpace = hubAddressSpace;
            return this;
        }
        @CustomType.Setter
        public Builder hubConnectedSubnets(List<String> hubConnectedSubnets) {
            if (hubConnectedSubnets == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "hubConnectedSubnets");
            }
            this.hubConnectedSubnets = hubConnectedSubnets;
            return this;
        }
        public Builder hubConnectedSubnets(String... hubConnectedSubnets) {
            return hubConnectedSubnets(List.of(hubConnectedSubnets));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder presharedKey(String presharedKey) {
            if (presharedKey == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "presharedKey");
            }
            this.presharedKey = presharedKey;
            return this;
        }
        @CustomType.Setter
        public Builder vwanGwAddresses(List<String> vwanGwAddresses) {
            if (vwanGwAddresses == null) {
              throw new MissingRequiredPropertyException("GetVWanConfigResult", "vwanGwAddresses");
            }
            this.vwanGwAddresses = vwanGwAddresses;
            return this;
        }
        public Builder vwanGwAddresses(String... vwanGwAddresses) {
            return vwanGwAddresses(List.of(vwanGwAddresses));
        }
        public GetVWanConfigResult build() {
            final var _resultValue = new GetVWanConfigResult();
            _resultValue.azureVwanName = azureVwanName;
            _resultValue.azureVwanResourcegroup = azureVwanResourcegroup;
            _resultValue.azureVwanVpnsite = azureVwanVpnsite;
            _resultValue.bigipGwIp = bigipGwIp;
            _resultValue.hubAddressSpace = hubAddressSpace;
            _resultValue.hubConnectedSubnets = hubConnectedSubnets;
            _resultValue.id = id;
            _resultValue.presharedKey = presharedKey;
            _resultValue.vwanGwAddresses = vwanGwAddresses;
            return _resultValue;
        }
    }
}
