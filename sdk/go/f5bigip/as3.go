// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `As3` provides details about bigip as3 resource
//
// This resource is helpful to configure as3 declarative JSON on BIG-IP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := f5bigip.NewAs3(ctx, "as3-example1As3", &f5bigip.As3Args{
//				As3Json: readFileOrPanic("example1.json"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = f5bigip.NewAs3(ctx, "as3-example1Index/as3As3", &f5bigip.As3Args{
//				As3Json:      readFileOrPanic("example2.json"),
//				TenantFilter: pulumi.String("Sample_03"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// As3 resources can be imported using the partition name, e.g., ( use comma separated partition names if there are multiple partitions in as3 deployments )
//
// ```sh
//
//	$ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01
//
// ```
//
// ```sh
//
//	$ pulumi import f5bigip:index/as3:As3 bigip_as3.test Sample_http_01,Sample_non_http_01
//
// ```
//
// #### Import examples ( single and multiple partitions )
//
// ```sh
//
//	$ pulumi import f5bigip:index/as3:As3 test Sample_http_01
//
// ```
//
//	bigip_as3.testImporting from ID "Sample_http_01"... bigip_as3.testImport prepared!
//
//	Prepared bigip_as3 for import bigip_as3.testRefreshing state... [id=Sample_http_01] Import successful! The resources that were imported are shown above. These resources are now in your Terraform state and will henceforth be managed by Terraform. $ terraform show bigip_as3.testresource "bigip_as3" "test" {
//
//	as3_json
//
// = jsonencode(
//
//	{
//
//	action
//
// = "deploy"
//
//	class
//
//	= "AS3"
//
//	declaration = {
//
//	Sample_http_01 = {
//
//	A1
//
// = {
//
//	class
//
// = "Application"
//
//	jsessionid = {
//
//	class
//
//	= "Persist"
//
//	cookieMethod
//
// = "hash"
//
//	cookieName
//
// = "JSESSIONID"
//
//	persistenceMethod = "cookie"
//
//	}
//
//	service
//
// = {
//
//	class
//
// = "Service_HTTP"
//
//	persistenceMethods = [
//
//	{
//
//	use = "jsessionid"
//
//	},
//
//	]
//
//	pool
//
//	= "web_pool"
//
//	virtualAddresses
//
//	= [
//
//	"10.0.2.10",
//
//	]
//
//	}
//
//	web_pool
//
//	= {
//
//	class
//
// = "Pool"
//
//	members
//
// = [
//
//	{
//
//	serverAddresses = [
//
//	"192.0.2.10",
//
//	"192.0.2.11",
//
//	]
//
//	servicePort
//
//	= 80
//
//	},
//
//	]
//
//	monitors = [
//
//	"http",
//
//	]
//
//	}
//
//	}
//
//	class = "Tenant"
//
//	}
//
//	class
//
// = "ADC"
//
//	id
//
//	= "UDP_DNS_Sample"
//
//	label
//
// = "UDP_DNS_Sample"
//
//	remark
//
//	= "Sample of a UDP DNS Load Balancer Service"
//
//	schemaVersion
//
// = "3.0.0"
//
//	}
//
//	persist
//
//	= true
//
//	}
//
//	)
//
//	id
//
// = "Sample_http_01"
//
//	tenant_filter = "Sample_http_01"
//
//	tenant_list
//
//	= "Sample_http_01" }
//
// ```sh
//
//	$ pulumi import f5bigip:index/as3:As3 test Sample_http_01,Sample_non_http_01
//
// ```
//
//	bigip_as3.testImporting from ID "Sample_http_01,Sample_non_http_01"... bigip_as3.testImport prepared!
//
//	Prepared bigip_as3 for import bigip_as3.testRefreshing state... [id=Sample_http_01,Sample_non_http_01] Import successful! The resources that were imported are shown above. These resources are now in your Terraform state and will henceforth be managed by Terraform. $ terraform show bigip_as3.testresource "bigip_as3" "test" {
//
//	as3_json
//
// = jsonencode(
//
//	{
//
//	action
//
// = "deploy"
//
//	class
//
//	= "AS3"
//
//	declaration = {
//
//	Sample_http_01
//
//	= {
//
//	A1
//
// = {
//
//	class
//
// = "Application"
//
//	jsessionid = {
//
//	class
//
//	= "Persist"
//
//	cookieMethod
//
// = "hash"
//
//	cookieName
//
// = "JSESSIONID"
//
//	persistenceMethod = "cookie"
//
//	}
//
//	service
//
// = {
//
//	class
//
// = "Service_HTTP"
//
//	persistenceMethods = [
//
//	{
//
//	use = "jsessionid"
//
//	},
//
//	]
//
//	pool
//
//	= "web_pool"
//
//	virtualAddresses
//
//	= [
//
//	"10.0.2.10",
//
//	]
//
//	}
//
//	web_pool
//
//	= {
//
//	class
//
// = "Pool"
//
//	members
//
// = [
//
//	{
//
//	serverAddresses = [
//
//	"192.0.2.10",
//
//	"192.0.2.11",
//
//	]
//
//	servicePort
//
//	= 80
//
//	},
//
//	]
//
//	monitors = [
//
//	"http",
//
//	]
//
//	}
//
//	}
//
//	class = "Tenant"
//
//	}
//
//	Sample_non_http_01 = {
//
//	DNS_Service = {
//
//	Pool1
//
//	= {
//
//	class
//
// = "Pool"
//
//	members
//
// = [
//
//	{
//
//	serverAddresses = [
//
//	"10.1.10.100",
//
//	]
//
//	servicePort
//
//	= 53
//
//	},
//
//	{
//
//	serverAddresses = [
//
//	"10.1.10.101",
//
//	]
//
//	servicePort
//
//	= 53
//
//	},
//
//	]
//
//	monitors = [
//
//	"icmp",
//
//	]
//
//	}
//
//	class
//
//	= "Application"
//
//	service = {
//
//	class
//
// = "Service_UDP"
//
//	pool
//
//	= "Pool1"
//
//	virtualAddresses = [
//
//	"10.1.20.121",
//
//	]
//
//	virtualPort
//
// = 53
//
//	}
//
//	}
//
//	class
//
//	= "Tenant"
//
//	}
//
//	class
//
// = "ADC"
//
//	id
//
//	= "UDP_DNS_Sample"
//
//	label
//
// = "UDP_DNS_Sample"
//
//	remark
//
//	= "Sample of a UDP DNS Load Balancer Service"
//
//	schemaVersion
//
// = "3.0.0"
//
//	}
//
//	persist
//
//	= true
//
//	}
//
//	)
//
//	id
//
// = "Sample_http_01,Sample_non_http_01"
//
//	tenant_filter = "Sample_http_01,Sample_non_http_01"
//
//	tenant_list
//
//	= "Sample_http_01,Sample_non_http_01" } * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/composing-a-declaration.html
type As3 struct {
	pulumi.CustomResourceState

	// Name of Application
	ApplicationList pulumi.StringOutput `pulumi:"applicationList"`
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringPtrOutput `pulumi:"as3Json"`
	// Set True if you want to ignore metadata changes during update. By default it is set to false
	IgnoreMetadata pulumi.BoolPtrOutput `pulumi:"ignoreMetadata"`
	// ID of AS3 post declaration async task
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
	TenantFilter pulumi.StringPtrOutput `pulumi:"tenantFilter"`
	// Name of Tenant
	TenantList pulumi.StringOutput `pulumi:"tenantList"`
	// Name of Tenant
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrOutput `pulumi:"tenantName"`
}

