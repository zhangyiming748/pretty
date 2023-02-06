package panics_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/testutil/assert"
	"github.com/zhangyiming748/pretty/testutil/panics"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}
