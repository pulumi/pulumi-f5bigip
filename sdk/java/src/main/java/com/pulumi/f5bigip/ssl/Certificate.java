// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ssl;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ssl.CertificateArgs;
import com.pulumi.f5bigip.ssl.inputs.CertificateState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ssl.Certificate` This resource will import SSL certificates on BIG-IP LTM.
 * Certificates can be imported from certificate files on the local disk, in PEM format
 * 
 */
@ResourceType(type="f5bigip:ssl/certificate:Certificate")
public class Certificate extends com.pulumi.resources.CustomResource {
    /**
     * Content of certificate on Disk
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return Content of certificate on Disk
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * Full Path Name of ssl certificate
     * 
     */
    @Export(name="fullPath", refs={String.class}, tree="[0]")
    private Output<String> fullPath;

    /**
     * @return Full Path Name of ssl certificate
     * 
     */
    public Output<String> fullPath() {
        return this.fullPath;
    }
    /**
     * Specifies the issuer certificate.
     * 
     */
    @Export(name="issuerCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issuerCert;

    /**
     * @return Specifies the issuer certificate.
     * 
     */
    public Output<Optional<String>> issuerCert() {
        return Codegen.optional(this.issuerCert);
    }
    /**
     * Specifies the type of monitoring used.
     * 
     */
    @Export(name="monitoringType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> monitoringType;

    /**
     * @return Specifies the type of monitoring used.
     * 
     */
    public Output<Optional<String>> monitoringType() {
        return Codegen.optional(this.monitoringType);
    }
    /**
     * Name of the SSL Certificate to be Imported on to BIGIP
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the SSL Certificate to be Imported on to BIGIP
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the OCSP responder.
     * 
     */
    @Export(name="ocsp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ocsp;

    /**
     * @return Specifies the OCSP responder.
     * 
     */
    public Output<Optional<String>> ocsp() {
        return Codegen.optional(this.ocsp);
    }
    /**
     * Partition of ssl certificate
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Partition of ssl certificate
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Certificate(String name) {
        this(name, CertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Certificate(String name, CertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Certificate(String name, CertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ssl/certificate:Certificate", name, args == null ? CertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Certificate(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ssl/certificate:Certificate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "content"
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
    public static Certificate get(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Certificate(name, id, state, options);
    }
}
