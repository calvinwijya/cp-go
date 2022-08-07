package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {

		if v, found := m[target-num]; found {
			fmt.Println(idx)
			fmt.Println(v)
			return []int{v, idx}
		}

		m[num] = idx

	}
	return nil
}

func main() {
	input := []int{2, 23, 5, 6, 7}
	fmt.Println(twoSum(input, 13))
}
