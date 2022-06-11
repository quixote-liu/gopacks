package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func main() {
	randInt()
}

func randInt() {
	maxBigInt := big.NewInt(math.MaxInt64)
	rint, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rint: %v\n", rint)
}
