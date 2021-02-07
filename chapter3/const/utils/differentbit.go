/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package utils

func Differ(s1 *[32]byte, s2 *[32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		count += popCount(s1[i] ^ s2[i])
	}
	return count
}

func popCount(x uint8) int {
	var res int
	for x != 0 {
		x &= x - 1
		res++
	}
	return res
}


