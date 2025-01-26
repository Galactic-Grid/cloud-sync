# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOBIN=$(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN=$(shell go env GOPATH)/bin
endif

.PHONY: all
all: build

.PHONY: generate
generate: controller-gen ## Generate code
	$(CONTROLLER_GEN) object paths=./...

.PHONY: build
build: generate ## Build the project
	go build -o bin/manager ./...

.PHONY: controller-gen
controller-gen: ## Download controller-gen if necessary
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@latest)

CONTROLLER_GEN=$(GOBIN)/controller-gen

# go-get-tool will 'go install' any package with custom version if needed
define go-get-tool
@[ -f $(1) ] || { \
set -e; \
echo "Downloading $(2)"; \
GOBIN=$(GOBIN) go install $(2); \
}
endef
