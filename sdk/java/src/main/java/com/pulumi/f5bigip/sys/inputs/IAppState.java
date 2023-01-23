// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.f5bigip.sys.inputs.IAppListArgs;
import com.pulumi.f5bigip.sys.inputs.IAppMetadataArgs;
import com.pulumi.f5bigip.sys.inputs.IAppTableArgs;
import com.pulumi.f5bigip.sys.inputs.IAppVariableArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IAppState extends com.pulumi.resources.ResourceArgs {

    public static final IAppState Empty = new IAppState();

    /**
     * Address of the Iapp which needs to be Iappensed
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Address of the Iapp which needs to be Iappensed
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="devicegroup")
    private @Nullable Output<String> devicegroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> devicegroup() {
        return Optional.ofNullable(this.devicegroup);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="executeAction")
    private @Nullable Output<String> executeAction;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> executeAction() {
        return Optional.ofNullable(this.executeAction);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="inheritedDevicegroup")
    private @Nullable Output<String> inheritedDevicegroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> inheritedDevicegroup() {
        return Optional.ofNullable(this.inheritedDevicegroup);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="inheritedTrafficGroup")
    private @Nullable Output<String> inheritedTrafficGroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> inheritedTrafficGroup() {
        return Optional.ofNullable(this.inheritedTrafficGroup);
    }

    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    @Import(name="jsonfile")
    private @Nullable Output<String> jsonfile;

    /**
     * @return Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    public Optional<Output<String>> jsonfile() {
        return Optional.ofNullable(this.jsonfile);
    }

    @Import(name="lists")
    private @Nullable Output<List<IAppListArgs>> lists;

    public Optional<Output<List<IAppListArgs>>> lists() {
        return Optional.ofNullable(this.lists);
    }

    @Import(name="metadatas")
    private @Nullable Output<List<IAppMetadataArgs>> metadatas;

    public Optional<Output<List<IAppMetadataArgs>>> metadatas() {
        return Optional.ofNullable(this.metadatas);
    }

    /**
     * Name of the iApp.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the iApp.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Address of the Iapp which needs to be Iappensed
     * 
     */
    @Import(name="partition")
    private @Nullable Output<String> partition;

    /**
     * @return Address of the Iapp which needs to be Iappensed
     * 
     */
    public Optional<Output<String>> partition() {
        return Optional.ofNullable(this.partition);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="strictUpdates")
    private @Nullable Output<String> strictUpdates;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> strictUpdates() {
        return Optional.ofNullable(this.strictUpdates);
    }

    @Import(name="tables")
    private @Nullable Output<List<IAppTableArgs>> tables;

    public Optional<Output<List<IAppTableArgs>>> tables() {
        return Optional.ofNullable(this.tables);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="template")
    private @Nullable Output<String> template;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> template() {
        return Optional.ofNullable(this.template);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="templateModified")
    private @Nullable Output<String> templateModified;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> templateModified() {
        return Optional.ofNullable(this.templateModified);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="templatePrerequisiteErrors")
    private @Nullable Output<String> templatePrerequisiteErrors;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> templatePrerequisiteErrors() {
        return Optional.ofNullable(this.templatePrerequisiteErrors);
    }

    /**
     * BIG-IP password
     * 
     */
    @Import(name="trafficGroup")
    private @Nullable Output<String> trafficGroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Optional<Output<String>> trafficGroup() {
        return Optional.ofNullable(this.trafficGroup);
    }

    @Import(name="variables")
    private @Nullable Output<List<IAppVariableArgs>> variables;

    public Optional<Output<List<IAppVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private IAppState() {}

    private IAppState(IAppState $) {
        this.description = $.description;
        this.devicegroup = $.devicegroup;
        this.executeAction = $.executeAction;
        this.inheritedDevicegroup = $.inheritedDevicegroup;
        this.inheritedTrafficGroup = $.inheritedTrafficGroup;
        this.jsonfile = $.jsonfile;
        this.lists = $.lists;
        this.metadatas = $.metadatas;
        this.name = $.name;
        this.partition = $.partition;
        this.strictUpdates = $.strictUpdates;
        this.tables = $.tables;
        this.template = $.template;
        this.templateModified = $.templateModified;
        this.templatePrerequisiteErrors = $.templatePrerequisiteErrors;
        this.trafficGroup = $.trafficGroup;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IAppState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IAppState $;

        public Builder() {
            $ = new IAppState();
        }

        public Builder(IAppState defaults) {
            $ = new IAppState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Address of the Iapp which needs to be Iappensed
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Address of the Iapp which needs to be Iappensed
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param devicegroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder devicegroup(@Nullable Output<String> devicegroup) {
            $.devicegroup = devicegroup;
            return this;
        }

        /**
         * @param devicegroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder devicegroup(String devicegroup) {
            return devicegroup(Output.of(devicegroup));
        }

        /**
         * @param executeAction BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder executeAction(@Nullable Output<String> executeAction) {
            $.executeAction = executeAction;
            return this;
        }

        /**
         * @param executeAction BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder executeAction(String executeAction) {
            return executeAction(Output.of(executeAction));
        }

        /**
         * @param inheritedDevicegroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder inheritedDevicegroup(@Nullable Output<String> inheritedDevicegroup) {
            $.inheritedDevicegroup = inheritedDevicegroup;
            return this;
        }

        /**
         * @param inheritedDevicegroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder inheritedDevicegroup(String inheritedDevicegroup) {
            return inheritedDevicegroup(Output.of(inheritedDevicegroup));
        }

        /**
         * @param inheritedTrafficGroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder inheritedTrafficGroup(@Nullable Output<String> inheritedTrafficGroup) {
            $.inheritedTrafficGroup = inheritedTrafficGroup;
            return this;
        }

        /**
         * @param inheritedTrafficGroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder inheritedTrafficGroup(String inheritedTrafficGroup) {
            return inheritedTrafficGroup(Output.of(inheritedTrafficGroup));
        }

        /**
         * @param jsonfile Refer to the Json file which will be deployed on F5 BIG-IP.
         * 
         * @return builder
         * 
         */
        public Builder jsonfile(@Nullable Output<String> jsonfile) {
            $.jsonfile = jsonfile;
            return this;
        }

        /**
         * @param jsonfile Refer to the Json file which will be deployed on F5 BIG-IP.
         * 
         * @return builder
         * 
         */
        public Builder jsonfile(String jsonfile) {
            return jsonfile(Output.of(jsonfile));
        }

        public Builder lists(@Nullable Output<List<IAppListArgs>> lists) {
            $.lists = lists;
            return this;
        }

        public Builder lists(List<IAppListArgs> lists) {
            return lists(Output.of(lists));
        }

        public Builder lists(IAppListArgs... lists) {
            return lists(List.of(lists));
        }

        public Builder metadatas(@Nullable Output<List<IAppMetadataArgs>> metadatas) {
            $.metadatas = metadatas;
            return this;
        }

        public Builder metadatas(List<IAppMetadataArgs> metadatas) {
            return metadatas(Output.of(metadatas));
        }

        public Builder metadatas(IAppMetadataArgs... metadatas) {
            return metadatas(List.of(metadatas));
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the iApp.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partition Address of the Iapp which needs to be Iappensed
         * 
         * @return builder
         * 
         */
        public Builder partition(@Nullable Output<String> partition) {
            $.partition = partition;
            return this;
        }

        /**
         * @param partition Address of the Iapp which needs to be Iappensed
         * 
         * @return builder
         * 
         */
        public Builder partition(String partition) {
            return partition(Output.of(partition));
        }

        /**
         * @param strictUpdates BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder strictUpdates(@Nullable Output<String> strictUpdates) {
            $.strictUpdates = strictUpdates;
            return this;
        }

        /**
         * @param strictUpdates BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder strictUpdates(String strictUpdates) {
            return strictUpdates(Output.of(strictUpdates));
        }

        public Builder tables(@Nullable Output<List<IAppTableArgs>> tables) {
            $.tables = tables;
            return this;
        }

        public Builder tables(List<IAppTableArgs> tables) {
            return tables(Output.of(tables));
        }

        public Builder tables(IAppTableArgs... tables) {
            return tables(List.of(tables));
        }

        /**
         * @param template BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<String> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder template(String template) {
            return template(Output.of(template));
        }

        /**
         * @param templateModified BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder templateModified(@Nullable Output<String> templateModified) {
            $.templateModified = templateModified;
            return this;
        }

        /**
         * @param templateModified BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder templateModified(String templateModified) {
            return templateModified(Output.of(templateModified));
        }

        /**
         * @param templatePrerequisiteErrors BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder templatePrerequisiteErrors(@Nullable Output<String> templatePrerequisiteErrors) {
            $.templatePrerequisiteErrors = templatePrerequisiteErrors;
            return this;
        }

        /**
         * @param templatePrerequisiteErrors BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder templatePrerequisiteErrors(String templatePrerequisiteErrors) {
            return templatePrerequisiteErrors(Output.of(templatePrerequisiteErrors));
        }

        /**
         * @param trafficGroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(@Nullable Output<String> trafficGroup) {
            $.trafficGroup = trafficGroup;
            return this;
        }

        /**
         * @param trafficGroup BIG-IP password
         * 
         * @return builder
         * 
         */
        public Builder trafficGroup(String trafficGroup) {
            return trafficGroup(Output.of(trafficGroup));
        }

        public Builder variables(@Nullable Output<List<IAppVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        public Builder variables(List<IAppVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        public Builder variables(IAppVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public IAppState build() {
            return $;
        }
    }

}