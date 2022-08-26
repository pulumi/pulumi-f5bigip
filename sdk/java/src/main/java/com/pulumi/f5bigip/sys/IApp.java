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
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.sys.IApp;
 * import com.pulumi.f5bigip.sys.IAppArgs;
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
 *         var simplehttp = new IApp(&#34;simplehttp&#34;, IAppArgs.builder()        
 *             .name(&#34;simplehttp&#34;)
 *             .jsonfile(Files.readString(Paths.get(&#34;simplehttp.json&#34;)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Json File
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 *  * `description` - User defined description.
 *  * `deviceGroup` - The name of the device group that the application service is assigned to.
 *  * `executeAction` - Run the specified template action associated with the application.
 *  * `inheritedDevicegroup`- Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use &#39;device-group default&#39; or &#39;device-group non-default&#39; to set this.
 *  * `inheritedTrafficGroup` - Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use &#39;traffic-group default&#39; or &#39;traffic-group non-default&#39; to set this.
 *  * `partition` - Displays the administrative partition within which the application resides.
 *  * `strictUpdates` - Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system&#39;s application management interfaces.
 *  * `template` - The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
 *  * `templateModified` - Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
 *  * `templatePrerequisiteErrors` - Indicates any missing prerequisites associated with the template that defines this application.
 *  * `trafficGroup` - The name of the traffic group that the application service is assigned to.
 *  * `lists` - string values
 *  * `metadata` - User defined generic data for the application service. It is a name and value pair.
 *  * `tables` - Values provided like pool name, nodes etc.
 *  * `variables` - Name, values, encrypted or not
 * 
 */
@ResourceType(type="f5bigip:sys/iApp:IApp")
public class IApp extends com.pulumi.resources.CustomResource {
    /**
     * Address of the Iapp which needs to be Iappensed
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Address of the Iapp which needs to be Iappensed
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="devicegroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> devicegroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> devicegroup() {
        return Codegen.optional(this.devicegroup);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="executeAction", type=String.class, parameters={})
    private Output</* @Nullable */ String> executeAction;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> executeAction() {
        return Codegen.optional(this.executeAction);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="inheritedDevicegroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> inheritedDevicegroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> inheritedDevicegroup() {
        return Codegen.optional(this.inheritedDevicegroup);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="inheritedTrafficGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> inheritedTrafficGroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> inheritedTrafficGroup() {
        return Codegen.optional(this.inheritedTrafficGroup);
    }
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    @Export(name="jsonfile", type=String.class, parameters={})
    private Output</* @Nullable */ String> jsonfile;

    /**
     * @return Refer to the Json file which will be deployed on F5 BIG-IP.
     * 
     */
    public Output<Optional<String>> jsonfile() {
        return Codegen.optional(this.jsonfile);
    }
    @Export(name="lists", type=List.class, parameters={IAppList.class})
    private Output</* @Nullable */ List<IAppList>> lists;

    public Output<Optional<List<IAppList>>> lists() {
        return Codegen.optional(this.lists);
    }
    @Export(name="metadatas", type=List.class, parameters={IAppMetadata.class})
    private Output</* @Nullable */ List<IAppMetadata>> metadatas;

    public Output<Optional<List<IAppMetadata>>> metadatas() {
        return Codegen.optional(this.metadatas);
    }
    /**
     * Name of the iApp.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output</* @Nullable */ String> name;

    /**
     * @return Name of the iApp.
     * 
     */
    public Output<Optional<String>> name() {
        return Codegen.optional(this.name);
    }
    /**
     * Address of the Iapp which needs to be Iappensed
     * 
     */
    @Export(name="partition", type=String.class, parameters={})
    private Output</* @Nullable */ String> partition;

    /**
     * @return Address of the Iapp which needs to be Iappensed
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="strictUpdates", type=String.class, parameters={})
    private Output</* @Nullable */ String> strictUpdates;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> strictUpdates() {
        return Codegen.optional(this.strictUpdates);
    }
    @Export(name="tables", type=List.class, parameters={IAppTable.class})
    private Output</* @Nullable */ List<IAppTable>> tables;

    public Output<Optional<List<IAppTable>>> tables() {
        return Codegen.optional(this.tables);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="template", type=String.class, parameters={})
    private Output</* @Nullable */ String> template;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> template() {
        return Codegen.optional(this.template);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="templateModified", type=String.class, parameters={})
    private Output</* @Nullable */ String> templateModified;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> templateModified() {
        return Codegen.optional(this.templateModified);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="templatePrerequisiteErrors", type=String.class, parameters={})
    private Output</* @Nullable */ String> templatePrerequisiteErrors;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> templatePrerequisiteErrors() {
        return Codegen.optional(this.templatePrerequisiteErrors);
    }
    /**
     * BIG-IP password
     * 
     */
    @Export(name="trafficGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> trafficGroup;

    /**
     * @return BIG-IP password
     * 
     */
    public Output<Optional<String>> trafficGroup() {
        return Codegen.optional(this.trafficGroup);
    }
    @Export(name="variables", type=List.class, parameters={IAppVariable.class})
    private Output</* @Nullable */ List<IAppVariable>> variables;

    public Output<Optional<List<IAppVariable>>> variables() {
        return Codegen.optional(this.variables);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IApp(String name) {
        this(name, IAppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IApp(String name, @Nullable IAppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IApp(String name, @Nullable IAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/iApp:IApp", name, args == null ? IAppArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IApp(String name, Output<String> id, @Nullable IAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:sys/iApp:IApp", name, state, makeResourceOptions(options, id));
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
    public static IApp get(String name, Output<String> id, @Nullable IAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IApp(name, id, state, options);
    }
}
