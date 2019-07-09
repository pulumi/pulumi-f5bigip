// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `bigip_ltm_dns` Configures DNS server on F5 BIG-IP
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_dns.html.markdown.
type Dns struct {
	s *pulumi.ResourceState
}

// NewDns registers a new resource with the given unique name, arguments, and options.
func NewDns(ctx *pulumi.Context,
	name string, args *DnsArgs, opts ...pulumi.ResourceOpt) (*Dns, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["nameServers"] = nil
		inputs["numberOfDots"] = nil
		inputs["searches"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["nameServers"] = args.NameServers
		inputs["numberOfDots"] = args.NumberOfDots
		inputs["searches"] = args.Searches
	}
	s, err := ctx.RegisterResource("f5bigip:sys/dns:Dns", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dns{s: s}, nil
}

// GetDns gets an existing Dns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDns(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DnsState, opts ...pulumi.ResourceOpt) (*Dns, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["nameServers"] = state.NameServers
		inputs["numberOfDots"] = state.NumberOfDots
		inputs["searches"] = state.Searches
	}
	s, err := ctx.ReadResource("f5bigip:sys/dns:Dns", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Dns{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Dns) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Dns) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Name of the Dns Servers
func (r *Dns) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Name or IP address of the DNS server
func (r *Dns) NameServers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nameServers"])
}

// Configures the number of dots needed in a name before an initial absolute query will be made.
func (r *Dns) NumberOfDots() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["numberOfDots"])
}

// Specify what domains you want to search
func (r *Dns) Searches() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["searches"])
}

// Input properties used for looking up and filtering Dns resources.
type DnsState struct {
	// Name of the Dns Servers
	Description interface{}
	// Name or IP address of the DNS server
	NameServers interface{}
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots interface{}
	// Specify what domains you want to search
	Searches interface{}
}

// The set of arguments for constructing a Dns resource.
type DnsArgs struct {
	// Name of the Dns Servers
	Description interface{}
	// Name or IP address of the DNS server
	NameServers interface{}
	// Configures the number of dots needed in a name before an initial absolute query will be made.
	NumberOfDots interface{}
	// Specify what domains you want to search
	Searches interface{}
}
