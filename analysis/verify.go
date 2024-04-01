package analysis

import (
	"log"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
	"golang.org/x/exp/constraints"
)

// IsArraySortedAscending sprawdza, czy tablica jest posortowana rosnąco.
func IsArraySortedAscending[T constraints.Ordered](array []T) bool {
  n := len(array)
  for i := 1; i < n; i++ {
    if array[i - 1] > array[i] {
      return false
    }
  }
  return true
}

// IsArraySortedDescending sprawdza, czy tablica jest posortowana malejąco.
func IsArraySortedDescending[T constraints.Ordered](array []T) bool {
  n := len(array)
  for i := 1; i < n; i++ {
    if array[i - 1] < array[i] {
      return false
    }
  }
  return true
}

// AssertSortedAscendingArray upewnia się, czy tablica jest posortowana rosnąco.
// Jeśli nie jest, wypisuje komunikat błędu i kończy działanie programu sygnalizując niepowodzenie.
func AssertSortedAscendingArray[T constraints.Ordered](array []T) {
  if IsArraySortedAscending(array) {
    log.Println(utils.GreenColor("[+] Tablica posortowana rosnąco poprawnie"))
  } else {
    log.Fatal(utils.RedColor("[!!] Tablica nie została posortowana rosnąco poprawnie!"))
  }
}


// AssertSortedDescendingArray upewnia się, czy tablica jest posortowana malejąco.
// Jeśli nie jest, wypisuje komunikat błędu i kończy działanie programu sygnalizując niepowodzenie.
func AssertSortedDescendingArray[T constraints.Ordered](array []T) {
  if IsArraySortedDescending(array) {
    log.Println(utils.GreenColor("[+] Tablica posortowana malejąco poprawnie"))
  } else {
    log.Fatal(utils.RedColor("[!!] Tablica nie została posortowana malejąco poprawnie!"))
  }
}
