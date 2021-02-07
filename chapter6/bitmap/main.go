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
	"go_programing/chapter6/bitmap/intset"
)

func main() {
	var set intset.IntSet
	set.Add(1)
	set.Add(0)
	set.Add(9)
	set.Add(64)
	set.Add(145)
	var t intset.IntSet
	t.Add(206)
	t.Add(306)
	t.Add(406)
	fmt.Println(set.Contains(9))
	fmt.Println(&set)
	fmt.Println(&t)
	set.UnionWith(&t)
	fmt.Println(&set)
	fmt.Println(set.Len())
	set.Remove(206)
	fmt.Println(&set)
	fmt.Println(set.Len())
	cp := set.Copy()
	set.Clear()
	fmt.Println(&set)
	fmt.Println(set.Len())
	fmt.Println(&cp)
	cp.AddAll(12, 13, 14, 15, 16)
	fmt.Println(&cp)
	i := cp.IntersectWith(&t)
	fmt.Println(&i)
	differ := cp.DifferenceWith(&t)
	fmt.Println(&differ)
	sd := cp.SymmetricDifference(&t)
	fmt.Println(&sd)
	for _, v := range cp.Elems() {
		fmt.Printf("%d\t", v)
	}
	fmt.Printf("\n%d\n", 32 << (^uint(0) >> 63))
}
