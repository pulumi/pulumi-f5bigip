// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWafEntityParameter(ctx *pulumi.Context, args *GetWafEntityParameterArgs, opts ...pulumi.InvokeOption) (*GetWafEntityParameterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWafEntityParameterResult
	err := ctx.Invoke("f5bigip:ssl/getWafEntityParameter:getWafEntityParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafEntityParameter.
type GetWafEntityParameterArgs struct {
	AllowEmptyType                 *bool                     `pulumi:"allowEmptyType"`
	AllowRepeatedParameterName     *bool                     `pulumi:"allowRepeatedParameterName"`
	AttackSignaturesCheck          *bool                     `pulumi:"attackSignaturesCheck"`
	CheckMaxValueLength            *bool                     `pulumi:"checkMaxValueLength"`
	CheckMinValueLength            *bool                     `pulumi:"checkMinValueLength"`
	DataType                       *string                   `pulumi:"dataType"`
	Description                    *string                   `pulumi:"description"`
	EnableRegularExpression        *bool                     `pulumi:"enableRegularExpression"`
	IsBase64                       *bool                     `pulumi:"isBase64"`
	IsCookie                       *bool                     `pulumi:"isCookie"`
	IsHeader                       *bool                     `pulumi:"isHeader"`
	Json                           *string                   `pulumi:"json"`
	Level                          *string                   `pulumi:"level"`
	Mandatory                      *bool                     `pulumi:"mandatory"`
	MaxValueLength                 *int                      `pulumi:"maxValueLength"`
	MetacharsOnParameterValueCheck *bool                     `pulumi:"metacharsOnParameterValueCheck"`
	MinValueLength                 *int                      `pulumi:"minValueLength"`
	Name                           string                    `pulumi:"name"`
	ParameterLocation              *string                   `pulumi:"parameterLocation"`
	PerformStaging                 *bool                     `pulumi:"performStaging"`
	SensitiveParameter             *bool                     `pulumi:"sensitiveParameter"`
	SignatureOverridesDisables     []int                     `pulumi:"signatureOverridesDisables"`
	Type                           *string                   `pulumi:"type"`
	Url                            *GetWafEntityParameterUrl `pulumi:"url"`
	ValueType                      *string                   `pulumi:"valueType"`
}

// A collection of values returned by getWafEntityParameter.
type GetWafEntityParameterResult struct {
	AllowEmptyType             *bool   `pulumi:"allowEmptyType"`
	AllowRepeatedParameterName *bool   `pulumi:"allowRepeatedParameterName"`
	AttackSignaturesCheck      *bool   `pulumi:"attackSignaturesCheck"`
	CheckMaxValueLength        *bool   `pulumi:"checkMaxValueLength"`
	CheckMinValueLength        *bool   `pulumi:"checkMinValueLength"`
	DataType                   *string `pulumi:"dataType"`
	Description                *string `pulumi:"description"`
	EnableRegularExpression    *bool   `pulumi:"enableRegularExpression"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                    `pulumi:"id"`
	IsBase64                       *bool                     `pulumi:"isBase64"`
	IsCookie                       *bool                     `pulumi:"isCookie"`
	IsHeader                       *bool                     `pulumi:"isHeader"`
	Json                           string                    `pulumi:"json"`
	Level                          *string                   `pulumi:"level"`
	Mandatory                      *bool                     `pulumi:"mandatory"`
	MaxValueLength                 *int                      `pulumi:"maxValueLength"`
	MetacharsOnParameterValueCheck *bool                     `pulumi:"metacharsOnParameterValueCheck"`
	MinValueLength                 *int                      `pulumi:"minValueLength"`
	Name                           string                    `pulumi:"name"`
	ParameterLocation              *string                   `pulumi:"parameterLocation"`
	PerformStaging                 *bool                     `pulumi:"performStaging"`
	SensitiveParameter             *bool                     `pulumi:"sensitiveParameter"`
	SignatureOverridesDisables     []int                     `pulumi:"signatureOverridesDisables"`
	Type                           *string                   `pulumi:"type"`
	Url                            *GetWafEntityParameterUrl `pulumi:"url"`
	ValueType                      *string                   `pulumi:"valueType"`
}

func GetWafEntityParameterOutput(ctx *pulumi.Context, args GetWafEntityParameterOutputArgs, opts ...pulumi.InvokeOption) GetWafEntityParameterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetWafEntityParameterResultOutput, error) {
			args := v.(GetWafEntityParameterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("f5bigip:ssl/getWafEntityParameter:getWafEntityParameter", args, GetWafEntityParameterResultOutput{}, options).(GetWafEntityParameterResultOutput), nil
		}).(GetWafEntityParameterResultOutput)
}

// A collection of arguments for invoking getWafEntityParameter.
type GetWafEntityParameterOutputArgs struct {
	AllowEmptyType                 pulumi.BoolPtrInput              `pulumi:"allowEmptyType"`
	AllowRepeatedParameterName     pulumi.BoolPtrInput              `pulumi:"allowRepeatedParameterName"`
	AttackSignaturesCheck          pulumi.BoolPtrInput              `pulumi:"attackSignaturesCheck"`
	CheckMaxValueLength            pulumi.BoolPtrInput              `pulumi:"checkMaxValueLength"`
	CheckMinValueLength            pulumi.BoolPtrInput              `pulumi:"checkMinValueLength"`
	DataType                       pulumi.StringPtrInput            `pulumi:"dataType"`
	Description                    pulumi.StringPtrInput            `pulumi:"description"`
	EnableRegularExpression        pulumi.BoolPtrInput              `pulumi:"enableRegularExpression"`
	IsBase64                       pulumi.BoolPtrInput              `pulumi:"isBase64"`
	IsCookie                       pulumi.BoolPtrInput              `pulumi:"isCookie"`
	IsHeader                       pulumi.BoolPtrInput              `pulumi:"isHeader"`
	Json                           pulumi.StringPtrInput            `pulumi:"json"`
	Level                          pulumi.StringPtrInput            `pulumi:"level"`
	Mandatory                      pulumi.BoolPtrInput              `pulumi:"mandatory"`
	MaxValueLength                 pulumi.IntPtrInput               `pulumi:"maxValueLength"`
	MetacharsOnParameterValueCheck pulumi.BoolPtrInput              `pulumi:"metacharsOnParameterValueCheck"`
	MinValueLength                 pulumi.IntPtrInput               `pulumi:"minValueLength"`
	Name                           pulumi.StringInput               `pulumi:"name"`
	ParameterLocation              pulumi.StringPtrInput            `pulumi:"parameterLocation"`
	PerformStaging                 pulumi.BoolPtrInput              `pulumi:"performStaging"`
	SensitiveParameter             pulumi.BoolPtrInput              `pulumi:"sensitiveParameter"`
	SignatureOverridesDisables     pulumi.IntArrayInput             `pulumi:"signatureOverridesDisables"`
	Type                           pulumi.StringPtrInput            `pulumi:"type"`
	Url                            GetWafEntityParameterUrlPtrInput `pulumi:"url"`
	ValueType                      pulumi.StringPtrInput            `pulumi:"valueType"`
}

func (GetWafEntityParameterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafEntityParameterArgs)(nil)).Elem()
}

// A collection of values returned by getWafEntityParameter.
type GetWafEntityParameterResultOutput struct{ *pulumi.OutputState }

func (GetWafEntityParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafEntityParameterResult)(nil)).Elem()
}

func (o GetWafEntityParameterResultOutput) ToGetWafEntityParameterResultOutput() GetWafEntityParameterResultOutput {
	return o
}

func (o GetWafEntityParameterResultOutput) ToGetWafEntityParameterResultOutputWithContext(ctx context.Context) GetWafEntityParameterResultOutput {
	return o
}

func (o GetWafEntityParameterResultOutput) AllowEmptyType() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.AllowEmptyType }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) AllowRepeatedParameterName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.AllowRepeatedParameterName }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) AttackSignaturesCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.AttackSignaturesCheck }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) CheckMaxValueLength() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.CheckMaxValueLength }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) CheckMinValueLength() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.CheckMinValueLength }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityParameterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityParameterResultOutput) EnableRegularExpression() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.EnableRegularExpression }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafEntityParameterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWafEntityParameterResultOutput) IsBase64() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.IsBase64 }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) IsCookie() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.IsCookie }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) IsHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.IsHeader }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetWafEntityParameterResultOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityParameterResultOutput) Mandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.Mandatory }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) MaxValueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *int { return v.MaxValueLength }).(pulumi.IntPtrOutput)
}

func (o GetWafEntityParameterResultOutput) MetacharsOnParameterValueCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.MetacharsOnParameterValueCheck }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) MinValueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *int { return v.MinValueLength }).(pulumi.IntPtrOutput)
}

func (o GetWafEntityParameterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetWafEntityParameterResultOutput) ParameterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.ParameterLocation }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityParameterResultOutput) PerformStaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.PerformStaging }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) SensitiveParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *bool { return v.SensitiveParameter }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityParameterResultOutput) SignatureOverridesDisables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) []int { return v.SignatureOverridesDisables }).(pulumi.IntArrayOutput)
}

func (o GetWafEntityParameterResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityParameterResultOutput) Url() GetWafEntityParameterUrlPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *GetWafEntityParameterUrl { return v.Url }).(GetWafEntityParameterUrlPtrOutput)
}

func (o GetWafEntityParameterResultOutput) ValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityParameterResult) *string { return v.ValueType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafEntityParameterResultOutput{})
}
