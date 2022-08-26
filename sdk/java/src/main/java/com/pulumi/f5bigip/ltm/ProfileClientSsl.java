// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileClientSslArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileClientSslState;
import com.pulumi.f5bigip.ltm.outputs.ProfileClientSslCertKeyChain;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileClientSsl` Manages client SSL profiles on a BIG-IP
 * 
 * Resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.ProfileClientSsl;
 * import com.pulumi.f5bigip.ltm.ProfileClientSslArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test_ClientSsl = new ProfileClientSsl(&#34;test-ClientSsl&#34;, ProfileClientSslArgs.builder()        
 *             .authenticate(&#34;always&#34;)
 *             .ciphers(&#34;DEFAULT&#34;)
 *             .defaultsFrom(&#34;/Common/clientssl&#34;)
 *             .name(&#34;/Common/test-ClientSsl&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileClientSsl:ProfileClientSsl")
public class ProfileClientSsl extends com.pulumi.resources.CustomResource {
    /**
     * Alert time out
     * 
     */
    @Export(name="alertTimeout", type=String.class, parameters={})
    private Output<String> alertTimeout;

    /**
     * @return Alert time out
     * 
     */
    public Output<String> alertTimeout() {
        return this.alertTimeout;
    }
    /**
     * Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
     * 
     */
    @Export(name="allowNonSsl", type=String.class, parameters={})
    private Output<String> allowNonSsl;

    /**
     * @return Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
     * 
     */
    public Output<String> allowNonSsl() {
        return this.allowNonSsl;
    }
    /**
     * Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
     * When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
     * 
     */
    @Export(name="authenticate", type=String.class, parameters={})
    private Output<String> authenticate;

    /**
     * @return Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
     * When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
     * 
     */
    public Output<String> authenticate() {
        return this.authenticate;
    }
    /**
     * Specifies the maximum number of certificates to be traversed in a client certificate chain
     * 
     */
    @Export(name="authenticateDepth", type=Integer.class, parameters={})
    private Output<Integer> authenticateDepth;

    /**
     * @return Specifies the maximum number of certificates to be traversed in a client certificate chain
     * 
     */
    public Output<Integer> authenticateDepth() {
        return this.authenticateDepth;
    }
    /**
     * Specifies the client certificate to use in SSL client certificate constrained delegation. This certificate will be used if client does not provide a cert during the SSL handshake. The default value is none.
     * 
     */
    @Export(name="c3dClientFallbackCert", type=String.class, parameters={})
    private Output<String> c3dClientFallbackCert;

    /**
     * @return Specifies the client certificate to use in SSL client certificate constrained delegation. This certificate will be used if client does not provide a cert during the SSL handshake. The default value is none.
     * 
     */
    public Output<String> c3dClientFallbackCert() {
        return this.c3dClientFallbackCert;
    }
    /**
     * Specifies the BIG-IP action when the OCSP responder returns unknown status. The default value is drop, which causes the onnection to be dropped. Conversely, you can specify ignore, which causes the connection to ignore the unknown status and continue.
     * 
     */
    @Export(name="c3dDropUnknownOcspStatus", type=String.class, parameters={})
    private Output<String> c3dDropUnknownOcspStatus;

    /**
     * @return Specifies the BIG-IP action when the OCSP responder returns unknown status. The default value is drop, which causes the onnection to be dropped. Conversely, you can specify ignore, which causes the connection to ignore the unknown status and continue.
     * 
     */
    public Output<String> c3dDropUnknownOcspStatus() {
        return this.c3dDropUnknownOcspStatus;
    }
    /**
     * Specifies the SSL client certificate constrained delegation OCSP object that the BIG-IP SSL should use to connect to the OCSP responder and check the client certificate status.
     * 
     */
    @Export(name="c3dOcsp", type=String.class, parameters={})
    private Output<String> c3dOcsp;

    /**
     * @return Specifies the SSL client certificate constrained delegation OCSP object that the BIG-IP SSL should use to connect to the OCSP responder and check the client certificate status.
     * 
     */
    public Output<String> c3dOcsp() {
        return this.c3dOcsp;
    }
    /**
     * Client certificate file path. Default None.
     * 
     */
    @Export(name="caFile", type=String.class, parameters={})
    private Output<String> caFile;

    /**
     * @return Client certificate file path. Default None.
     * 
     */
    public Output<String> caFile() {
        return this.caFile;
    }
    /**
     * Cache size (sessions).
     * 
     */
    @Export(name="cacheSize", type=Integer.class, parameters={})
    private Output<Integer> cacheSize;

    /**
     * @return Cache size (sessions).
     * 
     */
    public Output<Integer> cacheSize() {
        return this.cacheSize;
    }
    /**
     * Cache time out
     * 
     */
    @Export(name="cacheTimeout", type=Integer.class, parameters={})
    private Output<Integer> cacheTimeout;

    /**
     * @return Cache time out
     * 
     */
    public Output<Integer> cacheTimeout() {
        return this.cacheTimeout;
    }
    /**
     * Specifies a cert name for use.
     * 
     */
    @Export(name="cert", type=String.class, parameters={})
    private Output<String> cert;

    /**
     * @return Specifies a cert name for use.
     * 
     */
    public Output<String> cert() {
        return this.cert;
    }
    /**
     * Cert extension includes for ssl forward proxy
     * 
     */
    @Export(name="certExtensionIncludes", type=List.class, parameters={String.class})
    private Output<List<String>> certExtensionIncludes;

    /**
     * @return Cert extension includes for ssl forward proxy
     * 
     */
    public Output<List<String>> certExtensionIncludes() {
        return this.certExtensionIncludes;
    }
    @Export(name="certKeyChains", type=List.class, parameters={ProfileClientSslCertKeyChain.class})
    private Output<List<ProfileClientSslCertKeyChain>> certKeyChains;

    public Output<List<ProfileClientSslCertKeyChain>> certKeyChains() {
        return this.certKeyChains;
    }
    /**
     * Life span of the certificate in days for ssl forward proxy
     * 
     */
    @Export(name="certLifeSpan", type=Integer.class, parameters={})
    private Output<Integer> certLifeSpan;

    /**
     * @return Life span of the certificate in days for ssl forward proxy
     * 
     */
    public Output<Integer> certLifeSpan() {
        return this.certLifeSpan;
    }
    /**
     * Cert lookup by ip address and port enabled / disabled
     * 
     */
    @Export(name="certLookupByIpaddrPort", type=String.class, parameters={})
    private Output<String> certLookupByIpaddrPort;

    /**
     * @return Cert lookup by ip address and port enabled / disabled
     * 
     */
    public Output<String> certLookupByIpaddrPort() {
        return this.certLookupByIpaddrPort;
    }
    /**
     * Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
     * 
     */
    @Export(name="chain", type=String.class, parameters={})
    private Output<String> chain;

    /**
     * @return Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
     * 
     */
    public Output<String> chain() {
        return this.chain;
    }
    /**
     * Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
     * 
     */
    @Export(name="ciphers", type=String.class, parameters={})
    private Output<String> ciphers;

    /**
     * @return Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
     * 
     */
    public Output<String> ciphers() {
        return this.ciphers;
    }
    /**
     * client certificate name
     * 
     */
    @Export(name="clientCertCa", type=String.class, parameters={})
    private Output<String> clientCertCa;

    /**
     * @return client certificate name
     * 
     */
    public Output<String> clientCertCa() {
        return this.clientCertCa;
    }
    /**
     * Certificate revocation file name
     * 
     */
    @Export(name="crlFile", type=String.class, parameters={})
    private Output<String> crlFile;

    /**
     * @return Certificate revocation file name
     * 
     */
    public Output<String> crlFile() {
        return this.crlFile;
    }
    /**
     * Parent profile for this clientssl profile.Once this value has been set, it cannot be changed. Default value is `/Common/clientssl`. It Should Full path `/partition/profile_name`
     * 
     */
    @Export(name="defaultsFrom", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultsFrom;

    /**
     * @return Parent profile for this clientssl profile.Once this value has been set, it cannot be changed. Default value is `/Common/clientssl`. It Should Full path `/partition/profile_name`
     * 
     */
    public Output<Optional<String>> defaultsFrom() {
        return Codegen.optional(this.defaultsFrom);
    }
    /**
     * Forward proxy bypass default action. (enabled / disabled)
     * 
     */
    @Export(name="forwardProxyBypassDefaultAction", type=String.class, parameters={})
    private Output<String> forwardProxyBypassDefaultAction;

    /**
     * @return Forward proxy bypass default action. (enabled / disabled)
     * 
     */
    public Output<String> forwardProxyBypassDefaultAction() {
        return this.forwardProxyBypassDefaultAction;
    }
    /**
     * full path of the profile
     * 
     */
    @Export(name="fullPath", type=String.class, parameters={})
    private Output<String> fullPath;

    /**
     * @return full path of the profile
     * 
     */
    public Output<String> fullPath() {
        return this.fullPath;
    }
    /**
     * generation
     * 
     */
    @Export(name="generation", type=Integer.class, parameters={})
    private Output<Integer> generation;

    /**
     * @return generation
     * 
     */
    public Output<Integer> generation() {
        return this.generation;
    }
    /**
     * Generic alerts enabled / disabled.
     * 
     */
    @Export(name="genericAlert", type=String.class, parameters={})
    private Output<String> genericAlert;

    /**
     * @return Generic alerts enabled / disabled.
     * 
     */
    public Output<String> genericAlert() {
        return this.genericAlert;
    }
    /**
     * Handshake time out (seconds)
     * 
     */
    @Export(name="handshakeTimeout", type=String.class, parameters={})
    private Output<String> handshakeTimeout;

    /**
     * @return Handshake time out (seconds)
     * 
     */
    public Output<String> handshakeTimeout() {
        return this.handshakeTimeout;
    }
    /**
     * Inherit cert key chain
     * 
     */
    @Export(name="inheritCertKeychain", type=String.class, parameters={})
    private Output<String> inheritCertKeychain;

    /**
     * @return Inherit cert key chain
     * 
     */
    public Output<String> inheritCertKeychain() {
        return this.inheritCertKeychain;
    }
    /**
     * Contains a key name
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return Contains a key name
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     * 
     */
    @Export(name="modSslMethods", type=String.class, parameters={})
    private Output<String> modSslMethods;

    /**
     * @return ModSSL Methods enabled / disabled. Default is disabled.
     * 
     */
    public Output<String> modSslMethods() {
        return this.modSslMethods;
    }
    /**
     * ModSSL Methods enabled / disabled. Default is disabled.
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output<String> mode;

    /**
     * @return ModSSL Methods enabled / disabled. Default is disabled.
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Specifies the name of the profile.Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-clientssl-profile`.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Specifies the name of the profile.Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-clientssl-profile`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * name of partition
     * 
     */
    @Export(name="partition", type=String.class, parameters={})
    private Output<String> partition;

    /**
     * @return name of partition
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }
    /**
     * Client Certificate Constrained Delegation CA passphrase
     * 
     */
    @Export(name="passphrase", type=String.class, parameters={})
    private Output<String> passphrase;

    /**
     * @return Client Certificate Constrained Delegation CA passphrase
     * 
     */
    public Output<String> passphrase() {
        return this.passphrase;
    }
    /**
     * Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
     * 
     */
    @Export(name="peerCertMode", type=String.class, parameters={})
    private Output<String> peerCertMode;

    /**
     * @return Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
     * 
     */
    public Output<String> peerCertMode() {
        return this.peerCertMode;
    }
    /**
     * Proxy CA Cert
     * 
     */
    @Export(name="proxyCaCert", type=String.class, parameters={})
    private Output<String> proxyCaCert;

    /**
     * @return Proxy CA Cert
     * 
     */
    public Output<String> proxyCaCert() {
        return this.proxyCaCert;
    }
    /**
     * Proxy CA Key
     * 
     */
    @Export(name="proxyCaKey", type=String.class, parameters={})
    private Output<String> proxyCaKey;

    /**
     * @return Proxy CA Key
     * 
     */
    public Output<String> proxyCaKey() {
        return this.proxyCaKey;
    }
    /**
     * Proxy CA Passphrase
     * 
     */
    @Export(name="proxyCaPassphrase", type=String.class, parameters={})
    private Output<String> proxyCaPassphrase;

    /**
     * @return Proxy CA Passphrase
     * 
     */
    public Output<String> proxyCaPassphrase() {
        return this.proxyCaPassphrase;
    }
    /**
     * Proxy SSL enabled / disabled. Default is disabled.
     * 
     */
    @Export(name="proxySsl", type=String.class, parameters={})
    private Output<String> proxySsl;

    /**
     * @return Proxy SSL enabled / disabled. Default is disabled.
     * 
     */
    public Output<String> proxySsl() {
        return this.proxySsl;
    }
    /**
     * Proxy SSL passthrough enabled / disabled. Default is disabled.
     * 
     */
    @Export(name="proxySslPassthrough", type=String.class, parameters={})
    private Output<String> proxySslPassthrough;

    /**
     * @return Proxy SSL passthrough enabled / disabled. Default is disabled.
     * 
     */
    public Output<String> proxySslPassthrough() {
        return this.proxySslPassthrough;
    }
    /**
     * Renogotiate Period (seconds)
     * 
     */
    @Export(name="renegotiatePeriod", type=String.class, parameters={})
    private Output<String> renegotiatePeriod;

    /**
     * @return Renogotiate Period (seconds)
     * 
     */
    public Output<String> renegotiatePeriod() {
        return this.renegotiatePeriod;
    }
    /**
     * Renogotiate Size
     * 
     */
    @Export(name="renegotiateSize", type=String.class, parameters={})
    private Output<String> renegotiateSize;

    /**
     * @return Renogotiate Size
     * 
     */
    public Output<String> renegotiateSize() {
        return this.renegotiateSize;
    }
    /**
     * Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
     * 
     */
    @Export(name="renegotiation", type=String.class, parameters={})
    private Output<String> renegotiation;

    /**
     * @return Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
     * 
     */
    public Output<String> renegotiation() {
        return this.renegotiation;
    }
    /**
     * When `true`, client certificate is retained in SSL session.
     * 
     */
    @Export(name="retainCertificate", type=String.class, parameters={})
    private Output<String> retainCertificate;

    /**
     * @return When `true`, client certificate is retained in SSL session.
     * 
     */
    public Output<String> retainCertificate() {
        return this.retainCertificate;
    }
    /**
     * Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
     * When `request` is set the system request secure renegotation of SSL connections.
     * `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
     * The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
     * 
     */
    @Export(name="secureRenegotiation", type=String.class, parameters={})
    private Output<String> secureRenegotiation;

    /**
     * @return Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
     * When `request` is set the system request secure renegotation of SSL connections.
     * `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
     * The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
     * 
     */
    public Output<String> secureRenegotiation() {
        return this.secureRenegotiation;
    }
    /**
     * Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
     * 
     */
    @Export(name="serverName", type=String.class, parameters={})
    private Output<String> serverName;

    /**
     * @return Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
     * 
     */
    public Output<String> serverName() {
        return this.serverName;
    }
    /**
     * Session Mirroring (enabled / disabled)
     * 
     */
    @Export(name="sessionMirroring", type=String.class, parameters={})
    private Output<String> sessionMirroring;

    /**
     * @return Session Mirroring (enabled / disabled)
     * 
     */
    public Output<String> sessionMirroring() {
        return this.sessionMirroring;
    }
    /**
     * Session Ticket (enabled / disabled)
     * 
     */
    @Export(name="sessionTicket", type=String.class, parameters={})
    private Output<String> sessionTicket;

    /**
     * @return Session Ticket (enabled / disabled)
     * 
     */
    public Output<String> sessionTicket() {
        return this.sessionTicket;
    }
    /**
     * Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
     * There can be only one SSL profile with this setting enabled.
     * 
     */
    @Export(name="sniDefault", type=String.class, parameters={})
    private Output<String> sniDefault;

    /**
     * @return Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
     * There can be only one SSL profile with this setting enabled.
     * 
     */
    public Output<String> sniDefault() {
        return this.sniDefault;
    }
    /**
     * Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
     * 
     */
    @Export(name="sniRequire", type=String.class, parameters={})
    private Output<String> sniRequire;

    /**
     * @return Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
     * 
     */
    public Output<String> sniRequire() {
        return this.sniRequire;
    }
    /**
     * Enables or disables SSL client certificate constrained delegation. The default option is disabled. Conversely, you can specify enabled to use the SSL client certificate constrained delegation.
     * 
     */
    @Export(name="sslC3d", type=String.class, parameters={})
    private Output<String> sslC3d;

    /**
     * @return Enables or disables SSL client certificate constrained delegation. The default option is disabled. Conversely, you can specify enabled to use the SSL client certificate constrained delegation.
     * 
     */
    public Output<String> sslC3d() {
        return this.sslC3d;
    }
    /**
     * Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
     * 
     */
    @Export(name="sslForwardProxy", type=String.class, parameters={})
    private Output<String> sslForwardProxy;

    /**
     * @return Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
     * 
     */
    public Output<String> sslForwardProxy() {
        return this.sslForwardProxy;
    }
    /**
     * Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
     * 
     */
    @Export(name="sslForwardProxyBypass", type=String.class, parameters={})
    private Output<String> sslForwardProxyBypass;

    /**
     * @return Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
     * 
     */
    public Output<String> sslForwardProxyBypass() {
        return this.sslForwardProxyBypass;
    }
    /**
     * SSL sign hash (any, sha1, sha256, sha384)
     * 
     */
    @Export(name="sslSignHash", type=String.class, parameters={})
    private Output<String> sslSignHash;

    /**
     * @return SSL sign hash (any, sha1, sha256, sha384)
     * 
     */
    public Output<String> sslSignHash() {
        return this.sslSignHash;
    }
    /**
     * Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
     * 
     */
    @Export(name="strictResume", type=String.class, parameters={})
    private Output<String> strictResume;

    /**
     * @return Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
     * 
     */
    public Output<String> strictResume() {
        return this.strictResume;
    }
    /**
     * List of Enabled selection from a set of industry standard options for handling SSL processing.By default,
     * Don&#39;t insert empty fragments and No TLSv1.3 are listed as Enabled Options. `Usage` : tm_options    = [&#34;dont-insert-empty-fragments&#34;,&#34;no-tlsv1.3&#34;]
     * 
     */
    @Export(name="tmOptions", type=List.class, parameters={String.class})
    private Output<List<String>> tmOptions;

    /**
     * @return List of Enabled selection from a set of industry standard options for handling SSL processing.By default,
     * Don&#39;t insert empty fragments and No TLSv1.3 are listed as Enabled Options. `Usage` : tm_options    = [&#34;dont-insert-empty-fragments&#34;,&#34;no-tlsv1.3&#34;]
     * 
     */
    public Output<List<String>> tmOptions() {
        return this.tmOptions;
    }
    /**
     * Unclean Shutdown (enabled / disabled)
     * 
     */
    @Export(name="uncleanShutdown", type=String.class, parameters={})
    private Output<String> uncleanShutdown;

    /**
     * @return Unclean Shutdown (enabled / disabled)
     * 
     */
    public Output<String> uncleanShutdown() {
        return this.uncleanShutdown;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileClientSsl(String name) {
        this(name, ProfileClientSslArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileClientSsl(String name, ProfileClientSslArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileClientSsl(String name, ProfileClientSslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileClientSsl:ProfileClientSsl", name, args == null ? ProfileClientSslArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileClientSsl(String name, Output<String> id, @Nullable ProfileClientSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileClientSsl:ProfileClientSsl", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ProfileClientSsl get(String name, Output<String> id, @Nullable ProfileClientSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileClientSsl(name, id, state, options);
    }
}
