package pretty

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Available struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	LongCode string `json:"longCode"`
}

func TestPrintDump(t *testing.T) {
	var as []Available
	b := Read0()
	err := json.Unmarshal(b, &as)
	if err != nil {
		return
	}
	P(as)
}
func Read0() []byte {
	f, err := os.ReadFile("example.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return f
}
