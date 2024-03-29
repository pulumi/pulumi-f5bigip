// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetMonitorPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMonitorPlainArgs Empty = new GetMonitorPlainArgs();

    /**
     * Name of the ltm monitor
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the ltm monitor
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * partition of the ltm monitor
     * 
     */
    @Import(name="partition", required=true)
    private String partition;

    /**
     * @return partition of the ltm monitor
     * 
     */
    public String partition() {
        return this.partition;
    }

    private GetMonitorPlainArgs() {}

    private GetMonitorPlainArgs(GetMonitorPlainArgs $) {
        this.name = $.name;
        this.partition = $.partition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMonitorPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMonitorPlainArgs $;

        public Builder() {
            $ = new GetMonitorPlainArgs();
        }

        public Builder(GetMonitorPlainArgs defaults) {
            $ = new GetMonitorPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param partition partition of the ltm monitor
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            $.partition = partition;
            return this;
        }

        public GetMonitorPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetMonitorPlainArgs", "name");
            }
            if ($.partition == null) {
                throw new MissingRequiredPropertyException("GetMonitorPlainArgs", "partition");
            }
            return $;
        }
    }

}
