package generator

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/sort"
	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// GenerateAscendingSortedArrayInt genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedArrayInt(n int, verbose bool) []int {
  array := make([]int, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = i + intSeed
  }

  return array
}

// GenerateDescendingSortedArrayInt genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedArrayInt(n int, verbose bool) []int {
  array := make([]int, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = n - i - 1 + intSeed
  }

  return array
}

// GenerateAscendingSortedArrayInt32 genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedArrayInt32(n int, verbose bool) []int32 {
  array := make([]int32, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = int32(i + intSeed)
  }

  return array
}

// GenerateDescendingSortedArrayInt32 genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedArrayInt32(n int, verbose bool) []int32 {
  array := make([]int32, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = int32(n - i - 1 + intSeed)
  }

  return array
}


// GenerateAscendingSortedArrayInt64 genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedArrayInt64(n int, verbose bool) []int64 {
  array := make([]int64, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = int64(i + intSeed)
  }

  return array
}

// GenerateDescendingSortedArrayInt64 genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedArrayInt64(n int, verbose bool) []int64 {
  array := make([]int64, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))
  intSeed := rng.Int() % 100

  for i := 0; i < n; i++ {
    array[i] = int64(n - i - 1 + intSeed)
  }

  return array
}

// GenerateAscendingSortedArrayFloat32 generuje posortowany rosnąco
// ciąg liczb typu float32
func GenerateAscendingSortedArrayFloat32(n int, verbose bool) []float32 {
  array := make([]float32, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float32() * 100

  for i := 0; i < n; i++ {
    array[i] = float32(i) + floatSeed
  }

  return array
}

// GenerateDescendingSortedArrayFloat32 generuje posortowany malejąco
// ciąg liczb typu float32
func GenerateDescendingSortedArrayFloat32(n int, verbose bool) []float32 {
  array := make([]float32, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float32() * 100

  for i := 0; i < n; i++ {
    array[i] = float32(n - i) - floatSeed
  }

  return array
}


// GenerateAscendingSortedArrayFloat64 generuje posortowany rosnąco
// ciąg liczb typu float64
func GenerateAscendingSortedArrayFloat64(n int, verbose bool) []float64 {
  array := make([]float64, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float64() * 100

  for i := 0; i < n; i++ {
    array[i] = float64(i) + floatSeed
  }

  return array
}

// GenerateDescendingSortedArrayFloat64 generuje posortowany malejąco
// ciąg liczb typu float64
func GenerateDescendingSortedArrayFloat64(n int, verbose bool) []float64 {
  array := make([]float64, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float64() * 100

  for i := 0; i < n; i++ {
    array[i] = float64(n - i) - floatSeed
  }

  return array
}

// GenerateAscendingSortedArrayByte generuje posortowany rosnąco
// ciąg bajtów o rozmiarze n
func GenerateAscendingSortedArrayByte(n int, verbose bool) []byte {
  array := make([]byte, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  if n <= 256 {
    for i := 0; i < n; i++ {
      array[i] = byte(i)
    }
  } else {
    times := n / 255 + 1
    val := 0
    for i := 0; i < n; i += times {
      for j := 0; j < times && i + j < n; j++ {
        array[i + j] = byte(val)
      }
      if val < 256 {
        val++
      }
    }
  }

  return array
}

// GenerateDescendingSortedArrayByte generuje posortowany malejąco
// ciąg bajtów o rozmiarze n
func GenerateDescendingSortedArrayByte(n int, verbose bool) []byte {
  array := make([]byte, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  if n <= 256 {
    for i := 0; i < n; i++ {
      array[i] = byte(n - i - 1)
    }
  } else {
    times := n / 255 + 1
    val := 255
    for i := 0; i < n; i += times {
      for j := 0; j < times && i + j < n; j++ {
        array[i + j] = byte(val)
      }
      if val > 0 {
        val--
      }
    }
  }

  return array
}

// GenerateAscendingSortedArrayString generuje posortowaną rosnąco
// tablicę ciągów znaków o rozmiarze n
func GenerateAscendingSortedArrayString(n int, verbose bool) []string {
  array := GenerateRandomArrayString(n, 10, false)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej rosnąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sort.QuickSort(array, 0, n - 1, sort.GetPivotHigh)

  return array
}

// GenerateDescendingSortedArrayString generuje posortowaną malejąco
// tablicę ciągów znakóœ o rozmiarze n
func GenerateDescendingSortedArrayString(n int, verbose bool) []string {
  array := GenerateRandomArrayString(n, 10, false)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie posortowanej malejąco tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  sort.QuickSort(array, 0, n - 1, sort.GetPivotHigh)

  slices.Reverse(array)

  return array
}
