// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualServerState extends com.pulumi.resources.ResourceArgs {

    public static final VirtualServerState Empty = new VirtualServerState();

    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     * 
     */
    @Import(name="clientProfiles")
    private @Nullable Output<List<String>> clientProfiles;

    /**
     * @return List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     * 
     */
    public Optional<Output<List<String>>> clientProfiles() {
        return Optional.ofNullable(this.clientProfiles);
    }

    @Import(name="defaultPersistenceProfile")
    private @Nullable Output<String> defaultPersistenceProfile;

    public Optional<Output<String>> defaultPersistenceProfile() {
        return Optional.ofNullable(this.defaultPersistenceProfile);
    }

    /**
     * Description of Virtual server
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of Virtual server
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Destination IP
     * 
     */
    @Import(name="destination")
    private @Nullable Output<String> destination;

    /**
     * @return Destination IP
     * 
     */
    public Optional<Output<String>> destination() {
        return Optional.ofNullable(this.destination);
    }

    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     * 
     */
    @Import(name="fallbackPersistenceProfile")
    private @Nullable Output<String> fallbackPersistenceProfile;

    /**
     * @return Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     * 
     */
    public Optional<Output<String>> fallbackPersistenceProfile() {
        return Optional.ofNullable(this.fallbackPersistenceProfile);
    }

    /**
     * Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
     * 
     */
    @Import(name="firewallEnforcedPolicy")
    private @Nullable Output<String> firewallEnforcedPolicy;

    /**
     * @return Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
     * 
     */
    public Optional<Output<String>> firewallEnforcedPolicy() {
        return Optional.ofNullable(this.firewallEnforcedPolicy);
    }

    /**
     * Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
     * 
     */
    @Import(name="ipProtocol")
    private @Nullable Output<String> ipProtocol;

    /**
     * @return Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
     * 
     */
    public Optional<Output<String>> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    /**
     * The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     * 
     */
    @Import(name="irules")
    private @Nullable Output<List<String>> irules;

    /**
     * @return The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     * 
     */
    public Optional<Output<List<String>>> irules() {
        return Optional.ofNullable(this.irules);
    }

    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     * 
     */
    @Import(name="mask")
    private @Nullable Output<String> mask;

    /**
     * @return Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     * 
     */
    public Optional<Output<String>> mask() {
        return Optional.ofNullable(this.mask);
    }

    /**
     * Name of the virtual server
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the virtual server
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="perFlowRequestAccessPolicy")
    private @Nullable Output<String> perFlowRequestAccessPolicy;

    public Optional<Output<String>> perFlowRequestAccessPolicy() {
        return Optional.ofNullable(this.perFlowRequestAccessPolicy);
    }

    /**
     * List of persistence profiles associated with the Virtual Server.
     * 
     */
    @Import(name="persistenceProfiles")
    private @Nullable Output<List<String>> persistenceProfiles;

    /**
     * @return List of persistence profiles associated with the Virtual Server.
     * 
     */
    public Optional<Output<List<String>>> persistenceProfiles() {
        return Optional.ofNullable(this.persistenceProfiles);
    }

    /**
     * Specifies the policies for the virtual server.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return Specifies the policies for the virtual server.
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * Default pool name
     * 
     */
    @Import(name="pool")
    private @Nullable Output<String> pool;

    /**
     * @return Default pool name
     * 
     */
    public Optional<Output<String>> pool() {
        return Optional.ofNullable(this.pool);
    }

    /**
     * Listen port for the virtual server
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Listen port for the virtual server
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     * 
     */
    @Import(name="profiles")
    private @Nullable Output<List<String>> profiles;

    /**
     * @return List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     * 
     */
    public Optional<Output<List<String>>> profiles() {
        return Optional.ofNullable(this.profiles);
    }

    /**
     * Specifies the log profile applied to the virtual server.
     * 
     */
    @Import(name="securityLogProfiles")
    private @Nullable Output<List<String>> securityLogProfiles;

    /**
     * @return Specifies the log profile applied to the virtual server.
     * 
     */
    public Optional<Output<List<String>>> securityLogProfiles() {
        return Optional.ofNullable(this.securityLogProfiles);
    }

    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     * 
     */
    @Import(name="serverProfiles")
    private @Nullable Output<List<String>> serverProfiles;

    /**
     * @return List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     * 
     */
    public Optional<Output<List<String>>> serverProfiles() {
        return Optional.ofNullable(this.serverProfiles);
    }

    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
     * 
     */
    @Import(name="snatpool")
    private @Nullable Output<String> snatpool;

    /**
     * @return Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
     * 
     */
    public Optional<Output<String>> snatpool() {
        return Optional.ofNullable(this.snatpool);
    }

    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return Specifies an IP address or network from which the virtual server will accept traffic.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    /**
     * Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
     * 
     */
    @Import(name="sourceAddressTranslation")
    private @Nullable Output<String> sourceAddressTranslation;

    /**
     * @return Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
     * 
     */
    public Optional<Output<String>> sourceAddressTranslation() {
        return Optional.ofNullable(this.sourceAddressTranslation);
    }

    /**
     * Specifies whether the system preserves the source port of the connection. The default is `preserve`.
     * 
     */
    @Import(name="sourcePort")
    private @Nullable Output<String> sourcePort;

    /**
     * @return Specifies whether the system preserves the source port of the connection. The default is `preserve`.
     * 
     */
    public Optional<Output<String>> sourcePort() {
        return Optional.ofNullable(this.sourcePort);
    }

    /**
     * Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Specifies destination traffic matching information to which the virtual server sends traffic
     * 
     */
    @Import(name="trafficmatchingCriteria")
    private @Nullable Output<String> trafficmatchingCriteria;

    /**
     * @return Specifies destination traffic matching information to which the virtual server sends traffic
     * 
     */
    public Optional<Output<String>> trafficmatchingCriteria() {
        return Optional.ofNullable(this.trafficmatchingCriteria);
    }

    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     * 
     */
    @Import(name="translateAddress")
    private @Nullable Output<String> translateAddress;

    /**
     * @return Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     * 
     */
    public Optional<Output<String>> translateAddress() {
        return Optional.ofNullable(this.translateAddress);
    }

    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     * 
     */
    @Import(name="translatePort")
    private @Nullable Output<String> translatePort;

    /**
     * @return Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     * 
     */
    public Optional<Output<String>> translatePort() {
        return Optional.ofNullable(this.translatePort);
    }

    /**
     * The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
     * 
     */
    @Import(name="vlans")
    private @Nullable Output<List<String>> vlans;

    /**
     * @return The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
     * 
     */
    public Optional<Output<List<String>>> vlans() {
        return Optional.ofNullable(this.vlans);
    }

    /**
     * Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     * 
     */
    @Import(name="vlansEnabled")
    private @Nullable Output<Boolean> vlansEnabled;

    /**
     * @return Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     * 
     */
    public Optional<Output<Boolean>> vlansEnabled() {
        return Optional.ofNullable(this.vlansEnabled);
    }

    private VirtualServerState() {}

    private VirtualServerState(VirtualServerState $) {
        this.clientProfiles = $.clientProfiles;
        this.defaultPersistenceProfile = $.defaultPersistenceProfile;
        this.description = $.description;
        this.destination = $.destination;
        this.fallbackPersistenceProfile = $.fallbackPersistenceProfile;
        this.firewallEnforcedPolicy = $.firewallEnforcedPolicy;
        this.ipProtocol = $.ipProtocol;
        this.irules = $.irules;
        this.mask = $.mask;
        this.name = $.name;
        this.perFlowRequestAccessPolicy = $.perFlowRequestAccessPolicy;
        this.persistenceProfiles = $.persistenceProfiles;
        this.policies = $.policies;
        this.pool = $.pool;
        this.port = $.port;
        this.profiles = $.profiles;
        this.securityLogProfiles = $.securityLogProfiles;
        this.serverProfiles = $.serverProfiles;
        this.snatpool = $.snatpool;
        this.source = $.source;
        this.sourceAddressTranslation = $.sourceAddressTranslation;
        this.sourcePort = $.sourcePort;
        this.state = $.state;
        this.trafficmatchingCriteria = $.trafficmatchingCriteria;
        this.translateAddress = $.translateAddress;
        this.translatePort = $.translatePort;
        this.vlans = $.vlans;
        this.vlansEnabled = $.vlansEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualServerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualServerState $;

        public Builder() {
            $ = new VirtualServerState();
        }

        public Builder(VirtualServerState defaults) {
            $ = new VirtualServerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientProfiles List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
         * 
         * @return builder
         * 
         */
        public Builder clientProfiles(@Nullable Output<List<String>> clientProfiles) {
            $.clientProfiles = clientProfiles;
            return this;
        }

        /**
         * @param clientProfiles List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
         * 
         * @return builder
         * 
         */
        public Builder clientProfiles(List<String> clientProfiles) {
            return clientProfiles(Output.of(clientProfiles));
        }

        /**
         * @param clientProfiles List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
         * 
         * @return builder
         * 
         */
        public Builder clientProfiles(String... clientProfiles) {
            return clientProfiles(List.of(clientProfiles));
        }

        public Builder defaultPersistenceProfile(@Nullable Output<String> defaultPersistenceProfile) {
            $.defaultPersistenceProfile = defaultPersistenceProfile;
            return this;
        }

        public Builder defaultPersistenceProfile(String defaultPersistenceProfile) {
            return defaultPersistenceProfile(Output.of(defaultPersistenceProfile));
        }

        /**
         * @param description Description of Virtual server
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of Virtual server
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param destination Destination IP
         * 
         * @return builder
         * 
         */
        public Builder destination(@Nullable Output<String> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination Destination IP
         * 
         * @return builder
         * 
         */
        public Builder destination(String destination) {
            return destination(Output.of(destination));
        }

        /**
         * @param fallbackPersistenceProfile Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
         * 
         * @return builder
         * 
         */
        public Builder fallbackPersistenceProfile(@Nullable Output<String> fallbackPersistenceProfile) {
            $.fallbackPersistenceProfile = fallbackPersistenceProfile;
            return this;
        }

        /**
         * @param fallbackPersistenceProfile Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
         * 
         * @return builder
         * 
         */
        public Builder fallbackPersistenceProfile(String fallbackPersistenceProfile) {
            return fallbackPersistenceProfile(Output.of(fallbackPersistenceProfile));
        }

        /**
         * @param firewallEnforcedPolicy Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
         * 
         * @return builder
         * 
         */
        public Builder firewallEnforcedPolicy(@Nullable Output<String> firewallEnforcedPolicy) {
            $.firewallEnforcedPolicy = firewallEnforcedPolicy;
            return this;
        }

        /**
         * @param firewallEnforcedPolicy Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
         * 
         * @return builder
         * 
         */
        public Builder firewallEnforcedPolicy(String firewallEnforcedPolicy) {
            return firewallEnforcedPolicy(Output.of(firewallEnforcedPolicy));
        }

        /**
         * @param ipProtocol Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(@Nullable Output<String> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param ipProtocol Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(String ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        /**
         * @param irules The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
         * 
         * @return builder
         * 
         */
        public Builder irules(@Nullable Output<List<String>> irules) {
            $.irules = irules;
            return this;
        }

        /**
         * @param irules The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
         * 
         * @return builder
         * 
         */
        public Builder irules(List<String> irules) {
            return irules(Output.of(irules));
        }

        /**
         * @param irules The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
         * 
         * @return builder
         * 
         */
        public Builder irules(String... irules) {
            return irules(List.of(irules));
        }

        /**
         * @param mask Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
         * 
         * @return builder
         * 
         */
        public Builder mask(@Nullable Output<String> mask) {
            $.mask = mask;
            return this;
        }

        /**
         * @param mask Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
         * 
         * @return builder
         * 
         */
        public Builder mask(String mask) {
            return mask(Output.of(mask));
        }

        /**
         * @param name Name of the virtual server
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the virtual server
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder perFlowRequestAccessPolicy(@Nullable Output<String> perFlowRequestAccessPolicy) {
            $.perFlowRequestAccessPolicy = perFlowRequestAccessPolicy;
            return this;
        }

        public Builder perFlowRequestAccessPolicy(String perFlowRequestAccessPolicy) {
            return perFlowRequestAccessPolicy(Output.of(perFlowRequestAccessPolicy));
        }

        /**
         * @param persistenceProfiles List of persistence profiles associated with the Virtual Server.
         * 
         * @return builder
         * 
         */
        public Builder persistenceProfiles(@Nullable Output<List<String>> persistenceProfiles) {
            $.persistenceProfiles = persistenceProfiles;
            return this;
        }

        /**
         * @param persistenceProfiles List of persistence profiles associated with the Virtual Server.
         * 
         * @return builder
         * 
         */
        public Builder persistenceProfiles(List<String> persistenceProfiles) {
            return persistenceProfiles(Output.of(persistenceProfiles));
        }

        /**
         * @param persistenceProfiles List of persistence profiles associated with the Virtual Server.
         * 
         * @return builder
         * 
         */
        public Builder persistenceProfiles(String... persistenceProfiles) {
            return persistenceProfiles(List.of(persistenceProfiles));
        }

        /**
         * @param policies Specifies the policies for the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies Specifies the policies for the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies Specifies the policies for the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param pool Default pool name
         * 
         * @return builder
         * 
         */
        public Builder pool(@Nullable Output<String> pool) {
            $.pool = pool;
            return this;
        }

        /**
         * @param pool Default pool name
         * 
         * @return builder
         * 
         */
        public Builder pool(String pool) {
            return pool(Output.of(pool));
        }

        /**
         * @param port Listen port for the virtual server
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Listen port for the virtual server
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param profiles List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
         * 
         * @return builder
         * 
         */
        public Builder profiles(@Nullable Output<List<String>> profiles) {
            $.profiles = profiles;
            return this;
        }

        /**
         * @param profiles List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
         * 
         * @return builder
         * 
         */
        public Builder profiles(List<String> profiles) {
            return profiles(Output.of(profiles));
        }

        /**
         * @param profiles List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
         * 
         * @return builder
         * 
         */
        public Builder profiles(String... profiles) {
            return profiles(List.of(profiles));
        }

        /**
         * @param securityLogProfiles Specifies the log profile applied to the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(@Nullable Output<List<String>> securityLogProfiles) {
            $.securityLogProfiles = securityLogProfiles;
            return this;
        }

        /**
         * @param securityLogProfiles Specifies the log profile applied to the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(List<String> securityLogProfiles) {
            return securityLogProfiles(Output.of(securityLogProfiles));
        }

        /**
         * @param securityLogProfiles Specifies the log profile applied to the virtual server.
         * 
         * @return builder
         * 
         */
        public Builder securityLogProfiles(String... securityLogProfiles) {
            return securityLogProfiles(List.of(securityLogProfiles));
        }

        /**
         * @param serverProfiles List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
         * 
         * @return builder
         * 
         */
        public Builder serverProfiles(@Nullable Output<List<String>> serverProfiles) {
            $.serverProfiles = serverProfiles;
            return this;
        }

        /**
         * @param serverProfiles List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
         * 
         * @return builder
         * 
         */
        public Builder serverProfiles(List<String> serverProfiles) {
            return serverProfiles(Output.of(serverProfiles));
        }

        /**
         * @param serverProfiles List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
         * 
         * @return builder
         * 
         */
        public Builder serverProfiles(String... serverProfiles) {
            return serverProfiles(List.of(serverProfiles));
        }

        /**
         * @param snatpool Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
         * 
         * @return builder
         * 
         */
        public Builder snatpool(@Nullable Output<String> snatpool) {
            $.snatpool = snatpool;
            return this;
        }

        /**
         * @param snatpool Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
         * 
         * @return builder
         * 
         */
        public Builder snatpool(String snatpool) {
            return snatpool(Output.of(snatpool));
        }

        /**
         * @param source Specifies an IP address or network from which the virtual server will accept traffic.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source Specifies an IP address or network from which the virtual server will accept traffic.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        /**
         * @param sourceAddressTranslation Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
         * 
         * @return builder
         * 
         */
        public Builder sourceAddressTranslation(@Nullable Output<String> sourceAddressTranslation) {
            $.sourceAddressTranslation = sourceAddressTranslation;
            return this;
        }

        /**
         * @param sourceAddressTranslation Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
         * 
         * @return builder
         * 
         */
        public Builder sourceAddressTranslation(String sourceAddressTranslation) {
            return sourceAddressTranslation(Output.of(sourceAddressTranslation));
        }

        /**
         * @param sourcePort Specifies whether the system preserves the source port of the connection. The default is `preserve`.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(@Nullable Output<String> sourcePort) {
            $.sourcePort = sourcePort;
            return this;
        }

        /**
         * @param sourcePort Specifies whether the system preserves the source port of the connection. The default is `preserve`.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(String sourcePort) {
            return sourcePort(Output.of(sourcePort));
        }

        /**
         * @param state Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param trafficmatchingCriteria Specifies destination traffic matching information to which the virtual server sends traffic
         * 
         * @return builder
         * 
         */
        public Builder trafficmatchingCriteria(@Nullable Output<String> trafficmatchingCriteria) {
            $.trafficmatchingCriteria = trafficmatchingCriteria;
            return this;
        }

        /**
         * @param trafficmatchingCriteria Specifies destination traffic matching information to which the virtual server sends traffic
         * 
         * @return builder
         * 
         */
        public Builder trafficmatchingCriteria(String trafficmatchingCriteria) {
            return trafficmatchingCriteria(Output.of(trafficmatchingCriteria));
        }

        /**
         * @param translateAddress Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
         * 
         * @return builder
         * 
         */
        public Builder translateAddress(@Nullable Output<String> translateAddress) {
            $.translateAddress = translateAddress;
            return this;
        }

        /**
         * @param translateAddress Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
         * 
         * @return builder
         * 
         */
        public Builder translateAddress(String translateAddress) {
            return translateAddress(Output.of(translateAddress));
        }

        /**
         * @param translatePort Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
         * 
         * @return builder
         * 
         */
        public Builder translatePort(@Nullable Output<String> translatePort) {
            $.translatePort = translatePort;
            return this;
        }

        /**
         * @param translatePort Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
         * 
         * @return builder
         * 
         */
        public Builder translatePort(String translatePort) {
            return translatePort(Output.of(translatePort));
        }

        /**
         * @param vlans The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
         * 
         * @return builder
         * 
         */
        public Builder vlans(@Nullable Output<List<String>> vlans) {
            $.vlans = vlans;
            return this;
        }

        /**
         * @param vlans The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
         * 
         * @return builder
         * 
         */
        public Builder vlans(List<String> vlans) {
            return vlans(Output.of(vlans));
        }

        /**
         * @param vlans The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
         * 
         * @return builder
         * 
         */
        public Builder vlans(String... vlans) {
            return vlans(List.of(vlans));
        }

        /**
         * @param vlansEnabled Enables the virtual server on the VLANs specified by the `vlans` option.
         * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
         * 
         * @return builder
         * 
         */
        public Builder vlansEnabled(@Nullable Output<Boolean> vlansEnabled) {
            $.vlansEnabled = vlansEnabled;
            return this;
        }

        /**
         * @param vlansEnabled Enables the virtual server on the VLANs specified by the `vlans` option.
         * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
         * 
         * @return builder
         * 
         */
        public Builder vlansEnabled(Boolean vlansEnabled) {
            return vlansEnabled(Output.of(vlansEnabled));
        }

        public VirtualServerState build() {
            return $;
        }
    }

}