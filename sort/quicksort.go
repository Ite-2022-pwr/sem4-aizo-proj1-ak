package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// PivotChoosingMethod to typ funkcyjny określający funkcję służącą do wybierania osi
// dla sortowania szybkiego
type PivotChoosingMethod[T constraints.Ordered] func(array []T, low, high int) T

// QuickSort sortuje tablicę metodą sortowania szybkiego
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

// GetPivotLow wybiera jako oś lewy element z danego zakresu
func GetPivotLow[T constraints.Ordered](array []T, low, _ int) T {
  return array[low]
}

// GetPivotHigh wybiera jako oś prawy element z danego zakresu
func GetPivotHigh[T constraints.Ordered](array []T, _, high int) T {
  return array[high]
}

// GetPivotMiddle wybiera jako oś środkowy element z danego zakresu
func GetPivotMiddle[T constraints.Ordered](array []T, low, high int) T {
  return array[(low + high) / 2]
}

// GetPivotRandom wbyiera jako oś losowy element z danego zakresu
func GetPivotRandom[T constraints.Ordered](array []T, low, high int) T {
  return array[rand.Intn(high - low + 1) + low]
}

