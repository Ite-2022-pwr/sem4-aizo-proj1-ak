// Package generator zawiera funckje generujące dane wejściowe
package generator

import (
	"log"
	"math/rand"
	"time"
)

// GenerateRandomArrayInt generuje n losowych liczb typu int modulo mod
func GenerateRandomArrayInt(n, mod int) []int {
  array := make([]int, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int() % mod
  }

  return array
}


// GenerateRandomArrayInt32 generuje n losowych liczb typu int32 modulo mod
func GenerateRandomArrayInt32(n int, mod int32) []int32 {
  array := make([]int32, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int31() % mod
  }

  return array
}


// GenerateRandomArrayInt64 generuje n losowych liczb typu int64 modulo mod
func GenerateRandomArrayInt64(n int, mod int64) []int64 {
  array := make([]int64, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int63() % mod
  }

  return array
}

// GenerateRandomArrayFloat32 generuje n losowych liczb typu float32
func GenerateRandomArrayFloat32(n int) []float32 {
  array := make([]float32, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Float32()
  }

  return array
}


// GenerateRandomArrayFloat64 generuje n losowych liczb typu float64
func GenerateRandomArrayFloat64(n int) []float64 {
  array := make([]float64, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Float64()
  }

  return array
}

// GenerateRandomArrayByte generuje n losowych bajtów
func GenerateRandomArrayByte(n int) []byte {
  array := make([]byte, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  if _, err := rng.Read(array); err != nil {
    log.Fatal("[!!] Błąd przy generowaniu tablicy bajtów:", err)
  } 

  return array
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
  byteArray := make([]byte, length)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < length; i++ {
    byteArray[i] = letters[rng.Intn(len(letters))]  // wybierz losową literę
  }

  return string(byteArray)
}

// GenerateRandomArrayString generuje n losowych ciągów znaków o maksymalnej długości maxLength
func GenerateRandomArrayString(n, maxLength int) []string {
  array := make([]string, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = generateRandomString(rng.Intn(maxLength - 1) + 1)
  }

  return array
}
