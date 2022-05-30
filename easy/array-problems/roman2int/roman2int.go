package main

import "fmt"

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if i + 1 == l {
			res += m[string(s[i])]
		}
		if i + 1 < l { 
			if m[string(s[i])] < m[string(s[i+1])] {
				res -= m[string(s[i])]
			} else {
				res += m[string(s[i])]
			}
		}
	}
    return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}