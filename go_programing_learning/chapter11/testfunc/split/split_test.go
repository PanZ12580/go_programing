/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input, seq string
		want int
	}{
		{"a,b,c,d", ",", 4},
		{"a:b:c", ":", 3},
		{"fa fdf 45 dsff 4511 adsf", " ", 6},
		{"45a512a78a112ada", "a", 6},
	}

	for _, test := range tests {
		if got := len(strings.Split(test.input, test.seq)); got != test.want {
			t.Errorf("len(strings.Split(%q, %q)) = %d\twant: %d", test.input, test.seq, got, test.want)
		}
	}
}
