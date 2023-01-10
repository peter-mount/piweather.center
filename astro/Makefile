# We don't build anything as this is a library so instead we perform any tests.

# List of packages to perform tests
MODULES = astro/coord astro/ephemeris astro/julian astro/sidereal astro/sun

all: clean
	@$(foreach MODULE,$(MODULES),$(GO_TEST) $(PACKAGE_PREFIX)/$(MODULE)&&)exit 0

clean:
	@$(GO) clean -testcache

dist:
