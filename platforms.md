# Supported Platforms

The following platforms are supported by virtual of how the build system works:

| Operating System | CPU Architectures |
| ---------------- | ----------------- |
| aix | ppc64 |
| darwin | amd64 arm64 |
| dragonfly | amd64 |
| freebsd | 386 amd64 arm6 arm64 arm7 |
| illumos | amd64 |
| js | wasm |
| linux | 386 amd64 arm6 arm64 arm7 loong64 mips mips64 mips64le mipsle ppc64 ppc64le riscv64 s390x |
| netbsd | 386 amd64 arm6 arm64 arm7 |
| openbsd | 386 amd64 arm6 arm64 arm7 mips64 |
| plan9 | 386 amd64 arm6 arm7 |
| solaris | amd64 |
| windows | 386 amd64 arm6 arm64 arm7 |

This is all non-mobile platforms supported by go version `1.19`

This page is automatically generated from the output of `go tool dist list`
