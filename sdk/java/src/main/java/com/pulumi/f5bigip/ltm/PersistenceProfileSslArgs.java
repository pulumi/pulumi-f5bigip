// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PersistenceProfileSslArgs extends com.pulumi.resources.ResourceArgs {

    public static final PersistenceProfileSslArgs Empty = new PersistenceProfileSslArgs();

    @Import(name="appService")
    private @Nullable Output<String> appService;

    public Optional<Output<String>> appService() {
        return Optional.ofNullable(this.appService);
    }

    /**
     * Inherit defaults from parent profile
     * 
     */
    @Import(name="defaultsFrom", required=true)
    private Output<String> defaultsFrom;

    /**
     * @return Inherit defaults from parent profile
     * 
     */
    public Output<String> defaultsFrom() {
        return this.defaultsFrom;
    }

    /**
     * To enable _ disable match across pools with given persistence record
     * 
     */
    @Import(name="matchAcrossPools")
    private @Nullable Output<String> matchAcrossPools;

    /**
     * @return To enable _ disable match across pools with given persistence record
     * 
     */
    public Optional<Output<String>> matchAcrossPools() {
        return Optional.ofNullable(this.matchAcrossPools);
    }

    /**
     * To enable _ disable match across services with given persistence record
     * 
     */
    @Import(name="matchAcrossServices")
    private @Nullable Output<String> matchAcrossServices;

    /**
     * @return To enable _ disable match across services with given persistence record
     * 
     */
    public Optional<Output<String>> matchAcrossServices() {
        return Optional.ofNullable(this.matchAcrossServices);
    }

    /**
     * To enable _ disable match across services with given persistence record
     * 
     */
    @Import(name="matchAcrossVirtuals")
    private @Nullable Output<String> matchAcrossVirtuals;

    /**
     * @return To enable _ disable match across services with given persistence record
     * 
     */
    public Optional<Output<String>> matchAcrossVirtuals() {
        return Optional.ofNullable(this.matchAcrossVirtuals);
    }

    /**
     * To enable _ disable
     * 
     */
    @Import(name="mirror")
    private @Nullable Output<String> mirror;

    /**
     * @return To enable _ disable
     * 
     */
    public Optional<Output<String>> mirror() {
        return Optional.ofNullable(this.mirror);
    }

    /**
     * Name of the persistence profile
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the persistence profile
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     * 
     */
    @Import(name="overrideConnLimit")
    private @Nullable Output<String> overrideConnLimit;

    /**
     * @return To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     * 
     */
    public Optional<Output<String>> overrideConnLimit() {
        return Optional.ofNullable(this.overrideConnLimit);
    }

    /**
     * Timeout for persistence of the session
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Timeout for persistence of the session
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private PersistenceProfileSslArgs() {}

    private PersistenceProfileSslArgs(PersistenceProfileSslArgs $) {
        this.appService = $.appService;
        this.defaultsFrom = $.defaultsFrom;
        this.matchAcrossPools = $.matchAcrossPools;
        this.matchAcrossServices = $.matchAcrossServices;
        this.matchAcrossVirtuals = $.matchAcrossVirtuals;
        this.mirror = $.mirror;
        this.name = $.name;
        this.overrideConnLimit = $.overrideConnLimit;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PersistenceProfileSslArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PersistenceProfileSslArgs $;

        public Builder() {
            $ = new PersistenceProfileSslArgs();
        }

        public Builder(PersistenceProfileSslArgs defaults) {
            $ = new PersistenceProfileSslArgs(Objects.requireNonNull(defaults));
        }

        public Builder appService(@Nullable Output<String> appService) {
            $.appService = appService;
            return this;
        }

        public Builder appService(String appService) {
            return appService(Output.of(appService));
        }

        /**
         * @param defaultsFrom Inherit defaults from parent profile
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(Output<String> defaultsFrom) {
            $.defaultsFrom = defaultsFrom;
            return this;
        }

        /**
         * @param defaultsFrom Inherit defaults from parent profile
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(String defaultsFrom) {
            return defaultsFrom(Output.of(defaultsFrom));
        }

        /**
         * @param matchAcrossPools To enable _ disable match across pools with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossPools(@Nullable Output<String> matchAcrossPools) {
            $.matchAcrossPools = matchAcrossPools;
            return this;
        }

        /**
         * @param matchAcrossPools To enable _ disable match across pools with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossPools(String matchAcrossPools) {
            return matchAcrossPools(Output.of(matchAcrossPools));
        }

        /**
         * @param matchAcrossServices To enable _ disable match across services with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossServices(@Nullable Output<String> matchAcrossServices) {
            $.matchAcrossServices = matchAcrossServices;
            return this;
        }

        /**
         * @param matchAcrossServices To enable _ disable match across services with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossServices(String matchAcrossServices) {
            return matchAcrossServices(Output.of(matchAcrossServices));
        }

        /**
         * @param matchAcrossVirtuals To enable _ disable match across services with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossVirtuals(@Nullable Output<String> matchAcrossVirtuals) {
            $.matchAcrossVirtuals = matchAcrossVirtuals;
            return this;
        }

        /**
         * @param matchAcrossVirtuals To enable _ disable match across services with given persistence record
         * 
         * @return builder
         * 
         */
        public Builder matchAcrossVirtuals(String matchAcrossVirtuals) {
            return matchAcrossVirtuals(Output.of(matchAcrossVirtuals));
        }

        /**
         * @param mirror To enable _ disable
         * 
         * @return builder
         * 
         */
        public Builder mirror(@Nullable Output<String> mirror) {
            $.mirror = mirror;
            return this;
        }

        /**
         * @param mirror To enable _ disable
         * 
         * @return builder
         * 
         */
        public Builder mirror(String mirror) {
            return mirror(Output.of(mirror));
        }

        /**
         * @param name Name of the persistence profile
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the persistence profile
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param overrideConnLimit To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
         * limits remain hard limits and are not overridden.
         * 
         * @return builder
         * 
         */
        public Builder overrideConnLimit(@Nullable Output<String> overrideConnLimit) {
            $.overrideConnLimit = overrideConnLimit;
            return this;
        }

        /**
         * @param overrideConnLimit To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
         * limits remain hard limits and are not overridden.
         * 
         * @return builder
         * 
         */
        public Builder overrideConnLimit(String overrideConnLimit) {
            return overrideConnLimit(Output.of(overrideConnLimit));
        }

        /**
         * @param timeout Timeout for persistence of the session
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Timeout for persistence of the session
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public PersistenceProfileSslArgs build() {
            if ($.defaultsFrom == null) {
                throw new MissingRequiredPropertyException("PersistenceProfileSslArgs", "defaultsFrom");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("PersistenceProfileSslArgs", "name");
            }
            return $;
        }
    }

}
