// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryPlainArgs;
import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryPlainArgs;
import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryPlainArgs;
import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryPlainArgs;
import com.pulumi.f5bigip.fast.outputs.GetAwsServiceDiscoveryResult;
import com.pulumi.f5bigip.fast.outputs.GetAzureServiceDiscoveryResult;
import com.pulumi.f5bigip.fast.outputs.GetConsulServiceDiscoveryResult;
import com.pulumi.f5bigip.fast.outputs.GetGceServiceDiscoveryResult;
import java.util.concurrent.CompletableFuture;

public final class FastFunctions {
    /**
     * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey("testawstagkey")
     *             .tagValue("testawstagvalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAwsServiceDiscoveryResult> getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs args) {
        return getAwsServiceDiscovery(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey("testawstagkey")
     *             .tagValue("testawstagvalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAwsServiceDiscoveryResult> getAwsServiceDiscoveryPlain(GetAwsServiceDiscoveryPlainArgs args) {
        return getAwsServiceDiscoveryPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey("testawstagkey")
     *             .tagValue("testawstagvalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAwsServiceDiscoveryResult> getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", TypeShape.of(GetAwsServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey("testawstagkey")
     *             .tagValue("testawstagvalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAwsServiceDiscoveryResult> getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", TypeShape.of(GetAwsServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAwsServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey("testawstagkey")
     *             .tagValue("testawstagvalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAwsServiceDiscoveryResult> getAwsServiceDiscoveryPlain(GetAwsServiceDiscoveryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", TypeShape.of(GetAwsServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup("testazurerg")
     *             .subscriptionId("testazuresid")
     *             .tagKey("testazuretag")
     *             .tagValue("testazurevalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAzureServiceDiscoveryResult> getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs args) {
        return getAzureServiceDiscovery(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup("testazurerg")
     *             .subscriptionId("testazuresid")
     *             .tagKey("testazuretag")
     *             .tagValue("testazurevalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAzureServiceDiscoveryResult> getAzureServiceDiscoveryPlain(GetAzureServiceDiscoveryPlainArgs args) {
        return getAzureServiceDiscoveryPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup("testazurerg")
     *             .subscriptionId("testazuresid")
     *             .tagKey("testazuretag")
     *             .tagValue("testazurevalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAzureServiceDiscoveryResult> getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", TypeShape.of(GetAzureServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup("testazurerg")
     *             .subscriptionId("testazuresid")
     *             .tagKey("testazuretag")
     *             .tagValue("testazurevalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAzureServiceDiscoveryResult> getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", TypeShape.of(GetAzureServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetAzureServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup("testazurerg")
     *             .subscriptionId("testazuresid")
     *             .tagKey("testazuretag")
     *             .tagValue("testazurevalue")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAzureServiceDiscoveryResult> getAzureServiceDiscoveryPlain(GetAzureServiceDiscoveryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", TypeShape.of(GetAzureServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .uri("https://192.0.2.100:8500/v1/catalog/nodes")
     *             .port(8080)
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConsulServiceDiscoveryResult> getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs args) {
        return getConsulServiceDiscovery(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .uri("https://192.0.2.100:8500/v1/catalog/nodes")
     *             .port(8080)
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetConsulServiceDiscoveryResult> getConsulServiceDiscoveryPlain(GetConsulServiceDiscoveryPlainArgs args) {
        return getConsulServiceDiscoveryPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .uri("https://192.0.2.100:8500/v1/catalog/nodes")
     *             .port(8080)
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConsulServiceDiscoveryResult> getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", TypeShape.of(GetConsulServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .uri("https://192.0.2.100:8500/v1/catalog/nodes")
     *             .port(8080)
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConsulServiceDiscoveryResult> getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", TypeShape.of(GetConsulServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetConsulServiceDiscoveryArgs;
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
     *         final var TC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .uri("https://192.0.2.100:8500/v1/catalog/nodes")
     *             .port(8080)
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetConsulServiceDiscoveryResult> getConsulServiceDiscoveryPlain(GetConsulServiceDiscoveryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", TypeShape.of(GetConsulServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .tagKey("testgcetag")
     *             .tagValue("testgcevalue")
     *             .region("testgceregion")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetGceServiceDiscoveryResult> getGceServiceDiscovery(GetGceServiceDiscoveryArgs args) {
        return getGceServiceDiscovery(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .tagKey("testgcetag")
     *             .tagValue("testgcevalue")
     *             .region("testgceregion")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetGceServiceDiscoveryResult> getGceServiceDiscoveryPlain(GetGceServiceDiscoveryPlainArgs args) {
        return getGceServiceDiscoveryPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .tagKey("testgcetag")
     *             .tagValue("testgcevalue")
     *             .region("testgceregion")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetGceServiceDiscoveryResult> getGceServiceDiscovery(GetGceServiceDiscoveryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", TypeShape.of(GetGceServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .tagKey("testgcetag")
     *             .tagValue("testgcevalue")
     *             .region("testgceregion")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetGceServiceDiscoveryResult> getGceServiceDiscovery(GetGceServiceDiscoveryArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", TypeShape.of(GetGceServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
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
     * import com.pulumi.f5bigip.fast.FastFunctions;
     * import com.pulumi.f5bigip.fast.inputs.GetGceServiceDiscoveryArgs;
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
     *         final var TC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .tagKey("testgcetag")
     *             .tagValue("testgcevalue")
     *             .region("testgceregion")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetGceServiceDiscoveryResult> getGceServiceDiscoveryPlain(GetGceServiceDiscoveryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", TypeShape.of(GetGceServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
}
