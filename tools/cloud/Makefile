
all: $(BUILD)/clouddetect

$(BUILD)/clouddetect: bin/*.go */*.go *.go
	$(GO_BUILD) -o $@ $<

clean:
