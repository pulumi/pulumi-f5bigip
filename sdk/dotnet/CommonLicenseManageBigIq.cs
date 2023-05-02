// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.CommonLicenseManageBigIq` This Resource is used for BIGIP/Provider License Management from BIGIQ
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // MANAGED Regkey Pool
    ///     var testExampleCommonLicenseManageBigIq = new F5BigIP.CommonLicenseManageBigIq("testExampleCommonLicenseManageBigIq", new()
    ///     {
    ///         BigiqAddress = @var.Bigiq,
    ///         BigiqUser = @var.Bigiq_un,
    ///         BigiqPassword = @var.Bigiq_pw,
    ///         LicensePoolname = "regkeypool_name",
    ///         AssignmentType = "MANAGED",
    ///     });
    /// 
    ///     // UNMANAGED Regkey Pool
    ///     var testExampleIndex_commonLicenseManageBigIqCommonLicenseManageBigIq = new F5BigIP.CommonLicenseManageBigIq("testExampleIndex/commonLicenseManageBigIqCommonLicenseManageBigIq", new()
    ///     {
    ///         BigiqAddress = @var.Bigiq,
    ///         BigiqUser = @var.Bigiq_un,
    ///         BigiqPassword = @var.Bigiq_pw,
    ///         LicensePoolname = "regkeypool_name",
    ///         AssignmentType = "UNMANAGED",
    ///     });
    /// 
    ///     // UNMANAGED Utility Pool
    ///     var testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq = new F5BigIP.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq", new()
    ///     {
    ///         BigiqAddress = @var.Bigiq,
    ///         BigiqUser = @var.Bigiq_un,
    ///         BigiqPassword = @var.Bigiq_pw,
    ///         LicensePoolname = "utilitypool_name",
    ///         AssignmentType = "UNMANAGED",
    ///         UnitOfMeasure = "yearly",
    ///         Skukeyword1 = "BTHSM200M",
    ///     });
    /// 
    ///     // UNREACHABLE Regkey Pool
    ///     var testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq1 = new F5BigIP.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq1", new()
    ///     {
    ///         BigiqAddress = "xxx.xxx.xxx.xxx",
    ///         BigiqUser = "xxxx",
    ///         BigiqPassword = "xxxxx",
    ///         LicensePoolname = "regkey_pool_name",
    ///         AssignmentType = "UNREACHABLE",
    ///         MacAddress = "FA:16:3E:1B:6D:32",
    ///         Hypervisor = "azure",
    ///     });
    /// 
    ///     // MANAGED Purchased Pool
    ///     var testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq2 = new F5BigIP.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq2", new()
    ///     {
    ///         BigiqAddress = @var.Bigiq,
    ///         BigiqUser = @var.Bigiq_un,
    ///         BigiqPassword = @var.Bigiq_pw,
    ///         LicensePoolname = "purchased_pool_name",
    ///         AssignmentType = "MANAGED",
    ///     });
    /// 
    ///     // UNMANAGED Purchased Pool
    ///     var testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq3 = new F5BigIP.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq3", new()
    ///     {
    ///         BigiqAddress = @var.Bigiq,
    ///         BigiqUser = @var.Bigiq_un,
    ///         BigiqPassword = @var.Bigiq_pw,
    ///         LicensePoolname = "purchased_pool_name",
    ///         AssignmentType = "UNMANAGED",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq")]
    public partial class CommonLicenseManageBigIq : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
        /// </summary>
        [Output("assignmentType")]
        public Output<string> AssignmentType { get; private set; } = null!;

        /// <summary>
        /// BIGIQ License Manager IP Address, variable type `string`
        /// </summary>
        [Output("bigiqAddress")]
        public Output<string> BigiqAddress { get; private set; } = null!;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        [Output("bigiqLoginRef")]
        public Output<string?> BigiqLoginRef { get; private set; } = null!;

        /// <summary>
        /// BIGIQ License Manager password.  variable type `string`
        /// </summary>
        [Output("bigiqPassword")]
        public Output<string> BigiqPassword { get; private set; } = null!;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        [Output("bigiqPort")]
        public Output<string?> BigiqPort { get; private set; } = null!;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        [Output("bigiqTokenAuth")]
        public Output<bool?> BigiqTokenAuth { get; private set; } = null!;

        /// <summary>
        /// BIGIQ License Manager username, variable type `string`
        /// </summary>
        [Output("bigiqUser")]
        public Output<string> BigiqUser { get; private set; } = null!;

        /// <summary>
        /// Status of Licence Assignment
        /// </summary>
        [Output("deviceLicenseStatus")]
        public Output<string> DeviceLicenseStatus { get; private set; } = null!;

        /// <summary>
        /// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
        /// </summary>
        [Output("hypervisor")]
        public Output<string?> Hypervisor { get; private set; } = null!;

        /// <summary>
        /// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// A name given to the license pool. type `string`
        /// </summary>
        [Output("licensePoolname")]
        public Output<string> LicensePoolname { get; private set; } = null!;

        /// <summary>
        /// MAC address of the BIG-IP. type `string`
        /// </summary>
        [Output("macAddress")]
        public Output<string?> MacAddress { get; private set; } = null!;

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Output("skukeyword1")]
        public Output<string?> Skukeyword1 { get; private set; } = null!;

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Output("skukeyword2")]
        public Output<string?> Skukeyword2 { get; private set; } = null!;

        /// <summary>
        /// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
        /// </summary>
        [Output("tenant")]
        public Output<string?> Tenant { get; private set; } = null!;

        /// <summary>
        /// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
        /// </summary>
        [Output("unitOfMeasure")]
        public Output<string?> UnitOfMeasure { get; private set; } = null!;


        /// <summary>
        /// Create a CommonLicenseManageBigIq resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CommonLicenseManageBigIq(string name, CommonLicenseManageBigIqArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq", name, args ?? new CommonLicenseManageBigIqArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CommonLicenseManageBigIq(string name, Input<string> id, CommonLicenseManageBigIqState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "bigiqLoginRef",
                    "bigiqPassword",
                    "bigiqPort",
                    "bigiqTokenAuth",
                    "bigiqUser",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CommonLicenseManageBigIq resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CommonLicenseManageBigIq Get(string name, Input<string> id, CommonLicenseManageBigIqState? state = null, CustomResourceOptions? options = null)
        {
            return new CommonLicenseManageBigIq(name, id, state, options);
        }
    }

    public sealed class CommonLicenseManageBigIqArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
        /// </summary>
        [Input("assignmentType", required: true)]
        public Input<string> AssignmentType { get; set; } = null!;

        /// <summary>
        /// BIGIQ License Manager IP Address, variable type `string`
        /// </summary>
        [Input("bigiqAddress", required: true)]
        public Input<string> BigiqAddress { get; set; } = null!;

        [Input("bigiqLoginRef")]
        private Input<string>? _bigiqLoginRef;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        public Input<string>? BigiqLoginRef
        {
            get => _bigiqLoginRef;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqLoginRef = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPassword", required: true)]
        private Input<string>? _bigiqPassword;

        /// <summary>
        /// BIGIQ License Manager password.  variable type `string`
        /// </summary>
        public Input<string>? BigiqPassword
        {
            get => _bigiqPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPort")]
        private Input<string>? _bigiqPort;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        public Input<string>? BigiqPort
        {
            get => _bigiqPort;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPort = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqTokenAuth")]
        private Input<bool>? _bigiqTokenAuth;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        public Input<bool>? BigiqTokenAuth
        {
            get => _bigiqTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqUser", required: true)]
        private Input<string>? _bigiqUser;

        /// <summary>
        /// BIGIQ License Manager username, variable type `string`
        /// </summary>
        public Input<string>? BigiqUser
        {
            get => _bigiqUser;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqUser = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Status of Licence Assignment
        /// </summary>
        [Input("deviceLicenseStatus")]
        public Input<string>? DeviceLicenseStatus { get; set; }

        /// <summary>
        /// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
        /// </summary>
        [Input("hypervisor")]
        public Input<string>? Hypervisor { get; set; }

        /// <summary>
        /// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A name given to the license pool. type `string`
        /// </summary>
        [Input("licensePoolname", required: true)]
        public Input<string> LicensePoolname { get; set; } = null!;

        /// <summary>
        /// MAC address of the BIG-IP. type `string`
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Input("skukeyword1")]
        public Input<string>? Skukeyword1 { get; set; }

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Input("skukeyword2")]
        public Input<string>? Skukeyword2 { get; set; }

        /// <summary>
        /// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
        /// </summary>
        [Input("unitOfMeasure")]
        public Input<string>? UnitOfMeasure { get; set; }

        public CommonLicenseManageBigIqArgs()
        {
        }
        public static new CommonLicenseManageBigIqArgs Empty => new CommonLicenseManageBigIqArgs();
    }

    public sealed class CommonLicenseManageBigIqState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
        /// </summary>
        [Input("assignmentType")]
        public Input<string>? AssignmentType { get; set; }

        /// <summary>
        /// BIGIQ License Manager IP Address, variable type `string`
        /// </summary>
        [Input("bigiqAddress")]
        public Input<string>? BigiqAddress { get; set; }

        [Input("bigiqLoginRef")]
        private Input<string>? _bigiqLoginRef;

        /// <summary>
        /// BIGIQ Login reference for token authentication
        /// </summary>
        public Input<string>? BigiqLoginRef
        {
            get => _bigiqLoginRef;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqLoginRef = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPassword")]
        private Input<string>? _bigiqPassword;

        /// <summary>
        /// BIGIQ License Manager password.  variable type `string`
        /// </summary>
        public Input<string>? BigiqPassword
        {
            get => _bigiqPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqPort")]
        private Input<string>? _bigiqPort;

        /// <summary>
        /// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        /// </summary>
        public Input<string>? BigiqPort
        {
            get => _bigiqPort;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqPort = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqTokenAuth")]
        private Input<bool>? _bigiqTokenAuth;

        /// <summary>
        /// type `bool`, if set to `true` enables Token based Authentication,default is `false`
        /// </summary>
        public Input<bool>? BigiqTokenAuth
        {
            get => _bigiqTokenAuth;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqTokenAuth = Output.Tuple<Input<bool>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bigiqUser")]
        private Input<string>? _bigiqUser;

        /// <summary>
        /// BIGIQ License Manager username, variable type `string`
        /// </summary>
        public Input<string>? BigiqUser
        {
            get => _bigiqUser;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bigiqUser = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Status of Licence Assignment
        /// </summary>
        [Input("deviceLicenseStatus")]
        public Input<string>? DeviceLicenseStatus { get; set; }

        /// <summary>
        /// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
        /// </summary>
        [Input("hypervisor")]
        public Input<string>? Hypervisor { get; set; }

        /// <summary>
        /// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A name given to the license pool. type `string`
        /// </summary>
        [Input("licensePoolname")]
        public Input<string>? LicensePoolname { get; set; }

        /// <summary>
        /// MAC address of the BIG-IP. type `string`
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Input("skukeyword1")]
        public Input<string>? Skukeyword1 { get; set; }

        /// <summary>
        /// An optional offering name. type `string`
        /// </summary>
        [Input("skukeyword2")]
        public Input<string>? Skukeyword2 { get; set; }

        /// <summary>
        /// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
        /// </summary>
        [Input("unitOfMeasure")]
        public Input<string>? UnitOfMeasure { get; set; }

        public CommonLicenseManageBigIqState()
        {
        }
        public static new CommonLicenseManageBigIqState Empty => new CommonLicenseManageBigIqState();
    }
}
