// Package generator zawiera funckje generujące dane wejściowe
package generator

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Ite-2022-pwr/sem4-aizo-proj1-ak/utils"
)

// GenerateRandomArrayInt generuje n losowych liczb typu int modulo mod
func GenerateRandomArrayInt(n, mod int, verbose bool) []int {
  array := make([]int, n)
  
  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int() % mod
  }

  return array
}


// GenerateRandomArrayInt32 generuje n losowych liczb typu int32 modulo mod
func GenerateRandomArrayInt32(n int, mod int32, verbose bool) []int32 {
  array := make([]int32, n)
  
  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int31() % mod
  }

  return array
}


// GenerateRandomArrayInt64 generuje n losowych liczb typu int64 modulo mod
func GenerateRandomArrayInt64(n int, mod int64, verbose bool) []int64 {
  array := make([]int64, n)
  
  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int63() % mod
  }

  return array
}

// GenerateRandomArrayFloat32 generuje n losowych liczb typu float32
func GenerateRandomArrayFloat32(n int, mul float32, verbose bool) []float32 {
  array := make([]float32, n)
  
  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Float32() * mul
  }

  return array
}


// GenerateRandomArrayFloat64 generuje n losowych liczb typu float64
func GenerateRandomArrayFloat64(n int, mul float64, verbose bool) []float64 {
  array := make([]float64, n)
  
  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Float64() * mul
  }

  return array
}

// GenerateRandomArrayByte generuje n losowych bajtów
func GenerateRandomArrayByte(n int, verbose bool) []byte {
  array := make([]byte, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  if _, err := rng.Read(array); err != nil {
    log.Fatal("[!!] Błąd przy generowaniu tablicy bajtów:", err)
  } 

  return array
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString generuje losowy ciąg znaków o zadanej długości
func GenerateRandomString(length int) string {
  byteArray := make([]byte, length)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < length; i++ {
    byteArray[i] = letters[rng.Intn(len(letters))]  // wybierz losową literę
  }

  return string(byteArray)
}

// GenerateRandomArrayString generuje n losowych ciągów znaków o maksymalnej długości maxLength
func GenerateRandomArrayString(n, maxLength int, verbose bool) []string {
  array := make([]string, n)

  if verbose {
    log.Println(fmt.Sprintf("[*] Generowanie losowej tablicy typu %v o rozmiarze %v", utils.YellowColor(fmt.Sprintf("%T", array)), utils.YellowColor(n)))
  }

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = GenerateRandomString(rng.Intn(maxLength - 1) + 1)
  }

  return array
}
