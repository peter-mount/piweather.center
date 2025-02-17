module github.com/peter-mount/piweather.center

go 1.24

toolchain go1.24.0

require (
	github.com/alecthomas/participle/v2 v2.1.1
	github.com/eclipse/paho.mqtt.golang v1.5.0
	github.com/fsnotify/fsnotify v1.8.0
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3
	github.com/llgcode/draw2d v0.0.0-20240627062922-0ed1ff131195
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/peter-mount/go-anim v0.0.0-20250218200716-17e2f5f48f5c
	github.com/peter-mount/go-build v0.0.0-20250218200125-f187f75a6a5d
	github.com/peter-mount/go-kernel/v2 v2.0.3-0.20250218195942-5604474bedd7
	github.com/peter-mount/go-mastodon v0.0.0-20221228215100-3fcdfd9b124a
	github.com/peter-mount/go-script v0.0.0-20250218200359-943ffa62e818
	github.com/peter-mount/nre-feeds v0.0.0-20240201140817-fd78167946e5
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/soniakeys/meeus/v3 v3.0.1
	github.com/soniakeys/unit v1.0.0
	go.bug.st/serial v1.6.2
	golang.org/x/image v0.24.0
	golang.org/x/text v0.22.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
	gopkg.in/yaml.v3 v3.0.1
)

//replace github.com/peter-mount/go-script v0.0.0-20241218090358-129a6c764bf4 => ../script

//replace github.com/peter-mount/go-anim v0.0.0-20241218114807-a958c4339040 => ../go-anim

require (
	github.com/creack/goselect v0.1.2 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/peter-mount/go.uuid v1.2.1-0.20180103174451-36e9d2ebbde5 // indirect
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
