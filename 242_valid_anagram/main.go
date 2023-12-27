package main

import (
    "fmt"
)

// Given two strings s and t return true if t is an anagram of s and false otherwise

// Sum up the unicode values of a given string
func SumUnicode(str string) (int) {
    var sum int = 0
    runeOfStr := []rune(str)

    for _, val := range runeOfStr {
        sum += int(val)
    }
    return sum
}

// compare the sum of unicode characters and return true if they are the same
// or return false if not;
func IsAnagram(s string, t string) (bool) {
    sVal := SumUnicode(s)
   tVal := SumUnicode(t)

   if sVal == tVal {
       return true
   } else {
       return false
   }
}

func main() {
    s := "rat"
    t := "car"

    fmt.Println(IsAnagram(s, t))
}
