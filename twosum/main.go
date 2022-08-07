package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {

		if v, found := m[target-num]; found {
			return []int{v, idx}
		}

		m[num] = idx

	}
	return nil
}

func main() {
	input := []int{2, 7, 11, 15}
	fmt.Println(twoSum(input, 9))
}

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
