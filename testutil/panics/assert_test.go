package panics_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/panics"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}
