properties([
  buildDiscarder(
    logRotator(
      artifactDaysToKeepStr: '',
      artifactNumToKeepStr: '',
      daysToKeepStr: '',
      numToKeepStr: '10'
    )
  ),
  disableConcurrentBuilds(),
  disableResume(),
  pipelineTriggers([
    cron('H H * * *')
  ])
])
node("go") {
  stage( 'Checkout' ) {
    checkout scm
  }
  stage( 'Init' ) {
    sh 'make clean init test'
  }
  stage( 'aix' ) {
        sh 'make -f Makefile.gen aix_ppc64'
  }
  stage( 'darwin' ) {
    parallel(
      amd64: {
        sh 'make -f Makefile.gen darwin_amd64'
      },
      arm64: {
        sh 'make -f Makefile.gen darwin_arm64'
      }
    )
  }
  stage( 'dragonfly' ) {
        sh 'make -f Makefile.gen dragonfly_amd64'
  }
  stage( 'freebsd' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen freebsd_386'
      },
      amd64: {
        sh 'make -f Makefile.gen freebsd_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen freebsd_arm6'
      },
      arm64: {
        sh 'make -f Makefile.gen freebsd_arm64'
      },
      arm7: {
        sh 'make -f Makefile.gen freebsd_arm7'
      }
    )
  }
  stage( 'illumos' ) {
        sh 'make -f Makefile.gen illumos_amd64'
  }
  stage( 'linux' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen linux_386'
      },
      amd64: {
        sh 'make -f Makefile.gen linux_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen linux_arm6'
      },
      arm64: {
        sh 'make -f Makefile.gen linux_arm64'
      },
      arm7: {
        sh 'make -f Makefile.gen linux_arm7'
      },
      loong64: {
        sh 'make -f Makefile.gen linux_loong64'
      },
      mips: {
        sh 'make -f Makefile.gen linux_mips'
      },
      mips64: {
        sh 'make -f Makefile.gen linux_mips64'
      },
      mips64le: {
        sh 'make -f Makefile.gen linux_mips64le'
      },
      mipsle: {
        sh 'make -f Makefile.gen linux_mipsle'
      },
      ppc64: {
        sh 'make -f Makefile.gen linux_ppc64'
      },
      ppc64le: {
        sh 'make -f Makefile.gen linux_ppc64le'
      },
      riscv64: {
        sh 'make -f Makefile.gen linux_riscv64'
      },
      s390x: {
        sh 'make -f Makefile.gen linux_s390x'
      }
    )
  }
  stage( 'netbsd' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen netbsd_386'
      },
      amd64: {
        sh 'make -f Makefile.gen netbsd_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen netbsd_arm6'
      },
      arm64: {
        sh 'make -f Makefile.gen netbsd_arm64'
      },
      arm7: {
        sh 'make -f Makefile.gen netbsd_arm7'
      }
    )
  }
  stage( 'openbsd' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen openbsd_386'
      },
      amd64: {
        sh 'make -f Makefile.gen openbsd_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen openbsd_arm6'
      },
      arm64: {
        sh 'make -f Makefile.gen openbsd_arm64'
      },
      arm7: {
        sh 'make -f Makefile.gen openbsd_arm7'
      },
      mips64: {
        sh 'make -f Makefile.gen openbsd_mips64'
      }
    )
  }
  stage( 'plan9' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen plan9_386'
      },
      amd64: {
        sh 'make -f Makefile.gen plan9_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen plan9_arm6'
      },
      arm7: {
        sh 'make -f Makefile.gen plan9_arm7'
      }
    )
  }
  stage( 'solaris' ) {
        sh 'make -f Makefile.gen solaris_amd64'
  }
  stage( 'windows' ) {
    parallel(
      386: {
        sh 'make -f Makefile.gen windows_386'
      },
      amd64: {
        sh 'make -f Makefile.gen windows_amd64'
      },
      arm6: {
        sh 'make -f Makefile.gen windows_arm6'
      },
      arm64: {
        sh 'make -f Makefile.gen windows_arm64'
      },
      arm7: {
        sh 'make -f Makefile.gen windows_arm7'
      }
    )
  }
}
