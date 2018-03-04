package gotesttools

import (
	"testing"
)

func TestAssertEqual(t *testing.T) {
  AssertEqual(t, true, true)
}

func TestAssertTrue(t *testing.T) {
  AssertTrue(t, 1 == 1)
}

func TestAssertPanic(t *testing.T) {
  assert := AssertPanic(t)
  defer assert()
  panic("Expected panic")
}
