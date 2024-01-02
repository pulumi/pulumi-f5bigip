// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ltm.inputs.GetDataGroupRecordArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDataGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDataGroupArgs Empty = new GetDataGroupArgs();

    /**
     * Name of the datagroup
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the datagroup
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * partition of the datagroup
     * 
     */
    @Import(name="partition", required=true)
    private Output<String> partition;

    /**
     * @return partition of the datagroup
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }

    /**
     * Specifies record of type (string/ip/integer)
     * 
     */
    @Import(name="records")
    private @Nullable Output<List<GetDataGroupRecordArgs>> records;

    /**
     * @return Specifies record of type (string/ip/integer)
     * 
     */
    public Optional<Output<List<GetDataGroupRecordArgs>>> records() {
        return Optional.ofNullable(this.records);
    }

    /**
     * The Data Group type (string, ip, integer)&#34;
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The Data Group type (string, ip, integer)&#34;
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private GetDataGroupArgs() {}

    private GetDataGroupArgs(GetDataGroupArgs $) {
        this.name = $.name;
        this.partition = $.partition;
        this.records = $.records;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDataGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDataGroupArgs $;

        public Builder() {
            $ = new GetDataGroupArgs();
        }

        public Builder(GetDataGroupArgs defaults) {
            $ = new GetDataGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition partition of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder partition(Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition partition of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param records Specifies record of type (string/ip/integer)
         * 
         * @return builder
         * 
         */
        public Builder records(@Nullable Output<List<GetDataGroupRecordArgs>> records) {
            $.records = records;
            return this;
        }

        /**
         * @param records Specifies record of type (string/ip/integer)
         * 
         * @return builder
         * 
         */
        public Builder records(List<GetDataGroupRecordArgs> records) {
            return records(Output.of(records));
        }

        /**
         * @param records Specifies record of type (string/ip/integer)
         * 
         * @return builder
         * 
         */
        public Builder records(GetDataGroupRecordArgs... records) {
            return records(List.of(records));
        }

        /**
         * @param type The Data Group type (string, ip, integer)&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The Data Group type (string, ip, integer)&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetDataGroupArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDataGroupArgs", "name");
            }
            if ($.partition == null) {
                throw new MissingRequiredPropertyException("GetDataGroupArgs", "partition");
            }
            return $;
        }
    }

}
