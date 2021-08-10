# Repo info
GIT_COMMIT      ?= git-$(shell git rev-parse --short HEAD)

ERR		= echo ${TIME} ${RED}[FAIL]${CNone}
OK		= echo ${TIME} ${GREEN}[ OK ]${CNone}

PROJECT_VERSION_VAR    := github.com/oam-dev/velacp/pkg/version.Version
PROJECT_GITVERSION_VAR := github.com/oam-dev/velacp/pkg/version.GitRevision
LDFLAGS             ?= "-X $(PROJECT_VERSION_VAR)=$(PROJECT_VERSION) -X $(PROJECT_GITVERSION_VAR)=$(GIT_COMMIT)"

TARGETS     := darwin/amd64 linux/amd64 windows/amd64
DIST_DIRS   := find * -name "*-*" -type d -exec

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

build: reviewable
	go build -o test/testapi ./test/main.go

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./test/... ./apis/...

# Run go vet against code
vet:
	go vet ./pkg/... ./test/... ./apis/...

reviewable: fmt vet
	go mod tidy


