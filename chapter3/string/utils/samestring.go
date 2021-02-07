/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package utils

func IsSameString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	l := len(s1)
	var r byte
	for i := 0; i < l; i++ {
		r ^= s1[i] ^ s2[i]
	}
	return r == 0
}
