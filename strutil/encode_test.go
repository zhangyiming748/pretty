package strutil_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/strutil"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Eq(t, tests.want, strutil.EscapeHTML(tests.give))

	ret := strutil.EscapeJS("<script>var a = 23;</script>")
	assert.NotContains(t, ret, "<script>")
	assert.NotContains(t, ret, "</script>")
}

func TestAddSlashes(t *testing.T) {
	assert.Eq(t, "", strutil.AddSlashes(""))
	assert.Eq(t, "", strutil.StripSlashes(""))

	assert.Eq(t, `{\"key\": 123}`, strutil.AddSlashes(`{"key": 123}`))
	assert.Eq(t, `{"key": 123}`, strutil.StripSlashes(`{\"key\": 123}`))
	assert.Eq(t, `path\to`, strutil.StripSlashes(`path\\to`))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Eq("a.com/?name%3D%E4%BD%A0%E5%A5%BD", strutil.URLEncode("a.com/?name=你好"))
	is.Eq("a.com/?name=你好", strutil.URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Eq("a.com", strutil.URLEncode("a.com"))
	is.Eq("a.com", strutil.URLDecode("a.com"))
}

func TestBaseDecode(t *testing.T) {
	is := assert.New(t)

	is.Eq("MFRGG===", strutil.B32Encode("abc"))
	is.Eq("abc", strutil.B32Decode("MFRGG==="))

	is.Eq("YWJj", strutil.B64Encode("abc"))
	is.Eq("abc", strutil.B64Decode("YWJj"))
	is.Eq([]byte("YWJj"), strutil.B64EncodeBytes([]byte("abc")))
}
