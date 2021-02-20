/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package utils

import (
	"bytes"
)

func Comma(str string) string {
	if len(str) < 3 {
		return str
	}

	var buf bytes.Buffer
	if str[0] == '+' || str[0] == '-' {
		buf.WriteByte(str[0])
		str = str[1:]
	}

	flag := true
	for i, v := range str {
		if v == '.' {
			flag = false
		}
		buf.WriteRune(v)
		if flag && (i + 1) % 3 == 0 && i != len(str) - 1 && i + 1 < len(str) - 1 && str[i + 1] != '.' {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
