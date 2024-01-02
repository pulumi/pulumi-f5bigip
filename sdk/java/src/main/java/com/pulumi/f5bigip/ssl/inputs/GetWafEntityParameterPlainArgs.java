// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ssl.inputs.GetWafEntityParameterUrl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetWafEntityParameterPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetWafEntityParameterPlainArgs Empty = new GetWafEntityParameterPlainArgs();

    @Import(name="allowEmptyType")
    private @Nullable Boolean allowEmptyType;

    public Optional<Boolean> allowEmptyType() {
        return Optional.ofNullable(this.allowEmptyType);
    }

    @Import(name="allowRepeatedParameterName")
    private @Nullable Boolean allowRepeatedParameterName;

    public Optional<Boolean> allowRepeatedParameterName() {
        return Optional.ofNullable(this.allowRepeatedParameterName);
    }

    @Import(name="attackSignaturesCheck")
    private @Nullable Boolean attackSignaturesCheck;

    public Optional<Boolean> attackSignaturesCheck() {
        return Optional.ofNullable(this.attackSignaturesCheck);
    }

    @Import(name="checkMaxValueLength")
    private @Nullable Boolean checkMaxValueLength;

    public Optional<Boolean> checkMaxValueLength() {
        return Optional.ofNullable(this.checkMaxValueLength);
    }

    @Import(name="checkMinValueLength")
    private @Nullable Boolean checkMinValueLength;

    public Optional<Boolean> checkMinValueLength() {
        return Optional.ofNullable(this.checkMinValueLength);
    }

    @Import(name="dataType")
    private @Nullable String dataType;

    public Optional<String> dataType() {
        return Optional.ofNullable(this.dataType);
    }

    @Import(name="description")
    private @Nullable String description;

    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enableRegularExpression")
    private @Nullable Boolean enableRegularExpression;

    public Optional<Boolean> enableRegularExpression() {
        return Optional.ofNullable(this.enableRegularExpression);
    }

    @Import(name="isBase64")
    private @Nullable Boolean isBase64;

    public Optional<Boolean> isBase64() {
        return Optional.ofNullable(this.isBase64);
    }

    @Import(name="isCookie")
    private @Nullable Boolean isCookie;

    public Optional<Boolean> isCookie() {
        return Optional.ofNullable(this.isCookie);
    }

    @Import(name="isHeader")
    private @Nullable Boolean isHeader;

    public Optional<Boolean> isHeader() {
        return Optional.ofNullable(this.isHeader);
    }

    @Import(name="json")
    private @Nullable String json;

    public Optional<String> json() {
        return Optional.ofNullable(this.json);
    }

    @Import(name="level")
    private @Nullable String level;

    public Optional<String> level() {
        return Optional.ofNullable(this.level);
    }

    @Import(name="mandatory")
    private @Nullable Boolean mandatory;

    public Optional<Boolean> mandatory() {
        return Optional.ofNullable(this.mandatory);
    }

    @Import(name="metacharsOnParameterValueCheck")
    private @Nullable Boolean metacharsOnParameterValueCheck;

    public Optional<Boolean> metacharsOnParameterValueCheck() {
        return Optional.ofNullable(this.metacharsOnParameterValueCheck);
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="parameterLocation")
    private @Nullable String parameterLocation;

    public Optional<String> parameterLocation() {
        return Optional.ofNullable(this.parameterLocation);
    }

    @Import(name="performStaging")
    private @Nullable Boolean performStaging;

    public Optional<Boolean> performStaging() {
        return Optional.ofNullable(this.performStaging);
    }

    @Import(name="sensitiveParameter")
    private @Nullable Boolean sensitiveParameter;

    public Optional<Boolean> sensitiveParameter() {
        return Optional.ofNullable(this.sensitiveParameter);
    }

    @Import(name="signatureOverridesDisables")
    private @Nullable List<Integer> signatureOverridesDisables;

    public Optional<List<Integer>> signatureOverridesDisables() {
        return Optional.ofNullable(this.signatureOverridesDisables);
    }

    @Import(name="type")
    private @Nullable String type;

    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="url")
    private @Nullable GetWafEntityParameterUrl url;

    public Optional<GetWafEntityParameterUrl> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="valueType")
    private @Nullable String valueType;

    public Optional<String> valueType() {
        return Optional.ofNullable(this.valueType);
    }

    private GetWafEntityParameterPlainArgs() {}

    private GetWafEntityParameterPlainArgs(GetWafEntityParameterPlainArgs $) {
        this.allowEmptyType = $.allowEmptyType;
        this.allowRepeatedParameterName = $.allowRepeatedParameterName;
        this.attackSignaturesCheck = $.attackSignaturesCheck;
        this.checkMaxValueLength = $.checkMaxValueLength;
        this.checkMinValueLength = $.checkMinValueLength;
        this.dataType = $.dataType;
        this.description = $.description;
        this.enableRegularExpression = $.enableRegularExpression;
        this.isBase64 = $.isBase64;
        this.isCookie = $.isCookie;
        this.isHeader = $.isHeader;
        this.json = $.json;
        this.level = $.level;
        this.mandatory = $.mandatory;
        this.metacharsOnParameterValueCheck = $.metacharsOnParameterValueCheck;
        this.name = $.name;
        this.parameterLocation = $.parameterLocation;
        this.performStaging = $.performStaging;
        this.sensitiveParameter = $.sensitiveParameter;
        this.signatureOverridesDisables = $.signatureOverridesDisables;
        this.type = $.type;
        this.url = $.url;
        this.valueType = $.valueType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetWafEntityParameterPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafEntityParameterPlainArgs $;

        public Builder() {
            $ = new GetWafEntityParameterPlainArgs();
        }

        public Builder(GetWafEntityParameterPlainArgs defaults) {
            $ = new GetWafEntityParameterPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowEmptyType(@Nullable Boolean allowEmptyType) {
            $.allowEmptyType = allowEmptyType;
            return this;
        }

        public Builder allowRepeatedParameterName(@Nullable Boolean allowRepeatedParameterName) {
            $.allowRepeatedParameterName = allowRepeatedParameterName;
            return this;
        }

        public Builder attackSignaturesCheck(@Nullable Boolean attackSignaturesCheck) {
            $.attackSignaturesCheck = attackSignaturesCheck;
            return this;
        }

        public Builder checkMaxValueLength(@Nullable Boolean checkMaxValueLength) {
            $.checkMaxValueLength = checkMaxValueLength;
            return this;
        }

        public Builder checkMinValueLength(@Nullable Boolean checkMinValueLength) {
            $.checkMinValueLength = checkMinValueLength;
            return this;
        }

        public Builder dataType(@Nullable String dataType) {
            $.dataType = dataType;
            return this;
        }

        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        public Builder enableRegularExpression(@Nullable Boolean enableRegularExpression) {
            $.enableRegularExpression = enableRegularExpression;
            return this;
        }

        public Builder isBase64(@Nullable Boolean isBase64) {
            $.isBase64 = isBase64;
            return this;
        }

        public Builder isCookie(@Nullable Boolean isCookie) {
            $.isCookie = isCookie;
            return this;
        }

        public Builder isHeader(@Nullable Boolean isHeader) {
            $.isHeader = isHeader;
            return this;
        }

        public Builder json(@Nullable String json) {
            $.json = json;
            return this;
        }

        public Builder level(@Nullable String level) {
            $.level = level;
            return this;
        }

        public Builder mandatory(@Nullable Boolean mandatory) {
            $.mandatory = mandatory;
            return this;
        }

        public Builder metacharsOnParameterValueCheck(@Nullable Boolean metacharsOnParameterValueCheck) {
            $.metacharsOnParameterValueCheck = metacharsOnParameterValueCheck;
            return this;
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder parameterLocation(@Nullable String parameterLocation) {
            $.parameterLocation = parameterLocation;
            return this;
        }

        public Builder performStaging(@Nullable Boolean performStaging) {
            $.performStaging = performStaging;
            return this;
        }

        public Builder sensitiveParameter(@Nullable Boolean sensitiveParameter) {
            $.sensitiveParameter = sensitiveParameter;
            return this;
        }

        public Builder signatureOverridesDisables(@Nullable List<Integer> signatureOverridesDisables) {
            $.signatureOverridesDisables = signatureOverridesDisables;
            return this;
        }

        public Builder signatureOverridesDisables(Integer... signatureOverridesDisables) {
            return signatureOverridesDisables(List.of(signatureOverridesDisables));
        }

        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public Builder url(@Nullable GetWafEntityParameterUrl url) {
            $.url = url;
            return this;
        }

        public Builder valueType(@Nullable String valueType) {
            $.valueType = valueType;
            return this;
        }

        public GetWafEntityParameterPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterPlainArgs", "name");
            }
            return $;
        }
    }

}
