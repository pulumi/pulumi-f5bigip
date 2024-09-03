// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class As3Args extends com.pulumi.resources.ResourceArgs {

    public static final As3Args Empty = new As3Args();

    /**
     * List of applications currently deployed on the Big-Ip
     * 
     */
    @Import(name="applicationList")
    private @Nullable Output<String> applicationList;

    /**
     * @return List of applications currently deployed on the Big-Ip
     * 
     */
    public Optional<Output<String>> applicationList() {
        return Optional.ofNullable(this.applicationList);
    }

    /**
     * Path/Filename of Declarative AS3 JSON which is a json file used with builtin `file` function
     * 
     */
    @Import(name="as3Json")
    private @Nullable Output<String> as3Json;

    /**
     * @return Path/Filename of Declarative AS3 JSON which is a json file used with builtin `file` function
     * 
     */
    public Optional<Output<String>> as3Json() {
        return Optional.ofNullable(this.as3Json);
    }

    /**
     * Set True if you want to ignore metadata changes during update. By default it is set to false
     * 
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     * 
     * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
     * 
     */
    @Import(name="ignoreMetadata")
    private @Nullable Output<Boolean> ignoreMetadata;

    /**
     * @return Set True if you want to ignore metadata changes during update. By default it is set to false
     * 
     * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
     * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
     * 
     * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
     * 
     */
    public Optional<Output<Boolean>> ignoreMetadata() {
        return Optional.ofNullable(this.ignoreMetadata);
    }

    /**
     * ID of AS3 post declaration async task
     * 
     */
    @Import(name="taskId")
    private @Nullable Output<String> taskId;

    /**
     * @return ID of AS3 post declaration async task
     * 
     */
    public Optional<Output<String>> taskId() {
        return Optional.ofNullable(this.taskId);
    }

    /**
     * If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     * 
     */
    @Import(name="tenantFilter")
    private @Nullable Output<String> tenantFilter;

    /**
     * @return If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
     * 
     */
    public Optional<Output<String>> tenantFilter() {
        return Optional.ofNullable(this.tenantFilter);
    }

    /**
     * List of tenants currently deployed on the Big-Ip
     * 
     */
    @Import(name="tenantList")
    private @Nullable Output<String> tenantList;

    /**
     * @return List of tenants currently deployed on the Big-Ip
     * 
     */
    public Optional<Output<String>> tenantList() {
        return Optional.ofNullable(this.tenantList);
    }

    /**
     * Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
     * 
     */
    @Import(name="tenantName")
    private @Nullable Output<String> tenantName;

    /**
     * @return Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
     * 
     */
    public Optional<Output<String>> tenantName() {
        return Optional.ofNullable(this.tenantName);
    }

    private As3Args() {}

    private As3Args(As3Args $) {
        this.applicationList = $.applicationList;
        this.as3Json = $.as3Json;
        this.ignoreMetadata = $.ignoreMetadata;
        this.taskId = $.taskId;
        this.tenantFilter = $.tenantFilter;
        this.tenantList = $.tenantList;
        this.tenantName = $.tenantName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(As3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private As3Args $;

        public Builder() {
            $ = new As3Args();
        }

        public Builder(As3Args defaults) {
            $ = new As3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationList List of applications currently deployed on the Big-Ip
         * 
         * @return builder
         * 
         */
        public Builder applicationList(@Nullable Output<String> applicationList) {
            $.applicationList = applicationList;
            return this;
        }

        /**
         * @param applicationList List of applications currently deployed on the Big-Ip
         * 
         * @return builder
         * 
         */
        public Builder applicationList(String applicationList) {
            return applicationList(Output.of(applicationList));
        }

        /**
         * @param as3Json Path/Filename of Declarative AS3 JSON which is a json file used with builtin `file` function
         * 
         * @return builder
         * 
         */
        public Builder as3Json(@Nullable Output<String> as3Json) {
            $.as3Json = as3Json;
            return this;
        }

        /**
         * @param as3Json Path/Filename of Declarative AS3 JSON which is a json file used with builtin `file` function
         * 
         * @return builder
         * 
         */
        public Builder as3Json(String as3Json) {
            return as3Json(Output.of(as3Json));
        }

        /**
         * @param ignoreMetadata Set True if you want to ignore metadata changes during update. By default it is set to false
         * 
         * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
         * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
         * 
         * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
         * 
         * @return builder
         * 
         */
        public Builder ignoreMetadata(@Nullable Output<Boolean> ignoreMetadata) {
            $.ignoreMetadata = ignoreMetadata;
            return this;
        }

        /**
         * @param ignoreMetadata Set True if you want to ignore metadata changes during update. By default it is set to false
         * 
         * * `as3_example1.json` - Example  AS3 Declarative JSON file with single tenant
         * * `as3_example2.json` - Example  AS3 Declarative JSON file with multiple tenants
         * 
         * * `perApplication_example` - Per Application Example - JSON file with multiple Applications (and no Tenant Details)
         * 
         * @return builder
         * 
         */
        public Builder ignoreMetadata(Boolean ignoreMetadata) {
            return ignoreMetadata(Output.of(ignoreMetadata));
        }

        /**
         * @param taskId ID of AS3 post declaration async task
         * 
         * @return builder
         * 
         */
        public Builder taskId(@Nullable Output<String> taskId) {
            $.taskId = taskId;
            return this;
        }

        /**
         * @param taskId ID of AS3 post declaration async task
         * 
         * @return builder
         * 
         */
        public Builder taskId(String taskId) {
            return taskId(Output.of(taskId));
        }

        /**
         * @param tenantFilter If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
         * 
         * @return builder
         * 
         */
        public Builder tenantFilter(@Nullable Output<String> tenantFilter) {
            $.tenantFilter = tenantFilter;
            return this;
        }

        /**
         * @param tenantFilter If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
         * 
         * @return builder
         * 
         */
        public Builder tenantFilter(String tenantFilter) {
            return tenantFilter(Output.of(tenantFilter));
        }

        /**
         * @param tenantList List of tenants currently deployed on the Big-Ip
         * 
         * @return builder
         * 
         */
        public Builder tenantList(@Nullable Output<String> tenantList) {
            $.tenantList = tenantList;
            return this;
        }

        /**
         * @param tenantList List of tenants currently deployed on the Big-Ip
         * 
         * @return builder
         * 
         */
        public Builder tenantList(String tenantList) {
            return tenantList(Output.of(tenantList));
        }

        /**
         * @param tenantName Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
         * 
         * @return builder
         * 
         */
        public Builder tenantName(@Nullable Output<String> tenantName) {
            $.tenantName = tenantName;
            return this;
        }

        /**
         * @param tenantName Name of Tenant. This name is used only in the case of Per-Application Deployment. If it is not provided, then a random name would be generated.
         * 
         * @return builder
         * 
         */
        public Builder tenantName(String tenantName) {
            return tenantName(Output.of(tenantName));
        }

        public As3Args build() {
            return $;
        }
    }

}
