// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.VirtualServerArgs;
import com.pulumi.f5bigip.ltm.inputs.VirtualServerState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.VirtualServer` Configures Virtual Server
 * 
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (example: `/Common/test-virtualserver` ) or `partition + directory + name` of the resource (example: `/Common/test/test-virtualserver` ).
 * When including directory in `fullpath` we have to make sure it is created in the given partition before using it.
 * 
 * ## Importing
 * 
 * An existing virtual-server can be imported into this resource by supplying virtual-server Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_virtual_server.http /Common/terraform_vs_http
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/virtualServer:VirtualServer")
public class VirtualServer extends com.pulumi.resources.CustomResource {
    /**
     * List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     * 
     */
    @Export(name="clientProfiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> clientProfiles;

    /**
     * @return List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
     * 
     */
    public Output<Optional<List<String>>> clientProfiles() {
        return Codegen.optional(this.clientProfiles);
    }
    /**
     * Specifies the maximum number of connections allowed for the virtual server.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionLimit;

    /**
     * @return Specifies the maximum number of connections allowed for the virtual server.
     * 
     */
    public Output<Integer> connectionLimit() {
        return this.connectionLimit;
    }
    @Export(name="defaultPersistenceProfile", refs={String.class}, tree="[0]")
    private Output<String> defaultPersistenceProfile;

    public Output<String> defaultPersistenceProfile() {
        return this.defaultPersistenceProfile;
    }
    /**
     * Description of Virtual server
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of Virtual server
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Destination IP
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destination;

    /**
     * @return Destination IP
     * 
     */
    public Output<Optional<String>> destination() {
        return Codegen.optional(this.destination);
    }
    /**
     * Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     * 
     */
    @Export(name="fallbackPersistenceProfile", refs={String.class}, tree="[0]")
    private Output<String> fallbackPersistenceProfile;

    /**
     * @return Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
     * 
     */
    public Output<String> fallbackPersistenceProfile() {
        return this.fallbackPersistenceProfile;
    }
    /**
     * Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
     * 
     */
    @Export(name="firewallEnforcedPolicy", refs={String.class}, tree="[0]")
    private Output<String> firewallEnforcedPolicy;

    /**
     * @return Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
     * 
     */
    public Output<String> firewallEnforcedPolicy() {
        return this.firewallEnforcedPolicy;
    }
    /**
     * Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
     * 
     */
    @Export(name="ipProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipProtocol;

    /**
     * @return Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
     * 
     */
    public Output<Optional<String>> ipProtocol() {
        return Codegen.optional(this.ipProtocol);
    }
    /**
     * The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     * 
     */
    @Export(name="irules", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> irules;

    /**
     * @return The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
     * 
     */
    public Output<Optional<List<String>>> irules() {
        return Codegen.optional(this.irules);
    }
    /**
     * Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     * 
     */
    @Export(name="mask", refs={String.class}, tree="[0]")
    private Output<String> mask;

    /**
     * @return Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
     * 
     */
    public Output<String> mask() {
        return this.mask;
    }
    /**
     * Name of the virtual server
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the virtual server
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="perFlowRequestAccessPolicy", refs={String.class}, tree="[0]")
    private Output<String> perFlowRequestAccessPolicy;

    public Output<String> perFlowRequestAccessPolicy() {
        return this.perFlowRequestAccessPolicy;
    }
    /**
     * List of persistence profiles associated with the Virtual Server.
     * 
     */
    @Export(name="persistenceProfiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> persistenceProfiles;

    /**
     * @return List of persistence profiles associated with the Virtual Server.
     * 
     */
    public Output<Optional<List<String>>> persistenceProfiles() {
        return Codegen.optional(this.persistenceProfiles);
    }
    /**
     * Specifies the policies for the virtual server.
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> policies;

    /**
     * @return Specifies the policies for the virtual server.
     * 
     */
    public Output<Optional<List<String>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * Default pool name
     * 
     */
    @Export(name="pool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pool;

    /**
     * @return Default pool name
     * 
     */
    public Output<Optional<String>> pool() {
        return Codegen.optional(this.pool);
    }
    /**
     * Listen port for the virtual server
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Listen port for the virtual server
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     * 
     */
    @Export(name="profiles", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> profiles;

    /**
     * @return List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
     * 
     */
    public Output<List<String>> profiles() {
        return this.profiles;
    }
    /**
     * Specifies the log profile applied to the virtual server.
     * 
     */
    @Export(name="securityLogProfiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityLogProfiles;

    /**
     * @return Specifies the log profile applied to the virtual server.
     * 
     */
    public Output<Optional<List<String>>> securityLogProfiles() {
        return Codegen.optional(this.securityLogProfiles);
    }
    /**
     * List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     * 
     */
    @Export(name="serverProfiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> serverProfiles;

    /**
     * @return List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
     * 
     */
    public Output<Optional<List<String>>> serverProfiles() {
        return Codegen.optional(this.serverProfiles);
    }
    /**
     * Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
     * 
     */
    @Export(name="snatpool", refs={String.class}, tree="[0]")
    private Output<String> snatpool;

    /**
     * @return Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
     * 
     */
    public Output<String> snatpool() {
        return this.snatpool;
    }
    /**
     * Specifies an IP address or network from which the virtual server will accept traffic.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Specifies an IP address or network from which the virtual server will accept traffic.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
     * 
     */
    @Export(name="sourceAddressTranslation", refs={String.class}, tree="[0]")
    private Output<String> sourceAddressTranslation;

    /**
     * @return Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
     * 
     */
    public Output<String> sourceAddressTranslation() {
        return this.sourceAddressTranslation;
    }
    /**
     * Specifies whether the system preserves the source port of the connection. The default is `preserve`.
     * 
     */
    @Export(name="sourcePort", refs={String.class}, tree="[0]")
    private Output<String> sourcePort;

    /**
     * @return Specifies whether the system preserves the source port of the connection. The default is `preserve`.
     * 
     */
    public Output<String> sourcePort() {
        return this.sourcePort;
    }
    /**
     * Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * Specifies destination traffic matching information to which the virtual server sends traffic
     * 
     */
    @Export(name="trafficmatchingCriteria", refs={String.class}, tree="[0]")
    private Output<String> trafficmatchingCriteria;

    /**
     * @return Specifies destination traffic matching information to which the virtual server sends traffic
     * 
     */
    public Output<String> trafficmatchingCriteria() {
        return this.trafficmatchingCriteria;
    }
    /**
     * Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     * 
     */
    @Export(name="translateAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> translateAddress;

    /**
     * @return Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
     * 
     */
    public Output<Optional<String>> translateAddress() {
        return Codegen.optional(this.translateAddress);
    }
    /**
     * Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     * 
     */
    @Export(name="translatePort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> translatePort;

    /**
     * @return Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
     * 
     */
    public Output<Optional<String>> translatePort() {
        return Codegen.optional(this.translatePort);
    }
    /**
     * The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
     * 
     */
    @Export(name="vlans", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> vlans;

    /**
     * @return The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlan_enabled`
     * 
     */
    public Output<Optional<List<String>>> vlans() {
        return Codegen.optional(this.vlans);
    }
    /**
     * Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     * 
     */
    @Export(name="vlansEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> vlansEnabled;

    /**
     * @return Enables the virtual server on the VLANs specified by the `vlans` option.
     * By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
     * 
     */
    public Output<Optional<Boolean>> vlansEnabled() {
        return Codegen.optional(this.vlansEnabled);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualServer(java.lang.String name) {
        this(name, VirtualServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualServer(java.lang.String name, VirtualServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualServer(java.lang.String name, VirtualServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/virtualServer:VirtualServer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VirtualServer(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/virtualServer:VirtualServer", name, state, makeResourceOptions(options, id), false);
    }

    private static VirtualServerArgs makeArgs(VirtualServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VirtualServerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VirtualServer get(java.lang.String name, Output<java.lang.String> id, @Nullable VirtualServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualServer(name, id, state, options);
    }
}
