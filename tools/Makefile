# Where builds for a specific platform are placed
export BUILD	= $(BUILDS)/$(GOOS)/$(GOARCH)$(GOARM)

# Distribution archives
DIST_TAR		= $(BUILDS)/$(GOOS)_$(GOARCH)$(GOARM).tgz

# Common GO build for a specific platform
export GO_BUILD = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) $(GO) build

MODULES			= cloud
all:
	@$(MKDIR) $(BUILD)
	@$(foreach MODULE,$(MODULES),$(MAKE) -C $(MODULE) all;)

clean:
	@$(foreach MODULE,$(MODULES),$(MAKE) -C $(MODULE) clean;)
	@$(RM) $(DIST_TAR)

dist: $(DIST_TAR)

# TAR archive with some kungfu to change the path in the archive to something
# more useful
$(DIST_TAR):
	@(ECHO) "Generating $@"
	@(TAR) -P --transform "s|^$(BUILD)|piweather.center|" -cvzpf $(DIST_TAR) $(BUILD)
