package generator

import (
	"math/rand"
	"time"
)

func GenerateRandomArrayInt(n, mod int) []int {
  array := make([]int, n)
  
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < n; i++ {
    array[i] = rng.Int() % mod
  }

  return array
}
