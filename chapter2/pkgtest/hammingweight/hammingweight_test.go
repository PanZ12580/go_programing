/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package hammingweight

import "testing"

func BenchmarkHammingWeight1(b *testing.B) {
	/*for i := 0; i < b.N; i++ {
		HammingWeight1(105)
	}*/
	bench(b, 1000)
}


func BenchmarkHammingWeight2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HammingWeight2(105)
	}
}

func BenchmarkHammingWeight3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HammingWeight3(105)
	}
}

func bench(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		HammingWeight1(105)
	}
}