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

// `SslKeyCert` This resource will import SSL certificate and key on BIG-IP LTM.
// The certificate and the key can be imported from files on the local disk, in PEM format
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
//			_, err := f5bigip.NewSslKeyCert(ctx, "testkeycert", &f5bigip.SslKeyCertArgs{
//				Partition:   pulumi.String("Common"),
//				KeyName:     pulumi.String("ssl-test-key"),
//				KeyContent:  readFileOrPanic("key.pem"),
//				CertName:    pulumi.String("ssl-test-cert"),
//				CertContent: readFileOrPanic("certificate.pem"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SslKeyCert struct {
	pulumi.CustomResourceState

	// The content of the cert.
	CertContent pulumi.StringOutput `pulumi:"certContent"`
	// full path of the SSL certificate on the BIGIP.
	CertFullPath pulumi.StringOutput `pulumi:"certFullPath"`
	// Name of the SSL certificate to be Imported on to BIGIP.
	CertName pulumi.StringOutput `pulumi:"certName"`
	// The content of the key.
	KeyContent pulumi.StringOutput `pulumi:"keyContent"`
	// full path of the SSL key on the BIGIP.
	KeyFullPath pulumi.StringOutput `pulumi:"keyFullPath"`
	// Name of the SSL key to be Imported on to BIGIP.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Partition on to SSL certificate and key to be imported.
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Passphrase on the SSL key.
	Passphrase pulumi.StringPtrOutput `pulumi:"passphrase"`
}

