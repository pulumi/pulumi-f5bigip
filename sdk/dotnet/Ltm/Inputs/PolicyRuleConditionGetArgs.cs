// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class PolicyRuleConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<bool>? Address { get; set; }

        [Input("all")]
        public Input<bool>? All { get; set; }

        [Input("appService")]
        public Input<string>? AppService { get; set; }

        [Input("browserType")]
        public Input<bool>? BrowserType { get; set; }

        [Input("browserVersion")]
        public Input<bool>? BrowserVersion { get; set; }

        [Input("caseInsensitive")]
        public Input<bool>? CaseInsensitive { get; set; }

        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        [Input("cipher")]
        public Input<bool>? Cipher { get; set; }

        [Input("cipherBits")]
        public Input<bool>? CipherBits { get; set; }

        [Input("clientSsl")]
        public Input<bool>? ClientSsl { get; set; }

        [Input("code")]
        public Input<bool>? Code { get; set; }

        [Input("commonName")]
        public Input<bool>? CommonName { get; set; }

        [Input("contains")]
        public Input<bool>? Contains { get; set; }

        [Input("continent")]
        public Input<bool>? Continent { get; set; }

        [Input("countryCode")]
        public Input<bool>? CountryCode { get; set; }

        [Input("countryName")]
        public Input<bool>? CountryName { get; set; }

        [Input("cpuUsage")]
        public Input<bool>? CpuUsage { get; set; }

        [Input("deviceMake")]
        public Input<bool>? DeviceMake { get; set; }

        [Input("deviceModel")]
        public Input<bool>? DeviceModel { get; set; }

        [Input("domain")]
        public Input<bool>? Domain { get; set; }

        [Input("endsWith")]
        public Input<bool>? EndsWith { get; set; }

        [Input("equals")]
        public Input<bool>? Equals { get; set; }

        [Input("expiry")]
        public Input<bool>? Expiry { get; set; }

        [Input("extension")]
        public Input<bool>? Extension { get; set; }

        [Input("external")]
        public Input<bool>? External { get; set; }

        [Input("geoip")]
        public Input<bool>? Geoip { get; set; }

        [Input("greater")]
        public Input<bool>? Greater { get; set; }

        [Input("greaterOrEqual")]
        public Input<bool>? GreaterOrEqual { get; set; }

        [Input("host")]
        public Input<bool>? Host { get; set; }

        [Input("httpBasicAuth")]
        public Input<bool>? HttpBasicAuth { get; set; }

        [Input("httpCookie")]
        public Input<bool>? HttpCookie { get; set; }

        [Input("httpHeader")]
        public Input<bool>? HttpHeader { get; set; }

        [Input("httpHost")]
        public Input<bool>? HttpHost { get; set; }

        [Input("httpMethod")]
        public Input<bool>? HttpMethod { get; set; }

        [Input("httpReferer")]
        public Input<bool>? HttpReferer { get; set; }

        [Input("httpSetCookie")]
        public Input<bool>? HttpSetCookie { get; set; }

        [Input("httpStatus")]
        public Input<bool>? HttpStatus { get; set; }

        [Input("httpUri")]
        public Input<bool>? HttpUri { get; set; }

        [Input("httpUserAgent")]
        public Input<bool>? HttpUserAgent { get; set; }

        [Input("httpVersion")]
        public Input<bool>? HttpVersion { get; set; }

        [Input("index")]
        public Input<int>? Index { get; set; }

        [Input("internal")]
        public Input<bool>? Internal { get; set; }

        [Input("isp")]
        public Input<bool>? Isp { get; set; }

        [Input("last15secs")]
        public Input<bool>? Last15secs { get; set; }

        [Input("last1min")]
        public Input<bool>? Last1min { get; set; }

        [Input("last5mins")]
        public Input<bool>? Last5mins { get; set; }

        [Input("less")]
        public Input<bool>? Less { get; set; }

        [Input("lessOrEqual")]
        public Input<bool>? LessOrEqual { get; set; }

        [Input("local")]
        public Input<bool>? Local { get; set; }

        [Input("major")]
        public Input<bool>? Major { get; set; }

        [Input("matches")]
        public Input<bool>? Matches { get; set; }

        [Input("minor")]
        public Input<bool>? Minor { get; set; }

        [Input("missing")]
        public Input<bool>? Missing { get; set; }

        [Input("mss")]
        public Input<bool>? Mss { get; set; }

        [Input("not")]
        public Input<bool>? Not { get; set; }

        [Input("org")]
        public Input<bool>? Org { get; set; }

        [Input("password")]
        public Input<bool>? Password { get; set; }

        [Input("path")]
        public Input<bool>? Path { get; set; }

        [Input("pathSegment")]
        public Input<bool>? PathSegment { get; set; }

        [Input("port")]
        public Input<bool>? Port { get; set; }

        [Input("present")]
        public Input<bool>? Present { get; set; }

        [Input("protocol")]
        public Input<bool>? Protocol { get; set; }

        [Input("queryParameter")]
        public Input<bool>? QueryParameter { get; set; }

        [Input("queryString")]
        public Input<bool>? QueryString { get; set; }

        [Input("regionCode")]
        public Input<bool>? RegionCode { get; set; }

        [Input("regionName")]
        public Input<bool>? RegionName { get; set; }

        [Input("remote")]
        public Input<bool>? Remote { get; set; }

        [Input("request")]
        public Input<bool>? Request { get; set; }

        [Input("response")]
        public Input<bool>? Response { get; set; }

        [Input("routeDomain")]
        public Input<bool>? RouteDomain { get; set; }

        [Input("rtt")]
        public Input<bool>? Rtt { get; set; }

        [Input("scheme")]
        public Input<bool>? Scheme { get; set; }

        [Input("serverName")]
        public Input<bool>? ServerName { get; set; }

        [Input("sslCert")]
        public Input<bool>? SslCert { get; set; }

        [Input("sslClientHello")]
        public Input<bool>? SslClientHello { get; set; }

        [Input("sslExtension")]
        public Input<bool>? SslExtension { get; set; }

        [Input("sslServerHandshake")]
        public Input<bool>? SslServerHandshake { get; set; }

        [Input("sslServerHello")]
        public Input<bool>? SslServerHello { get; set; }

        [Input("startsWith")]
        public Input<bool>? StartsWith { get; set; }

        [Input("tcp")]
        public Input<bool>? Tcp { get; set; }

        [Input("text")]
        public Input<bool>? Text { get; set; }

        /// <summary>
        /// If Rule is used then you need to provide the tm_name it can be any value
        /// </summary>
        [Input("tmName")]
        public Input<string>? TmName { get; set; }

        [Input("unnamedQueryParameter")]
        public Input<bool>? UnnamedQueryParameter { get; set; }

        [Input("userAgentToken")]
        public Input<bool>? UserAgentToken { get; set; }

        [Input("username")]
        public Input<bool>? Username { get; set; }

        [Input("value")]
        public Input<bool>? Value { get; set; }

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("version")]
        public Input<bool>? Version { get; set; }

        [Input("vlan")]
        public Input<bool>? Vlan { get; set; }

        [Input("vlanId")]
        public Input<bool>? VlanId { get; set; }

        public PolicyRuleConditionGetArgs()
        {
        }
    }
}
