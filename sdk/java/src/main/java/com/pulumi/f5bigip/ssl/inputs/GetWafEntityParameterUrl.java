// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetWafEntityParameterUrl extends com.pulumi.resources.InvokeArgs {

    public static final GetWafEntityParameterUrl Empty = new GetWafEntityParameterUrl();

    @Import(name="method", required=true)
    private String method;

    public String method() {
        return this.method;
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="protocol", required=true)
    private String protocol;

    public String protocol() {
        return this.protocol;
    }

    @Import(name="type", required=true)
    private String type;

    public String type() {
        return this.type;
    }

    private GetWafEntityParameterUrl() {}

    private GetWafEntityParameterUrl(GetWafEntityParameterUrl $) {
        this.method = $.method;
        this.name = $.name;
        this.protocol = $.protocol;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWafEntityParameterUrl defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafEntityParameterUrl $;

        public Builder() {
            $ = new GetWafEntityParameterUrl();
        }

        public Builder(GetWafEntityParameterUrl defaults) {
            $ = new GetWafEntityParameterUrl(Objects.requireNonNull(defaults));
        }

        public Builder method(String method) {
            $.method = method;
            return this;
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder protocol(String protocol) {
            $.protocol = protocol;
            return this;
        }

        public Builder type(String type) {
            $.type = type;
            return this;
        }

        public GetWafEntityParameterUrl build() {
            if ($.method == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterUrl", "method");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterUrl", "name");
            }
            if ($.protocol == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterUrl", "protocol");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterUrl", "type");
            }
            return $;
        }
    }

}