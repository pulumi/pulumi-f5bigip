// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.CommandArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.CommandState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.Command` Run TMSH commands on F5 devices
 * 
 * This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device
 * 
 */
@ResourceType(type="f5bigip:index/command:Command")
public class Command extends com.pulumi.resources.CustomResource {
    /**
     * The resulting output from the `commands` executed.
     * 
     */
    @Export(name="commandResults", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> commandResults;

    /**
     * @return The resulting output from the `commands` executed.
     * 
     */
    public Output<List<String>> commandResults() {
        return this.commandResults;
    }
    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    @Export(name="commands", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> commands;

    /**
     * @return The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    public Output<List<String>> commands() {
        return this.commands;
    }
    @Export(name="when", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> when;

    public Output<Optional<String>> when() {
        return Codegen.optional(this.when);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Command(java.lang.String name) {
        this(name, CommandArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Command(java.lang.String name, CommandArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Command(java.lang.String name, CommandArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/command:Command", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Command(java.lang.String name, Output<java.lang.String> id, @Nullable CommandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/command:Command", name, state, makeResourceOptions(options, id), false);
    }

    private static CommandArgs makeArgs(CommandArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CommandArgs.Empty : args;
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
    public static Command get(java.lang.String name, Output<java.lang.String> id, @Nullable CommandState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Command(name, id, state, options);
    }
}
