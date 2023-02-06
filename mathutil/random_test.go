package mathutil_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangyiming748/pretty/mathutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestRandomInt(t *testing.T) {
	min, max := 1000, 9999

	for i := 0; i < 5; i++ {
		val := mathutil.RandomInt(min, max)
		fmt.Println(val)
		assert.True(t, val >= min)
		assert.True(t, val <= max)

		seed := time.Now().UnixNano()
		val = mathutil.RandomIntWithSeed(min, max, seed)
		assert.True(t, val >= min)
	}

	assert.True(t, mathutil.RandInt(min, max) > 999)
}
