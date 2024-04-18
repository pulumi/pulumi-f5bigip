// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ssl.Certificate` This resource will import SSL certificates on BIG-IP LTM.
// Certificates can be imported from certificate files on the local disk, in PEM format
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ssl"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "servercert.crt",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ssl.NewCertificate(ctx, "test-cert", &ssl.CertificateArgs{
//				Name:      pulumi.String("servercert.crt"),
//				Content:   invokeFile.Result,
//				Partition: pulumi.String("Common"),
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
type Certificate struct {
	pulumi.CustomResourceState

	// Content of certificate on Disk
	Content pulumi.StringOutput `pulumi:"content"`
	// Full Path Name of ssl certificate
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// Specifies the issuer certificate.
	IssuerCert pulumi.StringPtrOutput `pulumi:"issuerCert"`
	// Specifies the type of monitoring used.
	MonitoringType pulumi.StringPtrOutput `pulumi:"monitoringType"`
	// Name of the SSL Certificate to be Imported on to BIGIP
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the OCSP responder.
	Ocsp pulumi.StringPtrOutput `pulumi:"ocsp"`
	// Partition of ssl certificate
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Content != nil {
		args.Content = pulumi.ToSecret(args.Content).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"content",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("f5bigip:ssl/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("f5bigip:ssl/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Content of certificate on Disk
	Content *string `pulumi:"content"`
	// Full Path Name of ssl certificate
	FullPath *string `pulumi:"fullPath"`
	// Specifies the issuer certificate.
	IssuerCert *string `pulumi:"issuerCert"`
	// Specifies the type of monitoring used.
	MonitoringType *string `pulumi:"monitoringType"`
	// Name of the SSL Certificate to be Imported on to BIGIP
	Name *string `pulumi:"name"`
	// Specifies the OCSP responder.
	Ocsp *string `pulumi:"ocsp"`
	// Partition of ssl certificate
	Partition *string `pulumi:"partition"`
}

type CertificateState struct {
	// Content of certificate on Disk
	Content pulumi.StringPtrInput
	// Full Path Name of ssl certificate
	FullPath pulumi.StringPtrInput
	// Specifies the issuer certificate.
	IssuerCert pulumi.StringPtrInput
	// Specifies the type of monitoring used.
	MonitoringType pulumi.StringPtrInput
	// Name of the SSL Certificate to be Imported on to BIGIP
	Name pulumi.StringPtrInput
	// Specifies the OCSP responder.
	Ocsp pulumi.StringPtrInput
	// Partition of ssl certificate
	Partition pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Content of certificate on Disk
	Content string `pulumi:"content"`
	// Full Path Name of ssl certificate
	FullPath *string `pulumi:"fullPath"`
	// Specifies the issuer certificate.
	IssuerCert *string `pulumi:"issuerCert"`
	// Specifies the type of monitoring used.
	MonitoringType *string `pulumi:"monitoringType"`
	// Name of the SSL Certificate to be Imported on to BIGIP
	Name string `pulumi:"name"`
	// Specifies the OCSP responder.
	Ocsp *string `pulumi:"ocsp"`
	// Partition of ssl certificate
	Partition *string `pulumi:"partition"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Content of certificate on Disk
	Content pulumi.StringInput
	// Full Path Name of ssl certificate
	FullPath pulumi.StringPtrInput
	// Specifies the issuer certificate.
	IssuerCert pulumi.StringPtrInput
	// Specifies the type of monitoring used.
	MonitoringType pulumi.StringPtrInput
	// Name of the SSL Certificate to be Imported on to BIGIP
	Name pulumi.StringInput
	// Specifies the OCSP responder.
	Ocsp pulumi.StringPtrInput
	// Partition of ssl certificate
	Partition pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Content of certificate on Disk
func (o CertificateOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Full Path Name of ssl certificate
func (o CertificateOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.FullPath }).(pulumi.StringOutput)
}

// Specifies the issuer certificate.
func (o CertificateOutput) IssuerCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.IssuerCert }).(pulumi.StringPtrOutput)
}

// Specifies the type of monitoring used.
func (o CertificateOutput) MonitoringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.MonitoringType }).(pulumi.StringPtrOutput)
}

// Name of the SSL Certificate to be Imported on to BIGIP
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the OCSP responder.
func (o CertificateOutput) Ocsp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Ocsp }).(pulumi.StringPtrOutput)
}

// Partition of ssl certificate
func (o CertificateOutput) Partition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Partition }).(pulumi.StringPtrOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
