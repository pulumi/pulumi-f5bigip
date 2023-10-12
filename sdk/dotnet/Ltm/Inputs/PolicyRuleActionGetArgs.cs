// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class PolicyRuleActionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        [Input("application")]
        public Input<string>? Application { get; set; }

        [Input("asm")]
        public Input<bool>? Asm { get; set; }

        [Input("avr")]
        public Input<bool>? Avr { get; set; }

        [Input("cache")]
        public Input<bool>? Cache { get; set; }

        [Input("carp")]
        public Input<bool>? Carp { get; set; }

        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("classify")]
        public Input<bool>? Classify { get; set; }

        [Input("clonePool")]
        public Input<string>? ClonePool { get; set; }

        [Input("code")]
        public Input<int>? Code { get; set; }

        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        [Input("connection")]
        public Input<bool>? Connection { get; set; }

        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("cookieHash")]
        public Input<bool>? CookieHash { get; set; }

        [Input("cookieInsert")]
        public Input<bool>? CookieInsert { get; set; }

        [Input("cookiePassive")]
        public Input<bool>? CookiePassive { get; set; }

        [Input("cookieRewrite")]
        public Input<bool>? CookieRewrite { get; set; }

        [Input("decompress")]
        public Input<bool>? Decompress { get; set; }

        [Input("defer")]
        public Input<bool>? Defer { get; set; }

        [Input("destinationAddress")]
        public Input<bool>? DestinationAddress { get; set; }

        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        [Input("expirySecs")]
        public Input<int>? ExpirySecs { get; set; }

        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("extension")]
        public Input<string>? Extension { get; set; }

        [Input("facility")]
        public Input<string>? Facility { get; set; }

        [Input("forward")]
        public Input<bool>? Forward { get; set; }

        [Input("fromProfile")]
        public Input<string>? FromProfile { get; set; }

        [Input("hash")]
        public Input<bool>? Hash { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("http")]
        public Input<bool>? Http { get; set; }

        [Input("httpBasicAuth")]
        public Input<bool>? HttpBasicAuth { get; set; }

        [Input("httpCookie")]
        public Input<bool>? HttpCookie { get; set; }

        [Input("httpHeader")]
        public Input<bool>? HttpHeader { get; set; }

        [Input("httpHost")]
        public Input<bool>? HttpHost { get; set; }

        [Input("httpReferer")]
        public Input<bool>? HttpReferer { get; set; }

        [Input("httpReply")]
        public Input<bool>? HttpReply { get; set; }

        [Input("httpSetCookie")]
        public Input<bool>? HttpSetCookie { get; set; }

        [Input("httpUri")]
        public Input<bool>? HttpUri { get; set; }

        [Input("ifile")]
        public Input<string>? Ifile { get; set; }

        [Input("insert")]
        public Input<bool>? Insert { get; set; }

        [Input("internalVirtual")]
        public Input<string>? InternalVirtual { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("l7dos")]
        public Input<bool>? L7dos { get; set; }

        [Input("length")]
        public Input<int>? Length { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("log")]
        public Input<bool>? Log { get; set; }

        [Input("ltmPolicy")]
        public Input<bool>? LtmPolicy { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        [Input("nexthop")]
        public Input<string>? Nexthop { get; set; }

        [Input("node")]
        public Input<string>? Node { get; set; }

        [Input("offset")]
        public Input<int>? Offset { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("pem")]
        public Input<bool>? Pem { get; set; }

        [Input("persist")]
        public Input<bool>? Persist { get; set; }

        [Input("pin")]
        public Input<bool>? Pin { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("pool")]
        public Input<string>? Pool { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("profile")]
        public Input<string>? Profile { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("queryString")]
        public Input<string>? QueryString { get; set; }

        [Input("rateclass")]
        public Input<string>? Rateclass { get; set; }

        [Input("redirect")]
        public Input<bool>? Redirect { get; set; }

        [Input("remove")]
        public Input<bool>? Remove { get; set; }

        [Input("replace")]
        public Input<bool>? Replace { get; set; }

        [Input("request")]
        public Input<bool>? Request { get; set; }

        [Input("requestAdapt")]
        public Input<bool>? RequestAdapt { get; set; }

        [Input("reset")]
        public Input<bool>? Reset { get; set; }

        [Input("response")]
        public Input<bool>? Response { get; set; }

        [Input("responseAdapt")]
        public Input<bool>? ResponseAdapt { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("script")]
        public Input<string>? Script { get; set; }

        [Input("select")]
        public Input<bool>? Select { get; set; }

        [Input("serverSsl")]
        public Input<bool>? ServerSsl { get; set; }

        [Input("setVariable")]
        public Input<bool>? SetVariable { get; set; }

        [Input("shutdown")]
        public Input<bool>? Shutdown { get; set; }

        [Input("snat")]
        public Input<string>? Snat { get; set; }

        [Input("snatpool")]
        public Input<string>? Snatpool { get; set; }

        [Input("sourceAddress")]
        public Input<bool>? SourceAddress { get; set; }

        [Input("sslClientHello")]
        public Input<bool>? SslClientHello { get; set; }

        [Input("sslServerHandshake")]
        public Input<bool>? SslServerHandshake { get; set; }

        [Input("sslServerHello")]
        public Input<bool>? SslServerHello { get; set; }

        [Input("sslSessionId")]
        public Input<bool>? SslSessionId { get; set; }

        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("tcl")]
        public Input<bool>? Tcl { get; set; }

        [Input("tcpNagle")]
        public Input<bool>? TcpNagle { get; set; }

        [Input("text")]
        public Input<string>? Text { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("tmName")]
        public Input<string>? TmName { get; set; }

        [Input("uie")]
        public Input<bool>? Uie { get; set; }

        [Input("universal")]
        public Input<bool>? Universal { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        [Input("virtual")]
        public Input<string>? Virtual { get; set; }

        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        [Input("wam")]
        public Input<bool>? Wam { get; set; }

        [Input("write")]
        public Input<bool>? Write { get; set; }

        public PolicyRuleActionGetArgs()
        {
        }
        public static new PolicyRuleActionGetArgs Empty => new PolicyRuleActionGetArgs();
    }
}
