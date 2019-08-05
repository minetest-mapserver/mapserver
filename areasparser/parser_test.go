package areasparser

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	a, err := ParseFile("testdata/areas.dat")

	if err != nil {
		t.Fatal(err)
	}

	j, err := json.MarshalIndent(a, "", " ")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(j[:]))

}
