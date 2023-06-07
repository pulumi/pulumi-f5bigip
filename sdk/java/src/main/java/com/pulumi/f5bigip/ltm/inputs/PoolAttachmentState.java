// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PoolAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final PoolAttachmentState Empty = new PoolAttachmentState();

    /**
     * Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
     * 
     */
    @Import(name="connectionLimit")
    private @Nullable Output<Integer> connectionLimit;

    /**
     * @return Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
     * 
     */
    public Optional<Output<Integer>> connectionLimit() {
        return Optional.ofNullable(this.connectionLimit);
    }

    /**
     * Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
     * 
     */
    @Import(name="connectionRateLimit")
    private @Nullable Output<String> connectionRateLimit;

    /**
     * @return Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
     * 
     */
    public Optional<Output<String>> connectionRateLimit() {
        return Optional.ofNullable(this.connectionRateLimit);
    }

    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     * 
     */
    @Import(name="dynamicRatio")
    private @Nullable Output<Integer> dynamicRatio;

    /**
     * @return Specifies the fixed ratio value used for a node during ratio load balancing.
     * 
     */
    public Optional<Output<Integer>> dynamicRatio() {
        return Optional.ofNullable(this.dynamicRatio);
    }

    /**
     * Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
     * 
     */
    @Import(name="fqdnAutopopulate")
    private @Nullable Output<String> fqdnAutopopulate;

    /**
     * @return Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
     * 
     */
    public Optional<Output<String>> fqdnAutopopulate() {
        return Optional.ofNullable(this.fqdnAutopopulate);
    }

    /**
     * Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
     * 
     */
    @Import(name="monitor")
    private @Nullable Output<String> monitor;

    /**
     * @return Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
     * 
     */
    public Optional<Output<String>> monitor() {
        return Optional.ofNullable(this.monitor);
    }

    /**
     * Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
     * 
     */
    @Import(name="node")
    private @Nullable Output<String> node;

    /**
     * @return Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
     * 
     */
    public Optional<Output<String>> node() {
        return Optional.ofNullable(this.node);
    }

    /**
     * Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
     * 
     */
    @Import(name="pool")
    private @Nullable Output<String> pool;

    /**
     * @return Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
     * 
     */
    public Optional<Output<String>> pool() {
        return Optional.ofNullable(this.pool);
    }

    /**
     * Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
     * 
     */
    @Import(name="priorityGroup")
    private @Nullable Output<Integer> priorityGroup;

    /**
     * @return Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
     * 
     */
    public Optional<Output<Integer>> priorityGroup() {
        return Optional.ofNullable(this.priorityGroup);
    }

    /**
     * &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
     * 
     */
    @Import(name="ratio")
    private @Nullable Output<Integer> ratio;

    /**
     * @return &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
     * 
     */
    public Optional<Output<Integer>> ratio() {
        return Optional.ofNullable(this.ratio);
    }

    /**
     * Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private PoolAttachmentState() {}

    private PoolAttachmentState(PoolAttachmentState $) {
        this.connectionLimit = $.connectionLimit;
        this.connectionRateLimit = $.connectionRateLimit;
        this.dynamicRatio = $.dynamicRatio;
        this.fqdnAutopopulate = $.fqdnAutopopulate;
        this.monitor = $.monitor;
        this.node = $.node;
        this.pool = $.pool;
        this.priorityGroup = $.priorityGroup;
        this.ratio = $.ratio;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PoolAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PoolAttachmentState $;

        public Builder() {
            $ = new PoolAttachmentState();
        }

        public Builder(PoolAttachmentState defaults) {
            $ = new PoolAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionLimit Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(@Nullable Output<Integer> connectionLimit) {
            $.connectionLimit = connectionLimit;
            return this;
        }

        /**
         * @param connectionLimit Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(Integer connectionLimit) {
            return connectionLimit(Output.of(connectionLimit));
        }

        /**
         * @param connectionRateLimit Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
         * 
         * @return builder
         * 
         */
        public Builder connectionRateLimit(@Nullable Output<String> connectionRateLimit) {
            $.connectionRateLimit = connectionRateLimit;
            return this;
        }

        /**
         * @param connectionRateLimit Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
         * 
         * @return builder
         * 
         */
        public Builder connectionRateLimit(String connectionRateLimit) {
            return connectionRateLimit(Output.of(connectionRateLimit));
        }

        /**
         * @param dynamicRatio Specifies the fixed ratio value used for a node during ratio load balancing.
         * 
         * @return builder
         * 
         */
        public Builder dynamicRatio(@Nullable Output<Integer> dynamicRatio) {
            $.dynamicRatio = dynamicRatio;
            return this;
        }

        /**
         * @param dynamicRatio Specifies the fixed ratio value used for a node during ratio load balancing.
         * 
         * @return builder
         * 
         */
        public Builder dynamicRatio(Integer dynamicRatio) {
            return dynamicRatio(Output.of(dynamicRatio));
        }

        /**
         * @param fqdnAutopopulate Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
         * 
         * @return builder
         * 
         */
        public Builder fqdnAutopopulate(@Nullable Output<String> fqdnAutopopulate) {
            $.fqdnAutopopulate = fqdnAutopopulate;
            return this;
        }

        /**
         * @param fqdnAutopopulate Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
         * 
         * @return builder
         * 
         */
        public Builder fqdnAutopopulate(String fqdnAutopopulate) {
            return fqdnAutopopulate(Output.of(fqdnAutopopulate));
        }

        /**
         * @param monitor Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
         * 
         * @return builder
         * 
         */
        public Builder monitor(@Nullable Output<String> monitor) {
            $.monitor = monitor;
            return this;
        }

        /**
         * @param monitor Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
         * 
         * @return builder
         * 
         */
        public Builder monitor(String monitor) {
            return monitor(Output.of(monitor));
        }

        /**
         * @param node Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
         * 
         * @return builder
         * 
         */
        public Builder node(@Nullable Output<String> node) {
            $.node = node;
            return this;
        }

        /**
         * @param node Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
         * 
         * @return builder
         * 
         */
        public Builder node(String node) {
            return node(Output.of(node));
        }

        /**
         * @param pool Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
         * 
         * @return builder
         * 
         */
        public Builder pool(@Nullable Output<String> pool) {
            $.pool = pool;
            return this;
        }

        /**
         * @param pool Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
         * 
         * @return builder
         * 
         */
        public Builder pool(String pool) {
            return pool(Output.of(pool));
        }

        /**
         * @param priorityGroup Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
         * 
         * @return builder
         * 
         */
        public Builder priorityGroup(@Nullable Output<Integer> priorityGroup) {
            $.priorityGroup = priorityGroup;
            return this;
        }

        /**
         * @param priorityGroup Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
         * 
         * @return builder
         * 
         */
        public Builder priorityGroup(Integer priorityGroup) {
            return priorityGroup(Output.of(priorityGroup));
        }

        /**
         * @param ratio &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
         * 
         * @return builder
         * 
         */
        public Builder ratio(@Nullable Output<Integer> ratio) {
            $.ratio = ratio;
            return this;
        }

        /**
         * @param ratio &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
         * 
         * @return builder
         * 
         */
        public Builder ratio(Integer ratio) {
            return ratio(Output.of(ratio));
        }

        /**
         * @param state Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public PoolAttachmentState build() {
            return $;
        }
    }

}
