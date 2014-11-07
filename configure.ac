# autoconf script for building the pi weather station
#
AC_INIT([piweather], [0.1], [peter@retep.org])
AC_MSG_NOTICE([Pi Weather Station])

# Generate global_config.h.in if it's not present or older than configure
AC_CONFIG_HEADERS([global_config.h])
if test ! -f global_config.h.in -o global_config.h.in -ot configure
then
    AC_MSG_NOTICE([Generating global_config.h.in])
    autoheader
fi

# Now run make to generate piweather_build.h to ensure it exists
# otherwise IDE's will complain. It will always be rebuilt by
# make as required.
if test ! -f piweather_build.h.in -o piweather_build.h.in -ot configure
then
    AC_MSG_NOTICE([Generating piweather_build.h])
    bash ./piweather_build.sh
fi

# Check for gcc
AC_PROG_CC
AC_LANG(C)

# Math library
#AC_CHECK_LIB(m, cos)

# ======================================================================
DIRS=""

# Project base for include paths
LDFLAGS="$LDFLAGS -L../build"

# ======================================================================
# Camera support
# ======================================================================
AH_TEMPLATE(HAVE_CAMERA, [Camera support])
camera=0
AC_ARG_WITH([sqlite],
    AS_HELP_STRING([--without-camera], [disable camera support]))
AS_IF([test "x$with_camera" != "xno"],
    [
        camera=1
    ],[
        camera=0
    ])
AS_IF([test $camera == 1],
    [
        AC_MSG_NOTICE([Enabling camera support])
        AC_DEFINE(HAVE_CAMERA)
        LDFLAGS="$LDFLAGS -lsqlite3"
    ],[
        AC_MSG_NOTICE([Disabled camera support])
    ])

# ======================================================================
# libmicrohttpd needed for our webserver
# ======================================================================
AH_TEMPLATE(HAVE_MICROHTTPD, [microhttpd embedded webserver])
microhttpd=1
AC_CHECK_HEADER(microhttpd.h, [], [microhttpd=0])
AC_CHECK_LIB(microhttpd, MHD_start_daemon, [], [microhttpd=0])
if test $microhttpd == 1
then
    AC_DEFINE(HAVE_MICROHTTPD)
    LDFLAGS="$LDFLAGS -lmicrohttpd"
fi

# ======================================================================
# libsqlite optional for logging data
# ======================================================================
AH_TEMPLATE(HAVE_SQLITE, [sqlite embedded sql database])
AC_ARG_WITH([sqlite],
    AS_HELP_STRING([--without-sqlite], [disable sqlite support]))
AS_IF([test "x$with_sqlite" != "xno"],
    [
        sqlite=1
        AC_CHECK_HEADER(sqlite3.h, [], [sqlite=0])
        AC_CHECK_LIB(sqlite3, sqlite3_close, [], [sqlite=0])
    ],[
        sqlite=0
    ])
AS_IF([test $sqlite == 1],
    [
        AC_MSG_NOTICE([Enabling sqlite support])
        AC_DEFINE(HAVE_SQLITE)
        LDFLAGS="$LDFLAGS -lsqlite3"
    ],[
        AC_MSG_NOTICE([Disabled sqlite support])
    ])

# ======================================================================
# libcurl3 optional for logging data to 3rd party websites
# ======================================================================
AH_TEMPLATE(HAVE_CURL, [Curl HTTP client])
AC_ARG_WITH([curl],
    AS_HELP_STRING([--without-curl], [disable Curl HTTP support]))
AS_IF([test "x$with_curl" != "xno"],
    [
        curl=1
        AC_CHECK_HEADER(curl/curl.h, [], [curl=0])
        AC_CHECK_LIB(curl, curl_easy_cleanup, [], [curl=0])
    ],[
        curl=0
    ])
AS_IF([test $curl == 1],
    [
        AC_MSG_NOTICE([Enabling Curl HTTP support])
        AC_DEFINE(HAVE_CURL)
        LDFLAGS="$LDFLAGS -lcurl"
    ],[
        AC_MSG_NOTICE([Disabled Curl HTTP support])
    ])
AC_SUBST([CURL],[${curl}])

# ======================================================================
# librabbitmq optional for logging data to a RabbitMQ server
#
# Note, the DIR in use here is rabbitmq-logger and not rabbitmq as thats
# the external libraries name
# ======================================================================
AH_TEMPLATE(HAVE_RABBITMQ, [RabbitMQ messaging server])
AC_ARG_WITH([rabbitmq],
    AS_HELP_STRING([--without-rabbitmq], [disable RabbitMQ support]))
