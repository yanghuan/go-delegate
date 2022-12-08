package delegate

import (
	"sync/atomic"
	"unsafe"
)

type FnX func(args ...interface{})

type Delegate struct {
	multicastDelegate
}

func (d Delegate) Combine(f ...FnX) Delegate {
	m := d.combine(unsafe.Pointer(&f))
	return Delegate{m}
}

func (d Delegate) CombineDelegate(v Delegate) Delegate {
	m := d.combineDelegate(v.multicastDelegate)
	return Delegate{m}
}

func (d Delegate) Remove(f ...FnX) Delegate {
	m := d.remove(unsafe.Pointer(&f))
	return Delegate{m}
}

func (d Delegate) RemoveDelegate(v Delegate) Delegate {
	m := d.removeDelegate(v.multicastDelegate)
	return Delegate{m}
}

func (d Delegate) Invoke(args ...interface{}) {
	for _, invocation := range d.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*FnX)(unsafe.Pointer(&funcVar)))(args...)
	}
}

type multicastDelegate struct {
	invocations []invocation
}

type invocation struct {
	funcPtr, targetPtr uintptr
}

func (i invocation) equals(other invocation) bool {
	return i.funcPtr == other.funcPtr && i.targetPtr == other.targetPtr
}

func (d multicastDelegate) combine(fnPointers unsafe.Pointer) multicastDelegate {
	fns := *(*[]unsafe.Pointer)(fnPointers)
	count := len(fns)
	if count == 0 {
		return d
	}

	if count == 1 {
		fnPointer := fns[0]
		if fnPointer == nil {
			return d
		}

		return d.combineInvocation(getInvocation(fnPointer))
	}

	invocations := make([]invocation, 0, 9)
	for _, fnPointer := range fns {
		if fnPointer == nil {
			continue
		}
		invocations = append(invocations, getInvocation(fnPointer))
	}
	return d.combineDelegate(multicastDelegate{invocations: invocations})
}

func (d multicastDelegate) combineInvocation(a invocation) multicastDelegate {
	return d.combineDelegate(multicastDelegate{invocations: []invocation{a}})
}

func (d multicastDelegate) combineDelegate(follow multicastDelegate) multicastDelegate {
	followLen := len(follow.invocations)
	if followLen == 0 {
		return d
	}

	var invocations []invocation
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
			invocations = make([]invocation, resultLen)
			copy(invocations, d.invocations)
			copy(invocations[length:], follow.invocations)
		}
	} else {
		invocations = append(d.invocations, follow.invocations...)
	}

	return multicastDelegate{invocations: invocations}
}

func (d multicastDelegate) remove(fnPointers unsafe.Pointer) multicastDelegate {
	fns := *(*[]unsafe.Pointer)(fnPointers)
	count := len(fns)
	if count == 0 {
		return d
	}

	if count == 1 {
		fnPointer := fns[0]
		if fnPointer == nil {
			return d
		}

		return d.removeInvocation(getInvocation(fnPointer))
	}

	result := d
	for i := count - 1; i >= 0; i++ {
		fnPointer := fns[i]
		if fnPointer == nil {
			continue
		}
		result = result.removeInvocation(getInvocation(fnPointer))
	}
	return result
}

func (d multicastDelegate) removeInvocation(a invocation) multicastDelegate {
	return d.removeDelegate(multicastDelegate{invocations: []invocation{a}})
}

func (d multicastDelegate) removeDelegate(follow multicastDelegate) multicastDelegate {
	followLen := len(follow.invocations)
	if followLen == 0 {
		return d
	}

	length := len(d.invocations)
	diffLength := length - followLen
	for i := diffLength; i >= 0; i-- {
		if equalsInvocations(d.invocations, follow.invocations, i, followLen) {
			var invocations []invocation
			if i == 0 {
				invocations = d.invocations[followLen:]
			} else if i == diffLength {
				invocations = d.invocations[:diffLength]
			} else {
				invocations := make([]invocation, diffLength)
				copy(invocations, d.invocations[:i])
				copy(invocations[i:], d.invocations[i+followLen:])
			}
			return multicastDelegate{invocations}
		}
	}

	return d
}

func equalsInvocations(a, b []invocation, start, count int) bool {
	for i := 0; i < count; i++ {
		if !a[start+i].equals(b[i]) {
			return false
		}
	}
	return true
}

func trySetSlot(invocations []invocation, index int, value invocation) bool {
	cur := &invocations[index]
	if cur.funcPtr == 0 && atomic.CompareAndSwapUintptr((*uintptr)(unsafe.Pointer(cur)), 0, value.funcPtr) {
		cur.targetPtr = value.targetPtr
		return true
	}

	if cur.funcPtr != 0 && cur.equals(value) {
		return true
	}

	return false
}

func getInvocation(f unsafe.Pointer) invocation {
	return *(*invocation)(f)
}
