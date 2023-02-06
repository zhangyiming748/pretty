package sysutil_test

import (
	"github.com/zhangyiming748/pretty"
	"testing"

	"github.com/zhangyiming748/pretty/sysutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, sysutil.GoVersion())

	info, err := sysutil.ParseGoVersion("go version go1.19.2 darwin/amd64")
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	assert.Eq(t, "1.19.2", info.Version)
	assert.Eq(t, "darwin", info.GoOS)
	assert.Eq(t, "amd64", info.Arch)

	info, err = sysutil.OsGoInfo()
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	pretty.P(info)
}
