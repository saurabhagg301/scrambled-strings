package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedValue    = flag.Int("expectedValue", 4, "expected value/result for the input test file")
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 1, output[0].LineNum)
	assert.Equal(t, *expectedValue, output[0].Count)
}