AS_IF([test "x$with_rabbitmq" != "xno"],
    [
        rabbitmq=1
        AC_CHECK_HEADER(amqp.h, [], [rabbitmq=0])
        AC_CHECK_LIB(rabbitmq, amqp_new_connection, [], [rabbitmq=0])
    ],[
        rabbitmq=0
    ])
AS_IF([test $rabbitmq == 1],
    [
        AC_MSG_NOTICE([Enabling RabbitMQ support])
        AC_DEFINE(HAVE_RABBITMQ)
        LDFLAGS="$LDFLAGS -lrabbitmq"
    ],[
        AC_MSG_NOTICE([Disabled RabbitMQ support])
    ])
AC_SUBST([RABBITMQ],[${rabbitmq}])

# ======================================================================
# i2c sensors
#
# To enable ensure that i2c_bcm2708 kernel module is running or
# use --with-i2c as it's safe to build
# ======================================================================
AH_TEMPLATE(HAVE_I2C, [i2c based sensors])
AC_ARG_WITH([i2c],
    AS_HELP_STRING([--with-i2c],[force i2c support even if not enabled on this machine]))
AS_IF([test "x$with_i2c" != "xno"],
    [
        i2c=1
    ],[
        i2c=$(lsmod | grep -c i2c_bcm2708)
        AC_CHECK_FILE([/dev/i2c-0],[],[i2c=0])
        AC_CHECK_FILE([/dev/i2c-1],[],[i2c=0])
    ])
AS_IF([test $i2c == 1],
    [
        AC_MSG_NOTICE([Enabling I2C sensors])
        AC_DEFINE(HAVE_I2C)
    ],[
        AC_MSG_NOTICE([Disabled I2C sensors])
    ])
AC_SUBST([I2C],[${i2c}])

# ======================================================================
# wiringPi needed for GPIO
# ======================================================================
AH_TEMPLATE(HAVE_WIRING_PI, [wiringPi GPIO])
AC_ARG_WITH([wiringpi],
    AS_HELP_STRING([--without-wiringpi], [disable WiringPi support]))
AS_IF([test "x$with_wiringpi" != "xno"],
    [
        wiringpi=1
        AC_CHECK_HEADER(wiringPi.h, [], [wiringpi=0])
        AC_CHECK_LIB(wiringPi, wiringPiSetup, [], [wiringpi=0])
    ],[
        wiringpi=0
    ])
AS_IF([test $wiringpi == 1],
    [
        AC_MSG_NOTICE([Enabling wiringPi based sensors])
        AC_DEFINE(HAVE_WIRING_PI)
        LDFLAGS="$LDFLAGS -lwiringPi"
    ],[
        AC_MSG_NOTICE([Disabled wiringPi based sensors])
    ])
AC_SUBST([WIRING_PI],[${wiringpi}])

# ======================================================================
# gd required for annotating images
# ======================================================================
AH_TEMPLATE(HAVE_LIB_GD, [GD for annotating images])
AC_CHECK_LIB([gd], [gdFreeFontCache], [gd=1], [gd=0])
if test $gd == 1
then
    AC_DEFINE(HAVE_LIB_GD)
    LDFLAGS="$LDFLAGS -lgd"
fi

# ======================================================================
# Work out what to compile
# ======================================================================

# Webserver support
AS_IF([test $microhttpd == 1], [DIRS="$DIRS webserver"])

# Camera support
AS_IF([test $camera == 1],
    [
        DIRS="$DIRS renderers camera"
    ])

# Image Filter library
DIRS="$DIRS imagefilter"

# RabbitMQ
AS_IF([test $rabbitmq == 1],[DIRS="$DIRS rabbitmqapi"])

# logger's
DIRS="$DIRS logger"

# sensors
DIRS="$DIRS clouds sensors"

# ======================================================================
# Finish off generating all files

# add libraries based on the directory names, in reverse order to aid linking
for d in $DIRS
do
    LDFLAGS="-l$d $LDFLAGS"
done

# Finally add lib & weatherstation, doesn't matter compilation order but
# needs to be linked in a specific order to allow linking
DIRS="lib astro scheduler $DIRS weatherstation"
LDFLAGS="-lweatherstation $LDFLAGS -lscheduler -lastro -lweather"

AC_SUBST([DIRECTORIES],[$DIRS])
AC_SUBST([CPPFLAGS],[$CPPFLAGS])
AC_SUBST([LDFLAGS],[$LDFLAGS])

MAKEFILES="SubMakefile $(find . -type f -name Makefile.in|sed s/\\.in//g)"
AC_CONFIG_FILES([$MAKEFILES])

AC_OUTPUT()

AC_MSG_NOTICE([Now run make clean;make])
