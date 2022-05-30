package idgener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSonyflakeGen_Next(t *testing.T) {
	gen, _ := NewSonyflakeGen()
	id1, err := gen.Next()
	if err != nil {
		assert.FailNow(t, "NewSnowflakeGen gen next get error", err)
		return
	}
	id2, err := gen.Next()
	if err != nil {
		assert.FailNow(t, "NewSnowflakeGen gen next get error", err)
		return
	}
	assert.NotEqual(t, id1, id2)
}
