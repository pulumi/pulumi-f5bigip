// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final KeyArgs Empty = new KeyArgs();

    /**
     * Content of SSL certificate key present on local Disk
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return Content of SSL certificate key present on local Disk
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * Full Path Name of ssl key
     * 
     */
    @Import(name="fullPath")
    private @Nullable Output<String> fullPath;

    /**
     * @return Full Path Name of ssl key
     * 
     */
    public Optional<Output<String>> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }

    /**
     * Name of the SSL Certificate key to be Imported on to BIGIP
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the SSL Certificate key to be Imported on to BIGIP
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Partition of ssl certificate key
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Partition of ssl certificate key
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    /**
     * Passphrase on key.
     * 
     */
    @Import(name="passphrase")
    private @Nullable Output<String> passphrase;

    /**
     * @return Passphrase on key.
     * 
     */
    public Optional<Output<String>> passphrase() {
        return Optional.ofNullable(this.passphrase);
    }

    private KeyArgs() {}

    private KeyArgs(KeyArgs $) {
        this.content = $.content;
        this.fullPath = $.fullPath;
        this.name = $.name;
        this.partition = $.partition;
        this.passphrase = $.passphrase;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeyArgs $;

        public Builder() {
            $ = new KeyArgs();
        }

        public Builder(KeyArgs defaults) {
            $ = new KeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param content Content of SSL certificate key present on local Disk
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content Content of SSL certificate key present on local Disk
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param fullPath Full Path Name of ssl key
         * 
         * @return builder
         * 
         */
        public Builder fullPath(@Nullable Output<String> fullPath) {
            $.fullPath = fullPath;
            return this;
        }

        /**
         * @param fullPath Full Path Name of ssl key
         * 
         * @return builder
         * 
         */
        public Builder fullPath(String fullPath) {
            return fullPath(Output.of(fullPath));
        }

        /**
         * @param name Name of the SSL Certificate key to be Imported on to BIGIP
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the SSL Certificate key to be Imported on to BIGIP
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition Partition of ssl certificate key
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Partition of ssl certificate key
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param passphrase Passphrase on key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(@Nullable Output<String> passphrase) {
            $.passphrase = passphrase;
            return this;
        }

        /**
         * @param passphrase Passphrase on key.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(String passphrase) {
            return passphrase(Output.of(passphrase));
        }

        public KeyArgs build() {
            if ($.content == null) {
                throw new MissingRequiredPropertyException("KeyArgs", "content");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("KeyArgs", "name");
            }
            return $;
        }
    }

}
