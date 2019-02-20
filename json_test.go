package json

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type A struct {
	A int    `json:"a"`
	B string `json:"b"`
}

var a A

func Test(t *testing.T) {
	j := `{"a":1,"b":"a"}`
	Decode(strings.NewReader(j), &a)
	fmt.Println(a)
	Encode(os.Stdout, &a)
}
