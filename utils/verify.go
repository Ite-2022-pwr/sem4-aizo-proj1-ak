package utils

import "golang.org/x/exp/constraints"

// IsArraySorted sprawdza, czy tablica jest posortowana
func IsArraySorted[T constraints.Ordered](array []T) bool {
  n := len(array)
  for i := 1; i < n; i++ {
    if array[i - 1] > array[i] {
      return false
    }
  }
  return true
}
