package search

import (
	"fmt"
	"strings"
)

//Search looks for indexes of proverbs that contains the word w. It returns a slice with indexes of prverbs.
func Search(w string, s []string) []int {
	word := strings.ToLower(w)
	ans := []int{}
	for i, v := range s {
		v = strings.ToLower(v)
		if strings.Contains(v, word) {
			ans = append(ans, i)
		}
	}
	return ans
}

//PrintAll prints all proverbs within the file.
func PrintAll(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
