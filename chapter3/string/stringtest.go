/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package main

import (
	"fmt"
	"go_programing/chapter3/string/utils"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	var count int
	for range s {
		count++
	}
	fmt.Println("------------------------------")
	fmt.Println(count)
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))
	fmt.Println("------------------------------")

	for i, c := range s {
		fmt.Printf("%d\t%q\t%[2]d\t%[2]x\n", i, c)
	}
	fmt.Println("------------------------------")
	ru := []rune(s)
	fmt.Printf("%x\n", ru)


//	test basename
	var name string = "abd"
	base := utils.Basename(name)
	fmt.Println(base)

//	test comma
	fmt.Println("------------------------------")
	ints := "-123456789.56456"
	fmt.Println(utils.Comma(ints))

//	test sameString
	fmt.Println("------------------------------")
	s1 := "abhijkw啊"
	s2 := "hiw啊kjab"
	fmt.Println(utils.IsSameString(s1, s2))
}
