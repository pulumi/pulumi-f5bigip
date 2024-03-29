// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class SnatPoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final SnatPoolArgs Empty = new SnatPoolArgs();

    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     * 
     */
    @Import(name="members", required=true)
    private Output<List<String>> members;

    /**
     * @return Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }

    /**
     * Name of the snatpool
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the snatpool
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private SnatPoolArgs() {}

    private SnatPoolArgs(SnatPoolArgs $) {
        this.members = $.members;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnatPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnatPoolArgs $;

        public Builder() {
            $ = new SnatPoolArgs();
        }

        public Builder(SnatPoolArgs defaults) {
            $ = new SnatPoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param members Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
         * 
         * @return builder
         * 
         */
        public Builder members(Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param name Name of the snatpool
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the snatpool
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public SnatPoolArgs build() {
            if ($.members == null) {
                throw new MissingRequiredPropertyException("SnatPoolArgs", "members");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SnatPoolArgs", "name");
            }
            return $;
        }
    }

}
