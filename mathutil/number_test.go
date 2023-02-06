package mathutil_test

import (
	"testing"
	"time"

	"github.com/zhangyiming748/pretty/mathutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestIsNumeric(t *testing.T) {
	assert.True(t, mathutil.IsNumeric('3'))
	assert.False(t, mathutil.IsNumeric('a'))
}

func TestPercent(t *testing.T) {
	assert.Eq(t, float64(34), mathutil.Percent(34, 100))
	assert.Eq(t, float64(0), mathutil.Percent(34, 0))
	assert.Eq(t, float64(-100), mathutil.Percent(34, -34))
}

func TestElapsedTime(t *testing.T) {
	nt := time.Now().Add(-time.Second * 3)
	num := mathutil.ElapsedTime(nt)

	assert.Eq(t, 3000, int(mathutil.MustFloat(num)))
}

func TestDataSize(t *testing.T) {
	assert.Eq(t, "3.38K", mathutil.DataSize(3456))
}

func TestHowLongAgo(t *testing.T) {
	assert.Eq(t, "57 mins", mathutil.HowLongAgo(3456))
}
