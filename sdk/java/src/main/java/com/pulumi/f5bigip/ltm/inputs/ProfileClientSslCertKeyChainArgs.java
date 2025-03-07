// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProfileClientSslCertKeyChainArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProfileClientSslCertKeyChainArgs Empty = new ProfileClientSslCertKeyChainArgs();

    /**
     * Specifies the name of the certificate that the system uses for client-side SSL processing. The default is `default`
     * 
     */
    @Import(name="cert")
    private @Nullable Output<String> cert;

    /**
     * @return Specifies the name of the certificate that the system uses for client-side SSL processing. The default is `default`
     * 
     */
    public Optional<Output<String>> cert() {
        return Optional.ofNullable(this.cert);
    }

    /**
     * Specifies a certificate chain file that a server can use for authentication. The default is `None`.
     * 
     */
    @Import(name="chain")
    private @Nullable Output<String> chain;

    /**
     * @return Specifies a certificate chain file that a server can use for authentication. The default is `None`.
     * 
     */
    public Optional<Output<String>> chain() {
        return Optional.ofNullable(this.chain);
    }

    /**
     * Specifies the file name of the SSL key. The default is `default`
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return Specifies the file name of the SSL key. The default is `default`
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Name of Cert-key-chain
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of Cert-key-chain
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Type the name of the pass phrase used to encrypt the key.
     * 
     */
    @Import(name="passphrase")
    private @Nullable Output<String> passphrase;

    /**
     * @return Type the name of the pass phrase used to encrypt the key.
     * 
     */
    public Optional<Output<String>> passphrase() {
        return Optional.ofNullable(this.passphrase);
    }

    private ProfileClientSslCertKeyChainArgs() {}

    private ProfileClientSslCertKeyChainArgs(ProfileClientSslCertKeyChainArgs $) {
        this.cert = $.cert;
        this.chain = $.chain;
        this.key = $.key;
        this.name = $.name;
        this.passphrase = $.passphrase;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProfileClientSslCertKeyChainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProfileClientSslCertKeyChainArgs $;

        public Builder() {
            $ = new ProfileClientSslCertKeyChainArgs();
        }

        public Builder(ProfileClientSslCertKeyChainArgs defaults) {
            $ = new ProfileClientSslCertKeyChainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cert Specifies the name of the certificate that the system uses for client-side SSL processing. The default is `default`
         * 
         * @return builder
         * 
         */
        public Builder cert(@Nullable Output<String> cert) {
            $.cert = cert;
            return this;
        }

        /**
         * @param cert Specifies the name of the certificate that the system uses for client-side SSL processing. The default is `default`
         * 
         * @return builder
         * 
         */
        public Builder cert(String cert) {
            return cert(Output.of(cert));
        }

        /**
         * @param chain Specifies a certificate chain file that a server can use for authentication. The default is `None`.
         * 
         * @return builder
         * 
         */
        public Builder chain(@Nullable Output<String> chain) {
            $.chain = chain;
            return this;
        }

        /**
         * @param chain Specifies a certificate chain file that a server can use for authentication. The default is `None`.
         * 
         * @return builder
         * 
         */
        public Builder chain(String chain) {
            return chain(Output.of(chain));
        }

        /**
         * @param key Specifies the file name of the SSL key. The default is `default`
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Specifies the file name of the SSL key. The default is `default`
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param name Name of Cert-key-chain
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of Cert-key-chain
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param passphrase Type the name of the pass phrase used to encrypt the key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(@Nullable Output<String> passphrase) {
            $.passphrase = passphrase;
            return this;
        }

        /**
         * @param passphrase Type the name of the pass phrase used to encrypt the key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(String passphrase) {
            return passphrase(Output.of(passphrase));
        }

        public ProfileClientSslCertKeyChainArgs build() {
            return $;
        }
    }

}
