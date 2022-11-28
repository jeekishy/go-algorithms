package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompressString(t *testing.T) {
	assert.Equal(t, "a2b2c3", compressString([]byte("aabbccc")))
	assert.Equal(t, "a", compressString([]byte("a")))
	assert.Equal(t, "ab3", compressString([]byte("abbb")))
	assert.Equal(t, "ab12", compressString([]byte("abbbbbbbbbbbb")))
}
