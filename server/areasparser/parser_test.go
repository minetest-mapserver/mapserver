package areasparser

import (
  "testing"
  "fmt"
  "encoding/json"

)

func TestParse(t *testing.T){

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
