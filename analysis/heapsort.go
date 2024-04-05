package analysis

import (
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// HeapSortAnalysis mierzy czas sortowania przez kopcowanie.
func (sa *SortAnalyzer[T]) HeapSortAnalysis() ([]T, float64) {
  array := sa.GetDataCopy()
  
  log.Println("Typ danych:", utils.YellowColor(sa.DataTypyName))
  log.Println("Rozmiar Tablicy:", utils.YellowColor(sa.DataLength))
  if sa.DataLength > 20 {
    log.Println("Tablica (pierwsze 20 elementów):", utils.YellowColor(array[:20]))
  } else {
    log.Println("Tablica:", utils.YellowColor(array))
  }

  prompt := utils.BlueColor("Sortowanie przez kopcowanie")
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)
  start := time.Now()

  sort.HeapSort(array)

  timeElapsed := utils.PrintTimeElapsed(start, prompt)

  AssertSortedAscendingArray(array)

  return array, timeElapsed
}
