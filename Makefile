# Makefile for piweather.center

# The repository name/package prefix.
# This should match the value of module in go.mod
export PACKAGE_PREFIX = github.com/peter-mount/piweather.center

# List of modules to build.
#
# This list does not include tools as that's entered last during the build
# to ensure tests are run.
MODULES		= astro image log util

# Platforms to build.
#
# This is an array of os:architecture:armVersion
#
# armVersion is usually "" as it's only used for the "arm" architecture and can have values: 5, 6 or 7
#export PLATFORMS ?=

export PLATFORMSold ?= \
	linux:amd64: linux:arm64: linux:arm:6 linux:arm:7 linux:386: linux:s390x: linux:riscv64: \
	darwin:amd64: darwin:arm64: \
	freebsd:amd64: freebsd:arm64: freebsd:arm: \
	netbsd:amd64: netbsd:arm64: netbsd:arm: \
	openbsd:amd64: openbsd:arm64: openbsd:arm: \
	windows:amd64: windows:arm64:

# Where to place build artifacts
export BUILDS 	= $(shell pwd)/builds
export DIST		= $(shell pwd)/dist

# Tool names
export CP     	= @cp -p
export ECHO		= echo
export GO		= go
export MKDIR  	= @mkdir -p
export TAR		= tar

# Append -test.v to GO_TEST to show status of each test.
# Without it, only shows total time per module if they pass
export GO_TEST	?= $(GO) test

.PHONY: all build clean dist init test validate-go-version resolve-platforms

# Used to separate commands in foreach.
# NOTE this MUST have 2 empty lines between define and endef for it to work!
define \n


endef

all: init
	@$(MKDIR) -pv $(BUILDS)
	$(foreach MODULE,$(MODULES) tools,@$(MAKE) -C $(MODULE) all${\n})

clean:
	@$(GO) clean -testcache
	@$(RM) -r $(BUILDS) $(DIST)

dist: all
	@$(MKDIR) -pv $(DIST)
	$(foreach MODULE,$(MODULES) tools,@$(MAKE) -C $(MODULE) dist${\n})

init: validate-go-version resolve-platforms
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

# This discovers all platforms supported by the locally installed go compiler.
# This will only expand then if the PLATFORMS environment variable was not set
# when invoking make
resolve-platforms:
ifeq ("$(PLATFORMS)","")
	$(eval DISC_PLATFORMS=)
	$(foreach DISC_PLATFORM,$(shell $(GO) tool dist list), \
		$(eval GOOS=$(word 1,$(subst /, ,$(DISC_PLATFORM)))) \
		$(if $(filter android,$(GOOS)),,\
			$(if $(filter ios,$(GOOS)),,\
				$(eval GOARCH=$(word 2,$(subst /, ,$(DISC_PLATFORM)))) \
				$(foreach GOARM, \
					$(if $(filter arm,$(GOARCH)),6 7,:), \
					$(eval DISC_PLATFORMS=$(DISC_PLATFORMS) $(GOOS):$(GOARCH):$(GOARM)) \
				) \
			)\
		)\
	)
	$(eval export PLATFORMS=$(DISC_PLATFORMS))
endif

#	$(eval GO_PLATFORMS=$(shell go tool dist list)) \
