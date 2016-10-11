package time

import (
	"testing"
	tt "time"

	"github.com/stretchr/testify/assert"
)

func TestUTC(t *testing.T) {

	assert := assert.New(t)

	var ts = UTC

	assert.InDelta(tt.Now().UTC().Unix(), ts().Unix(), 1)

	ts = func() tt.Time {
		return tt.Unix(0, 0)
	}

	assert.Equal(1970, ts().Year())

}
