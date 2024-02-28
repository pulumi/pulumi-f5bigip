// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `bigipLtmRewriteProfile` Configures ltm policies to manage traffic assigned to a virtual server
//
// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-profile`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ltm.NewProfileRewrite(ctx, "test-profile", &ltm.ProfileRewriteArgs{
//				BypassLists: pulumi.StringArray{
//					pulumi.String("http://notouch.com"),
//				},
//				CaFile:       pulumi.String("/Common/ca-bundle.crt"),
//				CacheType:    pulumi.String("cache-img-css-js"),
//				CrlFile:      pulumi.String("none"),
//				DefaultsFrom: pulumi.String("/Common/rewrite"),
//				Name:         pulumi.String("/Common/tf_profile"),
//				RewriteLists: pulumi.StringArray{
//					pulumi.String("http://some.com"),
//				},
//				RewriteMode:    pulumi.String("portal"),
//				SigningCert:    pulumi.String("/Common/default.crt"),
//				SigningKey:     pulumi.String("/Common/default.key"),
//				SplitTunneling: pulumi.String("true"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ltm.NewProfileRewrite(ctx, "test-profile2", &ltm.ProfileRewriteArgs{
//				CookieRules: ltm.ProfileRewriteCookieRuleArray{
//					&ltm.ProfileRewriteCookieRuleArgs{
//						ClientDomain: pulumi.String("wrong.com"),
//						ClientPath:   pulumi.String("/this/"),
//						RuleName:     pulumi.String("cookie1"),
//						ServerDomain: pulumi.String("wrong.com"),
//						ServerPath:   pulumi.String("/this/"),
//					},
//					&ltm.ProfileRewriteCookieRuleArgs{
//						ClientDomain: pulumi.String("incorrect.com"),
//						ClientPath:   pulumi.String("/this/"),
//						RuleName:     pulumi.String("cookie2"),
//						ServerDomain: pulumi.String("absolute.com"),
//						ServerPath:   pulumi.String("/this/"),
//					},
//				},
//				DefaultsFrom: pulumi.String("/Common/rewrite"),
//				Name:         pulumi.String("/Common/tf_profile_translate"),
//				Requests: ltm.ProfileRewriteRequestArray{
//					&ltm.ProfileRewriteRequestArgs{
//						InsertXfwdFor:      pulumi.String("enabled"),
//						InsertXfwdHost:     pulumi.String("disabled"),
//						InsertXfwdProtocol: pulumi.String("enabled"),
//						RewriteHeaders:     pulumi.String("disabled"),
//					},
//				},
//				Responses: ltm.ProfileRewriteResponseArray{
//					&ltm.ProfileRewriteResponseArgs{
//						RewriteContent: pulumi.String("enabled"),
//						RewriteHeaders: pulumi.String("disabled"),
//					},
//				},
//				RewriteMode: pulumi.String("uri-translation"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ProfileRewrite struct {
	pulumi.CustomResourceState

	// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
	BypassLists pulumi.StringArrayOutput `pulumi:"bypassLists"`
	// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
	CaFile pulumi.StringOutput `pulumi:"caFile"`
	// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
	CacheType pulumi.StringPtrOutput `pulumi:"cacheType"`
	// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
	CookieRules ProfileRewriteCookieRuleArrayOutput `pulumi:"cookieRules"`
	// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
	CrlFile pulumi.StringPtrOutput `pulumi:"crlFile"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
	Name pulumi.StringOutput `pulumi:"name"`
	// Block type. Each request is block type with following arguments.
	Requests ProfileRewriteRequestArrayOutput `pulumi:"requests"`
	// Block type. Each request is block type with following arguments.
	Responses ProfileRewriteResponseArrayOutput `pulumi:"responses"`
	// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
	RewriteLists pulumi.StringArrayOutput `pulumi:"rewriteLists"`
	// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
	RewriteMode pulumi.StringOutput `pulumi:"rewriteMode"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
	SigningCert pulumi.StringOutput `pulumi:"signingCert"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
	SigningKey pulumi.StringOutput `pulumi:"signingKey"`
	// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
	SigningKeyPassword pulumi.StringOutput `pulumi:"signingKeyPassword"`
	// Specifies the type of Client caching. Valid choices are: `true, false`
	SplitTunneling pulumi.StringOutput `pulumi:"splitTunneling"`
}

// NewProfileRewrite registers a new resource with the given unique name, arguments, and options.
func NewProfileRewrite(ctx *pulumi.Context,
	name string, args *ProfileRewriteArgs, opts ...pulumi.ResourceOption) (*ProfileRewrite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.RewriteMode == nil {
		return nil, errors.New("invalid value for required argument 'RewriteMode'")
	}
	if args.SigningKeyPassword != nil {
		args.SigningKeyPassword = pulumi.ToSecret(args.SigningKeyPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"signingKeyPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProfileRewrite
	err := ctx.RegisterResource("f5bigip:ltm/profileRewrite:ProfileRewrite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileRewrite gets an existing ProfileRewrite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileRewrite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileRewriteState, opts ...pulumi.ResourceOption) (*ProfileRewrite, error) {
	var resource ProfileRewrite
	err := ctx.ReadResource("f5bigip:ltm/profileRewrite:ProfileRewrite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileRewrite resources.
type profileRewriteState struct {
	// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
	BypassLists []string `pulumi:"bypassLists"`
	// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
	CaFile *string `pulumi:"caFile"`
	// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
	CacheType *string `pulumi:"cacheType"`
	// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
	CookieRules []ProfileRewriteCookieRule `pulumi:"cookieRules"`
	// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
	CrlFile *string `pulumi:"crlFile"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
	Name *string `pulumi:"name"`
	// Block type. Each request is block type with following arguments.
	Requests []ProfileRewriteRequest `pulumi:"requests"`
	// Block type. Each request is block type with following arguments.
	Responses []ProfileRewriteResponse `pulumi:"responses"`
	// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
	RewriteLists []string `pulumi:"rewriteLists"`
	// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
	RewriteMode *string `pulumi:"rewriteMode"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
	SigningCert *string `pulumi:"signingCert"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
	SigningKey *string `pulumi:"signingKey"`
	// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
	SigningKeyPassword *string `pulumi:"signingKeyPassword"`
	// Specifies the type of Client caching. Valid choices are: `true, false`
	SplitTunneling *string `pulumi:"splitTunneling"`
}

