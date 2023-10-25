// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.
type IApp struct {
	pulumi.CustomResourceState

	// User defined description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// BIG-IP password
	Devicegroup pulumi.StringOutput `pulumi:"devicegroup"`
	// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
	ExecuteAction pulumi.StringOutput `pulumi:"executeAction"`
	// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
	InheritedDevicegroup pulumi.StringPtrOutput `pulumi:"inheritedDevicegroup"`
	// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
	InheritedTrafficGroup pulumi.StringPtrOutput `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile pulumi.StringOutput `pulumi:"jsonfile"`
	// string values
	Lists IAppListArrayOutput `pulumi:"lists"`
	// User defined generic data for the application service. It is a name and value pair.
	Metadatas IAppMetadataArrayOutput `pulumi:"metadatas"`
	// Name of the iApp.
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the administrative partition within which the application resides.
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
	StrictUpdates pulumi.StringPtrOutput `pulumi:"strictUpdates"`
	Tables        IAppTableArrayOutput   `pulumi:"tables"`
	// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
	TemplateModified pulumi.StringPtrOutput `pulumi:"templateModified"`
	// Indicates any missing prerequisites associated with the template that defines this application.
	TemplatePrerequisiteErrors pulumi.StringPtrOutput `pulumi:"templatePrerequisiteErrors"`
	// The name of the traffic group that the application service is assigned to.
	TrafficGroup pulumi.StringPtrOutput  `pulumi:"trafficGroup"`
	Variables    IAppVariableArrayOutput `pulumi:"variables"`
}

// NewIApp registers a new resource with the given unique name, arguments, and options.
func NewIApp(ctx *pulumi.Context,
	name string, args *IAppArgs, opts ...pulumi.ResourceOption) (*IApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Jsonfile == nil {
		return nil, errors.New("invalid value for required argument 'Jsonfile'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// User defined description.
	Description *string `pulumi:"description"`
	// BIG-IP password
	Devicegroup *string `pulumi:"devicegroup"`
	// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
	ExecuteAction *string `pulumi:"executeAction"`
	// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
	InheritedDevicegroup *string `pulumi:"inheritedDevicegroup"`
	// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
	InheritedTrafficGroup *string `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile *string `pulumi:"jsonfile"`
	// string values
	Lists []IAppList `pulumi:"lists"`
	// User defined generic data for the application service. It is a name and value pair.
	Metadatas []IAppMetadata `pulumi:"metadatas"`
	// Name of the iApp.
	Name *string `pulumi:"name"`
	// Displays the administrative partition within which the application resides.
	Partition *string `pulumi:"partition"`
	// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
	StrictUpdates *string     `pulumi:"strictUpdates"`
	Tables        []IAppTable `pulumi:"tables"`
	// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
	Template *string `pulumi:"template"`
	// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
	TemplateModified *string `pulumi:"templateModified"`
	// Indicates any missing prerequisites associated with the template that defines this application.
	TemplatePrerequisiteErrors *string `pulumi:"templatePrerequisiteErrors"`
	// The name of the traffic group that the application service is assigned to.
	TrafficGroup *string        `pulumi:"trafficGroup"`
	Variables    []IAppVariable `pulumi:"variables"`
}

type IAppState struct {
	// User defined description.
	Description pulumi.StringPtrInput
	// BIG-IP password
	Devicegroup pulumi.StringPtrInput
	// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
	ExecuteAction pulumi.StringPtrInput
	// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
	InheritedDevicegroup pulumi.StringPtrInput
	// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
	InheritedTrafficGroup pulumi.StringPtrInput
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile pulumi.StringPtrInput
	// string values
	Lists IAppListArrayInput
	// User defined generic data for the application service. It is a name and value pair.
	Metadatas IAppMetadataArrayInput
	// Name of the iApp.
	Name pulumi.StringPtrInput
	// Displays the administrative partition within which the application resides.
	Partition pulumi.StringPtrInput
	// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
	StrictUpdates pulumi.StringPtrInput
	Tables        IAppTableArrayInput
	// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
	Template pulumi.StringPtrInput
	// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
	TemplateModified pulumi.StringPtrInput
	// Indicates any missing prerequisites associated with the template that defines this application.
	TemplatePrerequisiteErrors pulumi.StringPtrInput
	// The name of the traffic group that the application service is assigned to.
	TrafficGroup pulumi.StringPtrInput
	Variables    IAppVariableArrayInput
}

func (IAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*iappState)(nil)).Elem()
}

