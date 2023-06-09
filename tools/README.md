# tools

This directory contains the sources for each individual command line tool.

## cloud

Standalone utility that will calculate the cloud coverage from a single image

## ephemeris

Standalone utility that generates an ephemeris for a specific date or range of dates.
This ephemeris contains the rise/set times of the Sun, Moon and major planets for
a specific location

## videoident

This extends my videoident project to generate static images or time-lapse videos
which contain captured images of the sky, calculated cloud cover with overlays of
captured data.

## dataencoder

This is a tool used by the build process to generate static data files required by the
various tools.

## weatherbot

This is a standalone tool which implements a bot to post weather reports to Mastodon.
An example is [@me15weather@area51.social](https://area51.social/@me15weather) which
posts the current weather every hour.

## weatherstation

This is the core service which both captures data from sensors but also makes them
available to other services.