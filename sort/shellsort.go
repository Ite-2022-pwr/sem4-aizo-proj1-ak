package sort

import (
	"math"

	"golang.org/x/exp/constraints"
)

type GapCalculatingFormula func(n, k int) int

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

func CalculateShellGap(n, k int) int {
  return int(math.Floor(float64(n) / math.Pow(2.0, float64(k))))
}

func CalculateFrankLazarusGap(n, k int) int {
  return 2 * int(math.Floor(float64(n) / math.Pow(2.0, float64(k + 1)))) + 1
}
