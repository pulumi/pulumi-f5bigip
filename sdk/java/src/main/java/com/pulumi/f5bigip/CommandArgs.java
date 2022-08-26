// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CommandArgs extends com.pulumi.resources.ResourceArgs {

    public static final CommandArgs Empty = new CommandArgs();

    /**
     * The resulting output from the `commands` executed
     * 
     */
    @Import(name="commandResults")
    private @Nullable Output<List<String>> commandResults;

    /**
     * @return The resulting output from the `commands` executed
     * 
     */
    public Optional<Output<List<String>>> commandResults() {
        return Optional.ofNullable(this.commandResults);
    }

    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    @Import(name="commands", required=true)
    private Output<List<String>> commands;

    /**
     * @return The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    public Output<List<String>> commands() {
        return this.commands;
    }

    @Import(name="when")
    private @Nullable Output<String> when;

    public Optional<Output<String>> when() {
        return Optional.ofNullable(this.when);
    }

    private CommandArgs() {}

    private CommandArgs(CommandArgs $) {
        this.commandResults = $.commandResults;
        this.commands = $.commands;
        this.when = $.when;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CommandArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CommandArgs $;

        public Builder() {
            $ = new CommandArgs();
        }

        public Builder(CommandArgs defaults) {
            $ = new CommandArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param commandResults The resulting output from the `commands` executed
         * 
         * @return builder
         * 
         */
        public Builder commandResults(@Nullable Output<List<String>> commandResults) {
            $.commandResults = commandResults;
            return this;
        }

        /**
         * @param commandResults The resulting output from the `commands` executed
         * 
         * @return builder
         * 
         */
        public Builder commandResults(List<String> commandResults) {
            return commandResults(Output.of(commandResults));
        }

        /**
         * @param commandResults The resulting output from the `commands` executed
         * 
         * @return builder
         * 
         */
        public Builder commandResults(String... commandResults) {
            return commandResults(List.of(commandResults));
        }

        /**
         * @param commands The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
         * 
         * @return builder
         * 
         */
        public Builder commands(Output<List<String>> commands) {
            $.commands = commands;
            return this;
        }

        /**
         * @param commands The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
         * 
         * @return builder
         * 
         */
        public Builder commands(List<String> commands) {
            return commands(Output.of(commands));
        }

        /**
         * @param commands The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
         * 
         * @return builder
         * 
         */
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }

        public Builder when(@Nullable Output<String> when) {
            $.when = when;
            return this;
        }

        public Builder when(String when) {
            return when(Output.of(when));
        }

        public CommandArgs build() {
            $.commands = Objects.requireNonNull($.commands, "expected parameter 'commands' to be non-null");
            return $;
        }
    }

}
