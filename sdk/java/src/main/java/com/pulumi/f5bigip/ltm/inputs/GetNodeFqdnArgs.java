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


public final class GetNodeFqdnArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetNodeFqdnArgs Empty = new GetNodeFqdnArgs();

    /**
     * The FQDN node&#39;s address family.
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return The FQDN node&#39;s address family.
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * Specifies if the node should scale to the IP address set returned by DNS.
     * 
     */
    @Import(name="autopopulate", required=true)
    private Output<String> autopopulate;

    /**
     * @return Specifies if the node should scale to the IP address set returned by DNS.
     * 
     */
    public Output<String> autopopulate() {
        return this.autopopulate;
    }

    /**
     * The number of attempts to resolve a domain name.
     * 
     */
    @Import(name="downinterval", required=true)
    private Output<Integer> downinterval;

    /**
     * @return The number of attempts to resolve a domain name.
     * 
     */
    public Output<Integer> downinterval() {
        return this.downinterval;
    }

    /**
     * The amount of time before sending the next DNS query.
     * 
     */
    @Import(name="interval", required=true)
    private Output<String> interval;

    /**
     * @return The amount of time before sending the next DNS query.
     * 
     */
    public Output<String> interval() {
        return this.interval;
    }

    /**
     * Name of the node.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the node.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetNodeFqdnArgs() {}

    private GetNodeFqdnArgs(GetNodeFqdnArgs $) {
        this.addressFamily = $.addressFamily;
        this.autopopulate = $.autopopulate;
        this.downinterval = $.downinterval;
        this.interval = $.interval;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNodeFqdnArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNodeFqdnArgs $;

        public Builder() {
            $ = new GetNodeFqdnArgs();
        }

        public Builder(GetNodeFqdnArgs defaults) {
            $ = new GetNodeFqdnArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily The FQDN node&#39;s address family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily The FQDN node&#39;s address family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param autopopulate Specifies if the node should scale to the IP address set returned by DNS.
         * 
         * @return builder
         * 
         */
        public Builder autopopulate(Output<String> autopopulate) {
            $.autopopulate = autopopulate;
            return this;
        }

        /**
         * @param autopopulate Specifies if the node should scale to the IP address set returned by DNS.
         * 
         * @return builder
         * 
         */
        public Builder autopopulate(String autopopulate) {
            return autopopulate(Output.of(autopopulate));
        }

        /**
         * @param downinterval The number of attempts to resolve a domain name.
         * 
         * @return builder
         * 
         */
        public Builder downinterval(Output<Integer> downinterval) {
            $.downinterval = downinterval;
            return this;
        }

        /**
         * @param downinterval The number of attempts to resolve a domain name.
         * 
         * @return builder
         * 
         */
        public Builder downinterval(Integer downinterval) {
            return downinterval(Output.of(downinterval));
        }

        /**
         * @param interval The amount of time before sending the next DNS query.
         * 
         * @return builder
         * 
         */
        public Builder interval(Output<String> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval The amount of time before sending the next DNS query.
         * 
         * @return builder
         * 
         */
        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param name Name of the node.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the node.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetNodeFqdnArgs build() {
            $.autopopulate = Objects.requireNonNull($.autopopulate, "expected parameter 'autopopulate' to be non-null");
            $.downinterval = Objects.requireNonNull($.downinterval, "expected parameter 'downinterval' to be non-null");
            $.interval = Objects.requireNonNull($.interval, "expected parameter 'interval' to be non-null");
            return $;
        }
    }

}
