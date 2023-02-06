package strutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/strutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestBuffer_WriteAny(t *testing.T) {
	buf := strutil.NewBuffer()

	buf.QuietWritef("ab-%s", "c")
	buf.QuietWriteByte('d')
	assert.Eq(t, "ab-cd", buf.ResetAndGet())

	buf.QuietWriteString("ab", "-", "cd")
	buf.MustWriteString("-ef")
	assert.Eq(t, "ab-cd-ef", buf.ResetAndGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetAndGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())
}
