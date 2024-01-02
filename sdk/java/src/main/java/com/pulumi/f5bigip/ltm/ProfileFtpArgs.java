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


public final class ProfileFtpArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProfileFtpArgs Empty = new ProfileFtpArgs();

    /**
     * Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
     * 
     */
    @Import(name="allowActiveMode")
    private @Nullable Output<String> allowActiveMode;

    /**
     * @return Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
     * 
     */
    public Optional<Output<String>> allowActiveMode() {
        return Optional.ofNullable(this.allowActiveMode);
    }

    /**
     * Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
     * 
     */
    @Import(name="allowFtps")
    private @Nullable Output<String> allowFtps;

    /**
     * @return Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
     * 
     */
    public Optional<Output<String>> allowFtps() {
        return Optional.ofNullable(this.allowFtps);
    }

    /**
     * The application service to which the object belongs.
     * 
     */
    @Import(name="appService")
    private @Nullable Output<String> appService;

    /**
     * @return The application service to which the object belongs.
     * 
     */
    public Optional<Output<String>> appService() {
        return Optional.ofNullable(this.appService);
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
     * User defined description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User defined description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
     * 
     */
    @Import(name="enforceTlssessionReuse")
    private @Nullable Output<String> enforceTlssessionReuse;

    /**
     * @return Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
     * 
     */
    public Optional<Output<String>> enforceTlssessionReuse() {
        return Optional.ofNullable(this.enforceTlssessionReuse);
    }

    /**
     * Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
     * 
     */
    @Import(name="ftpsMode")
    private @Nullable Output<String> ftpsMode;

    /**
     * @return Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
     * 
     */
    public Optional<Output<String>> ftpsMode() {
        return Optional.ofNullable(this.ftpsMode);
    }

    /**
     * Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     * 
     */
    @Import(name="inheritParentProfile")
    private @Nullable Output<String> inheritParentProfile;

    /**
     * @return Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     * 
     */
    public Optional<Output<String>> inheritParentProfile() {
        return Optional.ofNullable(this.inheritParentProfile);
    }

    /**
     * inherent vlan list
     * 
     */
    @Import(name="inheritVlanList")
    private @Nullable Output<String> inheritVlanList;

    /**
     * @return inherent vlan list
     * 
     */
    public Optional<Output<String>> inheritVlanList() {
        return Optional.ofNullable(this.inheritVlanList);
    }

    /**
     * Configures the ALG log profile that controls logging
     * 
     */
    @Import(name="logProfile")
    private @Nullable Output<String> logProfile;

    /**
     * @return Configures the ALG log profile that controls logging
     * 
     */
    public Optional<Output<String>> logProfile() {
        return Optional.ofNullable(this.logProfile);
    }

    /**
     * Configures the log publisher that handles events logging for this profile
     * 
     */
    @Import(name="logPublisher")
    private @Nullable Output<String> logPublisher;

    /**
     * @return Configures the log publisher that handles events logging for this profile
     * 
     */
    public Optional<Output<String>> logPublisher() {
        return Optional.ofNullable(this.logPublisher);
    }

    /**
     * Name of the profile_ftp
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the profile_ftp
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

    /**
     * Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     * 
     */
    @Import(name="security")
    private @Nullable Output<String> security;

    /**
     * @return Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     * 
     */
    public Optional<Output<String>> security() {
        return Optional.ofNullable(this.security);
    }

    /**
     * Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
     * 
     */
    @Import(name="translateExtended")
    private @Nullable Output<String> translateExtended;

    /**
     * @return Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
     * 
     */
    public Optional<Output<String>> translateExtended() {
        return Optional.ofNullable(this.translateExtended);
    }

    private ProfileFtpArgs() {}

    private ProfileFtpArgs(ProfileFtpArgs $) {
        this.allowActiveMode = $.allowActiveMode;
        this.allowFtps = $.allowFtps;
        this.appService = $.appService;
        this.defaultsFrom = $.defaultsFrom;
        this.description = $.description;
        this.enforceTlssessionReuse = $.enforceTlssessionReuse;
        this.ftpsMode = $.ftpsMode;
        this.inheritParentProfile = $.inheritParentProfile;
        this.inheritVlanList = $.inheritVlanList;
        this.logProfile = $.logProfile;
        this.logPublisher = $.logPublisher;
        this.name = $.name;
        this.partition = $.partition;
        this.port = $.port;
        this.security = $.security;
        this.translateExtended = $.translateExtended;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProfileFtpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProfileFtpArgs $;

        public Builder() {
            $ = new ProfileFtpArgs();
        }

        public Builder(ProfileFtpArgs defaults) {
            $ = new ProfileFtpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowActiveMode Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
         * 
         * @return builder
         * 
         */
        public Builder allowActiveMode(@Nullable Output<String> allowActiveMode) {
            $.allowActiveMode = allowActiveMode;
            return this;
        }

        /**
         * @param allowActiveMode Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
         * 
         * @return builder
         * 
         */
        public Builder allowActiveMode(String allowActiveMode) {
            return allowActiveMode(Output.of(allowActiveMode));
        }

        /**
         * @param allowFtps Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
         * 
         * @return builder
         * 
         */
        public Builder allowFtps(@Nullable Output<String> allowFtps) {
            $.allowFtps = allowFtps;
            return this;
        }

        /**
         * @param allowFtps Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
         * 
         * @return builder
         * 
         */
        public Builder allowFtps(String allowFtps) {
            return allowFtps(Output.of(allowFtps));
        }

        /**
         * @param appService The application service to which the object belongs.
         * 
         * @return builder
         * 
         */
        public Builder appService(@Nullable Output<String> appService) {
            $.appService = appService;
            return this;
        }

        /**
         * @param appService The application service to which the object belongs.
         * 
         * @return builder
         * 
         */
        public Builder appService(String appService) {
            return appService(Output.of(appService));
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
         * @param description User defined description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User defined description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enforceTlssessionReuse Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
         * 
         * @return builder
         * 
         */
        public Builder enforceTlssessionReuse(@Nullable Output<String> enforceTlssessionReuse) {
            $.enforceTlssessionReuse = enforceTlssessionReuse;
            return this;
        }

        /**
         * @param enforceTlssessionReuse Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
         * 
         * @return builder
         * 
         */
        public Builder enforceTlssessionReuse(String enforceTlssessionReuse) {
            return enforceTlssessionReuse(Output.of(enforceTlssessionReuse));
        }

        /**
         * @param ftpsMode Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
         * 
         * @return builder
         * 
         */
        public Builder ftpsMode(@Nullable Output<String> ftpsMode) {
            $.ftpsMode = ftpsMode;
            return this;
        }

        /**
         * @param ftpsMode Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
         * 
         * @return builder
         * 
         */
        public Builder ftpsMode(String ftpsMode) {
            return ftpsMode(Output.of(ftpsMode));
        }

        /**
         * @param inheritParentProfile Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
         * FastL4 only.
         * 
         * @return builder
         * 
         */
        public Builder inheritParentProfile(@Nullable Output<String> inheritParentProfile) {
            $.inheritParentProfile = inheritParentProfile;
            return this;
        }

        /**
         * @param inheritParentProfile Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
         * FastL4 only.
         * 
         * @return builder
         * 
         */
        public Builder inheritParentProfile(String inheritParentProfile) {
            return inheritParentProfile(Output.of(inheritParentProfile));
        }

        /**
         * @param inheritVlanList inherent vlan list
         * 
         * @return builder
         * 
         */
        public Builder inheritVlanList(@Nullable Output<String> inheritVlanList) {
            $.inheritVlanList = inheritVlanList;
            return this;
        }

        /**
         * @param inheritVlanList inherent vlan list
         * 
         * @return builder
         * 
         */
        public Builder inheritVlanList(String inheritVlanList) {
            return inheritVlanList(Output.of(inheritVlanList));
        }

        /**
         * @param logProfile Configures the ALG log profile that controls logging
         * 
         * @return builder
         * 
         */
        public Builder logProfile(@Nullable Output<String> logProfile) {
            $.logProfile = logProfile;
            return this;
        }

        /**
         * @param logProfile Configures the ALG log profile that controls logging
         * 
         * @return builder
         * 
         */
        public Builder logProfile(String logProfile) {
            return logProfile(Output.of(logProfile));
        }

        /**
         * @param logPublisher Configures the log publisher that handles events logging for this profile
         * 
         * @return builder
         * 
         */
        public Builder logPublisher(@Nullable Output<String> logPublisher) {
            $.logPublisher = logPublisher;
            return this;
        }

        /**
         * @param logPublisher Configures the log publisher that handles events logging for this profile
         * 
         * @return builder
         * 
         */
        public Builder logPublisher(String logPublisher) {
            return logPublisher(Output.of(logPublisher));
        }

        /**
         * @param name Name of the profile_ftp
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the profile_ftp
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

        /**
         * @param port Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param security Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
         * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
         * 
         * @return builder
         * 
         */
        public Builder security(@Nullable Output<String> security) {
            $.security = security;
            return this;
        }

        /**
         * @param security Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
         * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
         * 
         * @return builder
         * 
         */
        public Builder security(String security) {
            return security(Output.of(security));
        }

        /**
         * @param translateExtended Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
         * 
         * @return builder
         * 
         */
        public Builder translateExtended(@Nullable Output<String> translateExtended) {
            $.translateExtended = translateExtended;
            return this;
        }

        /**
         * @param translateExtended Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
         * 
         * @return builder
         * 
         */
        public Builder translateExtended(String translateExtended) {
            return translateExtended(Output.of(translateExtended));
        }

        public ProfileFtpArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("ProfileFtpArgs", "name");
            }
            return $;
        }
    }

}
