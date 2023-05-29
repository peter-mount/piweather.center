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
    sh 'make -f Makefile.gen aix'
  }
  stage( 'darwin' ) {
    sh 'make -f Makefile.gen darwin'
  }
  stage( 'dragonfly' ) {
    sh 'make -f Makefile.gen dragonfly'
  }
  stage( 'freebsd' ) {
    sh 'make -f Makefile.gen freebsd'
  }
  stage( 'illumos' ) {
    sh 'make -f Makefile.gen illumos'
  }
  stage( 'linux' ) {
    sh 'make -f Makefile.gen linux'
  }
  stage( 'netbsd' ) {
    sh 'make -f Makefile.gen netbsd'
  }
  stage( 'openbsd' ) {
    sh 'make -f Makefile.gen openbsd'
  }
  stage( 'plan9' ) {
    sh 'make -f Makefile.gen plan9'
  }
  stage( 'solaris' ) {
    sh 'make -f Makefile.gen solaris'
  }
  stage( 'windows' ) {
    sh 'make -f Makefile.gen windows'
  }
}
