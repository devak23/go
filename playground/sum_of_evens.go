package main

import (
	"fmt"
	"github.com/DylanMeeus/hasgo/types"
	"math"
)

func SumOfEvenMain() {
	var sum int64

	sum = SumOfEvensImperative()
	fmt.Printf("Sum of evens imperative way = %d\n", sum)
	sum = SumOfEvensFunctional()
	fmt.Printf("Sum of events functional way = %d\n", sum)
}

func SumOfEvensImperative() int64 {
	sum := int64(0)
	for i := -10; i <= 10; i++ {
		x := int(math.Abs(float64(i)))
		if x%2 == 0 {
			sum += int64(i)
		}
	}

	return sum
}

func SumOfEvensFunctional() int64 {
	isEven := func(i int64) bool {
		return i%2 == 0
	}
	return types.IntRange(-10, 10).Abs().Filter(isEven).Sum()
}
