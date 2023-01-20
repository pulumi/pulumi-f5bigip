// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.NetIkePeerArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.NetIkePeerState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.NetIkePeer` Manages a ike_peer configuration
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.NetIkePeer;
 * import com.pulumi.f5bigip.NetIkePeerArgs;
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
 *         var example1 = new NetIkePeer(&#34;example1&#34;, NetIkePeerArgs.builder()        
 *             .localAddress(&#34;192.16.81.240&#34;)
 *             .name(&#34;example1&#34;)
 *             .profile(&#34;/Common/dslite&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:index/netIkePeer:NetIkePeer")
public class NetIkePeer extends com.pulumi.resources.CustomResource {
    /**
     * The application service that the object belongs to
     * 
     */
    @Export(name="appService", type=String.class, parameters={})
    private Output</* @Nullable */ String> appService;

    /**
     * @return The application service that the object belongs to
     * 
     */
    public Output<Optional<String>> appService() {
        return Codegen.optional(this.appService);
    }
    /**
     * the trusted root and intermediate certificate authorities
     * 
     */
    @Export(name="caCertFile", type=String.class, parameters={})
    private Output<String> caCertFile;

    /**
     * @return the trusted root and intermediate certificate authorities
     * 
     */
    public Output<String> caCertFile() {
        return this.caCertFile;
    }
    /**
     * Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     * 
     */
    @Export(name="crlFile", type=String.class, parameters={})
    private Output<String> crlFile;

    /**
     * @return Specifies the file name of the Certificate Revocation List. Only supported in IKEv1
     * 
     */
    public Output<String> crlFile() {
        return this.crlFile;
    }
    /**
     * User defined description
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return User defined description
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Specifies the number of seconds between Dead Peer Detection messages
     * 
     */
    @Export(name="dpdDelay", type=Integer.class, parameters={})
    private Output<Integer> dpdDelay;

    /**
     * @return Specifies the number of seconds between Dead Peer Detection messages
     * 
     */
    public Output<Integer> dpdDelay() {
        return this.dpdDelay;
    }
    /**
     * Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     * 
     */
    @Export(name="generatePolicy", type=String.class, parameters={})
    private Output<String> generatePolicy;

    /**
     * @return Enable or disable the generation of Security Policy Database entries(SPD) when the device is the responder of the IKE remote node
     * 
     */
    public Output<String> generatePolicy() {
        return this.generatePolicy;
    }
    /**
     * Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     * 
     */
    @Export(name="lifetime", type=Integer.class, parameters={})
    private Output<Integer> lifetime;

    /**
     * @return Defines the lifetime in minutes of an IKE SA which will be proposed in the phase 1 negotiations
     * 
     */
    public Output<Integer> lifetime() {
        return this.lifetime;
    }
    /**
     * Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     * 
     */
    @Export(name="mode", type=String.class, parameters={})
    private Output<String> mode;

    /**
     * @return Defines the exchange mode for phase 1 when racoon is the initiator, or the acceptable exchange mode when racoon is the responder
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Specifies the name of the certificate file object
     * 
     */
    @Export(name="myCertFile", type=String.class, parameters={})
    private Output<String> myCertFile;

    /**
     * @return Specifies the name of the certificate file object
     * 
     */
    public Output<String> myCertFile() {
        return this.myCertFile;
    }
    /**
     * Specifies the name of the certificate key file object
     * 
     */
    @Export(name="myCertKeyFile", type=String.class, parameters={})
    private Output<String> myCertKeyFile;

    /**
     * @return Specifies the name of the certificate key file object
     * 
     */
    public Output<String> myCertKeyFile() {
        return this.myCertKeyFile;
    }
    /**
     * Specifies the passphrase of the key used for my-cert-key-file
     * 
     */
    @Export(name="myCertKeyPassphrase", type=String.class, parameters={})
    private Output<String> myCertKeyPassphrase;

    /**
     * @return Specifies the passphrase of the key used for my-cert-key-file
     * 
     */
    public Output<String> myCertKeyPassphrase() {
        return this.myCertKeyPassphrase;
    }
    /**
     * Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     * 
     */
    @Export(name="myIdType", type=String.class, parameters={})
    private Output<String> myIdType;

    /**
     * @return Specifies the identifier type sent to the remote host to use in the phase 1 negotiation
     * 
     */
    public Output<String> myIdType() {
        return this.myIdType;
    }
    /**
     * Specifies the identifier value sent to the remote host in the phase 1 negotiation
     * 
     */
    @Export(name="myIdValue", type=String.class, parameters={})
    private Output<String> myIdValue;

    /**
     * @return Specifies the identifier value sent to the remote host in the phase 1 negotiation
     * 
     */
    public Output<String> myIdValue() {
        return this.myIdValue;
    }
    /**
     * Name of the ike_peer
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the ike_peer
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Enables use of the NAT-Traversal IPsec extension
     * 
     */
    @Export(name="natTraversal", type=String.class, parameters={})
    private Output<String> natTraversal;

    /**
     * @return Enables use of the NAT-Traversal IPsec extension
     * 
     */
    public Output<String> natTraversal() {
        return this.natTraversal;
    }
    /**
     * Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     * 
     */
    @Export(name="passive", type=String.class, parameters={})
    private Output<String> passive;

    /**
     * @return Specifies whether the local IKE agent can be the initiator of the IKE negotiation with this ike-peer
     * 
     */
    public Output<String> passive() {
        return this.passive;
    }
    /**
     * Specifies the peer’s certificate for authentication
     * 
     */
    @Export(name="peersCertFile", type=String.class, parameters={})
    private Output<String> peersCertFile;

    /**
     * @return Specifies the peer’s certificate for authentication
     * 
     */
    public Output<String> peersCertFile() {
        return this.peersCertFile;
    }
    /**
     * Specifies that the only peers-cert-type supported is certfile
     * 
     */
    @Export(name="peersCertType", type=String.class, parameters={})
    private Output<String> peersCertType;

    /**
     * @return Specifies that the only peers-cert-type supported is certfile
     * 
     */
    public Output<String> peersCertType() {
        return this.peersCertType;
    }
    /**
     * Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     * 
     */
    @Export(name="peersIdType", type=String.class, parameters={})
    private Output<String> peersIdType;

    /**
     * @return Specifies which of address, fqdn, asn1dn, user-fqdn or keyid-tag types to use as peers-id-type
     * 
     */
    public Output<String> peersIdType() {
        return this.peersIdType;
    }
    /**
     * Specifies the peer’s identifier to be received
     * 
     */
    @Export(name="peersIdValue", type=String.class, parameters={})
    private Output<String> peersIdValue;

    /**
     * @return Specifies the peer’s identifier to be received
     * 
     */
    public Output<String> peersIdValue() {
        return this.peersIdValue;
    }
    /**
     * Specifies the authentication method used for phase 1 negotiation
     * 
     */
    @Export(name="phase1AuthMethod", type=String.class, parameters={})
    private Output<String> phase1AuthMethod;

    /**
     * @return Specifies the authentication method used for phase 1 negotiation
     * 
     */
    public Output<String> phase1AuthMethod() {
        return this.phase1AuthMethod;
    }
    /**
     * Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     * 
     */
    @Export(name="phase1EncryptAlgorithm", type=String.class, parameters={})
    private Output<String> phase1EncryptAlgorithm;

    /**
     * @return Specifies the encryption algorithm used for the isakmp phase 1 negotiation
     * 
     */
    public Output<String> phase1EncryptAlgorithm() {
        return this.phase1EncryptAlgorithm;
    }
    /**
     * Defines the hash algorithm used for the isakmp phase 1 negotiation
     * 
     */
    @Export(name="phase1HashAlgorithm", type=String.class, parameters={})
    private Output<String> phase1HashAlgorithm;

    /**
     * @return Defines the hash algorithm used for the isakmp phase 1 negotiation
     * 
     */
    public Output<String> phase1HashAlgorithm() {
        return this.phase1HashAlgorithm;
    }
    /**
     * Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     * 
     */
    @Export(name="phase1PerfectForwardSecrecy", type=String.class, parameters={})
    private Output<String> phase1PerfectForwardSecrecy;

    /**
     * @return Defines the Diffie-Hellman group for key exchange to provide perfect forward secrecy
     * 
     */
    public Output<String> phase1PerfectForwardSecrecy() {
        return this.phase1PerfectForwardSecrecy;
    }
    /**
     * Specifies the preshared key for ISAKMP SAs
     * 
     */
    @Export(name="presharedKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> presharedKey;

    /**
     * @return Specifies the preshared key for ISAKMP SAs
     * 
     */
    public Output<Optional<String>> presharedKey() {
        return Codegen.optional(this.presharedKey);
    }
    /**
     * Display the encrypted preshared-key for the IKE remote node
     * 
     */
    @Export(name="presharedKeyEncrypted", type=String.class, parameters={})
    private Output<String> presharedKeyEncrypted;

    /**
     * @return Display the encrypted preshared-key for the IKE remote node
     * 
     */
    public Output<String> presharedKeyEncrypted() {
        return this.presharedKeyEncrypted;
    }
    /**
     * Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     * 
     */
    @Export(name="prf", type=String.class, parameters={})
    private Output<String> prf;

    /**
     * @return Specifies the pseudo-random function used to derive keying material for all cryptographic operations
     * 
     */
    public Output<String> prf() {
        return this.prf;
    }
    /**
     * If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     * 
     */
    @Export(name="proxySupport", type=String.class, parameters={})
    private Output<String> proxySupport;

    /**
     * @return If this value is enabled, both values of ID payloads in the phase 2 exchange are used as the addresses of end-point of IPsec-SAs
     * 
     */
    public Output<String> proxySupport() {
        return this.proxySupport;
    }
    /**
     * Specifies the IP address of the IKE remote node
     * 
     */
    @Export(name="remoteAddress", type=String.class, parameters={})
    private Output<String> remoteAddress;

    /**
     * @return Specifies the IP address of the IKE remote node
     * 
     */
    public Output<String> remoteAddress() {
        return this.remoteAddress;
    }
    /**
     * Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     * 
     */
    @Export(name="replayWindowSize", type=Integer.class, parameters={})
    private Output<Integer> replayWindowSize;

    /**
     * @return Specifies the replay window size of the IPsec SAs negotiated with the IKE remote node
     * 
     */
    public Output<Integer> replayWindowSize() {
        return this.replayWindowSize;
    }
    /**
     * Enables or disables this IKE remote node
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return Enables or disables this IKE remote node
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Specifies the names of the traffic-selector objects associated with this ike-peer
     * 
     */
    @Export(name="trafficSelectors", type=List.class, parameters={String.class})
    private Output<List<String>> trafficSelectors;

    /**
     * @return Specifies the names of the traffic-selector objects associated with this ike-peer
     * 
     */
    public Output<List<String>> trafficSelectors() {
        return this.trafficSelectors;
    }
    /**
     * Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     * 
     */
    @Export(name="verifyCert", type=String.class, parameters={})
    private Output<String> verifyCert;

    /**
     * @return Specifies whether to verify the certificate chain of the remote peer based on the trusted certificates in ca-cert-file
     * 
     */
    public Output<String> verifyCert() {
        return this.verifyCert;
    }
    /**
     * Specifies which version of IKE to be used
     * 
     */
    @Export(name="versions", type=List.class, parameters={String.class})
    private Output<List<String>> versions;

    /**
     * @return Specifies which version of IKE to be used
     * 
     */
    public Output<List<String>> versions() {
        return this.versions;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetIkePeer(String name) {
        this(name, NetIkePeerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetIkePeer(String name, NetIkePeerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetIkePeer(String name, NetIkePeerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/netIkePeer:NetIkePeer", name, args == null ? NetIkePeerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetIkePeer(String name, Output<String> id, @Nullable NetIkePeerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/netIkePeer:NetIkePeer", name, state, makeResourceOptions(options, id));
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
    public static NetIkePeer get(String name, Output<String> id, @Nullable NetIkePeerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetIkePeer(name, id, state, options);
    }
}
