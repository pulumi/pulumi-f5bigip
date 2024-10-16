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


public final class NodeFqdnArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodeFqdnArgs Empty = new NodeFqdnArgs();

    /**
     * Specifies the node&#39;s address family. The default is &#39;unspecified&#39;, or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return Specifies the node&#39;s address family. The default is &#39;unspecified&#39;, or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * Specifies whether the node should scale to the IP address set returned by DNS.
     * 
     */
    @Import(name="autopopulate")
    private @Nullable Output<String> autopopulate;

    /**
     * @return Specifies whether the node should scale to the IP address set returned by DNS.
     * 
     */
    public Optional<Output<String>> autopopulate() {
        return Optional.ofNullable(this.autopopulate);
    }

    /**
     * Specifies the number of attempts to resolve a domain name. The default is 5.
     * 
     */
    @Import(name="downinterval")
    private @Nullable Output<Integer> downinterval;

    /**
     * @return Specifies the number of attempts to resolve a domain name. The default is 5.
     * 
     */
    public Optional<Output<Integer>> downinterval() {
        return Optional.ofNullable(this.downinterval);
    }

    /**
     * Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
     * 
     */
    @Import(name="interval")
    private @Nullable Output<String> interval;

    /**
     * @return Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
     * 
     */
    public Optional<Output<String>> interval() {
        return Optional.ofNullable(this.interval);
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

    private NodeFqdnArgs() {}

    private NodeFqdnArgs(NodeFqdnArgs $) {
        this.addressFamily = $.addressFamily;
        this.autopopulate = $.autopopulate;
        this.downinterval = $.downinterval;
        this.interval = $.interval;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodeFqdnArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodeFqdnArgs $;

        public Builder() {
            $ = new NodeFqdnArgs();
        }

        public Builder(NodeFqdnArgs defaults) {
            $ = new NodeFqdnArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Specifies the node&#39;s address family. The default is &#39;unspecified&#39;, or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Specifies the node&#39;s address family. The default is &#39;unspecified&#39;, or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param autopopulate Specifies whether the node should scale to the IP address set returned by DNS.
         * 
         * @return builder
         * 
         */
        public Builder autopopulate(@Nullable Output<String> autopopulate) {
            $.autopopulate = autopopulate;
            return this;
        }

        /**
         * @param autopopulate Specifies whether the node should scale to the IP address set returned by DNS.
         * 
         * @return builder
         * 
         */
        public Builder autopopulate(String autopopulate) {
            return autopopulate(Output.of(autopopulate));
        }

        /**
         * @param downinterval Specifies the number of attempts to resolve a domain name. The default is 5.
         * 
         * @return builder
         * 
         */
        public Builder downinterval(@Nullable Output<Integer> downinterval) {
            $.downinterval = downinterval;
            return this;
        }

        /**
         * @param downinterval Specifies the number of attempts to resolve a domain name. The default is 5.
         * 
         * @return builder
         * 
         */
        public Builder downinterval(Integer downinterval) {
            return downinterval(Output.of(downinterval));
        }

        /**
         * @param interval Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<String> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
         * 
         * @return builder
         * 
         */
        public Builder interval(String interval) {
            return interval(Output.of(interval));
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

        public NodeFqdnArgs build() {
            return $;
        }
    }

}