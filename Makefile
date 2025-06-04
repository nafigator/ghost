.SILENT: help
.DEFAULT_GOAL:=help

export UID:=$(shell id -u)
export GID:=$(shell id -g)
export PROJECT:=ghost

ifndef PROJECT_VERSION
	PROJECT_VERSION:=$(shell git tag -l | tail -n 1)

	# If project haven't any tags
	ifeq ($(PROJECT_VERSION),)
		PROJECT_VERSION:=develop
	endif

	export PROJECT_VERSION
endif

ifndef PROJECT_REVISION
	export PROJECT_REVISION:=$(shell git log -n 1 --format=%h)
endif

ifndef IMAGE_TAG
	export IMAGE_TAG:=$(shell git tag --sort=version:refname | tail -n1 | sed -e 's/v//')
endif

ifndef GOPATH
	export GOPATH:=$(HOME)/.local/go
endif

$(shell test -d $(GOPATH) || mkdir -p $(GOPATH))

ifndef GOOS
	export GOOS:=linux
endif

ifndef GOARCH
	export GOARCH:=amd64
endif

ifndef GOAMD64
	export GOAMD64:=v2
endif

$(shell test -e $(HOME)/.ssh/conifg || touch $(HOME)/.ssh/conifg)
$(shell test -e $(HOME)/.ssh/known_hosts || touch $(HOME)/.ssh/known_hosts)
$(shell test -e $(HOME)/.giconfig || touch $(HOME)/.gitconfig)
$(shell test -d $(GOPATH)/pkg || mkdir -p $(GOPATH)/pkg)
$(shell test -d $(GOPATH)/mod || mkdir -p $(GOPATH)/mod)
$(shell test -f $(ZAP_CONFIG_PATH) || cp $(ZAP_CONFIG_PATH).orig $(ZAP_CONFIG_PATH))

# https://docs.docker.com/desktop/mac/networking/#ssh-agent-forwarding
ifeq ($(PLATFORM),Darwin)
	export CACHE_DIR:=$(HOME)/Library/Caches
	export SSH_AUTH_SOCK_PATH:=/run/host-services/ssh-auth.sock
else
	export CACHE_DIR:=$(HOME)/.cache
	export SSH_AUTH_SOCK_PATH:=$(SSH_AUTH_SOCK)
endif

$(shell test -d $(CACHE_DIR) || mkdir -p $(CACHE_DIR))

ifndef GOCACHE
	export GOCACHE:=$(CACHE_DIR)/go-build
endif

$(shell test -d $(GOCACHE) || mkdir -p $(GOCACHE))
$(shell test -d $(HOME)/.config || mkdir -p $(HOME)/.config)

export BUILD_TIME:=$(shell date +'%F %T %Z')

export DOCKER_MOUNT_POINT:=/go/src/github.com/nafigator/$(PROJECT)
export GO_IMAGE:=nafigat0r/go:1.24.3
export LINTER_IMAGE:=nafigat0r/golangci-lint:2.0.2
export TRIVY_IMAGE:=aquasec/trivy:0.63.0
export GOVULNCHECK_IMAGE:=nafigat0r/govulncheck:1.1.4
export LD_FLAGS:='-s -w \
	-extldflags=-static \
	-X "github.com/nafigator/ghost/internal/app.build=$(PROJECT_VERSION), rev.$(PROJECT_REVISION), build time: $(BUILD_TIME)"'

export GO_DOCKER_PARAMS:="-u $(UID):$(GID) \
	-e HOME=$(DOCKER_MOUNT_POINT) \
	-e XDG_CONFIG_HOME=/var/config \
	-e XDG_CACHE_HOME=/var/cache \
	-e GOCACHE=/var/cache/go-build \
	-e GOLANGCI_LINT_CACHE=/var/cache/golangci-lint \
	-e CGO_ENABLED=0 \
	-e GOAMD64=$(GOAMD64) \
	-e SSH_AUTH_SOCK=/run/ssh-agent.sock \
	-v $(SSH_AUTH_SOCK_PATH):/run/ssh-agent.sock \
	-v $(HOME)/.ssh/config:/etc/ssh/ssh_config \
	-v /etc/passwd:/etc/passwd:ro \
	-v /etc/group:/etc/group:ro \
	-v $(HOME)/.ssh/known_hosts:/etc/ssh/ssh_known_hosts \
	-v $(HOME)/.gitconfig:/etc/gitconfig \
	-v $(HOME)/.config:/var/config \
	-v $(GOPATH)/pkg:/go/pkg:Z \
	-v $(GOPATH)/mod:/go/mod:Z \
	-v $(CURDIR):$(DOCKER_MOUNT_POINT) \
	-v $(CACHE_DIR):/var/cache \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w $(DOCKER_MOUNT_POINT) \
	--network host"

