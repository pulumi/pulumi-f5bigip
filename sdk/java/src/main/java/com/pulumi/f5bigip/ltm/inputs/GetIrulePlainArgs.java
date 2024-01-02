// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIrulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIrulePlainArgs Empty = new GetIrulePlainArgs();

    /**
     * Irule configured on bigip
     * 
     */
    @Import(name="irule")
    private @Nullable String irule;

    /**
     * @return Irule configured on bigip
     * 
     */
    public Optional<String> irule() {
        return Optional.ofNullable(this.irule);
    }

    /**
     * Name of the irule
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the irule
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * partition of the ltm irule
     * 
     */
    @Import(name="partition", required=true)
    private String partition;

    /**
     * @return partition of the ltm irule
     * 
     */
    public String partition() {
        return this.partition;
    }

    private GetIrulePlainArgs() {}

    private GetIrulePlainArgs(GetIrulePlainArgs $) {
        this.irule = $.irule;
        this.name = $.name;
        this.partition = $.partition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIrulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIrulePlainArgs $;

        public Builder() {
            $ = new GetIrulePlainArgs();
        }

        public Builder(GetIrulePlainArgs defaults) {
            $ = new GetIrulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param irule Irule configured on bigip
         * 
         * @return builder
         * 
         */
        public Builder irule(@Nullable String irule) {
            $.irule = irule;
            return this;
        }

        /**
         * @param name Name of the irule
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param partition partition of the ltm irule
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            $.partition = partition;
            return this;
        }

        public GetIrulePlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetIrulePlainArgs", "name");
            }
            if ($.partition == null) {
                throw new MissingRequiredPropertyException("GetIrulePlainArgs", "partition");
            }
            return $;
        }
    }

}
