// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.RequestLogProfileArgs;
import com.pulumi.f5bigip.ltm.inputs.RequestLogProfileState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.RequestLogProfile` Resource used for Configures request logging using the Request Logging profile
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
 * import com.pulumi.f5bigip.ltm.RequestLogProfile;
 * import com.pulumi.f5bigip.ltm.RequestLogProfileArgs;
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
 *         var request_log_profile_tc1_child = new RequestLogProfile("request-log-profile-tc1-child", RequestLogProfileArgs.builder()
 *             .name("/Common/request-log-profile-tc1-child")
 *             .defaultsFrom(request_log_profile_tc1.name())
 *             .requestLogging("disabled")
 *             .requestlogPool("/Common/pool2")
 *             .requestlogErrorPool("/Common/pool1")
 *             .requestlogProtocol("mds-tcp")
 *             .requestlogErrorProtocol("mds-tcp")
 *             .responselogProtocol("mds-tcp")
 *             .responselogErrorProtocol("mds-tcp")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * BIG-IP LTM Request Log profiles can be imported using the `name`, e.g.
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import f5bigip:ltm/requestLogProfile:RequestLogProfile test-request-log /Common/test-request-log
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/requestLogProfile:RequestLogProfile")
public class RequestLogProfile extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultsFrom;

    /**
     * @return Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log` profile.
     * 
     */
    public Output<Optional<String>> defaultsFrom() {
        return Codegen.optional(this.defaultsFrom);
    }
    /**
     * Specifies user-defined description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Specifies user-defined description.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Request Logging profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/request-log-profile-tc1`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Defines the pool associated with logging request errors. The default is None.
     * 
     */
    @Export(name="proxyResponse", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxyResponse;

    /**
     * @return Defines the pool associated with logging request errors. The default is None.
     * 
     */
    public Output<Optional<String>> proxyResponse() {
        return Codegen.optional(this.proxyResponse);
    }
    /**
     * Defines the pool associated with logging request errors. The default is None.
     * 
     */
    @Export(name="proxycloseOnError", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxycloseOnError;

    /**
     * @return Defines the pool associated with logging request errors. The default is None.
     * 
     */
    public Output<Optional<String>> proxycloseOnError() {
        return Codegen.optional(this.proxycloseOnError);
    }
    /**
     * Defines the pool associated with logging request errors. The default is None.
     * 
     */
    @Export(name="proxyrespondOnLoggingerror", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxyrespondOnLoggingerror;

    /**
     * @return Defines the pool associated with logging request errors. The default is None.
     * 
     */
    public Output<Optional<String>> proxyrespondOnLoggingerror() {
        return Codegen.optional(this.proxyrespondOnLoggingerror);
    }
    /**
     * Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     * 
     */
    @Export(name="requestLogging", refs={String.class}, tree="[0]")
    private Output<String> requestLogging;

    /**
     * @return Enables or disables request logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     * 
     */
    public Output<String> requestLogging() {
        return this.requestLogging;
    }
    /**
     * Defines the pool associated with logging request errors. The default is None.
     * 
     */
    @Export(name="requestlogErrorPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogErrorPool;

    /**
     * @return Defines the pool associated with logging request errors. The default is None.
     * 
     */
    public Output<Optional<String>> requestlogErrorPool() {
        return Codegen.optional(this.requestlogErrorPool);
    }
    /**
     * Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    @Export(name="requestlogErrorProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogErrorProtocol;

    /**
     * @return Specifies the protocol to be used for high-speed logging of request errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    public Output<Optional<String>> requestlogErrorProtocol() {
        return Codegen.optional(this.requestlogErrorProtocol);
    }
    /**
     * Specifies the directives and entries to be logged for request errors.
     * 
     */
    @Export(name="requestlogErrorTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogErrorTemplate;

    /**
     * @return Specifies the directives and entries to be logged for request errors.
     * 
     */
    public Output<Optional<String>> requestlogErrorTemplate() {
        return Codegen.optional(this.requestlogErrorTemplate);
    }
    /**
     * Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
     * 
     */
    @Export(name="requestlogPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogPool;

    /**
     * @return Defines the pool to send logs to. Typically, the pool will contain one or more syslog servers. It is recommended that you create a pool specifically for logging requests. The default is `none`.
     * 
     */
    public Output<Optional<String>> requestlogPool() {
        return Codegen.optional(this.requestlogPool);
    }
    /**
     * Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    @Export(name="requestlogProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogProtocol;

    /**
     * @return Specifies the protocol to be used for high-speed logging of requests. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    public Output<Optional<String>> requestlogProtocol() {
        return Codegen.optional(this.requestlogProtocol);
    }
    /**
     * Specifies the directives and entries to be logged. More infor on requestlog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     * 
     */
    @Export(name="requestlogTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestlogTemplate;

    /**
     * @return Specifies the directives and entries to be logged. More infor on requestlog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     * 
     */
    public Output<Optional<String>> requestlogTemplate() {
        return Codegen.optional(this.requestlogTemplate);
    }
    /**
     * Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     * 
     */
    @Export(name="responseLogging", refs={String.class}, tree="[0]")
    private Output<String> responseLogging;

    /**
     * @return Enables or disables response logging. The default is `disabled`, possible values are `enabled` and `disabled`.
     * 
     */
    public Output<String> responseLogging() {
        return this.responseLogging;
    }
    /**
     * Defines the pool associated with logging response errors. The default is `none`.
     * 
     */
    @Export(name="responselogErrorPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogErrorPool;

    /**
     * @return Defines the pool associated with logging response errors. The default is `none`.
     * 
     */
    public Output<Optional<String>> responselogErrorPool() {
        return Codegen.optional(this.responselogErrorPool);
    }
    /**
     * Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    @Export(name="responselogErrorProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogErrorProtocol;

    /**
     * @return Specifies the protocol to be used for high-speed logging of response errors. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    public Output<Optional<String>> responselogErrorProtocol() {
        return Codegen.optional(this.responselogErrorProtocol);
    }
    /**
     * Specifies the directives and entries to be logged for request errors.
     * 
     */
    @Export(name="responselogErrorTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogErrorTemplate;

    /**
     * @return Specifies the directives and entries to be logged for request errors.
     * 
     */
    public Output<Optional<String>> responselogErrorTemplate() {
        return Codegen.optional(this.responselogErrorTemplate);
    }
    /**
     * Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
     * 
     */
    @Export(name="responselogPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogPool;

    /**
     * @return Defines the pool to send logs to. Typically, the pool contains one or more syslog servers. It is recommended that you create a pool specifically for logging responses. The default is `none`.
     * 
     */
    public Output<Optional<String>> responselogPool() {
        return Codegen.optional(this.responselogPool);
    }
    /**
     * Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    @Export(name="responselogProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogProtocol;

    /**
     * @return Specifies the protocol to be used for high-speed logging of responses. The default is `mds-udp`,possible values are `mds-udp` and `mds-tcp`.
     * 
     */
    public Output<Optional<String>> responselogProtocol() {
        return Codegen.optional(this.responselogProtocol);
    }
    /**
     * Specifies the directives and entries to be logged. More infor on responselog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     * 
     */
    @Export(name="responselogTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> responselogTemplate;

    /**
     * @return Specifies the directives and entries to be logged. More infor on responselog_template can be found [here](https://techdocs.f5.com/en-us/bigip-15-0-0/external-monitoring-of-big-ip-systems-implementations/configuring-request-logging.html). how to use can be find [here](https://my.f5.com/manage/s/article/K00847516).
     * 
     */
    public Output<Optional<String>> responselogTemplate() {
        return Codegen.optional(this.responselogTemplate);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RequestLogProfile(java.lang.String name) {
        this(name, RequestLogProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RequestLogProfile(java.lang.String name, RequestLogProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RequestLogProfile(java.lang.String name, RequestLogProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RequestLogProfile(java.lang.String name, Output<java.lang.String> id, @Nullable RequestLogProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/requestLogProfile:RequestLogProfile", name, state, makeResourceOptions(options, id), false);
    }

    private static RequestLogProfileArgs makeArgs(RequestLogProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RequestLogProfileArgs.Empty : args;
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
    public static RequestLogProfile get(java.lang.String name, Output<java.lang.String> id, @Nullable RequestLogProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RequestLogProfile(name, id, state, options);
    }
}
