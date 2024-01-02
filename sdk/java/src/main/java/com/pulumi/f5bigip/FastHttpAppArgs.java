// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.inputs.FastHttpAppMonitorArgs;
import com.pulumi.f5bigip.inputs.FastHttpAppPoolMemberArgs;
import com.pulumi.f5bigip.inputs.FastHttpAppVirtualServerArgs;
import com.pulumi.f5bigip.inputs.FastHttpAppWafSecurityPolicyArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FastHttpAppArgs extends com.pulumi.resources.ResourceArgs {

    public static final FastHttpAppArgs Empty = new FastHttpAppArgs();

    /**
     * Name of the FAST HTTPS application.
     * 
     */
    @Import(name="application", required=true)
    private Output<String> application;

    /**
     * @return Name of the FAST HTTPS application.
     * 
     */
    public Output<String> application() {
        return this.application;
    }

    /**
     * List of LTM Policies to be applied FAST HTTP Application.
     * 
     */
    @Import(name="endpointLtmPolicies")
    private @Nullable Output<List<String>> endpointLtmPolicies;

    /**
     * @return List of LTM Policies to be applied FAST HTTP Application.
     * 
     */
    public Optional<Output<List<String>>> endpointLtmPolicies() {
        return Optional.ofNullable(this.endpointLtmPolicies);
    }

    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     * 
     */
    @Import(name="existingMonitor")
    private @Nullable Output<String> existingMonitor;

    /**
     * @return Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     * 
     */
    public Optional<Output<String>> existingMonitor() {
        return Optional.ofNullable(this.existingMonitor);
    }

    /**
     * Select an existing BIG-IP Pool
     * 
     */
    @Import(name="existingPool")
    private @Nullable Output<String> existingPool;

    /**
     * @return Select an existing BIG-IP Pool
     * 
     */
    public Optional<Output<String>> existingPool() {
        return Optional.ofNullable(this.existingPool);
    }

    /**
     * Name of an existing BIG-IP SNAT pool.
     * 
     */
    @Import(name="existingSnatPool")
    private @Nullable Output<String> existingSnatPool;

    /**
     * @return Name of an existing BIG-IP SNAT pool.
     * 
     */
    public Optional<Output<String>> existingSnatPool() {
        return Optional.ofNullable(this.existingSnatPool);
    }

    /**
     * Name of an existing WAF Security policy.
     * 
     */
    @Import(name="existingWafSecurityPolicy")
    private @Nullable Output<String> existingWafSecurityPolicy;

    /**
     * @return Name of an existing WAF Security policy.
     * 
     */
    public Optional<Output<String>> existingWafSecurityPolicy() {
        return Optional.ofNullable(this.existingWafSecurityPolicy);
    }

    /**
     * Type of fallback persistence record to be created for each new client connection.
     * 
     */
    @Import(name="fallbackPersistence")
    private @Nullable Output<String> fallbackPersistence;

    /**
     * @return Type of fallback persistence record to be created for each new client connection.
     * 
     */
    public Optional<Output<String>> fallbackPersistence() {
        return Optional.ofNullable(this.fallbackPersistence);
    }

    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     * 
     */
    @Import(name="loadBalancingMode")
    private @Nullable Output<String> loadBalancingMode;

    /**
     * @return A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     * 
     */
    public Optional<Output<String>> loadBalancingMode() {
        return Optional.ofNullable(this.loadBalancingMode);
    }

    /**
     * `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    @Import(name="monitor")
    private @Nullable Output<FastHttpAppMonitorArgs> monitor;

    /**
     * @return `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    public Optional<Output<FastHttpAppMonitorArgs>> monitor() {
        return Optional.ofNullable(this.monitor);
    }

    /**
     * Name of an existing BIG-IP persistence profile to be used.
     * 
     */
    @Import(name="persistenceProfile")
    private @Nullable Output<String> persistenceProfile;

    /**
     * @return Name of an existing BIG-IP persistence profile to be used.
     * 
     */
    public Optional<Output<String>> persistenceProfile() {
        return Optional.ofNullable(this.persistenceProfile);
    }

    /**
     * Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
     * 
     */
    @Import(name="persistenceType")
    private @Nullable Output<String> persistenceType;

    /**
     * @return Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
     * 
     */
    public Optional<Output<String>> persistenceType() {
        return Optional.ofNullable(this.persistenceType);
    }

    /**
     * `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    @Import(name="poolMembers")
    private @Nullable Output<List<FastHttpAppPoolMemberArgs>> poolMembers;

    /**
     * @return `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    public Optional<Output<List<FastHttpAppPoolMemberArgs>>> poolMembers() {
        return Optional.ofNullable(this.poolMembers);
    }

    /**
     * List of security log profiles to be used for FAST application
     * 
     */
    @Import(name="securityLogProfiles")
    private @Nullable Output<List<String>> securityLogProfiles;

    /**
     * @return List of security log profiles to be used for FAST application
     * 
     */
    public Optional<Output<List<String>>> securityLogProfiles() {
        return Optional.ofNullable(this.securityLogProfiles);
    }

    /**
     * List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
     * 
     */
    @Import(name="serviceDiscoveries")
    private @Nullable Output<List<String>> serviceDiscoveries;

    /**
     * @return List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
     * 
     */
    public Optional<Output<List<String>>> serviceDiscoveries() {
        return Optional.ofNullable(this.serviceDiscoveries);
    }

    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     * 
     */
    @Import(name="slowRampTime")
    private @Nullable Output<Integer> slowRampTime;

    /**
     * @return Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     * 
     */
    public Optional<Output<Integer>> slowRampTime() {
        return Optional.ofNullable(this.slowRampTime);
    }

    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    @Import(name="snatPoolAddresses")
    private @Nullable Output<List<String>> snatPoolAddresses;

    /**
     * @return List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    public Optional<Output<List<String>>> snatPoolAddresses() {
        return Optional.ofNullable(this.snatPoolAddresses);
    }

    /**
     * Name of the FAST HTTPS application tenant.
     * 
     */
    @Import(name="tenant", required=true)
    private Output<String> tenant;

    /**
     * @return Name of the FAST HTTPS application tenant.
     * 
     */
    public Output<String> tenant() {
        return this.tenant;
    }

    /**
     * `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    @Import(name="virtualServer")
    private @Nullable Output<FastHttpAppVirtualServerArgs> virtualServer;

    /**
     * @return `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    public Optional<Output<FastHttpAppVirtualServerArgs>> virtualServer() {
        return Optional.ofNullable(this.virtualServer);
    }

    /**
     * `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    @Import(name="wafSecurityPolicy")
    private @Nullable Output<FastHttpAppWafSecurityPolicyArgs> wafSecurityPolicy;

    /**
     * @return `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    public Optional<Output<FastHttpAppWafSecurityPolicyArgs>> wafSecurityPolicy() {
        return Optional.ofNullable(this.wafSecurityPolicy);
    }

    private FastHttpAppArgs() {}

    private FastHttpAppArgs(FastHttpAppArgs $) {
        this.application = $.application;
        this.endpointLtmPolicies = $.endpointLtmPolicies;
        this.existingMonitor = $.existingMonitor;
        this.existingPool = $.existingPool;
        this.existingSnatPool = $.existingSnatPool;
        this.existingWafSecurityPolicy = $.existingWafSecurityPolicy;
        this.fallbackPersistence = $.fallbackPersistence;
        this.loadBalancingMode = $.loadBalancingMode;
        this.monitor = $.monitor;
        this.persistenceProfile = $.persistenceProfile;
        this.persistenceType = $.persistenceType;
        this.poolMembers = $.poolMembers;
        this.securityLogProfiles = $.securityLogProfiles;
        this.serviceDiscoveries = $.serviceDiscoveries;
        this.slowRampTime = $.slowRampTime;
        this.snatPoolAddresses = $.snatPoolAddresses;
        this.tenant = $.tenant;
        this.virtualServer = $.virtualServer;
        this.wafSecurityPolicy = $.wafSecurityPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FastHttpAppArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FastHttpAppArgs $;

        public Builder() {
            $ = new FastHttpAppArgs();
        }

        public Builder(FastHttpAppArgs defaults) {
            $ = new FastHttpAppArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param application Name of the FAST HTTPS application.
         * 
         * @return builder
         * 
         */
        public Builder application(Output<String> application) {
            $.application = application;
            return this;
        }

        /**
         * @param application Name of the FAST HTTPS application.
         * 
         * @return builder
         * 
         */
        public Builder application(String application) {
            return application(Output.of(application));
        }

        /**
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTP Application.
         * 
         * @return builder
         * 
         */
        public Builder endpointLtmPolicies(@Nullable Output<List<String>> endpointLtmPolicies) {
            $.endpointLtmPolicies = endpointLtmPolicies;
            return this;
        }

        /**
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTP Application.
         * 
         * @return builder
         * 
         */
        public Builder endpointLtmPolicies(List<String> endpointLtmPolicies) {
            return endpointLtmPolicies(Output.of(endpointLtmPolicies));
        }

        /**
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTP Application.
         * 
         * @return builder
         * 
         */
        public Builder endpointLtmPolicies(String... endpointLtmPolicies) {
            return endpointLtmPolicies(List.of(endpointLtmPolicies));
        }

        /**
         * @param existingMonitor Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
         * 
         * @return builder
         * 
         */
        public Builder existingMonitor(@Nullable Output<String> existingMonitor) {
            $.existingMonitor = existingMonitor;
            return this;
        }

        /**
         * @param existingMonitor Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
         * 
         * @return builder
         * 
         */
        public Builder existingMonitor(String existingMonitor) {
            return existingMonitor(Output.of(existingMonitor));
        }

        /**
         * @param existingPool Select an existing BIG-IP Pool
         * 
         * @return builder
         * 
         */
        public Builder existingPool(@Nullable Output<String> existingPool) {
            $.existingPool = existingPool;
            return this;
        }

        /**
         * @param existingPool Select an existing BIG-IP Pool
         * 
         * @return builder
         * 
         */
        public Builder existingPool(String existingPool) {
            return existingPool(Output.of(existingPool));
        }

        /**
         * @param existingSnatPool Name of an existing BIG-IP SNAT pool.
         * 
         * @return builder
         * 
         */
        public Builder existingSnatPool(@Nullable Output<String> existingSnatPool) {
            $.existingSnatPool = existingSnatPool;
            return this;
        }

        /**
         * @param existingSnatPool Name of an existing BIG-IP SNAT pool.
         * 
         * @return builder
         * 
         */
        public Builder existingSnatPool(String existingSnatPool) {
            return existingSnatPool(Output.of(existingSnatPool));
        }

        /**
         * @param existingWafSecurityPolicy Name of an existing WAF Security policy.
         * 
         * @return builder
         * 
         */
        public Builder existingWafSecurityPolicy(@Nullable Output<String> existingWafSecurityPolicy) {
            $.existingWafSecurityPolicy = existingWafSecurityPolicy;
            return this;
        }

        /**
         * @param existingWafSecurityPolicy Name of an existing WAF Security policy.
         * 
         * @return builder
         * 
         */
        public Builder existingWafSecurityPolicy(String existingWafSecurityPolicy) {
            return existingWafSecurityPolicy(Output.of(existingWafSecurityPolicy));
        }

        /**
         * @param fallbackPersistence Type of fallback persistence record to be created for each new client connection.
         * 
         * @return builder
         * 
         */
        public Builder fallbackPersistence(@Nullable Output<String> fallbackPersistence) {
            $.fallbackPersistence = fallbackPersistence;
            return this;
        }

        /**
         * @param fallbackPersistence Type of fallback persistence record to be created for each new client connection.
         * 
         * @return builder
         * 
         */
        public Builder fallbackPersistence(String fallbackPersistence) {
            return fallbackPersistence(Output.of(fallbackPersistence));
        }

        /**
         * @param loadBalancingMode A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingMode(@Nullable Output<String> loadBalancingMode) {
            $.loadBalancingMode = loadBalancingMode;
            return this;
        }

        /**
         * @param loadBalancingMode A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingMode(String loadBalancingMode) {
            return loadBalancingMode(Output.of(loadBalancingMode));
        }

        /**
         * @param monitor `monitor` block takes input for FAST-Generated Pool Monitor.
         * See Pool Monitor below for more details.
         * 
         * @return builder
         * 
         */
        public Builder monitor(@Nullable Output<FastHttpAppMonitorArgs> monitor) {
            $.monitor = monitor;
            return this;
        }

        /**
         * @param monitor `monitor` block takes input for FAST-Generated Pool Monitor.
         * See Pool Monitor below for more details.
         * 
         * @return builder
         * 
         */
        public Builder monitor(FastHttpAppMonitorArgs monitor) {
            return monitor(Output.of(monitor));
        }

        /**
         * @param persistenceProfile Name of an existing BIG-IP persistence profile to be used.
         * 
         * @return builder
         * 
         */
        public Builder persistenceProfile(@Nullable Output<String> persistenceProfile) {
            $.persistenceProfile = persistenceProfile;
            return this;
        }

        /**
         * @param persistenceProfile Name of an existing BIG-IP persistence profile to be used.
         * 
         * @return builder
         * 
         */
        public Builder persistenceProfile(String persistenceProfile) {
            return persistenceProfile(Output.of(persistenceProfile));
        }

        /**
         * @param persistenceType Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
         * 
         * @return builder
         * 
         */
        public Builder persistenceType(@Nullable Output<String> persistenceType) {
            $.persistenceType = persistenceType;
            return this;
        }

        /**
         * @param persistenceType Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
         * 
         * @return builder
         * 
         */
        public Builder persistenceType(String persistenceType) {
            return persistenceType(Output.of(persistenceType));
        }

        /**
         * @param poolMembers `pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder poolMembers(@Nullable Output<List<FastHttpAppPoolMemberArgs>> poolMembers) {
            $.poolMembers = poolMembers;
            return this;
        }

        /**
         * @param poolMembers `pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder poolMembers(List<FastHttpAppPoolMemberArgs> poolMembers) {
            return poolMembers(Output.of(poolMembers));
        }

        /**
         * @param poolMembers `pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder poolMembers(FastHttpAppPoolMemberArgs... poolMembers) {
            return poolMembers(List.of(poolMembers));
        }

        /**
         * @param securityLogProfiles List of security log profiles to be used for FAST application
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(@Nullable Output<List<String>> securityLogProfiles) {
            $.securityLogProfiles = securityLogProfiles;
            return this;
        }

        /**
         * @param securityLogProfiles List of security log profiles to be used for FAST application
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(List<String> securityLogProfiles) {
            return securityLogProfiles(Output.of(securityLogProfiles));
        }

        /**
         * @param securityLogProfiles List of security log profiles to be used for FAST application
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(String... securityLogProfiles) {
            return securityLogProfiles(List.of(securityLogProfiles));
        }

        /**
         * @param serviceDiscoveries List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
         * 
         * @return builder
         * 
         */
        public Builder serviceDiscoveries(@Nullable Output<List<String>> serviceDiscoveries) {
            $.serviceDiscoveries = serviceDiscoveries;
            return this;
        }

        /**
         * @param serviceDiscoveries List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
         * 
         * @return builder
         * 
         */
        public Builder serviceDiscoveries(List<String> serviceDiscoveries) {
            return serviceDiscoveries(Output.of(serviceDiscoveries));
        }

        /**
         * @param serviceDiscoveries List of different cloud service discovery config provided as string, provided `service_discovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
         * 
         * @return builder
         * 
         */
        public Builder serviceDiscoveries(String... serviceDiscoveries) {
            return serviceDiscoveries(List.of(serviceDiscoveries));
        }

        /**
         * @param slowRampTime Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
         * 
         * @return builder
         * 
         */
        public Builder slowRampTime(@Nullable Output<Integer> slowRampTime) {
            $.slowRampTime = slowRampTime;
            return this;
        }

        /**
         * @param slowRampTime Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
         * 
         * @return builder
         * 
         */
        public Builder slowRampTime(Integer slowRampTime) {
            return slowRampTime(Output.of(slowRampTime));
        }

        /**
         * @param snatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder snatPoolAddresses(@Nullable Output<List<String>> snatPoolAddresses) {
            $.snatPoolAddresses = snatPoolAddresses;
            return this;
        }

        /**
         * @param snatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder snatPoolAddresses(List<String> snatPoolAddresses) {
            return snatPoolAddresses(Output.of(snatPoolAddresses));
        }

        /**
         * @param snatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder snatPoolAddresses(String... snatPoolAddresses) {
            return snatPoolAddresses(List.of(snatPoolAddresses));
        }

        /**
         * @param tenant Name of the FAST HTTPS application tenant.
         * 
         * @return builder
         * 
         */
        public Builder tenant(Output<String> tenant) {
            $.tenant = tenant;
            return this;
        }

        /**
         * @param tenant Name of the FAST HTTPS application tenant.
         * 
         * @return builder
         * 
         */
        public Builder tenant(String tenant) {
            return tenant(Output.of(tenant));
        }

        /**
         * @param virtualServer `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
         * See virtual server below for more details.
         * 
         * @return builder
         * 
         */
        public Builder virtualServer(@Nullable Output<FastHttpAppVirtualServerArgs> virtualServer) {
            $.virtualServer = virtualServer;
            return this;
        }

        /**
         * @param virtualServer `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
         * See virtual server below for more details.
         * 
         * @return builder
         * 
         */
        public Builder virtualServer(FastHttpAppVirtualServerArgs virtualServer) {
            return virtualServer(Output.of(virtualServer));
        }

        /**
         * @param wafSecurityPolicy `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
         * See WAF Security Policy below for more details.
         * 
         * @return builder
         * 
         */
        public Builder wafSecurityPolicy(@Nullable Output<FastHttpAppWafSecurityPolicyArgs> wafSecurityPolicy) {
            $.wafSecurityPolicy = wafSecurityPolicy;
            return this;
        }

        /**
         * @param wafSecurityPolicy `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
         * See WAF Security Policy below for more details.
         * 
         * @return builder
         * 
         */
        public Builder wafSecurityPolicy(FastHttpAppWafSecurityPolicyArgs wafSecurityPolicy) {
            return wafSecurityPolicy(Output.of(wafSecurityPolicy));
        }

        public FastHttpAppArgs build() {
            if ($.application == null) {
                throw new MissingRequiredPropertyException("FastHttpAppArgs", "application");
            }
            if ($.tenant == null) {
                throw new MissingRequiredPropertyException("FastHttpAppArgs", "tenant");
            }
            return $;
        }
    }

}
