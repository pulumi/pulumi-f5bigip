// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ssl.Certificate`) to get the ssl-certificate details available on BIG-IP
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ssl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := ssl.LookupCertificate(ctx, &ssl.LookupCertificateArgs{
//				Name:      "terraform_ssl_certificate",
//				Partition: "Common",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bigipSslCertificateName", test.Name)
//			return nil
//		})
//	}
//
// ```
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("f5bigip:ssl/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// Name of the ssl_certificate
	Name string `pulumi:"name"`
	// partition of the ltm ssl_certificate
	Partition string `pulumi:"partition"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	Certificate string `pulumi:"certificate"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of sslCertificate configured on bigip with full path
	Name string `pulumi:"name"`
	// Bigip partition in which ssl-certificate is configured
	Partition string `pulumi:"partition"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCertificateResultOutput, error) {
			args := v.(LookupCertificateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("f5bigip:ssl/getCertificate:getCertificate", args, LookupCertificateResultOutput{}, options).(LookupCertificateResultOutput), nil
		}).(LookupCertificateResultOutput)
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateOutputArgs struct {
	// Name of the ssl_certificate
	Name pulumi.StringInput `pulumi:"name"`
	// partition of the ltm ssl_certificate
	Partition pulumi.StringInput `pulumi:"partition"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getCertificate.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Certificate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of sslCertificate configured on bigip with full path
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Bigip partition in which ssl-certificate is configured
func (o LookupCertificateResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Partition }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
