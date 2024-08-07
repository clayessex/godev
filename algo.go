package algo

import "cmp"

// Create and return a new slice filled with the results of applying function f
// on every element of s
func Map[T any, O any](s []T, f func(T) O) []O {
	result := make([]O, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// Apply an accumulator function f to every element of s and return the final
// result. The accumulator is initialized with init, then with the result of
// subsequent iterations over s. The final result value returned from f is then
// returned from Reduce.
func Reduce[T any, O any](s []T, init O, f func(acc O, v T) O) O {
	result := init
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

// Create and return a new slice containing only the elements of s for which f
// returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Create and return a new slice such that s[0, middle) is swapped with
// s[middle, len(s))
func Rotate[T any](s []T, middle int) []T {
	s = append(s[middle:], s[:middle]...)
	return s
}

// Return the number of times that f returns true for each of the elements of s
func CountFunc[T any](s []T, f func(value T) bool) int {
	return Reduce(s, 0, func(acc int, v T) int {
		if f(v) {
			return acc + 1
		}
		return acc
	})
}

// Return the count of value in s using ==
func Count[T comparable](s []T, value T) int {
	return CountFunc(s, func(v T) bool {
		return v == value
	})
}

// Create and return a slice consisting of the two sorted slices a and b. The
// result is also sorted. Elements are ordered using <
func Merge[T cmp.Ordered](a, b []T) []T {
	return MergeFunc(a, b, func(x, y T) bool {
		return x < y
	})
}

// Create and return a slice consisting of the two sorted slices a and b. The
// result is also sorted. Elements are ordered using the function comp.
func MergeFunc[T any](a, b []T, comp func(x, y T) bool) []T {
	r := make([]T, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if comp(a[i], b[j]) {
			k := i
			for k != len(a) {
				if !comp(a[k], b[j]) {
					break
				}
				k++
			}
			r = append(r, a[i:k]...)
			i = k
		} else {
			k := j
			for k != len(b) {
				if !comp(b[k], a[i]) {
					break
				}
				k++
			}
			r = append(r, b[j:k]...)
			j = k
		}
	}

	if i < len(a) {
		r = append(r, a[i:]...)
	} else if j < len(b) {
		r = append(r, b[j:]...)
	}

	return r
}

// Return a result such that lo < v < hi
func Clamp[T cmp.Ordered](v, lo, hi T) T {
	return ClampFunc(v, lo, hi, func(a, b T) bool {
		return a < b
	})
}

// Return a result such that comp(lo, v) == true and comp(v, hi) == true
func ClampFunc[T any](v, lo, hi T, comp func(a, b T) bool) T {
	if comp(v, lo) {
		return lo
	} else if comp(hi, v) {
		return hi
	}
	return v
}

// Create and return a new slice consisting of the keys of map m
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Create and return a new slice consisting of the values of map m
func MapValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
