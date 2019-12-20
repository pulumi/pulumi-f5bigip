// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.F5BigIP
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("f5bigip");

        /// <summary>
        /// Domain name/IP of the BigIP
        /// </summary>
        public static string? Address { get; set; } = __config.Get("address");

        /// <summary>
        /// Login reference for token authentication (see BIG-IP REST docs for details)
        /// </summary>
        public static string? LoginRef { get; set; } = __config.Get("loginRef");

        /// <summary>
        /// The user's password
        /// </summary>
        public static string? Password { get; set; } = __config.Get("password");

        /// <summary>
        /// Management Port to connect to Bigip
        /// </summary>
        public static string? Port { get; set; } = __config.Get("port");

        /// <summary>
        /// Enable to use an external authentication source (LDAP, TACACS, etc)
        /// </summary>
        public static bool? TokenAuth { get; set; } = __config.GetBoolean("tokenAuth");

        /// <summary>
        /// Username with API access to the BigIP
        /// </summary>
        public static string? Username { get; set; } = __config.Get("username");

    }
    namespace ConfigTypes
    {
    }
}
