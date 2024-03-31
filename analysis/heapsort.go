package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// HeapSortAnalysis mierzy czas sortowania przez kopcowanie.
func (sa *SortAnalyzer[T]) HeapSortAnalysis() []T {
  array := sa.GetDataCopy()
  
  log.Println("Rozmiar Tablicy:", utils.YellowColor(sa.DataLength))
  log.Println("Tablica (max pierwsze 20 elementów):", utils.YellowColor(array[:20]))

  start := time.Now()
  prompt := utils.BlueColor(fmt.Sprintf("Sortowanie przez kopcowanie dla typu danych %T", array))
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)

  sort.HeapSort(array)

  utils.PrintTimeElapsed(start, prompt)

  AssertSortedArray(array)

  return array
}
