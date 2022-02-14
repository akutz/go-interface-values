#!/usr/bin/env bash

#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

## --------------------------------------
## General
## --------------------------------------

SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=help

# Use GOPROXY environment variable if set
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://proxy.golang.org
endif
export GOPROXY

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

# The help will print out all targets with their descriptions organized below
# their categories. The categories are represented by `##@` and the target
# descriptions by `##`.
#
# The awk commands is responsible to read the entire set of makefiles included
# in this invocation, looking for lines of the file as xyz: ## something, and
# then pretty-format the target and help. Then, if there's a line with ##@
# something, that gets pretty-printed as a category.
# 
# More info over the usage of ANSI control characters for terminal
# formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
#
# More info over awk command: http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


IMAGE_NAME ?= go-interface-values
IMAGE_TAG  ?= latest
IMAGE      ?= $(IMAGE_NAME):$(IMAGE_TAG)
PLATFORMS  ?= linux/amd64,linux/arm64
PUSH_ALL   ?=
IMAGE_RUN_FLAGS ?= -it --rm

# DOCKER_TARGETS is a list of targets that will have the -docker
# option to run in the container with the working directory mounted
# into the container.
DOCKER_TARGETS := generate-svgs

# SANDBOX_TARGETS is a list of targets that will have the -sandbox
# option to run in the container.
SANDBOX_TARGETS := asm bench sizes test

RUN_IN_PREFIX  := docker run $(IMAGE_RUN_FLAGS)
RUN_IN_DOCKER  := $(RUN_IN_PREFIX) -v "$$(pwd):/go-interface-values" $(IMAGE)
RUN_IN_SANDBOX := $(RUN_IN_PREFIX) $(IMAGE)


## --------------------------------------
## Images
## --------------------------------------
.PHONY: image-build
image-build: ## Build the docker image
	docker build -t $(IMAGE) .

.PHONY: image-build-all
image-build-all: ## Build the docker image for multiple platforms
	docker buildx build -t $(IMAGE) --platform $(PLATFORMS) $(PUSH_ALL) .

.PHONY: image-push
image-push: ## Push the docker image
	docker push $(IMAGE)

.PHONY: image-push-all
image-push-all: PUSH_ALL=--push
image-push-all: image-build-all
image-push-all: ## Push the docker image for multiple platforms

.PHONY: image-run
image-run: ## Launch the docker image
	$(RUN_IN_SANDBOX)


## --------------------------------------
## Generate
## --------------------------------------

# The command line to invoke ditaa. In the Docker image this
# will be set to "java /ditaa.jar".
DITAA ?= ditaa

.PHONY: generate-tests
generate-tests: ## Generate the mem tests
	cd tests/mem && python3 ../../hack/gen.py

.PHONY: generate-svgs
generate-svgs: ## Generate the svgs
	@find . -name '*.ascii' -type f -print0 | \
	xargs -0n1 $(DITAA) -E -o --background FFFFFF --svg --font-family courier


## --------------------------------------
## Lint
## --------------------------------------
.PHONY: lint-markdown
lint-markdown: ## Lint the project's markdown
	@find . -name "*.md" -type f -print0 | \
	xargs -0 markdownlint -c .markdownlint.yaml


## --------------------------------------
## Testing
## --------------------------------------
GCFLAGS := -gcflags "-l -N"

.PHONY: test
test: ## Run tests
	go version && go test -count 1 -v -run "^Test" ./...

.PHONY: test-m
test-lem-m: ## Print optimizations
	go version && go test -count 1 -v -c -gcflags -m ./tests/lem

.PHONY: bench
bench: ## Run benchmarks
	go version && \
	go test \
	  -v \
	  -count 1 \
	  -benchtime 1000x \
	  -run Mem -benchmem \
	  -bench BenchmarkMem \
	  ./tests/mem | \
	python3 hack/b2md.py

.PHONY: asm
asm: ## Print asm table
	go version && \
	cd ./tests/mem && \
	go tool compile -S -wb=false *.go | \
	python3 ../../hack/asm2md.py


## --------------------------------------
## Clean
## --------------------------------------
FILE_EXT_TO_CLEAN := .a .o .out .profile .test
FIND_EXT_TO_CLEAN := $(foreach e,$(FILE_EXT_TO_CLEAN),-name '*$e'$(if $(filter-out $e,$(lastword $(FILE_EXT_TO_CLEAN))), -or,))

.PHONY: clean
clean: ## Clean up artifacts
	@find . -type f \( $(FIND_EXT_TO_CLEAN) \) -delete
	@go clean -i -testcache ./...


## --------------------------------------
## Meta
## --------------------------------------
# Set up the sandbox targets.
SANDBOX_TARGETS := $(addsuffix -sandbox,$(SANDBOX_TARGETS))
.PHONY: $(SANDBOX_TARGETS)
$(SANDBOX_TARGETS):
	$(RUN_IN_SANDBOX) make $(subst -sandbox,,$@)

# Set up the docker targets.
DOCKER_TARGETS := $(addsuffix -docker,$(DOCKER_TARGETS))
.PHONY: $(DOCKER_TARGETS)
$(DOCKER_TARGETS):
	$(RUN_IN_DOCKER) make $(subst -docker,,$@)

