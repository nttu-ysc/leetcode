package problem151

import (
	"bytes"
)

func reverseWords(s string) string {
	var ans, tmp string

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(tmp) > 0 {
				ans = tmp + " " + ans
				tmp = ""
			}
		} else {
			tmp += string(s[i])
		}
	}

	if len(tmp) > 0 {
		ans = tmp + " " + ans
	}

	return ans[:len(ans)-1]
}

func reverseWords2(s string) string {
	var ans, tmp bytes.Buffer

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if tmp.Len() > 0 {
				ans = *bytes.NewBufferString(tmp.String() + " " + ans.String())
				tmp.Reset()
			}
		} else {
			tmp.WriteByte(s[i])
		}
	}

	if tmp.Len() > 0 {
		ans = *bytes.NewBufferString(tmp.String() + " " + ans.String())
	}

	return ans.String()[:ans.Len()-1]
}
