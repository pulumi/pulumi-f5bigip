// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `sys.Dns` Configures DNS server on F5 BIG-IP
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/sys"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sys.NewDns(ctx, "dns1", &sys.DnsArgs{
// 			Description: pulumi.String("/Common/DNS1"),
// 			NameServers: pulumi.StringArray{
// 				pulumi.String("1.1.1.1"),
// 			},
// 			NumberOfDots: pulumi.Int(2),
// 			Searches: pulumi.StringArray{
// 				pulumi.String("f5.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Dns struct {
	pulumi.CustomResourceState

	// Provide description for your DNS server
	Description pulumi.StringOutput `pulumi:"description"`
	// Name or IP address of the DNS server
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots pulumi.IntPtrOutput `pulumi:"numberOfDots"`
	// Specify what domains you want to search
	Searches pulumi.StringArrayOutput `pulumi:"searches"`
}

// NewDns registers a new resource with the given unique name, arguments, and options.
func NewDns(ctx *pulumi.Context,
	name string, args *DnsArgs, opts ...pulumi.ResourceOption) (*Dns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	var resource Dns
	err := ctx.RegisterResource("f5bigip:sys/dns:Dns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDns gets an existing Dns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsState, opts ...pulumi.ResourceOption) (*Dns, error) {
	var resource Dns
	err := ctx.ReadResource("f5bigip:sys/dns:Dns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dns resources.
type dnsState struct {
	// Provide description for your DNS server
	Description *string `pulumi:"description"`
	// Name or IP address of the DNS server
	NameServers []string `pulumi:"nameServers"`
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots *int `pulumi:"numberOfDots"`
	// Specify what domains you want to search
	Searches []string `pulumi:"searches"`
}

type DnsState struct {
	// Provide description for your DNS server
	Description pulumi.StringPtrInput
	// Name or IP address of the DNS server
	NameServers pulumi.StringArrayInput
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots pulumi.IntPtrInput
	// Specify what domains you want to search
	Searches pulumi.StringArrayInput
}

func (DnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsState)(nil)).Elem()
}

type dnsArgs struct {
	// Provide description for your DNS server
	Description string `pulumi:"description"`
	// Name or IP address of the DNS server
	NameServers []string `pulumi:"nameServers"`
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots *int `pulumi:"numberOfDots"`
	// Specify what domains you want to search
	Searches []string `pulumi:"searches"`
}

// The set of arguments for constructing a Dns resource.
type DnsArgs struct {
	// Provide description for your DNS server
	Description pulumi.StringInput
	// Name or IP address of the DNS server
	NameServers pulumi.StringArrayInput
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots pulumi.IntPtrInput
	// Specify what domains you want to search
	Searches pulumi.StringArrayInput
}

func (DnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsArgs)(nil)).Elem()
}

type DnsInput interface {
	pulumi.Input

	ToDnsOutput() DnsOutput
	ToDnsOutputWithContext(ctx context.Context) DnsOutput
}

func (*Dns) ElementType() reflect.Type {
	return reflect.TypeOf((*Dns)(nil))
}

func (i *Dns) ToDnsOutput() DnsOutput {
	return i.ToDnsOutputWithContext(context.Background())
}

func (i *Dns) ToDnsOutputWithContext(ctx context.Context) DnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsOutput)
}

func (i *Dns) ToDnsPtrOutput() DnsPtrOutput {
	return i.ToDnsPtrOutputWithContext(context.Background())
}

func (i *Dns) ToDnsPtrOutputWithContext(ctx context.Context) DnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPtrOutput)
}

type DnsPtrInput interface {
	pulumi.Input

	ToDnsPtrOutput() DnsPtrOutput
	ToDnsPtrOutputWithContext(ctx context.Context) DnsPtrOutput
}

type dnsPtrType DnsArgs

func (*dnsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dns)(nil))
}

func (i *dnsPtrType) ToDnsPtrOutput() DnsPtrOutput {
	return i.ToDnsPtrOutputWithContext(context.Background())
}

