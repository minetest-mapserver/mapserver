package areasparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	a, err := ParseFile("testdata/areas.json")
	assert.NoError(t, err)
	assert.True(t, len(a) > 1)

	area := a[0]
	assert.Equal(t, "ilai_house", area.Name)
	assert.Equal(t, "ilai", area.Owner)
	assert.NotNil(t, area.Pos1)
	assert.NotNil(t, area.Pos2)
	assert.Equal(t, 4970, area.Pos1.X)
}
