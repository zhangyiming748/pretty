package strutil_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/zhangyiming748/pretty/strutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestRandomChars(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomChars(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomCharsV2(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV2(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomCharsV3(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := strutil.RandomCharsV3(4)
		fmt.Println(str)

		assert.Len(t, str, 4)
	}
}

func TestRandomBytes(t *testing.T) {
	b, err := strutil.RandomBytes(3)

	// 1607400451937462000
	tsn := time.Now().UnixNano()
	rand.Seed(tsn)

	fmt.Println(tsn)
	fmt.Println(rand.Intn(12))
	fmt.Println(rand.Intn(12))

	fmt.Println(string(b))
	fmt.Println(base64.URLEncoding.EncodeToString(b))
	fmt.Println(base64.StdEncoding.EncodeToString(b))
	fmt.Println(hex.EncodeToString(b))
	assert.NoErr(t, err)
}

func TestRandomString(t *testing.T) {
	s, err := strutil.RandomString(3)

	fmt.Println(s)
	assert.NoErr(t, err)
	assert.True(t, len(s) > 3)
}
