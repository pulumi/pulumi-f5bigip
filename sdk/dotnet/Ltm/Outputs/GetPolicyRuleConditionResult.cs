// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Outputs
{

    [OutputType]
    public sealed class GetPolicyRuleConditionResult
    {
        public readonly bool Address;
        public readonly bool All;
        public readonly string AppService;
        public readonly bool BrowserType;
        public readonly bool BrowserVersion;
        public readonly bool CaseInsensitive;
        public readonly bool CaseSensitive;
        public readonly bool Cipher;
        public readonly bool CipherBits;
        public readonly bool ClientSsl;
        public readonly bool Code;
        public readonly bool CommonName;
        public readonly bool Contains;
        public readonly bool Continent;
        public readonly bool CountryCode;
        public readonly bool CountryName;
        public readonly bool CpuUsage;
        public readonly bool DeviceMake;
        public readonly bool DeviceModel;
        public readonly bool Domain;
        public readonly bool EndsWith;
        public readonly bool Equals;
        public readonly bool Expiry;
        public readonly bool Extension;
        public readonly bool External;
        public readonly bool Geoip;
        public readonly bool Greater;
        public readonly bool GreaterOrEqual;
        public readonly bool Host;
        public readonly bool HttpBasicAuth;
        public readonly bool HttpCookie;
        public readonly bool HttpHeader;
        public readonly bool HttpHost;
        public readonly bool HttpMethod;
        public readonly bool HttpReferer;
        public readonly bool HttpSetCookie;
        public readonly bool HttpStatus;
        public readonly bool HttpUri;
        public readonly bool HttpUserAgent;
        public readonly bool HttpVersion;
        public readonly int Index;
        public readonly bool Internal;
        public readonly bool Isp;
        public readonly bool Last15secs;
        public readonly bool Last1min;
        public readonly bool Last5mins;
        public readonly bool Less;
        public readonly bool LessOrEqual;
        public readonly bool Local;
        public readonly bool Major;
        public readonly bool Matches;
        public readonly bool Minor;
        public readonly bool Missing;
        public readonly bool Mss;
        public readonly bool Not;
        public readonly bool Org;
        public readonly bool Password;
        public readonly bool Path;
        public readonly bool PathSegment;
        public readonly bool Port;
        public readonly bool Present;
        public readonly bool Protocol;
        public readonly bool QueryParameter;
        public readonly bool QueryString;
        public readonly bool RegionCode;
        public readonly bool RegionName;
        public readonly bool Remote;
        public readonly bool Request;
        public readonly bool Response;
        public readonly bool RouteDomain;
        public readonly bool Rtt;
        public readonly bool Scheme;
        public readonly bool ServerName;
        public readonly bool SslCert;
        public readonly bool SslClientHello;
        public readonly bool SslExtension;
        public readonly bool SslServerHandshake;
        public readonly bool SslServerHello;
        public readonly bool StartsWith;
        public readonly bool Tcp;
        public readonly bool Text;
        public readonly string TmName;
        public readonly bool UnnamedQueryParameter;
        public readonly bool UserAgentToken;
        public readonly bool Username;
        public readonly bool Value;
        public readonly ImmutableArray<string> Values;
        public readonly bool Version;
        public readonly bool Vlan;
        public readonly bool VlanId;

        [OutputConstructor]
        private GetPolicyRuleConditionResult(
            bool address,

            bool all,

            string appService,

            bool browserType,

            bool browserVersion,

            bool caseInsensitive,

            bool caseSensitive,

            bool cipher,

            bool cipherBits,

            bool clientSsl,

            bool code,

            bool commonName,

            bool contains,

            bool continent,

            bool countryCode,

            bool countryName,

            bool cpuUsage,

            bool deviceMake,

            bool deviceModel,

            bool domain,

            bool endsWith,

            bool equals,

            bool expiry,

            bool extension,

            bool external,

            bool geoip,

            bool greater,

            bool greaterOrEqual,

            bool host,

            bool httpBasicAuth,

            bool httpCookie,

            bool httpHeader,

            bool httpHost,

            bool httpMethod,

            bool httpReferer,

            bool httpSetCookie,

            bool httpStatus,

            bool httpUri,

            bool httpUserAgent,

            bool httpVersion,

            int index,

            bool @internal,

            bool isp,

            bool last15secs,

            bool last1min,

            bool last5mins,

            bool less,

            bool lessOrEqual,

            bool local,

            bool major,

            bool matches,

            bool minor,

            bool missing,

            bool mss,

            bool not,

            bool org,

            bool password,

            bool path,

            bool pathSegment,

            bool port,

            bool present,

            bool protocol,

            bool queryParameter,

            bool queryString,

            bool regionCode,

            bool regionName,

            bool remote,

            bool request,

            bool response,

            bool routeDomain,

            bool rtt,

            bool scheme,

            bool serverName,

            bool sslCert,

            bool sslClientHello,

            bool sslExtension,

            bool sslServerHandshake,

            bool sslServerHello,

            bool startsWith,

            bool tcp,

            bool text,

            string tmName,

            bool unnamedQueryParameter,

            bool userAgentToken,

            bool username,

            bool value,

            ImmutableArray<string> values,

            bool version,

            bool vlan,

            bool vlanId)
        {
            Address = address;
            All = all;
            AppService = appService;
            BrowserType = browserType;
            BrowserVersion = browserVersion;
            CaseInsensitive = caseInsensitive;
            CaseSensitive = caseSensitive;
            Cipher = cipher;
            CipherBits = cipherBits;
            ClientSsl = clientSsl;
            Code = code;
            CommonName = commonName;
            Contains = contains;
            Continent = continent;
            CountryCode = countryCode;
            CountryName = countryName;
            CpuUsage = cpuUsage;
            DeviceMake = deviceMake;
            DeviceModel = deviceModel;
            Domain = domain;
            EndsWith = endsWith;
            Equals = equals;
            Expiry = expiry;
            Extension = extension;
            External = external;
            Geoip = geoip;
            Greater = greater;
            GreaterOrEqual = greaterOrEqual;
            Host = host;
            HttpBasicAuth = httpBasicAuth;
            HttpCookie = httpCookie;
            HttpHeader = httpHeader;
            HttpHost = httpHost;
            HttpMethod = httpMethod;
            HttpReferer = httpReferer;
            HttpSetCookie = httpSetCookie;
            HttpStatus = httpStatus;
            HttpUri = httpUri;
            HttpUserAgent = httpUserAgent;
            HttpVersion = httpVersion;
            Index = index;
            Internal = @internal;
            Isp = isp;
            Last15secs = last15secs;
            Last1min = last1min;
            Last5mins = last5mins;
            Less = less;
            LessOrEqual = lessOrEqual;
            Local = local;
            Major = major;
            Matches = matches;
            Minor = minor;
            Missing = missing;
            Mss = mss;
            Not = not;
            Org = org;
            Password = password;
            Path = path;
            PathSegment = pathSegment;
            Port = port;
            Present = present;
            Protocol = protocol;
            QueryParameter = queryParameter;
            QueryString = queryString;
            RegionCode = regionCode;
            RegionName = regionName;
            Remote = remote;
            Request = request;
            Response = response;
            RouteDomain = routeDomain;
            Rtt = rtt;
            Scheme = scheme;
            ServerName = serverName;
            SslCert = sslCert;
            SslClientHello = sslClientHello;
            SslExtension = sslExtension;
            SslServerHandshake = sslServerHandshake;
            SslServerHello = sslServerHello;
            StartsWith = startsWith;
            Tcp = tcp;
            Text = text;
            TmName = tmName;
            UnnamedQueryParameter = unnamedQueryParameter;
            UserAgentToken = userAgentToken;
            Username = username;
            Value = value;
            Values = values;
            Version = version;
            Vlan = vlan;
            VlanId = vlanId;
        }
    }
}
