package delegate

import (
	"sync/atomic"
	"unsafe"
)

type FnX func(args ...interface{})

type Delegate struct {
	multicastDelegate
}

func (d Delegate) Combine(a FnX) Delegate {
	if a == nil {
		return d
	}

	m := d.combine(getInvocation(a))
	return Delegate{m}
}

func (d Delegate) CombineDelegate(follow Delegate) Delegate {
	m := d.combineDelegate(follow.multicastDelegate)
	return Delegate{m}
}

func (d Delegate) Remove(a FnX) Delegate {
	if a == nil {
		return d
	}

	m := d.remove(getInvocation(a))
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

func (i invocation) Equals(other invocation) bool {
	return i.funcPtr == other.funcPtr && i.targetPtr == other.targetPtr
}

func (d multicastDelegate) combine(a invocation) multicastDelegate {
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

func (d multicastDelegate) remove(a invocation) multicastDelegate {
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
		if !a[start+i].Equals(b[i]) {
			return false
		}
	}
	return true
}

func trySetSlot(invocations []invocation, index int, invocation invocation) bool {
	if invocations[index].funcPtr == 0 && atomic.CompareAndSwapUintptr((*uintptr)(unsafe.Pointer(&invocations[index])), 0, invocation.funcPtr) {
		invocations[index].targetPtr = invocation.targetPtr
		return true
	}

	cur := invocations[index]
	if cur.funcPtr != 0 && cur.Equals(invocation) {
		return true
	}

	return false
}

func getInvocation(f interface{}) invocation {
	type funcHeader struct {
		funcPtr   uintptr
		targetPtr uintptr
	}
	type interfaceHeader struct {
		typ  uintptr
		data *funcHeader
	}
	address := (*interfaceHeader)(unsafe.Pointer(&f)).data
	return invocation{address.funcPtr, address.targetPtr}
}
