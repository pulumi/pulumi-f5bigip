// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.ltm.inputs.NodeFqdnArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NodeState extends com.pulumi.resources.ResourceArgs {

    public static final NodeState Empty = new NodeState();

    /**
     * IP or hostname of the node
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return IP or hostname of the node
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     * 
     */
    @Import(name="connectionLimit")
    private @Nullable Output<Integer> connectionLimit;

    /**
     * @return Specifies the maximum number of connections allowed for the node or node address.
     * 
     */
    public Optional<Output<Integer>> connectionLimit() {
        return Optional.ofNullable(this.connectionLimit);
    }

    /**
     * User-defined description give ltm_node
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User-defined description give ltm_node
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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

    @Import(name="fqdn")
    private @Nullable Output<NodeFqdnArgs> fqdn;

    public Optional<Output<NodeFqdnArgs>> fqdn() {
        return Optional.ofNullable(this.fqdn);
    }

    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     * 
     */
    @Import(name="monitor")
    private @Nullable Output<String> monitor;

    /**
     * @return specifies the name of the monitor or monitor rule that you want to associate with the node.
     * 
     */
    public Optional<Output<String>> monitor() {
        return Optional.ofNullable(this.monitor);
    }

    /**
     * Name of the node
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the node
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
     * 
     */
    @Import(name="rateLimit")
    private @Nullable Output<String> rateLimit;

    /**
     * @return Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
     * 
     */
    public Optional<Output<String>> rateLimit() {
        return Optional.ofNullable(this.rateLimit);
    }

    /**
     * Sets the ratio number for the node.
     * 
     */
    @Import(name="ratio")
    private @Nullable Output<Integer> ratio;

    /**
     * @return Sets the ratio number for the node.
     * 
     */
    public Optional<Output<Integer>> ratio() {
        return Optional.ofNullable(this.ratio);
    }

    /**
     * Enables or disables the node for new sessions. The default value is user-enabled.
     * 
     */
    @Import(name="session")
    private @Nullable Output<String> session;

    /**
     * @return Enables or disables the node for new sessions. The default value is user-enabled.
     * 
     */
    public Optional<Output<String>> session() {
        return Optional.ofNullable(this.session);
    }

    /**
     * Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
     * 
     * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
     * 
     * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private NodeState() {}

    private NodeState(NodeState $) {
        this.address = $.address;
        this.connectionLimit = $.connectionLimit;
        this.description = $.description;
        this.dynamicRatio = $.dynamicRatio;
        this.fqdn = $.fqdn;
        this.monitor = $.monitor;
        this.name = $.name;
        this.rateLimit = $.rateLimit;
        this.ratio = $.ratio;
        this.session = $.session;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodeState $;

        public Builder() {
            $ = new NodeState();
        }

        public Builder(NodeState defaults) {
            $ = new NodeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param address IP or hostname of the node
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address IP or hostname of the node
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param connectionLimit Specifies the maximum number of connections allowed for the node or node address.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(@Nullable Output<Integer> connectionLimit) {
            $.connectionLimit = connectionLimit;
            return this;
        }

        /**
         * @param connectionLimit Specifies the maximum number of connections allowed for the node or node address.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(Integer connectionLimit) {
            return connectionLimit(Output.of(connectionLimit));
        }

        /**
         * @param description User-defined description give ltm_node
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User-defined description give ltm_node
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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

        public Builder fqdn(@Nullable Output<NodeFqdnArgs> fqdn) {
            $.fqdn = fqdn;
            return this;
        }

        public Builder fqdn(NodeFqdnArgs fqdn) {
            return fqdn(Output.of(fqdn));
        }

        /**
         * @param monitor specifies the name of the monitor or monitor rule that you want to associate with the node.
         * 
         * @return builder
         * 
         */
        public Builder monitor(@Nullable Output<String> monitor) {
            $.monitor = monitor;
            return this;
        }

        /**
         * @param monitor specifies the name of the monitor or monitor rule that you want to associate with the node.
         * 
         * @return builder
         * 
         */
        public Builder monitor(String monitor) {
            return monitor(Output.of(monitor));
        }

        /**
         * @param name Name of the node
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the node
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rateLimit Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
         * 
         * @return builder
         * 
         */
        public Builder rateLimit(@Nullable Output<String> rateLimit) {
            $.rateLimit = rateLimit;
            return this;
        }

        /**
         * @param rateLimit Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
         * 
         * @return builder
         * 
         */
        public Builder rateLimit(String rateLimit) {
            return rateLimit(Output.of(rateLimit));
        }

        /**
         * @param ratio Sets the ratio number for the node.
         * 
         * @return builder
         * 
         */
        public Builder ratio(@Nullable Output<Integer> ratio) {
            $.ratio = ratio;
            return this;
        }

        /**
         * @param ratio Sets the ratio number for the node.
         * 
         * @return builder
         * 
         */
        public Builder ratio(Integer ratio) {
            return ratio(Output.of(ratio));
        }

        /**
         * @param session Enables or disables the node for new sessions. The default value is user-enabled.
         * 
         * @return builder
         * 
         */
        public Builder session(@Nullable Output<String> session) {
            $.session = session;
            return this;
        }

        /**
         * @param session Enables or disables the node for new sessions. The default value is user-enabled.
         * 
         * @return builder
         * 
         */
        public Builder session(String session) {
            return session(Output.of(session));
        }

        /**
         * @param state Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
         * 
         * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
         * 
         * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public NodeState build() {
            return $;
        }
    }

}
