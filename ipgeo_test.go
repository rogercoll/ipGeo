package ipgeo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeUrl(t *testing.T) {
	actual_result, _ := makeUrl("8.8.8.8", "juju")
	assert.Equal(t, "http://api.ipstack.com/8.8.8.8?access_key=juju", actual_result.URL.String())
}