type iappArgs struct {
	// User defined description.
	Description *string `pulumi:"description"`
	// BIG-IP password
	Devicegroup *string `pulumi:"devicegroup"`
	// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
	ExecuteAction *string `pulumi:"executeAction"`
	// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
	InheritedDevicegroup *string `pulumi:"inheritedDevicegroup"`
	// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
	InheritedTrafficGroup *string `pulumi:"inheritedTrafficGroup"`
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile string `pulumi:"jsonfile"`
	// string values
	Lists []IAppList `pulumi:"lists"`
	// User defined generic data for the application service. It is a name and value pair.
	Metadatas []IAppMetadata `pulumi:"metadatas"`
	// Name of the iApp.
	Name string `pulumi:"name"`
	// Displays the administrative partition within which the application resides.
	Partition *string `pulumi:"partition"`
	// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
	StrictUpdates *string     `pulumi:"strictUpdates"`
	Tables        []IAppTable `pulumi:"tables"`
	// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
	Template *string `pulumi:"template"`
	// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
	TemplateModified *string `pulumi:"templateModified"`
	// Indicates any missing prerequisites associated with the template that defines this application.
	TemplatePrerequisiteErrors *string `pulumi:"templatePrerequisiteErrors"`
	// The name of the traffic group that the application service is assigned to.
	TrafficGroup *string        `pulumi:"trafficGroup"`
	Variables    []IAppVariable `pulumi:"variables"`
}

// The set of arguments for constructing a IApp resource.
type IAppArgs struct {
	// User defined description.
	Description pulumi.StringPtrInput
	// BIG-IP password
	Devicegroup pulumi.StringPtrInput
	// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
	ExecuteAction pulumi.StringPtrInput
	// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
	InheritedDevicegroup pulumi.StringPtrInput
	// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
	InheritedTrafficGroup pulumi.StringPtrInput
	// Refer to the Json file which will be deployed on F5 BIG-IP.
	Jsonfile pulumi.StringInput
	// string values
	Lists IAppListArrayInput
	// User defined generic data for the application service. It is a name and value pair.
	Metadatas IAppMetadataArrayInput
	// Name of the iApp.
	Name pulumi.StringInput
	// Displays the administrative partition within which the application resides.
	Partition pulumi.StringPtrInput
	// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
	StrictUpdates pulumi.StringPtrInput
	Tables        IAppTableArrayInput
	// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
	Template pulumi.StringPtrInput
	// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
	TemplateModified pulumi.StringPtrInput
	// Indicates any missing prerequisites associated with the template that defines this application.
	TemplatePrerequisiteErrors pulumi.StringPtrInput
	// The name of the traffic group that the application service is assigned to.
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
	return reflect.TypeOf((**IApp)(nil)).Elem()
}

func (i *IApp) ToIAppOutput() IAppOutput {
	return i.ToIAppOutputWithContext(context.Background())
}

func (i *IApp) ToIAppOutputWithContext(ctx context.Context) IAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAppOutput)
}

func (i *IApp) ToOutput(ctx context.Context) pulumix.Output[*IApp] {
	return pulumix.Output[*IApp]{
		OutputState: i.ToIAppOutputWithContext(ctx).OutputState,
	}
}

// IAppArrayInput is an input type that accepts IAppArray and IAppArrayOutput values.
// You can construct a concrete instance of `IAppArrayInput` via:
//
//	IAppArray{ IAppArgs{...} }
type IAppArrayInput interface {
	pulumi.Input

	ToIAppArrayOutput() IAppArrayOutput
	ToIAppArrayOutputWithContext(context.Context) IAppArrayOutput
}

type IAppArray []IAppInput

func (IAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IApp)(nil)).Elem()
}

func (i IAppArray) ToIAppArrayOutput() IAppArrayOutput {
	return i.ToIAppArrayOutputWithContext(context.Background())
}

func (i IAppArray) ToIAppArrayOutputWithContext(ctx context.Context) IAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAppArrayOutput)
}

func (i IAppArray) ToOutput(ctx context.Context) pulumix.Output[[]*IApp] {
	return pulumix.Output[[]*IApp]{
		OutputState: i.ToIAppArrayOutputWithContext(ctx).OutputState,
	}
}

// IAppMapInput is an input type that accepts IAppMap and IAppMapOutput values.
// You can construct a concrete instance of `IAppMapInput` via:
//
//	IAppMap{ "key": IAppArgs{...} }
type IAppMapInput interface {
	pulumi.Input

	ToIAppMapOutput() IAppMapOutput
	ToIAppMapOutputWithContext(context.Context) IAppMapOutput
}

type IAppMap map[string]IAppInput

func (IAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IApp)(nil)).Elem()
}

func (i IAppMap) ToIAppMapOutput() IAppMapOutput {
	return i.ToIAppMapOutputWithContext(context.Background())
}

func (i IAppMap) ToIAppMapOutputWithContext(ctx context.Context) IAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAppMapOutput)
}

func (i IAppMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IApp] {
	return pulumix.Output[map[string]*IApp]{
		OutputState: i.ToIAppMapOutputWithContext(ctx).OutputState,
	}
}

