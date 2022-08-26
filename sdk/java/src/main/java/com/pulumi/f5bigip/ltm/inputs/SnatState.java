// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.ltm.inputs.SnatOriginArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SnatState extends com.pulumi.resources.ResourceArgs {

    public static final SnatState Empty = new SnatState();

    /**
     * -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default.
     * 
     */
    @Import(name="autolasthop")
    private @Nullable Output<String> autolasthop;

    /**
     * @return -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default.
     * 
     */
    public Optional<Output<String>> autolasthop() {
        return Optional.ofNullable(this.autolasthop);
    }

    /**
     * Fullpath
     * 
     */
    @Import(name="fullPath")
    private @Nullable Output<String> fullPath;

    /**
     * @return Fullpath
     * 
     */
    public Optional<Output<String>> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }

    /**
     * Enables or disables mirroring of SNAT connections.
     * 
     */
    @Import(name="mirror")
    private @Nullable Output<String> mirror;

    /**
     * @return Enables or disables mirroring of SNAT connections.
     * 
     */
    public Optional<Output<String>> mirror() {
        return Optional.ofNullable(this.mirror);
    }

    /**
     * Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
     * 
     */
    @Import(name="origins")
    private @Nullable Output<List<SnatOriginArgs>> origins;

    /**
     * @return Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
     * 
     */
    public Optional<Output<List<SnatOriginArgs>>> origins() {
        return Optional.ofNullable(this.origins);
    }

    /**
     * Partition or path to which the SNAT belongs
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Partition or path to which the SNAT belongs
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    /**
     * Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
     * 
     */
    @Import(name="snatpool")
    private @Nullable Output<String> snatpool;

    /**
     * @return Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
     * 
     */
    public Optional<Output<String>> snatpool() {
        return Optional.ofNullable(this.snatpool);
    }

    /**
     * Specifies how the SNAT object handles the client&#39;s source port. The default is `preserve`.
     * 
     */
    @Import(name="sourceport")
    private @Nullable Output<String> sourceport;

    /**
     * @return Specifies how the SNAT object handles the client&#39;s source port. The default is `preserve`.
     * 
     */
    public Optional<Output<String>> sourceport() {
        return Optional.ofNullable(this.sourceport);
    }

    /**
     * Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
     * 
     */
    @Import(name="translation")
    private @Nullable Output<String> translation;

    /**
     * @return Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
     * 
     */
    public Optional<Output<String>> translation() {
        return Optional.ofNullable(this.translation);
    }

    /**
     * Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
     * 
     */
    @Import(name="vlans")
    private @Nullable Output<List<String>> vlans;

    /**
     * @return Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
     * 
     */
    public Optional<Output<List<String>>> vlans() {
        return Optional.ofNullable(this.vlans);
    }

    /**
     * Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
     * 
     */
    @Import(name="vlansdisabled")
    private @Nullable Output<Boolean> vlansdisabled;

    /**
     * @return Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
     * 
     */
    public Optional<Output<Boolean>> vlansdisabled() {
        return Optional.ofNullable(this.vlansdisabled);
    }

    private SnatState() {}

    private SnatState(SnatState $) {
        this.autolasthop = $.autolasthop;
        this.fullPath = $.fullPath;
        this.mirror = $.mirror;
        this.name = $.name;
        this.origins = $.origins;
        this.partition = $.partition;
        this.snatpool = $.snatpool;
        this.sourceport = $.sourceport;
        this.translation = $.translation;
        this.vlans = $.vlans;
        this.vlansdisabled = $.vlansdisabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnatState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnatState $;

        public Builder() {
            $ = new SnatState();
        }

        public Builder(SnatState defaults) {
            $ = new SnatState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autolasthop -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default.
         * 
         * @return builder
         * 
         */
        public Builder autolasthop(@Nullable Output<String> autolasthop) {
            $.autolasthop = autolasthop;
            return this;
        }

        /**
         * @param autolasthop -(Optional) Specifies whether to automatically map last hop for pools or not. The default is to use next level&#39;s default.
         * 
         * @return builder
         * 
         */
        public Builder autolasthop(String autolasthop) {
            return autolasthop(Output.of(autolasthop));
        }

        /**
         * @param fullPath Fullpath
         * 
         * @return builder
         * 
         */
        public Builder fullPath(@Nullable Output<String> fullPath) {
            $.fullPath = fullPath;
            return this;
        }

        /**
         * @param fullPath Fullpath
         * 
         * @return builder
         * 
         */
        public Builder fullPath(String fullPath) {
            return fullPath(Output.of(fullPath));
        }

        /**
         * @param mirror Enables or disables mirroring of SNAT connections.
         * 
         * @return builder
         * 
         */
        public Builder mirror(@Nullable Output<String> mirror) {
            $.mirror = mirror;
            return this;
        }

        /**
         * @param mirror Enables or disables mirroring of SNAT connections.
         * 
         * @return builder
         * 
         */
        public Builder mirror(String mirror) {
            return mirror(Output.of(mirror));
        }

        /**
         * @param name Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param origins Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
         * 
         * @return builder
         * 
         */
        public Builder origins(@Nullable Output<List<SnatOriginArgs>> origins) {
            $.origins = origins;
            return this;
        }

        /**
         * @param origins Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
         * 
         * @return builder
         * 
         */
        public Builder origins(List<SnatOriginArgs> origins) {
            return origins(Output.of(origins));
        }

        /**
         * @param origins Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
         * 
         * @return builder
         * 
         */
        public Builder origins(SnatOriginArgs... origins) {
            return origins(List.of(origins));
        }

        /**
         * @param partition Partition or path to which the SNAT belongs
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Partition or path to which the SNAT belongs
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param snatpool Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
         * 
         * @return builder
         * 
         */
        public Builder snatpool(@Nullable Output<String> snatpool) {
            $.snatpool = snatpool;
            return this;
        }

        /**
         * @param snatpool Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
         * 
         * @return builder
         * 
         */
        public Builder snatpool(String snatpool) {
            return snatpool(Output.of(snatpool));
        }

        /**
         * @param sourceport Specifies how the SNAT object handles the client&#39;s source port. The default is `preserve`.
         * 
         * @return builder
         * 
         */
        public Builder sourceport(@Nullable Output<String> sourceport) {
            $.sourceport = sourceport;
            return this;
        }

        /**
         * @param sourceport Specifies how the SNAT object handles the client&#39;s source port. The default is `preserve`.
         * 
         * @return builder
         * 
         */
        public Builder sourceport(String sourceport) {
            return sourceport(Output.of(sourceport));
        }

        /**
         * @param translation Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
         * 
         * @return builder
         * 
         */
        public Builder translation(@Nullable Output<String> translation) {
            $.translation = translation;
            return this;
        }

        /**
         * @param translation Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
         * 
         * @return builder
         * 
         */
        public Builder translation(String translation) {
            return translation(Output.of(translation));
        }

        /**
         * @param vlans Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
         * 
         * @return builder
         * 
         */
        public Builder vlans(@Nullable Output<List<String>> vlans) {
            $.vlans = vlans;
            return this;
        }

        /**
         * @param vlans Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
         * 
         * @return builder
         * 
         */
        public Builder vlans(List<String> vlans) {
            return vlans(Output.of(vlans));
        }

        /**
         * @param vlans Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
         * 
         * @return builder
         * 
         */
        public Builder vlans(String... vlans) {
            return vlans(List.of(vlans));
        }

        /**
         * @param vlansdisabled Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
         * 
         * @return builder
         * 
         */
        public Builder vlansdisabled(@Nullable Output<Boolean> vlansdisabled) {
            $.vlansdisabled = vlansdisabled;
            return this;
        }

        /**
         * @param vlansdisabled Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
         * 
         * @return builder
         * 
         */
        public Builder vlansdisabled(Boolean vlansdisabled) {
            return vlansdisabled(Output.of(vlansdisabled));
        }

        public SnatState build() {
            return $;
        }
    }

}
