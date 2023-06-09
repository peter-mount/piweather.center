## Name
weatherbot

## Synopsis

    weatherbot [-host https://example.com] [-rootDir /config/path] [-v] [-test] -post id

## Description

**weatherbot** provides a tool that will post weather related statuses to Mastodon.
It's intended to be run from a crontab.

The <u>-host</u> option sets the http endpoint the bot will use to communicate
with the weather station.
If this is not set then it will default `http://127.0.0.1:8080`

The <u>-rootDir</u> option specifies the directory containing the bot's config.
If this is not set then it will default to the `etc` directory of the weather stations distribution.

The <u>-v</u> option will show debug information on the console.

The <u>-test</u> option puts the bot into test mode.
In this mode it will write what it would post to stdout rather than to Mastodon.

The <u>-post</u> option selects which post to submit.
You can define multiple posts depending on what you need.

For example, an hourly post which gives the current weather once an hour,
or a daily post which includes the statistics for the previous day. 

## Configuration

There's two configuration files required by weatherbot which need to be in the
directory pointed to by <u>-rootDir</u>.

### mastodon.yaml

    # Mastodon
    #debug: true
    server: "https://example.com"
    access_token: "server-provided-token"

The two main entries are:
* server - the hostname of the Mastodon instance to connect to
* access_token - an access token created on the Mastodon server for the account
you will be posting as. You can create a token by going to `Settings -> Development` in Mastodon
and selecting `New Application`.

### weatherbot.yaml

This is the main configuration file, which the bot uses to process posts.

This will be documented separately.

### crontab

The bot needs to be run at specific times of the day to make the posts.
The easiest method is with a traditional crontab entry.

The following is an example used during development:

    # Run @me15weather@area51.social every hour
    0 * * * * /usr/local/bin/weatherbot -rootDir /home/peter -v -post hourly
