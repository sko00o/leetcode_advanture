package tmpl3

import (
	"testing"

	"adventure/BinarySearch/bs/test"
)

func TestBinarySearch(t *testing.T) {
	test.CommonTest(t, binarySearch)
}