func (i *dnsPtrType) ToDnsPtrOutputWithContext(ctx context.Context) DnsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPtrOutput)
}

// DnsArrayInput is an input type that accepts DnsArray and DnsArrayOutput values.
// You can construct a concrete instance of `DnsArrayInput` via:
//
//          DnsArray{ DnsArgs{...} }
type DnsArrayInput interface {
	pulumi.Input

	ToDnsArrayOutput() DnsArrayOutput
	ToDnsArrayOutputWithContext(context.Context) DnsArrayOutput
}

type DnsArray []DnsInput

func (DnsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Dns)(nil))
}

func (i DnsArray) ToDnsArrayOutput() DnsArrayOutput {
	return i.ToDnsArrayOutputWithContext(context.Background())
}

func (i DnsArray) ToDnsArrayOutputWithContext(ctx context.Context) DnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsArrayOutput)
}

// DnsMapInput is an input type that accepts DnsMap and DnsMapOutput values.
// You can construct a concrete instance of `DnsMapInput` via:
//
//          DnsMap{ "key": DnsArgs{...} }
type DnsMapInput interface {
	pulumi.Input

	ToDnsMapOutput() DnsMapOutput
	ToDnsMapOutputWithContext(context.Context) DnsMapOutput
}

type DnsMap map[string]DnsInput

func (DnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Dns)(nil))
}

func (i DnsMap) ToDnsMapOutput() DnsMapOutput {
	return i.ToDnsMapOutputWithContext(context.Background())
}

func (i DnsMap) ToDnsMapOutputWithContext(ctx context.Context) DnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsMapOutput)
}

type DnsOutput struct {
	*pulumi.OutputState
}

func (DnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dns)(nil))
}

func (o DnsOutput) ToDnsOutput() DnsOutput {
	return o
}

func (o DnsOutput) ToDnsOutputWithContext(ctx context.Context) DnsOutput {
	return o
}

func (o DnsOutput) ToDnsPtrOutput() DnsPtrOutput {
	return o.ToDnsPtrOutputWithContext(context.Background())
}

func (o DnsOutput) ToDnsPtrOutputWithContext(ctx context.Context) DnsPtrOutput {
	return o.ApplyT(func(v Dns) *Dns {
		return &v
	}).(DnsPtrOutput)
}

type DnsPtrOutput struct {
	*pulumi.OutputState
}

func (DnsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dns)(nil))
}

func (o DnsPtrOutput) ToDnsPtrOutput() DnsPtrOutput {
	return o
}

func (o DnsPtrOutput) ToDnsPtrOutputWithContext(ctx context.Context) DnsPtrOutput {
	return o
}

type DnsArrayOutput struct{ *pulumi.OutputState }

func (DnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dns)(nil))
}

func (o DnsArrayOutput) ToDnsArrayOutput() DnsArrayOutput {
	return o
}

func (o DnsArrayOutput) ToDnsArrayOutputWithContext(ctx context.Context) DnsArrayOutput {
	return o
}

func (o DnsArrayOutput) Index(i pulumi.IntInput) DnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dns {
		return vs[0].([]Dns)[vs[1].(int)]
	}).(DnsOutput)
}

type DnsMapOutput struct{ *pulumi.OutputState }

func (DnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Dns)(nil))
}

func (o DnsMapOutput) ToDnsMapOutput() DnsMapOutput {
	return o
}

func (o DnsMapOutput) ToDnsMapOutputWithContext(ctx context.Context) DnsMapOutput {
	return o
}

func (o DnsMapOutput) MapIndex(k pulumi.StringInput) DnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Dns {
		return vs[0].(map[string]Dns)[vs[1].(string)]
	}).(DnsOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsOutput{})
	pulumi.RegisterOutputType(DnsPtrOutput{})
	pulumi.RegisterOutputType(DnsArrayOutput{})
	pulumi.RegisterOutputType(DnsMapOutput{})
}
