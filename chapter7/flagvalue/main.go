package main

import (
	"flag"
	"fmt"
	"go_programing/chapter2/tempconv"
)

type celsiusFlag struct {
	tempconv.Celsius
}

func (cf *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		cf.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		cf.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature: %s\n", s)
}

func CelsiusFlag(name string, defaultValue tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{defaultValue}
	flag.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}
