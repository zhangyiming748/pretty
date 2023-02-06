package sysutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/dump"
	"github.com/zhangyiming748/pretty/sysutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestUserDir(t *testing.T) {
	dir := sysutil.UserHomeDir()
	assert.NotEmpty(t, dir)
	dump.P(dir)

	dir1 := sysutil.HomeDir()
	assert.NotEmpty(t, dir1)

	dir2 := sysutil.UHomeDir()
	assert.NotEmpty(t, dir2)
	assert.Eq(t, dir1, dir2)

	dir = sysutil.UserDir("sub-path")
	assert.Contains(t, dir, "/sub-path")
	dump.P(dir)

	dir = sysutil.UserCacheDir("my-logs")
	assert.Contains(t, dir, ".cache/my-logs")
	dump.P(dir)

	dir = sysutil.UserConfigDir("my-conf")
	assert.Contains(t, dir, ".config/my-conf")
	dump.P(dir)

	rawPath := "~/.kite"
	assert.LenGt(t, sysutil.ExpandPath(rawPath), len(rawPath))
}

func TestWorkdir(t *testing.T) {
	assert.NotEmpty(t, sysutil.Workdir())
}

func TestLoginUser(t *testing.T) {
	cu := sysutil.LoginUser()
	assert.NotEmpty(t, cu)

	fu := sysutil.MustFindUser(cu.Username)
	assert.NotEmpty(t, fu)
	assert.Eq(t, cu.Uid, fu.Uid)
}
