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
	"go_programing/chapter2/pkgtest/hammingweight"
	_ "go_programing/chapter2/tempconv"
)

func main() {
/*	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Printf("CToF: %g -> %g\n", tempconv.BoilingC, tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))*/

	weight := hammingweight.HammingWeight1(951)
	fmt.Println(weight)
	fmt.Println(hammingweight.HammingWeight2(951))
	fmt.Println(hammingweight.HammingWeight3(951))

}
