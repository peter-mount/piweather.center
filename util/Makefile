# We don't build anything as this is a library so instead we perform any tests.

# List of packages to perform tests
MODULES = util/dir

# Used to separate commands in foreach.
# NOTE this MUST have 2 empty lines between define and endef for it to work!
define \n


endef

.PHONY: all dist

all:
	$(foreach MODULE,$(MODULES),@$(GO_TEST) $(PACKAGE_PREFIX)/$(MODULE)${\n})

dist:
