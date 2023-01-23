// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetMonitorArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMonitorArgs Empty = new GetMonitorArgs();

    /**
     * Name of the ltm monitor
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the ltm monitor
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * partition of the ltm monitor
     * 
     */
    @Import(name="partition", required=true)
    private Output<String> partition;

    /**
     * @return partition of the ltm monitor
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }

    private GetMonitorArgs() {}

    private GetMonitorArgs(GetMonitorArgs $) {
        this.name = $.name;
        this.partition = $.partition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMonitorArgs $;

        public Builder() {
            $ = new GetMonitorArgs();
        }

        public Builder(GetMonitorArgs defaults) {
            $ = new GetMonitorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition partition of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder partition(Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition partition of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        public GetMonitorArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.partition = Objects.requireNonNull($.partition, "expected parameter 'partition' to be non-null");
            return $;
        }
    }

}