// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetWafPbSuggestionsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Json string representing exported PB suggestions ready to be used in WAF policy declaration
     * 
     */
    private String json;
    private Integer minimumLearningScore;
    private String partition;
    /**
     * @return System generated id of the WAF policy
     * 
     */
    private String policyId;
    private String policyName;

    private GetWafPbSuggestionsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Json string representing exported PB suggestions ready to be used in WAF policy declaration
     * 
     */
    public String json() {
        return this.json;
    }
    public Integer minimumLearningScore() {
        return this.minimumLearningScore;
    }
    public String partition() {
        return this.partition;
    }
    /**
     * @return System generated id of the WAF policy
     * 
     */
    public String policyId() {
        return this.policyId;
    }
    public String policyName() {
        return this.policyName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetWafPbSuggestionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String json;
        private Integer minimumLearningScore;
        private String partition;
        private String policyId;
        private String policyName;
        public Builder() {}
        public Builder(GetWafPbSuggestionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.json = defaults.json;
    	      this.minimumLearningScore = defaults.minimumLearningScore;
    	      this.partition = defaults.partition;
    	      this.policyId = defaults.policyId;
    	      this.policyName = defaults.policyName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder json(String json) {
            if (json == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "json");
            }
            this.json = json;
            return this;
        }
        @CustomType.Setter
        public Builder minimumLearningScore(Integer minimumLearningScore) {
            if (minimumLearningScore == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "minimumLearningScore");
            }
            this.minimumLearningScore = minimumLearningScore;
            return this;
        }
        @CustomType.Setter
        public Builder partition(String partition) {
            if (partition == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "partition");
            }
            this.partition = partition;
            return this;
        }
        @CustomType.Setter
        public Builder policyId(String policyId) {
            if (policyId == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "policyId");
            }
            this.policyId = policyId;
            return this;
        }
        @CustomType.Setter
        public Builder policyName(String policyName) {
            if (policyName == null) {
              throw new MissingRequiredPropertyException("GetWafPbSuggestionsResult", "policyName");
            }
            this.policyName = policyName;
            return this;
        }
        public GetWafPbSuggestionsResult build() {
            final var _resultValue = new GetWafPbSuggestionsResult();
            _resultValue.id = id;
            _resultValue.json = json;
            _resultValue.minimumLearningScore = minimumLearningScore;
            _resultValue.partition = partition;
            _resultValue.policyId = policyId;
            _resultValue.policyName = policyName;
            return _resultValue;
        }
    }
}
