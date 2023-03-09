package slicekit

type EachParams[T any] struct {
	Index int
	Value T
	Break func()
}
