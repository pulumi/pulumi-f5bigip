module github.com/pulumi/pulumi-f5bigip

go 1.12

require (
	github.com/chzyer/logex v1.1.11-0.20160617073814-96a4d311aa9b // indirect
	github.com/coreos/bbolt v1.3.1-coreos.1 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/hashicorp/aws-sdk-go-base v0.3.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.2-0.20171022020050-c20a0b1b1ea9 // indirect
	github.com/hashicorp/terraform v0.12.0
	github.com/miekg/dns v1.0.14 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.22-0.20190702185104-ebceea93a5da
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190709052202-629f7c54269d
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-bigip v0.12.3
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)
