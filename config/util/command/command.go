package command

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

// Command represents a parsed command line, consisting of the executable path and any arguments.
// The command line is terminated by a new line, carriage return, form feed or eof
type Command interface {
	Position() lexer.Position
	Command() string
	Args() []string
}

type command struct {
	position lexer.Position
	command  string
	args     []string
}

// Command returns the command name
func (c command) Command() string {
	return c.command
}

// Args returns the arguments for the command
func (c command) Args() []string {
	return c.args
}

func (c command) Position() lexer.Position {
	return c.position
}

type commandParser struct {
	whiteSpaceToken lexer.TokenType
	punctToken      lexer.TokenType
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
		punctToken:      symbols["Punct"],
	}

	return parser.parseCommand
}

func (c *commandParser) parseCommand(pl *lexer.PeekingLexer) (Command, error) {
	s, pos, stop := c.parseAttribute(pl)
	ret := command{position: pos, command: s}

	for !stop {
		s, _, stop = c.parseAttribute(pl)
		ret.args = append(ret.args, s)
	}

	return ret, nil
}

func (c *commandParser) parseAttribute(pl *lexer.PeekingLexer) (string, lexer.Position, bool) {
	var s []string
	var stop bool

	var token *lexer.Token
	for {
		// Peek at the token. For the first one use Peek as we are expecting a token
		// but after that use RawPeek as we want to check for elided tokens, specifically whitespace
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
			stop = isLineTerminator(token.Value)
			break
		} else if token.Type == c.punctToken && token.Value == `\` {
			// Is the \ at the end of the line? If not then include \ but either way carry on parsing
			// NB: Next then RawPeek as we need to look for whitespace which is Elided
			_ = pl.Next()
			next := pl.RawPeek()

			if !isLineTerminator(next.Value) {
				s = append(s, token.String())
			}
		} else {
			s = append(s, token.String())
			_ = pl.Next()
		}
	}

	return strings.Join(s, ""), token.Pos, stop
}

func isLineTerminator(s string) bool {
	return strings.ContainsAny(s, "\n\r\f")
}