// NewSslKeyCert registers a new resource with the given unique name, arguments, and options.
func NewSslKeyCert(ctx *pulumi.Context,
	name string, args *SslKeyCertArgs, opts ...pulumi.ResourceOption) (*SslKeyCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertContent == nil {
		return nil, errors.New("invalid value for required argument 'CertContent'")
	}
	if args.CertName == nil {
		return nil, errors.New("invalid value for required argument 'CertName'")
	}
	if args.KeyContent == nil {
		return nil, errors.New("invalid value for required argument 'KeyContent'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.CertContent != nil {
		args.CertContent = pulumi.ToSecret(args.CertContent).(pulumi.StringInput)
	}
	if args.KeyContent != nil {
		args.KeyContent = pulumi.ToSecret(args.KeyContent).(pulumi.StringInput)
	}
	if args.Passphrase != nil {
		args.Passphrase = pulumi.ToSecret(args.Passphrase).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"certContent",
		"keyContent",
		"passphrase",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SslKeyCert
	err := ctx.RegisterResource("f5bigip:index/sslKeyCert:SslKeyCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslKeyCert gets an existing SslKeyCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslKeyCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslKeyCertState, opts ...pulumi.ResourceOption) (*SslKeyCert, error) {
	var resource SslKeyCert
	err := ctx.ReadResource("f5bigip:index/sslKeyCert:SslKeyCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslKeyCert resources.
type sslKeyCertState struct {
	// The content of the cert.
	CertContent *string `pulumi:"certContent"`
	// full path of the SSL certificate on the BIGIP.
	CertFullPath *string `pulumi:"certFullPath"`
	// Name of the SSL certificate to be Imported on to BIGIP.
	CertName *string `pulumi:"certName"`
	// The content of the key.
	KeyContent *string `pulumi:"keyContent"`
	// full path of the SSL key on the BIGIP.
	KeyFullPath *string `pulumi:"keyFullPath"`
	// Name of the SSL key to be Imported on to BIGIP.
	KeyName *string `pulumi:"keyName"`
	// Partition on to SSL certificate and key to be imported.
	Partition *string `pulumi:"partition"`
	// Passphrase on the SSL key.
	Passphrase *string `pulumi:"passphrase"`
}

type SslKeyCertState struct {
	// The content of the cert.
	CertContent pulumi.StringPtrInput
	// full path of the SSL certificate on the BIGIP.
	CertFullPath pulumi.StringPtrInput
	// Name of the SSL certificate to be Imported on to BIGIP.
	CertName pulumi.StringPtrInput
	// The content of the key.
	KeyContent pulumi.StringPtrInput
	// full path of the SSL key on the BIGIP.
	KeyFullPath pulumi.StringPtrInput
	// Name of the SSL key to be Imported on to BIGIP.
	KeyName pulumi.StringPtrInput
	// Partition on to SSL certificate and key to be imported.
	Partition pulumi.StringPtrInput
	// Passphrase on the SSL key.
	Passphrase pulumi.StringPtrInput
}

func (SslKeyCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslKeyCertState)(nil)).Elem()
}

type sslKeyCertArgs struct {
	// The content of the cert.
	CertContent string `pulumi:"certContent"`
	// full path of the SSL certificate on the BIGIP.
	CertFullPath *string `pulumi:"certFullPath"`
	// Name of the SSL certificate to be Imported on to BIGIP.
	CertName string `pulumi:"certName"`
	// The content of the key.
	KeyContent string `pulumi:"keyContent"`
	// full path of the SSL key on the BIGIP.
	KeyFullPath *string `pulumi:"keyFullPath"`
	// Name of the SSL key to be Imported on to BIGIP.
	KeyName string `pulumi:"keyName"`
	// Partition on to SSL certificate and key to be imported.
	Partition *string `pulumi:"partition"`
	// Passphrase on the SSL key.
	Passphrase *string `pulumi:"passphrase"`
}

// The set of arguments for constructing a SslKeyCert resource.
type SslKeyCertArgs struct {
	// The content of the cert.
	CertContent pulumi.StringInput
	// full path of the SSL certificate on the BIGIP.
	CertFullPath pulumi.StringPtrInput
	// Name of the SSL certificate to be Imported on to BIGIP.
	CertName pulumi.StringInput
	// The content of the key.
	KeyContent pulumi.StringInput
	// full path of the SSL key on the BIGIP.
	KeyFullPath pulumi.StringPtrInput
	// Name of the SSL key to be Imported on to BIGIP.
	KeyName pulumi.StringInput
	// Partition on to SSL certificate and key to be imported.
	Partition pulumi.StringPtrInput
	// Passphrase on the SSL key.
	Passphrase pulumi.StringPtrInput
}

func (SslKeyCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslKeyCertArgs)(nil)).Elem()
}

type SslKeyCertInput interface {
	pulumi.Input

	ToSslKeyCertOutput() SslKeyCertOutput
	ToSslKeyCertOutputWithContext(ctx context.Context) SslKeyCertOutput
}

func (*SslKeyCert) ElementType() reflect.Type {
	return reflect.TypeOf((**SslKeyCert)(nil)).Elem()
}

func (i *SslKeyCert) ToSslKeyCertOutput() SslKeyCertOutput {
	return i.ToSslKeyCertOutputWithContext(context.Background())
}

func (i *SslKeyCert) ToSslKeyCertOutputWithContext(ctx context.Context) SslKeyCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslKeyCertOutput)
}

// SslKeyCertArrayInput is an input type that accepts SslKeyCertArray and SslKeyCertArrayOutput values.
// You can construct a concrete instance of `SslKeyCertArrayInput` via:
//
//	SslKeyCertArray{ SslKeyCertArgs{...} }
type SslKeyCertArrayInput interface {
	pulumi.Input

	ToSslKeyCertArrayOutput() SslKeyCertArrayOutput
	ToSslKeyCertArrayOutputWithContext(context.Context) SslKeyCertArrayOutput
}

type SslKeyCertArray []SslKeyCertInput

func (SslKeyCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslKeyCert)(nil)).Elem()
}

func (i SslKeyCertArray) ToSslKeyCertArrayOutput() SslKeyCertArrayOutput {
	return i.ToSslKeyCertArrayOutputWithContext(context.Background())
}

func (i SslKeyCertArray) ToSslKeyCertArrayOutputWithContext(ctx context.Context) SslKeyCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslKeyCertArrayOutput)
}

// SslKeyCertMapInput is an input type that accepts SslKeyCertMap and SslKeyCertMapOutput values.
// You can construct a concrete instance of `SslKeyCertMapInput` via:
//
//	SslKeyCertMap{ "key": SslKeyCertArgs{...} }
type SslKeyCertMapInput interface {
	pulumi.Input

	ToSslKeyCertMapOutput() SslKeyCertMapOutput
	ToSslKeyCertMapOutputWithContext(context.Context) SslKeyCertMapOutput
}

