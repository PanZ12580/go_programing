/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package intset

import (
	"math/rand"
	"testing"
)

var (
	testMap map[int]bool
	set IntSet
	testInts = []int{1, 45, 8, 22, 22, 0, 8, 15, 69}
)

func init() {
	testMap = make(map[int]bool)
	for _, i := range testInts {
		testMap[i] = true
		set.Add(i)
	}
}

func TestIntSet_Add_And_Contains_And_Len(t *testing.T) {
	if l := len(testMap); l != set.Len() {
		t.Errorf("set.Len() != %d", l)
	}
	for _, i := range testInts {
		if !set.Contains(i) {
			t.Errorf("set.Contains(%d) = false", i)
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	union := []int{4, 5, 6, 10, 11}
	var newSet IntSet
	for _, i := range union {
		newSet.Add(i)
		testMap[i] = true
	}
	set.UnionWith(&newSet)

	if l := len(testMap); l != set.Len() {
		t.Errorf("set.Len() != %d", l)
	}

	newSlice := make([]int, 0)
	newSlice = append(newSlice, testInts...)
	newSlice = append(newSlice, union...)

	for _, n := range newSlice {
		if !set.Contains(n) {
			t.Errorf("set.Contains(%d) == false", n)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	remove := []int{22, 0, 8, 15, 600}
	for _, rm := range remove {
		set.Remove(rm)
		testMap[rm] = false
	}
	var valid int
	for _, v := range testMap {
		if v {
			valid++
		}
	}

	if valid != set.Len() {
		t.Errorf("set.Len() != %d", valid)
	}

	for _, rm := range remove {
		if set.Contains(rm) {
			t.Errorf("set.Contains(%d) == true", rm)
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	ints := []int{10, 8, 0, 56, 77, 2, 4, 5, 88}
	var differ IntSet
	differMap := make(map[int]bool)
	for _, i := range ints {
		differMap[i] = true
		differ.Add(i)
	}

	res := set.DifferenceWith(&differ)
	newMap := make(map[int]bool)

	for k, v := range testMap {
		if v && !differMap[k] {
			newMap[k] = true
		}
	}

	if l := len(newMap); l != res.Len() {
		t.Errorf("set.Len() != %d", l)
	}

	for k, v := range newMap {
		if v && !res.Contains(k) {
			t.Errorf("the differSet didn't contains the expected value: %d", k)
		}
	}
}

func BenchmarkIntSet_Add(b *testing.B) {
	b.StopTimer()
	var set IntSet
	rng := rand.New(rand.NewSource(1))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(1000))
	}
}

func BenchmarkIntSet_Add_Map(b *testing.B) {
	b.StopTimer()
	m := make(map[int]bool)
	rng := rand.New(rand.NewSource(1))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m[rng.Intn(1000)] = true
	}
}



