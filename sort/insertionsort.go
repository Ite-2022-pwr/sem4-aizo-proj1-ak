package sort

import "golang.org/x/exp/constraints"

// InsertionSort sortuje tablicę metodą przez wstawianie
func InsertionSort[T constraints.Ordered](array []T) {
  n := len(array)
  for i := 0; i < n; i++ {
    key := array[i]
    j := i - 1

    for j >= 0 && array[j] > key {
      array[j + 1] = array[j]
      j--
    }
    array[j + 1] = key
  }
}
