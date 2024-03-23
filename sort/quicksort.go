package sort

import "golang.org/x/exp/constraints"

type PivotChoosingMethod[T constraints.Ordered] func(array []T, low, high int) T

func QuickSort[T constraints.Ordered](array []T, low, high int, getPivot PivotChoosingMethod[T]) {
  if low < high {
    partitionIndex := partition(array, low, high, getPivot)
    QuickSort(array, low, partitionIndex - 1, getPivot)
    QuickSort(array, partitionIndex + 1, high, getPivot)
  }
}

func partition[T constraints.Ordered](array []T, low, high int, getPivot PivotChoosingMethod[T]) int {
  pivot := getPivot(array, low, high)

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

func GetPivotLow[T constraints.Ordered](array []T, low, high int) T {
  return array[low]
}

func GetPivotHigh[T constraints.Ordered](array []T, low, high int) T {
  return array[high]
}

func GetPivotMiddle[T constraints.Ordered](array []T, low, high int) T {
  return array[(low + high) / 2]
}
