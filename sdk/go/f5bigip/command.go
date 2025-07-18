// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Command` Run TMSH commands on F5 devices
//
// This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device
type Command struct {
	pulumi.CustomResourceState

	// The resulting output from the `commands` executed.
	CommandResults pulumi.StringArrayOutput `pulumi:"commandResults"`
	// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
	Commands pulumi.StringArrayOutput `pulumi:"commands"`
	When     pulumi.StringPtrOutput   `pulumi:"when"`
}

// NewCommand registers a new resource with the given unique name, arguments, and options.
func NewCommand(ctx *pulumi.Context,
	name string, args *CommandArgs, opts ...pulumi.ResourceOption) (*Command, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Commands == nil {
		return nil, errors.New("invalid value for required argument 'Commands'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Command
	err := ctx.RegisterResource("f5bigip:index/command:Command", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommand gets an existing Command resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommandState, opts ...pulumi.ResourceOption) (*Command, error) {
	var resource Command
	err := ctx.ReadResource("f5bigip:index/command:Command", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Command resources.
type commandState struct {
	// The resulting output from the `commands` executed.
	CommandResults []string `pulumi:"commandResults"`
	// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
	Commands []string `pulumi:"commands"`
	When     *string  `pulumi:"when"`
}

type CommandState struct {
	// The resulting output from the `commands` executed.
	CommandResults pulumi.StringArrayInput
	// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
	Commands pulumi.StringArrayInput
	When     pulumi.StringPtrInput
}

func (CommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*commandState)(nil)).Elem()
}

type commandArgs struct {
	// The resulting output from the `commands` executed.
	CommandResults []string `pulumi:"commandResults"`
	// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
	Commands []string `pulumi:"commands"`
	When     *string  `pulumi:"when"`
}

// The set of arguments for constructing a Command resource.
type CommandArgs struct {
	// The resulting output from the `commands` executed.
	CommandResults pulumi.StringArrayInput
	// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
	Commands pulumi.StringArrayInput
	When     pulumi.StringPtrInput
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commandArgs)(nil)).Elem()
}

type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(ctx context.Context) CommandOutput
}

func (*Command) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (i *Command) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i *Command) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}

// CommandArrayInput is an input type that accepts CommandArray and CommandArrayOutput values.
// You can construct a concrete instance of `CommandArrayInput` via:
//
//	CommandArray{ CommandArgs{...} }
type CommandArrayInput interface {
	pulumi.Input

	ToCommandArrayOutput() CommandArrayOutput
	ToCommandArrayOutputWithContext(context.Context) CommandArrayOutput
}

type CommandArray []CommandInput

func (CommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (i CommandArray) ToCommandArrayOutput() CommandArrayOutput {
	return i.ToCommandArrayOutputWithContext(context.Background())
}

func (i CommandArray) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandArrayOutput)
}

// CommandMapInput is an input type that accepts CommandMap and CommandMapOutput values.
// You can construct a concrete instance of `CommandMapInput` via:
//
//	CommandMap{ "key": CommandArgs{...} }
type CommandMapInput interface {
	pulumi.Input

	ToCommandMapOutput() CommandMapOutput
	ToCommandMapOutputWithContext(context.Context) CommandMapOutput
}

type CommandMap map[string]CommandInput

func (CommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (i CommandMap) ToCommandMapOutput() CommandMapOutput {
	return i.ToCommandMapOutputWithContext(context.Background())
}

func (i CommandMap) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandMapOutput)
}

type CommandOutput struct{ *pulumi.OutputState }

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

// The resulting output from the `commands` executed.
func (o CommandOutput) CommandResults() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Command) pulumi.StringArrayOutput { return v.CommandResults }).(pulumi.StringArrayOutput)
}

// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
func (o CommandOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Command) pulumi.StringArrayOutput { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o CommandOutput) When() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.When }).(pulumi.StringPtrOutput)
}

type CommandArrayOutput struct{ *pulumi.OutputState }

func (CommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (o CommandArrayOutput) ToCommandArrayOutput() CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) Index(i pulumi.IntInput) CommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Command {
		return vs[0].([]*Command)[vs[1].(int)]
	}).(CommandOutput)
}

type CommandMapOutput struct{ *pulumi.OutputState }

func (CommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (o CommandMapOutput) ToCommandMapOutput() CommandMapOutput {
	return o
}

func (o CommandMapOutput) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return o
}

func (o CommandMapOutput) MapIndex(k pulumi.StringInput) CommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Command {
		return vs[0].(map[string]*Command)[vs[1].(string)]
	}).(CommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommandInput)(nil)).Elem(), &Command{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandArrayInput)(nil)).Elem(), CommandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandMapInput)(nil)).Elem(), CommandMap{})
	pulumi.RegisterOutputType(CommandOutput{})
	pulumi.RegisterOutputType(CommandArrayOutput{})
	pulumi.RegisterOutputType(CommandMapOutput{})
}
