// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class ProfileHttpEnforcementArgs : global::Pulumi.ResourceArgs
    {
        [Input("knownMethods")]
        private InputList<string>? _knownMethods;

        /// <summary>
        /// Specifies which HTTP methods count as being known. Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. Default value is [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK].If no value is specified while creating, then default value will be assigned. In order to remove it, [""]  list is to be passed.
        /// </summary>
        public InputList<string> KnownMethods
        {
            get => _knownMethods ?? (_knownMethods = new InputList<string>());
            set => _knownMethods = value;
        }

        /// <summary>
        /// Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers.If no value is specified, then default value will be assigned.
        /// </summary>
        [Input("maxHeaderCount")]
        public Input<int>? MaxHeaderCount { get; set; }

        /// <summary>
        /// Specifies the maximum header size.The default value is 32768.If no string is specified, then default value will be assigned.
        /// </summary>
        [Input("maxHeaderSize")]
        public Input<int>? MaxHeaderSize { get; set; }

        /// <summary>
        /// Specifies whether to allow, reject or switch to pass-through mode when an unknown HTTP method is parsed. Default value is allow. If no string is specified, then default value will be assigned.
        /// </summary>
        [Input("unknownMethod")]
        public Input<string>? UnknownMethod { get; set; }

        public ProfileHttpEnforcementArgs()
        {
        }
        public static new ProfileHttpEnforcementArgs Empty => new ProfileHttpEnforcementArgs();
    }
}
