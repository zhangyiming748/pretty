package cmdr

import (
	"context"
	"fmt"
	"github.com/zhangyiming748/pretty"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/zhangyiming748/pretty/arrutil"
	"github.com/zhangyiming748/pretty/color"
	"github.com/zhangyiming748/pretty/internal/comfunc"
)

// Cmd struct
type Cmd struct {
	*exec.Cmd
	// Name of the command
	Name string
	// inited bool

	// BeforeRun hook
	BeforeRun func(c *Cmd)
	// AfterRun hook
	AfterRun func(c *Cmd, err error)
}

// WrapGoCmd instance
func WrapGoCmd(cmd *exec.Cmd) *Cmd {
	return &Cmd{Cmd: cmd}
}

// NewGitCmd instance
func NewGitCmd(subCmd string, args ...string) *Cmd {
	return NewCmd("git", subCmd).AddArgs(args)
}

// NewCmd instance
//
// see exec.Command
func NewCmd(bin string, args ...string) *Cmd {
	return &Cmd{
		Cmd: exec.Command(bin, args...),
	}
}

// CmdWithCtx create new instance with context.
//
// see exec.CommandContext
func CmdWithCtx(ctx context.Context, bin string, args ...string) *Cmd {
	return &Cmd{
		Cmd: exec.CommandContext(ctx, bin, args...),
	}
}

// PrintCmdline on before exec
func PrintCmdline(c *Cmd) {
	color.Yellowln(">", c.Cmdline())
}

// -------------------------------------------------
// config the command
// -------------------------------------------------

// Config the command
func (c *Cmd) Config(fn func(c *Cmd)) *Cmd {
	fn(c)
	return c
}

// OnBefore exec add hook
func (c *Cmd) OnBefore(fn func(c *Cmd)) *Cmd {
	c.BeforeRun = fn
	return c
}

// OnAfter exec add hook
func (c *Cmd) OnAfter(fn func(c *Cmd, err error)) *Cmd {
	c.AfterRun = fn
	return c
}

// WithBin name returns the current object
func (c *Cmd) WithBin(name string) *Cmd {
	c.Args[0] = name
	c.lookPath(name)
	return c
}

func (c *Cmd) lookPath(name string) {
	if filepath.Base(name) == name {
		lp, err := exec.LookPath(name)
		if lp != "" {
			// Update cmd.Path even if err is non-nil.
			// If err is ErrDot (especially on Windows), lp may include a resolved
			// extension (like .exe or .bat) that should be preserved.
			c.Path = lp
		}
		if err != nil {
			pretty.Panicf("look %q path error: %s", name, err.Error())
		}
	}
}

// WithGoCmd and returns the current instance.
func (c *Cmd) WithGoCmd(ec *exec.Cmd) *Cmd {
	c.Cmd = ec
	return c
}

// WithWorkDir returns the current object
func (c *Cmd) WithWorkDir(dir string) *Cmd {
	c.Dir = dir
	return c
}

// WorkDirOnNot set, returns the current object
func (c *Cmd) WorkDirOnNot(dir string) *Cmd {
	if c.Dir == "" {
		c.Dir = dir
	}
	return c
}

// OutputToOS output to OS stdout and error
func (c *Cmd) OutputToOS() *Cmd {
	return c.ToOSStdoutStderr()
}

