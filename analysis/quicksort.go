package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// QuickSortAnalysis mierzy czas sortowania szybkiego.
// Parametr pivot musi przyjmować jedną z wartości: "lewy", "prawy", "środkowy" lub "losowy".
func (sa *SortAnalyzer[T]) QuickSortAnalysis(pivotType string) []T {
  array := sa.GetDataCopy()

  log.Println("Tablica:", utils.YellowColor(array))

  var getPivot sort.PivotIndexChoosingMethod[T]
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
    log.Fatal(utils.RedColor(fmt.Sprintf("[!!] Błąd: Nieznany typ pivota dla QuickSort: %s. Dopuszczalne: lewy, prawy, środkowy, losowy", pivotType)))
  }

  start := time.Now()
  prompt := utils.BlueColor(fmt.Sprintf("Sortowanie szybkie (pivot: %s) dla typu danych %T", pivotType, array))
  log.Println("[*] Rozpoczynanie:", prompt)

  sort.QuickSort(array, 0, sa.DataLength - 1, getPivot)

  utils.PrintTimeElapsed(start, prompt)

  AssertSortedArray(array)

  return array
}
