// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `CommonLicenseManageBigIq` This Resource is used for BIGIP/Provider License Management from BIGIQ
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := f5bigip.NewCommonLicenseManageBigIq(ctx, "testExampleCommonLicenseManageBigIq", &f5bigip.CommonLicenseManageBigIqArgs{
// 			BigiqAddress:    pulumi.Any(_var.Bigiq),
// 			BigiqUser:       pulumi.Any(_var.Bigiq_un),
// 			BigiqPassword:   pulumi.Any(_var.Bigiq_pw),
// 			LicensePoolname: pulumi.String("regkeypool_name"),
// 			AssignmentType:  pulumi.String("MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = f5bigip.NewCommonLicenseManageBigIq(ctx, "testExampleIndex_commonLicenseManageBigIqCommonLicenseManageBigIq", &f5bigip.CommonLicenseManageBigIqArgs{
// 			BigiqAddress:    pulumi.Any(_var.Bigiq),
// 			BigiqUser:       pulumi.Any(_var.Bigiq_un),
// 			BigiqPassword:   pulumi.Any(_var.Bigiq_pw),
// 			LicensePoolname: pulumi.String("regkeypool_name"),
// 			AssignmentType:  pulumi.String("UNMANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = f5bigip.NewCommonLicenseManageBigIq(ctx, "testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq", &f5bigip.CommonLicenseManageBigIqArgs{
// 			BigiqAddress:    pulumi.Any(_var.Bigiq),
// 			BigiqUser:       pulumi.Any(_var.Bigiq_un),
// 			BigiqPassword:   pulumi.Any(_var.Bigiq_pw),
// 			LicensePoolname: pulumi.String("utilitypool_name"),
// 			AssignmentType:  pulumi.String("UNMANAGED"),
// 			UnitOfMeasure:   pulumi.String("yearly"),
// 			Skukeyword1:     pulumi.String("BTHSM200M"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = f5bigip.NewCommonLicenseManageBigIq(ctx, "testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq1", &f5bigip.CommonLicenseManageBigIqArgs{
// 			BigiqAddress:    pulumi.String("xxx.xxx.xxx.xxx"),
// 			BigiqUser:       pulumi.String("xxxx"),
// 			BigiqPassword:   pulumi.String("xxxxx"),
// 			LicensePoolname: pulumi.String("regkey_pool_name"),
// 			AssignmentType:  pulumi.String("UNREACHABLE"),
// 			MacAddress:      pulumi.String("FA:16:3E:1B:6D:32"),
// 			Hypervisor:      pulumi.String("azure"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CommonLicenseManageBigIq struct {
	pulumi.CustomResourceState

	// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
	AssignmentType pulumi.StringOutput `pulumi:"assignmentType"`
	// BIGIQ License Manager IP Address, variable type `string`
	BigiqAddress pulumi.StringOutput `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrOutput `pulumi:"bigiqLoginRef"`
	// BIGIQ License Manager password.  variable type `string`
	BigiqPassword pulumi.StringOutput `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrOutput `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrOutput `pulumi:"bigiqTokenAuth"`
	// BIGIQ License Manager username, variable type `string`
	BigiqUser pulumi.StringOutput `pulumi:"bigiqUser"`
	// Status of Licence Assignment
	DeviceLicenseStatus pulumi.StringOutput `pulumi:"deviceLicenseStatus"`
	// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
	Hypervisor pulumi.StringPtrOutput `pulumi:"hypervisor"`
	// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// A name given to the license pool. type `string`
	LicensePoolname pulumi.StringOutput `pulumi:"licensePoolname"`
	// MAC address of the BIG-IP. type `string`
	MacAddress pulumi.StringPtrOutput `pulumi:"macAddress"`
	// An optional offering name. type `string`
	Skukeyword1 pulumi.StringPtrOutput `pulumi:"skukeyword1"`
	// An optional offering name. type `string`
	Skukeyword2 pulumi.StringPtrOutput `pulumi:"skukeyword2"`
	// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
	Tenant pulumi.StringPtrOutput `pulumi:"tenant"`
	// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
	UnitOfMeasure pulumi.StringPtrOutput `pulumi:"unitOfMeasure"`
}

// NewCommonLicenseManageBigIq registers a new resource with the given unique name, arguments, and options.
func NewCommonLicenseManageBigIq(ctx *pulumi.Context,
	name string, args *CommonLicenseManageBigIqArgs, opts ...pulumi.ResourceOption) (*CommonLicenseManageBigIq, error) {
	if args == nil || args.AssignmentType == nil {
		return nil, errors.New("missing required argument 'AssignmentType'")
	}
	if args == nil || args.BigiqAddress == nil {
		return nil, errors.New("missing required argument 'BigiqAddress'")
	}
	if args == nil || args.BigiqPassword == nil {
		return nil, errors.New("missing required argument 'BigiqPassword'")
	}
	if args == nil || args.BigiqUser == nil {
		return nil, errors.New("missing required argument 'BigiqUser'")
	}
	if args == nil || args.LicensePoolname == nil {
		return nil, errors.New("missing required argument 'LicensePoolname'")
	}
	if args == nil {
		args = &CommonLicenseManageBigIqArgs{}
	}
	var resource CommonLicenseManageBigIq
	err := ctx.RegisterResource("f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommonLicenseManageBigIq gets an existing CommonLicenseManageBigIq resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommonLicenseManageBigIq(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommonLicenseManageBigIqState, opts ...pulumi.ResourceOption) (*CommonLicenseManageBigIq, error) {
	var resource CommonLicenseManageBigIq
	err := ctx.ReadResource("f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommonLicenseManageBigIq resources.
type commonLicenseManageBigIqState struct {
	// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
	AssignmentType *string `pulumi:"assignmentType"`
	// BIGIQ License Manager IP Address, variable type `string`
	BigiqAddress *string `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef *string `pulumi:"bigiqLoginRef"`
	// BIGIQ License Manager password.  variable type `string`
	BigiqPassword *string `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort *string `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth *bool `pulumi:"bigiqTokenAuth"`
	// BIGIQ License Manager username, variable type `string`
	BigiqUser *string `pulumi:"bigiqUser"`
	// Status of Licence Assignment
	DeviceLicenseStatus *string `pulumi:"deviceLicenseStatus"`
	// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
	Hypervisor *string `pulumi:"hypervisor"`
	// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
	Key *string `pulumi:"key"`
	// A name given to the license pool. type `string`
	LicensePoolname *string `pulumi:"licensePoolname"`
	// MAC address of the BIG-IP. type `string`
	MacAddress *string `pulumi:"macAddress"`
	// An optional offering name. type `string`
	Skukeyword1 *string `pulumi:"skukeyword1"`
	// An optional offering name. type `string`
	Skukeyword2 *string `pulumi:"skukeyword2"`
	// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
	Tenant *string `pulumi:"tenant"`
	// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
	UnitOfMeasure *string `pulumi:"unitOfMeasure"`
}

type CommonLicenseManageBigIqState struct {
	// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
	AssignmentType pulumi.StringPtrInput
	// BIGIQ License Manager IP Address, variable type `string`
	BigiqAddress pulumi.StringPtrInput
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrInput
	// BIGIQ License Manager password.  variable type `string`
	BigiqPassword pulumi.StringPtrInput
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrInput
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrInput
	// BIGIQ License Manager username, variable type `string`
	BigiqUser pulumi.StringPtrInput
	// Status of Licence Assignment
	DeviceLicenseStatus pulumi.StringPtrInput
	// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
	Hypervisor pulumi.StringPtrInput
	// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
	Key pulumi.StringPtrInput
	// A name given to the license pool. type `string`
	LicensePoolname pulumi.StringPtrInput
	// MAC address of the BIG-IP. type `string`
	MacAddress pulumi.StringPtrInput
	// An optional offering name. type `string`
	Skukeyword1 pulumi.StringPtrInput
	// An optional offering name. type `string`
	Skukeyword2 pulumi.StringPtrInput
	// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
	Tenant pulumi.StringPtrInput
	// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
	UnitOfMeasure pulumi.StringPtrInput
}

func (CommonLicenseManageBigIqState) ElementType() reflect.Type {
	return reflect.TypeOf((*commonLicenseManageBigIqState)(nil)).Elem()
}

type commonLicenseManageBigIqArgs struct {
	// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
	AssignmentType string `pulumi:"assignmentType"`
	// BIGIQ License Manager IP Address, variable type `string`
	BigiqAddress string `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef *string `pulumi:"bigiqLoginRef"`
	// BIGIQ License Manager password.  variable type `string`
	BigiqPassword string `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort *string `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth *bool `pulumi:"bigiqTokenAuth"`
	// BIGIQ License Manager username, variable type `string`
	BigiqUser string `pulumi:"bigiqUser"`
	// Status of Licence Assignment
	DeviceLicenseStatus *string `pulumi:"deviceLicenseStatus"`
	// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
	Hypervisor *string `pulumi:"hypervisor"`
	// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
	Key *string `pulumi:"key"`
	// A name given to the license pool. type `string`
	LicensePoolname string `pulumi:"licensePoolname"`
	// MAC address of the BIG-IP. type `string`
	MacAddress *string `pulumi:"macAddress"`
	// An optional offering name. type `string`
	Skukeyword1 *string `pulumi:"skukeyword1"`
	// An optional offering name. type `string`
	Skukeyword2 *string `pulumi:"skukeyword2"`
	// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
	Tenant *string `pulumi:"tenant"`
	// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
	UnitOfMeasure *string `pulumi:"unitOfMeasure"`
}

// The set of arguments for constructing a CommonLicenseManageBigIq resource.
type CommonLicenseManageBigIqArgs struct {
	// The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
	AssignmentType pulumi.StringInput
	// BIGIQ License Manager IP Address, variable type `string`
	BigiqAddress pulumi.StringInput
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrInput
	// BIGIQ License Manager password.  variable type `string`
	BigiqPassword pulumi.StringInput
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrInput
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrInput
	// BIGIQ License Manager username, variable type `string`
	BigiqUser pulumi.StringInput
	// Status of Licence Assignment
	DeviceLicenseStatus pulumi.StringPtrInput
	// Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
	Hypervisor pulumi.StringPtrInput
	// License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
	Key pulumi.StringPtrInput
	// A name given to the license pool. type `string`
	LicensePoolname pulumi.StringInput
	// MAC address of the BIG-IP. type `string`
	MacAddress pulumi.StringPtrInput
	// An optional offering name. type `string`
	Skukeyword1 pulumi.StringPtrInput
	// An optional offering name. type `string`
	Skukeyword2 pulumi.StringPtrInput
	// For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
	Tenant pulumi.StringPtrInput
	// The units used to measure billing. For example, “hourly” or “daily”. Type `string`
	UnitOfMeasure pulumi.StringPtrInput
}

func (CommonLicenseManageBigIqArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commonLicenseManageBigIqArgs)(nil)).Elem()
}

type CommonLicenseManageBigIqInput interface {
	pulumi.Input

	ToCommonLicenseManageBigIqOutput() CommonLicenseManageBigIqOutput
	ToCommonLicenseManageBigIqOutputWithContext(ctx context.Context) CommonLicenseManageBigIqOutput
}

func (CommonLicenseManageBigIq) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonLicenseManageBigIq)(nil)).Elem()
}

func (i CommonLicenseManageBigIq) ToCommonLicenseManageBigIqOutput() CommonLicenseManageBigIqOutput {
	return i.ToCommonLicenseManageBigIqOutputWithContext(context.Background())
}

func (i CommonLicenseManageBigIq) ToCommonLicenseManageBigIqOutputWithContext(ctx context.Context) CommonLicenseManageBigIqOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonLicenseManageBigIqOutput)
}

type CommonLicenseManageBigIqOutput struct {
	*pulumi.OutputState
}

func (CommonLicenseManageBigIqOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonLicenseManageBigIqOutput)(nil)).Elem()
}

func (o CommonLicenseManageBigIqOutput) ToCommonLicenseManageBigIqOutput() CommonLicenseManageBigIqOutput {
	return o
}

func (o CommonLicenseManageBigIqOutput) ToCommonLicenseManageBigIqOutputWithContext(ctx context.Context) CommonLicenseManageBigIqOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CommonLicenseManageBigIqOutput{})
}
