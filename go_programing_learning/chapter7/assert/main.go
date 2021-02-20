package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(sqlQuote(time.Hour))
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case float64, float32:
		return fmt.Sprintf("%f", x)
	case bool:
		if x {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type: %T: %[1]v\n", x))
	}
}