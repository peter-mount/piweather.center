
all: $(BUILD)/videogenerator

$(BUILD)/videogenerator: bin/*.go */*.go *.go
	$(GO_BUILD) -o $@ $<

clean:
