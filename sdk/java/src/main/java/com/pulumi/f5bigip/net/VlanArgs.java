// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.net;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.net.inputs.VlanInterfaceArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VlanArgs extends com.pulumi.resources.ResourceArgs {

    public static final VlanArgs Empty = new VlanArgs();

    /**
     * Specifies which interfaces you want this VLAN to use for traffic management.
     * 
     */
    @Import(name="interfaces")
    private @Nullable Output<List<VlanInterfaceArgs>> interfaces;

    /**
     * @return Specifies which interfaces you want this VLAN to use for traffic management.
     * 
     */
    public Optional<Output<List<VlanInterfaceArgs>>> interfaces() {
        return Optional.ofNullable(this.interfaces);
    }

    /**
     * Name of the vlan
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the vlan
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Specifies a number that the system adds into the header of any frame passing through the VLAN.
     * 
     */
    @Import(name="tag")
    private @Nullable Output<Integer> tag;

    /**
     * @return Specifies a number that the system adds into the header of any frame passing through the VLAN.
     * 
     */
    public Optional<Output<Integer>> tag() {
        return Optional.ofNullable(this.tag);
    }

    private VlanArgs() {}

    private VlanArgs(VlanArgs $) {
        this.interfaces = $.interfaces;
        this.name = $.name;
        this.tag = $.tag;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VlanArgs $;

        public Builder() {
            $ = new VlanArgs();
        }

        public Builder(VlanArgs defaults) {
            $ = new VlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param interfaces Specifies which interfaces you want this VLAN to use for traffic management.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(@Nullable Output<List<VlanInterfaceArgs>> interfaces) {
            $.interfaces = interfaces;
            return this;
        }

        /**
         * @param interfaces Specifies which interfaces you want this VLAN to use for traffic management.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(List<VlanInterfaceArgs> interfaces) {
            return interfaces(Output.of(interfaces));
        }

        /**
         * @param interfaces Specifies which interfaces you want this VLAN to use for traffic management.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(VlanInterfaceArgs... interfaces) {
            return interfaces(List.of(interfaces));
        }

        /**
         * @param name Name of the vlan
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the vlan
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tag Specifies a number that the system adds into the header of any frame passing through the VLAN.
         * 
         * @return builder
         * 
         */
        public Builder tag(@Nullable Output<Integer> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag Specifies a number that the system adds into the header of any frame passing through the VLAN.
         * 
         * @return builder
         * 
         */
        public Builder tag(Integer tag) {
            return tag(Output.of(tag));
        }

        public VlanArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
