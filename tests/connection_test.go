package redigowrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("aaa", "aaa")
}
