package cmdline_test

import (
	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/pretty/cliutil/cmdline"
	"github.com/zhangyiming748/pretty/testutil/assert"
	"strings"
	"testing"
)

func TestLineParser_Parse(t *testing.T) {
	args := cmdline.NewParser(`./app top sub -a ddd --xx "msg"`).Parse()
	assert.Len(t, args, 7)
	assert.Eq(t, "msg", args[6])

	args = cmdline.NewParser(`./app top sub -a ddd --xx "abc
def"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "abc\ndef", args[6])

	args = cmdline.NewParser(`./app top sub -a ddd --xx "abc
def ghi"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "abc\ndef ghi", args[6])

	args = cmdline.NewParser(`./app top sub --msg "has multi words"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "has multi words", args[4])

	args = cmdline.NewParser(`./app top sub --msg "has inner 'quote'"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "has inner 'quote'", args[4])

	args = cmdline.NewParser(`./app top sub --msg "'has' inner quote"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "'has' inner quote", args[4])

	args = cmdline.NewParser(`./app top sub --msg "has inner 'quote' words"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "has inner 'quote' words", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "has 'inner quote' words"`)
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "has 'inner quote' words", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "has 'inner quote words'"`)
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "has 'inner quote words'", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "'has inner quote' words"`)
	pretty.P(args)
	assert.Len(t, args, 5)
	assert.Eq(t, "'has inner quote' words", args[4])

	args = cmdline.ParseLine(" ")
	assert.Len(t, args, 0)

	args = cmdline.ParseLine("./app")
	assert.Len(t, args, 1)
}

func TestParseLine_errLine(t *testing.T) {
	// exception line string.
	args := cmdline.NewParser(`./app top sub -a ddd --xx msg"`).Parse()
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg`)
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg text`)
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg text", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg "text"`)
	pretty.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg text", args[6])
}

func TestLineParser_BinAndArgs(t *testing.T) {
	p := cmdline.NewParser("git status")
	b, a := p.BinAndArgs()
	assert.Eq(t, "git", b)
	assert.Eq(t, "status", strings.Join(a, " "))

	p = cmdline.NewParser("git")
	b, a = p.BinAndArgs()
	assert.Eq(t, "git", b)
	assert.Empty(t, a)
}
