// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CommandState extends com.pulumi.resources.ResourceArgs {

    public static final CommandState Empty = new CommandState();

    /**
     * The resulting output from the `commands` executed.
     * 
     */
    @Import(name="commandResults")
    private @Nullable Output<List<String>> commandResults;

    /**
     * @return The resulting output from the `commands` executed.
     * 
     */
    public Optional<Output<List<String>>> commandResults() {
        return Optional.ofNullable(this.commandResults);
    }

    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    @Import(name="commands")
    private @Nullable Output<List<String>> commands;

    /**
     * @return The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
     * 
     */
    public Optional<Output<List<String>>> commands() {
        return Optional.ofNullable(this.commands);
    }

    @Import(name="when")
    private @Nullable Output<String> when;

    public Optional<Output<String>> when() {
        return Optional.ofNullable(this.when);
    }

    private CommandState() {}

    private CommandState(CommandState $) {
        this.commandResults = $.commandResults;
        this.commands = $.commands;
        this.when = $.when;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CommandState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CommandState $;

        public Builder() {
            $ = new CommandState();
        }

        public Builder(CommandState defaults) {
            $ = new CommandState(Objects.requireNonNull(defaults));
        }

        /**
         * @param commandResults The resulting output from the `commands` executed.
         * 
         * @return builder
         * 
         */
        public Builder commandResults(@Nullable Output<List<String>> commandResults) {
            $.commandResults = commandResults;
            return this;
        }

        /**
         * @param commandResults The resulting output from the `commands` executed.
         * 
         * @return builder
         * 
         */
        public Builder commandResults(List<String> commandResults) {
            return commandResults(Output.of(commandResults));
        }

        /**
         * @param commandResults The resulting output from the `commands` executed.
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
        public Builder commands(@Nullable Output<List<String>> commands) {
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

        public CommandState build() {
            return $;
        }
    }

}
