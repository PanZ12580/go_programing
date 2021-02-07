/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package format

import (
	"strings"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	tests := []struct {
		value interface{}
		want string
	}{
		{1 * time.Nanosecond, "1"},
		{1, "1"},
		{[]int64{1, 2}, "[]int64 0x"},
		{[2]int{1, 2}, "[2]intvalue"},
	}

	for _, test := range tests {
		if got := Any(test.value); got != test.want && !strings.Contains(got, test.want) {
			t.Logf("value: %q, got %q, want %q\n", test.value, got, test.want)
		}
	}
}
