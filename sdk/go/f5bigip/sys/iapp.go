// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.
//
// ## Example Usage
// ### Json File
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
//  * `description` - User defined description.
//  * `deviceGroup` - The name of the device group that the application service is assigned to.
//  * `executeAction` - Run the specified template action associated with the application.
//  * `inheritedDevicegroup`- Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
//  * `inheritedTrafficGroup` - Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
//  * `partition` - Displays the administrative partition within which the application resides.
//  * `strictUpdates` - Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
//  * `template` - The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
//  * `templateModified` - Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
//  * `templatePrerequisiteErrors` - Indicates any missing prerequisites associated with the template that defines this application.
//  * `trafficGroup` - The name of the traffic group that the application service is assigned to.
//  * `lists` - string values
//  * `metadata` - User defined generic data for the application service. It is a name and value pair.
//  * `tables` - Values provided like pool name, nodes etc.
//  * `variables` - Name, values, encrypted or not
type IApp struct {
	pulumi.CustomResourceState

	// Address of the Iapp which needs to be Iappensed
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// BIG-IP password
	Devicegroup pulumi.StringPtrOutput `pulumi:"devicegroup"`
	// BIG-IP password
	ExecuteAction pulumi.StringPtrOutput `pulumi:"executeAction"`
	// BIG-IP password
	InheritedDevicegroup pulumi.StringPtrOutput `pulumi:"inheritedDevicegroup"`
	// BIG-IP password
	InheritedTrafficGroup pulumi.StringPtrOutput `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile  pulumi.StringPtrOutput  `pulumi:"jsonfile"`
	Lists     IAppListArrayOutput     `pulumi:"lists"`
	Metadatas IAppMetadataArrayOutput `pulumi:"metadatas"`
	// Name of the iApp.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Address of the Iapp which needs to be Iappensed
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// BIG-IP password
	StrictUpdates pulumi.StringPtrOutput `pulumi:"strictUpdates"`
	Tables        IAppTableArrayOutput   `pulumi:"tables"`
	// BIG-IP password
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// BIG-IP password
	TemplateModified pulumi.StringPtrOutput `pulumi:"templateModified"`
	// BIG-IP password
	TemplatePrerequisiteErrors pulumi.StringPtrOutput `pulumi:"templatePrerequisiteErrors"`
	// BIG-IP password
	TrafficGroup pulumi.StringPtrOutput  `pulumi:"trafficGroup"`
	Variables    IAppVariableArrayOutput `pulumi:"variables"`
}

// NewIApp registers a new resource with the given unique name, arguments, and options.
func NewIApp(ctx *pulumi.Context,
	name string, args *IAppArgs, opts ...pulumi.ResourceOption) (*IApp, error) {
	if args == nil {
		args = &IAppArgs{}
	}

	var resource IApp
	err := ctx.RegisterResource("f5bigip:sys/iApp:IApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIApp gets an existing IApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAppState, opts ...pulumi.ResourceOption) (*IApp, error) {
	var resource IApp
	err := ctx.ReadResource("f5bigip:sys/iApp:IApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IApp resources.
type iappState struct {
	// Address of the Iapp which needs to be Iappensed
	Description *string `pulumi:"description"`
	// BIG-IP password
	Devicegroup *string `pulumi:"devicegroup"`
	// BIG-IP password
	ExecuteAction *string `pulumi:"executeAction"`
	// BIG-IP password
	InheritedDevicegroup *string `pulumi:"inheritedDevicegroup"`
	// BIG-IP password
	InheritedTrafficGroup *string `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile  *string        `pulumi:"jsonfile"`
	Lists     []IAppList     `pulumi:"lists"`
	Metadatas []IAppMetadata `pulumi:"metadatas"`
	// Name of the iApp.
	Name *string `pulumi:"name"`
	// Address of the Iapp which needs to be Iappensed
	Partition *string `pulumi:"partition"`
	// BIG-IP password
	StrictUpdates *string     `pulumi:"strictUpdates"`
	Tables        []IAppTable `pulumi:"tables"`
	// BIG-IP password
	Template *string `pulumi:"template"`
	// BIG-IP password
	TemplateModified *string `pulumi:"templateModified"`
	// BIG-IP password
	TemplatePrerequisiteErrors *string `pulumi:"templatePrerequisiteErrors"`
	// BIG-IP password
	TrafficGroup *string        `pulumi:"trafficGroup"`
	Variables    []IAppVariable `pulumi:"variables"`
}

