/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package function

func Reverse(arr *[5]int) {
	for i, j := 0, len(*arr) - 1; i < j; i, j = i + 1, j - 1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func Rotate(s []int, i int) []int {
	l := len(s)
	res := make([]int, l)
	idx1, idx2 := 0, l - 1
	for right, left := i, i - 1; right < l || left >= 0; right, left = right + 1, left - 1 {
		if right < l {
			res[idx1] = s[right]
			idx1++
		}
		if left >= 0 {
			res[idx2] = s[left]
			idx2--
		}
	}

	return res
}

func EliminateRepeat(s *[]string) {
	idx := 1
	flag := (*s)[0]
	for i := 1; i < len(*s); i++ {
		if (*s)[i] == flag {
			continue
		} else {
			(*s)[idx] = (*s)[i]
			idx++
			flag = (*s)[i]
		}
	}
	*s = (*s)[:idx]
}

func EliminateRepeatSpace(b *[]byte) {
	idx := 1
	isSpace := func(b byte) bool {
		switch b {
		case '\t', '\n', '\v', '\f', '\r', ' ':
			return true
		}
		return false
	}

	flag := false
	for i := 1; i < len(*b); i++ {
		if isSpace((*b)[i]) {
			flag = true
			continue
		} else {
			if flag {
				idx++
				flag = false
			}
			(*b)[idx] = (*b)[i]
			idx++
		}
	}
	*b = (*b)[:idx]
}

func ReverseByteArr(b *[]byte) {
	for i, j := 0, len(*b) - 1; i < j; i, j = i + 1, j - 1 {
		(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
	}
}