// NewAs3 registers a new resource with the given unique name, arguments, and options.
func NewAs3(ctx *pulumi.Context,
	name string, args *As3Args, opts ...pulumi.ResourceOption) (*As3, error) {
	if args == nil {
		args = &As3Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource As3
	err := ctx.RegisterResource("f5bigip:index/as3:As3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAs3 gets an existing As3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAs3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *As3State, opts ...pulumi.ResourceOption) (*As3, error) {
	var resource As3
	err := ctx.ReadResource("f5bigip:index/as3:As3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering As3 resources.
type as3State struct {
	// Name of Application
	ApplicationList *string `pulumi:"applicationList"`
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json *string `pulumi:"as3Json"`
	// Set True if you want to ignore metadata changes during update. By default it is set to false
	IgnoreMetadata *bool `pulumi:"ignoreMetadata"`
	// ID of AS3 post declaration async task
	TaskId *string `pulumi:"taskId"`
	// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
	TenantFilter *string `pulumi:"tenantFilter"`
	// Name of Tenant
	TenantList *string `pulumi:"tenantList"`
	// Name of Tenant
	//
	// Deprecated: this attribute is no longer in use
	TenantName *string `pulumi:"tenantName"`
}

type As3State struct {
	// Name of Application
	ApplicationList pulumi.StringPtrInput
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringPtrInput
	// Set True if you want to ignore metadata changes during update. By default it is set to false
	IgnoreMetadata pulumi.BoolPtrInput
	// ID of AS3 post declaration async task
	TaskId pulumi.StringPtrInput
	// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
	TenantFilter pulumi.StringPtrInput
	// Name of Tenant
	TenantList pulumi.StringPtrInput
	// Name of Tenant
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrInput
}

func (As3State) ElementType() reflect.Type {
	return reflect.TypeOf((*as3State)(nil)).Elem()
}

type as3Args struct {
	// Name of Application
	ApplicationList *string `pulumi:"applicationList"`
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json *string `pulumi:"as3Json"`
	// Set True if you want to ignore metadata changes during update. By default it is set to false
	IgnoreMetadata *bool `pulumi:"ignoreMetadata"`
	// ID of AS3 post declaration async task
	TaskId *string `pulumi:"taskId"`
	// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
	TenantFilter *string `pulumi:"tenantFilter"`
	// Name of Tenant
	TenantList *string `pulumi:"tenantList"`
	// Name of Tenant
	//
	// Deprecated: this attribute is no longer in use
	TenantName *string `pulumi:"tenantName"`
}

// The set of arguments for constructing a As3 resource.
type As3Args struct {
	// Name of Application
	ApplicationList pulumi.StringPtrInput
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringPtrInput
	// Set True if you want to ignore metadata changes during update. By default it is set to false
	IgnoreMetadata pulumi.BoolPtrInput
	// ID of AS3 post declaration async task
	TaskId pulumi.StringPtrInput
	// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
	TenantFilter pulumi.StringPtrInput
	// Name of Tenant
	TenantList pulumi.StringPtrInput
	// Name of Tenant
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrInput
}

func (As3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*as3Args)(nil)).Elem()
}

type As3Input interface {
	pulumi.Input

	ToAs3Output() As3Output
	ToAs3OutputWithContext(ctx context.Context) As3Output
}

func (*As3) ElementType() reflect.Type {
	return reflect.TypeOf((**As3)(nil)).Elem()
}

func (i *As3) ToAs3Output() As3Output {
	return i.ToAs3OutputWithContext(context.Background())
}

func (i *As3) ToAs3OutputWithContext(ctx context.Context) As3Output {
	return pulumi.ToOutputWithContext(ctx, i).(As3Output)
}

// As3ArrayInput is an input type that accepts As3Array and As3ArrayOutput values.
// You can construct a concrete instance of `As3ArrayInput` via:
//
//	As3Array{ As3Args{...} }
type As3ArrayInput interface {
	pulumi.Input

	ToAs3ArrayOutput() As3ArrayOutput
	ToAs3ArrayOutputWithContext(context.Context) As3ArrayOutput
}

type As3Array []As3Input

func (As3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*As3)(nil)).Elem()
}

