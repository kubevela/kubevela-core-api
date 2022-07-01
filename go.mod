module github.com/oam-dev/kubevela-core-api

go 1.16

require (
	cuelang.org/go v0.2.2
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210722005935-0b469fcc77cd
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-cmp v0.5.5
	github.com/oam-dev/cluster-gateway v1.4.0
	github.com/oam-dev/terraform-controller v0.4.6
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	k8s.io/api v0.23.1
	k8s.io/apiextensions-apiserver v0.23.0
	k8s.io/apimachinery v0.23.1
	k8s.io/client-go v0.23.1
	k8s.io/klog/v2 v2.30.0
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.2.8
	sigs.k8s.io/yaml v1.3.0
)

replace sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.24
