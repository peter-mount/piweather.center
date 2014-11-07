#!/bin/bash

# Jenkins build id - default to the date
if [[ "x$BUILD_ID" == "x" ]]
then
    BUILD_ID="$(date)"
fi

# Jenkins build number - manual if not from jenkins
if [[ "x$BUILD_NUMBER" == "x" ]]
then
    BUILD_NUMBER="manual"
    BUILD_VERSION="X"
else
    BUILD_VERSION="$BUILD_NUMBER"
fi

# Jenkins node name - defaults to hostname
if [[ "x$NODE_NAME" == "x" ]]
then
    NODE_NAME=$HOSTNAME
fi

cat piweather_build.h.in | \
    sed "s/@BUILD_ID@/$BUILD_ID/g" | \
    sed "s/@NODE_NAME@/$NODE_NAME/g" | \
    sed "s/@BUILD_NUMBER@/$BUILD_NUMBER/g" | \
    sed "s/@BUILD_VERSION@/$BUILD_VERSION/g" \
    > piweather_build.h
