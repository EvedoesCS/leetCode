package main

import (
    "fmt"
)

// Given an array of nums return true if any value appears at least twice in the array and return
// false if every element is distinct; 

// Checks for duplicate values in an int splice
func ContainsDuplicates(nums []int) (bool) {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
    }
    return false
}

func main() {
    nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
   
    fmt.Printf("[]int := %v\n", nums)
   fmt.Println(ContainsDuplicates(nums))
}
