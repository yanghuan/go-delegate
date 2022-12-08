package delegate

import "unsafe"

type Fn func()

type Action struct {
	multicastDelegate
}

func (a Action) Equals(other Action) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action) Combine(f ...Fn) Action {
	m := a.combine(unsafe.Pointer(&f))
	return Action{m}
}

func (a Action) CombineDelegate(v Action) Action {
	m := a.combineDelegate(v.multicastDelegate)
	return Action{m}
}

func (a Action) Remove(f ...Fn) Action {
	m := a.remove(unsafe.Pointer(&f))
	return Action{m}
}

func (a Action) RemoveDelegate(v Action) Action {
	m := a.removeDelegate(v.multicastDelegate)
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

func (a Action1[T]) Equals(other Action1[T]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action1[T]) Combine(f ...Fn1[T]) Action1[T] {
	m := a.combine(unsafe.Pointer(&f))
	return Action1[T]{m}
}

func (a Action1[T]) CombineDelegate(v Action1[T]) Action1[T] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action1[T]{m}
}

func (a Action1[T]) Remove(f ...Fn1[T]) Action1[T] {
	m := a.remove(unsafe.Pointer(&f))
	return Action1[T]{m}
}

func (a Action1[T]) RemoveDelegate(v Action1[T]) Action1[T] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action1[T]{m}
}

func (a Action1[T]) Invoke(x T) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn1[T])(unsafe.Pointer(&funcVar)))(x)
	}
}

type Fn2[T1, T2 any] func(T1, T2)

type Action2[T1, T2 any] struct {
	multicastDelegate
}

func (a Action2[T1, T2]) Equals(other Action2[T1, T2]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action2[T1, T2]) Combine(f ...Fn2[T1, T2]) Action2[T1, T2] {
	m := a.combine(unsafe.Pointer(&f))
	return Action2[T1, T2]{m}
}

func (a Action2[T1, T2]) CombineDelegate(v Action2[T1, T2]) Action2[T1, T2] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action2[T1, T2]{m}
}

func (a Action2[T1, T2]) Remove(f ...Fn2[T1, T2]) Action2[T1, T2] {
	m := a.remove(unsafe.Pointer(&f))
	return Action2[T1, T2]{m}
}

func (a Action2[T1, T2]) RemoveDelegate(v Action2[T1, T2]) Action2[T1, T2] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action2[T1, T2]{m}
}

func (a Action2[T1, T2]) Invoke(x1 T1, x2 T2) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn2[T1, T2])(unsafe.Pointer(&funcVar)))(x1, x2)
	}
}

type Fn3[T1, T2, T3 any] func(T1, T2, T3)

type Action3[T1, T2, T3 any] struct {
	multicastDelegate
}

func (a Action3[T1, T2, T3]) Equals(other Action3[T1, T2, T3]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action3[T1, T2, T3]) Combine(f ...Fn3[T1, T2, T3]) Action3[T1, T2, T3] {
	m := a.combine(unsafe.Pointer(&f))
	return Action3[T1, T2, T3]{m}
}

func (a Action3[T1, T2, T3]) CombineDelegate(v Action3[T1, T2, T3]) Action3[T1, T2, T3] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action3[T1, T2, T3]{m}
}

func (a Action3[T1, T2, T3]) Remove(f ...Fn3[T1, T2, T3]) Action3[T1, T2, T3] {
	m := a.remove(unsafe.Pointer(&f))
	return Action3[T1, T2, T3]{m}
}

func (a Action3[T1, T2, T3]) RemoveDelegate(v Action3[T1, T2, T3]) Action3[T1, T2, T3] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action3[T1, T2, T3]{m}
}

func (a Action3[T1, T2, T3]) Invoke(x1 T1, x2 T2, x3 T3) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn3[T1, T2, T3])(unsafe.Pointer(&funcVar)))(x1, x2, x3)
	}
}

type Fn4[T1, T2, T3, T4 any] func(T1, T2, T3, T4)

type Action4[T1, T2, T3, T4 any] struct {
	multicastDelegate
}

func (a Action4[T1, T2, T3, T4]) Equals(other Action4[T1, T2, T3, T4]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action4[T1, T2, T3, T4]) Combine(f ...Fn4[T1, T2, T3, T4]) Action4[T1, T2, T3, T4] {
	m := a.combine(unsafe.Pointer(&f))
	return Action4[T1, T2, T3, T4]{m}
}

func (a Action4[T1, T2, T3, T4]) CombineDelegate(v Action2[T1, T2]) Action4[T1, T2, T3, T4] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action4[T1, T2, T3, T4]{m}
}

func (a Action4[T1, T2, T3, T4]) Remove(f ...Fn4[T1, T2, T3, T4]) Action4[T1, T2, T3, T4] {
	m := a.remove(unsafe.Pointer(&f))
	return Action4[T1, T2, T3, T4]{m}
}

func (a Action4[T1, T2, T3, T4]) RemoveDelegate(v Action4[T1, T2, T3, T4]) Action4[T1, T2, T3, T4] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action4[T1, T2, T3, T4]{m}
}

func (a Action4[T1, T2, T3, T4]) Invoke(x1 T1, x2 T2, x3 T3, x4 T4) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn4[T1, T2, T3, T4])(unsafe.Pointer(&funcVar)))(x1, x2, x3, x4)
	}
}

