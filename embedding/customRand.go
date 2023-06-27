package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

// outer type implementing a method that has the same name as a method implemented by the inner type,
// the outher type method will be given priority
func (cr *customRand) Intn(n int) int {
	cr.count++
	return cr.Rand.Intn(n) + 1
}

func (cr *customRand) GetCount() int {
	return cr.count
}

func main() {
	cr := NewCustomRand(time.Now().UnixNano())

	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())
}
