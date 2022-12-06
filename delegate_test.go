package delegate

import (
	"testing"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

type testDelegateHelper struct {
	s *string
}

func (h *testDelegateHelper) f1(_ ...interface{}) {
	*h.s = *h.s + "a"
}

func (h *testDelegateHelper) f2(_ ...interface{}) {
	*h.s = *h.s + "b"
}

func (h *testDelegateHelper) f3(_ ...interface{}) {
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

		d := Delegate{}.Combine(f1).Combine(f2).Combine(f3)
		d.Invoke()
		So(s, ShouldEqual, "abc")
	})

	Convey("TestCombineMethod", t, func() {
		var s string
		h := &testDelegateHelper{s: &s}

		d := Delegate{}.Combine(h.f1).Combine(h.f2).Combine(h.f3)
		d.Invoke()
		So(s, ShouldEqual, "abc")
	})

	Convey("TestCombineSliceConflict", t, func() {
		var s string
		h := &testDelegateHelper{s: &s}

		d := Delegate{}
		println(uintptr(unsafe.Pointer(&d)))

		d = d.Combine(h.f1).Combine(h.f2).Combine(h.f3)
		e := d.Combine(h.f1)
		d = d.Combine(h.f2)
		d.Invoke()
		e.Invoke()
		So(s, ShouldEqual, "abcbabca")
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

		d := Delegate{}.Combine(f1)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.Combine(f1).CombineDelegate(Delegate{}.Combine(f2).Combine(f3))
		d.Invoke()
		So(s, ShouldEqual, "abc")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2).CombineDelegate(Delegate{}.Combine(f2).Combine(f3))
		d.Invoke()
		So(s, ShouldEqual, "abbc")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2).Remove(f1)
		d.Invoke()
		So(s, ShouldEqual, "b")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2).Remove(f2)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2).Combine(f1).Remove(f1)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.Combine(f1).Combine(f2).Combine(f3).Remove(f1).Remove(f2)
		d.Invoke()
		So(s, ShouldEqual, "c")
		s = ""
	})

	Convey("TestCombineRemoveMethod", t, func() {
		var s string
		h := &testDelegateHelper{s: &s}

		d := Delegate{}.Combine(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.Combine(h.f1).CombineDelegate(Delegate{}.Combine(h.f2).Combine(h.f3))
		d.Invoke()
		So(s, ShouldEqual, "abc")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2).CombineDelegate(Delegate{}.Combine(h.f2).Combine(h.f3))
		d.Invoke()
		So(s, ShouldEqual, "abbc")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2).Remove(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "b")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2).Remove(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "a")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2).Combine(h.f1).Remove(h.f1)
		d.Invoke()
		So(s, ShouldEqual, "ab")
		s = ""

		d = Delegate{}.Combine(h.f1).Combine(h.f2).Combine(h.f3).Remove(h.f1).Remove(h.f2)
		d.Invoke()
		So(s, ShouldEqual, "c")
		s = ""
	})
}
