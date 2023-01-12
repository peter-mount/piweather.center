
MODULES	= cloud ephemeris videogenerator

.PHONY:	all dist

# Used to separate commands in foreach.
# NOTE this MUST have 2 empty lines between define and endef for it to work!
define \n


endef

all:
	$(foreach PLATFORM,$(PLATFORMS), \
		$(eval GOOS=$(word 1,$(subst :, ,$(PLATFORM)))) \
		$(eval GOARCH=$(word 2,$(subst :, ,$(PLATFORM)))) \
		$(eval GOARM=$(word 3,$(subst :, ,$(PLATFORM)))) \
		$(eval BUILD=$(BUILDS)/$(GOOS)/$(GOARCH)$(GOARM)) \
		@(echo "Platform $(GOOS) $(GOARCH)$(GOARM)")${\n}\
		@$(MKDIR) $(BUILD)${\n}\
		$(foreach MODULE,$(MODULES),\
			@printf "    Compile %s\n" $(MODULE)${\n}\
			@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) $(GO) build -o $(BUILD)/$(MODULE) $(MODULE)/bin/main.go${\n}\
		)\
	)

dist:
	$(foreach PLATFORM,$(PLATFORMS), \
		$(eval GOOS=$(word 1,$(subst :, ,$(PLATFORM)))) \
		$(eval GOARCH=$(word 2,$(subst :, ,$(PLATFORM)))) \
		$(eval GOARM=$(word 3,$(subst :, ,$(PLATFORM)))) \
		$(eval BUILD=$(BUILDS)/$(GOOS)/$(GOARCH)$(GOARM)) \
		$(eval DIST_TAR=$(GOOS)_$(GOARCH)$(GOARM).tgz) \
		@printf "Generating %s\n" $(DIST_TAR)${\n}\
		@$(TAR) -P --transform "s|^$(BUILD)|piweather.center|" -czpf $(DIST)/$(DIST_TAR) $(BUILD)${\n}\
	)
