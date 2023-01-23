// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetWafEntityUrlMethodOverride extends com.pulumi.resources.InvokeArgs {

    public static final GetWafEntityUrlMethodOverride Empty = new GetWafEntityUrlMethodOverride();

    /**
     * Specifies that the system allows or disallows a method for this URL
     * 
     */
    @Import(name="allow", required=true)
    private Boolean allow;

    /**
     * @return Specifies that the system allows or disallows a method for this URL
     * 
     */
    public Boolean allow() {
        return this.allow;
    }

    /**
     * Specifies an HTTP method.
     * 
     */
    @Import(name="method", required=true)
    private String method;

    /**
     * @return Specifies an HTTP method.
     * 
     */
    public String method() {
        return this.method;
    }

    private GetWafEntityUrlMethodOverride() {}

    private GetWafEntityUrlMethodOverride(GetWafEntityUrlMethodOverride $) {
        this.allow = $.allow;
        this.method = $.method;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWafEntityUrlMethodOverride defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafEntityUrlMethodOverride $;

        public Builder() {
            $ = new GetWafEntityUrlMethodOverride();
        }

        public Builder(GetWafEntityUrlMethodOverride defaults) {
            $ = new GetWafEntityUrlMethodOverride(Objects.requireNonNull(defaults));
        }

        /**
         * @param allow Specifies that the system allows or disallows a method for this URL
         * 
         * @return builder
         * 
         */
        public Builder allow(Boolean allow) {
            $.allow = allow;
            return this;
        }

        /**
         * @param method Specifies an HTTP method.
         * 
         * @return builder
         * 
         */
        public Builder method(String method) {
            $.method = method;
            return this;
        }

        public GetWafEntityUrlMethodOverride build() {
            $.allow = Objects.requireNonNull($.allow, "expected parameter 'allow' to be non-null");
            $.method = Objects.requireNonNull($.method, "expected parameter 'method' to be non-null");
            return $;
        }
    }

}