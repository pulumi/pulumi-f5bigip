// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `BigIqAs3` provides details about bigiq as3 resource
//
// This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			// Example Usage for json file
//			_, err := f5bigip.NewBigIqAs3(ctx, "exampletask", &f5bigip.BigIqAs3Args{
//				As3Json:       readFileOrPanic("bigiq_example.json"),
//				BigiqAddress:  pulumi.String("xx.xx.xxx.xx"),
//				BigiqPassword: pulumi.String("xxxxxxxxx"),
//				BigiqUser:     pulumi.String("xxxxx"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type BigIqAs3 struct {
	pulumi.CustomResourceState

	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringOutput `pulumi:"as3Json"`
	// Address of the BIG-IQ to which your targer BIG-IP is attached
	BigiqAddress pulumi.StringOutput `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrOutput `pulumi:"bigiqLoginRef"`
	// Password of the BIG-IQ to which your targer BIG-IP is attached
	BigiqPassword pulumi.StringOutput `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrOutput `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrOutput `pulumi:"bigiqTokenAuth"`
	// User name  of the BIG-IQ to which your targer BIG-IP is attached
	BigiqUser pulumi.StringOutput `pulumi:"bigiqUser"`
	// Set True if you want to ignore metadata changes during update. By default it is set to `true`
	//
	// * `bigiq_example.json` - Example  AS3 Declarative JSON file
	//
	// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
	//
	// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
	IgnoreMetadata pulumi.BoolPtrOutput `pulumi:"ignoreMetadata"`
	// Name of Tenant
	TenantList pulumi.StringOutput `pulumi:"tenantList"`
}

// NewBigIqAs3 registers a new resource with the given unique name, arguments, and options.
func NewBigIqAs3(ctx *pulumi.Context,
	name string, args *BigIqAs3Args, opts ...pulumi.ResourceOption) (*BigIqAs3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.As3Json == nil {
		return nil, errors.New("invalid value for required argument 'As3Json'")
	}
	if args.BigiqAddress == nil {
		return nil, errors.New("invalid value for required argument 'BigiqAddress'")
	}
	if args.BigiqPassword == nil {
		return nil, errors.New("invalid value for required argument 'BigiqPassword'")
	}
	if args.BigiqUser == nil {
		return nil, errors.New("invalid value for required argument 'BigiqUser'")
	}
	if args.BigiqLoginRef != nil {
		args.BigiqLoginRef = pulumi.ToSecret(args.BigiqLoginRef).(pulumi.StringPtrInput)
	}
	if args.BigiqPassword != nil {
		args.BigiqPassword = pulumi.ToSecret(args.BigiqPassword).(pulumi.StringInput)
	}
	if args.BigiqPort != nil {
		args.BigiqPort = pulumi.ToSecret(args.BigiqPort).(pulumi.StringPtrInput)
	}
	if args.BigiqTokenAuth != nil {
		args.BigiqTokenAuth = pulumi.ToSecret(args.BigiqTokenAuth).(pulumi.BoolPtrInput)
	}
	if args.BigiqUser != nil {
		args.BigiqUser = pulumi.ToSecret(args.BigiqUser).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bigiqLoginRef",
		"bigiqPassword",
		"bigiqPort",
		"bigiqTokenAuth",
		"bigiqUser",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BigIqAs3
	err := ctx.RegisterResource("f5bigip:index/bigIqAs3:BigIqAs3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigIqAs3 gets an existing BigIqAs3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigIqAs3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigIqAs3State, opts ...pulumi.ResourceOption) (*BigIqAs3, error) {
	var resource BigIqAs3
	err := ctx.ReadResource("f5bigip:index/bigIqAs3:BigIqAs3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigIqAs3 resources.
type bigIqAs3State struct {
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json *string `pulumi:"as3Json"`
	// Address of the BIG-IQ to which your targer BIG-IP is attached
	BigiqAddress *string `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef *string `pulumi:"bigiqLoginRef"`
	// Password of the BIG-IQ to which your targer BIG-IP is attached
	BigiqPassword *string `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort *string `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth *bool `pulumi:"bigiqTokenAuth"`
	// User name  of the BIG-IQ to which your targer BIG-IP is attached
	BigiqUser *string `pulumi:"bigiqUser"`
	// Set True if you want to ignore metadata changes during update. By default it is set to `true`
	//
	// * `bigiq_example.json` - Example  AS3 Declarative JSON file
	//
	// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
	//
	// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
	IgnoreMetadata *bool `pulumi:"ignoreMetadata"`
	// Name of Tenant
	TenantList *string `pulumi:"tenantList"`
}

