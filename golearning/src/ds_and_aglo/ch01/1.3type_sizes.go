package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt8 := math.MaxInt8 // number of bytes contained in maxInt8
	minInt8 := math.MinInt8
	fmt.Println("maxInt8 :: ", maxInt8)
	fmt.Println("minInt8 :: ", minInt8)

	maxInt16 := math.MaxInt16 // number of bytes contained in maxInt16
	minInt16 := math.MinInt16
	fmt.Println("maxInt16 :: ", maxInt16)
	fmt.Println("minInt16 :: ", minInt16)

	maxInt32 := math.MaxInt32 // number of bytes contained in maxInt32
	minInt32 := math.MinInt32
	fmt.Println("maxInt32 :: ", maxInt32)
	fmt.Println("minInt32 :: ", minInt32)

	maxInt64 := math.MaxInt64 // number of bytes contained in maxInt64
	minInt64 := math.MinInt64
	fmt.Println("maxInt64 :: ", maxInt64)
	fmt.Println("minInt64 :: ", minInt64)

	maxUint64 := math.MaxUint64 // number of bytes contained in maxUint64
	fmt.Println("maxUint64 :: ", maxUint64)

	maxUint32 := math.MaxUint32 // number of bytes contained in maxUint32
	fmt.Println("maxUint32 :: ", maxUint32)

	maxUint16 := math.MaxUint16 // number of bytes contained in maxUint16
	fmt.Println("maxUint16 :: ", maxUint16)

}
