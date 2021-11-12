// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewProfileHttpCompress(ctx, "sjhttpcompression", &ltm.ProfileHttpCompressArgs{
// 			ContentTypeExcludes: pulumi.StringArray{
// 				pulumi.String("nicecontentexclude.com"),
// 			},
// 			ContentTypeIncludes: pulumi.StringArray{
// 				pulumi.String("nicecontent.com"),
// 			},
// 			DefaultsFrom: pulumi.String("/Common/httpcompression"),
// 			Name:         pulumi.String("/Common/sjhttpcompression2"),
// 			UriExcludes: pulumi.StringArray{
// 				pulumi.String("www.abc.f5.com"),
// 				pulumi.String("www.abc2.f5.com"),
// 			},
// 			UriIncludes: pulumi.StringArray{
// 				pulumi.String("www.xyzbc.cisco.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProfileHttpCompress struct {
	pulumi.CustomResourceState

	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExcludes pulumi.StringArrayOutput `pulumi:"contentTypeExcludes"`
	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeIncludes pulumi.StringArrayOutput `pulumi:"contentTypeIncludes"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Name of the profile_httpcompress
	Name pulumi.StringOutput `pulumi:"name"`
	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExcludes pulumi.StringArrayOutput `pulumi:"uriExcludes"`
	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriIncludes pulumi.StringArrayOutput `pulumi:"uriIncludes"`
}

// NewProfileHttpCompress registers a new resource with the given unique name, arguments, and options.
func NewProfileHttpCompress(ctx *pulumi.Context,
	name string, args *ProfileHttpCompressArgs, opts ...pulumi.ResourceOption) (*ProfileHttpCompress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource ProfileHttpCompress
	err := ctx.RegisterResource("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileHttpCompress gets an existing ProfileHttpCompress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileHttpCompress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileHttpCompressState, opts ...pulumi.ResourceOption) (*ProfileHttpCompress, error) {
	var resource ProfileHttpCompress
	err := ctx.ReadResource("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileHttpCompress resources.
type profileHttpCompressState struct {
	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExcludes []string `pulumi:"contentTypeExcludes"`
	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeIncludes []string `pulumi:"contentTypeIncludes"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Name of the profile_httpcompress
	Name *string `pulumi:"name"`
	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExcludes []string `pulumi:"uriExcludes"`
	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriIncludes []string `pulumi:"uriIncludes"`
}

type ProfileHttpCompressState struct {
	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExcludes pulumi.StringArrayInput
	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeIncludes pulumi.StringArrayInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Name of the profile_httpcompress
	Name pulumi.StringPtrInput
	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExcludes pulumi.StringArrayInput
	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriIncludes pulumi.StringArrayInput
}

func (ProfileHttpCompressState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileHttpCompressState)(nil)).Elem()
}

type profileHttpCompressArgs struct {
	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExcludes []string `pulumi:"contentTypeExcludes"`
	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeIncludes []string `pulumi:"contentTypeIncludes"`
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Name of the profile_httpcompress
	Name string `pulumi:"name"`
	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExcludes []string `pulumi:"uriExcludes"`
	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriIncludes []string `pulumi:"uriIncludes"`
}

// The set of arguments for constructing a ProfileHttpCompress resource.
type ProfileHttpCompressArgs struct {
	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExcludes pulumi.StringArrayInput
	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeIncludes pulumi.StringArrayInput
	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom pulumi.StringPtrInput
	// Name of the profile_httpcompress
	Name pulumi.StringInput
	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExcludes pulumi.StringArrayInput
	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriIncludes pulumi.StringArrayInput
}

func (ProfileHttpCompressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileHttpCompressArgs)(nil)).Elem()
}

type ProfileHttpCompressInput interface {
	pulumi.Input

	ToProfileHttpCompressOutput() ProfileHttpCompressOutput
	ToProfileHttpCompressOutputWithContext(ctx context.Context) ProfileHttpCompressOutput
}

func (*ProfileHttpCompress) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileHttpCompress)(nil))
}

func (i *ProfileHttpCompress) ToProfileHttpCompressOutput() ProfileHttpCompressOutput {
	return i.ToProfileHttpCompressOutputWithContext(context.Background())
}

func (i *ProfileHttpCompress) ToProfileHttpCompressOutputWithContext(ctx context.Context) ProfileHttpCompressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressOutput)
}

func (i *ProfileHttpCompress) ToProfileHttpCompressPtrOutput() ProfileHttpCompressPtrOutput {
	return i.ToProfileHttpCompressPtrOutputWithContext(context.Background())
}

func (i *ProfileHttpCompress) ToProfileHttpCompressPtrOutputWithContext(ctx context.Context) ProfileHttpCompressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressPtrOutput)
}

type ProfileHttpCompressPtrInput interface {
	pulumi.Input

	ToProfileHttpCompressPtrOutput() ProfileHttpCompressPtrOutput
	ToProfileHttpCompressPtrOutputWithContext(ctx context.Context) ProfileHttpCompressPtrOutput
}

type profileHttpCompressPtrType ProfileHttpCompressArgs

func (*profileHttpCompressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileHttpCompress)(nil))
}

func (i *profileHttpCompressPtrType) ToProfileHttpCompressPtrOutput() ProfileHttpCompressPtrOutput {
	return i.ToProfileHttpCompressPtrOutputWithContext(context.Background())
}

func (i *profileHttpCompressPtrType) ToProfileHttpCompressPtrOutputWithContext(ctx context.Context) ProfileHttpCompressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressPtrOutput)
}

