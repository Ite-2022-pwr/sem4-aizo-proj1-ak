package sort

import (
	"math"

	"golang.org/x/exp/constraints"
)

// GapCalculatingFormula to typ określający funkcję służącą
// do obliczania odstępów w sortowaniu Shella
type GapCalculatingFormula func(n, k int) int

// ShellSort sortuje tablicę metodą Donalda Shella
func ShellSort[T constraints.Ordered](array []T, calculateGap GapCalculatingFormula) {
  n, k := len(array), 1

  for gap := calculateGap(n, k); gap >= 1; gap = calculateGap(n, k) {
    for i := gap; i < n; i++ {
      temp := array[i]
      var j int
      for j = i; j >= gap && array[j - gap] > temp; j -= gap {
        array[j] = array[j - gap]
      }
      array[j] = temp
    }
    k++
    if gap == 1 {
      break
    }
  }
}

// CalculateShellGap oblicza odstęp zgodnie ze wzorem: floor(N / 2^k)
// Autor: Shell, 1959r.
func CalculateShellGap(n, k int) int {
  return int(math.Floor(float64(n) / math.Pow(2.0, float64(k))))
}

// CalculateFrankLazarusGap oblicza odstęp zgodnie zgodnie wzorem: 2 * floor(N / 2^(k + 1)) + 1 
// Autorzy: Frank, Lazarus, 1960r.
func CalculateFrankLazarusGap(n, k int) int {
  return 2 * int(math.Floor(float64(n) / math.Pow(2.0, float64(k + 1)))) + 1
}
