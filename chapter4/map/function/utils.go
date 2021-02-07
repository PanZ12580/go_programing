package function

import (
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"
)

func InputCharCount() {
	countMap := make(map[rune]int)
	differMap := make(map[string]int)
	var utfLenCount [utf8.UTFMax + 1]int
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		log.Fatalf("error happend: %q\n", err)
	}
	rArr := []rune(s)
	for _, v := range rArr {
		countMap[v]++
		utfLenCount[utf8.RuneLen(v)]++
		if unicode.IsLetter(v) {
			differMap["letter"]++
		} else if unicode.IsDigit(v) {
			differMap["digit"]++
		} else if unicode.IsSpace(v) {
			differMap["space"]++
		} else {
			differMap["other"]++
		}
	}
	fmt.Printf("char\tcount\n")
	for k, v := range countMap {
		fmt.Printf("%-4c\t%d\n", k, v)
	}
	fmt.Printf("length\tcount\n")
	for i, c := range utfLenCount {
		if i > 0 {
			fmt.Printf("%-4d\t%d\n", i, c)
		}
	}
	fmt.Printf("type\tcount\n")
	for k, v := range differMap {
		fmt.Printf("%q\t%d\n", k, v)
	}
}
