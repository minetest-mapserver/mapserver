package travelnetparser

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	//TODO: test takes about 3 seconds for 350kb data :/
	a, err := ParseFile("testdata/mod_travelnet.data")

	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(a, "", " ")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j[:]))

}
