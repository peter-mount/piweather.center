#
# By default this will build the project on every non-mobile platform
# supported by the installed go environment.
#
# To limit a build to a single environment, you can force it to just a
# single platform by prefixing make with:
#
# PLATFORMS=linux:amd64: make clean all
#
# Just change the entry for your OS and CPU. These are listed in platforms.md
#
# Note: For 32 bit arm processors the 3rd parameter is important.
# e.g. use linux:arm:6 or linux:arm:7
#
# For all other processors, including arm64, leave the third field blank.
#
# For a parallel builds you can use the -j parameter to make as usual.
#
# e.g.: make -j 8 clean all
#
# Pick a value suitable to the number of cores/thread your machine has.
# This is useful for a full build of all platforms as it will build all
# of the binaries in parallel speeding up the full build.
#

.PHONY: all clean init test build apt

all: init test build

init:
	@echo "GO MOD   tidy";go mod tidy
	@echo "GO MOD   download";go mod download
	@echo "GENERATE build";\
	CGO_ENABLED=0 go build -o build tools/build/bin/main.go
	@./build -build Makefile.gen -build-platform "$(PLATFORMS)" -d builds -dist dist -build-archiveArtifacts "dist/*" -block blocklist.yaml

clean: init
	@${MAKE} --no-print-directory -f Makefile.gen clean

test: init
	@${MAKE} --no-print-directory -f Makefile.gen test

build: init #test
	@${MAKE} --no-print-directory -f Makefile.gen all

# Builds apt packages for specific platforms
apt:
	@PLATFORMS="linux:amd64: linux:arm64: linux:arm:6 linux:arm:7" ${MAKE} --no-print-directory build
	#rm -rf builds/apt dist/*.deb
	@${MAKE} --no-print-directory apt-builds

apt-builds: dist/piweather-common.deb \
			dist/piweather-amd64.deb \
			dist/piweather-arm6.deb \
			dist/piweather-arm7.deb

dist/piweather-common.deb:
	# dpkg control files
	mkdir -pv builds/apt/common/DEBIAN
	cp -rp apt/common/* builds/apt/common/DEBIAN
	# Copyright - TODO format correctly for DEB
	mkdir -pv builds/apt/common/usr/share/doc/piweathercenter-common/
	cp LICENSE builds/apt/common/usr/share/doc/piweathercenter-common/copyright
	# Now common files
	mkdir -pv builds/apt/common/usr/share/piweather/
	cp -rp builds/linux/amd64/share/* builds/apt/common/usr/share/piweather/
	# temp hack until go-anim converted to new layout
	cp -rp builds/linux/amd64/lib/font builds/apt/common/usr/share/piweather/
	dpkg --build builds/apt/common dist/piweather-common.deb

dist/piweather-amd64.deb:
	# dpkg control files
	mkdir -pv builds/apt/amd64/DEBIAN
	cp -rp apt/amd64/* builds/apt/amd64/DEBIAN
	# Copyright - TODO format correctly for DEB
	mkdir -pv builds/apt/amd64/usr/share/doc/piweathercenter-amd64/
	cp LICENSE builds/apt/amd64/usr/share/doc/piweathercenter-amd64/copyright
	# Now common files
	mkdir -pv builds/apt/amd64/usr/bin
	cp -rp builds/linux/amd64/bin/* builds/apt/amd64/usr/bin/
	dpkg --build builds/apt/amd64 dist/piweather-amd64.deb

dist/piweather-arm6.deb:
	# dpkg control files
	mkdir -pv builds/apt/arm6/DEBIAN
	cp -rp apt/arm6/* builds/apt/arm6/DEBIAN
	# Copyright - TODO format correctly for DEB
	mkdir -pv builds/apt/arm6/usr/share/doc/piweathercenter-arm6/
	cp LICENSE builds/apt/arm6/usr/share/doc/piweathercenter-arm6/copyright
	# Now common files
	mkdir -pv builds/apt/arm6/usr/bin
	cp -rp builds/linux/arm6/bin/* builds/apt/arm6/usr/bin/
	dpkg --build builds/apt/arm6 dist/piweather-arm6.deb

dist/piweather-arm7.deb:
	# dpkg control files
	mkdir -pv builds/apt/arm7/DEBIAN
	cp -rp apt/arm7/* builds/apt/arm7/DEBIAN
	# Copyright - TODO format correctly for DEB
	mkdir -pv builds/apt/arm7/usr/share/doc/piweathercenter-arm7/
	cp LICENSE builds/apt/arm7/usr/share/doc/piweathercenter-arm7/copyright
	# Now common files
	mkdir -pv builds/apt/arm7/usr/bin
	cp -rp builds/linux/arm7/bin/* builds/apt/arm7/usr/bin/
	dpkg --build builds/apt/arm7 dist/piweather-arm7.deb