type Fn5[T1, T2, T3, T4, T5 any] func(T1, T2, T3, T4, T5)

type Action5[T1, T2, T3, T4, T5 any] struct {
	multicastDelegate
}

func (a Action5[T1, T2, T3, T4, T5]) Equals(other Action5[T1, T2, T3, T4, T5]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action5[T1, T2, T3, T4, T5]) Combine(f ...Fn5[T1, T2, T3, T4, T5]) Action5[T1, T2, T3, T4, T5] {
	m := a.combine(unsafe.Pointer(&f))
	return Action5[T1, T2, T3, T4, T5]{m}
}

func (a Action5[T1, T2, T3, T4, T5]) CombineDelegate(v Action5[T1, T2, T3, T4, T5]) Action5[T1, T2, T3, T4, T5] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action5[T1, T2, T3, T4, T5]{m}
}

func (a Action5[T1, T2, T3, T4, T5]) Remove(f ...Fn5[T1, T2, T3, T4, T5]) Action5[T1, T2, T3, T4, T5] {
	m := a.remove(unsafe.Pointer(&f))
	return Action5[T1, T2, T3, T4, T5]{m}
}

func (a Action5[T1, T2, T3, T4, T5]) RemoveDelegate(v Action5[T1, T2, T3, T4, T5]) Action5[T1, T2, T3, T4, T5] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action5[T1, T2, T3, T4, T5]{m}
}

func (a Action5[T1, T2, T3, T4, T5]) Invoke(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn5[T1, T2, T3, T4, T5])(unsafe.Pointer(&funcVar)))(x1, x2, x3, x4, x5)
	}
}

type Fn6[T1, T2, T3, T4, T5, T6 any] func(T1, T2, T3, T4, T5, T6)

type Action6[T1, T2, T3, T4, T5, T6 any] struct {
	multicastDelegate
}

func (a Action6[T1, T2, T3, T4, T5, T6]) Equals(other Action6[T1, T2, T3, T4, T5, T6]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action6[T1, T2, T3, T4, T5, T6]) Combine(f ...Fn6[T1, T2, T3, T4, T5, T6]) Action6[T1, T2, T3, T4, T5, T6] {
	m := a.combine(unsafe.Pointer(&f))
	return Action6[T1, T2, T3, T4, T5, T6]{m}
}

func (a Action6[T1, T2, T3, T4, T5, T6]) CombineDelegate(v Action6[T1, T2, T3, T4, T5, T6]) Action6[T1, T2, T3, T4, T5, T6] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action6[T1, T2, T3, T4, T5, T6]{m}
}

func (a Action6[T1, T2, T3, T4, T5, T6]) Remove(f ...Fn6[T1, T2, T3, T4, T5, T6]) Action6[T1, T2, T3, T4, T5, T6] {
	m := a.remove(unsafe.Pointer(&f))
	return Action6[T1, T2, T3, T4, T5, T6]{m}
}

func (a Action6[T1, T2, T3, T4, T5, T6]) RemoveDelegate(v Action6[T1, T2, T3, T4, T5, T6]) Action6[T1, T2, T3, T4, T5, T6] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action6[T1, T2, T3, T4, T5, T6]{m}
}

func (a Action6[T1, T2, T3, T4, T5, T6]) Invoke(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn5[T1, T2, T3, T4, T5])(unsafe.Pointer(&funcVar)))(x1, x2, x3, x4, x5)
	}
}

type Fn7[T1, T2, T3, T4, T5, T6, T7 any] func(T1, T2, T3, T4, T5, T6, T7)

type Action7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	multicastDelegate
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) Equals(other Action7[T1, T2, T3, T4, T5, T6, T7]) bool {
	return a.multicastDelegate.equals(other.multicastDelegate)
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) Combine(f ...Fn7[T1, T2, T3, T4, T5, T6, T7]) Action7[T1, T2, T3, T4, T5, T6, T7] {
	m := a.combine(unsafe.Pointer(&f))
	return Action7[T1, T2, T3, T4, T5, T6, T7]{m}
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) CombineDelegate(v Action5[T1, T2, T3, T4, T5]) Action7[T1, T2, T3, T4, T5, T6, T7] {
	m := a.combineDelegate(v.multicastDelegate)
	return Action7[T1, T2, T3, T4, T5, T6, T7]{m}
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) Remove(f ...Fn5[T1, T2, T3, T4, T5]) Action7[T1, T2, T3, T4, T5, T6, T7] {
	m := a.remove(unsafe.Pointer(&f))
	return Action7[T1, T2, T3, T4, T5, T6, T7]{m}
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) RemoveDelegate(v Action5[T1, T2, T3, T4, T5]) Action7[T1, T2, T3, T4, T5, T6, T7] {
	m := a.removeDelegate(v.multicastDelegate)
	return Action7[T1, T2, T3, T4, T5, T6, T7]{m}
}

func (a Action7[T1, T2, T3, T4, T5, T6, T7]) Invoke(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5, x6 T6, x7 T7) {
	for _, invocation := range a.invocations {
		funcVar := unsafe.Pointer(&invocation)
		(*(*Fn7[T1, T2, T3, T4, T5, T6, T7])(unsafe.Pointer(&funcVar)))(x1, x2, x3, x4, x5, x6, x7)
	}
}
