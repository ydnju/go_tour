package basictypes

import (
	"bytes"
	"strings"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func NoRecursiveComma(s string) string {
	var buf bytes.Buffer
	div := len(s) % 3
	buf.WriteString(s[:div])
	for ; div < len(s); div += 3 {
		buf.WriteString(",")
		buf.WriteString(s[div : div+3])
	}
	return buf.String()
}
