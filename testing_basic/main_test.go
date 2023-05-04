package testing_basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberDescrptionOne(t *testing.T) {
	var res string
	res = NumberDescrptionOne(-1)
	assert.Equal(t, res, "negative", "It should be negative")

	res = NumberDescrptionOne(5)
	assert.Equal(t, res, "single digit", "It should be single digit")

	res = NumberDescrptionOne(10)
	assert.Equal(t, res, "double digit", "It should be double digit")

	res = NumberDescrptionOne(100)
	assert.Equal(t, res, "double digit", "It should be others")

	res = NumberDescrptionOne(200)
	assert.Equal(t, res, "double digit", "It should be others")
}
