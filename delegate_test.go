package delegate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testDelegateHelper struct {
	s *string
}

func (h *testDelegateHelper) f1(args ...interface{}) {
	*h.s = *h.s + "a"
}

func (h *testDelegateHelper) f2(args ...interface{}) {
	*h.s = *h.s + "b"
}

func (h *testDelegateHelper) f3(args ...interface{}) {
	*h.s = *h.s + "c"
}

func TestDelegateCombine(t *testing.T) {
	Convey("TestCombineFunc", t, func() {
		var s string
		f1 := func(args ...interface{}) {
			s += "a"
		}

		f2 := func(args ...interface{}) {
			s += "b"
		}

		f3 := func(args ...interface{}) {
			s += "c"
		}

		d := Delegate{}.CombineAction(f1).CombineAction(f2).CombineAction(f3)
		d.Invoke()
		So(s, ShouldEqual, "abc")
	})

	Convey("TestCombineMethod", t, func() {
		var s string
		h := &testDelegateHelper{s: &s}

		d := Delegate{}.CombineAction(h.f1).CombineAction(h.f2).CombineAction(h.f3)
		d.Invoke()
		So(s, ShouldEqual, "abc")
	})
}

func TestDelegateRemove(t *testing.T) {
	Convey("TestCombineRemoveFunc", t, func() {
		var s string
		f1 := func(args ...interface{}) {
			s += "a"
		}

		f2 := func(args ...interface{}) {
			s += "b"
		}

		f3 := func(args ...interface{}) {
			s += "c"
		}

		d := Delegate{}.CombineAction(f1)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineDelegate(Delegate{}.CombineAction(f2).CombineAction(f3))
		d.Invoke()
		So(s, ShouldEqual, "abc")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2).CombineDelegate(Delegate{}.CombineAction(f2).CombineAction(f3))
		d.Invoke()
		So(s, ShouldEqual, "abbc")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2).RemoveAction(f1)
		d.Invoke()
		So(s, ShouldEqual, "b")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2).RemoveAction(f2)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2).CombineAction(f1).RemoveAction(f1)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.CombineAction(f1).CombineAction(f2).CombineAction(f3).RemoveAction(f1).RemoveAction(f2)
		d.Invoke()
		So(s, ShouldEqual, "c")
		s = ""
	})

	Convey("TestCombineRemoveMethod", t, func() {
		var s string
		h := &testDelegateHelper{s: &s}

		d := Delegate{}.CombineAction(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineDelegate(Delegate{}.CombineAction(h.f2).CombineAction(h.f3))
		d.Invoke()
		So(s, ShouldEqual, "abc")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2).CombineDelegate(Delegate{}.CombineAction(h.f2).CombineAction(h.f3))
		d.Invoke()
		So(s, ShouldEqual, "abbc")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2).RemoveAction(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "b")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2).RemoveAction(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2).CombineAction(h.f1).RemoveAction(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.CombineAction(h.f1).CombineAction(h.f2).CombineAction(h.f3).RemoveAction(h.f1).RemoveAction(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "c")
		s = ""
	})
}
