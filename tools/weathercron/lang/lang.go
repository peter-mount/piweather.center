package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/location"
)

type Script struct {
	Pos      lexer.Position
	Location *location.Location
	Rules    []*Rule `parser:"@@*"`
	state    *State
}

func (s *Script) merge(b *Script) (*Script, error) {
	if s == nil {
		return b, nil
	}

	s.Rules = append(s.Rules, b.Rules...)
	s.state.merge(b.state)

	return s, nil
}

type Rule struct {
	Pos  lexer.Position
	Task *TaskRule `parser:"@@"`
}

type TaskRule struct {
	Pos       lexer.Position
	Schedules []*Schedule `parser:"'task' (@@+)"`
	Task      *Task       `parser:"@@"`
}

type Schedule struct {
	Pos   lexer.Position
	At    *At    `parser:"( @@"`
	Every *Every `parser:"| @@"`
	Cron  *Cron  `parser:"| @@)"`
}

type Cron struct {
	Pos        lexer.Position
	Definition string `parser:"'cron' @String"`
}

type At struct {
	Pos        lexer.Position
	Definition string `parser:"'at' @String"`
}

type Every struct {
	Pos        lexer.Position
	Definition string `parser:"'every' @String"`
	Cron       string // Parsed version of Definition
}

type Task struct {
	Pos lexer.Position
	Run string `parser:"'run' @String"`
}
