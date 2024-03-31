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
 
  start := time.Now()
  prompt := fmt.Sprintf("Sortowanie przez kopcowanie dla typu danych %T", array)
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)

  sort.HeapSort(array)

  utils.PrintTimeElapsed(start, prompt)

  if IsArraySorted(array) {
    log.Println("[+] Tablica posortowana poprawnie")
  } else {
    log.Fatal("[!!] Tablica nie została posortowana poprawnie!")
  }

  return array
}
