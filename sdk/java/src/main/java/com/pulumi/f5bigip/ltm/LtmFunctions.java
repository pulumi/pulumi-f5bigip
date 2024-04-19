// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.inputs.GetDataGroupArgs;
import com.pulumi.f5bigip.ltm.inputs.GetDataGroupPlainArgs;
import com.pulumi.f5bigip.ltm.inputs.GetIruleArgs;
import com.pulumi.f5bigip.ltm.inputs.GetIrulePlainArgs;
import com.pulumi.f5bigip.ltm.inputs.GetMonitorArgs;
import com.pulumi.f5bigip.ltm.inputs.GetMonitorPlainArgs;
import com.pulumi.f5bigip.ltm.inputs.GetNodeArgs;
import com.pulumi.f5bigip.ltm.inputs.GetNodePlainArgs;
import com.pulumi.f5bigip.ltm.inputs.GetPolicyArgs;
import com.pulumi.f5bigip.ltm.inputs.GetPolicyPlainArgs;
import com.pulumi.f5bigip.ltm.inputs.GetPoolArgs;
import com.pulumi.f5bigip.ltm.inputs.GetPoolPlainArgs;
import com.pulumi.f5bigip.ltm.outputs.GetDataGroupResult;
import com.pulumi.f5bigip.ltm.outputs.GetIruleResult;
import com.pulumi.f5bigip.ltm.outputs.GetMonitorResult;
import com.pulumi.f5bigip.ltm.outputs.GetNodeResult;
import com.pulumi.f5bigip.ltm.outputs.GetPolicyResult;
import com.pulumi.f5bigip.ltm.outputs.GetPoolResult;
import java.util.concurrent.CompletableFuture;