DOCKER_RUN_INTERACTIVE:=docker run -ti --rm \
	"$(GO_DOCKER_PARAMS)" \
	$(GO_IMAGE)

DOCKER_BASH_INTERACTIVE:=docker run -ti --rm \
	"$(GO_DOCKER_PARAMS)" \
	--entrypoint='/bin/bash' \
	$(GO_IMAGE)

.PHONY: dc
dc: #? Docker Command
	@echo "Run docker command: $(command)"
	@$(DOCKER_RUN_INTERACTIVE) $(command)

.PHONY: deps
deps: tidy #? Run go mod tidy and vendor
	@echo "Run go mod vendor..."
	@$(DOCKER_BASH_INTERACTIVE) -c "go mod vendor"

.PHONY: ps
ps: #? Show current running containers
	@docker-compose ps

.PHONY: tidy
tidy: #? Run go mod tidy
	@echo "Run go mod tidy..."
	@$(DOCKER_RUN_INTERACTIVE) mod tidy

.PHONY: log
log: #? Container log
	@docker logs -f $(PROJECT)

.PHONY: lint
lint: #? Run Go linter
	@echo "Running golangci-lint..."
	@docker run -ti --rm \
		"$(GO_DOCKER_PARAMS)" \
		$(LINTER_IMAGE) run

.PHONY: fix-alignment
fix-alignment: #? Fix linter alignment issues
	@echo "Running golangci-lint..."
	@docker run -ti --rm \
		"$(GO_DOCKER_PARAMS)" \
		$(LINTER_IMAGE) run --enable-only govet --fix

.PHONY: build
build: deps #? Build binary
	@echo "Build binary..."
	@echo "Environment: GOOS:$(GOOS), GOARCH:$(GOARCH), GOAMD64:$(GOAMD64)"
	@$(DOCKER_RUN_INTERACTIVE) build \
		-ldflags=$(LD_FLAGS) \
		-o $(DOCKER_MOUNT_POINT)/bin/ghost \
		$(DOCKER_MOUNT_POINT)/cmd/main.go

.PHONY: image
image: #? Build docker image
	@echo "Build docker image..."
	@DOCKER_BUILDKIT=1 docker build \
		--progress=plain \
		--force-rm \
		--no-cache \
		--build-arg LD_FLAGS=$(LD_FLAGS) \
		--build-arg PROJECT_REVISION=$(PROJECT_REVISION) \
		--tag nafigat0r/ghost:$(IMAGE_TAG) \
		--file .docker/Dockerfile .
	@docker image prune -f --filter label=stage=builder

.PHONY: trivy
trivy: deps #? Security checks by trivy
	@echo "Check Go project..."
	@docker run --rm -ti \
		"$(GO_DOCKER_PARAMS)" \
		$(TRIVY_IMAGE) --cache-dir=/var/cache fs ./

.PHONY: govulncheck
govulncheck: deps #? Security checks by govulncheck
	@echo "Check Go project..."
	@docker run --rm -ti \
		"$(GO_DOCKER_PARAMS)" \
		-v $(CURDIR):/var/govulncheck \
		$(GOVULNCHECK_IMAGE) -show verbose ./...

.PHONY: help
help: #? Show this message
	@printf "\033[34;01mApplication management:\033[0m\n"
	@awk '/^@?[a-zA-Z\-\_0-9]+:/ { \
		nb = sub( /^#\? /, "", helpMsg ); \
		if(nb == 0) { \
			helpMsg = $$0; \
			nb = sub( /^[^:]*:.* #\? /, "", helpMsg ); \
		} \
		if (nb) \
			printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg; \
		} \
		{ helpMsg = $$0 }' \
	$(MAKEFILE_LIST) | column -ts: