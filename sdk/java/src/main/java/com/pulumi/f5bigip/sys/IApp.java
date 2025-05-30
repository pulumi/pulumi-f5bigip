// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.sys;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.sys.IAppArgs;
import com.pulumi.f5bigip.sys.inputs.IAppState;
import com.pulumi.f5bigip.sys.outputs.IAppList;
import com.pulumi.f5bigip.sys.outputs.IAppMetadata;
import com.pulumi.f5bigip.sys.outputs.IAppTable;
import com.pulumi.f5bigip.sys.outputs.IAppVariable;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.sys.IApp;
 * import com.pulumi.f5bigip.sys.IAppArgs;
 * import com.pulumi.std.StdFunctions;
 * import com.pulumi.std.inputs.FileArgs;
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
 *         var simplehttp = new IApp("simplehttp", IAppArgs.builder()
 *             .name("simplehttp")
 *             .jsonfile(StdFunctions.file(FileArgs.builder()
 *                 .input("simplehttp.json")
 *                 .build()).result())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Json File
 * 
 */
@ResourceType(type="f5bigip:sys/iApp:IApp")
public class IApp extends com.pulumi.resources.CustomResource {
    /**
     * User defined description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User defined description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="devicegroup", refs={String.class}, tree="[0]")
    private Output<String> devicegroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<String> devicegroup() {
        return this.devicegroup;
    }
    /**
     * Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `execute_action` attribute take precedence over `json` value
     * 
     */
    @Export(name="executeAction", refs={String.class}, tree="[0]")
    private Output<String> executeAction;

    /**
     * @return Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `execute_action` attribute take precedence over `json` value
     * 
     */
    public Output<String> executeAction() {
        return this.executeAction;
    }
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use &#39;device-group default&#39; or &#39;device-group non-default&#39; to set this.
     * 
     */
    @Export(name="inheritedDevicegroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inheritedDevicegroup;

    /**
     * @return Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use &#39;device-group default&#39; or &#39;device-group non-default&#39; to set this.
     * 
     */
    public Output<Optional<String>> inheritedDevicegroup() {
        return Codegen.optional(this.inheritedDevicegroup);
    }
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use &#39;traffic-group default&#39; or &#39;traffic-group non-default&#39; to set this.
     * 
     */
    @Export(name="inheritedTrafficGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inheritedTrafficGroup;

    /**
     * @return Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use &#39;traffic-group default&#39; or &#39;traffic-group non-default&#39; to set this.
     * 
     */
    public Output<Optional<String>> inheritedTrafficGroup() {
        return Codegen.optional(this.inheritedTrafficGroup);
    }
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    @Export(name="jsonfile", refs={String.class}, tree="[0]")
    private Output<String> jsonfile;

    /**
     * @return Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    public Output<String> jsonfile() {
        return this.jsonfile;
    }
    /**
     * string values
     * 
     */
    @Export(name="lists", refs={List.class,IAppList.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IAppList>> lists;

    /**
     * @return string values
     * 
     */
    public Output<Optional<List<IAppList>>> lists() {
        return Codegen.optional(this.lists);
    }
    /**
     * User defined generic data for the application service. It is a name and value pair.
     * 
     */
    @Export(name="metadatas", refs={List.class,IAppMetadata.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IAppMetadata>> metadatas;

    /**
     * @return User defined generic data for the application service. It is a name and value pair.
     * 
     */
    public Output<Optional<List<IAppMetadata>>> metadatas() {
        return Codegen.optional(this.metadatas);
    }
    /**
     * Name of the iApp.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the iApp.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Displays the administrative partition within which the application resides.
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Displays the administrative partition within which the application resides.
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system&#39;s application management interfaces.
     * 
     */
    @Export(name="strictUpdates", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> strictUpdates;

    /**
     * @return Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system&#39;s application management interfaces.
     * 
     */
    public Output<Optional<String>> strictUpdates() {
        return Codegen.optional(this.strictUpdates);
    }
    @Export(name="tables", refs={List.class,IAppTable.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IAppTable>> tables;

    public Output<Optional<List<IAppTable>>> tables() {
        return Codegen.optional(this.tables);
    }
    /**
     * The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
     * 
     */
    @Export(name="template", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> template;

    /**
     * @return The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
     * 
     */
    public Output<Optional<String>> template() {
        return Codegen.optional(this.template);
    }
    /**
     * Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
     * 
     */
    @Export(name="templateModified", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateModified;

    /**
     * @return Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
     * 
     */
    public Output<Optional<String>> templateModified() {
        return Codegen.optional(this.templateModified);
    }
    /**
     * Indicates any missing prerequisites associated with the template that defines this application.
     * 
     */
    @Export(name="templatePrerequisiteErrors", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templatePrerequisiteErrors;

    /**
     * @return Indicates any missing prerequisites associated with the template that defines this application.
     * 
     */
    public Output<Optional<String>> templatePrerequisiteErrors() {
        return Codegen.optional(this.templatePrerequisiteErrors);
    }
    /**
     * The name of the traffic group that the application service is assigned to.
     * 
     */
    @Export(name="trafficGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trafficGroup;

    /**
     * @return The name of the traffic group that the application service is assigned to.
     * 
     */
    public Output<Optional<String>> trafficGroup() {
        return Codegen.optional(this.trafficGroup);
    }
    @Export(name="variables", refs={List.class,IAppVariable.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IAppVariable>> variables;

    public Output<Optional<List<IAppVariable>>> variables() {
        return Codegen.optional(this.variables);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IApp(java.lang.String name) {
        this(name, IAppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IApp(java.lang.String name, IAppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IApp(java.lang.String name, IAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/iApp:IApp", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IApp(java.lang.String name, Output<java.lang.String> id, @Nullable IAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/iApp:IApp", name, state, makeResourceOptions(options, id), false);
    }

    private static IAppArgs makeArgs(IAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IAppArgs.Empty : args;
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
    public static IApp get(java.lang.String name, Output<java.lang.String> id, @Nullable IAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IApp(name, id, state, options);
    }
}
