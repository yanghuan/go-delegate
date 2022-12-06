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

	m := d.combineDelegate(multicastDelegate{invocations: []invocation{getInvocation(a)}})
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
	lastIndex := len(d.invocations) - 1
	i := lastIndex
	for ; i >= 0; i-- {
		if funcPtrEqual(d.invocations[i], a) {
			break
		}
	}

	if i == -1 {
		return d
	}

	if i == lastIndex {
		return multicastDelegate{invocations: d.invocations[:lastIndex]}
	}

	if i == 0 {
		return multicastDelegate{invocations: d.invocations[1:]}
	}

	actions := make([]invocation, lastIndex)
	copy(actions, d.invocations[:i])
	copy(actions[i:], d.invocations[i+1:])
	return multicastDelegate{invocations: actions}
}

func trySetSlot(invocations []invocation, index int, invocation invocation) bool {
	if invocations[index].funcPtr == 0 && atomic.CompareAndSwapUintptr((*uintptr)(unsafe.Pointer(&invocations[index])), 0, invocation.funcPtr) {
		invocations[index].targetPtr = invocation.targetPtr
		return true
	}

	cur := invocations[index]
	if cur.funcPtr != 0 && funcPtrEqual(cur, invocation) {
		return true
	}

	return false
}

func funcPtrEqual(f1, f2 invocation) bool {
	return f1.funcPtr == f2.funcPtr && f1.targetPtr == f2.targetPtr
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
