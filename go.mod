module github.com/oam-dev/kubevela-core-api

go 1.13

require (
	cuelang.org/go v0.2.2
	github.com/crossplane/crossplane-runtime v0.10.0
	github.com/davecgh/go-spew v1.1.1
	github.com/go-logr/logr v0.1.0
	github.com/google/go-cmp v0.5.2
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/controller-tools v0.2.4
)

// clint-go had a buggy release, https://github.com/kubernetes/client-go/issues/749
replace k8s.io/client-go => k8s.io/client-go v0.18.8
