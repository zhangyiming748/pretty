package cflag_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhangyiming748/pretty/cflag"
	"github.com/zhangyiming748/pretty/cliutil"
	"github.com/zhangyiming748/pretty/errorx"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func Example() {
	opts := struct {
		age  int
		name string
		str1 string
		bol  bool
	}{}

	c := cflag.New(func(c *cflag.CFlags) {
		c.Desc = "this is a demo command"
		c.Version = "0.5.1"
	})
	c.IntVar(&opts.age, "age", 0, "this is a int option;;a")
	c.StringVar(&opts.name, "name", "", "this is a string option and required;true")
	c.StringVar(&opts.str1, "str1", "def-val", "this is a string option with default value;;s")

	c.AddArg("arg1", "this is arg1", true, nil)
	c.AddArg("arg2", "this is arg2", true, nil)
	c.AddArg("arg3", "this is arg3 with default", false, "def-val")

	c.Func = func(c *cflag.CFlags) error {
		// do something ...

		cliutil.Infoln("hello, this is", c.Name())
		cliutil.Infoln("option.age =", opts.age)
		cliutil.Infoln("option.name =", opts.name)
		cliutil.Infoln("option.str1 =", opts.str1)
		cliutil.Infoln("arg1 =", c.Arg("arg1").String())
		cliutil.Infoln("arg2 =", c.Arg("arg2").String())
		cliutil.Infoln("arg3 =", c.Arg("arg3").String())

		return nil
	}

	// c.MustParse(os.Args[1:])
	c.MustParse(nil)
}

func TestSetDebug(t *testing.T) {
	cflag.SetDebug(true)
	assert.True(t, cflag.Debug)
	cflag.SetDebug(false)
}

var opts = struct {
	int  int
	str  string
	str1 string
	bol  bool
}{}

func TestNew(t *testing.T) {
	c := cflag.New(
		cflag.WithDesc("desc for the console command"),
		cflag.WithVersion("1.0.2"),
	)
	c.IntVar(&opts.int, "int", 0, "this is a int option;true;i")
	c.StringVar(&opts.str, "str", "", "this is a string option;;s")
	c.StringVar(&opts.str1, "str1", "def-val", "this is a string option with default;;s1")
	c.AddValidator("int", func(val any) error {
		iv := val.(int)
		if iv < 10 {
			return errorx.Raw("value should >= 10")
		}
		return nil
	})
	c.LongHelp = "this is a long help\nthis is a long help\nthis is a long help"
	c.Example = "this is some example for {{cmd}}\nthis is some example for {{cmd}}\nthis is some example for {{cmd}}"

	c.AddArg("ag1", "this is a int option", false, nil)
	c.AddArg("arg3", "this is arg2 with default", false, "def-val")

	inArgs := []string{"--help"}
	err := c.Parse(inArgs)
	assert.NoErr(t, err)

	inArgs = []string{"--int", "23"}
	err = c.Parse(inArgs)
	assert.NoErr(t, err)
	assert.Eq(t, 23, opts.int)

	// use validate
	inArgs = []string{"--int", "3"}
	err = c.Parse(inArgs)
	assert.Err(t, err)
	assert.Eq(t, "flag option 'int': value should >= 10", err.Error())
}

func TestCFlags_Parse(t *testing.T) {
	var opts = struct {
		int  int
		str  string
		str1 string
		bol  bool
	}{}

	c := cflag.New(func(c *cflag.CFlags) {
		c.Desc = "this is a demo command"
		c.Version = "0.5.1"
	})
	c.IntVar(&opts.int, "int", 0, "this is a int option;false;i")

	assert.PanicsMsg(t, func() {
		c.AddShortcuts("notExist", "d,e")
	}, "cflag: option 'notExist' is not registered")

	// assert.PanicsMsg(t, func() {
	// 	c.AddShortcuts("int", "i,n")
	// }, "cflag: option 'notExist' is not registered")

	osArgs := os.Args
	os.Args = []string{"./myapp", "ag1", "ag2"}

	c.QuickRun()
	assert.Eq(t, "[ag1 ag2]", fmt.Sprint(c.RemainArgs()))

	os.Args = osArgs
}
