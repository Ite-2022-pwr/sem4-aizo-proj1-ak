package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](array []T, low, high int) {
  if low < high {
    partitionIndex := partition(array, low, high)
    QuickSort(array, low, partitionIndex-1)
    QuickSort(array, partitionIndex + 1, high)
  }
}

func partition[T constraints.Ordered](array []T, low, high int) int {
  // pivot := array[low] // not sorting strings 
  pivot := array[high] // index out of range
  // pivot := array[(low + high) / 2]

  i, j := low, high
  
  for i < j {
    for array[j] > pivot {
      j--
    }
    for array[i] < pivot {
      i++
    }

    if i < j {
      array[i], array[j] = array[j], array[i]
    }
  }
  return j
}
