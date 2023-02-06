package structs_test

import (
	"testing"

	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/pretty/structs"
	"github.com/zhangyiming748/pretty/testutil/assert"
)

func TestTryToMap(t *testing.T) {
	mp, err := structs.TryToMap(nil)
	assert.Empty(t, mp)
	assert.NoErr(t, err)

	type User struct {
		Name string
		Age  int
		city string
	}

	u := User{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp, err = structs.TryToMap(u)
	assert.NoErr(t, err)
	pretty.P(mp)
	assert.Contains(t, mp, "Name")
	assert.Contains(t, mp, "Age")
	assert.NotContains(t, mp, "city")

	mp, err = structs.TryToMap(&u)
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)
	// pretty.P(mp)

	mp = structs.MustToMap(&u)
	assert.NotEmpty(t, mp)
	// pretty.P(mp)

	assert.Panics(t, func() {
		structs.MustToMap("abc")
	})

}

func TestToMap_useTag(t *testing.T) {
	type User1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		city string
	}

	u1 := &User1{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp := structs.ToMap(u1)
	pretty.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age"})
	assert.NotContains(t, mp, "city")
}

func TestToMap_nestStruct(t *testing.T) {
	type Extra struct {
		City   string `json:"city"`
		Github string `json:"github"`
	}
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Extra Extra  `json:"extra"`
	}

	u := &User{
		Name: "inhere",
		Age:  30,
		Extra: Extra{
			City:   "chengdu",
			Github: "https://github.com/inhere",
		},
	}

	mp := structs.MustToMap(u)
	pretty.P(mp)
	assert.ContainsKeys(t, mp, []string{"name", "age", "extra"})
	assert.ContainsKeys(t, mp["extra"], []string{"city", "github"})
}

func TestTryToMap_customTag(t *testing.T) {
	type User struct {
		Name     string `export:"name"`
		Age      int    `export:"age"`
		FullName string `export:"full_name"`
	}

	u1 := User{
		Name:     "inhere",
		Age:      34,
		FullName: "inhere xyz",
	}

	mp, err := structs.TryToMap(u1, func(opt *structs.MapOptions) {
		opt.TagName = "export"
	})
	assert.NoErr(t, err)
	assert.NotEmpty(t, mp)

	assert.ContainsKeys(t, mp, []string{"name", "age", "full_name"})
}
