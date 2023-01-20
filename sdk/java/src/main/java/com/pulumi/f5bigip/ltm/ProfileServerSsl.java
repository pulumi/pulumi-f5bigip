// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileServerSslArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileServerSslState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileServerSsl` Manages server SSL profiles on a BIG-IP
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
 * import com.pulumi.f5bigip.ltm.ProfileServerSsl;
 * import com.pulumi.f5bigip.ltm.ProfileServerSslArgs;
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
 *         var test_ServerSsl = new ProfileServerSsl(&#34;test-ServerSsl&#34;, ProfileServerSslArgs.builder()        
 *             .authenticate(&#34;always&#34;)
 *             .ciphers(&#34;DEFAULT&#34;)
 *             .defaultsFrom(&#34;/Common/serverssl&#34;)
 *             .name(&#34;/Common/test-ServerSsl&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileServerSsl:ProfileServerSsl")
public class ProfileServerSsl extends com.pulumi.resources.CustomResource {
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
     * Specifies the frequency of server authentication for an SSL session.When `once`,specifies that the system authenticates the server once for an SSL session.
     * When `always`, specifies that the system authenticates the server once for an SSL session and also upon reuse of that session.
     * 
     */
    @Export(name="authenticate", type=String.class, parameters={})
    private Output<String> authenticate;

    /**
     * @return Specifies the frequency of server authentication for an SSL session.When `once`,specifies that the system authenticates the server once for an SSL session.
     * When `always`, specifies that the system authenticates the server once for an SSL session and also upon reuse of that session.
     * 
     */
    public Output<String> authenticate() {
        return this.authenticate;
    }
    /**
     * Client certificate chain traversal depth. Default 9.
     * 
     */
    @Export(name="authenticateDepth", type=Integer.class, parameters={})
    private Output<Integer> authenticateDepth;

    /**
     * @return Client certificate chain traversal depth. Default 9.
     * 
     */
    public Output<Integer> authenticateDepth() {
        return this.authenticateDepth;
    }
    /**
     * Specifies the name of the certificate file that is used as the certification authority certificate when SSL client certificate constrained delegation is enabled. The certificate should be generated and installed by you on the system. When selecting this option, type a certificate file name.
     * 
     */
    @Export(name="c3dCaCert", type=String.class, parameters={})
    private Output</* @Nullable */ String> c3dCaCert;

    /**
     * @return Specifies the name of the certificate file that is used as the certification authority certificate when SSL client certificate constrained delegation is enabled. The certificate should be generated and installed by you on the system. When selecting this option, type a certificate file name.
     * 
     */
    public Output<Optional<String>> c3dCaCert() {
        return Codegen.optional(this.c3dCaCert);
    }
    /**
     * Specifies the name of the key file that is used as the certification authority key when SSL client certificate constrained delegation is enabled. The key should be generated and installed by you on the system. When selecting this option, type a key file name.
     * 
     */
    @Export(name="c3dCaKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> c3dCaKey;

    /**
     * @return Specifies the name of the key file that is used as the certification authority key when SSL client certificate constrained delegation is enabled. The key should be generated and installed by you on the system. When selecting this option, type a key file name.
     * 
     */
    public Output<Optional<String>> c3dCaKey() {
        return Codegen.optional(this.c3dCaKey);
    }
    /**
     * CA Passphrase. Default
     * 
     */
    @Export(name="c3dCaPassphrase", type=String.class, parameters={})
    private Output<String> c3dCaPassphrase;

    /**
     * @return CA Passphrase. Default
     * 
     */
    public Output<String> c3dCaPassphrase() {
        return this.c3dCaPassphrase;
    }
    /**
     * Certificate Extensions List. Default
     * 
     */
    @Export(name="c3dCertExtensionCustomOids", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> c3dCertExtensionCustomOids;

    /**
     * @return Certificate Extensions List. Default
     * 
     */
    public Output<Optional<List<String>>> c3dCertExtensionCustomOids() {
        return Codegen.optional(this.c3dCertExtensionCustomOids);
    }
    /**
     * Specifies the extensions of the client certificates to be included in the generated certificates using SSL client certificate constrained delegation. For example, { basic-constraints }. The default value is { basic-constraints extended-key-usage key-usage subject-alternative-name }. The extensions are:
     * 
     */
    @Export(name="c3dCertExtensionIncludes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> c3dCertExtensionIncludes;

    /**
     * @return Specifies the extensions of the client certificates to be included in the generated certificates using SSL client certificate constrained delegation. For example, { basic-constraints }. The default value is { basic-constraints extended-key-usage key-usage subject-alternative-name }. The extensions are:
     * 
     */
    public Output<Optional<List<String>>> c3dCertExtensionIncludes() {
        return Codegen.optional(this.c3dCertExtensionIncludes);
    }
    /**
     * Certificate Lifespan. Default
     * 
     */
    @Export(name="c3dCertLifespan", type=Integer.class, parameters={})
    private Output<Integer> c3dCertLifespan;

    /**
     * @return Certificate Lifespan. Default
     * 
     */
    public Output<Integer> c3dCertLifespan() {
        return this.c3dCertLifespan;
    }
    /**
     * CA Passphrase. Default enabled
     * 
     */
    @Export(name="c3dCertificateExtensions", type=String.class, parameters={})
    private Output<String> c3dCertificateExtensions;

    /**
     * @return CA Passphrase. Default enabled
     * 
     */
    public Output<String> c3dCertificateExtensions() {
        return this.c3dCertificateExtensions;
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
     * Specifies the name of the certificate that the system uses for server-side SSL processing.
     * 
     */
    @Export(name="cert", type=String.class, parameters={})
    private Output</* @Nullable */ String> cert;

    /**
     * @return Specifies the name of the certificate that the system uses for server-side SSL processing.
     * 
     */
    public Output<Optional<String>> cert() {
        return Codegen.optional(this.cert);
    }
    /**
     * Specifies the certificates-key chain to associate with the SSL profile
     * 
     */
    @Export(name="chain", type=String.class, parameters={})
    private Output</* @Nullable */ String> chain;

    /**
     * @return Specifies the certificates-key chain to associate with the SSL profile
     * 
     */
    public Output<Optional<String>> chain() {
        return Codegen.optional(this.chain);
    }
    /**
     * Specifies the cipher group for the SSL server profile. It is mutually exclusive with the argument, `ciphers`. The default value is `none`.
     * 
     */
    @Export(name="cipherGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> cipherGroup;

    /**
     * @return Specifies the cipher group for the SSL server profile. It is mutually exclusive with the argument, `ciphers`. The default value is `none`.
     * 
     */
    public Output<Optional<String>> cipherGroup() {
        return Codegen.optional(this.cipherGroup);
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
     * The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
     * 
     */
    @Export(name="defaultsFrom", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultsFrom;

    /**
     * @return The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
     * 
     */
    public Output<Optional<String>> defaultsFrom() {
        return Codegen.optional(this.defaultsFrom);
    }
    /**
     * Response if the cert is expired (drop / ignore).
     * 
     */
    @Export(name="expireCertResponseControl", type=String.class, parameters={})
    private Output<String> expireCertResponseControl;

    /**
     * @return Response if the cert is expired (drop / ignore).
     * 
     */
    public Output<String> expireCertResponseControl() {
        return this.expireCertResponseControl;
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
     * Specifies the file name of the SSL key.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output</* @Nullable */ String> key;

    /**
     * @return Specifies the file name of the SSL key.
     * 
     */
    public Output<Optional<String>> key() {
        return Codegen.optional(this.key);
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
     * Specifies the name of the profile.Name of Profile should be full path,full path is the combination of the `partition + profile name`. For example `/Common/test-serverssl-profile`.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Specifies the name of the profile.Name of Profile should be full path,full path is the combination of the `partition + profile name`. For example `/Common/test-serverssl-profile`.
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
     * Enables or disables SSL forward proxy bypass on receiving
     * handshake_failure, protocol_version or unsupported_extension alert message during the serverside SSL handshake. When enabled and there is an SSL handshake_failure, protocol_version or unsupported_extension alert during the serverside SSL handshake, SSL traffic bypasses the BIG-IP system untouched, without decryption/encryption. The default value is disabled. Conversely, you can specify enabled to use this feature.
     * 
     */
    @Export(name="sslC3d", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslC3d;

    /**
     * @return Enables or disables SSL forward proxy bypass on receiving
     * handshake_failure, protocol_version or unsupported_extension alert message during the serverside SSL handshake. When enabled and there is an SSL handshake_failure, protocol_version or unsupported_extension alert during the serverside SSL handshake, SSL traffic bypasses the BIG-IP system untouched, without decryption/encryption. The default value is disabled. Conversely, you can specify enabled to use this feature.
     * 
     */
    public Output<Optional<String>> sslC3d() {
        return Codegen.optional(this.sslC3d);
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
     * Unclean Shutdown (drop / ignore)
     * 
     */
    @Export(name="untrustedCertResponseControl", type=String.class, parameters={})
    private Output<String> untrustedCertResponseControl;

    /**
     * @return Unclean Shutdown (drop / ignore)
     * 
     */
    public Output<String> untrustedCertResponseControl() {
        return this.untrustedCertResponseControl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileServerSsl(String name) {
        this(name, ProfileServerSslArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileServerSsl(String name, ProfileServerSslArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileServerSsl(String name, ProfileServerSslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, args == null ? ProfileServerSslArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileServerSsl(String name, Output<String> id, @Nullable ProfileServerSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "passphrase"
            ))
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
    public static ProfileServerSsl get(String name, Output<String> id, @Nullable ProfileServerSslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileServerSsl(name, id, state, options);
    }
}
