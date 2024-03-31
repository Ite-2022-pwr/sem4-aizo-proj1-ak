package measurement

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

func (sm *SortMeter[T]) QuickSort(pivot string) []T {
  array := sm.GetDataCopy()

  var getPivot sort.PivotChoosingMethod[T]
  switch pivot {
  case "lewy":
    getPivot = sort.GetPivotLow
  case "prawy":
    getPivot = sort.GetPivotHigh
  case "środek":
    getPivot = sort.GetPivotMiddle
  case "losowy":
    getPivot = sort.GetPivotRandom
  default:
    log.Fatalf("Błąd: Nieznany typ pivota dla QuickSort: %s", pivot)
  }

  start := time.Now()
  prompt := fmt.Sprintf("QuickSort (pivot: %s) dla typu danych %T", pivot, array)
  log.Printf("Rozpoczynanie %s", prompt)
  defer utils.PrintTimeElapsed(start, prompt)

  sort.QuickSort(array, 0, sm.DataLength - 1, getPivot)

  return array
}
