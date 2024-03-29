// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.ssl.inputs.GetWafEntityParameterUrlArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetWafEntityParameterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetWafEntityParameterArgs Empty = new GetWafEntityParameterArgs();

    @Import(name="allowEmptyType")
    private @Nullable Output<Boolean> allowEmptyType;

    public Optional<Output<Boolean>> allowEmptyType() {
        return Optional.ofNullable(this.allowEmptyType);
    }

    @Import(name="allowRepeatedParameterName")
    private @Nullable Output<Boolean> allowRepeatedParameterName;

    public Optional<Output<Boolean>> allowRepeatedParameterName() {
        return Optional.ofNullable(this.allowRepeatedParameterName);
    }

    @Import(name="attackSignaturesCheck")
    private @Nullable Output<Boolean> attackSignaturesCheck;

    public Optional<Output<Boolean>> attackSignaturesCheck() {
        return Optional.ofNullable(this.attackSignaturesCheck);
    }

    @Import(name="checkMaxValueLength")
    private @Nullable Output<Boolean> checkMaxValueLength;

    public Optional<Output<Boolean>> checkMaxValueLength() {
        return Optional.ofNullable(this.checkMaxValueLength);
    }

    @Import(name="checkMinValueLength")
    private @Nullable Output<Boolean> checkMinValueLength;

    public Optional<Output<Boolean>> checkMinValueLength() {
        return Optional.ofNullable(this.checkMinValueLength);
    }

    @Import(name="dataType")
    private @Nullable Output<String> dataType;

    public Optional<Output<String>> dataType() {
        return Optional.ofNullable(this.dataType);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enableRegularExpression")
    private @Nullable Output<Boolean> enableRegularExpression;

    public Optional<Output<Boolean>> enableRegularExpression() {
        return Optional.ofNullable(this.enableRegularExpression);
    }

    @Import(name="isBase64")
    private @Nullable Output<Boolean> isBase64;

    public Optional<Output<Boolean>> isBase64() {
        return Optional.ofNullable(this.isBase64);
    }

    @Import(name="isCookie")
    private @Nullable Output<Boolean> isCookie;

    public Optional<Output<Boolean>> isCookie() {
        return Optional.ofNullable(this.isCookie);
    }

    @Import(name="isHeader")
    private @Nullable Output<Boolean> isHeader;

    public Optional<Output<Boolean>> isHeader() {
        return Optional.ofNullable(this.isHeader);
    }

    @Import(name="json")
    private @Nullable Output<String> json;

    public Optional<Output<String>> json() {
        return Optional.ofNullable(this.json);
    }

    @Import(name="level")
    private @Nullable Output<String> level;

    public Optional<Output<String>> level() {
        return Optional.ofNullable(this.level);
    }

    @Import(name="mandatory")
    private @Nullable Output<Boolean> mandatory;

    public Optional<Output<Boolean>> mandatory() {
        return Optional.ofNullable(this.mandatory);
    }

    @Import(name="maxValueLength")
    private @Nullable Output<Integer> maxValueLength;

    public Optional<Output<Integer>> maxValueLength() {
        return Optional.ofNullable(this.maxValueLength);
    }

    @Import(name="metacharsOnParameterValueCheck")
    private @Nullable Output<Boolean> metacharsOnParameterValueCheck;

    public Optional<Output<Boolean>> metacharsOnParameterValueCheck() {
        return Optional.ofNullable(this.metacharsOnParameterValueCheck);
    }

    @Import(name="minValueLength")
    private @Nullable Output<Integer> minValueLength;

    public Optional<Output<Integer>> minValueLength() {
        return Optional.ofNullable(this.minValueLength);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    @Import(name="parameterLocation")
    private @Nullable Output<String> parameterLocation;

    public Optional<Output<String>> parameterLocation() {
        return Optional.ofNullable(this.parameterLocation);
    }

    @Import(name="performStaging")
    private @Nullable Output<Boolean> performStaging;

    public Optional<Output<Boolean>> performStaging() {
        return Optional.ofNullable(this.performStaging);
    }

    @Import(name="sensitiveParameter")
    private @Nullable Output<Boolean> sensitiveParameter;

    public Optional<Output<Boolean>> sensitiveParameter() {
        return Optional.ofNullable(this.sensitiveParameter);
    }

    @Import(name="signatureOverridesDisables")
    private @Nullable Output<List<Integer>> signatureOverridesDisables;

    public Optional<Output<List<Integer>>> signatureOverridesDisables() {
        return Optional.ofNullable(this.signatureOverridesDisables);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="url")
    private @Nullable Output<GetWafEntityParameterUrlArgs> url;

    public Optional<Output<GetWafEntityParameterUrlArgs>> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="valueType")
    private @Nullable Output<String> valueType;

    public Optional<Output<String>> valueType() {
        return Optional.ofNullable(this.valueType);
    }

    private GetWafEntityParameterArgs() {}

    private GetWafEntityParameterArgs(GetWafEntityParameterArgs $) {
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
        this.maxValueLength = $.maxValueLength;
        this.metacharsOnParameterValueCheck = $.metacharsOnParameterValueCheck;
        this.minValueLength = $.minValueLength;
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
    public static Builder builder(GetWafEntityParameterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetWafEntityParameterArgs $;

        public Builder() {
            $ = new GetWafEntityParameterArgs();
        }

        public Builder(GetWafEntityParameterArgs defaults) {
            $ = new GetWafEntityParameterArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowEmptyType(@Nullable Output<Boolean> allowEmptyType) {
            $.allowEmptyType = allowEmptyType;
            return this;
        }

        public Builder allowEmptyType(Boolean allowEmptyType) {
            return allowEmptyType(Output.of(allowEmptyType));
        }

        public Builder allowRepeatedParameterName(@Nullable Output<Boolean> allowRepeatedParameterName) {
            $.allowRepeatedParameterName = allowRepeatedParameterName;
            return this;
        }

        public Builder allowRepeatedParameterName(Boolean allowRepeatedParameterName) {
            return allowRepeatedParameterName(Output.of(allowRepeatedParameterName));
        }

        public Builder attackSignaturesCheck(@Nullable Output<Boolean> attackSignaturesCheck) {
            $.attackSignaturesCheck = attackSignaturesCheck;
            return this;
        }

        public Builder attackSignaturesCheck(Boolean attackSignaturesCheck) {
            return attackSignaturesCheck(Output.of(attackSignaturesCheck));
        }

        public Builder checkMaxValueLength(@Nullable Output<Boolean> checkMaxValueLength) {
            $.checkMaxValueLength = checkMaxValueLength;
            return this;
        }

        public Builder checkMaxValueLength(Boolean checkMaxValueLength) {
            return checkMaxValueLength(Output.of(checkMaxValueLength));
        }

        public Builder checkMinValueLength(@Nullable Output<Boolean> checkMinValueLength) {
            $.checkMinValueLength = checkMinValueLength;
            return this;
        }

        public Builder checkMinValueLength(Boolean checkMinValueLength) {
            return checkMinValueLength(Output.of(checkMinValueLength));
        }

        public Builder dataType(@Nullable Output<String> dataType) {
            $.dataType = dataType;
            return this;
        }

        public Builder dataType(String dataType) {
            return dataType(Output.of(dataType));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder enableRegularExpression(@Nullable Output<Boolean> enableRegularExpression) {
            $.enableRegularExpression = enableRegularExpression;
            return this;
        }

        public Builder enableRegularExpression(Boolean enableRegularExpression) {
            return enableRegularExpression(Output.of(enableRegularExpression));
        }

        public Builder isBase64(@Nullable Output<Boolean> isBase64) {
            $.isBase64 = isBase64;
            return this;
        }

        public Builder isBase64(Boolean isBase64) {
            return isBase64(Output.of(isBase64));
        }

        public Builder isCookie(@Nullable Output<Boolean> isCookie) {
            $.isCookie = isCookie;
            return this;
        }

        public Builder isCookie(Boolean isCookie) {
            return isCookie(Output.of(isCookie));
        }

        public Builder isHeader(@Nullable Output<Boolean> isHeader) {
            $.isHeader = isHeader;
            return this;
        }

        public Builder isHeader(Boolean isHeader) {
            return isHeader(Output.of(isHeader));
        }

        public Builder json(@Nullable Output<String> json) {
            $.json = json;
            return this;
        }

        public Builder json(String json) {
            return json(Output.of(json));
        }

        public Builder level(@Nullable Output<String> level) {
            $.level = level;
            return this;
        }

        public Builder level(String level) {
            return level(Output.of(level));
        }

        public Builder mandatory(@Nullable Output<Boolean> mandatory) {
            $.mandatory = mandatory;
            return this;
        }

        public Builder mandatory(Boolean mandatory) {
            return mandatory(Output.of(mandatory));
        }

        public Builder maxValueLength(@Nullable Output<Integer> maxValueLength) {
            $.maxValueLength = maxValueLength;
            return this;
        }

        public Builder maxValueLength(Integer maxValueLength) {
            return maxValueLength(Output.of(maxValueLength));
        }

        public Builder metacharsOnParameterValueCheck(@Nullable Output<Boolean> metacharsOnParameterValueCheck) {
            $.metacharsOnParameterValueCheck = metacharsOnParameterValueCheck;
            return this;
        }

        public Builder metacharsOnParameterValueCheck(Boolean metacharsOnParameterValueCheck) {
            return metacharsOnParameterValueCheck(Output.of(metacharsOnParameterValueCheck));
        }

        public Builder minValueLength(@Nullable Output<Integer> minValueLength) {
            $.minValueLength = minValueLength;
            return this;
        }

        public Builder minValueLength(Integer minValueLength) {
            return minValueLength(Output.of(minValueLength));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder parameterLocation(@Nullable Output<String> parameterLocation) {
            $.parameterLocation = parameterLocation;
            return this;
        }

        public Builder parameterLocation(String parameterLocation) {
            return parameterLocation(Output.of(parameterLocation));
        }

        public Builder performStaging(@Nullable Output<Boolean> performStaging) {
            $.performStaging = performStaging;
            return this;
        }

        public Builder performStaging(Boolean performStaging) {
            return performStaging(Output.of(performStaging));
        }

        public Builder sensitiveParameter(@Nullable Output<Boolean> sensitiveParameter) {
            $.sensitiveParameter = sensitiveParameter;
            return this;
        }

        public Builder sensitiveParameter(Boolean sensitiveParameter) {
            return sensitiveParameter(Output.of(sensitiveParameter));
        }

        public Builder signatureOverridesDisables(@Nullable Output<List<Integer>> signatureOverridesDisables) {
            $.signatureOverridesDisables = signatureOverridesDisables;
            return this;
        }

        public Builder signatureOverridesDisables(List<Integer> signatureOverridesDisables) {
            return signatureOverridesDisables(Output.of(signatureOverridesDisables));
        }

        public Builder signatureOverridesDisables(Integer... signatureOverridesDisables) {
            return signatureOverridesDisables(List.of(signatureOverridesDisables));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder url(@Nullable Output<GetWafEntityParameterUrlArgs> url) {
            $.url = url;
            return this;
        }

        public Builder url(GetWafEntityParameterUrlArgs url) {
            return url(Output.of(url));
        }

        public Builder valueType(@Nullable Output<String> valueType) {
            $.valueType = valueType;
            return this;
        }

        public Builder valueType(String valueType) {
            return valueType(Output.of(valueType));
        }

        public GetWafEntityParameterArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetWafEntityParameterArgs", "name");
            }
            return $;
        }
    }

}
