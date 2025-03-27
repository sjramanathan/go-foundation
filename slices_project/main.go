package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	fmt.Println("len", len(s))

	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v \n", s2)
	fmt.Printf("s2 = %v \n", s2)

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v \n", s3)

	s3 = append(s3, 100)
	fmt.Printf("s3 = %#v \n", s3)
	fmt.Printf("s2 = %#v \n", s2)

	fmt.Printf("s2: len=%d, cap=%d \n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d \n", len(s3), cap(s3))

	fmt.Printf("m1: %v\n", median([]float64{1, 4, 5, 4, 3, 2}))
	fmt.Printf("m2: %v\n", median([]float64{1, 2, 3, 4, 4, 3, 2, 1}))
}

func appendInt(slice []int, value int) []int {
	index := len(slice)

	if len(slice) < cap(slice) {
		slice = slice[:len(slice)+1]
	} else {

		fmt.Printf("reallocate: %d -> %d ]n", len(slice), 2*len(slice)+1)
		newSlice := make([]int, len(slice)*2+1)
		copy(newSlice, slice) // shallow copy
		slice = newSlice[:len(slice)+1]
	}

	slice[index] = value
	return slice
}

func concat(s1, s2 []string) []string {
	s3 := make([]string, len(s1)+len(s2), 2*(len(s1)+len(s2))+1)

	copy(s3, s1)
	copy(s3[len(s1):], s2)

	return s3
}

func median(values []float64) float64 {
	sort.Float64s(values)

	if len(values)%2 == 1 {
		return values[len(values)/2]
	}

	return (values[len(values)/2] + values[(len(values)/2)+1]) / 2
}
