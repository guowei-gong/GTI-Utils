package xreflect_test

import (
	"github.com/guowei-gong/GTI-Utils/xreflect"
	"testing"
)

func TestIsNil(t *testing.T) {
	var b1 bool
	var b2 *bool
	t.Log(xreflect.IsNil(b1))
	t.Log(xreflect.IsNil(&b1))
	t.Log(xreflect.IsNil(b2))
	t.Log(xreflect.IsNil(&b2))
}
