// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IAppMetadata {
    /**
     * @return Name of origin
     * 
     */
    private @Nullable String persists;
    /**
     * @return Name of origin
     * 
     */
    private @Nullable String value;

    private IAppMetadata() {}
    /**
     * @return Name of origin
     * 
     */
    public Optional<String> persists() {
        return Optional.ofNullable(this.persists);
    }
    /**
     * @return Name of origin
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IAppMetadata defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String persists;
        private @Nullable String value;
        public Builder() {}
        public Builder(IAppMetadata defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.persists = defaults.persists;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder persists(@Nullable String persists) {

            this.persists = persists;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {

            this.value = value;
            return this;
        }
        public IAppMetadata build() {
            final var _resultValue = new IAppMetadata();
            _resultValue.persists = persists;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}