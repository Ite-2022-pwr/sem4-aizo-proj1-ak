// Package measurement służy do przeprowadzania pomiarów algorytmów sortujących z pakietu sort
package analysis

import "golang.org/x/exp/constraints"

// SortMeter to struktura odpowiedzialna za przeprowadzanie pomiarów
// czasu poszczególnych sortowań: InsertionSort, QuickSort, ShellSort i HeapSort
type SortMeter[T constraints.Ordered] struct {
  Data []T
  DataLength int
}

// GetDataCopy kopiuje i zwraca dane (tablicę) obiektu SortMeter
func (sm *SortMeter[T]) GetDataCopy() []T {
  newArray := make([]T, sm.DataLength)
  copy(newArray, sm.Data)
  return newArray
}

// NewSortMeter zwraca nowy obiekt SortMeter utworzony na podstawie zadanej tablicy
func NewSortMeter[T constraints.Ordered](array []T) *SortMeter[T] {
  return &SortMeter[T]{Data: array, DataLength: len(array)}
}
