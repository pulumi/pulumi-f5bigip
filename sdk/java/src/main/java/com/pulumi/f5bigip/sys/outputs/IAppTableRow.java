// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class IAppTableRow {
    private @Nullable List<String> rows;

    private IAppTableRow() {}
    public List<String> rows() {
        return this.rows == null ? List.of() : this.rows;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IAppTableRow defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> rows;
        public Builder() {}
        public Builder(IAppTableRow defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.rows = defaults.rows;
        }

        @CustomType.Setter
        public Builder rows(@Nullable List<String> rows) {

            this.rows = rows;
            return this;
        }
        public Builder rows(String... rows) {
            return rows(List.of(rows));
        }
        public IAppTableRow build() {
            final var _resultValue = new IAppTableRow();
            _resultValue.rows = rows;
            return _resultValue;
        }
    }
}
