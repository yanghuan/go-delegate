package delegate

import "unsafe"

type Fn func()

type Action struct {
	multicastDelegate
}

func (a Action) Combine(f Fn) Action {
	if f == nil {
		return a
	}

	m := a.combine(getInvocation(f))
	return Action{m}
}

func (a Action) Remove(f Fn) Action {
	if f == nil {
		return a
	}

	m := a.remove(getInvocation(f))
	return Action{m}
}

func (a Action) Invoke() {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn)(unsafe.Pointer(&funcVar)))()
	}
}

type Fn1[T any] func(T)

type Action1[T any] struct {
	multicastDelegate
}

func (a Action1[T]) Combine(f Fn1[T]) Action1[T] {
	if f == nil {
		return a
	}

	m := a.combine(getInvocation(f))
	return Action1[T]{m}
}

func (a Action1[T]) Remove(f Fn1[T]) Action1[T] {
	if f == nil {
		return a
	}

	m := a.remove(getInvocation(f))
	return Action1[T]{m}
}

func (a Action1[T]) Invoke(x T) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn1[T])(unsafe.Pointer(&funcVar)))(x)
	}
}
