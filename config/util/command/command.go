package command

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

// Command represents a parsed command line, consisting of the executable path and any arguments.
// The command line is terminated by a new line, carriage return, form feed or eof
type Command interface {
	Command() string
	Args() []string
}

type command struct {
	command string
	args    []string
}

// Command returns the command name
func (c command) Command() string {
	return c.command
}

// Args returns the arguments for the command
func (c command) Args() []string {
	return c.args
}

// Parser handles the parsing of the Command.
//
// To register this Parser with participle, use the following to create a parser option:
//
// participle.ParseTypeWith[command.Command](command.Parser)
func Parser(pl *lexer.PeekingLexer) (Command, error) {
	var ret command

	s, stop := commandParser2(pl)
	ret.command = s

	for !stop {
		s, stop = commandParser2(pl)
		ret.args = append(ret.args, s)
	}

	return ret, nil
}

func commandParser2(pl *lexer.PeekingLexer) (string, bool) {
	var s []string
	var stop bool

	for {
		// Peek at the token. For the first one use Peek as we are expecting a token
		// but after that use RawPeek as we want to check for elided tokens, specifically whitespace
		var token *lexer.Token
		if len(s) == 0 {
			token = pl.Peek()
		} else {
			token = pl.RawPeek()
		}

		if token.EOF() {
			stop = true
			break
		}

		switch token.Value[0] {
		// new line, carriage return or form feed terminates the command
		case '\n', '\r', '\f':
			return strings.Join(s, ""), true

		// space, tab, U+0085 (NEL), U+00A0 (NBSP) ends the argument
		case '\t', '\v', ' ', 0x85, 0xA0:
			return strings.Join(s, ""), false

		default:
			s = append(s, token.String())
			_ = pl.Next()
		}
	}

	return strings.Join(s, ""), stop
}
