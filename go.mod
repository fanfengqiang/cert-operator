module github.com/fanfengqiang/cert-operator

require (
	contrib.go.opencensus.io/exporter/ocagent v0.4.9 // indirect
	github.com/Azure/azure-sdk-for-go v31.1.0+incompatible // indirect
	github.com/Azure/go-autorest v11.5.2+incompatible // indirect
	github.com/JamesClonk/vultr v2.0.1+incompatible // indirect
	github.com/OpenDNS/vegadns2client v0.0.0-20180418235048-a3fa4a771d87 // indirect
	github.com/akamai/AkamaiOPEN-edgegrid-golang v0.8.0 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190711074514-48dbb0a91dbc // indirect
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/aws/aws-sdk-go v1.20.18 // indirect
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.9.2 // indirect
	github.com/coreos/prometheus-operator v0.26.0 // indirect
	github.com/cpu/goacmedns v0.0.1 // indirect
	github.com/decker502/dnspod-go v0.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/dnsimple/dnsimple-go v0.30.0 // indirect
	github.com/emicklei/go-restful v2.8.1+incompatible // indirect
	github.com/exoscale/egoscale v0.18.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-acme/lego v2.6.1-0.20190624180855-ac65f6c6a9a4+incompatible
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-logr/logr v0.1.0 // indirect
	github.com/go-logr/zapr v0.1.0 // indirect
	github.com/go-openapi/spec v0.18.0
	github.com/golang/groupcache v0.0.0-20180924190550-6f2cf27854a4 // indirect
	github.com/golang/mock v1.2.0 // indirect
	github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20190318015731-ff9851476e98 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.8.5 // indirect
	github.com/iij/doapi v0.0.0-20190504054126-0bbf12d6d7df // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/kolo/xmlrpc v0.0.0-20190514182600-74b23a09d7ea // indirect
	github.com/labbsr0x/bindman-dns-webhook v1.0.0 // indirect
	github.com/labbsr0x/goh v0.0.0-20190610190554-60aa50bcbca7 // indirect
	github.com/linode/linodego v0.10.0 // indirect
	github.com/miekg/dns v1.1.15 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/namedotcom/go v0.0.0-20180403034216-08470befbe04 // indirect
	github.com/nrdcg/auroradns v1.0.0 // indirect
	github.com/nrdcg/goinwx v0.6.0 // indirect
	github.com/nrdcg/namesilo v0.2.0 // indirect
	github.com/operator-framework/operator-sdk v0.8.2-0.20190522220659-031d71ef8154
	github.com/oracle/oci-go-sdk v5.14.0+incompatible // indirect
	github.com/ovh/go-ovh v0.0.0-20181109152953-ba5adb4cf014 // indirect
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/sacloud/libsacloud v1.26.1 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/timewasted/linode v0.0.0-20160829202747-37e84520dcf7 // indirect
	github.com/transip/gotransip v5.8.2+incompatible // indirect
	go.opencensus.io v0.19.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	gopkg.in/ns1/ns1-go.v2 v2.0.0-20190703192230-737a440630af // indirect
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
	k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go v2.0.0-alpha.0.0.20181126152608-d082d5923d3c+incompatible
	k8s.io/code-generator v0.0.0-20180823001027-3dcf91f64f63
	k8s.io/gengo v0.0.0-20190128074634-0689ccc1d7d6
	k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
	sigs.k8s.io/controller-runtime v0.1.10
	sigs.k8s.io/controller-tools v0.1.10
	sigs.k8s.io/testing_frameworks v0.1.0 // indirect
)

// Pinned to kubernetes-1.13.1
replace (
	k8s.io/api => k8s.io/api v0.0.0-20181213150558-05914d821849
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/client-go => k8s.io/client-go v0.0.0-20181213151034-8d9ed539ba31
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.8.1
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20181117043124-c2090bec4d9b
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.10
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)