public final class LtmFunctions {
    /**
     * Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetDataGroupArgs;
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
     *         final var DG-TC3 = LtmFunctions.getDataGroup(GetDataGroupArgs.builder()
     *             .name(&#34;test-dg&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDataGroupResult> getDataGroup(GetDataGroupArgs args) {
        return getDataGroup(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetDataGroupArgs;
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
     *         final var DG-TC3 = LtmFunctions.getDataGroup(GetDataGroupArgs.builder()
     *             .name(&#34;test-dg&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDataGroupResult> getDataGroupPlain(GetDataGroupPlainArgs args) {
        return getDataGroupPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetDataGroupArgs;
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
     *         final var DG-TC3 = LtmFunctions.getDataGroup(GetDataGroupArgs.builder()
     *             .name(&#34;test-dg&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDataGroupResult> getDataGroup(GetDataGroupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getDataGroup:getDataGroup", TypeShape.of(GetDataGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetDataGroupArgs;
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
     *         final var DG-TC3 = LtmFunctions.getDataGroup(GetDataGroupArgs.builder()
     *             .name(&#34;test-dg&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDataGroupResult> getDataGroupPlain(GetDataGroupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getDataGroup:getDataGroup", TypeShape.of(GetDataGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetIruleArgs;
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
     *         final var test = LtmFunctions.getIrule(GetIruleArgs.builder()
     *             .name(&#34;terraform_irule&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipIrule&#34;, test.applyValue(getIruleResult -&gt; getIruleResult.irule()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIruleResult> getIrule(GetIruleArgs args) {
        return getIrule(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetIruleArgs;
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
     *         final var test = LtmFunctions.getIrule(GetIruleArgs.builder()
     *             .name(&#34;terraform_irule&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipIrule&#34;, test.applyValue(getIruleResult -&gt; getIruleResult.irule()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIruleResult> getIrulePlain(GetIrulePlainArgs args) {
        return getIrulePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetIruleArgs;
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
     *         final var test = LtmFunctions.getIrule(GetIruleArgs.builder()
     *             .name(&#34;terraform_irule&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipIrule&#34;, test.applyValue(getIruleResult -&gt; getIruleResult.irule()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIruleResult> getIrule(GetIruleArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getIrule:getIrule", TypeShape.of(GetIruleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetIruleArgs;
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
     *         final var test = LtmFunctions.getIrule(GetIruleArgs.builder()
     *             .name(&#34;terraform_irule&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipIrule&#34;, test.applyValue(getIruleResult -&gt; getIruleResult.irule()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIruleResult> getIrulePlain(GetIrulePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getIrule:getIrule", TypeShape.of(GetIruleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetMonitorArgs;
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
     *         final var monitor-TC1 = LtmFunctions.getMonitor(GetMonitorArgs.builder()
     *             .name(&#34;test-monitor&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetMonitorResult> getMonitor(GetMonitorArgs args) {
        return getMonitor(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetMonitorArgs;
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
     *         final var monitor-TC1 = LtmFunctions.getMonitor(GetMonitorArgs.builder()
     *             .name(&#34;test-monitor&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetMonitorResult> getMonitorPlain(GetMonitorPlainArgs args) {
        return getMonitorPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetMonitorArgs;
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
     *         final var monitor-TC1 = LtmFunctions.getMonitor(GetMonitorArgs.builder()
     *             .name(&#34;test-monitor&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetMonitorResult> getMonitor(GetMonitorArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getMonitor:getMonitor", TypeShape.of(GetMonitorResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetMonitorArgs;
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
     *         final var monitor-TC1 = LtmFunctions.getMonitor(GetMonitorArgs.builder()
     *             .name(&#34;test-monitor&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetMonitorResult> getMonitorPlain(GetMonitorPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getMonitor:getMonitor", TypeShape.of(GetMonitorResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Node`) to get the ltm node details available on BIG-IP
     * 
     */
    public static Output<GetNodeResult> getNode(GetNodeArgs args) {
        return getNode(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Node`) to get the ltm node details available on BIG-IP
     * 
     */
    public static CompletableFuture<GetNodeResult> getNodePlain(GetNodePlainArgs args) {
        return getNodePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Node`) to get the ltm node details available on BIG-IP
     * 
     */
    public static Output<GetNodeResult> getNode(GetNodeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getNode:getNode", TypeShape.of(GetNodeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Node`) to get the ltm node details available on BIG-IP
     * 
     */
    public static CompletableFuture<GetNodeResult> getNodePlain(GetNodePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getNode:getNode", TypeShape.of(GetNodeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPolicyArgs;
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
     *         final var test = LtmFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name(&#34;/Common/test-policy&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipPolicy&#34;, test.applyValue(getPolicyResult -&gt; getPolicyResult.rules()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args) {
        return getPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPolicyArgs;
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
     *         final var test = LtmFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name(&#34;/Common/test-policy&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipPolicy&#34;, test.applyValue(getPolicyResult -&gt; getPolicyResult.rules()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args) {
        return getPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPolicyArgs;
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
     *         final var test = LtmFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name(&#34;/Common/test-policy&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipPolicy&#34;, test.applyValue(getPolicyResult -&gt; getPolicyResult.rules()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPolicyArgs;
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
     *         final var test = LtmFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name(&#34;/Common/test-policy&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;bigipPolicy&#34;, test.applyValue(getPolicyResult -&gt; getPolicyResult.rules()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPoolArgs;
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
     *         final var pool-Example = LtmFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;example-pool&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPoolResult> getPool(GetPoolArgs args) {
        return getPool(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPoolArgs;
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
     *         final var pool-Example = LtmFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;example-pool&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPoolResult> getPoolPlain(GetPoolPlainArgs args) {
        return getPoolPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPoolArgs;
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
     *         final var pool-Example = LtmFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;example-pool&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPoolResult> getPool(GetPoolArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("f5bigip:ltm/getPool:getPool", TypeShape.of(GetPoolResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
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
     * import com.pulumi.f5bigip.ltm.LtmFunctions;
     * import com.pulumi.f5bigip.ltm.inputs.GetPoolArgs;
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
     *         final var pool-Example = LtmFunctions.getPool(GetPoolArgs.builder()
     *             .name(&#34;example-pool&#34;)
     *             .partition(&#34;Common&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPoolResult> getPoolPlain(GetPoolPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("f5bigip:ltm/getPool:getPool", TypeShape.of(GetPoolResult.class), args, Utilities.withVersion(options));
    }
}