type BigIqAs3State struct {
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringPtrInput
	// Address of the BIG-IQ to which your targer BIG-IP is attached
	BigiqAddress pulumi.StringPtrInput
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrInput
	// Password of the BIG-IQ to which your targer BIG-IP is attached
	BigiqPassword pulumi.StringPtrInput
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrInput
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrInput
	// User name  of the BIG-IQ to which your targer BIG-IP is attached
	BigiqUser pulumi.StringPtrInput
	// Set True if you want to ignore metadata changes during update. By default it is set to `true`
	//
	// * `bigiq_example.json` - Example  AS3 Declarative JSON file
	//
	// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
	//
	// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
	IgnoreMetadata pulumi.BoolPtrInput
	// Name of Tenant
	TenantList pulumi.StringPtrInput
}

func (BigIqAs3State) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIqAs3State)(nil)).Elem()
}

type bigIqAs3Args struct {
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json string `pulumi:"as3Json"`
	// Address of the BIG-IQ to which your targer BIG-IP is attached
	BigiqAddress string `pulumi:"bigiqAddress"`
	// BIGIQ Login reference for token authentication
	BigiqLoginRef *string `pulumi:"bigiqLoginRef"`
	// Password of the BIG-IQ to which your targer BIG-IP is attached
	BigiqPassword string `pulumi:"bigiqPassword"`
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort *string `pulumi:"bigiqPort"`
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth *bool `pulumi:"bigiqTokenAuth"`
	// User name  of the BIG-IQ to which your targer BIG-IP is attached
	BigiqUser string `pulumi:"bigiqUser"`
	// Set True if you want to ignore metadata changes during update. By default it is set to `true`
	//
	// * `bigiq_example.json` - Example  AS3 Declarative JSON file
	//
	// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
	//
	// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
	IgnoreMetadata *bool `pulumi:"ignoreMetadata"`
	// Name of Tenant
	TenantList *string `pulumi:"tenantList"`
}

// The set of arguments for constructing a BigIqAs3 resource.
type BigIqAs3Args struct {
	// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
	As3Json pulumi.StringInput
	// Address of the BIG-IQ to which your targer BIG-IP is attached
	BigiqAddress pulumi.StringInput
	// BIGIQ Login reference for token authentication
	BigiqLoginRef pulumi.StringPtrInput
	// Password of the BIG-IQ to which your targer BIG-IP is attached
	BigiqPassword pulumi.StringInput
	// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
	BigiqPort pulumi.StringPtrInput
	// type `bool`, if set to `true` enables Token based Authentication,default is `false`
	BigiqTokenAuth pulumi.BoolPtrInput
	// User name  of the BIG-IQ to which your targer BIG-IP is attached
	BigiqUser pulumi.StringInput
	// Set True if you want to ignore metadata changes during update. By default it is set to `true`
	//
	// * `bigiq_example.json` - Example  AS3 Declarative JSON file
	//
	// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
	//
	// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
	IgnoreMetadata pulumi.BoolPtrInput
	// Name of Tenant
	TenantList pulumi.StringPtrInput
}

func (BigIqAs3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIqAs3Args)(nil)).Elem()
}

type BigIqAs3Input interface {
	pulumi.Input

	ToBigIqAs3Output() BigIqAs3Output
	ToBigIqAs3OutputWithContext(ctx context.Context) BigIqAs3Output
}

func (*BigIqAs3) ElementType() reflect.Type {
	return reflect.TypeOf((**BigIqAs3)(nil)).Elem()
}

func (i *BigIqAs3) ToBigIqAs3Output() BigIqAs3Output {
	return i.ToBigIqAs3OutputWithContext(context.Background())
}

func (i *BigIqAs3) ToBigIqAs3OutputWithContext(ctx context.Context) BigIqAs3Output {
	return pulumi.ToOutputWithContext(ctx, i).(BigIqAs3Output)
}

// BigIqAs3ArrayInput is an input type that accepts BigIqAs3Array and BigIqAs3ArrayOutput values.
// You can construct a concrete instance of `BigIqAs3ArrayInput` via:
//
//	BigIqAs3Array{ BigIqAs3Args{...} }
type BigIqAs3ArrayInput interface {
	pulumi.Input

	ToBigIqAs3ArrayOutput() BigIqAs3ArrayOutput
	ToBigIqAs3ArrayOutputWithContext(context.Context) BigIqAs3ArrayOutput
}

type BigIqAs3Array []BigIqAs3Input

func (BigIqAs3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BigIqAs3)(nil)).Elem()
}

func (i BigIqAs3Array) ToBigIqAs3ArrayOutput() BigIqAs3ArrayOutput {
	return i.ToBigIqAs3ArrayOutputWithContext(context.Background())
}

func (i BigIqAs3Array) ToBigIqAs3ArrayOutputWithContext(ctx context.Context) BigIqAs3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigIqAs3ArrayOutput)
}

// BigIqAs3MapInput is an input type that accepts BigIqAs3Map and BigIqAs3MapOutput values.
// You can construct a concrete instance of `BigIqAs3MapInput` via:
//
//	BigIqAs3Map{ "key": BigIqAs3Args{...} }
type BigIqAs3MapInput interface {
	pulumi.Input

	ToBigIqAs3MapOutput() BigIqAs3MapOutput
	ToBigIqAs3MapOutputWithContext(context.Context) BigIqAs3MapOutput
}