type IAppState struct {
	// Address of the Iapp which needs to be Iappensed
	Description pulumi.StringPtrInput
	// BIG-IP password
	Devicegroup pulumi.StringPtrInput
	// BIG-IP password
	ExecuteAction pulumi.StringPtrInput
	// BIG-IP password
	InheritedDevicegroup pulumi.StringPtrInput
	// BIG-IP password
	InheritedTrafficGroup pulumi.StringPtrInput
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile  pulumi.StringPtrInput
	Lists     IAppListArrayInput
	Metadatas IAppMetadataArrayInput
	// Name of the iApp.
	Name pulumi.StringPtrInput
	// Address of the Iapp which needs to be Iappensed
	Partition pulumi.StringPtrInput
	// BIG-IP password
	StrictUpdates pulumi.StringPtrInput
	Tables        IAppTableArrayInput
	// BIG-IP password
	Template pulumi.StringPtrInput
	// BIG-IP password
	TemplateModified pulumi.StringPtrInput
	// BIG-IP password
	TemplatePrerequisiteErrors pulumi.StringPtrInput
	// BIG-IP password
	TrafficGroup pulumi.StringPtrInput
	Variables    IAppVariableArrayInput
}

func (IAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*iappState)(nil)).Elem()
}

type iappArgs struct {
	// Address of the Iapp which needs to be Iappensed
	Description *string `pulumi:"description"`
	// BIG-IP password
	Devicegroup *string `pulumi:"devicegroup"`
	// BIG-IP password
	ExecuteAction *string `pulumi:"executeAction"`
	// BIG-IP password
	InheritedDevicegroup *string `pulumi:"inheritedDevicegroup"`
	// BIG-IP password
	InheritedTrafficGroup *string `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile  *string        `pulumi:"jsonfile"`
	Lists     []IAppList     `pulumi:"lists"`
	Metadatas []IAppMetadata `pulumi:"metadatas"`
	// Name of the iApp.
	Name *string `pulumi:"name"`
	// Address of the Iapp which needs to be Iappensed
	Partition *string `pulumi:"partition"`
	// BIG-IP password
	StrictUpdates *string     `pulumi:"strictUpdates"`
	Tables        []IAppTable `pulumi:"tables"`
	// BIG-IP password
	Template *string `pulumi:"template"`
	// BIG-IP password
	TemplateModified *string `pulumi:"templateModified"`
	// BIG-IP password
	TemplatePrerequisiteErrors *string `pulumi:"templatePrerequisiteErrors"`
	// BIG-IP password
	TrafficGroup *string        `pulumi:"trafficGroup"`
	Variables    []IAppVariable `pulumi:"variables"`
}

// The set of arguments for constructing a IApp resource.
type IAppArgs struct {
	// Address of the Iapp which needs to be Iappensed
	Description pulumi.StringPtrInput
	// BIG-IP password
	Devicegroup pulumi.StringPtrInput
	// BIG-IP password
	ExecuteAction pulumi.StringPtrInput
	// BIG-IP password
	InheritedDevicegroup pulumi.StringPtrInput
	// BIG-IP password
	InheritedTrafficGroup pulumi.StringPtrInput
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile  pulumi.StringPtrInput
	Lists     IAppListArrayInput
	Metadatas IAppMetadataArrayInput
	// Name of the iApp.
	Name pulumi.StringPtrInput
	// Address of the Iapp which needs to be Iappensed
	Partition pulumi.StringPtrInput
	// BIG-IP password
	StrictUpdates pulumi.StringPtrInput
	Tables        IAppTableArrayInput
	// BIG-IP password
	Template pulumi.StringPtrInput
	// BIG-IP password
	TemplateModified pulumi.StringPtrInput
	// BIG-IP password
	TemplatePrerequisiteErrors pulumi.StringPtrInput
	// BIG-IP password
	TrafficGroup pulumi.StringPtrInput
	Variables    IAppVariableArrayInput
}

func (IAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iappArgs)(nil)).Elem()
}

type IAppInput interface {
	pulumi.Input

	ToIAppOutput() IAppOutput
	ToIAppOutputWithContext(ctx context.Context) IAppOutput
}

func (*IApp) ElementType() reflect.Type {
	return reflect.TypeOf((*IApp)(nil))
}

func (i *IApp) ToIAppOutput() IAppOutput {
	return i.ToIAppOutputWithContext(context.Background())
}

func (i *IApp) ToIAppOutputWithContext(ctx context.Context) IAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAppOutput)
}

type IAppOutput struct {
	*pulumi.OutputState
}

func (IAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IApp)(nil))
}

func (o IAppOutput) ToIAppOutput() IAppOutput {
	return o
}

func (o IAppOutput) ToIAppOutputWithContext(ctx context.Context) IAppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IAppOutput{})
}
