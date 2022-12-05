package delegate

import "unsafe"

type ActionX func(args ...interface{})

type Delegate struct {
	actions []ActionX
}

func NewDelegate(actions ...ActionX) Delegate {
	length := len(actions)
	if length == 0 {
		return Delegate{actions: actions}
	}

	d := Delegate{actions: actions[:1]}
	for i := 1; i < length; i++ {
		d = d.CombineAction(actions[i])
	}
	return d
}

func (d Delegate) CombineAction(a ActionX) Delegate {
	if a == nil {
		return d
	}

	return d.CombineDelegate(Delegate{actions: []ActionX{a}})
}

func (d Delegate) CombineDelegate(follow Delegate) Delegate {
	followLen := len(follow.actions)
	if followLen == 0 {
		return d
	}

	var actions []ActionX
	length := len(d.actions)
	resultLen := length + followLen
	if resultLen <= cap(d.actions) {
		actions = d.actions[:resultLen]
		for i := 0; i < followLen; i++ {
			curIndex := length + i
			cur, a := actions[curIndex], follow.actions[i]
			if cur == nil {
				actions[curIndex] = a
			} else if !funcEqual(cur, a) {
				actions = nil
				break
			}
		}

		if actions == nil {
			actions = make([]ActionX, resultLen)
			copy(actions, d.actions)
			copy(actions[length:], follow.actions)
		}
	} else {
		actions = append(d.actions, follow.actions...)
	}

	return Delegate{actions: actions}
}

func (d Delegate) RemoveAction(a ActionX) Delegate {
	if a == nil {
		return d
	}

	lastIndex := len(d.actions) - 1
	i := lastIndex
	for ; i >= 0; i-- {
		if funcEqual(d.actions[i], a) {
			break
		}
	}

	if i == -1 {
		return d
	}

	if i == lastIndex {
		return Delegate{actions: d.actions[:lastIndex]}
	}

	if i == 0 {
		return Delegate{actions: d.actions[1:]}
	}

	actions := make([]ActionX, lastIndex)
	copy(actions, d.actions[:i])
	copy(actions[i:], d.actions[i+1:])
	return Delegate{actions: actions}
}

func (d Delegate) Invoke(args ...interface{}) {
	for _, a := range d.actions {
		a(args...)
	}
}

func getFuncPtrAndTargetPtr(f interface{}) (uintptr, uintptr) {
	type funcHeader struct {
		funcPtr   uintptr
		targetPtr uintptr
	}
	type interfaceHeader struct {
		typ  uintptr
		data *funcHeader
	}
	address := (*interfaceHeader)(unsafe.Pointer(&f)).data
	return address.funcPtr, address.targetPtr
}

func funcEqual(a interface{}, b interface{}) bool {
	x1, x2 := getFuncPtrAndTargetPtr(a)
	y1, y2 := getFuncPtrAndTargetPtr(b)
	return x1 == y1 && x2 == y2
}
