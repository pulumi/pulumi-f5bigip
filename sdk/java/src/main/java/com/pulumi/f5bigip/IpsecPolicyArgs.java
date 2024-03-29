// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpsecPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpsecPolicyArgs Empty = new IpsecPolicyArgs();

    /**
     * Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     * 
     */
    @Import(name="authAlgorithm")
    private @Nullable Output<String> authAlgorithm;

    /**
     * @return Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     * 
     */
    public Optional<Output<String>> authAlgorithm() {
        return Optional.ofNullable(this.authAlgorithm);
    }

    /**
     * Description of the IPSec policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the IPSec policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     * 
     */
    @Import(name="encryptAlgorithm")
    private @Nullable Output<String> encryptAlgorithm;

    /**
     * @return Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     * 
     */
    public Optional<Output<String>> encryptAlgorithm() {
        return Optional.ofNullable(this.encryptAlgorithm);
    }

    /**
     * Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
     * 
     */
    @Import(name="ipcomp")
    private @Nullable Output<String> ipcomp;

    /**
     * @return Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
     * 
     */
    public Optional<Output<String>> ipcomp() {
        return Optional.ofNullable(this.ipcomp);
    }

    /**
     * Specifies the length of time before the IKE security association expires, in kilobytes.
     * 
     */
    @Import(name="kbLifetime")
    private @Nullable Output<Integer> kbLifetime;

    /**
     * @return Specifies the length of time before the IKE security association expires, in kilobytes.
     * 
     */
    public Optional<Output<Integer>> kbLifetime() {
        return Optional.ofNullable(this.kbLifetime);
    }

    /**
     * Specifies the length of time before the IKE security association expires, in minutes.
     * 
     */
    @Import(name="lifetime")
    private @Nullable Output<Integer> lifetime;

    /**
     * @return Specifies the length of time before the IKE security association expires, in minutes.
     * 
     */
    public Optional<Output<Integer>> lifetime() {
        return Optional.ofNullable(this.lifetime);
    }

    /**
     * Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     * 
     */
    @Import(name="perfectForwardSecrecy")
    private @Nullable Output<String> perfectForwardSecrecy;

    /**
     * @return Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     * 
     */
    public Optional<Output<String>> perfectForwardSecrecy() {
        return Optional.ofNullable(this.perfectForwardSecrecy);
    }

    /**
     * Specifies the IPsec protocol. Valid choices are: `ah, esp`
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Specifies the IPsec protocol. Valid choices are: `ah, esp`
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    @Import(name="tunnelLocalAddress")
    private @Nullable Output<String> tunnelLocalAddress;

    /**
     * @return Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    public Optional<Output<String>> tunnelLocalAddress() {
        return Optional.ofNullable(this.tunnelLocalAddress);
    }

    /**
     * Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    @Import(name="tunnelRemoteAddress")
    private @Nullable Output<String> tunnelRemoteAddress;

    /**
     * @return Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     * 
     */
    public Optional<Output<String>> tunnelRemoteAddress() {
        return Optional.ofNullable(this.tunnelRemoteAddress);
    }

    private IpsecPolicyArgs() {}

    private IpsecPolicyArgs(IpsecPolicyArgs $) {
        this.authAlgorithm = $.authAlgorithm;
        this.description = $.description;
        this.encryptAlgorithm = $.encryptAlgorithm;
        this.ipcomp = $.ipcomp;
        this.kbLifetime = $.kbLifetime;
        this.lifetime = $.lifetime;
        this.mode = $.mode;
        this.name = $.name;
        this.perfectForwardSecrecy = $.perfectForwardSecrecy;
        this.protocol = $.protocol;
        this.tunnelLocalAddress = $.tunnelLocalAddress;
        this.tunnelRemoteAddress = $.tunnelRemoteAddress;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpsecPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpsecPolicyArgs $;

        public Builder() {
            $ = new IpsecPolicyArgs();
        }

        public Builder(IpsecPolicyArgs defaults) {
            $ = new IpsecPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authAlgorithm Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
         * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
         * 
         * @return builder
         * 
         */
        public Builder authAlgorithm(@Nullable Output<String> authAlgorithm) {
            $.authAlgorithm = authAlgorithm;
            return this;
        }

        /**
         * @param authAlgorithm Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
         * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
         * 
         * @return builder
         * 
         */
        public Builder authAlgorithm(String authAlgorithm) {
            return authAlgorithm(Output.of(authAlgorithm));
        }

        /**
         * @param description Description of the IPSec policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the IPSec policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param encryptAlgorithm Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
         * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
         * 
         * @return builder
         * 
         */
        public Builder encryptAlgorithm(@Nullable Output<String> encryptAlgorithm) {
            $.encryptAlgorithm = encryptAlgorithm;
            return this;
        }

        /**
         * @param encryptAlgorithm Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
         * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
         * 
         * @return builder
         * 
         */
        public Builder encryptAlgorithm(String encryptAlgorithm) {
            return encryptAlgorithm(Output.of(encryptAlgorithm));
        }

        /**
         * @param ipcomp Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
         * 
         * @return builder
         * 
         */
        public Builder ipcomp(@Nullable Output<String> ipcomp) {
            $.ipcomp = ipcomp;
            return this;
        }

        /**
         * @param ipcomp Specifies whether to use IPComp encapsulation. Valid choices are: `none&#34;, null&#34;, deflate`
         * 
         * @return builder
         * 
         */
        public Builder ipcomp(String ipcomp) {
            return ipcomp(Output.of(ipcomp));
        }

        /**
         * @param kbLifetime Specifies the length of time before the IKE security association expires, in kilobytes.
         * 
         * @return builder
         * 
         */
        public Builder kbLifetime(@Nullable Output<Integer> kbLifetime) {
            $.kbLifetime = kbLifetime;
            return this;
        }

        /**
         * @param kbLifetime Specifies the length of time before the IKE security association expires, in kilobytes.
         * 
         * @return builder
         * 
         */
        public Builder kbLifetime(Integer kbLifetime) {
            return kbLifetime(Output.of(kbLifetime));
        }

        /**
         * @param lifetime Specifies the length of time before the IKE security association expires, in minutes.
         * 
         * @return builder
         * 
         */
        public Builder lifetime(@Nullable Output<Integer> lifetime) {
            $.lifetime = lifetime;
            return this;
        }

        /**
         * @param lifetime Specifies the length of time before the IKE security association expires, in minutes.
         * 
         * @return builder
         * 
         */
        public Builder lifetime(Integer lifetime) {
            return lifetime(Output.of(lifetime));
        }

        /**
         * @param mode Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the IPSec policy,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param perfectForwardSecrecy Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
         * modp4096, modp6144, modp8192`
         * 
         * @return builder
         * 
         */
        public Builder perfectForwardSecrecy(@Nullable Output<String> perfectForwardSecrecy) {
            $.perfectForwardSecrecy = perfectForwardSecrecy;
            return this;
        }

        /**
         * @param perfectForwardSecrecy Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
         * modp4096, modp6144, modp8192`
         * 
         * @return builder
         * 
         */
        public Builder perfectForwardSecrecy(String perfectForwardSecrecy) {
            return perfectForwardSecrecy(Output.of(perfectForwardSecrecy));
        }

        /**
         * @param protocol Specifies the IPsec protocol. Valid choices are: `ah, esp`
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Specifies the IPsec protocol. Valid choices are: `ah, esp`
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param tunnelLocalAddress Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
         * 
         * @return builder
         * 
         */
        public Builder tunnelLocalAddress(@Nullable Output<String> tunnelLocalAddress) {
            $.tunnelLocalAddress = tunnelLocalAddress;
            return this;
        }

        /**
         * @param tunnelLocalAddress Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
         * 
         * @return builder
         * 
         */
        public Builder tunnelLocalAddress(String tunnelLocalAddress) {
            return tunnelLocalAddress(Output.of(tunnelLocalAddress));
        }

        /**
         * @param tunnelRemoteAddress Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
         * 
         * @return builder
         * 
         */
        public Builder tunnelRemoteAddress(@Nullable Output<String> tunnelRemoteAddress) {
            $.tunnelRemoteAddress = tunnelRemoteAddress;
            return this;
        }

        /**
         * @param tunnelRemoteAddress Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
         * 
         * @return builder
         * 
         */
        public Builder tunnelRemoteAddress(String tunnelRemoteAddress) {
            return tunnelRemoteAddress(Output.of(tunnelRemoteAddress));
        }

        public IpsecPolicyArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("IpsecPolicyArgs", "name");
            }
            return $;
        }
    }

}
