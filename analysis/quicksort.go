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
func (sa *SortAnalyzer[T]) QuickSortAnalysis(pivotType string) ([]T, float64) {
  array := sa.GetDataCopy()

  log.Println("Typ danych:", utils.YellowColor(sa.DataTypyName))
  log.Println("Rozmiar Tablicy:", utils.YellowColor(sa.DataLength))
  if sa.DataLength > 20 {
    log.Println("Tablica (pierwsze 20 elementów):", utils.YellowColor(array[:20]))
  } else {
    log.Println("Tablica:", utils.YellowColor(array))
  }

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
    log.Println(utils.RedColor(fmt.Sprintf("[!!] Błąd: Nieznany typ pivota dla QuickSort: %s. Dopuszczalne: lewy, prawy, środkowy, losowy", pivotType)))
    return nil, 0.0
  }

  prompt := utils.BlueColor(fmt.Sprintf("Sortowanie szybkie (pivot: %s)", pivotType))
  log.Println("[*] Rozpoczynanie:", prompt)
  start := time.Now()

  sort.QuickSort(array, 0, sa.DataLength - 1, getPivot)

  timeElapsed := utils.PrintTimeElapsed(start, prompt)

  AssertSortedAscendingArray(array)

  return array, timeElapsed
}