// ToOSStdoutStderr output to OS stdout and error
func (c *Cmd) ToOSStdoutStderr() *Cmd {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

// ToOSStdout output to OS stdout
func (c *Cmd) ToOSStdout() *Cmd {
	c.Stdout = os.Stdout
	c.Stderr = os.Stdout
	return c
}

// WithStdin returns the current argument
func (c *Cmd) WithStdin(in io.Reader) *Cmd {
	c.Stdin = in
	return c
}

// WithOutput returns the current instance
func (c *Cmd) WithOutput(out, errOut io.Writer) *Cmd {
	c.Stdout = out
	if errOut != nil {
		c.Stderr = errOut
	}
	return c
}

// WithAnyArgs add args and returns the current object.
func (c *Cmd) WithAnyArgs(args ...any) *Cmd {
	c.Args = append(c.Args, arrutil.SliceToStrings(args)...)
	return c
}

// AddArg add args and returns the current object
func (c *Cmd) AddArg(args ...string) *Cmd { return c.WithArg(args...) }

// WithArg add args and returns the current object. alias of the WithArg()
func (c *Cmd) WithArg(args ...string) *Cmd {
	c.Args = append(c.Args, args...)
	return c
}

// AddArgf add args and returns the current object. alias of the WithArgf()
func (c *Cmd) AddArgf(format string, args ...any) *Cmd {
	return c.WithArgf(format, args...)
}

// WithArgf add arg and returns the current object
func (c *Cmd) WithArgf(format string, args ...any) *Cmd {
	c.Args = append(c.Args, fmt.Sprintf(format, args...))
	return c
}

// ArgIf add arg and returns the current object
func (c *Cmd) ArgIf(arg string, exprOk bool) *Cmd {
	if exprOk {
		c.Args = append(c.Args, arg)
	}
	return c
}

// WithArgIf add arg and returns the current object
func (c *Cmd) WithArgIf(arg string, exprOk bool) *Cmd {
	return c.ArgIf(arg, exprOk)
}

// AddArgs for the git. alias of WithArgs()
func (c *Cmd) AddArgs(args []string) *Cmd { return c.WithArgs(args) }

// WithArgs for the git
func (c *Cmd) WithArgs(args []string) *Cmd {
	if len(args) > 0 {
		c.Args = append(c.Args, args...)
	}
	return c
}

// WithArgsIf add arg and returns the current object
func (c *Cmd) WithArgsIf(args []string, exprOk bool) *Cmd {
	if exprOk && len(args) > 0 {
		c.Args = append(c.Args, args...)
	}
	return c
}

// -------------------------------------------------
// helper command
// -------------------------------------------------

// IDString of the command
func (c *Cmd) IDString() string {
	if c.Name != "" {
		return c.Name
	}
	return c.BinOrPath()
}

// BinName of the command
func (c *Cmd) BinName() string {
	if len(c.Args) > 0 {
		return c.Args[0]
	}
	return ""
}

// BinOrPath of the command
func (c *Cmd) BinOrPath() string {
	if len(c.Args) > 0 {
		return c.Args[0]
	}
	return c.Path
}

// OnlyArgs of the command, not contains bin name.
func (c *Cmd) OnlyArgs() (ss []string) {
	if len(c.Args) > 1 {
		return c.Args[1:]
	}
	return
}

// ResetArgs for command, but will keep bin name.
func (c *Cmd) ResetArgs() {
	if len(c.Args) > 0 {
		c.Args = c.Args[0:1]
	} else {
		c.Args = c.Args[:0]
	}
}

// Cmdline to command line
func (c *Cmd) Cmdline() string {
	return comfunc.Cmdline(c.Args)
}

// Copy new instance from current command, with new args.
func (c *Cmd) Copy(args ...string) *Cmd {
	nc := *c

	// copy bin name.
	if len(c.Args) > 0 {
		nc.Args = append([]string{c.Args[0]}, args...)
	} else {
		nc.Args = args
	}

	return &nc
}

// GoCmd get exec.Cmd
func (c *Cmd) GoCmd() *exec.Cmd { return c.Cmd }

// -------------------------------------------------
// run command
// -------------------------------------------------

// Success run and return whether success
func (c *Cmd) Success() bool {
	return c.Run() == nil
}

// HasStdout output setting.
func (c *Cmd) HasStdout() bool {
	return c.Stdout != nil
}

// SafeLines run and return output as lines
func (c *Cmd) SafeLines() []string {
	ss, _ := c.OutputLines()
	return ss
}

// OutputLines run and return output as lines
func (c *Cmd) OutputLines() ([]string, error) {
	out, err := c.Output()
	if err != nil {
		return nil, err
	}
	return OutputLines(out), err
}

// SafeOutput run and return output
func (c *Cmd) SafeOutput() string {
	out, err := c.Output()
	if err != nil {
		return ""
	}
	return out
}

// Output run and return output
func (c *Cmd) Output() (string, error) {
	if c.BeforeRun != nil {
		c.BeforeRun(c)
	}

	output, err := c.Cmd.Output()

	if c.AfterRun != nil {
		c.AfterRun(c, err)
	}
	return string(output), err
}

// CombinedOutput run and return output, will combine stderr and stdout output
func (c *Cmd) CombinedOutput() (string, error) {
	if c.BeforeRun != nil {
		c.BeforeRun(c)
	}

	output, err := c.Cmd.CombinedOutput()

	if c.AfterRun != nil {
		c.AfterRun(c, err)
	}
	return string(output), err
}

// MustRun a command. will panic on error
func (c *Cmd) MustRun() {
	if err := c.Run(); err != nil {
		panic(err)
	}
}

// FlushRun runs command and flush output to stdout
func (c *Cmd) FlushRun() error {
	c.ToOSStdoutStderr()
	return c.Run()
}

// Run runs command
func (c *Cmd) Run() error {
	if c.BeforeRun != nil {
		c.BeforeRun(c)
	}

	// do running
	err := c.Cmd.Run()

	if c.AfterRun != nil {
		c.AfterRun(c, err)
	}
	return err

	// if IsWindows() {
	// 	return c.Spawn()
	// }
	// return c.Exec()
}

// Spawn runs command with spawn(3)
// func (c *Cmd) Spawn() error {
// 	return c.Cmd.Run()
// }
//
// // Exec runs command with exec(3)
// // Note that Windows doesn't support exec(3): http://golang.org/src/pkg/syscall/exec_windows.go#L339
// func (c *Cmd) Exec() error {
// 	binary, err := exec.LookPath(c.Path)
// 	if err != nil {
// 		return &exec.Error{
// 			Name: c.Path,
// 			Err:  errorx.Newf("%s not found in the system", c.Path),
// 		}
// 	}
//
// 	args := []string{binary}
// 	args = append(args, c.Args...)
//
// 	return syscall.Exec(binary, args, os.Environ())
// }
