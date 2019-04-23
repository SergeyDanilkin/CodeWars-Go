package main

import (
	"fmt"
	"sort"
)

func main() {
	var ls = []int{50, 55, 56, 57, 58}
	max := ChooseBestSum(163, 3, ls)
	fmt.Print(max)
}

func ChooseBestSum(t int, k int, ls []int) int {
	return getBestSum(t, k, ls)
}

//get best sum of array
func getBestSum(max int, cnt int, m []int) int {
	if len(m) < cnt {
		return -1
	}
	var cntCombs = countCombinations(len(m), cnt)
	var sums []int
	for i := 0; i < cntCombs; i++ {
		sums = append(sums, getSumsOfComb(i, cnt, m))
	}
	sort.Ints(sums)

	if sums[0] > max {
		return -1
	}

	for i := 0; i < len(sums); i++ {
		if sums[i] > max && i > 0 {
			return sums[i-1]
		}
		if sums[i] > max && i == 0 {
			return sums[i]
		}
	}

	return sums[len(sums)-1]
}

//get sums of all comb
func getSumsOfComb(index int, k int, ar []int) int {
	var res []int
	res = append(res, 0)
	var n = len(ar)
	var s = 0
	for t := 1; t <= k; t++ {
		var j = res[t-1] + 1
		for {
			if !((j < (n - k + t)) && ((s + countCombinations(n-j, k-t)) <= index)) {
				break
			}
			s += countCombinations(n-j, k-t)
			j++
		}
		res = append(res, j)
	}

	sum := 0
	for i := 1; i < len(res); i++ {
		sum += ar[res[i]-1]
	}
	return sum
}

// get factorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

//get count of combinations
func countCombinations(n int, k int) int {
	return factorial(n) / factorial(k) / factorial(n-k)
}

//get sum of array
func sum(ar []int) int {
	sum := 0
	for i := range ar {
		sum += ar[i]
	}
	return sum
}
