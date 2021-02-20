/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package hammingweight

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

func HammingWeight1(x uint64) int {
	/*return int(pc[byte(x >> (0 * 8))] +
		pc[byte(x >> (1 * 8))] +
		pc[byte(x >> (2 * 8))] +
		pc[byte(x >> (3 * 8))] +
		pc[byte(x >> (4 * 8))] +
		pc[byte(x >> (5 * 8))] +
		pc[byte(x >> (6 * 8))] +
		pc[byte(x >> (7 * 8))])*/
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x >> (i * 8))])
	}
	return count
}

func HammingWeight2(x int) int {
	var count int
	for x != 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func HammingWeight3(x int) int {
	var count int
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