// ProfileHttpCompressArrayInput is an input type that accepts ProfileHttpCompressArray and ProfileHttpCompressArrayOutput values.
// You can construct a concrete instance of `ProfileHttpCompressArrayInput` via:
//
//          ProfileHttpCompressArray{ ProfileHttpCompressArgs{...} }
type ProfileHttpCompressArrayInput interface {
	pulumi.Input

	ToProfileHttpCompressArrayOutput() ProfileHttpCompressArrayOutput
	ToProfileHttpCompressArrayOutputWithContext(context.Context) ProfileHttpCompressArrayOutput
}

type ProfileHttpCompressArray []ProfileHttpCompressInput

func (ProfileHttpCompressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileHttpCompress)(nil)).Elem()
}

func (i ProfileHttpCompressArray) ToProfileHttpCompressArrayOutput() ProfileHttpCompressArrayOutput {
	return i.ToProfileHttpCompressArrayOutputWithContext(context.Background())
}

func (i ProfileHttpCompressArray) ToProfileHttpCompressArrayOutputWithContext(ctx context.Context) ProfileHttpCompressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressArrayOutput)
}

// ProfileHttpCompressMapInput is an input type that accepts ProfileHttpCompressMap and ProfileHttpCompressMapOutput values.
// You can construct a concrete instance of `ProfileHttpCompressMapInput` via:
//
//          ProfileHttpCompressMap{ "key": ProfileHttpCompressArgs{...} }
type ProfileHttpCompressMapInput interface {
	pulumi.Input

	ToProfileHttpCompressMapOutput() ProfileHttpCompressMapOutput
	ToProfileHttpCompressMapOutputWithContext(context.Context) ProfileHttpCompressMapOutput
}

type ProfileHttpCompressMap map[string]ProfileHttpCompressInput

func (ProfileHttpCompressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileHttpCompress)(nil)).Elem()
}

func (i ProfileHttpCompressMap) ToProfileHttpCompressMapOutput() ProfileHttpCompressMapOutput {
	return i.ToProfileHttpCompressMapOutputWithContext(context.Background())
}

func (i ProfileHttpCompressMap) ToProfileHttpCompressMapOutputWithContext(ctx context.Context) ProfileHttpCompressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressMapOutput)
}

type ProfileHttpCompressOutput struct{ *pulumi.OutputState }

func (ProfileHttpCompressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileHttpCompress)(nil))
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressOutput() ProfileHttpCompressOutput {
	return o
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressOutputWithContext(ctx context.Context) ProfileHttpCompressOutput {
	return o
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressPtrOutput() ProfileHttpCompressPtrOutput {
	return o.ToProfileHttpCompressPtrOutputWithContext(context.Background())
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressPtrOutputWithContext(ctx context.Context) ProfileHttpCompressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileHttpCompress) *ProfileHttpCompress {
		return &v
	}).(ProfileHttpCompressPtrOutput)
}

type ProfileHttpCompressPtrOutput struct{ *pulumi.OutputState }

func (ProfileHttpCompressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileHttpCompress)(nil))
}

func (o ProfileHttpCompressPtrOutput) ToProfileHttpCompressPtrOutput() ProfileHttpCompressPtrOutput {
	return o
}

func (o ProfileHttpCompressPtrOutput) ToProfileHttpCompressPtrOutputWithContext(ctx context.Context) ProfileHttpCompressPtrOutput {
	return o
}

func (o ProfileHttpCompressPtrOutput) Elem() ProfileHttpCompressOutput {
	return o.ApplyT(func(v *ProfileHttpCompress) ProfileHttpCompress {
		if v != nil {
			return *v
		}
		var ret ProfileHttpCompress
		return ret
	}).(ProfileHttpCompressOutput)
}

type ProfileHttpCompressArrayOutput struct{ *pulumi.OutputState }

func (ProfileHttpCompressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileHttpCompress)(nil))
}

func (o ProfileHttpCompressArrayOutput) ToProfileHttpCompressArrayOutput() ProfileHttpCompressArrayOutput {
	return o
}

func (o ProfileHttpCompressArrayOutput) ToProfileHttpCompressArrayOutputWithContext(ctx context.Context) ProfileHttpCompressArrayOutput {
	return o
}

func (o ProfileHttpCompressArrayOutput) Index(i pulumi.IntInput) ProfileHttpCompressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileHttpCompress {
		return vs[0].([]ProfileHttpCompress)[vs[1].(int)]
	}).(ProfileHttpCompressOutput)
}

type ProfileHttpCompressMapOutput struct{ *pulumi.OutputState }

func (ProfileHttpCompressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProfileHttpCompress)(nil))
}

func (o ProfileHttpCompressMapOutput) ToProfileHttpCompressMapOutput() ProfileHttpCompressMapOutput {
	return o
}

func (o ProfileHttpCompressMapOutput) ToProfileHttpCompressMapOutputWithContext(ctx context.Context) ProfileHttpCompressMapOutput {
	return o
}

func (o ProfileHttpCompressMapOutput) MapIndex(k pulumi.StringInput) ProfileHttpCompressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProfileHttpCompress {
		return vs[0].(map[string]ProfileHttpCompress)[vs[1].(string)]
	}).(ProfileHttpCompressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileHttpCompressInput)(nil)).Elem(), &ProfileHttpCompress{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileHttpCompressPtrInput)(nil)).Elem(), &ProfileHttpCompress{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileHttpCompressArrayInput)(nil)).Elem(), ProfileHttpCompressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileHttpCompressMapInput)(nil)).Elem(), ProfileHttpCompressMap{})
	pulumi.RegisterOutputType(ProfileHttpCompressOutput{})
	pulumi.RegisterOutputType(ProfileHttpCompressPtrOutput{})
	pulumi.RegisterOutputType(ProfileHttpCompressArrayOutput{})
	pulumi.RegisterOutputType(ProfileHttpCompressMapOutput{})
}
