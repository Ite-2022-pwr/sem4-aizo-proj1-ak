package analysis

import (
	"log"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
	"golang.org/x/exp/constraints"
)

// IsArraySorted sprawdza, czy tablica jest posortowana
func IsArraySorted[T constraints.Ordered](array []T) bool {
  n := len(array)
  for i := 1; i < n; i++ {
    if array[i - 1] > array[i] {
      return false
    }
  }
  return true
}

// AssertSortedArray upewnia się, czy tablica jest posortowana.
// Jeśli nie jest, wypisuje komunikat błędu i kończy działanie programu sygnalizując niepowodzenie.
func AssertSortedArray[T constraints.Ordered](array []T) {
  if IsArraySorted(array) {
    log.Println(utils.GreenColor("[+] Tablica posortowana poprawnie"))
  } else {
    log.Fatal(utils.RedColor("[!!] Tablica nie została posortowana poprawnie!"))
  }
}
