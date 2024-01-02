// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ltm.inputs.GetDataGroupRecord;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDataGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDataGroupPlainArgs Empty = new GetDataGroupPlainArgs();

    /**
     * Name of the datagroup
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the datagroup
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * partition of the datagroup
     * 
     */
    @Import(name="partition", required=true)
    private String partition;

    /**
     * @return partition of the datagroup
     * 
     */
    public String partition() {
        return this.partition;
    }

    /**
     * Specifies record of type (string/ip/integer)
     * 
     */
    @Import(name="records")
    private @Nullable List<GetDataGroupRecord> records;

    /**
     * @return Specifies record of type (string/ip/integer)
     * 
     */
    public Optional<List<GetDataGroupRecord>> records() {
        return Optional.ofNullable(this.records);
    }

    /**
     * The Data Group type (string, ip, integer)&#34;
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The Data Group type (string, ip, integer)&#34;
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetDataGroupPlainArgs() {}

    private GetDataGroupPlainArgs(GetDataGroupPlainArgs $) {
        this.name = $.name;
        this.partition = $.partition;
        this.records = $.records;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDataGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDataGroupPlainArgs $;

        public Builder() {
            $ = new GetDataGroupPlainArgs();
        }

        public Builder(GetDataGroupPlainArgs defaults) {
            $ = new GetDataGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param partition partition of the datagroup
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param records Specifies record of type (string/ip/integer)
         * 
         * @return builder
         * 
         */
        public Builder records(@Nullable List<GetDataGroupRecord> records) {
            $.records = records;
            return this;
        }

        /**
         * @param records Specifies record of type (string/ip/integer)
         * 
         * @return builder
         * 
         */
        public Builder records(GetDataGroupRecord... records) {
            return records(List.of(records));
        }

        /**
         * @param type The Data Group type (string, ip, integer)&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetDataGroupPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDataGroupPlainArgs", "name");
            }
            if ($.partition == null) {
                throw new MissingRequiredPropertyException("GetDataGroupPlainArgs", "partition");
            }
            return $;
        }
    }

}
