package lang

import (
	"github.com/peter-mount/go-script/parser"
	"sync"
)

type State struct {
	mutex      sync.Mutex
	scriptInit parser.Initialiser // from go-script, initialiser for embedded scripts
}

func NewState() *State {
	return &State{
		scriptInit: parser.NewInitialiser(),
	}
}

func (s *State) Cleanup() *State {
	s.scriptInit = nil
	return s
}

func (s *State) merge(b *State) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
}
