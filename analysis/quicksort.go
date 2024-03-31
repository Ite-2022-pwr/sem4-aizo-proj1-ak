package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// QuickSort sortuje zawartość tablicy SortMeter.Data metodą sortowania szybkiego
// operując na kopii tej tablicy. Zwraca posortowaną tablicę.
// Parametr pivot musi zawierać jedną z wartości: "lewy", "prawy", "środkowy" lub "losowy"
func (sm *SortMeter[T]) QuickSort(pivotType string) []T {
  array := sm.GetDataCopy()

  var getPivot sort.PivotChoosingMethod[T]
  switch pivotType {
  case "lewy":
    getPivot = sort.GetPivotLow
  case "prawy":
    getPivot = sort.GetPivotHigh
  case "środkowy":
    getPivot = sort.GetPivotMiddle
  case "losowy":
    getPivot = sort.GetPivotRandom
  default:
    log.Fatalf("[!!] Błąd: Nieznany typ pivota dla QuickSort: %s. Dopuszczalne: lewy, prawy, środkowy, losowy", pivotType)
  }

  start := time.Now()
  prompt := fmt.Sprintf("Sortowanie szybkie (pivot: %s) dla typu danych %T", pivotType, array)
  log.Println("[*] Rozpoczynanie:", prompt)
  defer utils.PrintTimeElapsed(start, prompt)

  sort.QuickSort(array, 0, sm.DataLength - 1, getPivot)

  return array
}
