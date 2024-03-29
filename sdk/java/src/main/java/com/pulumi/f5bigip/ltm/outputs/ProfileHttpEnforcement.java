// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProfileHttpEnforcement {
    /**
     * @return Specifies which HTTP methods count as being known. Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. Default value is [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK].If no value is specified while creating, then default value will be assigned by BigIP. In order to remove it, [&#34;&#34;] list is to be passed. If known_methods is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value , we need to pass [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK] explicitly.
     * 
     */
    private @Nullable List<String> knownMethods;
    /**
     * @return Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers.If no value is specified while creating, then default value will be assigned by BigIP. If max_header_count is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;64&#34; explicitly.
     * 
     */
    private @Nullable Integer maxHeaderCount;
    /**
     * @return Specifies the maximum header size. The default value is 32768. If no string is specified while creating, then default value will be assigned by BigIP. If max_header_size is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;32768&#34; explicitly.
     * 
     */
    private @Nullable Integer maxHeaderSize;
    /**
     * @return Specifies whether to allow, reject or switch to pass-through mode when an unknown HTTP method is parsed. Default value is &#34;allow&#34;. If no string is specified while creating, then default value will be assigned by BigIP. If unknown_method is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;allow&#34; explicitly.
     * 
     */
    private @Nullable String unknownMethod;

    private ProfileHttpEnforcement() {}
    /**
     * @return Specifies which HTTP methods count as being known. Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. Default value is [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK].If no value is specified while creating, then default value will be assigned by BigIP. In order to remove it, [&#34;&#34;] list is to be passed. If known_methods is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value , we need to pass [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK] explicitly.
     * 
     */
    public List<String> knownMethods() {
        return this.knownMethods == null ? List.of() : this.knownMethods;
    }
    /**
     * @return Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers.If no value is specified while creating, then default value will be assigned by BigIP. If max_header_count is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;64&#34; explicitly.
     * 
     */
    public Optional<Integer> maxHeaderCount() {
        return Optional.ofNullable(this.maxHeaderCount);
    }
    /**
     * @return Specifies the maximum header size. The default value is 32768. If no string is specified while creating, then default value will be assigned by BigIP. If max_header_size is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;32768&#34; explicitly.
     * 
     */
    public Optional<Integer> maxHeaderSize() {
        return Optional.ofNullable(this.maxHeaderSize);
    }
    /**
     * @return Specifies whether to allow, reject or switch to pass-through mode when an unknown HTTP method is parsed. Default value is &#34;allow&#34;. If no string is specified while creating, then default value will be assigned by BigIP. If unknown_method is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass &#34;allow&#34; explicitly.
     * 
     */
    public Optional<String> unknownMethod() {
        return Optional.ofNullable(this.unknownMethod);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProfileHttpEnforcement defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> knownMethods;
        private @Nullable Integer maxHeaderCount;
        private @Nullable Integer maxHeaderSize;
        private @Nullable String unknownMethod;
        public Builder() {}
        public Builder(ProfileHttpEnforcement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.knownMethods = defaults.knownMethods;
    	      this.maxHeaderCount = defaults.maxHeaderCount;
    	      this.maxHeaderSize = defaults.maxHeaderSize;
    	      this.unknownMethod = defaults.unknownMethod;
        }

        @CustomType.Setter
        public Builder knownMethods(@Nullable List<String> knownMethods) {

            this.knownMethods = knownMethods;
            return this;
        }
        public Builder knownMethods(String... knownMethods) {
            return knownMethods(List.of(knownMethods));
        }
        @CustomType.Setter
        public Builder maxHeaderCount(@Nullable Integer maxHeaderCount) {

            this.maxHeaderCount = maxHeaderCount;
            return this;
        }
        @CustomType.Setter
        public Builder maxHeaderSize(@Nullable Integer maxHeaderSize) {

            this.maxHeaderSize = maxHeaderSize;
            return this;
        }
        @CustomType.Setter
        public Builder unknownMethod(@Nullable String unknownMethod) {

            this.unknownMethod = unknownMethod;
            return this;
        }
        public ProfileHttpEnforcement build() {
            final var _resultValue = new ProfileHttpEnforcement();
            _resultValue.knownMethods = knownMethods;
            _resultValue.maxHeaderCount = maxHeaderCount;
            _resultValue.maxHeaderSize = maxHeaderSize;
            _resultValue.unknownMethod = unknownMethod;
            return _resultValue;
        }
    }
}
