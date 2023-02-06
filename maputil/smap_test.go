package maputil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/maputil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestSMap_usage(t *testing.T) {
	mp := maputil.SMap{
		"k1": "23",
		"k2": "ab",
		"k3": "true",
		"k4": "1,2",
	}

	assert.True(t, mp.Has("k1"))
	assert.True(t, mp.HasValue("true"))
	assert.True(t, mp.Bool("k3"))
	assert.False(t, mp.IsEmpty())
	assert.False(t, mp.HasValue("not-exist"))
	assert.Len(t, mp.Keys(), 4)
	assert.Len(t, mp.Values(), 4)

	val, ok := mp.Value("k2")
	assert.True(t, ok)
	assert.Eq(t, "ab", val)

	// int
	assert.Eq(t, 23, mp.Int("k1"))
	assert.Eq(t, int64(23), mp.Int64("k1"))

	// str
	assert.Eq(t, "23", mp.Str("k1"))
	assert.Eq(t, "ab", mp.Get("k2"))

	// slice
	assert.Eq(t, []int{1, 2}, mp.Ints("k4"))
	assert.Eq(t, []string{"1", "2"}, mp.Strings("k4"))
	assert.Nil(t, mp.Strings("not-exist"))

	// not exists
	assert.False(t, mp.Bool("notExists"))
	assert.Eq(t, 0, mp.Int("notExists"))
	assert.Eq(t, int64(0), mp.Int64("notExists"))
	assert.Eq(t, "", mp.Str("notExists"))
	assert.Empty(t, mp.Ints("notExists"))
}
