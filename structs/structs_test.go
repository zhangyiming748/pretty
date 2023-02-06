package structs_test

import (
	"testing"

	"github.com/zhangyiming748/pretty/structs"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestIsExported(t *testing.T) {
	assert.True(t, structs.IsExported("Name"))
	assert.True(t, structs.IsExported("Abc12"))
	assert.True(t, structs.IsExported("A"))
	assert.False(t, structs.IsExported("name"))
	assert.False(t, structs.IsExported("_name"))
	assert.False(t, structs.IsExported("abc12"))
	assert.False(t, structs.IsExported("123abcd"))

	assert.False(t, structs.IsUnexported("Name"))
	assert.False(t, structs.IsUnexported("Abc12"))
	assert.True(t, structs.IsUnexported("name"))
	assert.True(t, structs.IsUnexported("_name"))
	assert.True(t, structs.IsUnexported("abc12"))
	assert.True(t, structs.IsUnexported("123abcd"))
}
