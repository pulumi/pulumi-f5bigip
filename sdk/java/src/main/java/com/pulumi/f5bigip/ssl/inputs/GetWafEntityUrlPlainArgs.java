// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ssl.inputs.GetWafEntityUrlMethodOverride;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetWafEntityUrlPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetWafEntityUrlPlainArgs Empty = new GetWafEntityUrlPlainArgs();

    /**
     * A description of the URL.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return A description of the URL.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Select a Method for the URL to create an API endpoint. Default is : *.
     * 
     */
    @Import(name="method")
    private @Nullable String method;

    /**
     * @return Select a Method for the URL to create an API endpoint. Default is : *.
     * 
     */
    public Optional<String> method() {
        return Optional.ofNullable(this.method);
    }

    /**
     * A list of methods that are allowed or disallowed for a specific URL.
     * 
     */
    @Import(name="methodOverrides")
    private @Nullable List<GetWafEntityUrlMethodOverride> methodOverrides;

    /**
     * @return A list of methods that are allowed or disallowed for a specific URL.
     * 
     */
    public Optional<List<GetWafEntityUrlMethodOverride>> methodOverrides() {
        return Optional.ofNullable(this.methodOverrides);
    }

    /**
     * WAF entity URL name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return WAF entity URL name.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
     * 
     */
    @Import(name="performStaging")
    private @Nullable Boolean performStaging;

    /**
     * @return If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
     * 
     */
    public Optional<Boolean> performStaging() {
        return Optional.ofNullable(this.performStaging);
    }

    /**
     * Specifies whether the protocol for the URL is &#39;http&#39; or &#39;https&#39;. Default is: http.
     * 
     */
    @Import(name="protocol")
    private @Nullable String protocol;

    /**
     * @return Specifies whether the protocol for the URL is &#39;http&#39; or &#39;https&#39;. Default is: http.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * List of Attack Signature Ids which are disabled for this particular URL.
     * 
     */
    @Import(name="signatureOverridesDisables")
    private @Nullable List<Integer> signatureOverridesDisables;

    /**
     * @return List of Attack Signature Ids which are disabled for this particular URL.
     * 
     */
    public Optional<List<Integer>> signatureOverridesDisables() {
        return Optional.ofNullable(this.signatureOverridesDisables);
    }

    /**
     * Specifies whether the parameter is an &#39;explicit&#39; or a &#39;wildcard&#39; attribute. Default is: wildcard.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return Specifies whether the parameter is an &#39;explicit&#39; or a &#39;wildcard&#39; attribute. Default is: wildcard.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetWafEntityUrlPlainArgs() {}

    private GetWafEntityUrlPlainArgs(GetWafEntityUrlPlainArgs $) {
        this.description = $.description;
        this.method = $.method;
        this.methodOverrides = $.methodOverrides;
        this.name = $.name;
        this.performStaging = $.performStaging;
        this.protocol = $.protocol;
        this.signatureOverridesDisables = $.signatureOverridesDisables;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWafEntityUrlPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafEntityUrlPlainArgs $;

        public Builder() {
            $ = new GetWafEntityUrlPlainArgs();
        }

        public Builder(GetWafEntityUrlPlainArgs defaults) {
            $ = new GetWafEntityUrlPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description of the URL.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param method Select a Method for the URL to create an API endpoint. Default is : *.
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable String method) {
            $.method = method;
            return this;
        }

        /**
         * @param methodOverrides A list of methods that are allowed or disallowed for a specific URL.
         * 
         * @return builder
         * 
         */
        public Builder methodOverrides(@Nullable List<GetWafEntityUrlMethodOverride> methodOverrides) {
            $.methodOverrides = methodOverrides;
            return this;
        }

        /**
         * @param methodOverrides A list of methods that are allowed or disallowed for a specific URL.
         * 
         * @return builder
         * 
         */
        public Builder methodOverrides(GetWafEntityUrlMethodOverride... methodOverrides) {
            return methodOverrides(List.of(methodOverrides));
        }

        /**
         * @param name WAF entity URL name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param performStaging If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
         * 
         * @return builder
         * 
         */
        public Builder performStaging(@Nullable Boolean performStaging) {
            $.performStaging = performStaging;
            return this;
        }

        /**
         * @param protocol Specifies whether the protocol for the URL is &#39;http&#39; or &#39;https&#39;. Default is: http.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable String protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param signatureOverridesDisables List of Attack Signature Ids which are disabled for this particular URL.
         * 
         * @return builder
         * 
         */
        public Builder signatureOverridesDisables(@Nullable List<Integer> signatureOverridesDisables) {
            $.signatureOverridesDisables = signatureOverridesDisables;
            return this;
        }

        /**
         * @param signatureOverridesDisables List of Attack Signature Ids which are disabled for this particular URL.
         * 
         * @return builder
         * 
         */
        public Builder signatureOverridesDisables(Integer... signatureOverridesDisables) {
            return signatureOverridesDisables(List.of(signatureOverridesDisables));
        }

        /**
         * @param type Specifies whether the parameter is an &#39;explicit&#39; or a &#39;wildcard&#39; attribute. Default is: wildcard.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetWafEntityUrlPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetWafEntityUrlPlainArgs", "name");
            }
            return $;
        }
    }

}
