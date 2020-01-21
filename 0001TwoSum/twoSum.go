package main

import "fmt"

func twoSumV2(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for index, item := range nums {
		if j, ok := sumMap[target - item]; ok {
			return []int{j, index}
		}

		sumMap[item] = index
	}
	return nil

}

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for index, item := range nums {
		sumMap[item] = index
	}
	for index, item := range nums {
		j, ok := sumMap[target - item]
		if ok && index != j {
			return []int{index, j}
		}
	}
	return nil

}
// 3-0 2-1 4-2

func main()  {
	num := []int{3,3}
	fmt.Println(twoSum(num,6))
}

