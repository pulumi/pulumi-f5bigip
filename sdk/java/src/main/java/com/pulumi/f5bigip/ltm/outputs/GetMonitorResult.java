// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMonitorResult {
    /**
     * @return Displays whether adaptive response time monitoring is enabled for this monitor.
     * 
     */
    private String adaptive;
    /**
     * @return Displays whether adaptive response time monitoring is enabled for this monitor.
     * 
     */
    private Integer adaptiveLimit;
    private String base;
    private String chaseReferrals;
    private String database;
    private String defaultsFrom;
    /**
     * @return id will be full path name of ltm monitor.
     * 
     */
    private String destination;
    private String filename;
    private String filter;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.
     * 
     */
    private Integer interval;
    /**
     * @return Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.
     * 
     */
    private Integer ipDscp;
    private String mandatoryAttributes;
    /**
     * @return Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     * 
     */
    private String manualResume;
    private String mode;
    private String name;
    private String partition;
    private String receiveDisable;
    /**
     * @return Instructs the system to mark the target resource down when the test is successful.
     * 
     */
    private String reverse;
    private String security;
    private Integer timeUntilUp;
    private Integer timeout;
    /**
     * @return Displays whether the monitor operates in transparent mode.
     * 
     */
    private String transparent;
    private String username;

    private GetMonitorResult() {}
    /**
     * @return Displays whether adaptive response time monitoring is enabled for this monitor.
     * 
     */
    public String adaptive() {
        return this.adaptive;
    }
    /**
     * @return Displays whether adaptive response time monitoring is enabled for this monitor.
     * 
     */
    public Integer adaptiveLimit() {
        return this.adaptiveLimit;
    }
    public String base() {
        return this.base;
    }
    public String chaseReferrals() {
        return this.chaseReferrals;
    }
    public String database() {
        return this.database;
    }
    public String defaultsFrom() {
        return this.defaultsFrom;
    }
    /**
     * @return id will be full path name of ltm monitor.
     * 
     */
    public String destination() {
        return this.destination;
    }
    public String filename() {
        return this.filename;
    }
    public String filter() {
        return this.filter;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.
     * 
     */
    public Integer interval() {
        return this.interval;
    }
    /**
     * @return Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.
     * 
     */
    public Integer ipDscp() {
        return this.ipDscp;
    }
    public String mandatoryAttributes() {
        return this.mandatoryAttributes;
    }
    /**
     * @return Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     * 
     */
    public String manualResume() {
        return this.manualResume;
    }
    public String mode() {
        return this.mode;
    }
    public String name() {
        return this.name;
    }
    public String partition() {
        return this.partition;
    }
    public String receiveDisable() {
        return this.receiveDisable;
    }
    /**
     * @return Instructs the system to mark the target resource down when the test is successful.
     * 
     */
    public String reverse() {
        return this.reverse;
    }
    public String security() {
        return this.security;
    }
    public Integer timeUntilUp() {
        return this.timeUntilUp;
    }
    public Integer timeout() {
        return this.timeout;
    }
    /**
     * @return Displays whether the monitor operates in transparent mode.
     * 
     */
    public String transparent() {
        return this.transparent;
    }
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMonitorResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adaptive;
        private Integer adaptiveLimit;
        private String base;
        private String chaseReferrals;
        private String database;
        private String defaultsFrom;
        private String destination;
        private String filename;
        private String filter;
        private String id;
        private Integer interval;
        private Integer ipDscp;
        private String mandatoryAttributes;
        private String manualResume;
        private String mode;
        private String name;
        private String partition;
        private String receiveDisable;
        private String reverse;
        private String security;
        private Integer timeUntilUp;
        private Integer timeout;
        private String transparent;
        private String username;
        public Builder() {}
        public Builder(GetMonitorResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adaptive = defaults.adaptive;
    	      this.adaptiveLimit = defaults.adaptiveLimit;
    	      this.base = defaults.base;
    	      this.chaseReferrals = defaults.chaseReferrals;
    	      this.database = defaults.database;
    	      this.defaultsFrom = defaults.defaultsFrom;
    	      this.destination = defaults.destination;
    	      this.filename = defaults.filename;
    	      this.filter = defaults.filter;
    	      this.id = defaults.id;
    	      this.interval = defaults.interval;
    	      this.ipDscp = defaults.ipDscp;
    	      this.mandatoryAttributes = defaults.mandatoryAttributes;
    	      this.manualResume = defaults.manualResume;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.partition = defaults.partition;
    	      this.receiveDisable = defaults.receiveDisable;
    	      this.reverse = defaults.reverse;
    	      this.security = defaults.security;
    	      this.timeUntilUp = defaults.timeUntilUp;
    	      this.timeout = defaults.timeout;
    	      this.transparent = defaults.transparent;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder adaptive(String adaptive) {
            if (adaptive == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "adaptive");
            }
            this.adaptive = adaptive;
            return this;
        }
        @CustomType.Setter
        public Builder adaptiveLimit(Integer adaptiveLimit) {
            if (adaptiveLimit == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "adaptiveLimit");
            }
            this.adaptiveLimit = adaptiveLimit;
            return this;
        }
        @CustomType.Setter
        public Builder base(String base) {
            if (base == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "base");
            }
            this.base = base;
            return this;
        }
        @CustomType.Setter
        public Builder chaseReferrals(String chaseReferrals) {
            if (chaseReferrals == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "chaseReferrals");
            }
            this.chaseReferrals = chaseReferrals;
            return this;
        }
        @CustomType.Setter
        public Builder database(String database) {
            if (database == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "database");
            }
            this.database = database;
            return this;
        }
        @CustomType.Setter
        public Builder defaultsFrom(String defaultsFrom) {
            if (defaultsFrom == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "defaultsFrom");
            }
            this.defaultsFrom = defaultsFrom;
            return this;
        }
        @CustomType.Setter
        public Builder destination(String destination) {
            if (destination == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "destination");
            }
            this.destination = destination;
            return this;
        }
        @CustomType.Setter
        public Builder filename(String filename) {
            if (filename == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "filename");
            }
            this.filename = filename;
            return this;
        }
        @CustomType.Setter
        public Builder filter(String filter) {
            if (filter == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "filter");
            }
            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder interval(Integer interval) {
            if (interval == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "interval");
            }
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder ipDscp(Integer ipDscp) {
            if (ipDscp == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "ipDscp");
            }
            this.ipDscp = ipDscp;
            return this;
        }
        @CustomType.Setter
        public Builder mandatoryAttributes(String mandatoryAttributes) {
            if (mandatoryAttributes == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "mandatoryAttributes");
            }
            this.mandatoryAttributes = mandatoryAttributes;
            return this;
        }
        @CustomType.Setter
        public Builder manualResume(String manualResume) {
            if (manualResume == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "manualResume");
            }
            this.manualResume = manualResume;
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            if (mode == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "mode");
            }
            this.mode = mode;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder partition(String partition) {
            if (partition == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "partition");
            }
            this.partition = partition;
            return this;
        }
        @CustomType.Setter
        public Builder receiveDisable(String receiveDisable) {
            if (receiveDisable == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "receiveDisable");
            }
            this.receiveDisable = receiveDisable;
            return this;
        }
        @CustomType.Setter
        public Builder reverse(String reverse) {
            if (reverse == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "reverse");
            }
            this.reverse = reverse;
            return this;
        }
        @CustomType.Setter
        public Builder security(String security) {
            if (security == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "security");
            }
            this.security = security;
            return this;
        }
        @CustomType.Setter
        public Builder timeUntilUp(Integer timeUntilUp) {
            if (timeUntilUp == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "timeUntilUp");
            }
            this.timeUntilUp = timeUntilUp;
            return this;
        }
        @CustomType.Setter
        public Builder timeout(Integer timeout) {
            if (timeout == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "timeout");
            }
            this.timeout = timeout;
            return this;
        }
        @CustomType.Setter
        public Builder transparent(String transparent) {
            if (transparent == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "transparent");
            }
            this.transparent = transparent;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetMonitorResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetMonitorResult build() {
            final var _resultValue = new GetMonitorResult();
            _resultValue.adaptive = adaptive;
            _resultValue.adaptiveLimit = adaptiveLimit;
            _resultValue.base = base;
            _resultValue.chaseReferrals = chaseReferrals;
            _resultValue.database = database;
            _resultValue.defaultsFrom = defaultsFrom;
            _resultValue.destination = destination;
            _resultValue.filename = filename;
            _resultValue.filter = filter;
            _resultValue.id = id;
            _resultValue.interval = interval;
            _resultValue.ipDscp = ipDscp;
            _resultValue.mandatoryAttributes = mandatoryAttributes;
            _resultValue.manualResume = manualResume;
            _resultValue.mode = mode;
            _resultValue.name = name;
            _resultValue.partition = partition;
            _resultValue.receiveDisable = receiveDisable;
            _resultValue.reverse = reverse;
            _resultValue.security = security;
            _resultValue.timeUntilUp = timeUntilUp;
            _resultValue.timeout = timeout;
            _resultValue.transparent = transparent;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
