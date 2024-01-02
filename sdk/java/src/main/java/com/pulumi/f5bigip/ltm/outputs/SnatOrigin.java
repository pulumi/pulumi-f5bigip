// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SnatOrigin {
    private @Nullable String appService;
    /**
     * @return Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
     * 
     */
    private @Nullable String name;

    private SnatOrigin() {}
    public Optional<String> appService() {
        return Optional.ofNullable(this.appService);
    }
    /**
     * @return Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SnatOrigin defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String appService;
        private @Nullable String name;
        public Builder() {}
        public Builder(SnatOrigin defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appService = defaults.appService;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder appService(@Nullable String appService) {

            this.appService = appService;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        public SnatOrigin build() {
            final var _resultValue = new SnatOrigin();
            _resultValue.appService = appService;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
