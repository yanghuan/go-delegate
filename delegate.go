package delegate

import (
	"sync/atomic"
	"unsafe"
)

type FnX func(args ...interface{})

type Delegate struct {
	multicastDelegate[FnX]
}

func (d Delegate) Equals(other Delegate) bool {
	return d.multicastDelegate.equals(other.multicastDelegate)
}

func (d Delegate) Combine(f ...FnX) Delegate {
	m := d.combine(f)
	return Delegate{m}
}

func (d Delegate) CombineDelegate(v Delegate) Delegate {
	m := d.combineDelegate(v.multicastDelegate)
	return Delegate{m}
}

func (d Delegate) Remove(f ...FnX) Delegate {
	m := d.remove(f)
	return Delegate{m}
}

func (d Delegate) RemoveDelegate(v Delegate) Delegate {
	m := d.removeDelegate(v.multicastDelegate)
	return Delegate{m}
}

func (d Delegate) GetInvocationList() []FnX {
	return d.invocations
}

func (d Delegate) Invoke(args ...interface{}) {
	for _, invocation := range d.invocations {
		invocation(args...)
	}
}

type multicastDelegate[F any] struct {
	invocations []F
}

type invocation struct {
	funcPtr, targetPtr uintptr
}

func (i invocation) equals(other invocation) bool {
	return i.funcPtr == other.funcPtr && i.targetPtr == other.targetPtr
}

func toPointer[F any](f F) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&f))
}

func toInvocation[F any](f F) invocation {
	return *(*invocation)(toPointer(f))
}

func equals[F any](a, b F) bool {
	return toInvocation(a).equals(toInvocation(b))
}

func isNil[F any](f F) bool {
	return toPointer(f) == nil
}

func (d multicastDelegate[F]) equals(other multicastDelegate[F]) bool {
	count, otherCount := len(d.invocations), len(other.invocations)
	if count != otherCount {
		return false
	}

	for i := 0; i < count; i++ {
		if !equals(d.invocations[i], other.invocations[i]) {
			return false
		}
	}

	return true
}

func (d multicastDelegate[F]) combine(fns []F) multicastDelegate[F] {
	count := len(fns)
	if count == 0 {
		return d
	}

	if count == 1 {
		f := fns[0]
		if isNil(f) {
			return d
		}

		return d.combineInvocation(f)
	}

	invocations := make([]F, 0, 9)
	for _, f := range fns {
		if isNil(f) {
			continue
		}
		invocations = append(invocations, f)
	}
	return d.combineDelegate(multicastDelegate[F]{invocations: invocations})
}

func (d multicastDelegate[F]) combineInvocation(f F) multicastDelegate[F] {
	return d.combineDelegate(multicastDelegate[F]{invocations: []F{f}})
}

func (d multicastDelegate[F]) combineDelegate(follow multicastDelegate[F]) multicastDelegate[F] {
	followLen := len(follow.invocations)
	if followLen == 0 {
		return d
	}

	var invocations []F
	length := len(d.invocations)
	resultLen := length + followLen
	if resultLen <= cap(d.invocations) {
		invocations = d.invocations[:resultLen]
		for i := 0; i < followLen; i++ {
			if !trySetSlot(invocations, length+i, follow.invocations[i]) {
				invocations = nil
				break
			}
		}

		if invocations == nil {
			invocations = make([]F, resultLen)
			copy(invocations, d.invocations)
			copy(invocations[length:], follow.invocations)
		}
	} else {
		invocations = append(d.invocations, follow.invocations...)
	}

	return multicastDelegate[F]{invocations: invocations}
}

func (d multicastDelegate[F]) remove(fns []F) multicastDelegate[F] {
	count := len(fns)
	if count == 0 {
		return d
	}

	if count == 1 {
		f := fns[0]
		if isNil(f) {
			return d
		}

		return d.removeInvocation(f)
	}

	result := d
	for i := count - 1; i >= 0; i++ {
		f := fns[i]
		if isNil(f) {
			continue
		}
		result = result.removeInvocation(f)
	}
	return result
}

func (d multicastDelegate[F]) removeInvocation(f F) multicastDelegate[F] {
	return d.removeDelegate(multicastDelegate[F]{invocations: []F{f}})
}

func (d multicastDelegate[F]) removeDelegate(follow multicastDelegate[F]) multicastDelegate[F] {
	followLen := len(follow.invocations)
	if followLen == 0 {
		return d
	}

	length := len(d.invocations)
	diffLength := length - followLen
	for i := diffLength; i >= 0; i-- {
		if equalsInvocations(d.invocations, follow.invocations, i, followLen) {
			var invocations []F
			if i == 0 {
				invocations = d.invocations[followLen:]
			} else if i == diffLength {
				invocations = d.invocations[:diffLength]
			} else {
				invocations := make([]F, diffLength)
				copy(invocations, d.invocations[:i])
				copy(invocations[i:], d.invocations[i+followLen:])
			}
			return multicastDelegate[F]{invocations}
		}
	}

	return d
}

func equalsInvocations[F any](a, b []F, start, count int) bool {
	for i := 0; i < count; i++ {
		if !equals(a[start+i], b[i]) {
			return false
		}
	}
	return true
}

func trySetSlot[F any](invocations []F, index int, value F) bool {
	cur := &invocations[index]
	if isNil(*cur) {
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(cur)), nil, toPointer(value)) {
			*cur = value
			return true
		}
	} else if equals(*cur, value) {
		return true
	}
	return false
}
