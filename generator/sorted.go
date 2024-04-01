package generator

import (
	"math/rand"
	"time"
)

// GenerateAscendingSortedInt genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedInt(n int) []int {
  array := make([]int, n)

  for i := 0; i < n; i++ {
    array[i] = i
  }

  return array
}

// GenerateDescendingSortedInt genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedInt(n int) []int {
  array := make([]int, n)

  for i := 0; i < n; i++ {
    array[i] = n - i - 1
  }

  return array
}

// GenerateAscendingSortedInt32 genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedInt32(n int) []int32 {
  array := make([]int32, n)

  for i := 0; i < n; i++ {
    array[i] = int32(i)
  }

  return array
}

// GenerateDescendingSortedInt64 genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedInt32(n int) []int32 {
  array := make([]int32, n)

  for i := 0; i < n; i++ {
    array[i] = int32(n - i - 1)
  }

  return array
}


// GenerateAscendingSortedInt64 genruje rosnący ciąg liczb [0, n)
func GenerateAscendingSortedInt64(n int) []int64 {
  array := make([]int64, n)

  for i := 0; i < n; i++ {
    array[i] = int64(i)
  }

  return array
}

// GenerateDescendingSortedInt64 genruje malejący ciąg liczb od n - 1 do 0
func GenerateDescendingSortedInt64(n int) []int64 {
  array := make([]int64, n)

  for i := 0; i < n; i++ {
    array[i] = int64(n - i - 1)
  }

  return array
}

// GenerateAscendingSortedFloat32 generuje posortowany rosnąco
// ciąg liczb typu float32
func GenerateAscendingSortedFloat32(n int) []float32 {
  array := make([]float32, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float32()
  floatSeed -= float32(int(floatSeed))  // część po przecinku

  for i := 0; i < n; i++ {
    array[i] = float32(i) + floatSeed
  }

  return array
}

// GenerateDescendingSortedFloat32 generuje posortowany malejąco
// ciąg liczb typu float32
func GenerateDescendingSortedFloat32(n int) []float32 {
  array := make([]float32, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float32()
  floatSeed -= float32(int(floatSeed))  // część po przecinku

  for i := 0; i < n; i++ {
    array[i] = float32(n - i) - floatSeed
  }

  return array
}


// GenerateAscendingSortedFloat64 generuje posortowany rosnąco
// ciąg liczb typu float64
func GenerateAscendingSortedFloat64(n int) []float64 {
  array := make([]float64, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float64()
  floatSeed -= float64(int(floatSeed))  // część po przecinku

  for i := 0; i < n; i++ {
    array[i] = float64(i) + floatSeed
  }

  return array
}

// GenerateDescendingSortedFloat64 generuje posortowany malejąco
// ciąg liczb typu float64
func GenerateDescendingSortedFloat64(n int) []float64 {
  array := make([]float64, n)

  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  floatSeed := rng.Float64()
  floatSeed -= float64(int(floatSeed))  // część po przecinku

  for i := 0; i < n; i++ {
    array[i] = float64(n - i) - floatSeed
  }

  return array
}
