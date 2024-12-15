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

type commandParser struct {
	whiteSpaceToken lexer.TokenType
}

// Parser handles the parsing of the Command.
//
// To register this Parser with participle, use the following to create a parser option:
//
// participle.ParseTypeWith[command.Command](command.Parser(l))
//
// where l is the lexer.Definition
func Parser(l lexer.Definition) func(pl *lexer.PeekingLexer) (Command, error) {
	symbols := l.Symbols()

	parser := &commandParser{
		whiteSpaceToken: symbols["Whitespace"],
	}

	return parser.parseCommand
}

func (c *commandParser) parseCommand(pl *lexer.PeekingLexer) (Command, error) {
	var ret command

	s, stop := c.parseAttribute(pl)
	ret.command = s

	for !stop {
		s, stop = c.parseAttribute(pl)
		ret.args = append(ret.args, s)
	}

	return ret, nil
}

func (c *commandParser) parseAttribute(pl *lexer.PeekingLexer) (string, bool) {
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

		if token.Type == c.whiteSpaceToken {
			// Stop parsing the attribute on a whitespace,
			// but terminate the command on newline, carriage return or form feed.
			stop = strings.ContainsAny(token.Value, "\n\r\f")
			break
		} else {
			s = append(s, token.String())
			_ = pl.Next()
		}
	}

	return strings.Join(s, ""), stop
}
