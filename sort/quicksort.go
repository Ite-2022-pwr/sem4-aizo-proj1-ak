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
	pivotIndex := getPivot(array, low, high)
	pivot := array[pivotIndex]
	i, j := low, high

	for i <= j {
		for array[i] < pivot {
			i++
		}

		for array[j] > pivot {
			j--
		}

		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	if low < j {
		QuickSort(array, low, j, getPivot)
	}

	if i < high {
		QuickSort(array, i, high, getPivot)
	}
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
	return rand.Intn(high-low+1) + low
}
