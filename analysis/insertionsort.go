package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// InsertionSort mierzy czas sortowania przez wstawianie
func (sm *SortMeter[T]) InsertionSort() []T {
  array := sm.GetDataCopy()
 
  start := time.Now()
  prompt := fmt.Sprintf("Sortowanie przez wstawianie dla typu danych %T", array)
  log.Printf("[*] Rozpoczynanie: %s\n", prompt)
  defer utils.PrintTimeElapsed(start, prompt)

  sort.InsertionSort(array) 

  return array
}