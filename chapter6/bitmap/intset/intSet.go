/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	bucket, idx := x / 64, uint(x % 64)
	for bucket >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[bucket] |= 1 << idx
}

func (s *IntSet) Contains(x int) bool {
	bucket, idx := x / 64, uint(x % 64)
	return bucket < len(s.words) && (s.words[bucket] & (1 << idx)) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, bucket := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, bucket)
		} else {
			s.words[i] |= bucket
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteRune('{')
	for i, bucket := range s.words {
		if bucket == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if s.words[i] & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteRune(' ')
				}
				fmt.Fprintf(&buf, "%d", 64 * i + j)
			}
		}
	}
	buf.WriteRune('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, bucket := range s.words {
		for j := 0; j < 64; j++ {
			if bucket & (1 << uint(j)) != 0 {
				count++
			}
		}
	}

	return count
}

func (s *IntSet) Remove(x int) {
	if s.Contains(x) {
		bucket, idx := x / 64, uint(x % 64)
		s.words[bucket] &= ^(1 << idx)
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() IntSet {
	cp := IntSet{}
	for i, bucket := range s.words {
		if i >= len(cp.words) {
			cp.words = append(cp.words, 0)
		}
		cp.words[i] = bucket
	}
	return cp
}

func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		bucket, idx := n / 64, uint(n % 64)
		for bucket >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[bucket] |= 1 << idx
	}
}

func (s *IntSet) IntersectWith(t *IntSet) IntSet {
	res := IntSet{}
	for i, bucket := range s.words {
		if i >= len(t.words) {
			break
		}
		res.words = append(res.words, t.words[i] & bucket)
	}
	return res
}

func (s *IntSet) DifferenceWith(t *IntSet) IntSet {
	res := IntSet{}
	for i, bucket := range s.words {
		if i >= len(t.words) {
			res.words = append(res.words, bucket)
		} else {
			res.words = append(res.words, bucket & (bucket ^ t.words[i]))
		}
	}
	return res
}

func (s *IntSet) SymmetricDifference(t *IntSet) IntSet {
	res := IntSet{}
	for i, bucket := range s.words {
		if i >= len(t.words) {
			res.words = append(res.words, bucket)
		} else {
			res.words = append(res.words, bucket ^ t.words[i])
		}
	}
	return res
}

func (s *IntSet) Elems() []int {
	var res []int
	for i, bucket := range s.words {
		if bucket == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if bucket & (1 << uint(j)) != 0 {
				res = append(res, i * 64 + j)
			}
		}
	}
	return res
}


