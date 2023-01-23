// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.inputs.FastHttpsAppMonitorArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppPoolMemberArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppTlsClientProfileArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppTlsServerProfileArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppVirtualServerArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppWafSecurityPolicyArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FastHttpsAppArgs extends com.pulumi.resources.ResourceArgs {

    public static final FastHttpsAppArgs Empty = new FastHttpsAppArgs();

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
     * List of LTM Policies to be applied FAST HTTPS Application.
     * 
     */
    @Import(name="endpointLtmPolicies")
    private @Nullable Output<List<String>> endpointLtmPolicies;

    /**
     * @return List of LTM Policies to be applied FAST HTTPS Application.
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
     * Name of an existing BIG-IP pool.
     * 
     */
    @Import(name="existingPool")
    private @Nullable Output<String> existingPool;

    /**
     * @return Name of an existing BIG-IP pool.
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
     * Name of an existing TLS client profile.
     * 
     */
    @Import(name="existingTlsClientProfile")
    private @Nullable Output<String> existingTlsClientProfile;

    /**
     * @return Name of an existing TLS client profile.
     * 
     */
    public Optional<Output<String>> existingTlsClientProfile() {
        return Optional.ofNullable(this.existingTlsClientProfile);
    }

    /**
     * Name of an existing TLS server profile.
     * 
     */
    @Import(name="existingTlsServerProfile")
    private @Nullable Output<String> existingTlsServerProfile;

    /**
     * @return Name of an existing TLS server profile.
     * 
     */
    public Optional<Output<String>> existingTlsServerProfile() {
        return Optional.ofNullable(this.existingTlsServerProfile);
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
    private @Nullable Output<FastHttpsAppMonitorArgs> monitor;

    /**
     * @return `monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppMonitorArgs>> monitor() {
        return Optional.ofNullable(this.monitor);
    }

    /**
     * `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    @Import(name="poolMembers")
    private @Nullable Output<List<FastHttpsAppPoolMemberArgs>> poolMembers;

    /**
     * @return `pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    public Optional<Output<List<FastHttpsAppPoolMemberArgs>>> poolMembers() {
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
     * `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     * 
     */
    @Import(name="tlsClientProfile")
    private @Nullable Output<FastHttpsAppTlsClientProfileArgs> tlsClientProfile;

    /**
     * @return `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
     * See TLS Client Profile below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppTlsClientProfileArgs>> tlsClientProfile() {
        return Optional.ofNullable(this.tlsClientProfile);
    }

    /**
     * `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     */
    @Import(name="tlsServerProfile")
    private @Nullable Output<FastHttpsAppTlsServerProfileArgs> tlsServerProfile;

    /**
     * @return `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppTlsServerProfileArgs>> tlsServerProfile() {
        return Optional.ofNullable(this.tlsServerProfile);
    }

    /**
     * `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    @Import(name="virtualServer")
    private @Nullable Output<FastHttpsAppVirtualServerArgs> virtualServer;

    /**
     * @return `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppVirtualServerArgs>> virtualServer() {
        return Optional.ofNullable(this.virtualServer);
    }

    /**
     * `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    @Import(name="wafSecurityPolicy")
    private @Nullable Output<FastHttpsAppWafSecurityPolicyArgs> wafSecurityPolicy;

    /**
     * @return `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
     * See WAF Security Policy below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppWafSecurityPolicyArgs>> wafSecurityPolicy() {
        return Optional.ofNullable(this.wafSecurityPolicy);
    }

    private FastHttpsAppArgs() {}

    private FastHttpsAppArgs(FastHttpsAppArgs $) {
        this.application = $.application;
        this.endpointLtmPolicies = $.endpointLtmPolicies;
        this.existingMonitor = $.existingMonitor;
        this.existingPool = $.existingPool;
        this.existingSnatPool = $.existingSnatPool;
        this.existingTlsClientProfile = $.existingTlsClientProfile;
        this.existingTlsServerProfile = $.existingTlsServerProfile;
        this.existingWafSecurityPolicy = $.existingWafSecurityPolicy;
        this.loadBalancingMode = $.loadBalancingMode;
        this.monitor = $.monitor;
        this.poolMembers = $.poolMembers;
        this.securityLogProfiles = $.securityLogProfiles;
        this.slowRampTime = $.slowRampTime;
        this.snatPoolAddresses = $.snatPoolAddresses;
        this.tenant = $.tenant;
        this.tlsClientProfile = $.tlsClientProfile;
        this.tlsServerProfile = $.tlsServerProfile;
        this.virtualServer = $.virtualServer;
        this.wafSecurityPolicy = $.wafSecurityPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FastHttpsAppArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FastHttpsAppArgs $;

        public Builder() {
            $ = new FastHttpsAppArgs();
        }

        public Builder(FastHttpsAppArgs defaults) {
            $ = new FastHttpsAppArgs(Objects.requireNonNull(defaults));
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
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTPS Application.
         * 
         * @return builder
         * 
         */
        public Builder endpointLtmPolicies(@Nullable Output<List<String>> endpointLtmPolicies) {
            $.endpointLtmPolicies = endpointLtmPolicies;
            return this;
        }

        /**
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTPS Application.
         * 
         * @return builder
         * 
         */
        public Builder endpointLtmPolicies(List<String> endpointLtmPolicies) {
            return endpointLtmPolicies(Output.of(endpointLtmPolicies));
        }

        /**
         * @param endpointLtmPolicies List of LTM Policies to be applied FAST HTTPS Application.
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
         * @param existingPool Name of an existing BIG-IP pool.
         * 
         * @return builder
         * 
         */
        public Builder existingPool(@Nullable Output<String> existingPool) {
            $.existingPool = existingPool;
            return this;
        }

        /**
         * @param existingPool Name of an existing BIG-IP pool.
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
         * @param existingTlsClientProfile Name of an existing TLS client profile.
         * 
         * @return builder
         * 
         */
        public Builder existingTlsClientProfile(@Nullable Output<String> existingTlsClientProfile) {
            $.existingTlsClientProfile = existingTlsClientProfile;
            return this;
        }

        /**
         * @param existingTlsClientProfile Name of an existing TLS client profile.
         * 
         * @return builder
         * 
         */
        public Builder existingTlsClientProfile(String existingTlsClientProfile) {
            return existingTlsClientProfile(Output.of(existingTlsClientProfile));
        }

        /**
         * @param existingTlsServerProfile Name of an existing TLS server profile.
         * 
         * @return builder
         * 
         */
        public Builder existingTlsServerProfile(@Nullable Output<String> existingTlsServerProfile) {
            $.existingTlsServerProfile = existingTlsServerProfile;
            return this;
        }

        /**
         * @param existingTlsServerProfile Name of an existing TLS server profile.
         * 
         * @return builder
         * 
         */
        public Builder existingTlsServerProfile(String existingTlsServerProfile) {
            return existingTlsServerProfile(Output.of(existingTlsServerProfile));
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
        public Builder monitor(@Nullable Output<FastHttpsAppMonitorArgs> monitor) {
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
        public Builder monitor(FastHttpsAppMonitorArgs monitor) {
            return monitor(Output.of(monitor));
        }

        /**
         * @param poolMembers `pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder poolMembers(@Nullable Output<List<FastHttpsAppPoolMemberArgs>> poolMembers) {
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
        public Builder poolMembers(List<FastHttpsAppPoolMemberArgs> poolMembers) {
            return poolMembers(Output.of(poolMembers));
        }

        /**
         * @param poolMembers `pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder poolMembers(FastHttpsAppPoolMemberArgs... poolMembers) {
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
         * @param tlsClientProfile `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
         * See TLS Client Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientProfile(@Nullable Output<FastHttpsAppTlsClientProfileArgs> tlsClientProfile) {
            $.tlsClientProfile = tlsClientProfile;
            return this;
        }

        /**
         * @param tlsClientProfile `tls_client_profile` block takes input for FAST-Generated TLS client Profile.
         * See TLS Client Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientProfile(FastHttpsAppTlsClientProfileArgs tlsClientProfile) {
            return tlsClientProfile(Output.of(tlsClientProfile));
        }

        /**
         * @param tlsServerProfile `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
         * See TLS Server Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerProfile(@Nullable Output<FastHttpsAppTlsServerProfileArgs> tlsServerProfile) {
            $.tlsServerProfile = tlsServerProfile;
            return this;
        }

        /**
         * @param tlsServerProfile `tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
         * See TLS Server Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerProfile(FastHttpsAppTlsServerProfileArgs tlsServerProfile) {
            return tlsServerProfile(Output.of(tlsServerProfile));
        }

        /**
         * @param virtualServer `virtual_server` block will provide `ip` and `port` options to be used for virtual server.
         * See virtual server below for more details.
         * 
         * @return builder
         * 
         */
        public Builder virtualServer(@Nullable Output<FastHttpsAppVirtualServerArgs> virtualServer) {
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
        public Builder virtualServer(FastHttpsAppVirtualServerArgs virtualServer) {
            return virtualServer(Output.of(virtualServer));
        }

        /**
         * @param wafSecurityPolicy `waf_security_policy` block takes input for FAST-Generated WAF Security Policy.
         * See WAF Security Policy below for more details.
         * 
         * @return builder
         * 
         */
        public Builder wafSecurityPolicy(@Nullable Output<FastHttpsAppWafSecurityPolicyArgs> wafSecurityPolicy) {
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
        public Builder wafSecurityPolicy(FastHttpsAppWafSecurityPolicyArgs wafSecurityPolicy) {
            return wafSecurityPolicy(Output.of(wafSecurityPolicy));
        }

        public FastHttpsAppArgs build() {
            $.application = Objects.requireNonNull($.application, "expected parameter 'application' to be non-null");
            $.tenant = Objects.requireNonNull($.tenant, "expected parameter 'tenant' to be non-null");
            return $;
        }
    }

}