// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProfileFastL4Args extends com.pulumi.resources.ResourceArgs {

    public static final ProfileFastL4Args Empty = new ProfileFastL4Args();

    /**
     * Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     * 
     */
    @Import(name="clientTimeout")
    private @Nullable Output<Integer> clientTimeout;

    /**
     * @return Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
     * 
     */
    public Optional<Output<Integer>> clientTimeout() {
        return Optional.ofNullable(this.clientTimeout);
    }

    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    @Import(name="defaultsFrom")
    private @Nullable Output<String> defaultsFrom;

    /**
     * @return Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    public Optional<Output<String>> defaultsFrom() {
        return Optional.ofNullable(this.defaultsFrom);
    }

    /**
     * Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     * 
     */
    @Import(name="explicitflowMigration")
    private @Nullable Output<String> explicitflowMigration;

    /**
     * @return Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
     * 
     */
    public Optional<Output<String>> explicitflowMigration() {
        return Optional.ofNullable(this.explicitflowMigration);
    }

    /**
     * Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     * 
     */
    @Import(name="hardwareSyncookie")
    private @Nullable Output<String> hardwareSyncookie;

    /**
     * @return Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
     * 
     */
    public Optional<Output<String>> hardwareSyncookie() {
        return Optional.ofNullable(this.hardwareSyncookie);
    }

    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    @Import(name="idleTimeout")
    private @Nullable Output<String> idleTimeout;

    /**
     * @return Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    public Optional<Output<String>> idleTimeout() {
        return Optional.ofNullable(this.idleTimeout);
    }

    /**
     * Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    @Import(name="iptosToclient")
    private @Nullable Output<String> iptosToclient;

    /**
     * @return Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    public Optional<Output<String>> iptosToclient() {
        return Optional.ofNullable(this.iptosToclient);
    }

    /**
     * Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    @Import(name="iptosToserver")
    private @Nullable Output<String> iptosToserver;

    /**
     * @return Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
     * 
     */
    public Optional<Output<String>> iptosToserver() {
        return Optional.ofNullable(this.iptosToserver);
    }

    /**
     * Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     * 
     */
    @Import(name="keepaliveInterval")
    private @Nullable Output<String> keepaliveInterval;

    /**
     * @return Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
     * 
     */
    public Optional<Output<String>> keepaliveInterval() {
        return Optional.ofNullable(this.keepaliveInterval);
    }

    /**
     * Name of the profile_fastl4
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the profile_fastl4
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Displays the administrative partition within which this profile resides
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Displays the administrative partition within which this profile resides
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    private ProfileFastL4Args() {}

    private ProfileFastL4Args(ProfileFastL4Args $) {
        this.clientTimeout = $.clientTimeout;
        this.defaultsFrom = $.defaultsFrom;
        this.explicitflowMigration = $.explicitflowMigration;
        this.hardwareSyncookie = $.hardwareSyncookie;
        this.idleTimeout = $.idleTimeout;
        this.iptosToclient = $.iptosToclient;
        this.iptosToserver = $.iptosToserver;
        this.keepaliveInterval = $.keepaliveInterval;
        this.name = $.name;
        this.partition = $.partition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProfileFastL4Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProfileFastL4Args $;

        public Builder() {
            $ = new ProfileFastL4Args();
        }

        public Builder(ProfileFastL4Args defaults) {
            $ = new ProfileFastL4Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientTimeout Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
         * 
         * @return builder
         * 
         */
        public Builder clientTimeout(@Nullable Output<Integer> clientTimeout) {
            $.clientTimeout = clientTimeout;
            return this;
        }

        /**
         * @param clientTimeout Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
         * 
         * @return builder
         * 
         */
        public Builder clientTimeout(Integer clientTimeout) {
            return clientTimeout(Output.of(clientTimeout));
        }

        /**
         * @param defaultsFrom Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(@Nullable Output<String> defaultsFrom) {
            $.defaultsFrom = defaultsFrom;
            return this;
        }

        /**
         * @param defaultsFrom Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
         * 
         * @return builder
         * 
         */
        public Builder defaultsFrom(String defaultsFrom) {
            return defaultsFrom(Output.of(defaultsFrom));
        }

        /**
         * @param explicitflowMigration Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
         * 
         * @return builder
         * 
         */
        public Builder explicitflowMigration(@Nullable Output<String> explicitflowMigration) {
            $.explicitflowMigration = explicitflowMigration;
            return this;
        }

        /**
         * @param explicitflowMigration Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
         * 
         * @return builder
         * 
         */
        public Builder explicitflowMigration(String explicitflowMigration) {
            return explicitflowMigration(Output.of(explicitflowMigration));
        }

        /**
         * @param hardwareSyncookie Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
         * 
         * @return builder
         * 
         */
        public Builder hardwareSyncookie(@Nullable Output<String> hardwareSyncookie) {
            $.hardwareSyncookie = hardwareSyncookie;
            return this;
        }

        /**
         * @param hardwareSyncookie Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the &#34;/sys modify db&#34; command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
         * 
         * @return builder
         * 
         */
        public Builder hardwareSyncookie(String hardwareSyncookie) {
            return hardwareSyncookie(Output.of(hardwareSyncookie));
        }

        /**
         * @param idleTimeout Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(@Nullable Output<String> idleTimeout) {
            $.idleTimeout = idleTimeout;
            return this;
        }

        /**
         * @param idleTimeout Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(String idleTimeout) {
            return idleTimeout(Output.of(idleTimeout));
        }

        /**
         * @param iptosToclient Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
         * 
         * @return builder
         * 
         */
        public Builder iptosToclient(@Nullable Output<String> iptosToclient) {
            $.iptosToclient = iptosToclient;
            return this;
        }

        /**
         * @param iptosToclient Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
         * 
         * @return builder
         * 
         */
        public Builder iptosToclient(String iptosToclient) {
            return iptosToclient(Output.of(iptosToclient));
        }

        /**
         * @param iptosToserver Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
         * 
         * @return builder
         * 
         */
        public Builder iptosToserver(@Nullable Output<String> iptosToserver) {
            $.iptosToserver = iptosToserver;
            return this;
        }

        /**
         * @param iptosToserver Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
         * 
         * @return builder
         * 
         */
        public Builder iptosToserver(String iptosToserver) {
            return iptosToserver(Output.of(iptosToserver));
        }

        /**
         * @param keepaliveInterval Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
         * 
         * @return builder
         * 
         */
        public Builder keepaliveInterval(@Nullable Output<String> keepaliveInterval) {
            $.keepaliveInterval = keepaliveInterval;
            return this;
        }

        /**
         * @param keepaliveInterval Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
         * 
         * @return builder
         * 
         */
        public Builder keepaliveInterval(String keepaliveInterval) {
            return keepaliveInterval(Output.of(keepaliveInterval));
        }

        /**
         * @param name Name of the profile_fastl4
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the profile_fastl4
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition Displays the administrative partition within which this profile resides
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Displays the administrative partition within which this profile resides
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        public ProfileFastL4Args build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
