// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.fast;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
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
     * ```java
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
     *         final var tC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey(&#34;testawstagkey&#34;)
     *             .tagValue(&#34;testawstagvalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey(&#34;testawstagkey&#34;)
     *             .tagValue(&#34;testawstagvalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey(&#34;testawstagkey&#34;)
     *             .tagValue(&#34;testawstagvalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getAwsServiceDiscovery(GetAwsServiceDiscoveryArgs.builder()
     *             .tagKey(&#34;testawstagkey&#34;)
     *             .tagValue(&#34;testawstagvalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup(&#34;testazurerg&#34;)
     *             .subscriptionId(&#34;testazuresid&#34;)
     *             .tagKey(&#34;testazuretag&#34;)
     *             .tagValue(&#34;testazurevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup(&#34;testazurerg&#34;)
     *             .subscriptionId(&#34;testazuresid&#34;)
     *             .tagKey(&#34;testazuretag&#34;)
     *             .tagValue(&#34;testazurevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup(&#34;testazurerg&#34;)
     *             .subscriptionId(&#34;testazuresid&#34;)
     *             .tagKey(&#34;testazuretag&#34;)
     *             .tagValue(&#34;testazurevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getAzureServiceDiscovery(GetAzureServiceDiscoveryArgs.builder()
     *             .resourceGroup(&#34;testazurerg&#34;)
     *             .subscriptionId(&#34;testazuresid&#34;)
     *             .tagKey(&#34;testazuretag&#34;)
     *             .tagValue(&#34;testazurevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .port(8080)
     *             .uri(&#34;https://192.0.2.100:8500/v1/catalog/nodes&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .port(8080)
     *             .uri(&#34;https://192.0.2.100:8500/v1/catalog/nodes&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .port(8080)
     *             .uri(&#34;https://192.0.2.100:8500/v1/catalog/nodes&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC2 = FastFunctions.getConsulServiceDiscovery(GetConsulServiceDiscoveryArgs.builder()
     *             .port(8080)
     *             .uri(&#34;https://192.0.2.100:8500/v1/catalog/nodes&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .region(&#34;testgceregion&#34;)
     *             .tagKey(&#34;testgcetag&#34;)
     *             .tagValue(&#34;testgcevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .region(&#34;testgceregion&#34;)
     *             .tagKey(&#34;testgcetag&#34;)
     *             .tagValue(&#34;testgcevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .region(&#34;testgceregion&#34;)
     *             .tagKey(&#34;testgcetag&#34;)
     *             .tagValue(&#34;testgcevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
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
     * ```java
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
     *         final var tC3 = FastFunctions.getGceServiceDiscovery(GetGceServiceDiscoveryArgs.builder()
     *             .region(&#34;testgceregion&#34;)
     *             .tagKey(&#34;testgcetag&#34;)
     *             .tagValue(&#34;testgcevalue&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetGceServiceDiscoveryResult> getGceServiceDiscoveryPlain(GetGceServiceDiscoveryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", TypeShape.of(GetGceServiceDiscoveryResult.class), args, Utilities.withVersion(options));
    }
}
