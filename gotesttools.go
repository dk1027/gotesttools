package gotesttools

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expected %+v. Got: %+v", expected, actual))
	}
}

func AssertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Error(fmt.Sprintf("Expected True. Got False."))
	}
}

type assertPanicT struct {
	panicked bool
	t *testing.T
}

func check(this *assertPanicT) func(){
	return func() {
		if r := recover(); r != nil {
			this.panicked = true
		}
		AssertTrue(this.t, this.panicked)
	}
}

func AssertPanic(t *testing.T) func(){
	assert := &assertPanicT{t: t, panicked: false}
	return check(assert)
}
