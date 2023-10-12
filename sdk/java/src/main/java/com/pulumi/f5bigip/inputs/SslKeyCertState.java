// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SslKeyCertState extends com.pulumi.resources.ResourceArgs {

    public static final SslKeyCertState Empty = new SslKeyCertState();

    /**
     * The content of the cert.
     * 
     */
    @Import(name="certContent")
    private @Nullable Output<String> certContent;

    /**
     * @return The content of the cert.
     * 
     */
    public Optional<Output<String>> certContent() {
        return Optional.ofNullable(this.certContent);
    }

    /**
     * full path of the SSL certificate on the BIGIP.
     * 
     */
    @Import(name="certFullPath")
    private @Nullable Output<String> certFullPath;

    /**
     * @return full path of the SSL certificate on the BIGIP.
     * 
     */
    public Optional<Output<String>> certFullPath() {
        return Optional.ofNullable(this.certFullPath);
    }

    /**
     * Name of the SSL certificate to be Imported on to BIGIP.
     * 
     */
    @Import(name="certName")
    private @Nullable Output<String> certName;

    /**
     * @return Name of the SSL certificate to be Imported on to BIGIP.
     * 
     */
    public Optional<Output<String>> certName() {
        return Optional.ofNullable(this.certName);
    }

    /**
     * The content of the key.
     * 
     */
    @Import(name="keyContent")
    private @Nullable Output<String> keyContent;

    /**
     * @return The content of the key.
     * 
     */
    public Optional<Output<String>> keyContent() {
        return Optional.ofNullable(this.keyContent);
    }

    /**
     * full path of the SSL key on the BIGIP.
     * 
     */
    @Import(name="keyFullPath")
    private @Nullable Output<String> keyFullPath;

    /**
     * @return full path of the SSL key on the BIGIP.
     * 
     */
    public Optional<Output<String>> keyFullPath() {
        return Optional.ofNullable(this.keyFullPath);
    }

    /**
     * Name of the SSL key to be Imported on to BIGIP.
     * 
     */
    @Import(name="keyName")
    private @Nullable Output<String> keyName;

    /**
     * @return Name of the SSL key to be Imported on to BIGIP.
     * 
     */
    public Optional<Output<String>> keyName() {
        return Optional.ofNullable(this.keyName);
    }

    /**
     * Partition on to SSL certificate and key to be imported.
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Partition on to SSL certificate and key to be imported.
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    /**
     * Passphrase on the SSL key.
     * 
     */
    @Import(name="passphrase")
    private @Nullable Output<String> passphrase;

    /**
     * @return Passphrase on the SSL key.
     * 
     */
    public Optional<Output<String>> passphrase() {
        return Optional.ofNullable(this.passphrase);
    }

    private SslKeyCertState() {}

    private SslKeyCertState(SslKeyCertState $) {
        this.certContent = $.certContent;
        this.certFullPath = $.certFullPath;
        this.certName = $.certName;
        this.keyContent = $.keyContent;
        this.keyFullPath = $.keyFullPath;
        this.keyName = $.keyName;
        this.partition = $.partition;
        this.passphrase = $.passphrase;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SslKeyCertState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SslKeyCertState $;

        public Builder() {
            $ = new SslKeyCertState();
        }

        public Builder(SslKeyCertState defaults) {
            $ = new SslKeyCertState(Objects.requireNonNull(defaults));
        }

        /**
         * @param certContent The content of the cert.
         * 
         * @return builder
         * 
         */
        public Builder certContent(@Nullable Output<String> certContent) {
            $.certContent = certContent;
            return this;
        }

        /**
         * @param certContent The content of the cert.
         * 
         * @return builder
         * 
         */
        public Builder certContent(String certContent) {
            return certContent(Output.of(certContent));
        }

        /**
         * @param certFullPath full path of the SSL certificate on the BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder certFullPath(@Nullable Output<String> certFullPath) {
            $.certFullPath = certFullPath;
            return this;
        }

        /**
         * @param certFullPath full path of the SSL certificate on the BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder certFullPath(String certFullPath) {
            return certFullPath(Output.of(certFullPath));
        }

        /**
         * @param certName Name of the SSL certificate to be Imported on to BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder certName(@Nullable Output<String> certName) {
            $.certName = certName;
            return this;
        }

        /**
         * @param certName Name of the SSL certificate to be Imported on to BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder certName(String certName) {
            return certName(Output.of(certName));
        }

        /**
         * @param keyContent The content of the key.
         * 
         * @return builder
         * 
         */
        public Builder keyContent(@Nullable Output<String> keyContent) {
            $.keyContent = keyContent;
            return this;
        }

        /**
         * @param keyContent The content of the key.
         * 
         * @return builder
         * 
         */
        public Builder keyContent(String keyContent) {
            return keyContent(Output.of(keyContent));
        }

        /**
         * @param keyFullPath full path of the SSL key on the BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder keyFullPath(@Nullable Output<String> keyFullPath) {
            $.keyFullPath = keyFullPath;
            return this;
        }

        /**
         * @param keyFullPath full path of the SSL key on the BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder keyFullPath(String keyFullPath) {
            return keyFullPath(Output.of(keyFullPath));
        }

        /**
         * @param keyName Name of the SSL key to be Imported on to BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder keyName(@Nullable Output<String> keyName) {
            $.keyName = keyName;
            return this;
        }

        /**
         * @param keyName Name of the SSL key to be Imported on to BIGIP.
         * 
         * @return builder
         * 
         */
        public Builder keyName(String keyName) {
            return keyName(Output.of(keyName));
        }

        /**
         * @param partition Partition on to SSL certificate and key to be imported.
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Partition on to SSL certificate and key to be imported.
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param passphrase Passphrase on the SSL key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(@Nullable Output<String> passphrase) {
            $.passphrase = passphrase;
            return this;
        }

        /**
         * @param passphrase Passphrase on the SSL key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(String passphrase) {
            return passphrase(Output.of(passphrase));
        }

        public SslKeyCertState build() {
            return $;
        }
    }

}