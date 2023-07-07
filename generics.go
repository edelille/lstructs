package lstructs

import (
	"golang.org/x/exp/constraints"
)

type EqualsFn[T comparable] func(x, y T) bool

type LessFn[T comparable] func(x, y T) bool

func Equals[T comparable](x, y T) bool {
	return x == y
}

func Less[T constraints.Ordered](x, y T) bool {
	return x < y
}

func Swap[T any](arr []T, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
