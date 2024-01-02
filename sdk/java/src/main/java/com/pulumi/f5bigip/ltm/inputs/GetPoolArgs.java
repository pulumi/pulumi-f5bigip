// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetPoolArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPoolArgs Empty = new GetPoolArgs();

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

    private GetPoolArgs() {}

    private GetPoolArgs(GetPoolArgs $) {
        this.name = $.name;
        this.partition = $.partition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPoolArgs $;

        public Builder() {
            $ = new GetPoolArgs();
        }

        public Builder(GetPoolArgs defaults) {
            $ = new GetPoolArgs(Objects.requireNonNull(defaults));
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

        public GetPoolArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetPoolArgs", "name");
            }
            if ($.partition == null) {
                throw new MissingRequiredPropertyException("GetPoolArgs", "partition");
            }
            return $;
        }
    }

}
