// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FastHttpAppMonitor {
    /**
     * @return Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     * 
     */
    private @Nullable Integer interval;
    /**
     * @return set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
     * 
     */
    private @Nullable Boolean monitorAuth;
    /**
     * @return password for web access on FAST-Generated Pool Monitor.
     * 
     */
    private @Nullable String password;
    /**
     * @return The presence of this string anywhere in the HTTP response implies availability.
     * 
     */
    private @Nullable String response;
    /**
     * @return Specify data to be sent during each health check for FAST-Generated Pool Monitor.
     * 
     */
    private @Nullable String sendString;
    /**
     * @return username for web access on FAST-Generated Pool Monitor.
     * 
     */
    private @Nullable String username;

    private FastHttpAppMonitor() {}
    /**
     * @return Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     * 
     */
    public Optional<Integer> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
     * 
     */
    public Optional<Boolean> monitorAuth() {
        return Optional.ofNullable(this.monitorAuth);
    }
    /**
     * @return password for web access on FAST-Generated Pool Monitor.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The presence of this string anywhere in the HTTP response implies availability.
     * 
     */
    public Optional<String> response() {
        return Optional.ofNullable(this.response);
    }
    /**
     * @return Specify data to be sent during each health check for FAST-Generated Pool Monitor.
     * 
     */
    public Optional<String> sendString() {
        return Optional.ofNullable(this.sendString);
    }
    /**
     * @return username for web access on FAST-Generated Pool Monitor.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FastHttpAppMonitor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer interval;
        private @Nullable Boolean monitorAuth;
        private @Nullable String password;
        private @Nullable String response;
        private @Nullable String sendString;
        private @Nullable String username;
        public Builder() {}
        public Builder(FastHttpAppMonitor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.interval = defaults.interval;
    	      this.monitorAuth = defaults.monitorAuth;
    	      this.password = defaults.password;
    	      this.response = defaults.response;
    	      this.sendString = defaults.sendString;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder interval(@Nullable Integer interval) {
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder monitorAuth(@Nullable Boolean monitorAuth) {
            this.monitorAuth = monitorAuth;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder response(@Nullable String response) {
            this.response = response;
            return this;
        }
        @CustomType.Setter
        public Builder sendString(@Nullable String sendString) {
            this.sendString = sendString;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public FastHttpAppMonitor build() {
            final var _resultValue = new FastHttpAppMonitor();
            _resultValue.interval = interval;
            _resultValue.monitorAuth = monitorAuth;
            _resultValue.password = password;
            _resultValue.response = response;
            _resultValue.sendString = sendString;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
