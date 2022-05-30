package idgener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowflakeGen_Next(t *testing.T) {
	gen, err := NewSnowflakeGen()
	if err != nil {
		assert.FailNow(t, "NewSnowflakeGen get error", err)
		return
	}
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
