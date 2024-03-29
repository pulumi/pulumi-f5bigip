// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.sys.inputs.IAppTableRowArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IAppTableArgs extends com.pulumi.resources.ResourceArgs {

    public static final IAppTableArgs Empty = new IAppTableArgs();

    @Import(name="columnNames")
    private @Nullable Output<List<String>> columnNames;

    public Optional<Output<List<String>>> columnNames() {
        return Optional.ofNullable(this.columnNames);
    }

    /**
     * Name of origin
     * 
     */
    @Import(name="encryptedColumns")
    private @Nullable Output<String> encryptedColumns;

    /**
     * @return Name of origin
     * 
     */
    public Optional<Output<String>> encryptedColumns() {
        return Optional.ofNullable(this.encryptedColumns);
    }

    /**
     * Name of the iApp.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the iApp.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="rows")
    private @Nullable Output<List<IAppTableRowArgs>> rows;

    public Optional<Output<List<IAppTableRowArgs>>> rows() {
        return Optional.ofNullable(this.rows);
    }

    private IAppTableArgs() {}

    private IAppTableArgs(IAppTableArgs $) {
        this.columnNames = $.columnNames;
        this.encryptedColumns = $.encryptedColumns;
        this.name = $.name;
        this.rows = $.rows;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IAppTableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IAppTableArgs $;

        public Builder() {
            $ = new IAppTableArgs();
        }

        public Builder(IAppTableArgs defaults) {
            $ = new IAppTableArgs(Objects.requireNonNull(defaults));
        }

        public Builder columnNames(@Nullable Output<List<String>> columnNames) {
            $.columnNames = columnNames;
            return this;
        }

        public Builder columnNames(List<String> columnNames) {
            return columnNames(Output.of(columnNames));
        }

        public Builder columnNames(String... columnNames) {
            return columnNames(List.of(columnNames));
        }

        /**
         * @param encryptedColumns Name of origin
         * 
         * @return builder
         * 
         */
        public Builder encryptedColumns(@Nullable Output<String> encryptedColumns) {
            $.encryptedColumns = encryptedColumns;
            return this;
        }

        /**
         * @param encryptedColumns Name of origin
         * 
         * @return builder
         * 
         */
        public Builder encryptedColumns(String encryptedColumns) {
            return encryptedColumns(Output.of(encryptedColumns));
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder rows(@Nullable Output<List<IAppTableRowArgs>> rows) {
            $.rows = rows;
            return this;
        }

        public Builder rows(List<IAppTableRowArgs> rows) {
            return rows(Output.of(rows));
        }

        public Builder rows(IAppTableRowArgs... rows) {
            return rows(List.of(rows));
        }

        public IAppTableArgs build() {
            return $;
        }
    }

}
