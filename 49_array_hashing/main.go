package main

import (
    "fmt"
)

// Given an array of strings 'strs', grousp the anagrams together. Return in any order;

// func totals and the values of each rune in a strings as a unicode character
// value is returned as an int;
func sumUnicode(s string) (int) {
    iterator := []rune(s)
    total := 0

    for _, val := range iterator {
        total += int(val)
    }
    return total
}


// func compares the unicode values of two strings and returns a bool
func compareStrings(s string, t string) (bool) {
    sVal := sumUnicode(s)
    tVal := sumUnicode(t)

    if sVal == tVal {
        return true
    } 
    return false
}

// func groups a []string into groups of anagrams
func groupAnagrams(strs []string) [][]string {
    groups := [][]string{}
    
    for i := 0; i < len(strs); i++ {
        fmt.Println(strs)
        if strs[i] == "NIL" {
            continue
        }
        anagramGroup := []string{}
        anagramGroup = append(anagramGroup, strs[i])
        for j := i + 1; j < len(strs) - i; j++ {
            if strs[j] == "NIL" {
                continue
            } else if compareStrings(strs[i], strs[j]) == true {
                anagramGroup = append(anagramGroup, strs[j])
                strs[j] = "NIL"
            }
        }
        strs[i] = "NIL"
        groups = append(groups, anagramGroup)
    }
    return groups
}

func main() {
    strs := []string {"eat", "tea", "tan", "ate", "nat", "bat"}
    fmt.Println(groupAnagrams(strs))
}
