package luaparser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := New()
	m, err := p.ParseMap(`return {a=1, b=true, c="abc"}`)

	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	if m["a"].(int) != 1 {
		t.Fatal("parsing error")
	}

	if !m["b"].(bool) {
		t.Fatal("parsing error")
	}

	if m["c"].(string) != "abc" {
		t.Fatal("parsing error")
	}

}