type IAppOutput struct{ *pulumi.OutputState }

func (IAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IApp)(nil)).Elem()
}

func (o IAppOutput) ToIAppOutput() IAppOutput {
	return o
}

func (o IAppOutput) ToIAppOutputWithContext(ctx context.Context) IAppOutput {
	return o
}

func (o IAppOutput) ToOutput(ctx context.Context) pulumix.Output[*IApp] {
	return pulumix.Output[*IApp]{
		OutputState: o.OutputState,
	}
}

// User defined description.
func (o IAppOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// BIG-IP password
func (o IAppOutput) Devicegroup() pulumi.StringOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringOutput { return v.Devicegroup }).(pulumi.StringOutput)
}

// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
func (o IAppOutput) ExecuteAction() pulumi.StringOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringOutput { return v.ExecuteAction }).(pulumi.StringOutput)
}

// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
func (o IAppOutput) InheritedDevicegroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.InheritedDevicegroup }).(pulumi.StringPtrOutput)
}

// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
func (o IAppOutput) InheritedTrafficGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.InheritedTrafficGroup }).(pulumi.StringPtrOutput)
}

// Refer to the Json file which will be deployed on F5 BIG-IP.
func (o IAppOutput) Jsonfile() pulumi.StringOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringOutput { return v.Jsonfile }).(pulumi.StringOutput)
}

// string values
func (o IAppOutput) Lists() IAppListArrayOutput {
	return o.ApplyT(func(v *IApp) IAppListArrayOutput { return v.Lists }).(IAppListArrayOutput)
}

// User defined generic data for the application service. It is a name and value pair.
func (o IAppOutput) Metadatas() IAppMetadataArrayOutput {
	return o.ApplyT(func(v *IApp) IAppMetadataArrayOutput { return v.Metadatas }).(IAppMetadataArrayOutput)
}

// Name of the iApp.
func (o IAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Displays the administrative partition within which the application resides.
func (o IAppOutput) Partition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.Partition }).(pulumi.StringPtrOutput)
}

// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
func (o IAppOutput) StrictUpdates() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.StrictUpdates }).(pulumi.StringPtrOutput)
}

func (o IAppOutput) Tables() IAppTableArrayOutput {
	return o.ApplyT(func(v *IApp) IAppTableArrayOutput { return v.Tables }).(IAppTableArrayOutput)
}

// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
func (o IAppOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
func (o IAppOutput) TemplateModified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.TemplateModified }).(pulumi.StringPtrOutput)
}

// Indicates any missing prerequisites associated with the template that defines this application.
func (o IAppOutput) TemplatePrerequisiteErrors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.TemplatePrerequisiteErrors }).(pulumi.StringPtrOutput)
}

// The name of the traffic group that the application service is assigned to.
func (o IAppOutput) TrafficGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IApp) pulumi.StringPtrOutput { return v.TrafficGroup }).(pulumi.StringPtrOutput)
}

func (o IAppOutput) Variables() IAppVariableArrayOutput {
	return o.ApplyT(func(v *IApp) IAppVariableArrayOutput { return v.Variables }).(IAppVariableArrayOutput)
}

type IAppArrayOutput struct{ *pulumi.OutputState }

func (IAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IApp)(nil)).Elem()
}

func (o IAppArrayOutput) ToIAppArrayOutput() IAppArrayOutput {
	return o
}

func (o IAppArrayOutput) ToIAppArrayOutputWithContext(ctx context.Context) IAppArrayOutput {
	return o
}

func (o IAppArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IApp] {
	return pulumix.Output[[]*IApp]{
		OutputState: o.OutputState,
	}
}

func (o IAppArrayOutput) Index(i pulumi.IntInput) IAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IApp {
		return vs[0].([]*IApp)[vs[1].(int)]
	}).(IAppOutput)
}

type IAppMapOutput struct{ *pulumi.OutputState }

func (IAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IApp)(nil)).Elem()
}

func (o IAppMapOutput) ToIAppMapOutput() IAppMapOutput {
	return o
}

func (o IAppMapOutput) ToIAppMapOutputWithContext(ctx context.Context) IAppMapOutput {
	return o
}

func (o IAppMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IApp] {
	return pulumix.Output[map[string]*IApp]{
		OutputState: o.OutputState,
	}
}

func (o IAppMapOutput) MapIndex(k pulumi.StringInput) IAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IApp {
		return vs[0].(map[string]*IApp)[vs[1].(string)]
	}).(IAppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAppInput)(nil)).Elem(), &IApp{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAppArrayInput)(nil)).Elem(), IAppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAppMapInput)(nil)).Elem(), IAppMap{})
	pulumi.RegisterOutputType(IAppOutput{})
	pulumi.RegisterOutputType(IAppArrayOutput{})
	pulumi.RegisterOutputType(IAppMapOutput{})
}
