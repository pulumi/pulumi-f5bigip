// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

func (ProfileHttpCompress) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileHttpCompress)(nil)).Elem()
}

func (i ProfileHttpCompress) ToProfileHttpCompressOutput() ProfileHttpCompressOutput {
	return i.ToProfileHttpCompressOutputWithContext(context.Background())
}

func (i ProfileHttpCompress) ToProfileHttpCompressOutputWithContext(ctx context.Context) ProfileHttpCompressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileHttpCompressOutput)
}

type ProfileHttpCompressOutput struct {
	*pulumi.OutputState
}

func (ProfileHttpCompressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileHttpCompressOutput)(nil)).Elem()
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressOutput() ProfileHttpCompressOutput {
	return o
}

func (o ProfileHttpCompressOutput) ToProfileHttpCompressOutputWithContext(ctx context.Context) ProfileHttpCompressOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfileHttpCompressOutput{})
}
