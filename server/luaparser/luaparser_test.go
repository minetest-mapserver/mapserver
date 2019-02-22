package luaparser

import (
	"fmt"
	"testing"
)

func TestParseList(t *testing.T) {
	p := New()
	m, err := p.ParseList(`return {{["x"]=1},{["y"]=2}}`)

	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	if len(m) != 2 {
		t.Fatalf("wrong length: %d", len(m))
	}

	v1 := m[0]
	fmt.Println(v1)

	if v1["x"].(int) != 1 {
		t.Fatal("[0][x] does not match")
	}

	v2 := m[1]
	if v2["y"].(int) != 2 {
		t.Fatal("[1][y] does not match")
	}

}

func TestParseMap(t *testing.T) {
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