func (i As3Array) ToAs3ArrayOutput() As3ArrayOutput {
	return i.ToAs3ArrayOutputWithContext(context.Background())
}

func (i As3Array) ToAs3ArrayOutputWithContext(ctx context.Context) As3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(As3ArrayOutput)
}

// As3MapInput is an input type that accepts As3Map and As3MapOutput values.
// You can construct a concrete instance of `As3MapInput` via:
//
//	As3Map{ "key": As3Args{...} }
type As3MapInput interface {
	pulumi.Input

	ToAs3MapOutput() As3MapOutput
	ToAs3MapOutputWithContext(context.Context) As3MapOutput
}

type As3Map map[string]As3Input

func (As3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*As3)(nil)).Elem()
}

func (i As3Map) ToAs3MapOutput() As3MapOutput {
	return i.ToAs3MapOutputWithContext(context.Background())
}

func (i As3Map) ToAs3MapOutputWithContext(ctx context.Context) As3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(As3MapOutput)
}

type As3Output struct{ *pulumi.OutputState }

func (As3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**As3)(nil)).Elem()
}

func (o As3Output) ToAs3Output() As3Output {
	return o
}

func (o As3Output) ToAs3OutputWithContext(ctx context.Context) As3Output {
	return o
}

// Name of Application
func (o As3Output) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v *As3) pulumi.StringOutput { return v.ApplicationList }).(pulumi.StringOutput)
}

// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
func (o As3Output) As3Json() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *As3) pulumi.StringPtrOutput { return v.As3Json }).(pulumi.StringPtrOutput)
}

// Set True if you want to ignore metadata changes during update. By default it is set to false
func (o As3Output) IgnoreMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *As3) pulumi.BoolPtrOutput { return v.IgnoreMetadata }).(pulumi.BoolPtrOutput)
}

// ID of AS3 post declaration async task
func (o As3Output) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *As3) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// If there are multiple tenants on a BIG-IP, this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified.
func (o As3Output) TenantFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *As3) pulumi.StringPtrOutput { return v.TenantFilter }).(pulumi.StringPtrOutput)
}

// Name of Tenant
func (o As3Output) TenantList() pulumi.StringOutput {
	return o.ApplyT(func(v *As3) pulumi.StringOutput { return v.TenantList }).(pulumi.StringOutput)
}

// Name of Tenant
//
// Deprecated: this attribute is no longer in use
func (o As3Output) TenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *As3) pulumi.StringPtrOutput { return v.TenantName }).(pulumi.StringPtrOutput)
}

type As3ArrayOutput struct{ *pulumi.OutputState }

func (As3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*As3)(nil)).Elem()
}

func (o As3ArrayOutput) ToAs3ArrayOutput() As3ArrayOutput {
	return o
}

func (o As3ArrayOutput) ToAs3ArrayOutputWithContext(ctx context.Context) As3ArrayOutput {
	return o
}

func (o As3ArrayOutput) Index(i pulumi.IntInput) As3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *As3 {
		return vs[0].([]*As3)[vs[1].(int)]
	}).(As3Output)
}

type As3MapOutput struct{ *pulumi.OutputState }

func (As3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*As3)(nil)).Elem()
}

func (o As3MapOutput) ToAs3MapOutput() As3MapOutput {
	return o
}

func (o As3MapOutput) ToAs3MapOutputWithContext(ctx context.Context) As3MapOutput {
	return o
}

func (o As3MapOutput) MapIndex(k pulumi.StringInput) As3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *As3 {
		return vs[0].(map[string]*As3)[vs[1].(string)]
	}).(As3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*As3Input)(nil)).Elem(), &As3{})
	pulumi.RegisterInputType(reflect.TypeOf((*As3ArrayInput)(nil)).Elem(), As3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*As3MapInput)(nil)).Elem(), As3Map{})
	pulumi.RegisterOutputType(As3Output{})
	pulumi.RegisterOutputType(As3ArrayOutput{})
	pulumi.RegisterOutputType(As3MapOutput{})
}
