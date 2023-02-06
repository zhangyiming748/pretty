package finder_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/fsutil/finder"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestFilterFunc(t *testing.T) {
	fn := finder.FilterFunc(func(filePath, filename string) bool {
		return false
	})

	assert.False(t, fn("path/some.txt", "some.txt"))
}
