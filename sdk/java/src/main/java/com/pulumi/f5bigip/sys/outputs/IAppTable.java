// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.f5bigip.sys.outputs.IAppTableRow;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IAppTable {
    private @Nullable List<String> columnNames;
    private @Nullable String encryptedColumns;
    /**
     * @return Name of the iApp.
     * 
     */
    private @Nullable String name;
    private @Nullable List<IAppTableRow> rows;

    private IAppTable() {}
    public List<String> columnNames() {
        return this.columnNames == null ? List.of() : this.columnNames;
    }
    public Optional<String> encryptedColumns() {
        return Optional.ofNullable(this.encryptedColumns);
    }
    /**
     * @return Name of the iApp.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public List<IAppTableRow> rows() {
        return this.rows == null ? List.of() : this.rows;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IAppTable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> columnNames;
        private @Nullable String encryptedColumns;
        private @Nullable String name;
        private @Nullable List<IAppTableRow> rows;
        public Builder() {}
        public Builder(IAppTable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.columnNames = defaults.columnNames;
    	      this.encryptedColumns = defaults.encryptedColumns;
    	      this.name = defaults.name;
    	      this.rows = defaults.rows;
        }

        @CustomType.Setter
        public Builder columnNames(@Nullable List<String> columnNames) {
            this.columnNames = columnNames;
            return this;
        }
        public Builder columnNames(String... columnNames) {
            return columnNames(List.of(columnNames));
        }
        @CustomType.Setter
        public Builder encryptedColumns(@Nullable String encryptedColumns) {
            this.encryptedColumns = encryptedColumns;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder rows(@Nullable List<IAppTableRow> rows) {
            this.rows = rows;
            return this;
        }
        public Builder rows(IAppTableRow... rows) {
            return rows(List.of(rows));
        }
        public IAppTable build() {
            final var _resultValue = new IAppTable();
            _resultValue.columnNames = columnNames;
            _resultValue.encryptedColumns = encryptedColumns;
            _resultValue.name = name;
            _resultValue.rows = rows;
            return _resultValue;
        }
    }
}
