// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.DoArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.DoState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="f5bigip:index/do:Do")
public class Do extends com.pulumi.resources.CustomResource {
    /**
     * IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Export(name="bigipAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigipAddress;

    /**
     * @return IP Address of BIGIP Host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Output<Optional<String>> bigipAddress() {
        return Codegen.optional(this.bigipAddress);
    }
    /**
     * Password of BIGIP host to be used for this resource
     * 
     */
    @Export(name="bigipPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigipPassword;

    /**
     * @return Password of BIGIP host to be used for this resource
     * 
     */
    public Output<Optional<String>> bigipPassword() {
        return Codegen.optional(this.bigipPassword);
    }
    /**
     * Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Export(name="bigipPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigipPort;

    /**
     * @return Port number of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Output<Optional<String>> bigipPort() {
        return Codegen.optional(this.bigipPort);
    }
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     * 
     */
    @Export(name="bigipTokenAuth", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bigipTokenAuth;

    /**
     * @return Enable to use an external authentication source (LDAP, TACACS, etc)
     * 
     */
    public Output<Optional<Boolean>> bigipTokenAuth() {
        return Codegen.optional(this.bigipTokenAuth);
    }
    /**
     * UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    @Export(name="bigipUser", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bigipUser;

    /**
     * @return UserName of BIGIP host to be used for this resource,this is optional parameter.
     * whenever we specify this parameter it gets overwrite provider configuration
     * 
     */
    public Output<Optional<String>> bigipUser() {
        return Codegen.optional(this.bigipUser);
    }
    /**
     * Name of the of the Declarative DO JSON file
     * 
     */
    @Export(name="doJson", refs={String.class}, tree="[0]")
    private Output<String> doJson;

    /**
     * @return Name of the of the Declarative DO JSON file
     * 
     */
    public Output<String> doJson() {
        return this.doJson;
    }
    /**
     * unique identifier for DO resource
     * 
     * @deprecated
     * this attribute is no longer in use
     * 
     */
    @Deprecated /* this attribute is no longer in use */
    @Export(name="tenantName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tenantName;

    /**
     * @return unique identifier for DO resource
     * 
     */
    public Output<Optional<String>> tenantName() {
        return Codegen.optional(this.tenantName);
    }
    /**
     * DO json
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return DO json
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Do(java.lang.String name) {
        this(name, DoArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Do(java.lang.String name, DoArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Do(java.lang.String name, DoArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/do:Do", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Do(java.lang.String name, Output<java.lang.String> id, @Nullable DoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/do:Do", name, state, makeResourceOptions(options, id), false);
    }

    private static DoArgs makeArgs(DoArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DoArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "bigipPassword",
                "bigipTokenAuth"
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
    public static Do get(java.lang.String name, Output<java.lang.String> id, @Nullable DoState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Do(name, id, state, options);
    }
}
