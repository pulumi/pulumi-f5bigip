// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package f5bigip

import (
	"strings"
	"unicode"

	"github.com/F5Networks/terraform-provider-bigip/bigip"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the F5 BigIP token components used below.
const (
	f5BigIPPkg     = "f5bigip"
	f5BigIPCMMod   = "CM"    // Centralized Management (CM)
	f5BigIPLTMMod  = "Ltm"   // Local Traffic Manager (LTM)
	f5BigIPNetMod  = "Net"   // Network
	f5BigIPSysMod  = "Sys"   // System
	f5BigIPSslMod  = "Ssl"   // Ssl
	f5BigIPMainMod = "Index" // Index

)

var namespaceMap = map[string]string{
	f5BigIPPkg: "F5BigIP",
}

// makeMember manufactures a type token for the F5 BigIP package and the given module and type.  It automatically
// uses the F5 BigIP package and names the file by simply lower casing the resource's first character.
func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(f5BigIPPkg + ":" + token + ":" + mem)
}

// makeType manufactures a type token for the F5 BigIP package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.
func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

// Provider returns additional overlaid schema and metadata associated with the F5 BigIP package.
func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		P:           shimv1.NewProvider(bigip.Provider().(*schema.Provider)),
		Name:        "bigip",
		Description: "A Pulumi package for creating and managing F5 BigIP resources.",
		Keywords:    []string{"pulumi", "f5", "bigip"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-f5bigip",
		GitHubOrg:   "F5Networks",
		Resources: map[string]*tfbridge.ResourceInfo{
			"bigip_cm_device":                       {Tok: makeResource(f5BigIPCMMod, "Device")},
			"bigip_cm_devicegroup":                  {Tok: makeResource(f5BigIPCMMod, "DeviceGroup")},
			"bigip_net_route":                       {Tok: makeResource(f5BigIPNetMod, "Route")},
			"bigip_net_selfip":                      {Tok: makeResource(f5BigIPNetMod, "SelfIp")},
			"bigip_net_vlan":                        {Tok: makeResource(f5BigIPNetMod, "Vlan")},
			"bigip_ltm_irule":                       {Tok: makeResource(f5BigIPLTMMod, "IRule")},
			"bigip_ltm_datagroup":                   {Tok: makeResource(f5BigIPLTMMod, "DataGroup")},
			"bigip_ltm_monitor":                     {Tok: makeResource(f5BigIPLTMMod, "Monitor")},
			"bigip_ltm_node":                        {Tok: makeResource(f5BigIPLTMMod, "Node")},
			"bigip_ltm_pool":                        {Tok: makeResource(f5BigIPLTMMod, "Pool")},
			"bigip_ltm_pool_attachment":             {Tok: makeResource(f5BigIPLTMMod, "PoolAttachment")},
			"bigip_ltm_policy":                      {Tok: makeResource(f5BigIPLTMMod, "Policy")},
			"bigip_ltm_profile_fasthttp":            {Tok: makeResource(f5BigIPLTMMod, "ProfileFastHttp")},
			"bigip_ltm_profile_fastl4":              {Tok: makeResource(f5BigIPLTMMod, "ProfileFastL4")},
			"bigip_ltm_profile_http":                {Tok: makeResource(f5BigIPLTMMod, "ProfileHttp")},
			"bigip_ltm_profile_http2":               {Tok: makeResource(f5BigIPLTMMod, "ProfileHttp2")},
			"bigip_ltm_profile_httpcompress":        {Tok: makeResource(f5BigIPLTMMod, "ProfileHttpCompress")},
			"bigip_ltm_profile_oneconnect":          {Tok: makeResource(f5BigIPLTMMod, "ProfileOneConnect")},
			"bigip_ltm_profile_tcp":                 {Tok: makeResource(f5BigIPLTMMod, "ProfileTcp")},
			"bigip_ltm_persistence_profile_srcaddr": {Tok: makeResource(f5BigIPLTMMod, "PersistenceProfileSrcAddr")},
			"bigip_ltm_persistence_profile_dstaddr": {Tok: makeResource(f5BigIPLTMMod, "PersistenceProfileDstAddr")},
			"bigip_ltm_persistence_profile_ssl":     {Tok: makeResource(f5BigIPLTMMod, "PersistenceProfileSsl")},
			"bigip_ltm_persistence_profile_cookie":  {Tok: makeResource(f5BigIPLTMMod, "PersistenceProfileCookie")},
			"bigip_ltm_snat":                        {Tok: makeResource(f5BigIPLTMMod, "Snat")},
			"bigip_ltm_snatpool":                    {Tok: makeResource(f5BigIPLTMMod, "SnatPool")},
			"bigip_ltm_virtual_address":             {Tok: makeResource(f5BigIPLTMMod, "VirtualAddress")},
			"bigip_ltm_virtual_server":              {Tok: makeResource(f5BigIPLTMMod, "VirtualServer")},
			"bigip_ltm_profile_client_ssl":          {Tok: makeResource(f5BigIPLTMMod, "ProfileClientSsl")},
			"bigip_ltm_profile_server_ssl":          {Tok: makeResource(f5BigIPLTMMod, "ProfileServerSsl")},
			"bigip_sys_dns":                         {Tok: makeResource(f5BigIPSysMod, "Dns")},
			"bigip_sys_iapp":                        {Tok: makeResource(f5BigIPSysMod, "IApp")},
			"bigip_sys_ntp":                         {Tok: makeResource(f5BigIPSysMod, "Ntp")},
			"bigip_sys_provision":                   {Tok: makeResource(f5BigIPSysMod, "Provision")},
			"bigip_sys_snmp":                        {Tok: makeResource(f5BigIPSysMod, "Snmp")},
			"bigip_sys_snmp_traps":                  {Tok: makeResource(f5BigIPSysMod, "SnmpTraps")},
			"bigip_sys_bigiplicense":                {Tok: makeResource(f5BigIPSysMod, "BigIpLicense")},
			"bigip_ssl_certificate":                 {Tok: makeResource(f5BigIPSslMod, "Certificate")},
			"bigip_ssl_key":                         {Tok: makeResource(f5BigIPSslMod, "Key")},
			"bigip_as3":                             {Tok: makeResource(f5BigIPMainMod, "As3")},
			"bigip_do":                              {Tok: makeResource(f5BigIPMainMod, "Do")},
			"bigip_bigiq_as3":                       {Tok: makeResource(f5BigIPMainMod, "BigIqAs3")},
			"bigip_command":                         {Tok: makeResource(f5BigIPMainMod, "Command")},
			"bigip_common_license_manage_bigiq":     {Tok: makeResource(f5BigIPMainMod, "CommonLicenseManageBigIq")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
			UsesIOClasses: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}
}
