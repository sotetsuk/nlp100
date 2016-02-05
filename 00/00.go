package main

import "fmt"

func Reverse(s string) string {
	runeS := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i + 1, j - 1 {
		runeS[i], runeS[j] = runeS[j], runeS[i]
	}
	return string(runeS)
}


func main() {
	fmt.Println(Reverse("stressed"))
}