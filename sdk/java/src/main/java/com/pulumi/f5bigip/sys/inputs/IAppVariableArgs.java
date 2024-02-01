// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IAppVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final IAppVariableArgs Empty = new IAppVariableArgs();

    /**
     * Name of origin
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<String> encrypted;

    /**
     * @return Name of origin
     * 
     */
    public Optional<Output<String>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * Name of the iApp.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the iApp.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name of origin
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return Name of origin
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private IAppVariableArgs() {}

    private IAppVariableArgs(IAppVariableArgs $) {
        this.encrypted = $.encrypted;
        this.name = $.name;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IAppVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IAppVariableArgs $;

        public Builder() {
            $ = new IAppVariableArgs();
        }

        public Builder(IAppVariableArgs defaults) {
            $ = new IAppVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param encrypted Name of origin
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<String> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted Name of origin
         * 
         * @return builder
         * 
         */
        public Builder encrypted(String encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param value Name of origin
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Name of origin
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public IAppVariableArgs build() {
            return $;
        }
    }

}
