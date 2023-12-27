package main

import (
    "fmt"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such
// that they add up to target.
// You may assume that each input would have exactly one solution and you may not use the same element 
// twice;

func TwoSum(nums []int, target int) []int {
    twoSum := []int{0, 0}
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                twoSum[0] = i
                twoSum[1] = j
                return twoSum
            }
        }
    }
    return twoSum
}

func main() {
    nums := []int{3, 3}
    target := 6

    fmt.Println(TwoSum(nums, target))
}