type ProfileRewriteState struct {
	// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
	BypassLists pulumi.StringArrayInput
	// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
	CaFile pulumi.StringPtrInput
	// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
	CacheType pulumi.StringPtrInput
	// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
	CookieRules ProfileRewriteCookieRuleArrayInput
	// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
	CrlFile pulumi.StringPtrInput
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
	DefaultsFrom pulumi.StringPtrInput
	// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
	Name pulumi.StringPtrInput
	// Block type. Each request is block type with following arguments.
	Requests ProfileRewriteRequestArrayInput
	// Block type. Each request is block type with following arguments.
	Responses ProfileRewriteResponseArrayInput
	// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
	RewriteLists pulumi.StringArrayInput
	// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
	RewriteMode pulumi.StringPtrInput
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
	SigningCert pulumi.StringPtrInput
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
	SigningKey pulumi.StringPtrInput
	// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
	SigningKeyPassword pulumi.StringPtrInput
	// Specifies the type of Client caching. Valid choices are: `true, false`
	SplitTunneling pulumi.StringPtrInput
}

func (ProfileRewriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileRewriteState)(nil)).Elem()
}

type profileRewriteArgs struct {
	// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
	BypassLists []string `pulumi:"bypassLists"`
	// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
	CaFile *string `pulumi:"caFile"`
	// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
	CacheType *string `pulumi:"cacheType"`
	// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
	CookieRules []ProfileRewriteCookieRule `pulumi:"cookieRules"`
	// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
	CrlFile *string `pulumi:"crlFile"`
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
	Name string `pulumi:"name"`
	// Block type. Each request is block type with following arguments.
	Requests []ProfileRewriteRequest `pulumi:"requests"`
	// Block type. Each request is block type with following arguments.
	Responses []ProfileRewriteResponse `pulumi:"responses"`
	// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
	RewriteLists []string `pulumi:"rewriteLists"`
	// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
	RewriteMode string `pulumi:"rewriteMode"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
	SigningCert *string `pulumi:"signingCert"`
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
	SigningKey *string `pulumi:"signingKey"`
	// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
	SigningKeyPassword *string `pulumi:"signingKeyPassword"`
	// Specifies the type of Client caching. Valid choices are: `true, false`
	SplitTunneling *string `pulumi:"splitTunneling"`
}

// The set of arguments for constructing a ProfileRewrite resource.
type ProfileRewriteArgs struct {
	// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
	BypassLists pulumi.StringArrayInput
	// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
	CaFile pulumi.StringPtrInput
	// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
	CacheType pulumi.StringPtrInput
	// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
	CookieRules ProfileRewriteCookieRuleArrayInput
	// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
	CrlFile pulumi.StringPtrInput
	// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
	DefaultsFrom pulumi.StringPtrInput
	// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
	Name pulumi.StringInput
	// Block type. Each request is block type with following arguments.
	Requests ProfileRewriteRequestArrayInput
	// Block type. Each request is block type with following arguments.
	Responses ProfileRewriteResponseArrayInput
	// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
	RewriteLists pulumi.StringArrayInput
	// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
	RewriteMode pulumi.StringInput
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
	SigningCert pulumi.StringPtrInput
	// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
	SigningKey pulumi.StringPtrInput
	// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
	SigningKeyPassword pulumi.StringPtrInput
	// Specifies the type of Client caching. Valid choices are: `true, false`
	SplitTunneling pulumi.StringPtrInput
}

func (ProfileRewriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileRewriteArgs)(nil)).Elem()
}

type ProfileRewriteInput interface {
	pulumi.Input

	ToProfileRewriteOutput() ProfileRewriteOutput
	ToProfileRewriteOutputWithContext(ctx context.Context) ProfileRewriteOutput
}

func (*ProfileRewrite) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileRewrite)(nil)).Elem()
}

func (i *ProfileRewrite) ToProfileRewriteOutput() ProfileRewriteOutput {
	return i.ToProfileRewriteOutputWithContext(context.Background())
}

func (i *ProfileRewrite) ToProfileRewriteOutputWithContext(ctx context.Context) ProfileRewriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRewriteOutput)
}

// ProfileRewriteArrayInput is an input type that accepts ProfileRewriteArray and ProfileRewriteArrayOutput values.
// You can construct a concrete instance of `ProfileRewriteArrayInput` via:
//
//	ProfileRewriteArray{ ProfileRewriteArgs{...} }
type ProfileRewriteArrayInput interface {
	pulumi.Input

	ToProfileRewriteArrayOutput() ProfileRewriteArrayOutput
	ToProfileRewriteArrayOutputWithContext(context.Context) ProfileRewriteArrayOutput
}

type ProfileRewriteArray []ProfileRewriteInput

func (ProfileRewriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileRewrite)(nil)).Elem()
}

func (i ProfileRewriteArray) ToProfileRewriteArrayOutput() ProfileRewriteArrayOutput {
	return i.ToProfileRewriteArrayOutputWithContext(context.Background())
}

func (i ProfileRewriteArray) ToProfileRewriteArrayOutputWithContext(ctx context.Context) ProfileRewriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRewriteArrayOutput)
}

// ProfileRewriteMapInput is an input type that accepts ProfileRewriteMap and ProfileRewriteMapOutput values.
// You can construct a concrete instance of `ProfileRewriteMapInput` via:
//
//	ProfileRewriteMap{ "key": ProfileRewriteArgs{...} }
type ProfileRewriteMapInput interface {
	pulumi.Input

	ToProfileRewriteMapOutput() ProfileRewriteMapOutput
	ToProfileRewriteMapOutputWithContext(context.Context) ProfileRewriteMapOutput
}

type ProfileRewriteMap map[string]ProfileRewriteInput

func (ProfileRewriteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileRewrite)(nil)).Elem()
}

func (i ProfileRewriteMap) ToProfileRewriteMapOutput() ProfileRewriteMapOutput {
	return i.ToProfileRewriteMapOutputWithContext(context.Background())
}

func (i ProfileRewriteMap) ToProfileRewriteMapOutputWithContext(ctx context.Context) ProfileRewriteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileRewriteMapOutput)
}

type ProfileRewriteOutput struct{ *pulumi.OutputState }

func (ProfileRewriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileRewrite)(nil)).Elem()
}

func (o ProfileRewriteOutput) ToProfileRewriteOutput() ProfileRewriteOutput {
	return o
}

func (o ProfileRewriteOutput) ToProfileRewriteOutputWithContext(ctx context.Context) ProfileRewriteOutput {
	return o
}

// Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
func (o ProfileRewriteOutput) BypassLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringArrayOutput { return v.BypassLists }).(pulumi.StringArrayOutput)
}

// Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
func (o ProfileRewriteOutput) CaFile() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.CaFile }).(pulumi.StringOutput)
}

// Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
func (o ProfileRewriteOutput) CacheType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringPtrOutput { return v.CacheType }).(pulumi.StringPtrOutput)
}

// Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
func (o ProfileRewriteOutput) CookieRules() ProfileRewriteCookieRuleArrayOutput {
	return o.ApplyT(func(v *ProfileRewrite) ProfileRewriteCookieRuleArrayOutput { return v.CookieRules }).(ProfileRewriteCookieRuleArrayOutput)
}

// Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
func (o ProfileRewriteOutput) CrlFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringPtrOutput { return v.CrlFile }).(pulumi.StringPtrOutput)
}

// Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
func (o ProfileRewriteOutput) DefaultsFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringPtrOutput { return v.DefaultsFrom }).(pulumi.StringPtrOutput)
}

// Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
func (o ProfileRewriteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Block type. Each request is block type with following arguments.
func (o ProfileRewriteOutput) Requests() ProfileRewriteRequestArrayOutput {
	return o.ApplyT(func(v *ProfileRewrite) ProfileRewriteRequestArrayOutput { return v.Requests }).(ProfileRewriteRequestArrayOutput)
}

// Block type. Each request is block type with following arguments.
func (o ProfileRewriteOutput) Responses() ProfileRewriteResponseArrayOutput {
	return o.ApplyT(func(v *ProfileRewrite) ProfileRewriteResponseArrayOutput { return v.Responses }).(ProfileRewriteResponseArrayOutput)
}

// Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
func (o ProfileRewriteOutput) RewriteLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringArrayOutput { return v.RewriteLists }).(pulumi.StringArrayOutput)
}

// Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
func (o ProfileRewriteOutput) RewriteMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.RewriteMode }).(pulumi.StringOutput)
}

// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
func (o ProfileRewriteOutput) SigningCert() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.SigningCert }).(pulumi.StringOutput)
}

// Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
func (o ProfileRewriteOutput) SigningKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.SigningKey }).(pulumi.StringOutput)
}

// Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
func (o ProfileRewriteOutput) SigningKeyPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.SigningKeyPassword }).(pulumi.StringOutput)
}

// Specifies the type of Client caching. Valid choices are: `true, false`
func (o ProfileRewriteOutput) SplitTunneling() pulumi.StringOutput {
	return o.ApplyT(func(v *ProfileRewrite) pulumi.StringOutput { return v.SplitTunneling }).(pulumi.StringOutput)
}

type ProfileRewriteArrayOutput struct{ *pulumi.OutputState }

func (ProfileRewriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProfileRewrite)(nil)).Elem()
}

func (o ProfileRewriteArrayOutput) ToProfileRewriteArrayOutput() ProfileRewriteArrayOutput {
	return o
}

func (o ProfileRewriteArrayOutput) ToProfileRewriteArrayOutputWithContext(ctx context.Context) ProfileRewriteArrayOutput {
	return o
}

func (o ProfileRewriteArrayOutput) Index(i pulumi.IntInput) ProfileRewriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProfileRewrite {
		return vs[0].([]*ProfileRewrite)[vs[1].(int)]
	}).(ProfileRewriteOutput)
}

type ProfileRewriteMapOutput struct{ *pulumi.OutputState }

func (ProfileRewriteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProfileRewrite)(nil)).Elem()
}

func (o ProfileRewriteMapOutput) ToProfileRewriteMapOutput() ProfileRewriteMapOutput {
	return o
}

func (o ProfileRewriteMapOutput) ToProfileRewriteMapOutputWithContext(ctx context.Context) ProfileRewriteMapOutput {
	return o
}

func (o ProfileRewriteMapOutput) MapIndex(k pulumi.StringInput) ProfileRewriteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProfileRewrite {
		return vs[0].(map[string]*ProfileRewrite)[vs[1].(string)]
	}).(ProfileRewriteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRewriteInput)(nil)).Elem(), &ProfileRewrite{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRewriteArrayInput)(nil)).Elem(), ProfileRewriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileRewriteMapInput)(nil)).Elem(), ProfileRewriteMap{})
	pulumi.RegisterOutputType(ProfileRewriteOutput{})
	pulumi.RegisterOutputType(ProfileRewriteArrayOutput{})
	pulumi.RegisterOutputType(ProfileRewriteMapOutput{})
}