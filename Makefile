# Makefile for piweather.center

# The repository name/package prefix.
# This should match the value of module in go.mod
export PACKAGE_PREFIX = github.com/peter-mount/piweather.center

# List of modules to build.
#
# This list does not include tools as that's entered last during the build
# as we build that one once for each platform listed in PLATFORMS below.
MODULES		= astro image log util

# Platforms to build.
# This is an array of os:architecture:armVersion
PLATFORMS	   ?= linux:amd64: linux:arm64: linux:arm:6 linux:arm:7

# Where to place build artifacts
export BUILDS 	= $(shell pwd)/builds
export DIST		= $(shell pwd)/dist

# Tool names
export CP     	= @cp -p
export ECHO		= echo
export GO		= go
export MKDIR  	= @mkdir -p -v
export TAR		= tar

# Append -test.v to GO_TEST to show status of each test.
# Without it, only shows total time per module if they pass
export GO_TEST	?= $(GO) test

.PHONY: all build clean dist init test validate-go-version

all: init
	@$(MKDIR) -pv $(BUILDS)
	@$(foreach MODULE,$(MODULES),$(MAKE) -C $(MODULE) all&&)exit 0
	@$(foreach PLATFORM,$(PLATFORMS), \
		GOOS=$(word 1,$(subst :, ,$(PLATFORM))) \
		GOARCH=$(word 2,$(subst :, ,$(PLATFORM))) \
		GOARM=$(word 3,$(subst :, ,$(PLATFORM))) \
		$(MAKE) -C tools &&\
	)$(ECHO) -n

clean:
	@$(GO) clean -testcache
	@$(foreach MODULE,$(MODULES) tools,$(MAKE) -C $(MODULE) clean;)
	@$(RM) -r $(BUILDS) $(DIST)

dist: all
	@$(MKDIR) -pv $(DIST)
	@$(foreach MODULE,$(MODULES) tools,$(MAKE) -C $(MODULE) dist;)
	@$(foreach PLATFORM,$(PLATFORMS), \
		GOOS=$(word 1,$(subst :, ,$(PLATFORM))) \
		GOARCH=$(word 2,$(subst :, ,$(PLATFORM))) \
		GOARM=$(word 3,$(subst :, ,$(PLATFORM))) \
		$(MAKE) -C tools dist;\
	)

init: validate-go-version
	@$(GO) mod download

# Validates the installed version of go against the version declared in go.mod
MINIMUM_SUPPORTED_GO_MAJOR_VERSION	= $(shell grep "^go" go.mod | cut -f2 -d' ' | cut -f1 -d'.')
MINIMUM_SUPPORTED_GO_MINOR_VERSION	= $(shell grep "^go" go.mod | cut -f2 -d' ' | cut -f2 -d'.')
GO_MAJOR_VERSION = $(shell $(GO) version | cut -f3 -d' ' | cut -c 3- | cut -f1 -d' ' | cut -f1 -d'.')
GO_MINOR_VERSION = $(shell $(GO) version | cut -f3 -d' ' | cut -c 3- | cut -f1 -d' ' | cut -f2 -d'.')
GO_VERSION_VALIDATION_ERR_MSG = Golang version $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION) is not supported, please update to at least $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION).$(MINIMUM_SUPPORTED_GO_MINOR_VERSION)
validate-go-version:
	@if [ $(GO_MAJOR_VERSION) -gt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		exit 0 ;\
	elif [ $(GO_MAJOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MINOR_VERSION) ] ; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	fi
