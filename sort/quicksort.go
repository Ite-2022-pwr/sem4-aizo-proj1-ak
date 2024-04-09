package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// PivotIndexChoosingMethod to typ funkcyjny określający funkcję służącą do wybierania osi
// dla sortowania szybkiego. Zwraca index wybranej osi.
type PivotIndexChoosingMethod[T constraints.Ordered] func(array []T, low, high int) int

// QuickSort mierzy czas sortowania szybkiego
func QuickSort[T constraints.Ordered](array []T, low, high int, getPivot PivotIndexChoosingMethod[T]) {
  if low < high {
    partitionIndex := partition(array, low, high, getPivot)
    QuickSort(array, low, partitionIndex - 1, getPivot)
    QuickSort(array, partitionIndex + 1, high, getPivot)
  }
}

func partition[T constraints.Ordered](array []T, low, high int, getPivot PivotIndexChoosingMethod[T]) int {
  pivotIndex := getPivot(array, low, high)
  array[low], array[pivotIndex] = array[pivotIndex], array[low]
  pivot := array[low]

  m := low
  for i := low + 1; i <= high; i++ {
    if array[i] < pivot {
      m++
      array[m], array[i] = array[i], array[m]
    }
  }
  array[low], array[m] = array[m], array[low]
  return m
}

// GetPivotLow wybiera jako oś lewy element z danego zakresu
func GetPivotLow[T constraints.Ordered](_ []T, low, _ int) int {
  return low
}

// GetPivotHigh wybiera jako oś prawy element z danego zakresu
func GetPivotHigh[T constraints.Ordered](_ []T, _, high int) int {
  return high
}

// GetPivotMiddle wybiera jako oś środkowy element z danego zakresu
func GetPivotMiddle[T constraints.Ordered](_ []T, low, high int) int {
  return (low + high) / 2
}

// GetPivotRandom wbyiera jako oś losowy element z danego zakresu
func GetPivotRandom[T constraints.Ordered](_ []T, low, high int) int {
  return rand.Intn(high - low + 1) + low
}

