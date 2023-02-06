package sysutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/sysutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestCallersInfo(t *testing.T) {
	cs := sysutil.CallersInfos(0, 2)
	// pretty.P(cs)
	assert.NotEmpty(t, cs)
	assert.Len(t, cs, 2)
	assert.StrContains(t, cs[0].String(), "goutil/sysutil/stack.go")
}
