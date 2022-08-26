// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.inputs.FastHttpsAppCreateTlsServerProfileArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppFastCreateMonitorArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppFastCreatePoolMemberArgs;
import com.pulumi.f5bigip.inputs.FastHttpsAppVirtualServerArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FastHttpsAppState extends com.pulumi.resources.ResourceArgs {

    public static final FastHttpsAppState Empty = new FastHttpsAppState();

    /**
     * Name of the FAST HTTPS application.
     * 
     */
    @Import(name="application")
    private @Nullable Output<String> application;

    /**
     * @return Name of the FAST HTTPS application.
     * 
     */
    public Optional<Output<String>> application() {
        return Optional.ofNullable(this.application);
    }

    /**
     * `create_tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     */
    @Import(name="createTlsServerProfile")
    private @Nullable Output<FastHttpsAppCreateTlsServerProfileArgs> createTlsServerProfile;

    /**
     * @return `create_tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
     * See TLS Server Profile below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppCreateTlsServerProfileArgs>> createTlsServerProfile() {
        return Optional.ofNullable(this.createTlsServerProfile);
    }

    /**
     * Name of an existing BIG-IP pool.
     * 
     */
    @Import(name="existPoolName")
    private @Nullable Output<String> existPoolName;

    /**
     * @return Name of an existing BIG-IP pool.
     * 
     */
    public Optional<Output<String>> existPoolName() {
        return Optional.ofNullable(this.existPoolName);
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
     * `fast_create_monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    @Import(name="fastCreateMonitor")
    private @Nullable Output<FastHttpsAppFastCreateMonitorArgs> fastCreateMonitor;

    /**
     * @return `fast_create_monitor` block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     * 
     */
    public Optional<Output<FastHttpsAppFastCreateMonitorArgs>> fastCreateMonitor() {
        return Optional.ofNullable(this.fastCreateMonitor);
    }

    /**
     * `fast_create_pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    @Import(name="fastCreatePoolMembers")
    private @Nullable Output<List<FastHttpsAppFastCreatePoolMemberArgs>> fastCreatePoolMembers;

    /**
     * @return `fast_create_pool_members` block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     * 
     */
    public Optional<Output<List<FastHttpsAppFastCreatePoolMemberArgs>>> fastCreatePoolMembers() {
        return Optional.ofNullable(this.fastCreatePoolMembers);
    }

    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    @Import(name="fastCreateSnatPoolAddresses")
    private @Nullable Output<List<String>> fastCreateSnatPoolAddresses;

    /**
     * @return List of address to be used for FAST-Generated SNAT Pool.
     * 
     */
    public Optional<Output<List<String>>> fastCreateSnatPoolAddresses() {
        return Optional.ofNullable(this.fastCreateSnatPoolAddresses);
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
     * Name of the FAST HTTPS application tenant.
     * 
     */
    @Import(name="tenant")
    private @Nullable Output<String> tenant;

    /**
     * @return Name of the FAST HTTPS application tenant.
     * 
     */
    public Optional<Output<String>> tenant() {
        return Optional.ofNullable(this.tenant);
    }

    /**
     * Name of an existing TLS server profile.
     * 
     */
    @Import(name="tlsServerProfileName")
    private @Nullable Output<String> tlsServerProfileName;

    /**
     * @return Name of an existing TLS server profile.
     * 
     */
    public Optional<Output<String>> tlsServerProfileName() {
        return Optional.ofNullable(this.tlsServerProfileName);
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

    private FastHttpsAppState() {}

    private FastHttpsAppState(FastHttpsAppState $) {
        this.application = $.application;
        this.createTlsServerProfile = $.createTlsServerProfile;
        this.existPoolName = $.existPoolName;
        this.existingMonitor = $.existingMonitor;
        this.existingSnatPool = $.existingSnatPool;
        this.fastCreateMonitor = $.fastCreateMonitor;
        this.fastCreatePoolMembers = $.fastCreatePoolMembers;
        this.fastCreateSnatPoolAddresses = $.fastCreateSnatPoolAddresses;
        this.loadBalancingMode = $.loadBalancingMode;
        this.slowRampTime = $.slowRampTime;
        this.tenant = $.tenant;
        this.tlsServerProfileName = $.tlsServerProfileName;
        this.virtualServer = $.virtualServer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FastHttpsAppState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FastHttpsAppState $;

        public Builder() {
            $ = new FastHttpsAppState();
        }

        public Builder(FastHttpsAppState defaults) {
            $ = new FastHttpsAppState(Objects.requireNonNull(defaults));
        }

        /**
         * @param application Name of the FAST HTTPS application.
         * 
         * @return builder
         * 
         */
        public Builder application(@Nullable Output<String> application) {
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
         * @param createTlsServerProfile `create_tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
         * See TLS Server Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder createTlsServerProfile(@Nullable Output<FastHttpsAppCreateTlsServerProfileArgs> createTlsServerProfile) {
            $.createTlsServerProfile = createTlsServerProfile;
            return this;
        }

        /**
         * @param createTlsServerProfile `create_tls_server_profile` block takes input for FAST-Generated TLS Server Profile.
         * See TLS Server Profile below for more details.
         * 
         * @return builder
         * 
         */
        public Builder createTlsServerProfile(FastHttpsAppCreateTlsServerProfileArgs createTlsServerProfile) {
            return createTlsServerProfile(Output.of(createTlsServerProfile));
        }

        /**
         * @param existPoolName Name of an existing BIG-IP pool.
         * 
         * @return builder
         * 
         */
        public Builder existPoolName(@Nullable Output<String> existPoolName) {
            $.existPoolName = existPoolName;
            return this;
        }

        /**
         * @param existPoolName Name of an existing BIG-IP pool.
         * 
         * @return builder
         * 
         */
        public Builder existPoolName(String existPoolName) {
            return existPoolName(Output.of(existPoolName));
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
         * @param fastCreateMonitor `fast_create_monitor` block takes input for FAST-Generated Pool Monitor.
         * See Pool Monitor below for more details.
         * 
         * @return builder
         * 
         */
        public Builder fastCreateMonitor(@Nullable Output<FastHttpsAppFastCreateMonitorArgs> fastCreateMonitor) {
            $.fastCreateMonitor = fastCreateMonitor;
            return this;
        }

        /**
         * @param fastCreateMonitor `fast_create_monitor` block takes input for FAST-Generated Pool Monitor.
         * See Pool Monitor below for more details.
         * 
         * @return builder
         * 
         */
        public Builder fastCreateMonitor(FastHttpsAppFastCreateMonitorArgs fastCreateMonitor) {
            return fastCreateMonitor(Output.of(fastCreateMonitor));
        }

        /**
         * @param fastCreatePoolMembers `fast_create_pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder fastCreatePoolMembers(@Nullable Output<List<FastHttpsAppFastCreatePoolMemberArgs>> fastCreatePoolMembers) {
            $.fastCreatePoolMembers = fastCreatePoolMembers;
            return this;
        }

        /**
         * @param fastCreatePoolMembers `fast_create_pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder fastCreatePoolMembers(List<FastHttpsAppFastCreatePoolMemberArgs> fastCreatePoolMembers) {
            return fastCreatePoolMembers(Output.of(fastCreatePoolMembers));
        }

        /**
         * @param fastCreatePoolMembers `fast_create_pool_members` block takes input for FAST-Generated Pool.
         * See Pool Members below for more details.
         * 
         * @return builder
         * 
         */
        public Builder fastCreatePoolMembers(FastHttpsAppFastCreatePoolMemberArgs... fastCreatePoolMembers) {
            return fastCreatePoolMembers(List.of(fastCreatePoolMembers));
        }

        /**
         * @param fastCreateSnatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder fastCreateSnatPoolAddresses(@Nullable Output<List<String>> fastCreateSnatPoolAddresses) {
            $.fastCreateSnatPoolAddresses = fastCreateSnatPoolAddresses;
            return this;
        }

        /**
         * @param fastCreateSnatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder fastCreateSnatPoolAddresses(List<String> fastCreateSnatPoolAddresses) {
            return fastCreateSnatPoolAddresses(Output.of(fastCreateSnatPoolAddresses));
        }

        /**
         * @param fastCreateSnatPoolAddresses List of address to be used for FAST-Generated SNAT Pool.
         * 
         * @return builder
         * 
         */
        public Builder fastCreateSnatPoolAddresses(String... fastCreateSnatPoolAddresses) {
            return fastCreateSnatPoolAddresses(List.of(fastCreateSnatPoolAddresses));
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
         * @param tenant Name of the FAST HTTPS application tenant.
         * 
         * @return builder
         * 
         */
        public Builder tenant(@Nullable Output<String> tenant) {
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
         * @param tlsServerProfileName Name of an existing TLS server profile.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerProfileName(@Nullable Output<String> tlsServerProfileName) {
            $.tlsServerProfileName = tlsServerProfileName;
            return this;
        }

        /**
         * @param tlsServerProfileName Name of an existing TLS server profile.
         * 
         * @return builder
         * 
         */
        public Builder tlsServerProfileName(String tlsServerProfileName) {
            return tlsServerProfileName(Output.of(tlsServerProfileName));
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

        public FastHttpsAppState build() {
            return $;
        }
    }

}
