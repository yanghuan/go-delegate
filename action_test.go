package delegate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestActionCombine(t *testing.T) {
	Convey("TestCombineFunc", t, func() {
		var s string
		f1 := func() {
			s += "a"
		}

		f2 := func() {
			s += "b"
		}

		f3 := func() {
			s += "c"
		}

		d := Action{}.Combine(f1).Combine(f2).Combine(f3)
		d.Invoke()
		So(s, ShouldEqual, "abc")
	})
}

func TestAction1Combine(t *testing.T) {
	Convey("TestCombineFunc", t, func() {
		var s string
		f1 := func(c string) {
			s += c + "a"
		}

		f2 := func(c string) {
			s += c + "b"
		}

		f3 := func(c string) {
			s += c + "c"
		}

		d := Action1[string]{}.Combine(f1).Combine(f2).Combine(f3)
		d.Invoke("T")
		So(s, ShouldEqual, "TaTbTc")
	})
}
