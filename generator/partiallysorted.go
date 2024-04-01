package generator

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// GeneratePartiallySortedArrayInt generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayInt(n, sortedPercentage int, verbose bool) []int {
  array := make([]int, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayInt(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int() % 100
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}

// GeneratePartiallySortedArrayInt32 generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayInt32(n, sortedPercentage int, verbose bool) []int32 {
  array := make([]int32, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayInt32(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int31() % 100
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}

// GeneratePartiallySortedArrayInt64 generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayInt64(n, sortedPercentage int, verbose bool) []int64 {
  array := make([]int64, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayInt64(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int63() % 1000
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}

// GeneratePartiallySortedArrayFloat32 generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayFloat32(n, sortedPercentage int, verbose bool) []float32 {
  array := make([]float32, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayFloat32(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Float32()
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}

// GeneratePartiallySortedArrayFloat64 generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayFloat64(n, sortedPercentage int, verbose bool) []float64 {
  array := make([]float64, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayFloat64(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Float64()
    if r < array[sortedNum - 1] {
      r += array[sortedNum - 1]
    }
    array[i] = r
  }

  return array
}

// GeneratePartiallySortedArrayByte generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayByte(n, sortedPercentage int, verbose bool) []byte {
  array := make([]byte, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayByte(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := sortedNum; i < n; i++ {
    r := rng.Int()
    if byte(r) < array[sortedNum - 1] {
      r = min(int(array[sortedNum - 1]) + r, 255)
    }
    array[i] = byte(r)
  }

  return array
}

// GeneratePartiallySortedArrayString generuje częściowo posortowaną tablicę
func GeneratePartiallySortedArrayString(n, sortedPercentage int, verbose bool) []string {
  array := make([]string, n)

  sortedNum := n * sortedPercentage / 100

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej w %v tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%v%%", sortedPercentage)), utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sortedPart := GenerateAscendingSortedArrayString(sortedNum, false)
  copy(array[:sortedNum], sortedPart)

  for i := sortedNum; i < n; i++ {
    r := GenerateRandomString(5)
    if r < array[sortedNum - 1] {
      r = array[sortedNum - 1] + r
    }
    array[i] = r
  }

  return array
}
