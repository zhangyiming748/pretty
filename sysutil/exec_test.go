package sysutil_test

import (
	"strings"
	"testing"

	"github.com/zhangyiming748/pretty/sysutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestExecCmd(t *testing.T) {
	ret, err := sysutil.ExecCmd("echo", []string{"OK"})
	assert.NoErr(t, err)
	// *nix: "OK\n" win: "OK\r\n"
	assert.Eq(t, "OK", strings.TrimSpace(ret))

	ret, err = sysutil.QuickExec("echo OK")
	assert.NoErr(t, err)
	assert.Eq(t, "OK", strings.TrimSpace(ret))

	ret, err = sysutil.ExecLine("echo OK1")
	assert.NoErr(t, err)
	assert.Eq(t, "OK1", strings.TrimSpace(ret))
}

func TestShellExec(t *testing.T) {
	ret, err := sysutil.ShellExec("echo OK")
	assert.NoErr(t, err)
	// *nix: "OK\n" win: "OK\r\n"
	assert.Eq(t, "OK", strings.TrimSpace(ret))

	ret, err = sysutil.ShellExec("echo OK", "bash")
	assert.NoErr(t, err)
	assert.Eq(t, "OK", strings.TrimSpace(ret))
}
