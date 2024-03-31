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
  array[high], array[pivotIndex] = array[pivotIndex], array[high]
  pivot := array[high]

  i := low - 1
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	i++
	array[i], array[high] = array[high], array[i]
	
  return i
}

// GetPivotLow wybiera jako oś lewy element z danego zakresu
func GetPivotLow[T constraints.Ordered](array []T, low, _ int) int {
  return low
}

// GetPivotHigh wybiera jako oś prawy element z danego zakresu
func GetPivotHigh[T constraints.Ordered](array []T, _, high int) int {
  return high
}

// GetPivotMiddle wybiera jako oś środkowy element z danego zakresu
func GetPivotMiddle[T constraints.Ordered](array []T, low, high int) int {
  return (low + high) / 2
}

// GetPivotRandom wbyiera jako oś losowy element z danego zakresu
func GetPivotRandom[T constraints.Ordered](array []T, low, high int) int {
  return rand.Intn(high - low + 1) + low
}

