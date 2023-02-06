package stdutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/stdutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestGo(t *testing.T) {
	err := stdutil.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}