type SslKeyCertMap map[string]SslKeyCertInput

func (SslKeyCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslKeyCert)(nil)).Elem()
}

func (i SslKeyCertMap) ToSslKeyCertMapOutput() SslKeyCertMapOutput {
	return i.ToSslKeyCertMapOutputWithContext(context.Background())
}

func (i SslKeyCertMap) ToSslKeyCertMapOutputWithContext(ctx context.Context) SslKeyCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslKeyCertMapOutput)
}

type SslKeyCertOutput struct{ *pulumi.OutputState }

func (SslKeyCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslKeyCert)(nil)).Elem()
}

func (o SslKeyCertOutput) ToSslKeyCertOutput() SslKeyCertOutput {
	return o
}

func (o SslKeyCertOutput) ToSslKeyCertOutputWithContext(ctx context.Context) SslKeyCertOutput {
	return o
}

// The content of the cert.
func (o SslKeyCertOutput) CertContent() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.CertContent }).(pulumi.StringOutput)
}

// full path of the SSL certificate on the BIGIP.
func (o SslKeyCertOutput) CertFullPath() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.CertFullPath }).(pulumi.StringOutput)
}

// Name of the SSL certificate to be Imported on to BIGIP.
func (o SslKeyCertOutput) CertName() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.CertName }).(pulumi.StringOutput)
}

// The content of the key.
func (o SslKeyCertOutput) KeyContent() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.KeyContent }).(pulumi.StringOutput)
}

// full path of the SSL key on the BIGIP.
func (o SslKeyCertOutput) KeyFullPath() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.KeyFullPath }).(pulumi.StringOutput)
}

// Name of the SSL key to be Imported on to BIGIP.
func (o SslKeyCertOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// Partition on to SSL certificate and key to be imported.
func (o SslKeyCertOutput) Partition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringPtrOutput { return v.Partition }).(pulumi.StringPtrOutput)
}

// Passphrase on the SSL key.
func (o SslKeyCertOutput) Passphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslKeyCert) pulumi.StringPtrOutput { return v.Passphrase }).(pulumi.StringPtrOutput)
}

type SslKeyCertArrayOutput struct{ *pulumi.OutputState }

func (SslKeyCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslKeyCert)(nil)).Elem()
}

func (o SslKeyCertArrayOutput) ToSslKeyCertArrayOutput() SslKeyCertArrayOutput {
	return o
}

func (o SslKeyCertArrayOutput) ToSslKeyCertArrayOutputWithContext(ctx context.Context) SslKeyCertArrayOutput {
	return o
}

func (o SslKeyCertArrayOutput) Index(i pulumi.IntInput) SslKeyCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SslKeyCert {
		return vs[0].([]*SslKeyCert)[vs[1].(int)]
	}).(SslKeyCertOutput)
}

type SslKeyCertMapOutput struct{ *pulumi.OutputState }

func (SslKeyCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslKeyCert)(nil)).Elem()
}

func (o SslKeyCertMapOutput) ToSslKeyCertMapOutput() SslKeyCertMapOutput {
	return o
}

func (o SslKeyCertMapOutput) ToSslKeyCertMapOutputWithContext(ctx context.Context) SslKeyCertMapOutput {
	return o
}

func (o SslKeyCertMapOutput) MapIndex(k pulumi.StringInput) SslKeyCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SslKeyCert {
		return vs[0].(map[string]*SslKeyCert)[vs[1].(string)]
	}).(SslKeyCertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslKeyCertInput)(nil)).Elem(), &SslKeyCert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslKeyCertArrayInput)(nil)).Elem(), SslKeyCertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslKeyCertMapInput)(nil)).Elem(), SslKeyCertMap{})
	pulumi.RegisterOutputType(SslKeyCertOutput{})
	pulumi.RegisterOutputType(SslKeyCertArrayOutput{})
	pulumi.RegisterOutputType(SslKeyCertMapOutput{})
}
