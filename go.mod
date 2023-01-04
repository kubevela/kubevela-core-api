module github.com/oam-dev/kubevela-core-api

go 1.16

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	cuelang.org/go v0.5.0-alpha.1
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210722005935-0b469fcc77cd
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/google/go-cmp v0.5.8
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/oam-dev/cluster-gateway v1.4.0
	github.com/oam-dev/terraform-controller v0.7.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.20.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be // indirect
	golang.org/x/net v0.0.0-20220906165146-f3363e06e74c // indirect
	golang.org/x/oauth2 v0.0.0-20220622183110-fd043fe589d2 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	google.golang.org/genproto v0.0.0-20220628213854-d9e0b6570c03 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	k8s.io/api v0.23.6
	k8s.io/apiextensions-apiserver v0.23.6
	k8s.io/apimachinery v0.23.6
	k8s.io/client-go v0.23.6
	k8s.io/klog/v2 v2.60.1
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
	open-cluster-management.io/api v0.7.0 // indirect
	sigs.k8s.io/apiserver-runtime v1.1.1 // indirect
	sigs.k8s.io/controller-runtime v0.11.2
	sigs.k8s.io/controller-tools v0.6.2
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/yaml v1.3.0
)

replace sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.24