type BigIqAs3Map map[string]BigIqAs3Input

func (BigIqAs3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BigIqAs3)(nil)).Elem()
}

func (i BigIqAs3Map) ToBigIqAs3MapOutput() BigIqAs3MapOutput {
	return i.ToBigIqAs3MapOutputWithContext(context.Background())
}

func (i BigIqAs3Map) ToBigIqAs3MapOutputWithContext(ctx context.Context) BigIqAs3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigIqAs3MapOutput)
}

type BigIqAs3Output struct{ *pulumi.OutputState }

func (BigIqAs3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BigIqAs3)(nil)).Elem()
}

func (o BigIqAs3Output) ToBigIqAs3Output() BigIqAs3Output {
	return o
}

func (o BigIqAs3Output) ToBigIqAs3OutputWithContext(ctx context.Context) BigIqAs3Output {
	return o
}

// Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
func (o BigIqAs3Output) As3Json() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringOutput { return v.As3Json }).(pulumi.StringOutput)
}

// Address of the BIG-IQ to which your targer BIG-IP is attached
func (o BigIqAs3Output) BigiqAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringOutput { return v.BigiqAddress }).(pulumi.StringOutput)
}

// BIGIQ Login reference for token authentication
func (o BigIqAs3Output) BigiqLoginRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringPtrOutput { return v.BigiqLoginRef }).(pulumi.StringPtrOutput)
}

// Password of the BIG-IQ to which your targer BIG-IP is attached
func (o BigIqAs3Output) BigiqPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringOutput { return v.BigiqPassword }).(pulumi.StringOutput)
}

// type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
func (o BigIqAs3Output) BigiqPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringPtrOutput { return v.BigiqPort }).(pulumi.StringPtrOutput)
}

// type `bool`, if set to `true` enables Token based Authentication,default is `false`
func (o BigIqAs3Output) BigiqTokenAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.BoolPtrOutput { return v.BigiqTokenAuth }).(pulumi.BoolPtrOutput)
}

// User name  of the BIG-IQ to which your targer BIG-IP is attached
func (o BigIqAs3Output) BigiqUser() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringOutput { return v.BigiqUser }).(pulumi.StringOutput)
}

// Set True if you want to ignore metadata changes during update. By default it is set to `true`
//
// * `bigiq_example.json` - Example  AS3 Declarative JSON file
//
// * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
//
// >  **Note:** This resource does not support `teanatFilter` parameter as BIG-IP As3 resource
func (o BigIqAs3Output) IgnoreMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.BoolPtrOutput { return v.IgnoreMetadata }).(pulumi.BoolPtrOutput)
}

// Name of Tenant
func (o BigIqAs3Output) TenantList() pulumi.StringOutput {
	return o.ApplyT(func(v *BigIqAs3) pulumi.StringOutput { return v.TenantList }).(pulumi.StringOutput)
}

type BigIqAs3ArrayOutput struct{ *pulumi.OutputState }

func (BigIqAs3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BigIqAs3)(nil)).Elem()
}

func (o BigIqAs3ArrayOutput) ToBigIqAs3ArrayOutput() BigIqAs3ArrayOutput {
	return o
}

func (o BigIqAs3ArrayOutput) ToBigIqAs3ArrayOutputWithContext(ctx context.Context) BigIqAs3ArrayOutput {
	return o
}

func (o BigIqAs3ArrayOutput) Index(i pulumi.IntInput) BigIqAs3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BigIqAs3 {
		return vs[0].([]*BigIqAs3)[vs[1].(int)]
	}).(BigIqAs3Output)
}

type BigIqAs3MapOutput struct{ *pulumi.OutputState }

func (BigIqAs3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BigIqAs3)(nil)).Elem()
}

func (o BigIqAs3MapOutput) ToBigIqAs3MapOutput() BigIqAs3MapOutput {
	return o
}

func (o BigIqAs3MapOutput) ToBigIqAs3MapOutputWithContext(ctx context.Context) BigIqAs3MapOutput {
	return o
}

func (o BigIqAs3MapOutput) MapIndex(k pulumi.StringInput) BigIqAs3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BigIqAs3 {
		return vs[0].(map[string]*BigIqAs3)[vs[1].(string)]
	}).(BigIqAs3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigIqAs3Input)(nil)).Elem(), &BigIqAs3{})
	pulumi.RegisterInputType(reflect.TypeOf((*BigIqAs3ArrayInput)(nil)).Elem(), BigIqAs3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BigIqAs3MapInput)(nil)).Elem(), BigIqAs3Map{})
	pulumi.RegisterOutputType(BigIqAs3Output{})
	pulumi.RegisterOutputType(BigIqAs3ArrayOutput{})
	pulumi.RegisterOutputType(BigIqAs3MapOutput{})
}
