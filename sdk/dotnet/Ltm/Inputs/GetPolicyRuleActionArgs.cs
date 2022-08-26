// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class GetPolicyRuleActionInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("appService", required: true)]
        public Input<string> AppService { get; set; } = null!;

        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        [Input("asm", required: true)]
        public Input<bool> Asm { get; set; } = null!;

        [Input("avr", required: true)]
        public Input<bool> Avr { get; set; } = null!;

        [Input("cache", required: true)]
        public Input<bool> Cache { get; set; } = null!;

        [Input("carp", required: true)]
        public Input<bool> Carp { get; set; } = null!;

        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        [Input("classify", required: true)]
        public Input<bool> Classify { get; set; } = null!;

        [Input("clonePool", required: true)]
        public Input<string> ClonePool { get; set; } = null!;

        [Input("code", required: true)]
        public Input<int> Code { get; set; } = null!;

        [Input("compress", required: true)]
        public Input<bool> Compress { get; set; } = null!;

        [Input("connection", required: true)]
        public Input<bool> Connection { get; set; } = null!;

        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("cookieHash", required: true)]
        public Input<bool> CookieHash { get; set; } = null!;

        [Input("cookieInsert", required: true)]
        public Input<bool> CookieInsert { get; set; } = null!;

        [Input("cookiePassive", required: true)]
        public Input<bool> CookiePassive { get; set; } = null!;

        [Input("cookieRewrite", required: true)]
        public Input<bool> CookieRewrite { get; set; } = null!;

        [Input("decompress", required: true)]
        public Input<bool> Decompress { get; set; } = null!;

        [Input("defer", required: true)]
        public Input<bool> Defer { get; set; } = null!;

        [Input("destinationAddress", required: true)]
        public Input<bool> DestinationAddress { get; set; } = null!;

        [Input("disable", required: true)]
        public Input<bool> Disable { get; set; } = null!;

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        [Input("expiry", required: true)]
        public Input<string> Expiry { get; set; } = null!;

        [Input("expirySecs", required: true)]
        public Input<int> ExpirySecs { get; set; } = null!;

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("extension", required: true)]
        public Input<string> Extension { get; set; } = null!;

        [Input("facility", required: true)]
        public Input<string> Facility { get; set; } = null!;

        [Input("forward")]
        public Input<bool>? Forward { get; set; }

        [Input("fromProfile", required: true)]
        public Input<string> FromProfile { get; set; } = null!;

        [Input("hash", required: true)]
        public Input<bool> Hash { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("http", required: true)]
        public Input<bool> Http { get; set; } = null!;

        [Input("httpBasicAuth", required: true)]
        public Input<bool> HttpBasicAuth { get; set; } = null!;

        [Input("httpCookie", required: true)]
        public Input<bool> HttpCookie { get; set; } = null!;

        [Input("httpHeader", required: true)]
        public Input<bool> HttpHeader { get; set; } = null!;

        [Input("httpHost")]
        public Input<bool>? HttpHost { get; set; }

        [Input("httpReferer", required: true)]
        public Input<bool> HttpReferer { get; set; } = null!;

        [Input("httpReply", required: true)]
        public Input<bool> HttpReply { get; set; } = null!;

        [Input("httpSetCookie", required: true)]
        public Input<bool> HttpSetCookie { get; set; } = null!;

        [Input("httpUri", required: true)]
        public Input<bool> HttpUri { get; set; } = null!;

        [Input("ifile", required: true)]
        public Input<string> Ifile { get; set; } = null!;

        [Input("insert", required: true)]
        public Input<bool> Insert { get; set; } = null!;

        [Input("internalVirtual", required: true)]
        public Input<string> InternalVirtual { get; set; } = null!;

        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("l7dos", required: true)]
        public Input<bool> L7dos { get; set; } = null!;

        [Input("length", required: true)]
        public Input<int> Length { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("log", required: true)]
        public Input<bool> Log { get; set; } = null!;

        [Input("ltmPolicy", required: true)]
        public Input<bool> LtmPolicy { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        [Input("netmask", required: true)]
        public Input<string> Netmask { get; set; } = null!;

        [Input("nexthop", required: true)]
        public Input<string> Nexthop { get; set; } = null!;

        [Input("node", required: true)]
        public Input<string> Node { get; set; } = null!;

        [Input("offset", required: true)]
        public Input<int> Offset { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("pem", required: true)]
        public Input<bool> Pem { get; set; } = null!;

        [Input("persist", required: true)]
        public Input<bool> Persist { get; set; } = null!;

        [Input("pin", required: true)]
        public Input<bool> Pin { get; set; } = null!;

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        [Input("pool", required: true)]
        public Input<string> Pool { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("priority", required: true)]
        public Input<string> Priority { get; set; } = null!;

        [Input("profile", required: true)]
        public Input<string> Profile { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("queryString", required: true)]
        public Input<string> QueryString { get; set; } = null!;

        [Input("rateclass", required: true)]
        public Input<string> Rateclass { get; set; } = null!;

        [Input("redirect", required: true)]
        public Input<bool> Redirect { get; set; } = null!;

        [Input("remove", required: true)]
        public Input<bool> Remove { get; set; } = null!;

        [Input("replace", required: true)]
        public Input<bool> Replace { get; set; } = null!;

        [Input("request", required: true)]
        public Input<bool> Request { get; set; } = null!;

        [Input("requestAdapt", required: true)]
        public Input<bool> RequestAdapt { get; set; } = null!;

        [Input("reset", required: true)]
        public Input<bool> Reset { get; set; } = null!;

        [Input("response", required: true)]
        public Input<bool> Response { get; set; } = null!;

        [Input("responseAdapt", required: true)]
        public Input<bool> ResponseAdapt { get; set; } = null!;

        [Input("scheme", required: true)]
        public Input<string> Scheme { get; set; } = null!;

        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        [Input("select", required: true)]
        public Input<bool> Select { get; set; } = null!;

        [Input("serverSsl", required: true)]
        public Input<bool> ServerSsl { get; set; } = null!;

        [Input("setVariable", required: true)]
        public Input<bool> SetVariable { get; set; } = null!;

        [Input("shutdown", required: true)]
        public Input<bool> Shutdown { get; set; } = null!;

        [Input("snat", required: true)]
        public Input<string> Snat { get; set; } = null!;

        [Input("snatpool", required: true)]
        public Input<string> Snatpool { get; set; } = null!;

        [Input("sourceAddress", required: true)]
        public Input<bool> SourceAddress { get; set; } = null!;

        [Input("sslClientHello", required: true)]
        public Input<bool> SslClientHello { get; set; } = null!;

        [Input("sslServerHandshake", required: true)]
        public Input<bool> SslServerHandshake { get; set; } = null!;

        [Input("sslServerHello", required: true)]
        public Input<bool> SslServerHello { get; set; } = null!;

        [Input("sslSessionId", required: true)]
        public Input<bool> SslSessionId { get; set; } = null!;

        [Input("status", required: true)]
        public Input<int> Status { get; set; } = null!;

        [Input("tcl", required: true)]
        public Input<bool> Tcl { get; set; } = null!;

        [Input("tcpNagle", required: true)]
        public Input<bool> TcpNagle { get; set; } = null!;

        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        [Input("tmName", required: true)]
        public Input<string> TmName { get; set; } = null!;

        [Input("uie", required: true)]
        public Input<bool> Uie { get; set; } = null!;

        [Input("universal", required: true)]
        public Input<bool> Universal { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        [Input("virtual", required: true)]
        public Input<string> Virtual { get; set; } = null!;

        [Input("vlan", required: true)]
        public Input<string> Vlan { get; set; } = null!;

        [Input("vlanId", required: true)]
        public Input<int> VlanId { get; set; } = null!;

        [Input("wam", required: true)]
        public Input<bool> Wam { get; set; } = null!;

        [Input("write", required: true)]
        public Input<bool> Write { get; set; } = null!;

        public GetPolicyRuleActionInputArgs()
        {
        }
        public static new GetPolicyRuleActionInputArgs Empty => new GetPolicyRuleActionInputArgs();
    }
}
