// Package main
package main

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	var sum int
	for _, val := range numbers {
		sum = sum + val
	}
	return sum
}

func SumAll(numberList ...[]int) []int {
	sumList := make([]int, 0, len(numberList))
	for _, v := range numberList {
		sumList = append(sumList, Sum(v))
	}
	return sumList
}

func SumAllTails(numberList ...[]int) []int {
	sumList := make([]int, 0, len(numberList))
	for _, v := range numberList {
		length := len(v)
		if length == 0 {
			sumList = append(sumList, 0)
			continue
		}
		tail := v[1:]
		sumList = append(sumList, Sum(tail))
	}
	return sumList
}

func SumAllHeads(numberList ...[]int) []int {
	sumList := make([]int, 0, len(numberList))
	for _, v := range numberList {
		length := len(v)
		if length == 0 {
			sumList = append(sumList, 0)
			continue
		}
		head := v[:len(v)-1]
		sumList = append(sumList, Sum(head))
	}
	return sumList
}
