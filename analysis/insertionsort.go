package analysis

import (
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// InsertionSortAnalysis mierzy czas sortowania przez wstawianie.
func (sa *SortAnalyzer[T]) InsertionSortAnalysis() []T {
  array := sa.GetDataCopy()

  log.Println("Typ danych:", utils.YellowColor(sa.DataTypyName))
  log.Println("Rozmiar Tablicy:", utils.YellowColor(sa.DataLength))
  if sa.DataLength > 20 {
    log.Println("Tablica (pierwsze 20 elementów):", utils.YellowColor(array[:20]))
  } else {
    log.Println("Tablica:", utils.YellowColor(array))
  }

  prompt := utils.BlueColor("Sortowanie przez wstawianie")
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)
  start := time.Now()

  sort.InsertionSort(array) 

  utils.PrintTimeElapsed(start, prompt)

  AssertSortedAscendingArray(array)

  return array
}
