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
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/F5Networks/terraform-provider-bigip/bigip"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-f5bigip/provider/v2/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the F5 BigIP token components used below.
const (
	f5BigIPPkg = "f5bigip"
	cmMod      = "CM"    // Centralized Management (CM)
	ltmMod     = "Ltm"   // Local Traffic Manager (LTM)
	netMod     = "Net"   // Network
	sysMod     = "Sys"   // System
	sslMod     = "Ssl"   // Ssl
	mainMod    = "Index" // Index

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

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
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
			"bigip_cm_device":                       {Tok: makeResource(cmMod, "Device")},
			"bigip_cm_devicegroup":                  {Tok: makeResource(cmMod, "DeviceGroup")},
			"bigip_net_route":                       {Tok: makeResource(netMod, "Route")},
			"bigip_net_selfip":                      {Tok: makeResource(netMod, "SelfIp")},
			"bigip_net_vlan":                        {Tok: makeResource(netMod, "Vlan")},
			"bigip_ltm_irule":                       {Tok: makeResource(ltmMod, "IRule")},
			"bigip_ltm_datagroup":                   {Tok: makeResource(ltmMod, "DataGroup")},
			"bigip_ltm_monitor":                     {Tok: makeResource(ltmMod, "Monitor")},
			"bigip_ltm_node":                        {Tok: makeResource(ltmMod, "Node")},
			"bigip_ltm_pool":                        {Tok: makeResource(ltmMod, "Pool")},
			"bigip_ltm_pool_attachment":             {Tok: makeResource(ltmMod, "PoolAttachment")},
			"bigip_ltm_policy":                      {Tok: makeResource(ltmMod, "Policy")},
			"bigip_ltm_profile_fasthttp":            {Tok: makeResource(ltmMod, "ProfileFastHttp")},
			"bigip_ltm_profile_fastl4":              {Tok: makeResource(ltmMod, "ProfileFastL4")},
			"bigip_ltm_profile_http":                {Tok: makeResource(ltmMod, "ProfileHttp")},
			"bigip_ltm_profile_http2":               {Tok: makeResource(ltmMod, "ProfileHttp2")},
			"bigip_ltm_profile_httpcompress":        {Tok: makeResource(ltmMod, "ProfileHttpCompress")},
			"bigip_ltm_profile_oneconnect":          {Tok: makeResource(ltmMod, "ProfileOneConnect")},
			"bigip_ltm_profile_tcp":                 {Tok: makeResource(ltmMod, "ProfileTcp")},
			"bigip_ltm_persistence_profile_srcaddr": {Tok: makeResource(ltmMod, "PersistenceProfileSrcAddr")},
			"bigip_ltm_persistence_profile_dstaddr": {Tok: makeResource(ltmMod, "PersistenceProfileDstAddr")},
			"bigip_ltm_persistence_profile_ssl":     {Tok: makeResource(ltmMod, "PersistenceProfileSsl")},
			"bigip_ltm_persistence_profile_cookie":  {Tok: makeResource(ltmMod, "PersistenceProfileCookie")},
			"bigip_ltm_snat":                        {Tok: makeResource(ltmMod, "Snat")},
			"bigip_ltm_snatpool":                    {Tok: makeResource(ltmMod, "SnatPool")},
			"bigip_ltm_virtual_address":             {Tok: makeResource(ltmMod, "VirtualAddress")},
			"bigip_ltm_virtual_server":              {Tok: makeResource(ltmMod, "VirtualServer")},
			"bigip_ltm_profile_client_ssl":          {Tok: makeResource(ltmMod, "ProfileClientSsl")},
			"bigip_ltm_profile_server_ssl":          {Tok: makeResource(ltmMod, "ProfileServerSsl")},
			"bigip_sys_dns":                         {Tok: makeResource(sysMod, "Dns")},
			"bigip_sys_iapp":                        {Tok: makeResource(sysMod, "IApp")},
			"bigip_sys_ntp":                         {Tok: makeResource(sysMod, "Ntp")},
			"bigip_sys_provision":                   {Tok: makeResource(sysMod, "Provision")},
			"bigip_sys_snmp":                        {Tok: makeResource(sysMod, "Snmp")},
			"bigip_sys_snmp_traps":                  {Tok: makeResource(sysMod, "SnmpTraps")},
			"bigip_sys_bigiplicense":                {Tok: makeResource(sysMod, "BigIpLicense")},
			"bigip_ssl_certificate":                 {Tok: makeResource(sslMod, "Certificate")},
			"bigip_ssl_key":                         {Tok: makeResource(sslMod, "Key")},
			"bigip_as3":                             {Tok: makeResource(mainMod, "As3")},
			"bigip_do":                              {Tok: makeResource(mainMod, "Do")},
			"bigip_bigiq_as3":                       {Tok: makeResource(mainMod, "BigIqAs3")},
			"bigip_command":                         {Tok: makeResource(mainMod, "Command")},
			"bigip_common_license_manage_bigiq":     {Tok: makeResource(mainMod, "CommonLicenseManageBigIq")},
			"bigip_event_service_discovery":         {Tok: makeResource(mainMod, "EventServiceDiscovery")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"bigip_ltm_datagroup":   {Tok: makeDataSource(ltmMod, "getDataGroup")},
			"bigip_ltm_irule":       {Tok: makeDataSource(ltmMod, "getIrule")},
			"bigip_ltm_monitor":     {Tok: makeDataSource(ltmMod, "getMonitor")},
			"bigip_ltm_pool":        {Tok: makeDataSource(ltmMod, "getPool")},
			"bigip_ltm_node":        {Tok: makeDataSource(ltmMod, "getNode")},
			"bigip_ssl_certificate": {Tok: makeDataSource(sslMod, "getCertificate")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", f5BigIPPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				f5BigIPPkg,
			),
			GenerateResourceContainerTypes: true,
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
