// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FastTcpAppPoolMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final FastTcpAppPoolMemberArgs Empty = new FastTcpAppPoolMemberArgs();

    /**
     * List of server address to be used for FAST-Generated Pool.
     * 
     */
    @Import(name="addresses", required=true)
    private Output<List<String>> addresses;

    /**
     * @return List of server address to be used for FAST-Generated Pool.
     * 
     */
    public Output<List<String>> addresses() {
        return this.addresses;
    }

    /**
     * connectionLimit value to be used for FAST-Generated Pool.
     * 
     */
    @Import(name="connectionLimit")
    private @Nullable Output<Integer> connectionLimit;

    /**
     * @return connectionLimit value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Output<Integer>> connectionLimit() {
        return Optional.ofNullable(this.connectionLimit);
    }

    /**
     * port number of serviceport to be used for FAST-Generated Pool.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return port number of serviceport to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * priorityGroup value to be used for FAST-Generated Pool.
     * 
     */
    @Import(name="priorityGroup")
    private @Nullable Output<Integer> priorityGroup;

    /**
     * @return priorityGroup value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Output<Integer>> priorityGroup() {
        return Optional.ofNullable(this.priorityGroup);
    }

    /**
     * shareNodes value to be used for FAST-Generated Pool.
     * 
     */
    @Import(name="shareNodes")
    private @Nullable Output<Boolean> shareNodes;

    /**
     * @return shareNodes value to be used for FAST-Generated Pool.
     * 
     */
    public Optional<Output<Boolean>> shareNodes() {
        return Optional.ofNullable(this.shareNodes);
    }

    private FastTcpAppPoolMemberArgs() {}

    private FastTcpAppPoolMemberArgs(FastTcpAppPoolMemberArgs $) {
        this.addresses = $.addresses;
        this.connectionLimit = $.connectionLimit;
        this.port = $.port;
        this.priorityGroup = $.priorityGroup;
        this.shareNodes = $.shareNodes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FastTcpAppPoolMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FastTcpAppPoolMemberArgs $;

        public Builder() {
            $ = new FastTcpAppPoolMemberArgs();
        }

        public Builder(FastTcpAppPoolMemberArgs defaults) {
            $ = new FastTcpAppPoolMemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addresses List of server address to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder addresses(Output<List<String>> addresses) {
            $.addresses = addresses;
            return this;
        }

        /**
         * @param addresses List of server address to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder addresses(List<String> addresses) {
            return addresses(Output.of(addresses));
        }

        /**
         * @param addresses List of server address to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder addresses(String... addresses) {
            return addresses(List.of(addresses));
        }

        /**
         * @param connectionLimit connectionLimit value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(@Nullable Output<Integer> connectionLimit) {
            $.connectionLimit = connectionLimit;
            return this;
        }

        /**
         * @param connectionLimit connectionLimit value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(Integer connectionLimit) {
            return connectionLimit(Output.of(connectionLimit));
        }

        /**
         * @param port port number of serviceport to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port port number of serviceport to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param priorityGroup priorityGroup value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder priorityGroup(@Nullable Output<Integer> priorityGroup) {
            $.priorityGroup = priorityGroup;
            return this;
        }

        /**
         * @param priorityGroup priorityGroup value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder priorityGroup(Integer priorityGroup) {
            return priorityGroup(Output.of(priorityGroup));
        }

        /**
         * @param shareNodes shareNodes value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder shareNodes(@Nullable Output<Boolean> shareNodes) {
            $.shareNodes = shareNodes;
            return this;
        }

        /**
         * @param shareNodes shareNodes value to be used for FAST-Generated Pool.
         * 
         * @return builder
         * 
         */
        public Builder shareNodes(Boolean shareNodes) {
            return shareNodes(Output.of(shareNodes));
        }

        public FastTcpAppPoolMemberArgs build() {
            if ($.addresses == null) {
                throw new MissingRequiredPropertyException("FastTcpAppPoolMemberArgs", "addresses");
            }
            return $;
        }
    }

}