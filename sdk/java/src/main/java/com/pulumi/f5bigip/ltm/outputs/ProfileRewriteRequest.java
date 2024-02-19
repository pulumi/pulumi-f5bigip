// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProfileRewriteRequest {
    /**
     * @return Enable to add the X-Forwarded For (XFF) header, to specify the originating IP address of the client. Valid choices are: `enabled, disabled`
     * 
     */
    private @Nullable String insertXfwdFor;
    /**
     * @return Enable to add the X-Forwarded Host header, to specify the originating host of the client. Valid choices are: `enabled, disabled`
     * 
     */
    private @Nullable String insertXfwdHost;
    /**
     * @return Enable to add the X-Forwarded Proto header, to specify the originating protocol of the client. Valid choices are: `enabled, disabled`
     * 
     */
    private @Nullable String insertXfwdProtocol;
    /**
     * @return Enable to rewrite headers in the response. Valid choices are: `enabled, disabled`
     * 
     */
    private @Nullable String rewriteHeaders;

    private ProfileRewriteRequest() {}
    /**
     * @return Enable to add the X-Forwarded For (XFF) header, to specify the originating IP address of the client. Valid choices are: `enabled, disabled`
     * 
     */
    public Optional<String> insertXfwdFor() {
        return Optional.ofNullable(this.insertXfwdFor);
    }
    /**
     * @return Enable to add the X-Forwarded Host header, to specify the originating host of the client. Valid choices are: `enabled, disabled`
     * 
     */
    public Optional<String> insertXfwdHost() {
        return Optional.ofNullable(this.insertXfwdHost);
    }
    /**
     * @return Enable to add the X-Forwarded Proto header, to specify the originating protocol of the client. Valid choices are: `enabled, disabled`
     * 
     */
    public Optional<String> insertXfwdProtocol() {
        return Optional.ofNullable(this.insertXfwdProtocol);
    }
    /**
     * @return Enable to rewrite headers in the response. Valid choices are: `enabled, disabled`
     * 
     */
    public Optional<String> rewriteHeaders() {
        return Optional.ofNullable(this.rewriteHeaders);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProfileRewriteRequest defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String insertXfwdFor;
        private @Nullable String insertXfwdHost;
        private @Nullable String insertXfwdProtocol;
        private @Nullable String rewriteHeaders;
        public Builder() {}
        public Builder(ProfileRewriteRequest defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.insertXfwdFor = defaults.insertXfwdFor;
    	      this.insertXfwdHost = defaults.insertXfwdHost;
    	      this.insertXfwdProtocol = defaults.insertXfwdProtocol;
    	      this.rewriteHeaders = defaults.rewriteHeaders;
        }

        @CustomType.Setter
        public Builder insertXfwdFor(@Nullable String insertXfwdFor) {

            this.insertXfwdFor = insertXfwdFor;
            return this;
        }
        @CustomType.Setter
        public Builder insertXfwdHost(@Nullable String insertXfwdHost) {

            this.insertXfwdHost = insertXfwdHost;
            return this;
        }
        @CustomType.Setter
        public Builder insertXfwdProtocol(@Nullable String insertXfwdProtocol) {

            this.insertXfwdProtocol = insertXfwdProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder rewriteHeaders(@Nullable String rewriteHeaders) {

            this.rewriteHeaders = rewriteHeaders;
            return this;
        }
        public ProfileRewriteRequest build() {
            final var _resultValue = new ProfileRewriteRequest();
            _resultValue.insertXfwdFor = insertXfwdFor;
            _resultValue.insertXfwdHost = insertXfwdHost;
            _resultValue.insertXfwdProtocol = insertXfwdProtocol;
            _resultValue.rewriteHeaders = rewriteHeaders;
            return _resultValue;
        }
    }
}
